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
	StockNumber int
	IsDeleted   bool
	Price       float64
}

func NewBook(id, stockCode, isbn_num, pageNum, stockNumber int, price float64, bookName, authorName string) *Book {
	// bu fonksiyon içinde projenizde yaratacağınız struct'ın her zaman geçerli olduğunu garanti edebilirsiniz ve istediğiniz kuralı burada kontrol edebilirsiniz

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

func (b *Book) Buy(count int) (string, error) {

	if b.IsDeleted {
		err := errors.New("Bu kitap silinmistir lutfen baska kitap almayi deneyiniz!")
		return "", err
	}

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

func (b *Book) Delete() (string, error) {

	if b.IsDeleted {
		err := errors.New("Bu kitap zaten silinmiş tekrar silemezsiniz!")
		return "", err
	}

	b.IsDeleted = true

	return b.Name + " Basariyla silindi", nil
}
