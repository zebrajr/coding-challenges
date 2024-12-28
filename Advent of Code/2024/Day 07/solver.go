package main

import (
    "time"
	"bufio"
	"log"
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

func generate_combinations(numbers int, operators []string) []string{
    if numbers == 0 {
        return []string {""}
    }

    combinations := []string{}
    for _, c := range operators{
        for _, prev := range generate_combinations(numbers -1, operators){
            combinations = append(combinations, prev + string(c))
        }
    }

    return combinations
}

func calculate_all_operations_for_equation(eq equation, operations []string) []string {
    var final_list []string
    list_combinations := generate_combinations(len(eq.numbers) - 1, operations)

    for _, c := range list_combinations {
        //log.Println(list_combinations, idx, c)
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
            if string(op) == "|" {
                temp_a := strconv.Itoa(current_result)
                temp_b := strconv.Itoa(eq.numbers[ido + 1])
                result_string := temp_a + temp_b
                current_result, _ = strconv.Atoi(result_string)
            }
        }
        //log.Println("Wanted vs Got: ", eq.target_number, current_result)
        if eq.target_number == current_result {
            return true
        }
    }
    return false
}


func time_track(start time.Time, name string){
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main() {
    defer time_track(time.Now(), "main")
    equations := load_puzzle_from_file(puzzle_file_name) 
    final_result := 0
    possible_operations := []string {"+", "*"}
    for _, eq := range equations {
        operations_list := calculate_all_operations_for_equation(eq, possible_operations)
        if is_equation_possible(eq, operations_list){
            final_result += eq.target_number
        }
    }
    log.Println("Final Result - Part A: ", final_result)

    // this is WRONG
    // the question clearly asks to use "||" as the operators
    // I don't want to refactor the code base to handle 2 characters for the operations
    //
    // the fix would be to instead of returning an array of strings eg: [ +*|| *+|| ]
    // to instead return a string of strings [ [+ * || ] [ * + || ] ]
    // which would then allow for safe iteration 
    final_result = 0
    possible_operations = []string {"+", "*", "|"}
    for _, eq := range equations {
        operations_list := calculate_all_operations_for_equation(eq, possible_operations)
        if is_equation_possible(eq, operations_list){
            final_result += eq.target_number
        }
    }
    log.Println("Final Result - Part B: ", final_result)
}
