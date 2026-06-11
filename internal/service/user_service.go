package service

import (
	"time"
)

// CalculateAge computes current age from a DOB string (YYYY-MM-DD format).
func CalculateAge(dob string) (int, error) {
	birthDate, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return 0, err 
	}

	today := time.Now()
	age := today.Year() - birthDate.Year()

	// Adjust for leap years and upcoming birthdays in the current year
	if today.Month() < birthDate.Month() || (today.Month() == birthDate.Month() && today.Day() < birthDate.Day()) {
		age--
	}

	return age, nil
}
