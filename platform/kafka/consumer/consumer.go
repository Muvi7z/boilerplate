package consumer

import (
	"context"
	"errors"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

type consumer struct {
	logger     Logger
	group      sarama.ConsumerGroup
	topics     []string
	middleware []Middleware
}

func NewConsumer(group sarama.ConsumerGroup, topics []string, logger Logger, middleware ...Middleware) *consumer {
	return &consumer{
		group:      group,
		topics:     topics,
		middleware: middleware,
		logger:     logger,
	}
}

func (c *consumer) Consume(ctx context.Context, handler MessageHandler) error {
	newGroupHandler := NewGroupHandler(handler, c.logger, c.middleware...)

	for {
		if err := c.group.Consume(ctx, c.topics, newGroupHandler); err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}

			c.logger.Error(ctx, "Kafka consumer error", zap.Error(err))
			return err
		}
	}
}
