package kafka

import (
	"context"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
)

type Consumer interface {
	Consume(ctx context.Context, handler consumer.MessageHandler) error
}

type Producer interface {
	Send(ctx context.Context, key, value []byte) error
}
