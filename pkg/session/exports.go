package session

import (
	"github.com/ishaqcherry9/sqlx/internal/sql/adaptor"
	"github.com/ishaqcherry9/sqlx/internal/sql/scanner"
)

var (
	Scan = scanner.Scan
	Open = adaptor.Open
)

type (
	Adaptor = adaptor.Adaptor
)
