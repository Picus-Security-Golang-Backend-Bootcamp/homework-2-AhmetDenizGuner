package models

import (
	"errors"
	"strconv"
)

type Book struct {
	Author
	Id          int
	Name        string
	StockCode   int
	ISBNnum     int
	PageNum     int
	StockNumber int //Stock Quantity
	IsDeleted   bool
	Price       float64
}

//struct constructor
func NewBook(id, stockCode, isbn_num, pageNum, stockNumber int, price float64, bookName, authorName string) *Book {

	author := NewAuthor(authorName)

	book := &Book{
		Id:          id,
		Name:        bookName,
		StockCode:   stockCode,
		ISBNnum:     isbn_num,
		PageNum:     pageNum,
		StockNumber: stockNumber,
		IsDeleted:   false,
		Author:      *author,
		Price:       price,
	}

	return book
}

//This function check some rules and decrease the stock quantity
func (b *Book) Buy(count int) (string, error) {

	//check book is deleted before
	if b.IsDeleted {
		err := errors.New("Bu kitap silinmistir lutfen baska kitap almayi deneyiniz!")
		return "", err
	}

	//check ıs there enough book to buy
	if count > b.StockNumber {
		err := errors.New("Yeterli sayida kitap yoktur lutfen daha az miktarda deneyiniz!")
		return "", err
	}

	b.StockNumber -= count

	return b.Name + " " + strconv.Itoa(count) + " adet satin alindi.", nil
}

type Deletable interface {
	Delete()
}

//This function check some rules and set the isDeleted field as true
func (b *Book) Delete() (string, error) {

	//check book is deleted before
	if b.IsDeleted {
		err := errors.New("Bu kitap zaten silinmiş tekrar silemezsiniz!")
		return "", err
	}

	b.IsDeleted = true

	return b.Name + " Basariyla silindi", nil
}
