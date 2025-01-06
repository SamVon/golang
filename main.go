// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	fmt.Println("Helllo, World!")
// 	id := uuid.New()

// 	fmt.Printf("Generated UUID: %s\n", id)
// }

// func main(){
// 	var score int = 50

// 	if score >= 70 {
// 		fmt.Printf("You passed!")
// 	} else {
// 		fmt.Printf("You failed!")
// 	}
// }

// func main(){
// 	var mySlice []int // Declared but not initialized

// 	// Appending data to the slice
// 	mySlice = append(mySlice, 10)
// 	mySlice = append(mySlice, 20, 30)
// 	mySlice = append(mySlice, 20, 30)

// 	fmt.Println(mySlice) //
// }

// convert array to slice
// func main() {
// 	var myArray [3]int // An array of 3 integers
// 	myArray[0] = 10    // Assign values
// 	myArray[1] = 20
// 	myArray[2] = 30

// 	// Converting array to slice
// 	mySlice := myArray[:]

// 	// Resizing slice by appending new elements
// 	mySlice = append(mySlice, 40, 50, 60)

// 	fmt.Println(mySlice) // Output will show a slice with 5 elements: [10 20 30 40 50]
// }

// func main() {
// 	// x := 10

// 	// // Declare a pointer to an integer and assign it the address of x
// 	// var p *int = &x

// 	// // Print the value of x and the value at the pointer p
// 	// fmt.Println("Value of x:", x)  // Output: Value of x: 10
// 	// fmt.Println("Value at p:", *p) // Output: Value at p: 10

// 	// // Modify the value at the pointer p
// 	// *p = 20

// 	// // x is modified since p points to x
// 	// fmt.Println("New value of x:", x) // Output: New value of x: 20

// 	// map
// 	myMap := make(map[string]int)

// 	myMap["apple"] = 5
// 	myMap["banana"] = 10
// 	myMap["orange"] = 8

// 	// show all index
// 	fmt.Println("Apples:", myMap)

// 	// dalete wite key in map
// 	delete(myMap, "banana")

// 	// iterate over the map
// 	for key, value := range myMap {
// 		fmt.Printf("%s -> %d\n", key, value)
// 	}

// 	// checking if a key exists
// 	val, ok := myMap["apple"]
// 	if ok {
// 		fmt.Println("Pear's value:", val)
// 	} else {
// 		fmt.Println("Pear not found inmap")
// 	}
// }

// struct
// type Student struct {
// 	Name    string
// 	Height  int
// 	weigth  int
// 	Grade   string
// 	Address Address
// }

// // Another struct type used in Student
// type Address struct {
// 	Street  string
// 	City    string
// 	Zipcade int
// }

// func main() {

// 	//create an instance of the student struct
// 	var student Student
// 	student.Name = "same"
// 	student.Height = 18
// 	student.weigth = 30
// 	student.Address = Address{
// 		Street:  "123 dongsard",
// 		City:    "Viantein",
// 		Zipcade: 12345,
// 	}

// 	fmt.Println(student)
// }

//===================== // ========= //
// create func add paramater a,b and return (a+b) when use func add
// func add(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	number1 := 3
// 	number2 := 5

// 	// input palamater number1,number2 send to func
// 	sumNumber := add(number1, number2)
// 	fmt.Println(sumNumber)
// }

//===================== // ========= //

// divide divides two integers and returns an error if the divisor is 0
// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Result:", result)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello World!")
// }

// func main() {
// 	http.HandleFunc("/hello", helloHandler)

// 	fmt.Printf("Starting server at port 8080\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
