package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Vote is used to save a campaign vote
type Vote struct {
	ID         uint      `json:"id", gorm:"primaryKey"`
	IP         string    `json:"ip", gorm:"index"`
	CampaignID uint      `json:"campaign_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func CreateVote(tx *gorm.DB, vote Vote) error {
	if err := tx.Create(&vote).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

type VoteCanVote struct {
	Quantity int
}

func CanVote(tx *gorm.DB, ip string, campaingID uint) error {
	var voteCanVote VoteCanVote
	if err := tx.Model(&Vote{}).Select("count(id) as quantity").Where("ip = ? and campaign_id = ?", ip, campaingID).Find(&voteCanVote).Error; err != nil {
		return err
	}

	// TODO: Verify if is correct with tests
	if voteCanVote.Quantity > 3 {
		return errors.New("You exceeded the limit of votes")
	}

	return nil
}
