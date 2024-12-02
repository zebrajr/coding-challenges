package main

import (
	"fmt"
	"sort"

	utils "advent_of_code/utils"
)

const input_file_name = "puzzle_input"

// Input is 2 arrays of lists of equal size
// Output is an array with each position distance AND the total of the distances
func calculate_total_distance_between_lists(
	list_a []int,
	list_b []int,
) int {

	far_apart_list := []int{}
	total_far_apart := 0

	for index, element := range list_a {
		_ = element
		difference := list_b[index] - list_a[index]
		if difference < 0 {
			difference = difference * -1
		}

		fmt.Println(list_a[index], list_b[index], difference)
		far_apart_list = append(far_apart_list, difference)
		total_far_apart += difference
	}
	return total_far_apart
}

func main() {
	input_file := utils.Load_input_file(input_file_name)
	// Ensure the file is closed afterwards to prevent resource leaks
	defer input_file.Close()

	list_a, list_b := utils.Parse_input_values_from_file(input_file)

	sort.Ints(list_a)
	sort.Ints(list_b)

	total_far_apart := calculate_total_distance_between_lists(
		list_a,
		list_b,
	)

	fmt.Println("Total Far Apart:", total_far_apart)

}
