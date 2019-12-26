package models

import (
	"github.com/jinzhu/gorm"
	"time"

	"github.com/gofrs/uuid"
)

// BaseModel defines the common columns that all db structs should hold, usually
// db structs based on this have no soft delete
type BaseModel struct {
	ID        uuid.UUID  `gorm:"primary_key;type:varchar(36)"`
	CreatedAt *time.Time `gorm:"index;type:datetime;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"index;type:datetime"`
}

// BaseModelSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `gorm:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

// BaseModelSeq defines the common columns that all db structs should hold, with
// an INT key
type BaseModelSeq struct {
	// Default values for PostgreSQL, change it for other DBMS
	ID        int        `gorm:"primary_key,auto_increment"`
	CreatedAt *time.Time `gorm:"index;type:datetime;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"index;type:datetime"`
}

// BaseModelSeqSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSeqSoftDelete struct {
	BaseModelSeq
	DeletedAt *time.Time `gorm:"index"`
}
