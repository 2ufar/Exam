package task

import (
	"fmt"
	"ExamFolder/store"
)

// Task 1: Print details of all customers, including their total cash and total spent.
func PrintCustomerDetails(customers []store.Customer) {
	totalCash := 0.0
	totalSpent := 0.0

	for _, customer := range customers {
		store.PrintCustomerInfo(customer)
		totalCash += customer.Cash
		totalSpent += customer.Basket.Total
	}

	fmt.Printf("Total Customer Cash: %.2f\n", totalCash)
	fmt.Printf("Total Amount Spent: %.2f\n", totalSpent)
}

// Task 2: Find the customer who spent the most.
func FindTopSpender(customers []store.Customer) store.Customer {
	if len(customers) == 0 {
		return store.Customer{}
	}

	topSpender := customers[0]

	for _, customer := range customers {
		if customer.Basket.Total > topSpender.Basket.Total {
			topSpender = customer
		}
	}

	return topSpender
}

// Task 3: Find the most expensive product among all products.
func FindMostExpensiveProduct(products []store.Product) store.Product {
	if len(products) == 0 {
		return store.Product{}
	}

	mostExpensive := products[0]

	for _, product := range products {
		if product.Price > mostExpensive.Price {
			mostExpensive = product
		}
	}

	return mostExpensive
}

// Task 4: Collect all products bought by customers.
func AllProducts(customers []store.Customer) []store.Product {
	var allProds []store.Product
	for _, customer := range customers {
		allProds = append(allProds, customer.Basket.Products...)
	}
	return allProds
}

// Task 5: Print details of the customer who spent the least.
func PrintLowestSpender(customers []store.Customer) {
	lowestSpender := FindLowestSpender(customers)

	if lowestSpender.ID == "" {
		fmt.Println("Customer not found.")
		return
	}

	fmt.Println("Customer with the least total purchase amount:")
	store.PrintCustomerInfo(lowestSpender)
}

// Task 6: Find the best-selling category among all products.
func FindBestSellingCategory(customers []store.Customer) string {
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

// Task 7: Find the product sold in maximum and minimum quantity among all customers.
func FindMinMaxSoldProducts(customers []store.Customer) (store.Product, store.Product) {
	allProds := AllProducts(customers)
	if len(allProds) == 0 {
		return store.Product{}, store.Product{}
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

// Task 8: Calculate and print the average quantity of products sold per customer.
func CalculateAndPrintAverageQuantitySold(customers []store.Customer) {
	totalSales := len(customers)
	totalQuantity := 0

	for _, customer := range customers {
		totalQuantity += len(customer.Basket.Products)
	}

	average := float64(totalQuantity) / float64(totalSales)
	fmt.Printf("Average Product Quantity: %d / %d = %.3f\n", totalQuantity, totalSales, average)
}

// Task 9: Find the customer who purchased the most number of products.
func FindTopCustomerByProductQuantity(customers []store.Customer) {
	if len(customers) == 0 {
		fmt.Println("Customer not found.")
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

	fmt.Println("Customer with the Most Products Purchased:")
	store.PrintCustomerInfo(topCustomer)
	fmt.Printf("Total number of products purchased: %d\n", totalQuantity)
}

// Task 10: Find the most sold product among all.
func FindMostSoldProduct(allProducts []store.Product) {
	if len(allProducts) == 0 {
		fmt.Println("No sold products found.")
		return
	}

	productCount := make(map[string]int)

	for _, product := range allProducts {
		productCount[product.ID]++
	}

	var mostSoldProductID string
	var maxCount int

	for productID, count := range productCount {
		if count > maxCount {
			mostSoldProductID = productID
			maxCount = count
		}
	}

	mostSoldProduct := FindProductByID(allProducts, mostSoldProductID)

	fmt.Println("Most Sold Product among Sold Products:")
	store.PrintProductInfo(mostSoldProduct)
}

// Task 11: Calculate and print the average spending of customers.
func CalculateAndPrintAverageSpending(customers []store.Customer) {
	if len(customers) == 0 {
		fmt.Println("Customer not found.")
		return
	}

	totalSpent := 0.0
	for _, customer := range customers {
		totalSpent += customer.Basket.Total
	}

	averageSpent := totalSpent / float64(len(customers))
	fmt.Printf("Average Total Spending per Customer: %.2f\n", averageSpent)

	topSpender := FindTopSpender(customers)
	fmt.Println("Top Spending Customer:")
	store.PrintCustomerInfo(topSpender)
}

// Task 12: Find the most profitable product category among all customers.
func FindMostProfitableCategory(customers []store.Customer) {
	if len(customers) == 0 {
		fmt.Println("Customer not found.")
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

	fmt.Printf("Most Profitable Category: %s (Total Profit: %.2f)\n", mostProfitableCategory, maxProfit)
}

// Task 13: Find the most expensive purchase made by each customer.
func FindMostExpensivePurchaseByCustomer(customers []store.Customer) {
	if len(customers) == 0 {
		fmt.Println("Customer not found.")
		return
	}

	for _, customer := range customers {
		mostExpensiveProduct := FindMostExpensiveProduct(customer.Basket.Products)

		if mostExpensiveProduct.ID != "" {
			fmt.Printf("%s %s's Most Expensive Purchase:\n", customer.FirstName, customer.LastName)
			store.PrintProductInfo(mostExpensiveProduct)
		} else {
			fmt.Printf("%s %s's Purchase Not Found.\n", customer.FirstName, customer.LastName)
		}
	}
}

// Task 14: Find the category in which each customer spent the most.
func FindMostExpensiveCategoryByCustomer(customers []store.Customer) {
	if len(customers) == 0 {
		fmt.Println("Customer not found.")
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
			fmt.Printf("%s %s's Most Expensive Category: %s\n", customer.FirstName, customer.LastName, mostExpensiveCategory)
			fmt.Printf("Total amount spent in this category: %.2f\n", maxSpending)
		} else {
			fmt.Printf("%s %s's Spending Category Not Found.\n", customer.FirstName, customer.LastName)
		}
	}
}

// Task 15: Print the total quantity sold for each product and overall.
func PrintTotalSoldQuantity(products []store.Product) {
	if len(products) == 0 {
		fmt.Println("Sold products not found.")
		return
	}

	productSoldQuantity := make(map[string]int)
	totalSoldQuantity := 0

	for _, product := range products {
		productSoldQuantity[product.Name] += product.Quantity
		totalSoldQuantity += product.Quantity
	}

	fmt.Println("Total Quantity Sold for Each Product:")
	for productName, quantity := range productSoldQuantity {
		fmt.Printf("%s: %d units\n", productName, quantity)
	}

	fmt.Printf("Total Quantity of Sold Products: %d units\n", totalSoldQuantity)
}

// Helper function: Find the customer who spent the least.
func FindLowestSpender(customers []store.Customer) store.Customer {
	if len(customers) == 0 {
		return store.Customer{}
	}

	lowestSpender := customers[0]

	for _, customer := range customers {
		if customer.Basket.Total < lowestSpender.Basket.Total {
			lowestSpender = customer
		}
	}

	return lowestSpender
}

// Helper function: Find a product by its ID.
func FindProductByID(products []store.Product, productID string) store.Product {
	for _, product := range products {
		if product.ID == productID {
			return product
		}
	}

	return store.Product{}
}