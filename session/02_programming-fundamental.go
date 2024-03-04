package session

import "fmt"

// if else
func conditions() {
	// conditions (temporary variable)
	var currentYear = 2024
	if age := currentYear - 1998; age < 17 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

// switch case
func switchs() {
	// switch
	var score = 7
	switch score {
	case 10:
		fmt.Println("perfect")
	case 8:
		fmt.Println("good")
	default:
		fmt.Println("not bad")
	}

	// switch with relational operators
	switch {
	case score == 10:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("u dont hv a good score yet")
		}
	}

	// switch with fallthrough keyword
	switch {
	case score == 10:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
		fallthrough
	case score > 5:
		fmt.Println("its ok but pls study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("u dont hv a good score yet")
		}
	}
}

// looping
func loopings() {
	// first way of looping
	for i := 0; i < 3; i++ {
		fmt.Println("number", i)
	}

	// second way of looping
	var i = 0
	for i < 3 {
		fmt.Println("number", i)
		i++
	}

	// third way of looping
	var i = 0
	for {
		fmt.Println("number", i)
		i++
		if i == 3 {
			break
		}
	}

	// break and continue keywords
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("number", i)
	}

	// nested looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Println(j, " ")
		}
		fmt.Println()
	}

	// looping label
outerloop:
	for i := 0; i < 3; i++ {
		fmt.Println("looping ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerloop
			}
			fmt.Println(j, " ")
		}
		fmt.Println()
	}

}

// array
func arrays() {
	// modify array element through index
	var numbers [7]int
	numbers = [7]int{1, 2, 3, 4, 5, 6, 7}
	var strings = [3]string{"a", "b", "c"}
	numbers[0] = 9
	strings[1] = "z"

	// loop array through elements
	for i, v := range strings {
		fmt.Printf("index: %d, value: %s \n", i, v)
	}

	fmt.Println(strings.Repeat("-", 25))

	for i := 0; i < len(strings); i++ {
		fmt.Printf("index: %d, value: %s \n", i, strings[i])
	}

	// multidimensional array
	balances := [2][3]int{{5,6,7},{8,9,10}}
	for _, arr := range balances {
		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

// slice
func slices() {
	// common way declaring a slice
	var fruits = []string{"a", "b", "c", "d", "e", "f", "g"}
	books := []string{"x", "y", "z"}

	// make function
	var musics = make([]string, 7)

	// append function
	musics = append(musics, "a", "b", "c")

	// append function with ellipsis
	musics = append(musics, books...)

	// copy function
	copies := copy(musics, books)

	// slicing
	fruits1 := fruits[1:4] 			// output: b, c, d
	fruits2 := fruits[0:]  			// output: a, b, c, d, e, f, g
	fruits3 := fruits[:3]  			// output: a, b, c
	fruits4 := fruits[:]   			// output: a, b, c, d, e, f, g
	fruits5 := fruits[:len(fruits)] // output: a, b, c, d, e, f, g

	// combining slicing and append
	fruits = append(fruits[:3], "z") // output: a, b, c, z 

	// backing array
	fruits6 = fruits[2:4] // output: c, d
	fruits6[0] = "x"      // output: x, d  and then output fruits: a, b, x, d, e, f, g bcs still in the same backing array

	// cap function
	fmt.Println("fruits cap: ", cap(fruits)) // 7
	fmt.Println("fruits len: ", len(fruits)) // 7

	fruits7 := fruits[0:3]
	fmt.Println("fruits cap: ", cap(fruits)) // 7
	fmt.Println("fruits len: ", len(fruits)) // 3

	fruits7 := fruits[1:0]
	fmt.Println("fruits cap: ", cap(fruits)) // 6
	fmt.Println("fruits len: ", len(fruits)) // 6

	// creating a new baking array
	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...) // ford, honda
	
	cars[0] = "nissan" // cars output: nissan, honda, audi, range rover
}

// map 
func maps() {

	// declaration
	var person map[string]string

	// initialization
	person = map[string]string{}

	person["name"] = "anna" 
	
	fmt.Println("name: ", person["name"]) // name: anna

	var person = map[string]string {
		"name": "anna",
		"age": "17",
	}

	// looking with map
	for key, value := range person {
		fmt.Println(key, ":", value)

	}

	// deleting value 
	delete(person, "age")

	// detecting value
	value, exist := person["age"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value doesnt exist")
	}

	// combining slice with map
	var people = []map[string]string{
		{"name": "anna", "age": "20"},
		{"name": "anna", "age": "20"},
		{"name": "anna", "age": "20"},
	}

	for i, val := range people {
		fmt.Printf("index: %d, name: %s, age: %s \n", i, val["name"], val["age"])
	}

}