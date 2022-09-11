package services

import (
	"errors"
	"time"

	"goTravel2.0/database"
	"goTravel2.0/models"
)

type Service struct {
	database *database.Database
}

func NewService(database *database.Database) *Service {
	return &Service{database}
}

func (s *Service) AddClothes(Clothes *models.Clothes) error {
	if Clothes.Underwear <= 0 {
		return errors.New("this field cannot be empty or less than 0")
	}
	if Clothes.Pants <= 0 {
		return errors.New("this field cannot be empty or less than 0")
	}
	if Clothes.Shirts <= 0 {
		return errors.New("this field cannot be empty or less than 0")
	}
	if Clothes.TShirts <= 0 {
		return errors.New("this field cannot be empty or less than 0")
	}
	if Clothes.Shoes <= 0 {
		return errors.New("this field cannot be empty or less than 0")
	}
	s.database.AddClothes(Clothes)

	return nil
}

func (s *Service) GetClothesById(ourID string) (*models.Clothes, error) {
	if ourID == "" {
		return nil, errors.New("this field cannot be empty or less")
	}

	result := s.database.GetClothesById(ourID)

	return &result, nil

}

func (s *Service) AddTravel(Travel *models.Travel) error {

	if len(Travel.Destination) <= 0 {
		return errors.New("this field cannot be empty or less")
	}
	now := time.Now()

	isNotValid := now.Before(Travel.Date)
	if isNotValid {
		return errors.New("you can't alter the sacred timeline")
	}
	if Travel.Budget <= 0 {
		return errors.New("this field cannot be empty or less to 0")
	}

	s.database.AddTravel(Travel)

	return nil
}

func (s *Service) SearchForTravel(searchString string) ([]models.Travel, error) {
	if len(searchString) <= 0 {
		return nil, errors.New("this field cannot be empty or less")
	}

	currentTravel := s.database.SearchForTravel(searchString)

	return currentTravel, nil

}

func (s *Service) GetTravelById(ourID string) (*models.Travel, error) {
	if ourID == "" {
		return nil, errors.New("this field cannot be empty or less to 0")
	}

	result := s.database.GetTravelById(ourID)

	return &result, nil

}

func (s *Service) UpdateTravel(OurTravel models.Travel) int64 {

	result := s.database.UpdateTravel(OurTravel)
	return result
}

func (s *Service) DeleteTravel(idToDelete string) int64 {
	affected := s.database.DeleteTravel(idToDelete)
	return affected
}
func (s *Service) DeleteClothes(idToDelete string) int64 {
	affected := s.database.DeleteClothes(idToDelete)
	return affected
}
