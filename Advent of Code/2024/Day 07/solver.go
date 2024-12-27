package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var puzzle_file_name = "puzzle_input.txt"

type equation struct {
    target_number int
    numbers []int
}

func load_puzzle_from_file(target_file_name string) []equation{
    file, _ := os.Open(target_file_name)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var equations_list []equation
    for scanner.Scan(){
        line := scanner.Text()
        parts := strings.Split(line, ":")

        final_number, _ := strconv.Atoi(parts[0])

        var new_numbers []int
        temp_numbers := strings.Split(parts[1], " ")
        for _, num := range temp_numbers{
            n, _ := strconv.Atoi(num)
            new_numbers = append(new_numbers, n)
        }

        var new_equation = equation {
            target_number: final_number,
            numbers: new_numbers,
        }
        equations_list = append(equations_list, new_equation)
    }
    return equations_list
}

func generate_combinations(numbers int) []string{
    if numbers == 0 {
        return []string {""}
    }

    combinations := []string{}
    for _, c := range []string {"+", "*"}{
        for _, prev := range generate_combinations(numbers -1){
            combinations = append(combinations, prev + string(c))
        }
    }

    return combinations
}

func calculate_all_operations_for_equation(eq equation) []string {
    var final_list []string
    combinations := generate_combinations(len(eq.numbers) - 1)

    for _, c := range combinations {
        final_list = append(final_list, c)
    }

    return final_list
}

func is_equation_possible(eq equation, operations []string) bool {
    for _, ops := range operations {
        current_result := 0
        for ido, op := range ops {
            if ido == 0 {
                current_result = eq.numbers[ido]
            }
            if string(op) == "+" {
                current_result = current_result + eq.numbers[ido + 1]
            }
            if string(op) == "*" {
                current_result = current_result * eq.numbers[ido + 1]
            }
        }
        //fmt.Println("Wanted vs Got: ", eq.target_number, current_result)
        if eq.target_number == current_result {
            return true
        }
    }
    return false
}

func main() {
    equations := load_puzzle_from_file(puzzle_file_name) 
    final_result := 0
    for _, eq := range equations {
        operations_list := calculate_all_operations_for_equation(eq)
        if is_equation_possible(eq, operations_list){
            final_result += eq.target_number
        }
    }
    fmt.Println("Final Result: ", final_result)
}
