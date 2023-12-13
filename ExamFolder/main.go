package main

import (
	"ExamFolder/store"
	"ExamFolder/task"
	"fmt"
)

func main() {
	filename := "dataJson/store_data.json"
	customers, err := store.ReadData(filename)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Task 1:")
	task.PrintCustomerDetails(customers)

	fmt.Println("\nTask 2:")
	topSpender := task.FindTopSpender(customers)
	task.PrintCustomerInfo(topSpender)

	fmt.Println("\nTask 3:")
	allProds := task.AllProducts(customers)
	mostExpensiveProduct := task.FindMostExpensiveProduct(allProds)
	task.PrintProductInfo(mostExpensiveProduct)

	fmt.Println("\nTask 4:")
	task.CalculateAndPrintAverageQuantitySold(customers)

	fmt.Println("\nTask 5:")
	task.PrintLowestSpender(customers)

	fmt.Println("\nTask 6:")
	bestSellingCategory := task.FindBestSellingCategory(customers)
	fmt.Println("En çok satan ürün kategorisi:", bestSellingCategory)

	fmt.Println("\nTask 7:")
	maxSold, minSold := task.FindMinMaxSoldProducts(customers)
	task.PrintProductInfo(maxSold)
	task.PrintProductInfo(minSold)

	fmt.Println("\nTask 8:")
	task.CalculateAndPrintAverageQuantitySold(customers)

	fmt.Println("\nTask 9:")
	task.FindTopCustomerByProductQuantity(customers)

	fmt.Println("\nTask 10:")
	task.FindMostSoldProduct(allProds)

	fmt.Println("\nTask 11:")
	task.CalculateAndPrintAverageSpending(customers)

	fmt.Println("\nTask 12:")
	task.FindMostProfitableCategory(customers)

	fmt.Println("\nTask 13:")
	task.FindMostExpensivePurchaseByCustomer(customers)

	fmt.Println("\nTask 14:")
	task.FindMostExpensiveCategoryByCustomer(customers)

	fmt.Println("\nTask 15:")
	task.PrintTotalSoldQuantity(allProds)
}
