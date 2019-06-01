package entity

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	Id            string `gorm:"primary_key;"`
	Name          string
	Description   string
	PriceUnits    int64
	PriceDecimals int64
	Stock         int64
}

func (u *Product) BeforeCreate(scope *gorm.Scope) (err error) {
	generatedUuid, err := uuid.NewV4()
	_ = scope.SetColumn("Id", generatedUuid.String())
	return nil
}
