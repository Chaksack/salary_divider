package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Salary struct {
	MonthlyIncome float64 `json:"monthly_income"`
}

type CategorizedIncome struct {
	Investment  float64 `json:"investment"`
	Expenditure float64 `json:"expenditure"`
	Savings     float64 `json:"savings"`
	Personal    float64 `json:"personal"`
}

func main() {
	app := fiber.New()

	app.Post("/calculate", func(c *fiber.Ctx) error {
		// Parse the request body
		salary := new(Salary)
		if err := c.BodyParser(salary); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Calculate the categorized income
		categorizedIncome := calculateCategorizedIncome(salary.MonthlyIncome)

		// Return the categorized income
		return c.JSON(categorizedIncome)
	})

	log.Fatal(app.Listen(":3000"))
}

func calculateCategorizedIncome(monthlyIncome float64) CategorizedIncome {
	// Calculate the amounts based on the given percentages
	investment := monthlyIncome * 0.5
	expenditure := monthlyIncome * 0.3
	savings := monthlyIncome * 0.2
	personal := monthlyIncome * 0.1

	// Create a CategorizedIncome struct
	categorizedIncome := CategorizedIncome{
		Investment:  investment,
		Expenditure: expenditure,
		Savings:     savings,
		Personal:    personal,
	}

	return categorizedIncome
}

