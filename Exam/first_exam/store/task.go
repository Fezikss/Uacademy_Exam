package store

import (
	"fmt"
	"math"
)

type Store struct {
	Data []Costumer `json:"costumer"`
}
type Costumer struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Cash      int    `json:"cash"`
	Baskets   Basket `json:"basket"`
}
type Basket struct {
	ID       string `json:"id"`
	Products []Product
	Total    int `json:"total"`
}
type Product struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

//task-1 barcha customerlarni umumiy mablagini va qancha summa harid qilganini hisonlang va korsating
/*
	umumiy mablag: 150,000 +120,00+...=470,000
	qancha summa harid qilgan : 59,000 + 42,000 + 85,000 = 186,000
*/
func (s *Store) Task1() {
	costumers := s.Data
	//fmt.Println(costumers)
	countMap := make(map[string]int)
	basketMap := make(map[string]int)
	sumUmumiy := 0
	sumHarid := 0
	for _, c := range costumers {
		countMap[c.ID] = c.Cash
		basketMap[c.Baskets.ID] = c.Baskets.Total
	}
	for _, s := range costumers {
		sumUmumiy += countMap[s.ID]
		sumHarid += basketMap[s.Baskets.ID]
	}
	fmt.Printf("Umumiy mablag: %d\nQancha summa harid qilgan: %d\n", sumUmumiy, sumHarid)
}

//task-2 eng kop pul sarflagan customerni aniqlang va ko'rsating
/*
	eng kop pul sarf etgan mijoz: Micheal Jordan(200,000 sum)
*/
func (s *Store) Task2() {
	clients := s.Data
	max := 0
	costumerMap := make(map[string]int)
	costumerName := make(map[string]string)
	costumerLastName := make(map[string]string)
	for _, costumer := range clients {
		costumerMap[costumer.ID] = costumer.Baskets.Total
		costumerName[costumer.ID] = costumer.FirstName
		costumerLastName[costumer.ID] = costumer.LastName
	}
	for _, c := range clients {
		if costumerMap[c.ID] > max {
			max = costumerMap[c.ID]
		}
	}
	for c := range costumerMap {
		if max == costumerMap[c] {
			fmt.Printf("Eng ko'p pul sarf etgan mijoz: %s %s (%d sum)", costumerName[c], costumerLastName[c], max)
		}
	}
}

// task-3 eng qimmat mahsulot steak(30,000 sum)
func (s *Store) Task3() {
	data := s.Data
	maxProduct := 0
	productName := make(map[string]string)
	countMap := make(map[string]int)
	for _, c := range data {
		for _, b := range c.Baskets.Products {
			countMap[b.ID] = b.Price
			productName[b.ID] = b.Name
		}
	}
	for _, p := range data {
		for _, b := range p.Baskets.Products {
			if countMap[b.ID] > maxProduct {
				maxProduct = countMap[b.ID]
			}
		}

	}
	for c := range countMap {
		if maxProduct == countMap[c] {
			fmt.Printf("Eng qimmat mahsulot: %s (%d sum)\n", productName[c], maxProduct)
		}
	}

}

// task4 barcha productlar un narxni hisonlang va korsating -> (12+4+5)/3 = ....
func (s *Store) Task4() {
	data := s.Data
	counter := 0
	sum := 0
	priceMap := make(map[string]int)
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			priceMap[p.ID] = p.Price
		}
	}
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			sum += priceMap[p.ID]
			counter++
		}
	}
	// fmt.Println(sum,counter)
	fmt.Println("Ortacha narx: ", float32(sum)/float32(counter))

}

// task5 eng arzon savdo qilgan customerni aniqlangva korsating
func (s *Store) Task5() {
	clients := s.Data
	costumerMap := make(map[string]int)
	costumerName := make(map[string]string)
	costumerLastName := make(map[string]string)
	for _, costumer := range clients {
		costumerMap[costumer.ID] = costumer.Baskets.Total
		costumerName[costumer.ID] = costumer.FirstName
		costumerLastName[costumer.ID] = costumer.LastName
	}
	min := math.MaxInt32
	for _, c := range clients {
		if costumerMap[c.ID] < min {
			min = costumerMap[c.ID]
		}
	}
	for c := range costumerMap {
		if min == costumerMap[c] {
			fmt.Printf("Eng kam pul sarf etgan mijoz: %s %s(%d sum)", costumerName[c], costumerLastName[c], min)
		}
	}
}

// task6 end ko'p sotilgan productlar categoriyasini aniqlang va chiqaring
func (s *Store) Task6() {
	data := s.Data
	category := make(map[string]string)
	quantity := make(map[string]int)
	maxSale := 0
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			category[p.ID] = p.Category
			quantity[p.ID] = p.Quantity
		}
	}
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			if quantity[p.ID] > maxSale {
				maxSale = quantity[p.ID]
			}
		}
	}
	for p := range quantity {
		if maxSale == quantity[p] {
			fmt.Println("Eng ko'p sotilgan categoriya:", category[p])
		}
	}

}

//task7 eng kop va eng kam sotilgan productlar nomini chiqarish
/*
	eng kop sotilgan mahsulot nomi: carrot (micheal jordan - 5ta)
	eng kam sotilgan mahsulot nomi: apple (dwayne johnson - 1ta)
*/
func (s *Store) Task7() {
	data := s.Data
	quantity := make(map[string]int)
	min := math.MaxInt32
	max := math.MinInt32
	clientName := make(map[string]string)
	clientLastName := make(map[string]string)
	productName := make(map[string]string)
	for _, d := range data {
		for _, b := range d.Baskets.Products {
			quantity[b.ID] = b.Quantity
			productName[b.ID] = b.Name
			clientName[b.ID] = d.FirstName
			clientLastName[b.ID] = d.LastName
		}
	}
	for _, d := range data {
		for _, b := range d.Baskets.Products {
			if max < quantity[b.ID] {
				max = quantity[b.ID]
			}
			if min > quantity[b.ID] {
				min = quantity[b.ID]
			}
		}
	}
	for i := range quantity {
		if max == quantity[i] {
			fmt.Printf("Eng kop sotilgan mahsulot nomi:%s (%s %s - %dta)\n", productName[i], clientName[i], clientLastName[i], max)
		}
	}
	for i := range quantity {
		count := 0
		if min == quantity[i] {
			count++
			if count == 1 {
				fmt.Printf("Eng kam sotilgan mahsulot nomi:%s (%s %s - %dta)\n", productName[i], clientName[i], clientLastName[i], min)
				break
			}
		}

	}

}

// task8 har bir savdo ucun ortacha mahsulot miqdorini hisoblang
func (s *Store) Task8() {
	data := s.Data
	counter := 0
	sum := 0
	priceMap := make(map[string]int)
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			priceMap[p.ID] = p.Quantity
		}
	}
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			sum += priceMap[p.ID]
			counter++
		}
	}
	//fmt.Println(sum,counter)
	fmt.Printf("Ortacha : %.3f\n", float32(sum)/float32(counter))
}

//task9 eng kop mahsulot mahsulot sotb olgan mijozni aniqlang
/*
	eng kop mahsulot mahsulot sotb olgan mijoz: Micheal Jordan (8ta mahsulot)
*/
func (s *Store) Task9() {
	custumerdata := s.Data
	engKoP := 0
	var customerName string
	var customerLastName string
	for _, customer := range custumerdata {
		totalProducts := 0
		for _, products := range customer.Baskets.Products {
			totalProducts += products.Quantity
			if totalProducts > engKoP {
				engKoP = totalProducts
				customerName = customer.FirstName
				customerLastName = customer.LastName
			}
		}
	}
	fmt.Printf("Eng ko'p mahsulot sotib olgan mijoz: %s %s (%d ta)\n", customerName, customerLastName, engKoP)

}

func (s *Store) Task10() {
	custumerdata := s.Data
	MahsulotSoni := make(map[string]int)

	for _, customer := range custumerdata {
		for _, product := range customer.Baskets.Products {
			MahsulotSoni[product.Name]++
		}
	}
	var KopkoringanMahsulotNomi string
	KopkoringanMahsulot := 0
	for productName, frequency := range MahsulotSoni {
		if frequency > KopkoringanMahsulot {
			KopkoringanMahsulot = frequency
			KopkoringanMahsulotNomi = productName
		}
	}

	if KopkoringanMahsulot > 0 {
		fmt.Printf("Eng ko'p savdolarda ko'rinadigan mahsulot: %s (%d ta)\n", KopkoringanMahsulotNomi, KopkoringanMahsulot)
	} else {
		fmt.Println("Mijozlar savatchalarida hech qanday mahsulot topilmadi.")
	}

}

func (s *Store) Task11() {
	custumerData := s.Data

	maxOrtachaXarajat := 0
	var engKopXarajatMijozIsmi, engKopXarajatMijozFamiliyasi string

	for _, mijoz := range custumerData {
		jamiXarajat := 0
		for _, mahsulot := range mijoz.Baskets.Products {
			jamiXarajat += mahsulot.Price * mahsulot.Quantity
		}

		ortachaXarajat := jamiXarajat / len(mijoz.Baskets.Products)
		if ortachaXarajat > maxOrtachaXarajat {
			maxOrtachaXarajat = ortachaXarajat
			engKopXarajatMijozIsmi = mijoz.FirstName
			engKopXarajatMijozFamiliyasi = mijoz.LastName
		}
	}

	fmt.Printf("O'rtacha savdo mablag'i: %d so'm\n", maxOrtachaXarajat)
	fmt.Printf("Eng kop xarajatga ega bo'lgan mijoz: %s %s\n", engKopXarajatMijozIsmi, engKopXarajatMijozFamiliyasi)

}

func (s *Store) Task12() {
	custumerData := s.Data
	daromadMap := make(map[string]int)

	for _, customer := range custumerData {

		for _, product := range customer.Baskets.Products {
			daromad := product.Quantity * product.Price
			daromadMap[product.Category] += daromad
		}
	}

	maxDaromad := 0
	var Kategoriyasi string
	for category, daromad := range daromadMap {
		if daromad > maxDaromad {
			maxDaromad = daromad
			Kategoriyasi = category
		}
	}

	fmt.Printf("Eng ko'p daromad olishgan toifa: %s (%d so'm)\n", Kategoriyasi, maxDaromad)

}

func (s *Store) Task13() {
	custumerData := s.Data
	var topProduct string
	var customerName string
	var customerLastName string
	for _, customer := range custumerData {
		maxPrice := 0
		for _, product := range customer.Baskets.Products {
			if product.Price > maxPrice {
				maxPrice = product.Price
				topProduct = product.Name
				customerName = customer.FirstName
				customerLastName = customer.LastName
			}

		}
		fmt.Printf("Eng qimmat mahsulot: %s (%s %s - %d) \n", topProduct, customerName, customerLastName, maxPrice)
	}

}

func (s *Store) Task14() {
	custumerData := s.Data
	var topCustomerName, topCustomerLastName, topProductName string

	for _, customer := range custumerData {
		maxTotalSpending := 0
		totalSpending := 0
		var productCategory string
		totalPrice := 0
		for _, product := range customer.Baskets.Products {
			spending := product.Quantity
			totalSpending += spending
			if spending > maxTotalSpending {
				totalPrice = product.Quantity * product.Price
				maxTotalSpending = spending
				topCustomerName = customer.FirstName
				topCustomerLastName = customer.LastName
				productCategory = product.Category
			}
		}
		topProductName = productCategory
		fmt.Printf("Eng ko'p xarid qilgan toifa: %s %s (%s - %d)\n", topCustomerName, topCustomerLastName, topProductName, totalPrice)
	}

}

func (s *Store) Task15() {
	custumerData := s.Data
	sum := 0
	for _, customer := range custumerData {
		for _, product := range customer.Baskets.Products {
			sum += product.Quantity
		}
	}
	fmt.Printf("Mahsulotlar umumiy %d ta sotilgan ", sum)

}
