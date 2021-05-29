package apigatewaylambdagoexample

import "context"

// Service implementation.
type Service struct {
	queue Queue
}

// NewService creates a service layer.
func NewService(queue Queue) *Service {
	return &Service{
		queue: queue,
	}
}

// CreateMessage creates a message in the queue
func (service *Service) CreateMessage(ctx context.Context, qURL string, content []byte) error {
	return service.queue.sendMessage(ctx, qURL, content)
}
