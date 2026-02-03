package session

import (
	"net/url"
	"time"

	"github.com/xoctopus/confx/pkg/types"
	"github.com/xoctopus/x/misc/must"
	"github.com/xoctopus/x/textx"

	"github.com/ishaqcherry9/sqlx/internal/sql/adaptor"
)

type EndpointOption struct {
	AutoMigration   bool `url:"-"`
	DryRun          bool `url:"-"`
	CreateTableOnly bool `url:"-"`

	MultiStatements bool

	MaxOpen     int            `url:"-"`
	MaxIdle     int            `url:"-"`
	MaxLifetime types.Duration `url:"-"`
	MaxIdleTime types.Duration `url:"-"`
}

func (o *EndpointOption) SetDefault() {
	must.NoError(textx.UnmarshalURL(url.Values{}, o))
	o.MultiStatements = true
	o.MaxOpen = 100
	o.MaxIdle = 50
	o.MaxLifetime = types.Duration(time.Hour)
	o.MaxIdleTime = types.Duration(time.Hour)
}

func (o *EndpointOption) Apply(d adaptor.DB) {
	if db, ok := d.(adaptor.D); ok {
		db.D().SetMaxOpenConns(o.MaxOpen)
		db.D().SetMaxIdleConns(o.MaxIdle)
		db.D().SetConnMaxLifetime(time.Duration(o.MaxLifetime))
		db.D().SetConnMaxIdleTime(time.Duration(o.MaxIdleTime))
	}
}

type AdaptorOption struct {
	ReadOnly bool
}

type OptionFunc func(*AdaptorOption)

func ReadOnly() OptionFunc {
	return func(o *AdaptorOption) {
		o.ReadOnly = true
	}
}
