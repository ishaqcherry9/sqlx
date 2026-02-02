package modeled

import (
	"github.com/ishaqcherry9/sqlx/internal"
	"github.com/ishaqcherry9/sqlx/pkg/builder"
)

type Newer[M internal.Model] struct{}

func (m *Newer[M]) Model() *M {
	return new(M)
}

type (
	Model               = internal.Model
	ModelNewer[M Model] internal.ModelNewer[M]
)

type OrderAddition[M Model] interface {
	builder.OrderAddition
}
