package coffees

import (
	coffeesUsecase "finalProject/business/coffees"
	"time"
)

type Coffees struct {
	Id             int
	Name           string
	ProcessingType string
	RoastingLevel  string
	Elevation      string
	BeanType       string
	Region         string
	CreatedAt      time.Time `gorm:"<-create"`
	UpdatedAt      time.Time
}

func fromDomain(domain *coffeesUsecase.Domain) *Coffees {
	return &Coffees{
		Id:             domain.Id,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		Name:           domain.Name,
		ProcessingType: domain.ProcessingType,
		RoastingLevel:  domain.RoastingLevel,
		Elevation:      domain.Elevation,
		BeanType:       domain.BeanType,
		Region:         domain.Region,
	}
}

func (rec *Coffees) toDomain() coffeesUsecase.Domain {
	return coffeesUsecase.Domain{
		Name:           rec.Name,
		ProcessingType: rec.ProcessingType,
		RoastingLevel:  rec.RoastingLevel,
		Elevation:      rec.Elevation,
		BeanType:       rec.BeanType,
		Region:         rec.Region,
	}
}
