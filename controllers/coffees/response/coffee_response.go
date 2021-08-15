package response

import (
	articles "finalProject/business/coffees"
	"time"
)

type Coffee struct {
	Id             int       `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	Name           string    `json:"name"`
	ProcessingType string    `json:"processing_type"`
	RoastingLevel  string    `json:"roasting_level"`
	Elevation      string    `json:"elevation"`
	BeanType       string    `json:"bean_type"`
	Region         string    `json:"region"`
}

func FromDomain(domain articles.Domain) Coffee {
	return Coffee{
		Id:             domain.Id,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Name:           domain.Name,
		ProcessingType: domain.ProcessingType,
		RoastingLevel:  domain.RoastingLevel,
		Elevation:      domain.Elevation,
		BeanType:       domain.BeanType,
		Region:         domain.Region,
	}
}
