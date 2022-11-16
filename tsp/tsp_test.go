package tsp

import (
	"testing"
)

func TestCalculateEuc2DDistance(t *testing.T) {
	distance := calculateEuc2DDistance(565, 575, 25, 185)

	if distance != 666 {
		t.Log("Distance should be 666, but got ", distance)
		t.Fail()
	}

}

func TestInstances(t *testing.T) {

	// Test instance read and calculate with literature validation
	fileName := "../instances/pcb442.tsp"
	literature_canonical_cost := 221440

	instance, err := ReadFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	solution := Canonical(instance)
	solution.Print()

	if solution.TotalCost() != literature_canonical_cost {
		t.Fatal("[TestInstances] Distance should be ", literature_canonical_cost, ", but got ", solution.TotalCost())
	}

}
