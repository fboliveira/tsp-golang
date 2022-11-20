package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"tsp/tsp"
	"tsp/util"
)

func commandLineParameters(arguments []string) tsp.Parameters {

	var parameters tsp.Parameters

	if len(arguments) == 1 {
		// fileName := "./instances/berlin52.tsp"
		parameters.FileName = "./instances/berlin-toy.tsp"
	} else {
		parameters.FileName = arguments[1]
	}

	return parameters

}

// go test ./...

func main() {

	//Provide seed
	rand.Seed(time.Now().Unix())

	parameters := commandLineParameters(os.Args)

	util.PrintHeader()

	instance, err := tsp.ReadFile(parameters.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	tsp.PrintInstance(instance, true)

	solution := tsp.Canonical(instance)
	solution.Print()
	solution.CalculateCost()
	fmt.Printf("Canonical - Revised Cost = %d\n", solution.TotalCost())

	source := solution.List.Front()
	destination := source.Next()

	solution.Swap(source, destination)

	cost := solution.TotalCost()
	calculatedCost := solution.CalculateTotalDistance()

	if cost != calculatedCost {
		fmt.Printf("Cost error at InsertFront on startup: Expected %d\tgot %d", calculatedCost, cost)
	}
	// // Greedy
	// greedy := tsp.PartialGreedy(instance, 0.0)

	// greedy.Print()
	// greedy.CalculateCost()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Greedy - Revised Cost = %d\n", greedy.TotalCost())

	// // Partial Greedy
	// partial_greedy := tsp.PartialGreedy(instance, 0.3)

	// partial_greedy.Print()
	// partial_greedy.CalculateCost()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Partial Greedy - Revised Cost = %d\n", partial_greedy.TotalCost())

}
