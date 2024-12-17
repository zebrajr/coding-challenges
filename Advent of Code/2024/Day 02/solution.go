package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var puzzle_input = "puzzle_input"
var min_difference = 1
var max_difference = 3

func is_input_complete_safe(line_values []string) bool{
    is_complete_safe := true

    direction := "none"
    var previous_value *int
    for _, level := range line_values{
        // convert to integer
        difference := 0
        new_direction := "none"
        ilevel, err := strconv.Atoi(level)
        if err != nil {
            log.Fatal("Error converting value to integer: ", err)
        }

        // first case of running
        if previous_value == nil {
            previous_value = &ilevel
            continue
        }

        if ilevel > *previous_value {
            difference = ilevel - (*previous_value) 
            if direction == "none"{
                direction = "up"
            }
            new_direction = "up"
        }

        if ilevel < *previous_value {
            difference = *previous_value - ilevel
            if direction == "none"{
                direction = "down"
            }
            new_direction = "down"
        }

        if difference < min_difference {
            is_complete_safe = false
        }

        if difference > max_difference {
            is_complete_safe = false
        }

        if direction != new_direction {
            is_complete_safe = false
        }

        *previous_value = ilevel
    }
    return is_complete_safe
}

func is_safe_with_level_removed(line_values []string) bool{
    is_safe := false

    // go though each of the line values's elements
    for i, _ := range line_values{
        // make a new array with all values except the current chosen element
        var values_to_test []string
        for j, value := range line_values{
            if j != i{
               values_to_test = append(values_to_test, value) 
            }
        }
        
        // now check if the new array is valid
        if is_input_complete_safe(values_to_test) {
            is_safe = true
        }

        // return earlier if found
        if is_safe {
            return is_safe
        }
    }
    return is_safe
}

func main(){
    file, err := os.Open(puzzle_input)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    file_scanner := bufio.NewScanner(file)

    count_safe_reports := 0
    count_safe_reports_with_bad_levels := 0
    for file_scanner.Scan(){
        current_line := file_scanner.Text()
        line_values := strings.Fields(current_line)

        is_complete_safe := is_input_complete_safe(line_values)

        if is_complete_safe == true {
            count_safe_reports++
        }

        if is_complete_safe == false {
            safe_with_bad_level := is_safe_with_level_removed(line_values)
            if safe_with_bad_level {
                count_safe_reports_with_bad_levels++
            }
        }
    }

    fmt.Println("Safe Reports found: ", count_safe_reports)
    fmt.Println("Safe Reports with Bad Levels found: ", count_safe_reports_with_bad_levels)
    fmt.Println("Safe + Safe with Bad Levels found: ", (count_safe_reports + count_safe_reports_with_bad_levels))
}

