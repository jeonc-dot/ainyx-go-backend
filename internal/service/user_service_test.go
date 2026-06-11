package service

import (
	"testing"
	"time"
)

// TestCalculateAge ensures our dynamic age calculation handles edge cases 
// like upcoming birthdays perfectly.
func TestCalculateAge(t *testing.T) {
	today := time.Now()
	
	// Test Case 1: Birthday already happened this year
	// We simulate a birth date exactly 20 years and 1 month ago.
	pastDOB := today.AddDate(-20, -1, 0).Format("2006-01-02")
	age, err := CalculateAge(pastDOB)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if age != 20 {
		t.Errorf("expected 20, got %d", age)
	}

	// Test Case 2: Birthday hasn't happened yet this year (The Trap!)
	// We simulate a birth date 20 years ago, but 1 month in the future.
	futureDOB := today.AddDate(-20, 1, 0).Format("2006-01-02")
	age, err = CalculateAge(futureDOB)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Because their birthday hasn't happened yet, they should still be 19!
	if age != 19 {
		t.Errorf("expected 19, got %d", age)
	}
}
