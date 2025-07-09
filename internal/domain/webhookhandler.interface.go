package domain

import (
	"context"
)

type WebhookHandler interface {
	ProcessWebhook(ctx context.Context, payload []byte) error
}
