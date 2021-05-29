#!/bin/bash
# Script to create sqs queue

SQS_QUEUE_NAME=test-queue # Sqs queue name

awslocal sqs create-queue \
    --queue-name $SQS_QUEUE_NAME \
    && echo "Created" || echo "Failed to create"

echo "Sqs initialization completed"
