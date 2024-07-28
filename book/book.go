package book

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const FILE_NAME = "./book.json"

type Book struct {
	ID     int    `json:"id:`
	Title  string `json:"title"`
	Auther string `json:"auther"`
}

func getBooks() ([]Book, error) {
	var books []Book
	file, err := os.OpenFile(FILE_NAME, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return books, err
	}
	defer file.Close()
	booksJson, err := ioutil.ReadAll(file)
	if err != nil {
		return books, err
	}
	err = json.Unmarshal(booksJson, &books)
	if err != nil {
		return books, err
	}
	return books, nil
}

func setBooks(books []Book) error {
	jsonDatas, err := json.Marshal(books)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(FILE_NAME, jsonDatas, 0644)
	fmt.Println(err, books)

	return err
}

// 获取book all 或者 id
func GetBook(cmd *flag.FlagSet, _id *string, _all *bool) ([]Book, error) {
	cmd.Parse(os.Args[2:])
	id := *_id
	all := *_all
	var books []Book
	fmt.Println(id, all)

	if all == false && id == "" {
		return books, nil
	}
	intId, err := strconv.Atoi(id)
	fmt.Println("开始查询 all = %b, id = %s", all, id)
	// 解析
	books, err = getBooks()
	if all {
		return getBooks()
	}
	for i := 0; i < len(books); i++ {
		if books[i].ID == intId {
			return []Book{books[i]}, nil
		}
	}
	return books, err
}

func AddBook(cmd *flag.FlagSet, _id, _title, _auther *string, addBooks bool) {
	cmd.Parse(os.Args[2:])
	id := *_id
	title := *_title
	auther := *_auther
	if id == "" || title == "" || auther == "" {
		os.Exit(1)
	}
	books, _ := getBooks()
	var newBook Book
	var foundBook bool
	intId, _ := strconv.Atoi(id)
	if addBooks {
		newBook = Book{ID: intId, Title: title, Auther: auther}
		books = append(books, newBook)
	} else {
		for i, book := range books {
			if book.ID == intId {
				books[i] = Book{ID: intId, Title: title, Auther: auther}
				foundBook = true
			}
		}
		if !foundBook {
			os.Exit(1)
		}
	}
	setBooks(books)

}

func DeleteBook(cmd *flag.FlagSet, _id *string) {
	cmd.Parse(os.Args[2:])
	id, _ := strconv.Atoi(*_id)
	books, _ := getBooks()
	var deleteIndex int = -1
	for i, book := range books {
		if book.ID == id {
			deleteIndex = i
		}
	}
	if deleteIndex != -1 {
		books = append(books[:deleteIndex], books[deleteIndex:]...)
	}
	setBooks(books)
}
