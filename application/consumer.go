package application

import "context"

type Consumer interface {
    Consume(ctx context.Context) error
}
