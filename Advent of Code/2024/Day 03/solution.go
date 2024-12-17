package main

import (
	"bufio"
	"fmt"
    "strconv"
	"log"
	"os"
	"regexp"
)

var puzzle_input_file = "puzzle_input"
var regex_pattern_mul = `mul\((\d+),(\d+)\)`
var regex_pattern_numbers = `(\d+)`

func main(){
    file, err := os.Open(puzzle_input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    final_answer := 0
    file_scanner := bufio.NewScanner(file)
    mul_regex := regexp.MustCompile(regex_pattern_mul) 

    for file_scanner.Scan(){
        current_line := file_scanner.Text()
        all_matches := mul_regex.FindAllString(current_line, -1) 
        for _, value := range all_matches {
            numbers_regex := regexp.MustCompile(regex_pattern_numbers)
            numbers_matches := numbers_regex.FindAllString(value, -1)
            first_number, _ := strconv.Atoi(numbers_matches[0])
            second_number, _ := strconv.Atoi(numbers_matches[1])
            multiplication_result := first_number * second_number 
            final_answer += multiplication_result
        }
    }
    fmt.Println("Final Answer: ", final_answer)
}
