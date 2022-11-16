package tsp

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"tsp/util"
)

func ReadFile(file string) (*TSPInstance, error) {

	fileReader, err := os.Open(file)

	readData := false

	if err != nil {
		return nil, err
	}

	defer fileReader.Close()

	fileScanner := bufio.NewScanner(fileReader)
	fileScanner.Split(bufio.ScanLines)

	var instance TSPInstance

	for fileScanner.Scan() {

		line := fileScanner.Text()

		if readData {
			if strings.Compare(line, "EOF") == 0 {
				break
			}

			data := strings.Split(line, " ")
			var customer TSPData
			customer.id, _ = strconv.Atoi(data[0])
			customer.pointx, _ = strconv.ParseFloat(data[1], 64)
			customer.pointy, _ = strconv.ParseFloat(data[2], 64)
			instance.Data = append(instance.Data, customer)
			continue
		}

		node := strings.Split(line, ":")

		var value string

		if len(node) > 1 {
			value = strings.Trim(node[1], " ")
		}

		switch strings.Trim(node[0], " ") {

		case "NAME":
			instance.name = value
		case "TYPE":
			instance.problemType = value
		case "COMMENT":
			instance.comment = value
		case "DIMENSION":
			instance.Dimension, _ = strconv.Atoi(value)
		case "EDGE_WEIGHT_TYPE":
			instance.edgeType = value

			if instance.edgeType != "EUC_2D" {
				return nil, errors.New("EDGE_WEIGHT_TYPE: " + instance.edgeType + " is not implmented yet.")
			}
		case "NODE_COORD_SECTION":
			readData = true
			continue
		}

	}

	calculateDistanceMatrix(&instance)
	return &instance, nil

}

// From TSPLIB
func nint(value float64) int {
	return int(value + 0.5)
}

func calculateEuc2DDistance(px, py, qx, qy float64) int {

	dx := px - qx
	dy := py - qy

	dx = math.Pow(dx, 2)
	dy = math.Pow(dy, 2)

	distance := nint(math.Sqrt(dx + dy))

	return distance

}

func calculateDistanceMatrix(instance *TSPInstance) {

	// fmt.Println(instance.Data)

	instance.Distance = make([][]int, instance.Dimension)

	for i, ci := range instance.Data {
		for _, cj := range instance.Data {

			distance := calculateEuc2DDistance(ci.pointx, ci.pointy, cj.pointx, cj.pointy)
			instance.Distance[i] = append(instance.Distance[i], distance)

		}
	}

}

func PrintInstance(instance *TSPInstance, isToPrintData bool) {

	fmt.Printf("Instance...: %s\n", instance.name)
	fmt.Printf("Problem....: %s\n", instance.problemType)
	fmt.Printf("Comment....: %s\n", instance.comment)
	fmt.Printf("Dimension..: %d\n", instance.Dimension)
	fmt.Printf("Edge type..: %s\n", instance.edgeType)

	util.PrintLine()

	if isToPrintData {

		fmt.Println("\nData: ")
		fmt.Printf("Id\tPx\tPy\n")
		for _, data := range instance.Data {
			fmt.Printf("%d\t%g\t%g\n", data.id, data.pointx, data.pointy)
		}

		util.PrintLine()

		fmt.Println("\nDistance matrix:")
		for i := range instance.Distance {
			for j := range instance.Distance[i] {
				fmt.Printf("%d\t", instance.Distance[i][j])
			}
			fmt.Println()
		}

		util.PrintLine()

	}

}

// func CalculateCost(tour []int, instance *TSPInstance) (int, error) {

// 	cost := int(math.Inf(0))
// 	var err error

// 	if len(tour) != instance.Dimension {
// 		return cost, fmt.Errorf("the number of customer is different from the instance dimension! Expected: %d - Received: %d", instance.Dimension, len(tour))
// 	}

// 	cost = 0
// 	// Index to Index + 1
// 	for index, item := range tour {
// 		if index < len(tour)-1 {
// 			ci, cj := item, tour[index+1]
// 			cost += instance.Distance[ci-1][cj-1]

// 		}
// 	}

// 	// From last to first
// 	cost += instance.Distance[tour[len(tour)-1]-1][tour[0]-1]
// 	return cost, err

// }

// func validateTour(tour []int, instance *TSPInstance) (bool, error) {

// 	check := make([]int, instance.Dimension)
// 	check[tour[len(tour)-1]-1] = 1

// 	return true, nil

// }
