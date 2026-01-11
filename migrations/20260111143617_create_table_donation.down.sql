BEGIN;

DROP TABLE IF EXISTS webhook_events;
DROP TABLE IF EXISTS transaction_histories;
DROP TABLE IF EXISTS donations;
DROP TYPE transaction_category;
DROP TYPE transaction_type;
DROP TYPE payment_method;
DROP TYPE donation_status;

COMMIT;