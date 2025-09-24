// Package main provides arithmetic operations for command-line applications.
// This file contains subtraction operations with comprehensive documentation and examples.
package main

import (
	"fmt"
	"strconv"
)

// Subtract performs subtraction of two integers and returns the result.
// This function takes two integer parameters and returns their difference (a - b).
//
// Parameters:
//   - a: First integer operand (minuend)
//   - b: Second integer operand (subtrahend)
//
// Returns:
//   - int: The difference of a and b (a - b)
//
// Example:
//   result := Subtract(10, 3)
//   fmt.Println(result) // Output: 7
//
//   result2 := Subtract(5, 8)
//   fmt.Println(result2) // Output: -3
func Subtract(a, b int) int {
	return a - b
}

// SubtractFloat performs subtraction of two floating-point numbers and returns the result.
// This function handles decimal numbers with precision.
//
// Parameters:
//   - a: First float64 operand (minuend)
//   - b: Second float64 operand (subtrahend)
//
// Returns:
//   - float64: The difference of a and b (a - b)
//
// Example:
//   result := SubtractFloat(10.75, 3.25)
//   fmt.Printf("%.2f\n", result) // Output: 7.50
//
//   result2 := SubtractFloat(3.14, 6.28)
//   fmt.Printf("%.2f\n", result2) // Output: -3.14
func SubtractFloat(a, b float64) float64 {
	return a - b
}

// SubtractMultiple performs sequential subtraction of multiple integers from the first number.
// This function subtracts all subsequent numbers from the first number (a - b - c - d...).
//
// Parameters:
//   - first: The initial number (minuend)
//   - numbers: Variable number of integer arguments to subtract
//
// Returns:
//   - int: The result after subtracting all numbers from the first
//
// Example:
//   result := SubtractMultiple(100, 10, 5, 3)
//   fmt.Println(result) // Output: 82 (100 - 10 - 5 - 3)
//
//   result2 := SubtractMultiple(50, 20)
//   fmt.Println(result2) // Output: 30 (50 - 20)
func SubtractMultiple(first int, numbers ...int) int {
	result := first
	for _, num := range numbers {
		result -= num
	}
	return result
}

// SubtractFromStrings converts string representations of numbers to integers and subtracts them.
// This function is useful for command-line applications that receive numeric input as strings.
//
// Parameters:
//   - a: String representation of first number (minuend)
//   - b: String representation of second number (subtrahend)
//
// Returns:
//   - int: The difference of the converted numbers (a - b)
//   - error: Error if conversion fails
//
// Example:
//   result, err := SubtractFromStrings("100", "42")
//   if err != nil {
//       fmt.Println("Error:", err)
//   } else {
//       fmt.Println(result) // Output: 58
//   }
func SubtractFromStrings(a, b string) (int, error) {
	numA, err := strconv.Atoi(a)
	if err != nil {
		return 0, fmt.Errorf("failed to convert '%s' to integer: %w", a, err)
	}
	
	numB, err := strconv.Atoi(b)
	if err != nil {
		return 0, fmt.Errorf("failed to convert '%s' to integer: %w", b, err)
	}
	
	return Subtract(numA, numB), nil
}

// SubtractWithValidation performs subtraction with input validation and overflow checking.
// This function provides additional safety by checking for potential integer overflow scenarios.
//
// Parameters:
//   - a: First integer operand (minuend)
//   - b: Second integer operand (subtrahend)
//
// Returns:
//   - int: The difference of a and b
//   - error: Error if validation fails or overflow is detected
//
// Example:
//   result, err := SubtractWithValidation(100, 25)
//   if err != nil {
//       fmt.Println("Error:", err)
//   } else {
//       fmt.Println(result) // Output: 75
//   }
func SubtractWithValidation(a, b int) (int, error) {
	// Check for potential underflow when subtracting a positive number from a negative number
	// or when the result would be less than the minimum int value
	if b > 0 && a < 0 {
		// Check if a - b would underflow
		if a < (-2147483648 + b) {
			return 0, fmt.Errorf("subtraction would cause integer underflow: %d - %d", a, b)
		}
	}
	
	// Check for potential overflow when subtracting a negative number from a positive number
	if b < 0 && a > 0 {
		// Check if a - b would overflow (equivalent to a + |b|)
		if a > (2147483647 + b) {
			return 0, fmt.Errorf("subtraction would cause integer overflow: %d - %d", a, b)
		}
	}
	
	return a - b, nil
}

// AbsoluteSubtract performs subtraction and returns the absolute value of the result.
// This function always returns a non-negative result.
//
// Parameters:
//   - a: First integer operand
//   - b: Second integer operand
//
// Returns:
//   - int: The absolute difference between a and b
//
// Example:
//   result := AbsoluteSubtract(5, 10)
//   fmt.Println(result) // Output: 5 (|5 - 10| = 5)
//
//   result2 := AbsoluteSubtract(15, 3)
//   fmt.Println(result2) // Output: 12 (|15 - 3| = 12)
func AbsoluteSubtract(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

// demonstrateSubtractionOperations shows examples of all subtraction functions.
// This function serves as a comprehensive example of how to use the subtraction operations.
func demonstrateSubtractionOperations() {
	fmt.Println("=== Subtraction Operations Demo ===")
	
	// Basic integer subtraction
	fmt.Printf("Subtract(25, 10) = %d\n", Subtract(25, 10))
	fmt.Printf("Subtract(5, 8) = %d\n", Subtract(5, 8))
	
	// Float subtraction
	fmt.Printf("SubtractFloat(10.75, 3.25) = %.2f\n", SubtractFloat(10.75, 3.25))
	fmt.Printf("SubtractFloat(3.14, 6.28) = %.2f\n", SubtractFloat(3.14, 6.28))
	
	// Multiple number subtraction
	fmt.Printf("SubtractMultiple(100, 10, 5, 3) = %d\n", SubtractMultiple(100, 10, 5, 3))
	fmt.Printf("SubtractMultiple(50, 20, 15) = %d\n", SubtractMultiple(50, 20, 15))
	
	// String to number subtraction
	result, err := SubtractFromStrings("456", "123")
	if err != nil {
		fmt.Printf("Error in SubtractFromStrings: %v\n", err)
	} else {
		fmt.Printf("SubtractFromStrings(\"456\", \"123\") = %d\n", result)
	}
	
	// Example with error handling
	_, err = SubtractFromStrings("abc", "123")
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
	
	// Subtraction with validation
	validResult, err := SubtractWithValidation(100, 25)
	if err != nil {
		fmt.Printf("Error in SubtractWithValidation: %v\n", err)
	} else {
		fmt.Printf("SubtractWithValidation(100, 25) = %d\n", validResult)
	}
	
	// Absolute subtraction
	fmt.Printf("AbsoluteSubtract(5, 10) = %d\n", AbsoluteSubtract(5, 10))
	fmt.Printf("AbsoluteSubtract(15, 3) = %d\n", AbsoluteSubtract(15, 3))
	
	fmt.Println("=== End Demo ===")
}

// Uncomment the main function below to run the demonstration
/*
func main() {
	demonstrateSubtractionOperations()
}
*/
