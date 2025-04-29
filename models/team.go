package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Conference   string `gorm:"type:text"`
	Division     string `gorm:"type:text"`
	City         string `gorm:"type:text"`
	Name         string `gorm:"type:text"`
	FullName     string `gorm:"type:text"`
	Abbreviation string `gorm:"type:text"`
	LogoBase64   string `gorm:"type:text"`
}
