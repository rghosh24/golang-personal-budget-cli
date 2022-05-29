package main

import (
	"fmt"
	m2 "personal-budget/module2"
	"time"
)

var months = []time.Month{
	time.January,
	time.February,
	time.March,
	time.April,
	time.May,
	time.June,
	time.July,
	time.August,
	time.September,
	time.October,
	time.November,
	time.December,
	time.January,
}

func main() {
	bu, _ := m2.CreateBudget(time.January, 1000)
	bu.AddItem("bananas", 10.0)

	fmt.Println("Items in January:", len(bu.Items))
	fmt.Printf("Current cost for January: $%.2f \n", bu.CurrentCost())

	m2.CreateBudget(time.February, 1000)

	bu2 := m2.GetBudget(time.February)
	bu2.AddItem("bananas", 10.0)
	bu2.AddItem("coffee", 3.99)
	bu2.AddItem("gym", 50.0)
	bu2.RemoveItem("coffee")
	fmt.Println("Items in February:", len(bu2.Items))
	fmt.Printf("Current cost for February: $%.2f \n", bu2.CurrentCost())

	m2.CreateBudget(time.March, 300)
	bu3 := m2.GetBudget(time.March)
	bu3.AddItem("chipotle", 30.0)
	bu3.AddItem("ktp", 20.0)
	fmt.Println("Items in March:", len(bu3.Items))
	fmt.Printf("Current cost for March: $%.2f \n", bu3.CurrentCost())

	m2.CreateBudget(time.April, 200)
	bu4 := m2.GetBudget(time.April)
	bu4.AddItem("none", 0.00)
	fmt.Println("Items in April:", 0)
	fmt.Printf("Current cost for April: $%.2f \n", bu4.CurrentCost())

	m2.CreateBudget(time.May, 1000)
	bu5 := m2.GetBudget(time.May)
	bu5.AddItem("gym", 10.30)
	bu5.AddItem("coke", 2.00)
	bu5.AddItem("vitamin water", 1.40)
	fmt.Println("Items in May:", len(bu5.Items))
	fmt.Printf("Current cost for May: $%.2f \n", bu5.CurrentCost())

	m2.CreateBudget(time.June, 200)
	bu6 := m2.GetBudget(time.June)
	bu6.AddItem("new", 0.00)
	fmt.Println("Items in June:", len(bu6.Items))
	fmt.Printf("Current cost for June: $%.2f \n", bu6.CurrentCost())

	fmt.Println("Resetting Budget Report...")
	m2.InitializeReport()

	for _, month := range months {
		_, err := m2.CreateBudget(month, 100.00)
		if err == nil {
			fmt.Println("Budget created for", month)
		} else {
			fmt.Println("Error creating budget:", err)
		}
	}

	_, err := m2.CreateBudget(time.December, 100.00)
	fmt.Println(err)
}
