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
	CreatedAt time.Time  `gorm:"index;type:datetime;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"index;type:datetime"`
}

// BaseModelSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}
