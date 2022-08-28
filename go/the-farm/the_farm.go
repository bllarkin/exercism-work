package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	} else if cows < 0 {
		return 0.0, &SillyNephewError{cows: cows}
	}

	if err == nil {
		if amount >= 0.0 {
			return amount / float64(cows), err
		} else {
			return 0.0, errors.New("negative fodder")
		}
	} else if err == ErrScaleMalfunction {
		if amount >= 0.0 {
			amount = amount * 2
			return amount / float64(cows), nil
		} else {
			return 0.0, errors.New("negative fodder")
		}
	} else {
		return 0, err
	}
}
