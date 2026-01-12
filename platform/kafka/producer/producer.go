package producer

import (
	"context"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

type producer struct {
	logger       Logger
	topic        string
	syncProducer sarama.SyncProducer
}

func NewProducer(syncProducer sarama.SyncProducer, topic string, logger Logger) *producer {
	return &producer{
		logger:       logger,
		topic:        topic,
		syncProducer: syncProducer,
	}
}

func (p *producer) Send(ctx context.Context, key, value []byte) error {
	partition, offset, err := p.syncProducer.SendMessage(&sarama.ProducerMessage{
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
		Topic: p.topic,
	})
	if err != nil {
		p.logger.Error(ctx, "Failed to send message", zap.Error(err))
		return err
	}

	p.logger.Info(ctx, "Message sent",
		zap.String("topic", p.topic),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
		zap.String("key", string(key)),
		zap.String("value", string(value)),
	)

	return nil
}
