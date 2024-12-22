package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var puzzle_file_name = "puzzle_sample.txt"

var pages [][]int 
var page []int 

var rule_book = make(map[string] rule) 
type rule struct {
    before int
    after int
}

func load_data_from_file(file_name string){
    file, err := os.Open(file_name)
    if err != nil { return }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan(){
        line := scanner.Text()
        if strings.Contains(line, "|"){
            fields := strings.Split(line, "|")
            first, _ := strconv.Atoi(fields[0])
            second, _ := strconv.Atoi(fields[1])

            new_rule := rule {
                before: first,
                after: second,
            }
            uuid := fields[0] + fields[1]
            rule_book[uuid] = new_rule
        }

        if strings.Contains(line, ","){
            fields := strings.Split(line, ",")
            var new_page []int 
            for _, field := range fields{
                new_number, _ := strconv.Atoi(field)
                new_page = append(new_page, new_number)
            }
            pages = append(pages, new_page)
        }
    }
}

func check_pages_for_problem_part_1(){
    last_order := ""
    sum_of_valid_middle_values := 0
    
    for _, page := range pages{
        //fmt.Println("Checking Page: ", page)
        is_order_valid := true
        for _, order := range page {
            order_as_string := strconv.Itoa(order)
            rule_to_check_for := order_as_string + last_order
            _, ok := rule_book[rule_to_check_for] 
            // since we are comparing the new rule with the older
            // therefore making the rules inverted eg: new|previous
            // means we broke that specific rule
            if ok {
                //fmt.Println("Checked Page: ", page)
                //fmt.Println("Rule Broken: ", rule_to_check_for)
                is_order_valid = false
                break
            }
            last_order = order_as_string 
        }

        if is_order_valid {
            slice_length := len(page)
            middle_index := slice_length / 2
            middle_value := page[middle_index]
            //fmt.Println("Middle Value: ", middle_value)
            sum_of_valid_middle_values += middle_value
        }
        last_order = ""
    }
    fmt.Println("Final Sum of Valid Middle Values: ", sum_of_valid_middle_values)
}

func main(){
    load_data_from_file(puzzle_file_name)
    // fmt.Println(pages)
    // fmt.Println(rule_book)

    check_pages_for_problem_part_1()
}
