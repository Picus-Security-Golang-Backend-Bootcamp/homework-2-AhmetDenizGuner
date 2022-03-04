package searchhelper

import (
	"errors"
	"strconv"
	"strings"
	model "workspace/models"
)

//this method search the matches between the given string and bookList element, and return the matches list
func Search(bookList []model.Book, searchItems []string) ([]model.Book, error) {

	//creating search string using by program arguments
	//BuildSearchItem method get arguments and return the string these arguments
	searchItem := buildSearchItem(searchItems)

	//check string that will be searched is equal or bigger than  3 char
	if len(searchItem) < 3 {
		err := errors.New("Lütfen daha uzun bir keime giriniz!")
		return nil, err
	}

	searchResult := []model.Book{}

	for _, book := range bookList {
		if strings.Contains(strings.ToLower(book.Name), strings.ToLower(searchItem)) {
			searchResult = append(searchResult, book)
		} else if strings.Contains(strings.ToLower(book.Author.Name), strings.ToLower(searchItem)) {
			searchResult = append(searchResult, book)
		} else if strings.Contains(strings.ToLower(strconv.Itoa(book.ISBNnum)), strings.ToLower(searchItem)) {
			searchResult = append(searchResult, book)
		} else if strings.Contains(strings.ToLower(strconv.Itoa(book.StockCode)), strings.ToLower(searchItem)) {
			searchResult = append(searchResult, book)
		}
	}

	//set error, if there is no result
	if len(searchResult) == 0 {
		err := errors.New("Aradiginiz kitap bulunamadi lutfen baska bir kelime/kelimelerle deneyiniz!")
		return nil, err
	}
	return searchResult, nil
}

//this method return a string using by elements of given string list
func buildSearchItem(argumentSlice []string) string {
	return strings.TrimSpace(strings.Join(argumentSlice, " "))
}
