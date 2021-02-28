package models

import "gorm.io/gorm"

type Audit struct {
	ID     uint `json:"id",gorm:"primaryKey"`
	Action string
}

func CreateAudit(tx *gorm.DB, audit *Audit) error {
	if err := tx.Create(audit).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
