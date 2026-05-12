package thefarm

import "fmt"
import "errors"

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	fa, err := fc.FodderAmount(cows)
	if err != nil {
	//	return 0, fmt.Errorf("Failed to get fodder amount: %w", err)
	return 0, err
	}
	ff, err := fc.FatteningFactor()
	if err != nil {
	//	return 0, fmt.Errorf("Failed to get fattening factor: %w", err)
	return 0, err
	} 
	return fa*ff/float64(cows), nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	//if cows <= 0 {
	//	return 0, errors.New("invalid number of cows")
	if err := ValidateNumberOfCows(cows); err != nil {
		return 0, errors.New("invalid number of cows")
	} 
	return DivideFood(fc, cows)
}
// TODO: define the 'ValidateNumberOfCows' function

// 1: create a structure with the number of cows and the message string.
type InvalidCowsError struct {
	cows int
	message string
}

// 2: Now make an Error method on that structure, ith a pointer receiver.
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// 3: Now let's actually make the ValidateNumberOfCows function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	}
	if cows == 0 {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
