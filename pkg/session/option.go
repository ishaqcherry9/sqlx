package session

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/xoctopus/x/misc/must"
	"github.com/xoctopus/x/textx"
)

type EndpointOption struct {
	AutoMigration   bool `url:"-"`
	DryRun          bool `url:"-"`
	CreateTableOnly bool `url:"-"`

	MultiStatements bool

	MaxOpen     int           `url:"-"`
	MaxIdle     int           `url:"-"`
	MaxLifetime time.Duration `url:"-"`
	MaxIdleTime time.Duration `url:"-"`
}

func (o *EndpointOption) SetDefault() {
	must.NoError(textx.UnmarshalURL(url.Values{}, o))
	o.MultiStatements = true
	o.MaxOpen = 100
	o.MaxIdle = 50
	o.MaxLifetime = time.Hour
	o.MaxIdleTime = time.Hour
}

func (o *EndpointOption) Apply(db *sql.DB) {
	db.SetMaxOpenConns(o.MaxOpen)
	db.SetMaxIdleConns(o.MaxIdle)
	db.SetConnMaxLifetime(o.MaxLifetime)
	db.SetConnMaxIdleTime(o.MaxIdleTime)
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
