package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidateReview(review *models.Review) error {
	if review.Rating > 5 {
		return errors.New("The rating cannot be higher than 5.")
	}
	if review.Rating < 1 {
		return errors.New("The rating cannot be negative.")

	}
	return nil
}
