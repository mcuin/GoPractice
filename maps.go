package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])

	delete(y, 1)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	
	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"])
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
	
	/*elementsShort := map[string]string {
		"H" : "Hydrogen",
		"He" : "Helium",
		"Li" : "Lithium",
		"Be" : "Beryllium",
		"B" : "Boron",
		"C" : "Carbon",
		"N" : "Nitrogen",
		"O" : "Oxygen",
		"F" : "Flourine",
		"Ne" : "Neon",
	}*/

	elementsStates := map[string]map[string]string {
		"H" : map[string]string {
			"name" : "Hydrogen",
			"state" : "Gas",
		},
		"He" : map[string]string {
			"name" : "Helium",
			"state" : "Gas",
		},
		"Li" : map[string]string {
			"name" : "Lithium",
			"state" : "Solid",
		},
		"Be" : map[string]string {
			"name" : "Beryllium", 
			"state" : "Solid",
		},
		"B" : map[string]string {
			"name" : "Boron",
			"state" : "Solid",
		},
		"C" : map[string]string {
			"name" : "Cabon",
			"state" : "Solid",
		},
		"N" : map[string]string {
			"name" : "Nitrogen",
			"state" : "Gas",
		},
		"O" : map[string]string {
			"name" : "Oxygen",
			"state" : "Gas",
		},
		"F" : map[string]string {
			"name" : "Flourine",
			"state" : "Gas",
		},
		"Ne" : map[string]string {
			"name" : "Neon",
			"state" : "Gas",
		},
	}

	if el, ok := elementsStates["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
		
}
