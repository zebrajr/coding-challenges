package main

import (
	"fmt"
	"sort"

	utils "advent_of_code/utils"
)

const INPUT_FILE_NAME = "puzzle_input"
const FULL_VERBOSE = false

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

		if FULL_VERBOSE {
			fmt.Println(list_a[index], list_b[index], difference)
		}
		far_apart_list = append(far_apart_list, difference)
		total_far_apart += difference
	}

	fmt.Println("Total Far Apart:", total_far_apart)
	return total_far_apart
}

// Outputs a map of [int keys] to int values
func create_occurence_map_from_list(
	values_list []int,
) map[int]int {
	occurence_map := make(map[int]int)
	for _, number := range values_list {
		occurence_map[number]++
	}

	return occurence_map
}

func calculate_similarity_score(
	list_a []int,
	similarity_score_map map[int]int,
) {

	total_similarity_score := 0
	for _, element := range list_a {
		calculated_similarity_score := element * similarity_score_map[element]
		if FULL_VERBOSE {
			fmt.Println(calculated_similarity_score)
		}
		total_similarity_score += calculated_similarity_score
	}
	fmt.Println("Total Similarity Score:", total_similarity_score)
}

func main() {
	input_file := utils.Load_input_file(INPUT_FILE_NAME)
	// Ensure the file is closed afterwards to prevent resource leaks
	defer input_file.Close()

	list_a, list_b := utils.Parse_input_values_from_file(input_file)

	sort.Ints(list_a)
	sort.Ints(list_b)

	calculate_total_distance_between_lists(
		list_a,
		list_b,
	)

	occurence_map := create_occurence_map_from_list(list_b)
	calculate_similarity_score(list_a, occurence_map)

}
