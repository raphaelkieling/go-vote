package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/raphaelkieling/go-vote/models"
	"gorm.io/gorm"
)

type CampaignHandler struct {
	DB *gorm.DB
}

func (ch *CampaignHandler) Vote(c *fiber.Ctx) error {
	tx := ch.DB.Begin()
	id := c.Params("id")
	campaingID, err := strconv.Atoi(id)

	if err != nil {
		tx.Rollback()
		log.Println(err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error on convert number from parameter",
		})
	}

	ip := c.IP()

	err = models.CanVote(tx, ip, uint(campaingID))

	if err != nil {
		tx.Rollback()
		log.Println(err)

		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	vote := models.Vote{
		IP:         ip,
		CampaignID: uint(campaingID),
	}

	err = models.CreateVote(tx, vote)

	if err != nil {
		tx.Rollback()
		log.Println(err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error on create a new vote",
		})
	}

	campaign, err := models.GetCampaignById(tx, uint(campaingID))

	if err != nil {
		tx.Rollback()
		log.Println(err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Campaign not found",
		})
	}

	tx.Commit()

	return c.Status(200).JSON(fiber.Map{
		"data":    campaign,
		"message": "Vote ocurred with success",
	})
}

func (ch *CampaignHandler) Create(c *fiber.Ctx) error {
	tx := ch.DB.Begin()

	campaignDTO := new(models.Campaign)

	if err := c.BodyParser(campaignDTO); err != nil {
		tx.Rollback()
		return err
	}

	if err := models.ValidateCampaign(campaignDTO); err != nil {
		tx.Rollback()
		log.Println(err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	models.CreateCampaign(tx, campaignDTO)

	audit := &models.Audit{
		Action: fmt.Sprintf("Created a new campaing with a ID %d", campaignDTO.ID),
	}

	models.CreateAudit(tx, audit)

	tx.Commit()

	return c.JSON(campaignDTO)
}

func (ch *CampaignHandler) GetAll(c *fiber.Ctx) error {
	var campains []models.Campaign

	if err := ch.DB.Find(&campains).Error; err != nil {
		log.Println("Problema ao pegar as campanhas do banco")
		return err
	}

	return c.JSON(campains)
}
