package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

//var puzzle_input_file = "simple_puzzle"
var puzzle_input_file = "puzzle_input"
var multiplier_expression = `mul\((\d+),(\d+)\)`
var digits_expression = `(\d+)`
var do_dont_expression = `do\(\)|don't\(\)`
var string_do = `do()`
var string_dont = `don't()`


func find_dos_donts(input_string string) int {
    do_dont_regex := regexp.MustCompile(do_dont_expression)
    parts := do_dont_regex.FindAllStringIndex(input_string, -1)
    current_result := 0
    next_start_index := 0
    is_multiplier_enabled := true
    for _, part := range parts{
        match_start := part[0]
        match_end := part[1]
        // first match, so check start or string -> now
        if is_multiplier_enabled {
            string_to_check := input_string[next_start_index:match_start]
            mul_result := find_mul_in_string(string_to_check)
            current_result += mul_result
        }

        next_start_index = match_end

        // check which operation we got coming next
        new_operation := input_string[match_start:match_end]
        if new_operation == string_do{
            if is_multiplier_enabled == false {
                is_multiplier_enabled = true
            }
        }
        if new_operation == string_dont{
            if is_multiplier_enabled == true {
                is_multiplier_enabled = false 
            }
        }
    }

    // before exiting, handle the rest of the text
    if is_multiplier_enabled{
        string_to_check := input_string[next_start_index:]
        mul_result := find_mul_in_string(string_to_check)
        current_result += mul_result
    }

    return current_result
}


func find_mul_in_string(input_string string) int {
    final_answer := 0
    mul_regex := regexp.MustCompile(multiplier_expression)
    //all_matches := mul_regex.FindAllString(input_string, -1)
    all_matches := mul_regex.FindAllString(input_string, -1)
    for _, value := range all_matches{
            numbers_regex := regexp.MustCompile(digits_expression)
            numbers_matches := numbers_regex.FindAllString(value, -1)
            first_number, _ := strconv.Atoi(numbers_matches[0])
            second_number, _ := strconv.Atoi(numbers_matches[1])
            multiplication_result := first_number * second_number 
            final_answer += multiplication_result
    }
    return final_answer
}

func main(){
    // if we read the file as a buffer like we did in the day 01 and 02
    // we will get an error because we have multiple lines (6)
    // therefore making the `final_answer_advanced` higher then it should
    // since a new line will mean the mul(x,y) are valid
    // imho that should be the right answer
    input_file, err := os.ReadFile(puzzle_input_file)
    if err != nil {
        log.Fatal(err)
    }
    input_string := string(input_file)

    final_answer := find_mul_in_string(input_string)
    fmt.Println("Final Answer: ", final_answer)

    final_answer_advanced := find_dos_donts(input_string) 
    fmt.Println("Final Answer with do() / don't(): ", final_answer_advanced)
}
