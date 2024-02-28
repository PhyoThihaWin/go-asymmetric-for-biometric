package biometric

import "context"

type Repository interface {
	CreateBiometric(ctx context.Context, id string) error
}
