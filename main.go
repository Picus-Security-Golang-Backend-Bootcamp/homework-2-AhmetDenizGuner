package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	data "workspace/data"
	"workspace/internal/searchhelper"
	model "workspace/models"
)

var bookSlice []model.Book = []model.Book{}

func init() {
	bookSlice = data.InitiliazeBookSlice()
}

func main() {

	args := os.Args

	//if user didn't give any argument to the program, return this warning message that includes possible commands and stop program
	if len(args) == 1 {
		printDefaultErrorMessage(args[0])
		return
	}

	if strings.TrimSpace(args[1]) == "list" {
		if len(args) == 2 {
			bookPrinter(bookSlice)
		} else {
			fmt.Println("list komutu baska arguman alamaz!")
		}

	} else if strings.TrimSpace(args[1]) == "search" {

		//check is there any search string
		if len(args) < 3 {
			fmt.Println("Search komutunu kullanabilmek için lütfen arama yapilacak kelime veya kelimleri de giriniz!")
			return
		}

		//Searching string in book list
		//Search method use strings contains method and seek any match in book list and return the all matches
		resultSlices, err := searchhelper.Search(bookSlice, args[2:])

		//if resultSlices is empty Serach method will be return error message and program terminated
		if err != nil {
			fmt.Println(err)
			return
		}

		//if there is search result , bookPrinter function print the books that are in given slice
		bookPrinter(resultSlices)

	} else if strings.TrimSpace(args[1]) == "buy" {

		//check argument count is ok
		if len(args) != 4 {
			fmt.Println("Buy komutu icin hatali sayida arguman girdiniz, lutfen 2 tam sayi giriniz!")
			return
		}

		bookId, err1 := strconv.Atoi(args[2])
		orderCount, err2 := strconv.Atoi(args[3])

		//check arguments are integers
		if err1 != nil || err2 != nil {
			fmt.Println("Lutfen arguman degerlerini tam sayi giriniz!")
			return
		}

		if !checkBookExist(bookId) {
			return
		}

		message, err := bookSlice[bookId-1].Buy(orderCount)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(message)

	} else if strings.TrimSpace(args[1]) == "delete" {

		if len(args) != 3 {
			fmt.Println("Buy komutu icin hatali sayida arguman girdiniz, lutfen 1 tam sayi giriniz!")
			return
		}

		bookId, err1 := strconv.Atoi(args[2])

		//check arguments are integers
		if err1 != nil {
			fmt.Println("Lutfen arguman degerini tam sayi giriniz!")
			return
		}

		if !checkBookExist(bookId) {
			return
		}

		message, err := bookSlice[bookId-1].Delete()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(message)

	} else { // If user gives wrong command , terminate program and print error message.
		printDefaultErrorMessage(args[0])
		return
	}

}

func printDefaultErrorMessage(projectPath string) {
	//taking executable project name from path
	projectName := filepath.Base(projectPath)
	fmt.Printf("%s uygulamasinda kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n buy => satin alma islemi icin\n delete => silme islemi icin \n", projectName)
}

//this function print the list that given
func bookPrinter(books []model.Book) {
	for index, book := range books {
		fmt.Printf("%d) BookID: %d, Name: %s, Author: %s, Price: %.2f, ISBN: %d, Stock Code: %d, Stock Num: %d\n", index+1, book.Id, book.Name, book.Author.Name, book.Price, book.ISBNnum, book.StockCode, book.StockNumber)
	}
}

//This function checks the book added to system inittialy.
func checkBookExist(bookId int) bool {
	if bookId > len(bookSlice) {
		fmt.Println("Bu ID'de hicbir kitap sisteme daha once eklnmedi lutfen daha kucuk ID'lerle deneyiniz!")
		return false
	}
	return true
}
