package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Campaign to save a group of votes
type Campaign struct {
	ID          uint   `json:"id", gorm:"primaryKey"`
	Description string `validate:"required,min=3,max=250",json:"description"`
	Votes       []Vote `json:"votes", gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func ValidateCampaign(campaign *Campaign) error {
	validate := validator.New()
	err := validate.Struct(campaign)

	if err != nil {
		return err
	}

	return nil
}

func CreateCampaign(tx *gorm.DB, campaign *Campaign) error {
	if err := tx.Create(campaign).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func GetCampaignById(tx *gorm.DB, id uint) (Campaign, error) {
	campaign := Campaign{}
	if err := tx.Preload("Votes").First(&campaign, id).Error; err != nil {
		return Campaign{}, err
	}

	return campaign, nil
}

func UpdateCampaign(tx *gorm.DB, campaign *Campaign) error {
	if err := tx.Save(campaign).Error; err != nil {
		return err
	}

	return nil
}
