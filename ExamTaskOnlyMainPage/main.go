package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Ürün (Product) struct'ı
type Product struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Alışveriş Sepeti (Basket) struct'ı
type Basket struct {
	ID       string    `json:"id"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}

// Müşteri (Customer) struct'ı
type Customer struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Cash      float64 `json:"cash"`
	Basket    Basket  `json:"basket"`
}

// JSON'dan bilgileri okuma
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

// Müşteri ve alışveriş sepeti bilgilerini yazdırma
func printCustomerInfo(customer Customer) {
	fmt.Printf("Müşteri ID: %s, Ad: %s, Soyad: %s, Nakit Miktarı: %.2f\n",
		customer.ID, customer.FirstName, customer.LastName, customer.Cash)

	// Alışveriş sepetini yazdırma
	for _, product := range customer.Basket.Products {
		fmt.Printf("   Ürün ID: %s, Kategori: %s, Ad: %s, Fiyat: %.2f, Miktar: %d\n",
			product.ID, product.Category, product.Name, product.Price, product.Quantity)
	}

	fmt.Printf("   Toplam Tutar: %.2f\n", customer.Basket.Total)
	fmt.Println("------------------------------")
}

// Tüm müşteri bilgilerini yazdırma
func printAllCustomers(customers []Customer) {
	for _, customer := range customers {
		printCustomerInfo(customer)
	}
}

func main() {
	// JSON dosyasını okuma
	filename := "store_data.json"
	customers, err := readData(filename)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Task 1: Tüm müşteri bilgilerini ve toplam paraları yazdırma
	fmt.Println("Task 1: Tüm müşteri bilgilerini ve toplam paraları yazdırma")
	printCustomerDetails(customers)

	// Task 2: En çok para harcayan müşteriyi bulma ve yazdırma
	fmt.Println("\nTask 2: En çok para harcayan müşteriyi bulma ve yazdırma")
	topSpender := findTopSpender(customers)
	printCustomerInfo(topSpender)

	// Task 3: En pahalı ürünü bulma ve yazdırma
	fmt.Println("\nTask 3: En pahalı ürünü bulma ve yazdırma")
	allProds := allProducts(customers)
	mostExpensiveProduct := findMostExpensiveProduct(allProds)
	printProductInfo(mostExpensiveProduct)

	// Task 4: Tüm ürünlerin ortalama fiyatını hesapla ve görüntüle
	fmt.Println("\nTask 4: Tüm ürünlerin ortalama fiyatını hesapla ve görüntüle")
	calculateAndPrintAveragePrice(allProds)

	// Task 5: En düşük fiyatla toplamda satın alan müşteriyi bulma ve yazdırma
	fmt.Println("\nTask 5: En düşük fiyatla toplamda satın alan müşteriyi bulma ve yazdırma")
	printLowestSpender(customers)

	// Task 6: En çok satan ürün kategorisini belirleme ve yazdırma
	fmt.Println("\nTask 6: En çok satan ürün kategorisini belirleme ve yazdırma")
	bestSellingCategory := findBestSellingCategory(customers)
	fmt.Println("En çok satan ürün kategorisi:", bestSellingCategory)

	// Task 7: En çok ve en az satılan ürünleri yazdırma
    printMinMaxSoldProducts(customers)

	 // Task 8: Her satış için ortalama ürün miktarını hesapla ve görüntüleme
	fmt.Println("\nTask 8:")
	calculateAndPrintAverageQuantitySold(customers)

	  // Task 9: En çok ürün alan müşteriyi ve toplamda kaç ürün aldığını ekrana çıkart
	fmt.Println("\nTask 9: ")
	findTopCustomerByProductQuantity(customers)

	// Task 10: Satılmış ürünlerde en çok görülen ürünü bulma ve yazdırma
	fmt.Println("\nTask 10: ")
	findMostSoldProduct(allProds)

	// Task 11: Müşterilerin her birinin toplam harcadığı paranın ortalamasını çıkar ve en çok para harcayan müşteriyi ekrana çıkar
	fmt.Println("\nTask 11: ")
	calculateAndPrintAverageSpending(customers)

	// Task 12: En çok toplam gelir elde eden kategoriyi ekrana çıkart
	fmt.Println("\nTask 12: ")
	findMostProfitableCategory(customers)

	// Task 13: Her bir kullanıcının satın aldığı en pahalı eşyayı ekrana çıkart
	fmt.Println("\nTask 13: ")
	findMostExpensivePurchaseByCustomer(customers)

	// Task 14: Her bir kullanıcının satın aldığı en çok kategoriyi ve o kategoride ne kadar para harcadığını ekrana çıkart
	fmt.Println("\nTask 14: ")
	findMostExpensiveCategoryByCustomer(customers)

	// Task 15: Her bir üründen toplamda kaç tane satıldığını ve toplamda kaç tane ürün satıldığını ekrana çıkart
	fmt.Println("\nTask 15: ")
	printTotalSoldQuantity(allProds)
}


// Tüm müşteri bilgilerini yazdırma
func printCustomerDetails(customers []Customer) {
	totalCash := 0.0
	totalSpent := 0.0

	for _, customer := range customers {
		printCustomerInfo(customer)
		totalCash += customer.Cash
		totalSpent += customer.Basket.Total
	}

	fmt.Printf("Toplam Müşteri Nakit: %.2f\n", totalCash)
	fmt.Printf("Toplam Harcanan Tutar: %.2f\n", totalSpent)
}

// En çok para harcayan müşteriyi bulma
func findTopSpender(customers []Customer) Customer {
	if len(customers) == 0 {
		// Eğer müşteri yoksa, nil bir müşteri döndür
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

// En pahalı ürünü bulma
func findMostExpensiveProduct(products []Product) Product {
	if len(products) == 0 {
		// Eğer ürün yoksa, nil bir ürün döndür
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

// Tüm ürünleri getirme
func allProducts(customers []Customer) []Product {
    var allProds []Product
    for _, customer := range customers {
        allProds = append(allProds, customer.Basket.Products...)
    }
    return allProds
}



// Ürün bilgilerini yazdırma
func printProductInfo(product Product) {
	fmt.Printf("ID: %s\n", product.ID)
	fmt.Printf("Kategori: %s\n", product.Category)
	fmt.Printf("Ürün adı: %s\n", product.Name)
	fmt.Printf("Fiyat: %.2f\n", product.Price)
	fmt.Printf("Miktar: %d\n", product.Quantity)
	fmt.Println("------------------------------")
}

// Tüm ürünlerin ortalama fiyatını hesapla ve görüntüle
func calculateAndPrintAveragePrice(allProducts []Product) {
    if len(allProducts) == 0 {
        fmt.Println("Ürün bulunamadı.")
        return
    }

    var total float64
    for _, product := range allProducts {
        total += product.Price
    }

    average := total / float64(len(allProducts))
    fmt.Printf("Tüm ürünlerin ortalama fiyatı: %.2f\n", average)
}




// En düşük fiyatla toplamda satın alan müşteriyi bulma
func findLowestSpender(customers []Customer) Customer {
    if len(customers) == 0 {
        // Eğer müşteri yoksa, nil bir müşteri döndür
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

// En düşük fiyatla toplamda satın alan müşteriyi yazdırma
func printLowestSpender(customers []Customer) {
    lowestSpender := findLowestSpender(customers)

    if lowestSpender.ID == "" {
        fmt.Println("Müşteri bulunamadı.")
        return
    }

    fmt.Println("En düşük fiyatla toplamda satın alan müşteri:")
    printCustomerInfo(lowestSpender)
}

// En çok satan ürün kategorisini belirleme
func findBestSellingCategory(customers []Customer) string {
	if len(customers) == 0 {
		// Eğer müşteri yoksa, boş bir kategori döndür
		return ""
	}

	categoryCounts := make(map[string]int)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			categoryCounts[product.Category] += product.Quantity
		}
	}

	// En çok satan kategoriyi bulma
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

// En çok ve en az satılan ürünleri bulma
func findMinMaxSoldProducts(customers []Customer) (Product, Product) {
	allProds := allProducts(customers)
	if len(allProds) == 0 {
		// Eğer ürün yoksa, nil ürünler döndür
		return Product{}, Product{}
	}

	// En çok satılan ve en az satılan ürünleri bulma
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

// En çok ve en az satılan ürünleri yazdırma
func printMinMaxSoldProducts(customers []Customer) {
	maxSold, minSold := findMinMaxSoldProducts(customers)

	if maxSold.ID == "" || minSold.ID == "" {
		fmt.Println("Ürün bulunamadı.")
		return
	}

	fmt.Println("---------------------------------------")
	fmt.Println("Task 7: En çok ve en az satılan ürünler")
	fmt.Println("En çok satılan ürün:")
	printProductInfo(maxSold)
	fmt.Println("En az satılan ürün:")
	printProductInfo(minSold)

}

// Task 8: Her satış için ortalama ürün miktarını hesapla ve görüntüleme
func calculateAndPrintAverageQuantitySold(customers []Customer) {
	totalSales := len(customers)
	totalQuantity := 0

	for _, customer := range customers {
		totalQuantity += len(customer.Basket.Products)
	}

	average := float64(totalQuantity) / float64(totalSales)
	fmt.Printf("Ortalama ürün miktarı: %d / %d = %.3f\n", totalQuantity, totalSales, average)
}


// Sepetteki toplam ürün miktarını hesapla
func calculateBasketTotalQuantity(basket Basket) int {
    totalQuantity := 0
    for _, product := range basket.Products {
        totalQuantity += product.Quantity
    }
    return totalQuantity
}

// Task 9: En çok ürün alan müşteriyi ve toplamda kaç ürün aldığını ekrana çıkart
func findTopCustomerByProductQuantity(customers []Customer) {
    if len(customers) == 0 {
        fmt.Println("Müşteri bulunamadı.")
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

    fmt.Printf("En çok ürün alan müşteri:\n")
    printCustomerInfo(topCustomer)
    fmt.Printf("Toplamda aldığı ürün sayısı: %d\n", totalQuantity)
}

// Task 10: Satılmış ürünlerde en çok görülen ürünü bulma ve yazdırma
func findMostSoldProduct(allProducts []Product) {
	if len(allProducts) == 0 {
		fmt.Println("Satılmış ürün bulunamadı.")
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

	mostSoldProduct := findProductByID(allProducts, mostSoldProductID)

	fmt.Println("Task 10: Satılmış ürünlerde en çok görülen ürün:")
	printProductInfo(mostSoldProduct)
}

// ID'ye göre ürün bulma
func findProductByID(products []Product, productID string) Product {
    for _, product := range products {
        if product.ID == productID {
            return product
        }
    }
    // Eğer ürün bulunamazsa, boş bir ürün döndür
    return Product{}
}

// Task 11: Müşterilerin her birinin toplam harcadığı paranın ortalamasını çıkar ve en çok para harcayan müşteriyi ekrana çıkar
func calculateAndPrintAverageSpending(customers []Customer) {
	if len(customers) == 0 {
		fmt.Println("Müşteri bulunamadı.")
		return
	}

	totalSpent := 0.0
	for _, customer := range customers {
		totalSpent += customer.Basket.Total
	}

	averageSpent := totalSpent / float64(len(customers))
	fmt.Printf("Task 11: Müşterilerin her birinin toplam harcadığı paranın ortalaması: %.2f\n", averageSpent)

	// En çok para harcayan müşteriyi bulma
	topSpender := findTopSpender(customers)
	fmt.Println("En çok para harcayan müşteri:")
	printCustomerInfo(topSpender)
}

// Task 12: En çok toplam gelir elde eden (quantity*price) kategoriyi ekrana çıkart
func findMostProfitableCategory(customers []Customer) {
	if len(customers) == 0 {
		fmt.Println("Müşteri bulunamadı.")
		return
	}

	categoryProfits := make(map[string]float64)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			profit := float64(product.Quantity) * product.Price
			categoryProfits[product.Category] += profit
		}
	}

	// En çok gelir elde edilen kategoriyi bulma
	mostProfitableCategory := ""
	maxProfit := 0.0
	for category, profit := range categoryProfits {
		if profit > maxProfit {
			mostProfitableCategory = category
			maxProfit = profit
		}
	}

	fmt.Printf("Task 12: En çok toplam gelir elde eden kategori: %s (Toplam Gelir: %.2f)\n", mostProfitableCategory, maxProfit)
}

// Task 13: Her bir kullanıcının satın aldığı en pahalı eşyayı ekrana çıkart
func findMostExpensivePurchaseByCustomer(customers []Customer) {
	if len(customers) == 0 {
		fmt.Println("Müşteri bulunamadı.")
		return
	}

	for _, customer := range customers {
		mostExpensiveProduct := findMostExpensiveProduct(customer.Basket.Products)

		if mostExpensiveProduct.ID != "" {
			fmt.Printf("Task 13: %s %s adlı müşterinin satın aldığı en pahalı ürün:\n", customer.FirstName, customer.LastName)
			printProductInfo(mostExpensiveProduct)
		} else {
			fmt.Printf("Task 13: %s %s adlı müşterinin satın aldığı ürün bulunamadı.\n", customer.FirstName, customer.LastName)
		}
	}
}

// Task 14: Her bir kullanıcının satın aldığı en çok kategoriyi ve o kategoride ne kadar para harcadığını ekrana çıkart
func findMostExpensiveCategoryByCustomer(customers []Customer) {
	if len(customers) == 0 {
		fmt.Println("Müşteri bulunamadı.")
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
			fmt.Printf("Task 14: %s %s adlı müşterinin en çok harcama yaptığı kategori: %s\n", customer.FirstName, customer.LastName, mostExpensiveCategory)
			fmt.Printf("Bu kategoride harcanan toplam miktar: %.2f\n", maxSpending)
		} else {
			fmt.Printf("Task 14: %s %s adlı müşterinin harcama yaptığı kategori bulunamadı.\n", customer.FirstName, customer.LastName)
		}
	}
}

// Task 15: Her bir üründen toplamda kaç tane satıldığını ve toplamda kaç tane ürün satıldığını ekrana çıkart
func printTotalSoldQuantity(products []Product) {
	if len(products) == 0 {
		fmt.Println("Satılan ürün bulunamadı.")
		return
	}

	productSoldQuantity := make(map[string]int)
	totalSoldQuantity := 0

	for _, product := range products {
		productSoldQuantity[product.Name] += product.Quantity
		totalSoldQuantity += product.Quantity
	}

	fmt.Println("Task 15: Her bir üründen satılan toplam miktar:")
	for productName, quantity := range productSoldQuantity {
		fmt.Printf("%s: %d adet\n", productName, quantity)
	}

	fmt.Printf("Toplamda satılan ürün miktarı: %d adet\n", totalSoldQuantity)
}
