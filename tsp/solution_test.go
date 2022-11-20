package tsp

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSolutionOperations(t *testing.T) {

	// Test instance read and calculate with literature validation
	fileName := "../instances/pcb442.tsp"
	// fileName := "../instances/berlin-toy.tsp"
	// literature_canonical_cost := 221440

	instance, err := ReadFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	var solution Solution
	solution.Init(instance)

	customers := rand.Perm(instance.Dimension)

	solution.InsertFront(customers[0] + 1)
	customers = customers[1:]

	cost := solution.TotalCost()
	calculatedCost := solution.CalculateTotalDistance()

	if cost != calculatedCost {
		t.Fatalf("Cost error at InsertFront on startup: Expected %d\tgot %d", calculatedCost, cost)
	}

	for len(customers) > 0 {

		// Define operation
		operation := rand.Intn(3)
		var name string

		if solution.Len() < 3 {
			operation = 1
		}

		switch operation {
		case 0:
			solution.InsertFront(customers[0] + 1)
			name = "InsertFront"
		case 1:
			solution.InsertBack(customers[0] + 1)
			name = "InsertBack"
		default:
			item := solution.List.Front()
			move := rand.Intn(solution.Len() - 2)

			for i := 0; i < move; i++ {
				item = item.Next()
			}

			solution.InsertAfter(customers[0]+1, item)
			name = "InsertAfter " + strconv.Itoa(item.Value.(int))

		}

		customers = customers[1:]

		cost = solution.TotalCost()
		calculatedCost = solution.CalculateTotalDistance()

		// t.Logf("Cost = %s: Expected %d\tgot %d", name, calculatedCost, cost)

		if cost != calculatedCost {
			t.Fatalf("Cost error at %s: Expected %d\tgot %d", name, calculatedCost, cost)
		}

	}

}

func TestCostToSwap(t *testing.T) {

	fileName := "../instances/berlin-toy.tsp"
	instance, err := ReadFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	solution := Canonical(instance)

	for i := 0; i < instance.Dimension; i++ {

		// source_advance := rand.Intn(instance.Dimension)
		// destination_advance := source_advance

		// for source_advance == destination_advance {
		// 	destination_advance = rand.Intn(instance.Dimension)
		// }

		source := solution.List.Front()
		destination := source.Next()

		// for i := 0; i < source_advance; i++ {
		// 	source = source.Next()
		// }

		// for i := 0; i < destination_advance; i++ {
		// 	destination = destination.Next()
		// }

		t.Logf("Cost before = %d", solution.TotalCost())
		t.Logf("Source: %d\tDestination: %d", source.Value.(int), destination.Value.(int))

		solution.Swap(source, destination)
		costAfterSwap := solution.TotalCost()

		solution.CalculateCost()

		if costAfterSwap != solution.TotalCost() {
			solution.Print()
			t.Fatalf("Swap cost wrong - expected: %d\tgot: %d", solution.TotalCost(), costAfterSwap)
		}

	}

	for item := solution.List.Front(); item != nil; item = item.Next() {

	}

}
