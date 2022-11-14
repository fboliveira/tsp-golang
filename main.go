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

	tsp.PrintInstance(instance, false)

	// Greedy
	greedy := tsp.PartialGreedy(instance, 0.0)

	fmt.Println(greedy.Tour)
	fmt.Printf("Greedy Cost = %d\n", greedy.Cost)

	greedy.Cost, err = tsp.CalculateCost(greedy.Tour, instance)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Greedy - Revised Cost = %d\n", greedy.Cost)

	// Partial Greedy
	partial_greedy := tsp.PartialGreedy(instance, 0.3)

	fmt.Println(partial_greedy.Tour)
	fmt.Printf("Partial Greedy Cost = %d\n", partial_greedy.Cost)

	partial_greedy.Cost, err = tsp.CalculateCost(partial_greedy.Tour, instance)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Partial Greedy - Revised Cost = %d\n", partial_greedy.Cost)

}
