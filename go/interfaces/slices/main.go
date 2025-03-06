package main

import "fmt"

func main() {

	names := []string{"Jacob", "Jenn"}

	m := map[string]interface{}{
		"n": names,
	}

	_, isStrings := m["n"].([]string)
	_, isInterface := m["n"].(interface{})
	_, isInterfaceSlice := m["n"].([]interface{})
	fmt.Println("isStrings", isStrings, "isInterface", isInterface, "isInterfaceSlice", isInterfaceSlice)
}
