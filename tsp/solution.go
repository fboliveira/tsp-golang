package tsp

import (
	"container/list"
	"fmt"
	"tsp/util"
)

type Item struct {
	*list.Element
}

type Solution struct {
	List     list.List
	cost     int
	instance *TSPInstance
}

func (solution *Solution) TotalCost() int {
	if solution.Len() == 0 {
		solution.cost = 0
	}
	return solution.cost
}

func (solution *Solution) distance(customerI, customerJ int) int {
	return solution.instance.Distance[customerI-1][customerJ-1]
}

// Create a new/empty tour
func (solution *Solution) Init(instance *TSPInstance) *Solution {
	solution.List.Init()
	solution.instance = instance
	solution.cost = 0
	return solution
}

func (solution *Solution) Len() int {
	return solution.List.Len()
}

func CustomerValue(item *list.Element) int {
	return item.Value.(int)
}

func CustomerPosition(item *list.Element) int {
	return CustomerValue(item) - 1
}

func (solution *Solution) InsertAfter(customer int, item *list.Element) *Solution {

	if item == solution.List.Front() {
		solution.InsertFront(customer)
	}

	if item == solution.List.Back() {
		solution.InsertBack(customer)
	}

	// Front -> ... -> item -> afterItem -> Last (-> Front)
	// Remove cost from item to afterItem
	deltaCost := solution.cost - solution.distance(item.Value.(int), item.Next().Value.(int))

	new := solution.List.InsertAfter(customer, item)

	// Front -> ... -> item -> New -> afterNew (previous afterItem) -> Last (-> Front)
	// Add Cost from item to New
	deltaCost += solution.distance(item.Value.(int), customer)
	// Add Cost from New to afterItem
	deltaCost += solution.distance(new.Value.(int), new.Next().Value.(int))

	solution.cost = deltaCost

	return solution

}

func (solution *Solution) InsertBack(customer int) *Solution {

	deltaCost := solution.TotalCost()

	if solution.Len() > 0 {
		// Front -> ... -> Last (-> Front)
		// Remove cost from last to current front
		deltaCost -= solution.distance(solution.List.Back().Value.(int), solution.List.Front().Value.(int))
	}

	// Front -> ... -> Previous Back -> New (-> Front)
	solution.List.PushBack(customer)

	if solution.List.Back().Prev() != nil {
		deltaCost += solution.distance(solution.List.Back().Prev().Value.(int), solution.List.Back().Value.(int))
	}

	deltaCost += solution.distance(solution.List.Back().Value.(int), solution.List.Front().Value.(int))

	solution.cost = deltaCost

	return solution

}

func (solution *Solution) InsertFront(customer int) *Solution {

	deltaCost := solution.TotalCost()

	if solution.Len() == 0 {
		solution.List.PushBack(customer)
	} else {
		// Front -> ... -> Last (-> Front)
		// Remove cost from last to current front
		deltaCost -= solution.distance(solution.List.Back().Value.(int), solution.List.Front().Value.(int))

		// New -> Front -> ... -> Last (-> New)
		solution.List.InsertBefore(customer, solution.List.Front())

	}

	if solution.List.Front().Next() != nil {
		// New -> previous front -> ... -> Last
		deltaCost += solution.distance(solution.List.Front().Value.(int), solution.List.Front().Next().Value.(int))
	}

	deltaCost += solution.distance(solution.List.Back().Value.(int), solution.List.Front().Value.(int))

	solution.cost = deltaCost

	return solution
}

func (solution *Solution) CalculateTotalDistance() int {

	cost := 0
	// Index to Index + 1
	for customer := solution.List.Front(); customer != nil; customer = customer.Next() {
		if customer.Next() != nil {
			cost += solution.distance(customer.Value.(int), customer.Next().Value.(int))
		}
	}

	// From last to first
	cost += solution.distance(solution.List.Back().Value.(int), solution.List.Front().Value.(int))

	return cost

}

func (solution *Solution) CalculateCost() (*Solution, error) {

	var err error

	if solution.Len() != solution.instance.Dimension {
		return nil, fmt.Errorf("the number of customer is different from the instance dimension! Expected: %d - Received: %d", solution.instance.Dimension, solution.Len())
	}

	solution.cost = solution.CalculateTotalDistance()

	return solution, err

}

func (solution *Solution) Print() {

	util.PrintLine()

	for customer := solution.List.Front(); customer != nil; customer = customer.Next() {
		if customer.Next() != nil {
			fmt.Printf("[%d] -> ", customer.Value.(int))
		}
	}

	fmt.Printf("[%d]\n", solution.List.Back().Value.(int))
	fmt.Printf("Total cost = %d\n", solution.TotalCost())

	util.PrintLine()

}

func printList(list *list.List) {

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d\t", e.Value)
	}
	fmt.Println()
}

func (solution *Solution) CostToSwap(source *list.Element, destination *list.Element) int {
	if source == destination {
		return 0
	}

	deltaCost := 0

	if source == solution.List.Front() {
		// Remove distance from last to front
		deltaCost -= solution.distance(CustomerValue(solution.List.Back()), CustomerValue(solution.List.Front()))
		// Add distance from last to destination (as new front)
		deltaCost += solution.distance(CustomerValue(solution.List.Back()), CustomerValue(destination))
	}

	if destination == solution.List.Back() {
		// Remove distance from destionation to front
		deltaCost -= solution.distance(CustomerValue(destination), CustomerValue(solution.List.Front()))
		// Add distance from last to source (as new front)
		deltaCost += solution.distance(CustomerValue(solution.List.Back()), CustomerValue(source))
	}

	if source.Next() == destination {
		// [source] -> [destination]
		deltaCost -= solution.distance(CustomerValue(source), CustomerValue(destination))
	}

	if source.Prev() != nil {
		// Remove [previous source] -> [source]
		deltaCost -= solution.distance(CustomerValue(source.Prev()), CustomerValue(source))
		// Add [previous source] -> [destination]
		deltaCost += solution.distance(CustomerValue(source.Prev()), CustomerValue(destination))
	}

	if destination.Next() != nil {
		// Remove [destination] -> [next destination]
		deltaCost -= solution.distance(CustomerValue(destination), CustomerValue(destination.Next()))
		// Add [source] -> [next destination]
		deltaCost += solution.distance(CustomerValue(source), CustomerValue(destination.Next()))
	}

	return deltaCost

}

func (solution *Solution) Swap(source *list.Element, destination *list.Element) *Solution {

	deltaCost := solution.CostToSwap(source, destination)

	item := source.Value
	source.Value = destination.Value
	destination.Value = item

	solution.cost += deltaCost

	return solution

}
