package session

import (
	"fmt"
	"strings"
)

// function
func main() {
	greet("anna", "jakarta")

	message := "hello"
	fullname := []string{"anna", "frozen"}
	result, err := greeting(message, fullname)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(result)

	students := print("a", "b", "c", "d", "e")
	fmt.Printf("%v", students)

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	total := sum(numbers...)
	fmt.Println("total: ", total)

	profile("anna", "pasta", "chicken", "noodles")
}

// without return value
func greet(name, address string) {
	fmt.Println("hello there! my name is ", name)
	fmt.Println("i live in ", address)
}

// with return multiple value
func greeting(message string, fullname []string) (string, error) {
	if len(fullname) == 0 {
		return "", fmt.Errorf("fullname slice is empty")
	}
	joinStr := strings.Join(fullname, " ")
	result := fmt.Sprintf("%s %s", message, joinStr)
	return result, nil
}

// variadic function
func print(names ...string) []map[string]string {
	result := []map[string]string

	for i, value := range names {
		key := fmt.Sprintf("student %d", i+1)
		temp := map[string]string {
			key: value
		}

		result = append(result, temp)
	}

	return result
}

func sum(number ...int) int {
	total := 0
	for _, v := range number {
		total += v
	}
	return total
}

func profile(name string, foods ...string) {
	mergeFoods := strings.Join(foods, ",")
	fmt.Println("hello there! i'm  ", name)
	fmt.Println("i really love to eat ", mergeFoods)
}

// closure or anonymous function
func main() {
	// declare closure in variable
	evenNumbers := func(numbers ...int) []int {
		result := []int
		for _, v := range numbers {
			if v % 2 == 0{
				result = append(result, v)
			}
		}
		return result
	}
	var numbers = []int{4, 93, 77, 10, 52, 22, 34}
	fmt.Println(evenNumbers(numbers...)) // 4, 10, 52, 22, 34

	// IIFE or immediately-invoked function expression
	isPalindrome := func(str string) bool {
		temp:= ""
		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("katak")
	fmt.Println(isPalindrome)

	// closure as return value
	func findStudent(students []string) func(string) string {
		return func(s string) string {
			var (
				student string
				position int
			)
			for i, v := range students {
				if strings.ToLower(v) == strings.ToLower(s) {
					student = v
					position = i
					break
				}
			}
			if student == "" {
				return fmt.Sprintf("%s doesnt exist", s)
			}
			return fmt.Sprintf("we found %s at position %d", s, position+1)
		}
	}
	studentLists := []{"a", "b", "c", "d"}
	find := findStudent(studentLists)
	fmt.Println(find("a"))

	// closure as callback
	fuct findOddNumbers(numbers []int, callback func(int) bool) int {
		var totalOddNumbers int
		for _, v := range numbers {
			if callback(v) {
				totalOddNumbers++
			}
		}
		return totalOddNumbers
	}
	var numbers = []int{2, 5, 8, 10, 3, 99, 23}
	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})
	fmt.Println("total odd numbers: ", find)
}

// pointer (memory address)
var basic_var int = 4
var pointer_var *int = &basic_var // assigns a value to a pointer variable with ampersand (&)

fmt.Println("basic_var (value): ", basic_var) // displays the original value of an basic variable
fmt.Println("basic_var (memory address): ", &basic_var) // displays the memory address of an basic variable with ampersand (&)

fmt.Println("pointer_var (value): ", *pointer_var) // displays the original value of an pointer variable with an asterisk (*)
fmt.Println("pointer_var (memory address): ", pointer_var) // displays the memory address of an pointer variable 

// changing value through pointer
*pointer_var = 7

// pointer as a parameter
func main() {
	var a int = 10
	fmt.Println("before: ", a) // 10

	changeValue(&a)
	fmt.Println("after: ", a) // 20
}

func changeValue(number *int){
	*number = 20
}

// struct
type Employee struct {
	name 	 string
	age 	 int
	division string
}

// giving value to struct
employee1 := Employee{}
employee1.name = "key"
employee1.age = 20
employee1.division = "design"
fmt.Printf("employee1: %+v\n", employee1)

var employee2 Employee
employee2.name = "anna"
employee2.age = 17
employee2.division = "tech"
fmt.Println(employee2.name)
fmt.Println(employee2.age)
fmt.Println(employee2.division)

// initializing struct
var employee3 = Employee{name: "issa", age: 18, division: "finance"}
fmt.Printf("employee3: %+v\n", employee3)

// pointer to a struct
var employee4 *Employee = &employee1 //  variable has a value such as employee1

// embedded struct
type Person struct {
	name string
	age  int
}

type Employee struct {
	division string
	person   Person
}

func main() {
	employee := Employee{}
	employee.person.name = "anna"
	employee.person.age = 20
	employee.division = "curriculum developer"
	fmt.Printf("%+v", employee) // {division:curriculum developer person:{name:anna age:20}}
}

// anonymous struct without value initialization
var employee = struct {
	person   Person
	division string
}{}
employee.person = Person{name: "anna", age: 20}
employee.division = "developer"
fmt.Printf("%+v", employee)

// anonymous struct with value initialization
var employee = struct {
	person   Person
	division string
}{
	person: Person{name: "anna", age: 20},
	division: "finance",
}
fmt.Printf("%+v", employee)

// slice of struct
var people = []Person {
	{name: "anna", age:21},
	{name: "anna", age:21},
	{name: "anna", age:21},
}
for _, v := range people {
	fmt.Printf("%+v", v)
}

// slice of anonymous struct
var employee = []struct {
	person   Person
	division string
}{
	{person: Person{name: "anna", age:21}, division: "developer"},
	{person: Person{name: "anna", age:21}, division: "developer"},
	{person: Person{name: "anna", age:21}, division: "developer"},
	{person: Person{name: "anna", age:21}, division: "developer"},
}

// the way to export variables or data types is to start writing them with uppercase