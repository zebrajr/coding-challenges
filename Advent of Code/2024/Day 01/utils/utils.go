package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// load_input_file opens the specified file and returns a pointer to the os.File.
// Input parameter: file_name (string) - the name of the file to load.
// Output: *os.File - a pointer to the opened file, or nil if an error occurs.
func Load_input_file(file_name string) *os.File {
	// Open the file and assign it to input_file variable
	input_file, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil // Return nil if there was an error opening the file
	}

	return input_file // Return a pointer to the opened file
}

// parse_input_values_from_file reads each line from the provided file pointer and processes its values.
// Input parameter: target_file (*os.File) - a pointer to the file to read from
// Output: two integer arrays
func Parse_input_values_from_file(target_file *os.File) ([]int, []int) {
	list_a := []int{}
	list_b := []int{}

	// Create a new scanner to read from the file
	scanner := bufio.NewScanner(target_file)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		new_a_value, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}

		new_b_value, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		list_a = append(list_a, new_a_value)
		list_b = append(list_b, new_b_value)
	}

	return list_a, list_b
}
