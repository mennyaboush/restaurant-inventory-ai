// Lesson 1.1: Go Variables, Types, and Constants
// Run with: go run learn/01_basics/main.go
package main

import "fmt"

// Constants - values that never change
const (
	AppName    = "Restaurant Inventory"
	AppVersion = "0.1.0"
)

func main() {
	fmt.Println("🎓 Lesson 1.1: Go Variables & Types")
	fmt.Println("====================================")

	// ============================================
	// SECTION 1: Variable Declarations
	// ============================================
	fmt.Println("\n📖 SECTION 1: Variable Declarations")

	// Method 1: Full declaration with type
	var productName string = "Coca Cola 330ml Can"
	fmt.Println("Full declaration:", productName)

	// Method 2: Type inference (Go figures out the type)
	var quantity = 24 // Go knows this is int
	fmt.Println("Type inference:", quantity)

	// Method 3: Short declaration (most common, only inside functions)
	price := 45.50 // := means "declare and assign"
	fmt.Println("Short declaration:", price)

	// ============================================
	// SECTION 2: Basic Types
	// ============================================
	fmt.Println("\n📖 SECTION 2: Basic Types")

	// Integer types
	var boxes int = 3
	var units int = 5

	// Float types (decimal numbers)
	var pricePerBox float64 = 55.00

	// String (text)
	var category string = "משקאות" // Hebrew!

	// Boolean (true/false)
	var isActive bool = true

	fmt.Printf("Boxes: %d, Units: %d\n", boxes, units)
	fmt.Printf("Price: %.2f NIS\n", pricePerBox)
	fmt.Printf("Category: %s\n", category)
	fmt.Printf("Active: %t\n", isActive)

	// ============================================
	// SECTION 3: Zero Values
	// ============================================
	fmt.Println("\n�� SECTION 3: Zero Values (uninitialized variables)")

	var uninitInt int
	var uninitString string
	var uninitBool bool
	var uninitFloat float64

	fmt.Printf("int zero value: %d\n", uninitInt)         // 0
	fmt.Printf("string zero value: '%s'\n", uninitString) // "" (empty)
	fmt.Printf("bool zero value: %t\n", uninitBool)       // false
	fmt.Printf("float zero value: %f\n", uninitFloat)     // 0.000000

	// ============================================
	// SECTION 4: Calculations
	// ============================================
	fmt.Println("\n📖 SECTION 4: Calculations")

	boxSize := 24
	totalUnits := (boxes * boxSize) + units

	fmt.Printf("Boxes: %d × %d = %d units\n", boxes, boxSize, boxes*boxSize)
	fmt.Printf("Plus loose units: %d\n", units)
	fmt.Printf("Total: %d units\n", totalUnits)

	// Type conversion required for mixed types
	totalValue := float64(totalUnits) * (pricePerBox / float64(boxSize))
	fmt.Printf("Total value: %.2f NIS\n", totalValue)

	// ============================================
	// SECTION 5: Constants
	// ============================================
	fmt.Println("\n📖 SECTION 5: Constants")
	fmt.Printf("App: %s v%s\n", AppName, AppVersion)

	// Constants for our domain
	const (
		MovementIN         = "IN"
		MovementOUT        = "OUT"
		MovementWASTE      = "WASTE"
		MovementADJUSTMENT = "ADJUSTMENT"
	)
	fmt.Printf("Movement types: %s, %s, %s, %s\n",
		MovementIN, MovementOUT, MovementWASTE, MovementADJUSTMENT)

	// ============================================
	// SECTION 6: Printf Formatting
	// ============================================
	fmt.Println("\n📖 SECTION 6: Printf Formatting")
	fmt.Println("Common format verbs:")
	fmt.Printf("  %%s = string: %s\n", productName)
	fmt.Printf("  %%d = integer: %d\n", boxes)
	fmt.Printf("  %%f = float: %f\n", price)
	fmt.Printf("  %%.2f = float 2 decimals: %.2f\n", price)
	fmt.Printf("  %%t = boolean: %t\n", isActive)
	fmt.Printf("  %%T = type: %T\n", price)
	fmt.Printf("  %%v = any value: %v\n", boxes)

	fmt.Println("\n✅ Lesson 1.1 Complete!")
}
