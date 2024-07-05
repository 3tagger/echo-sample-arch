package seeder

import (
	"context"
	"math"
)

// SeedBatchSize is used to determine the number of instances created
// for each batches
const SeedBatchSize int = 100

// SeederExecutor is the logic on how to perform the seeding
type SeederExecutor interface {
	Run(ctx context.Context, n int) error
}

type seederExecutor[K comparable, V SeederService[K]] struct {
	seederService SeederService[K]
}

func NewSeederExecutor[K comparable, V SeederService[K]](seederService V) SeederExecutor {
	return &seederExecutor[K, V]{
		seederService: seederService,
	}
}

// SeederService is how the fake data is generated and inserted
type SeederService[K comparable] interface {
	Fake(ctx context.Context) K
	InsertMany(ctx context.Context, arr []K) error
}

func (e *seederExecutor[K, V]) Run(ctx context.Context, n int) error {
	for n > 0 {
		j := int(math.Min(float64(n), float64(SeedBatchSize)))

		arr := make([]K, j)

		for i := 0; i < j; i++ {
			u := e.seederService.Fake(ctx)

			arr[i] = u
		}
		if err := e.seederService.InsertMany(ctx, arr); err != nil {
			return err
		}

		n -= SeedBatchSize
	}
	return nil
}
