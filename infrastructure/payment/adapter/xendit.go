package adapter

import (
	"context"
	"fmt"

	"github.com/arvinpaundra/sesen-api/config"
	"github.com/arvinpaundra/sesen-api/domain/donation/constant"
	"github.com/arvinpaundra/sesen-api/domain/payment/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/payment/repository"
	"github.com/xendit/xendit-go/v7"
	"github.com/xendit/xendit-go/v7/payment_request"
	"go.uber.org/zap"
)

var _ repository.PaymentGatewayMapper = (*XenditPaymentAdapter)(nil)

type XenditPaymentAdapter struct {
	client *xendit.APIClient
	logger *zap.Logger
}

func NewXenditPaymentAdapter(secret string) *XenditPaymentAdapter {
	env := config.GetString("APP_ENV")
	if env == "" {
		env = "development"
	}

	var logger *zap.Logger
	if env == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	return &XenditPaymentAdapter{
		client: xendit.NewClient(secret),
		logger: logger,
	}
}

// Pay creates a payment request via Xendit for e-wallet or QRIS
func (x *XenditPaymentAdapter) Pay(ctx context.Context, payload request.PaymentRequestPayload) (string, error) {
	method := constant.PaymentMethod(payload.Method)

	switch method {
	case constant.PaymentMethodQris:
		return x.createQRISPayment(ctx, payload)
	case constant.PaymentMethodShopeepay,
		constant.PaymentMethodDana, constant.PaymentMethodLinkAja:
		return x.createEWalletPayment(ctx, payload)
	default:
		err := fmt.Errorf("unsupported payment method: %s", payload.Method)
		x.logger.Error(fmt.Sprintf("Failed to create payment: %v", err))
		return "", err
	}
}

// createEWalletPayment creates an e-wallet charge via Xendit
func (x *XenditPaymentAdapter) createEWalletPayment(ctx context.Context, payload request.PaymentRequestPayload) (string, error) {
	channelCode := x.mapPaymentMethodToEWalletChannel(payload.Method)

	ewalletParams := payment_request.NewEWalletParameters()
	ewalletParams.ChannelCode = &channelCode
	ewalletParams.ChannelProperties = &payment_request.EWalletChannelProperties{
		SuccessReturnUrl: &payload.SuccessRedirectURL,
		FailureReturnUrl: &payload.FailureRedirectURL,
	}

	amount := float64(payload.Amount)

	paymentMethod := payment_request.NewPaymentMethodParameters(
		payment_request.PAYMENTMETHODTYPE_EWALLET,
		payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE,
	)
	paymentMethod.Ewallet = *payment_request.NewNullableEWalletParameters(ewalletParams)

	paymentReq := payment_request.NewPaymentRequestParameters(payment_request.PAYMENTREQUESTCURRENCY_IDR)
	paymentReq.Amount = &amount
	paymentReq.ReferenceId = &payload.ReferenceID
	paymentReq.PaymentMethod = paymentMethod

	if payload.Description != nil {
		paymentReq.Description = *payment_request.NewNullableString(payload.Description)
	}

	resp, httpResp, err := x.client.PaymentRequestApi.CreatePaymentRequest(ctx).
		PaymentRequestParameters(*paymentReq).
		Execute()

	if err != nil {
		x.logger.Error(fmt.Sprintf("Failed to create e-wallet payment for reference %s: %v", payload.ReferenceID, err))
		if httpResp != nil {
			x.logger.Error(fmt.Sprintf("HTTP Status: %d", httpResp.StatusCode))
		}
		return "", fmt.Errorf("failed to create e-wallet payment: %w", err)
	}

	if len(resp.Actions) == 0 {
		err := fmt.Errorf("no payment actions returned from Xendit")
		x.logger.Error(fmt.Sprintf("Failed to create e-wallet payment for reference %s: %v", payload.ReferenceID, err))
		return "", err
	}

	// Return the payment URL/action URL for customer to complete payment
	paymentURL := ""
	for _, action := range resp.Actions {
		if action.UrlType == "DEEPLINK" && action.Url.IsSet() {
			paymentURL = *action.Url.Get()
			break
		}
		if action.Url.IsSet() {
			paymentURL = *action.Url.Get()
		}
	}

	if paymentURL == "" {
		err := fmt.Errorf("no payment URL returned from Xendit")
		x.logger.Error(fmt.Sprintf("Failed to create e-wallet payment for reference %s: %v", payload.ReferenceID, err))
		return "", err
	}

	return paymentURL, nil
}

// createQRISPayment creates a QRIS charge via Xendit
func (x *XenditPaymentAdapter) createQRISPayment(ctx context.Context, payload request.PaymentRequestPayload) (string, error) {
	qrCodeChannelCode := payment_request.QRCODECHANNELCODE_QRIS

	qrCodeParams := payment_request.NewQRCodeParameters()
	qrCodeParams.ChannelCode = *payment_request.NewNullableQRCodeChannelCode(&qrCodeChannelCode)

	amount := float64(payload.Amount)

	paymentMethod := payment_request.NewPaymentMethodParameters(
		payment_request.PAYMENTMETHODTYPE_QR_CODE,
		payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE,
	)
	paymentMethod.QrCode = *payment_request.NewNullableQRCodeParameters(qrCodeParams)

	paymentReq := payment_request.NewPaymentRequestParameters(payment_request.PAYMENTREQUESTCURRENCY_IDR)
	paymentReq.Amount = &amount
	paymentReq.ReferenceId = &payload.ReferenceID
	paymentReq.PaymentMethod = paymentMethod

	if payload.Description != nil {
		paymentReq.Description = *payment_request.NewNullableString(payload.Description)
	}

	resp, httpResp, err := x.client.PaymentRequestApi.CreatePaymentRequest(ctx).
		PaymentRequestParameters(*paymentReq).
		Execute()

	if err != nil {
		x.logger.Error(fmt.Sprintf("Failed to create QRIS payment for reference %s: %v", payload.ReferenceID, err))
		if httpResp != nil {
			x.logger.Error(fmt.Sprintf("HTTP Status: %d", httpResp.StatusCode))
		}
		return "", fmt.Errorf("failed to create QRIS payment: %w", err)
	}

	if len(resp.Actions) == 0 {
		err := fmt.Errorf("no payment actions returned from Xendit")
		x.logger.Error(fmt.Sprintf("Failed to create QRIS payment for reference %s: %v", payload.ReferenceID, err))
		return "", err
	}

	// Return the QR code string from actions
	qrCodeURL := ""
	for _, action := range resp.Actions {
		// For QRIS, the QR string is typically in the action URL
		if action.Url.IsSet() {
			qrCodeURL = *action.Url.Get()
			break
		}
	}

	if qrCodeURL == "" {
		err := fmt.Errorf("no QR code returned from Xendit")
		x.logger.Error(fmt.Sprintf("Failed to create QRIS payment for reference %s: %v", payload.ReferenceID, err))
		return "", err
	}

	return qrCodeURL, nil
}

// mapPaymentMethodToEWalletChannel maps donation payment method constants to Xendit e-wallet channel codes
func (x *XenditPaymentAdapter) mapPaymentMethodToEWalletChannel(method string) payment_request.EWalletChannelCode {
	switch constant.PaymentMethod(method) {
	case constant.PaymentMethodShopeepay:
		return payment_request.EWALLETCHANNELCODE_SHOPEEPAY
	case constant.PaymentMethodDana:
		return payment_request.EWALLETCHANNELCODE_DANA
	case constant.PaymentMethodLinkAja:
		return payment_request.EWALLETCHANNELCODE_LINKAJA
	default:
		// Unsupported e-wallet method
		return payment_request.EWALLETCHANNELCODE_DANA
	}
}
