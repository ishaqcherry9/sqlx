package types

import (
	"database/sql/driver"
	"time"

	"github.com/ishaqcherry9/sqlx/pkg/types/sqltime"
)

type CreationDatetime struct {
	// CreatedAt 创建日期时间(秒)
	CreatedAt sqltime.Datetime `db:"f_created_at,default=CURRENT_TIMESTAMP" json:"createdAt"`
}

func (c *CreationDatetime) MarkCreatedAt() {
	if c.CreatedAt.IsZero() {
		c.CreatedAt.Time = time.Now()
	}
}

type CreationModificationDatetime struct {
	CreationDatetime
	// UpdatedAt 更新日期时间(秒)
	UpdatedAt sqltime.Datetime `db:"f_updated_at,default=CURRENT_TIMESTAMP,onupdate=CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (cu *CreationModificationDatetime) MarkCreatedAt() {
	cu.MarkModifiedAt()
	if cu.CreatedAt.IsZero() {
		cu.CreatedAt = cu.UpdatedAt
	}
}

func (cu *CreationModificationDatetime) MarkModifiedAt() {
	if cu.UpdatedAt.IsZero() {
		cu.UpdatedAt.Time = time.Now()
	}
}

type CreationModificationDeletionDatetime struct {
	CreationModificationDatetime
	// DeletedAt 删除日期时间(秒)
	DeletedAt sqltime.Datetime `db:"f_deleted_at,default='0001-01-01 00:00:00'" json:"deletedAt"`
}

func (cud CreationModificationDeletionDatetime) SoftDeletion() (string, []string, driver.Value) {
	return "DeletedAt", []string{"UpdatedAt"}, sqltime.DatetimeZero
}

func (cud *CreationModificationDeletionDatetime) MarkDeletedAt() {
	cud.MarkModifiedAt()
	cud.DeletedAt = cud.UpdatedAt
}

type (
	OperationDatetime = CreationModificationDeletionDatetime
)
