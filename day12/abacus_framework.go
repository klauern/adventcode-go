package main

import "fmt"


func main() {

	fmt.Println("Output for input is ")
}

func Sum(input map[string]interface{}) int {
	total := 0
	m := input.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case int:
			total += vv
		case []interface{}:
			total += Sum(vv)
		default:
			panic(fmt.Sprintf("Can't evaluate %v for %v", k, vv))
		}
	}
	return total
}
