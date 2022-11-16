package tsp

import (
	"container/list"
	"math"
	"math/rand"
)

func removeFromList(list *list.List, value int) {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == value {
			// fmt.Printf("Removing %d\n", value)
			list.Remove(e)
			break
		}
	}
}

func PartialGreedy(instance *TSPInstance, alpha float64) Solution {

	var solution Solution
	solution.Init(instance)

	random_visit_ordem := rand.Perm(instance.Dimension)
	customers := list.New()

	for i := 0; i < instance.Dimension; i++ {
		customers.PushBack(random_visit_ordem[i] + 1)
	}

	// printList(customers)
	// solution.Tour = append(solution.Tour, customers.Front().Value.(int))
	solution.InsertFront(customers.Front().Value.(int))

	last_customer_index := customers.Front().Value.(int) - 1
	customers.Remove(customers.Front())

	// printList(customers)

	for customers.Len() > 0 {

		distance_min := math.Inf(0)
		distance_max := math.Inf(-1)

		for customer := customers.Front(); customer != nil; customer = customer.Next() {

			customer_index := customer.Value.(int) - 1

			if float64(instance.Distance[last_customer_index][customer_index]) > distance_max {
				distance_max = float64(instance.Distance[last_customer_index][customer_index])
			}

			if float64(instance.Distance[last_customer_index][customer_index]) < distance_min {
				distance_min = float64(instance.Distance[last_customer_index][customer_index])
			}

		}

		// fmt.Println(distance_min, distance_max)

		var restricted_candidates_list []int

		for customer := customers.Front(); customer != nil; customer = customer.Next() {

			customer_index := customer.Value.(int) - 1

			if float64(instance.Distance[last_customer_index][customer_index]) <= distance_min+alpha*(distance_max-distance_min) {
				restricted_candidates_list = append(restricted_candidates_list, customer_index)
			}
		}

		// fmt.Println("Candidate: ")
		// fmt.Println(restricted_candidates_list)

		candidate := restricted_candidates_list[rand.Intn(len(restricted_candidates_list))]

		removeFromList(customers, candidate+1)
		// printList(customers)

		// solution.Tour  = append(solution.Tour, candidate+1)
		solution.InsertBack(candidate + 1)
		// cost_to_add := instance.Distance[last_customer_index][candidate]
		// solution.Cost += cost_to_add

		// fmt.Printf("Cost Added = %.2f\tCost = %.2f\n", cost_to_add, solution.Cost)

		last_customer_index = candidate

	}

	// first := solution.Tour[0] - 1
	// last := solution.Tour[len(solution.Tour)-1] - 1
	// solution.Cost += instance.Distance[last][first]

	return solution

}

// TSPLIB instance canonical solution to verify distance calculating
func Canonical(instance *TSPInstance) Solution {

	var solution Solution
	solution.Init(instance)

	for i := 1; i <= instance.Dimension; i++ {
		solution.InsertBack(i)
	}

	return solution

}
