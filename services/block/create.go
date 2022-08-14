package services

import (
	"context"
	"odisha_gov_be/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Create() models.Block {
	block := models.Block{
		Name: "Angul",
	}
	block.InsertG(context.Background(), boil.Infer())
	return block
}
