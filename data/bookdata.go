package data

import (
	"math/rand"
	model "workspace/models"
)

var bookID int = 1

func InitiliazeBookSlice() []model.Book {
	// BookID, ISBN, Stock Code, Page Number, Stock Quantity, Price, "Book Name", "Author Name"
	book1 := model.NewBook(getNextBookId(), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(100, 500), 3, randomNumberGeneratorForPrice(), "Sherlock Holmes Red Leech", "Sir Arthur Conan Doyle")
	book2 := model.NewBook(getNextBookId(), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(100, 500), 7, randomNumberGeneratorForPrice(), "Crime And Punishment", "Dostoyevski")
	book3 := model.NewBook(getNextBookId(), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(100, 500), 18, randomNumberGeneratorForPrice(), "Sherlock Holmes The Sign of the Four", "Sir Arthur Conan Doyle")
	book4 := model.NewBook(getNextBookId(), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(100, 500), 6, randomNumberGeneratorForPrice(), "Animal Farm", "Geroge Orwell")
	book5 := model.NewBook(getNextBookId(), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(100, 500), 9, randomNumberGeneratorForPrice(), "Little Prince", "Antoine de Saint-Exupery")
	book6 := model.NewBook(getNextBookId(), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(100, 500), 35, randomNumberGeneratorForPrice(), "Sapiens", "Yuval Noah Harari")
	book7 := model.NewBook(getNextBookId(), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(100, 500), 2, randomNumberGeneratorForPrice(), "Harry Potter and the Philosopher's Stone", "J. K. Rowling")
	book8 := model.NewBook(getNextBookId(), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(1000000, 9999999), randomNumberGenerator(100, 500), 11, randomNumberGeneratorForPrice(), "Harry Potter and the Prisoner of Azkaban", "J. K. Rowling")

	return []model.Book{*book1, *book2, *book3, *book4, *book5, *book6, *book7, *book8}
}

//This function used for randomly generate isbn number ,stock code and page number
func randomNumberGenerator(min, max int) int {
	return rand.Intn(max-min) + min
}

//This function used for randomly generate book price
func randomNumberGeneratorForPrice() float64 {
	min := 10.0
	max := 99.99
	return min + rand.Float64()*(max-min)
}

//This function return the approprite bookID
func getNextBookId() int {
	bookID += 1
	return bookID - 1
}
