package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	Name       string
	Address    string
	Job        string
	Motivation string
}

func GetFriendByNumber(number int) (Friend, error) {

	data := map[int]Friend{
		1: {"Anna", "Address 01", "Developer", "Professional Development"},
		2: {"Elsa", "Address 02", "Data Scientist", "Professional Development"},
		3: {"Ron", "Address 03", "Engineer", "Professional Development"},
		4: {"Belle", "Address 04", "Marketing", "Professional Development"},
	}

	friend, exists := data[number]
	if !exists {
		return Friend{}, fmt.Errorf("friend with number %d not found", number)
	}

	return friend, nil
}

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("go run main.go [number]")
		return
	}

	number := args[1]

	numberInt, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("number must be a integer (1, 2, 3, 4)")
		return
	}

	friend, err := GetFriendByNumber(numberInt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("number [%d] \n", numberInt)
	fmt.Printf("name       : %s \n", friend.Name)
	fmt.Printf("address    : %s \n", friend.Address)
	fmt.Printf("job        : %s \n", friend.Job)
	fmt.Printf("motivation : %s \n", friend.Motivation)

}
