package lists

import "fmt"

type product struct {
	id    string
	title string
	price int
}

// func main() {
// 	// 1
// 	hobbies := [3]string{"eating", "cooking", "dancing"}
// 	fmt.Println(hobbies)

// 	// 2
// 	fmt.Println(hobbies[0])
// 	fmt.Println(hobbies[1:])

// 	// 3
// 	mainHobbies := hobbies[:2]
// 	fmt.Println(mainHobbies)

// 	// 4
// 	fmt.Println(cap(mainHobbies))
// 	mainHobbies = mainHobbies[1:3]
// 	fmt.Println(mainHobbies)

// 	// 5
// 	courseGoals := []string{"be good at golang", "get a job"}
// 	fmt.Println(courseGoals)

// 	// 6
// 	courseGoals[1] = "Learn all the details"
// 	courseGoals = append(courseGoals, "Know this very well")
// 	fmt.Println(courseGoals)

// 	// 7
// 	products := []product{
// 		{
// 			"id",
// 			"A first ",
// 			300,
// 		},
// 		{
// 			"ids",
// 			"A second product",
// 			4000,
// 		},
// 	}
// 	fmt.Println(products)
// 	newProd := product{
// 		"123",
// 		"A third prod",
// 		1000,
// 	}
// 	products = append(products, newProd)
// 	fmt.Println(products)

// 	book()
// }

func main() {
	prices := []float64{10.00, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 7.77, 66.55)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.56}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

// Create a program that manages a book collection.
// Initialize an array with at least three books, each represented by a struct with fields like title, author, and genre.
// Print the array of books to the command line.
// Output more data about the array:
// Print the details of the first book.
// Combine the details of the second and third books into a new list and print it.
// Create a slice based on the first book that contains the first and second books.
// Implement two different ways to create this slice.
// Re-slice the slice from the previous step and change it to contain the second and last books from the original array.
// Create a "dynamic array" that contains your favorite movies (at least 2 movies).
// Update the second movie to a different one and then add a third movie to the existing dynamic array.
// Bonus: Create a "MusicAlbum" struct with title, artist, and release year, and create a dynamic list of albums (at least 2 albums). Then, add a third album to the existing list of albums.
