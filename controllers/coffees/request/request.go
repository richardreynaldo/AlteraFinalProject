package request

import (
	"finalProject/business/coffees"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateCoffee struct {
	Name           string `json:"name"`
	ProcessingType string `json:"processing_type"`
	RoastingLevel  string `json:"roasting_level"`
	Elevation     string `json:"elevation"`
	BeanType       string `json:"bean_type"`
	Region         string `json:"region"`
}

func (req *CreateCoffee) ToDomain() *coffees.Domain {
	return &coffees.Domain{
		Name:           req.Name,
		ProcessingType: req.ProcessingType,
		RoastingLevel:  req.RoastingLevel,
		Elevation:     req.Elevation,
		BeanType:       req.BeanType,
		Region:         req.Region,
	}
}
