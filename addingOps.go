// Package main provides arithmetic operations for command-line applications.
// This file contains addition operations with comprehensive documentation and examples.
package main

import (
	"fmt"
	"strconv"
)

// Add performs addition of two integers and returns the result.
// This function takes two integer parameters and returns their sum.
//
// Parameters:
//   - a: First integer operand
//   - b: Second integer operand
//
// Returns:
//   - int: The sum of a and b
//
// Example:
//   result := Add(5, 3)
//   fmt.Println(result) // Output: 8
func Add(a, b int) int {
	return a + b
}

// AddFloat performs addition of two floating-point numbers and returns the result.
// This function handles decimal numbers with precision.
//
// Parameters:
//   - a: First float64 operand
//   - b: Second float64 operand
//
// Returns:
//   - float64: The sum of a and b
//
// Example:
//   result := AddFloat(3.14, 2.86)
//   fmt.Printf("%.2f\n", result) // Output: 6.00
func AddFloat(a, b float64) float64 {
	return a + b
}

// AddMultiple performs addition of multiple integers using variadic parameters.
// This function can accept any number of integer arguments and returns their sum.
//
// Parameters:
//   - numbers: Variable number of integer arguments
//
// Returns:
//   - int: The sum of all provided numbers
//
// Example:
//   result := AddMultiple(1, 2, 3, 4, 5)
//   fmt.Println(result) // Output: 15
//
//   result2 := AddMultiple(10, 20)
//   fmt.Println(result2) // Output: 30
func AddMultiple(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// AddFromStrings converts string representations of numbers to integers and adds them.
// This function is useful for command-line applications that receive numeric input as strings.
//
// Parameters:
//   - a: String representation of first number
//   - b: String representation of second number
//
// Returns:
//   - int: The sum of the converted numbers
//   - error: Error if conversion fails
//
// Example:
//   result, err := AddFromStrings("42", "58")
//   if err != nil {
//       fmt.Println("Error:", err)
//   } else {
//       fmt.Println(result) // Output: 100
//   }
func AddFromStrings(a, b string) (int, error) {
	numA, err := strconv.Atoi(a)
	if err != nil {
		return 0, fmt.Errorf("failed to convert '%s' to integer: %w", a, err)
	}
	
	numB, err := strconv.Atoi(b)
	if err != nil {
		return 0, fmt.Errorf("failed to convert '%s' to integer: %w", b, err)
	}
	
	return Add(numA, numB), nil
}

// demonstrateAdditionOperations shows examples of all addition functions.
// This function serves as a comprehensive example of how to use the addition operations.
func demonstrateAdditionOperations() {
	fmt.Println("=== Addition Operations Demo ===")
	
	// Basic integer addition
	fmt.Printf("Add(10, 25) = %d\n", Add(10, 25))
	
	// Float addition
	fmt.Printf("AddFloat(3.14, 2.86) = %.2f\n", AddFloat(3.14, 2.86))
	
	// Multiple number addition
	fmt.Printf("AddMultiple(1, 2, 3, 4, 5) = %d\n", AddMultiple(1, 2, 3, 4, 5))
	fmt.Printf("AddMultiple(100, 200, 300) = %d\n", AddMultiple(100, 200, 300))
	
	// String to number addition
	result, err := AddFromStrings("123", "456")
	if err != nil {
		fmt.Printf("Error in AddFromStrings: %v\n", err)
	} else {
		fmt.Printf("AddFromStrings(\"123\", \"456\") = %d\n", result)
	}
	
	// Example with error handling
	_, err = AddFromStrings("abc", "123")
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
	
	fmt.Println("=== End Demo ===")
}

// Uncomment the main function below to run the demonstration
/*
func main() {
	demonstrateAdditionOperations()
}
*/
