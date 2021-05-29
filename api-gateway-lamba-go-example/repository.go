package apigatewaylambdagoexample

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Queue interface
type Queue interface {
	sendMessage(ctx context.Context, qURL string, message []byte) error
}

// Sqs is a connection with an AWS SQS service
type Sqs struct {
	sqs *sqs.SQS
}

func (s Sqs) sendMessage(ctx context.Context, qURL string, message []byte) error {
	_, err := s.sqs.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageBody:  aws.String(string(message)),
		QueueUrl:     &qURL,
	})

	if err != nil {
		return err
	}

	return nil
}

// NewSqs will connect to an SQS in a certain AWS Region.
func NewSqs(region, endpoint string) (*Sqs, error) {
	config := createAwsConfig(region, endpoint)

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	return &Sqs{sqs: sqs.New(sess)}, nil
}

func createAwsConfig(region, endpoint string) *aws.Config {
	config := aws.NewConfig()

	if endpoint != "" {
		config.WithEndpoint(endpoint)
	}
	if region != "" {
		config.WithRegion(region)
	}
	return config
}
