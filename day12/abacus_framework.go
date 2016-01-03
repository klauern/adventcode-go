package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var f interface{}
	json.Unmarshal(file, &f)

	PrintOutput(f)
	fmt.Println("Output for input is ", Sum(f))
}


func PrintOutput(input interface{}) {
	m := input.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, "is a map:")
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func SumList(input []interface{}) int {
	total := 0
	for _, v := range input {
		switch v.(type) {
		case int:
			total += v.(int)
		}
	}
	return total
}

func SumMap(input map[string]interface{}) int {
	total := 0
	for k, v := range input {
		switch vv := v.(type) {
		case int:
			total += vv
		case []interface{}:
			for _, vvv := range vv {
				total += SumList(vvv.([]interface{}))
			}
		case map[string]interface{}:
			total += SumMap(v.(map[string]interface{}))
		default:
			panic(fmt.Sprintf("Can't evaluate %v for %v as %v", k, v, vv))
		}
	}
	return total
}

func Sum(input interface{}) int {
	total := 0
	switch input.(type) {
	case map[string]interface{}:
		total += SumMap(input.(map[string]interface{}))
	case []interface{}:
		total += SumList(input.([]interface{}))
	}
	return total
}
