package services

import (
	"context"
	"odisha_gov_be/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Create() models.District {
	district := models.District{
		Name: "Angul",
	}
	district.InsertG(context.Background(), boil.Infer())
	return district
}
