package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Product struct
type Product struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Basket struct
type Basket struct {
	ID       string    `json:"id"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}

// Customer struct
type Customer struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Cash      float64 `json:"cash"`
	Basket    Basket  `json:"basket"`
}

func main() {
	// Reading JSON file
	filename := "store_data.json"
	customers, err := readData(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Task 1: Print information about all customers and the total amount of money spent.
	fmt.Println("Task 1: Display information about all customers and the total amount spent on purchases:")
	printCustomerDetails(customers)

	// Task 2: Find and display the top spender.
	fmt.Println("\nTask 2: Top spending customer")
	topSpender := findTopSpender(customers)
	printCustomerInfo(topSpender)

	// Task 3: Find and display the most expensive product.
	fmt.Println("\nTask 3: Most expensive product")
	allProds := allProducts(customers)
	mostExpensiveProduct := findMostExpensiveProduct(allProds)
	printProductInfo(mostExpensiveProduct)

	// Task 4: Calculate and display the average price of all products.
	fmt.Println("\nTask 4: Average price of all products")
	calculateAndPrintAveragePrice(allProds)

	// Task 5: Find and display the customer who made the lowest purchase.
	fmt.Println("\nTask 5: Customer with the lowest purchase amount")
	printLowestSpender(customers)

	// Task 6: Find and display the best-selling category.
	fmt.Println("\nTask 6:")
	bestSellingCategory := findBestSellingCategory(customers)
	fmt.Println("Best-selling category:", bestSellingCategory)

	// Task 7: Display the most and least sold products.
	printMinMaxSoldProducts(customers)

	// Task 8: Calculate and display the average quantity of products sold per purchase:
	fmt.Println("\nTask 8:")
	calculateAndPrintAverageQuantitySold(customers)

	// Task 9: Identify and display the customer who purchased the most products and the total number of products purchased.
	fmt.Println("\nTask 9: ")
	findTopCustomerByProductQuantity(customers)

	// Task 10: Find and display the most frequently sold product among all purchases.
	// fmt.Println("\nTask 10: ")
	// findMostSoldProduct(allProds)

	// Task 11: Identify and display the customer who has the highest average spending.
	fmt.Println("\nTask 11: ")
	calculateAndPrintAverageSpending(customers)

	// Task 12: Display the category that generates the highest total revenue.
	fmt.Println("\nTask 12: ")
	findMostProfitableCategory(customers)

	// Task 13: Display the most expensive purchase made by each customer.
	fmt.Println("\nTask 13: ")
	findMostExpensivePurchaseByCustomer(customers)

	// Task 14: Display the most expensive category purchased by each customer and the amount spent in that category.
	fmt.Println("\nTask 14: ")
	findMostExpensiveCategoryByCustomer(customers)

	// Task 15: Display the total quantity sold for each product and the overall total number of products sold.
	// fmt.Println("\nTask 15: ")
	// printTotalSoldQuantity(allProds)
}

// Reading data from JSON
func readData(filename string) ([]Customer, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var customers []Customer
	err = json.Unmarshal(data, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// Printing customer and basket information
func printCustomerInfo(customer Customer) {
	fmt.Printf("Name: %s, Last Name: %s, Customer Cash: %.2f\n",
		customer.FirstName, customer.LastName, customer.Cash)

	// Printing the shopping basket
	for _, product := range customer.Basket.Products {
		fmt.Printf("   Category: %s, Name: %s, Price: %.2f, Quantity: %d\n",
			product.Category, product.Name, product.Price, product.Quantity)
	}

	fmt.Printf("   Total Basket Amount: %.2f\n", customer.Basket.Total)
	fmt.Println("------------------------------")
}

// Printing all customer information
func printAllCustomers(customers []Customer) {
	for _, customer := range customers {
		printCustomerInfo(customer)
	}
}

// Task 1: Print details of all customers
func printCustomerDetails(customers []Customer) {
	totalCash := 0.0
	totalSpent := 0.0

	for _, customer := range customers {
		printCustomerInfo(customer)
		totalCash += customer.Cash
		totalSpent += customer.Basket.Total
	}

	fmt.Printf("Total cash of all customers: %.2f\n", totalCash)
	fmt.Printf("Total spent overall: %.2f\n", totalSpent)
}

// Find the top spender
func findTopSpender(customers []Customer) Customer {
	// If there are no customers, return an empty customer
	if len(customers) == 0 {
		return Customer{}
	}

	topSpender := customers[0]

	for _, customer := range customers {
		if customer.Basket.Total > topSpender.Basket.Total {
			topSpender = customer
		}
	}

	return topSpender
}

// Find the most expensive product
func findMostExpensiveProduct(products []Product) Product {
	if len(products) == 0 {
		// If there are no products, return an empty product
		return Product{}
	}

	mostExpensive := products[0]

	for _, product := range products {
		if product.Price > mostExpensive.Price {
			mostExpensive = product
		}
	}

	return mostExpensive
}

// Get all products
func allProducts(customers []Customer) []Product {
	var allProds []Product
	for _, customer := range customers {
		allProds = append(allProds, customer.Basket.Products...)
	}
	return allProds
}

// Print product information
func printProductInfo(product Product) {
	fmt.Printf("Category: %s\n", product.Category)
	fmt.Printf("Product name: %s\n", product.Name)
	fmt.Printf("Price: %.0f\n", product.Price)
	fmt.Printf("Quantity: %d\n", product.Quantity)
	fmt.Println("------------------------------")
}

// Calculate and print the average price of all products
func calculateAndPrintAveragePrice(allProducts []Product) {
	if len(allProducts) == 0 {
		fmt.Println("No products found.")
		return
	}

	var total float64
	for _, product := range allProducts {
		total += product.Price
	}

	average := total / float64(len(allProducts))
	fmt.Printf("Average price of all products: %.0f\n", average)
}

// Find the customer who spent the least in total with the lowest price
func findLowestSpender(customers []Customer) Customer {
	if len(customers) == 0 {
		return Customer{}
	}

	lowestSpender := customers[0]

	for _, customer := range customers {
		if customer.Basket.Total < lowestSpender.Basket.Total {
			lowestSpender = customer
		}
	}

	return lowestSpender
}

// Print the customer who spent the least at the lowest price
func printLowestSpender(customers []Customer) {
	lowestSpender := findLowestSpender(customers)

	if lowestSpender.ID == "" {
		fmt.Println("Customer not found.")
		return
	}

	fmt.Println("Customer who spent the least:")
	printCustomerInfo(lowestSpender)
}

// Determine the best selling product category
func findBestSellingCategory(customers []Customer) string {
	if len(customers) == 0 {
		return ""
	}

	categoryCounts := make(map[string]int)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			categoryCounts[product.Category] += product.Quantity
		}
	}

	bestSellingCategory := ""
	maxQuantity := 0
	for category, quantity := range categoryCounts {
		if quantity > maxQuantity {
			bestSellingCategory = category
			maxQuantity = quantity
		}
	}

	return bestSellingCategory
}

// Find the most and least sold products
func findMinMaxSoldProducts(customers []Customer) (Product, Product) {
	allProds := allProducts(customers)
	if len(allProds) == 0 {
		return Product{}, Product{}
	}

	maxSold := allProds[0]
	minSold := allProds[0]

	for _, product := range allProds {
		if product.Quantity > maxSold.Quantity {
			maxSold = product
		} else if product.Quantity < minSold.Quantity {
			minSold = product
		}
	}

	return maxSold, minSold
}

func printMinMaxSoldProducts(customers []Customer) {
	maxSold, minSold := findMinMaxSoldProducts(customers)

	if maxSold.ID == "" || minSold.ID == "" {
		fmt.Println("Product not found.")
		return
	}

	fmt.Println("---------------------------------------")
	fmt.Println("Task 7:\nMost and least sold products")
	fmt.Println("Most sold product:")
	printProductInfo(maxSold)
	fmt.Println("Least sold product:")
	printProductInfo(minSold)
}

// Task 8: Calculate and display the average quantity of products sold per sale
func calculateAndPrintAverageQuantitySold(customers []Customer) {
	totalSales := len(customers)
	totalQuantity := 0

	for _, customer := range customers {
		totalQuantity += len(customer.Basket.Products)
	}

	average := float64(totalQuantity) / float64(totalSales)
	fmt.Printf("Average product quantity per sale: %d / %d = %.3f\n", totalQuantity, totalSales, average)
}

// Calculate the total quantity of products in the basket
func calculateBasketTotalQuantity(basket Basket) int {
	totalQuantity := 0
	for _, product := range basket.Products {
		totalQuantity += product.Quantity
	}
	return totalQuantity
}

// Task 9: Find and display the customer who purchased the most products
func findTopCustomerByProductQuantity(customers []Customer) {
	if len(customers) == 0 {
		fmt.Println("Customer not found")
		return
	}

	topCustomer := customers[0]
	totalQuantity := 0

	for _, customer := range customers {
		quantity := len(customer.Basket.Products)
		totalQuantity += quantity

		if quantity > len(topCustomer.Basket.Products) {
			topCustomer = customer
		}
	}

	fmt.Printf("Customer with the most purchased products:\n")
	printCustomerInfo(topCustomer)
	fmt.Printf("Total quantity purchased: %d\n", totalQuantity)
}

// Task 10: Find and display the most frequently sold product among sold products
func findMostSoldProduct(customers []Customer) {
	fmt.Println("\nTask 10: ")
	productFrequency := make(map[string]int)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			productFrequency[product.Name]++
		}
	}
	var mostSeenProductName string
	mostSeenCount := 0
	for productName, frequency := range productFrequency {
		if frequency > mostSeenCount {
			mostSeenCount = frequency
			mostSeenProductName = productName
		}
	}

	if mostSeenCount > 0 {
		fmt.Printf("Most frequently seen product in sales: %s (%d units)\n", mostSeenProductName, mostSeenCount)
	} else {
		fmt.Println("No products found in customers' baskets.")
	}
}

// Find products by ID
func findProductByID(products []Product, productID string) Product {
	for _, product := range products {
		if product.ID == productID {
			return product
		}
	}
	return Product{}
}

// Task 11: Find and display the customer with the highest average spending per sale
func calculateAndPrintAverageSpending(customers []Customer) {
	fmt.Println("\nTask 11: ")
	maxAverageSpending := 0
	var topSpenderFirstName, topSpenderLastName string

	for _, customer := range customers {
		totalSpending := 0
		for _, product := range customer.Basket.Products {
			totalSpending += int(product.Price) * product.Quantity
		}

		averageSpending := totalSpending / len(customer.Basket.Products)
		if averageSpending > maxAverageSpending {
			maxAverageSpending = averageSpending
			topSpenderFirstName = customer.FirstName
			topSpenderLastName = customer.LastName
		}
	}

	fmt.Printf("Average spending per sale: %d sum\n", maxAverageSpending)
	fmt.Printf("Customer with the highest spending: %s %s\n", topSpenderFirstName, topSpenderLastName)
}

// Task 12: Display the category that generates the highest overall profit (quantity*price)
func findMostProfitableCategory(customers []Customer) {
	fmt.Println("\nTask 12: ")
	if len(customers) == 0 {
		fmt.Println("Customer not found")
		return
	}

	categoryProfits := make(map[string]float64)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			profit := float64(product.Quantity) * product.Price
			categoryProfits[product.Category] += profit
		}
	}

	mostProfitableCategory := ""
	maxProfit := 0.0
	for category, profit := range categoryProfits {
		if profit > maxProfit {
			mostProfitableCategory = category
			maxProfit = profit
		}
	}

	fmt.Printf("Most profitable category: %s (Total profit: %.2f)\n", mostProfitableCategory, maxProfit)
}

// Task 13: Display the most expensive purchase made by each customer
func findMostExpensivePurchaseByCustomer(customers []Customer) {
	fmt.Println("\nTask 13: ")
	if len(customers) == 0 {
		fmt.Println("Customer not found")
		return
	}

	for _, customer := range customers {
		mostExpensiveProduct := findMostExpensiveProduct(customer.Basket.Products)

		if mostExpensiveProduct.ID != "" {
			fmt.Printf("Most expensive purchase made by %s %s:\n", customer.FirstName, customer.LastName)
			printProductInfo(mostExpensiveProduct)
		} else {
			fmt.Printf("No most expensive purchase found for %s %s.\n", customer.FirstName, customer.LastName)
		}
	}
}

// Task 14: Display the category where each customer made the most expensive purchase and the amount spent in that category
func findMostExpensiveCategoryByCustomer(customers []Customer) {
	fmt.Println("\nTask 14: ")
	if len(customers) == 0 {
		fmt.Println("Customer not found")
		return
	}

	for _, customer := range customers {
		categorySpending := make(map[string]float64)

		for _, product := range customer.Basket.Products {
			categorySpending[product.Category] += product.Price * float64(product.Quantity)
		}

		mostExpensiveCategory := ""
		maxSpending := 0.0

		for category, spending := range categorySpending {
			if spending > maxSpending {
				mostExpensiveCategory = category
				maxSpending = spending
			}
		}

		if mostExpensiveCategory != "" {
			fmt.Println("--------------------------------------------------------")
			fmt.Printf("Task 14: %s %s's most expensive category of purchase: %s\n", customer.FirstName, customer.LastName, mostExpensiveCategory)
			fmt.Printf("Total spent on all products in this category: %.2f\n", maxSpending)
			fmt.Println("--------------------------------------------------------")
		} else {
			fmt.Printf("Task 14: No most expensive category of purchase found for %s %s.\n", customer.FirstName, customer.LastName)
		}
	}
}

// Task 15: Display the total quantity sold for each product and the overall total number of products sold
func printTotalSoldQuantity(products []Product) {
	fmt.Println("\nTask 15: ")
	if len(products) == 0 {
		fmt.Println("No sold products found.")
		return
	}

	productSoldQuantity := make(map[string]int)
	totalSoldQuantity := 0

	for _, product := range products {
		productSoldQuantity[product.Name] += product.Quantity
		totalSoldQuantity += product.Quantity
	}

	fmt.Println("Task 15: Total quantity sold for each product:")
	for productName, quantity := range productSoldQuantity {
		fmt.Printf("%s: %d units\n", productName, quantity)
	}

	fmt.Printf("Total quantity of all sold products: %d units\n", totalSoldQuantity)
}
