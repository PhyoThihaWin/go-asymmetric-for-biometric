package biometric

import (
	"context"
)

type UseCase interface {
	CreateBiometric(ctx context.Context, id string) error
}
