package session

// check version
go version

// initializing project
mkdir <project-name>
cd <project-name>
go mod init <project-name>

// execution program files
cd <project-name>
go run main.go

-- mandatory condition
package main

// compile program files
go build -o <executable-name>
go build -o <executable-name>.exe

// download package
go get <package-name>

// update package
go get -u <package-name>


// variable with data types
var name string = "anna"
var age int
age = 20

// variable without data type
var address = "jakarta"

// variable with short declaration
motivation := "professional development"

// variable with multiple declaration
var one, two, three int = 1, 2, 3
var first, second, third int
first, second, third = 1, 2, 3
var name, address, age = "anna", "jakarta", 20

// underscore variable
var one, two, three int = 1, 2, 3
_, _, _ = one, two, three

-- note
if the variable is not used then use underscore variables

// go data type categories
basic type: number, string, boolean
aggregate type: array, struct
reference type: slice, pointer, map, function, channel
interface type: interface

// nil 
- nil is not a data type, but a value
- nil is an empty value
- data types that can be set to nil include:
	- pointer
	- function
	- slice
	- map
	- channel
	- empty interface or interface{}

// zero value
- zero value (default value of data type)
- zero value string: ""
- zero value bool: false
- zero value numeric non decimal: 0
- zero value numeric decimal: 0.0

// constant
const fullname string = "anna frozen"

