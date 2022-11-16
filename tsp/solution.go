package tsp

import (
	"container/list"
	"fmt"
	"tsp/util"
)

type Solution struct {
	list     list.List
	cost     int
	instance *TSPInstance
}

func (solution *Solution) TotalCost() int {
	if solution.Len() == 0 {
		solution.cost = 0
	}
	fmt.Printf("Total cost = %d\n", solution.cost)
	return solution.cost
}

func (solution *Solution) distance(customerI, customerJ int) int {
	fmt.Printf("Distance from [%d] to [%d] = %d\n", customerI, customerJ, solution.instance.Distance[customerI-1][customerJ-1])
	return solution.instance.Distance[customerI-1][customerJ-1]
}

// Create a new/empty tour
func (solution *Solution) Init(instance *TSPInstance) *Solution {
	solution.list.Init()
	solution.instance = instance
	solution.cost = 0
	return solution
}

func (solution *Solution) Len() int {
	return solution.list.Len()
}

func (solution *Solution) InsertAfter(customer int, item *list.Element) {

	if item == solution.list.Front() {
		solution.InsertFront(customer)
	}

	if item == solution.list.Back() {
		solution.InsertBack(customer)
	}

	// Front -> ... -> item -> afterItem -> Last (-> Front)
	// Remove cost from item to afterItem
	deltaCost := solution.cost - solution.distance(item.Value.(int), item.Next().Value.(int))

	new := solution.list.InsertAfter(customer, item)

	// Front -> ... -> item -> New -> afterNew (previous afterItem) -> Last (-> Front)
	// Add Cost from item to New
	deltaCost += solution.distance(item.Value.(int), customer)
	// Add Cost from New to afterItem
	deltaCost += solution.distance(new.Value.(int), new.Next().Value.(int))

	solution.cost = deltaCost

	// return solution

}

func (solution *Solution) InsertBack(customer int) {

	deltaCost := 0

	if solution.Len() > 0 {
		// Front -> ... -> Last (-> Front)
		// Remove cost from last to current front
		deltaCost = solution.TotalCost() - solution.distance(solution.list.Back().Value.(int), solution.list.Front().Value.(int))
		fmt.Println("Remove cost from last to current front: ", deltaCost, solution.list.Back().Value.(int), solution.list.Front().Value.(int))
	}

	// Front -> ... -> New (-> Front)
	solution.list.PushBack(customer)
	deltaCost += solution.distance(solution.list.Back().Value.(int), solution.list.Front().Value.(int))

	fmt.Printf("InsertBack->DeltaCost = %d\tCustomer = %d\n", deltaCost, customer)
	printList(&solution.list)
	solution.cost = deltaCost
	solution.TotalCost()

}

func (solution *Solution) InsertFront(customer int) *Solution {

	deltaCost := 0

	if solution.Len() == 0 {
		solution.list.PushBack(customer)
	} else {
		// Front -> ... -> Last (-> Front)
		// Remove cost from last to current front
		deltaCost = solution.cost - solution.distance(solution.list.Back().Value.(int), solution.list.Front().Value.(int))

		// New -> Front -> ... -> Last (-> New)
		solution.list.InsertBefore(customer, solution.list.Front())

	}

	deltaCost += solution.distance(solution.list.Back().Value.(int), solution.list.Front().Value.(int))

	solution.cost = deltaCost

	return solution
}

func (solution *Solution) CalculateCost() (*Solution, error) {

	var err error

	if solution.Len() != solution.instance.Dimension {
		return nil, fmt.Errorf("the number of customer is different from the instance dimension! Expected: %d - Received: %d", solution.instance.Dimension, solution.Len())
	}

	cost := 0
	// Index to Index + 1
	for customer := solution.list.Front(); customer != nil; customer = customer.Next() {
		if customer.Next() != nil {
			cost += solution.distance(customer.Value.(int), customer.Next().Value.(int))
		}
	}

	// From last to first
	cost += solution.distance(solution.list.Back().Value.(int), solution.list.Front().Value.(int))

	solution.cost = cost

	return solution, err

}

func (solution *Solution) Print() {

	util.PrintLine()

	for customer := solution.list.Front(); customer != nil; customer = customer.Next() {
		if customer.Next() != nil {
			fmt.Printf("[%d] -> ", customer.Value.(int))
		}
	}

	fmt.Printf("[%d]\n", solution.list.Back().Value.(int))
	fmt.Printf("Total cost = %d\n", solution.TotalCost())

	util.PrintLine()

}

func printList(list *list.List) {

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d\t", e.Value)
	}
	fmt.Println()
}
