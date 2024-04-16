package lists

import "fmt"

type BookType struct {
	title  string
	author string
	genre  string
}

type MusicAlbum struct {
	title  string
	artist string
	year   int
}

func book() {
	// Initialize an array with at least three books
	bookList := []BookType{
		{
			title:  "Harry Potter",
			author: "John Legend",
			genre:  "Scary Movies",
		},
		{
			title:  "Witches of Oz",
			author: "Angelina Jolie",
			genre:  "Horror",
		},
		{
			title:  "Woodpecker",
			author: "John Wick",
			genre:  "Fun",
		},
	}

	// Print the array of books
	fmt.Println("List of Books:")
	fmt.Println(bookList)

	// Print the details of the first book
	firstBook := bookList[0]
	fmt.Println("\nDetails of the First Book:")
	fmt.Println(firstBook)

	// Combine the details of the second and third books into a new list and print it
	secondAndThirdBooks := []BookType{bookList[1], bookList[2]}
	fmt.Println("\nCombined Details of Second and Third Books:")
	fmt.Println(secondAndThirdBooks)

	// Create a slice based on the first book that contains the first and second books
	// Implement two different ways to create this slice
	firstAndSecondBooks := []BookType{bookList[0], bookList[1]}
	firstAndSecondBooksAlternative := append([]BookType{}, bookList[:2]...)
	fmt.Println("\nSlice with First and Second Books:")
	fmt.Println(firstAndSecondBooks)
	fmt.Println(firstAndSecondBooksAlternative)

	// Re-slice the slice from the previous step and change it to contain the second and last books from the original array
	reSlicedBooks := firstAndSecondBooks[1:]
	reSlicedBooks = append(reSlicedBooks, bookList[len(bookList)-1])
	fmt.Println("\nRe-sliced Slice with Second and Last Books:")
	fmt.Println(reSlicedBooks)

	// Create a "dynamic array" that contains favorite movies
	movieList := []string{"Inception", "The Shawshank Redemption"}

	// Update the second movie to a different one
	movieList[1] = "Interstellar"

	// Add a third movie to the existing dynamic array
	movieList = append(movieList, "The Dark Knight")
	fmt.Println("\nList of Favorite Movies:")
	fmt.Println(movieList)

	// Bonus: Create a dynamic list of albums (at least 2 albums) and add a third album
	albumList := []MusicAlbum{
		{
			title:  "Dance",
			artist: "Wizkid",
			year:   2010,
		},
		{
			title:  "Jo",
			artist: "David",
			year:   2011,
		},
	}

	// Add a third album to the existing list of albums
	albumList = append(albumList, MusicAlbum{
		title:  "Top",
		artist: "Burna Boy",
		year:   2022,
	})
	fmt.Println("\nList of Music Albums:")
	fmt.Println(albumList)
}
