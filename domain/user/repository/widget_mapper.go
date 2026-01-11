package repository

import "context"

type WidgetMapper interface {
	CreateDefaultWidgets(ctx context.Context, userID, username string) error
}
