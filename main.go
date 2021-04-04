package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
)

type MyBoxItem struct {
	CodeOfBook string
	NamOfBook  string
	IsRented   bool
}

var messages = make(chan string)
var arr []MyBoxItem

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("LIBRARY SYSTEM")
	fmt.Print("cmd : ")
	for scanner.Scan() {

		arrayStr := strings.Fields(scanner.Text())
		switch {
		case strings.ToLower(arrayStr[0]) == "add":
			// add[#space]code_of_book(unique)[#space]name_of_book	(add	new	book	to	inventory
			m := MyBoxItem{CodeOfBook: arrayStr[1], NamOfBook: arrayStr[2], IsRented: false}
			go addBook(m)

			var message1 = <-messages
			fmt.Println(message1)
			fmt.Print("cmd : ")
		case strings.ToLower(arrayStr[0]) == "rent":
			// rent[#space]code_of_book	(for	update	status	of	book	rented)
			go rentBook(arrayStr[1])

			var message1 = <-messages
			fmt.Println(message1)
			fmt.Print("\ncmd : ")
		case strings.ToLower(arrayStr[0]) == "rented":
			// rented	(for	display	all	rented	books)
			listRentedBook()
			fmt.Print("\ncmd : ")
		case strings.ToLower(arrayStr[0]) == "return":
			// return[#space]code_of_book	(for	update	status	of	book	returned)
			go returnBook(arrayStr[1])

			var message1 = <-messages
			fmt.Println(message1)
			fmt.Print("\ncmd : ")
		case strings.ToLower(arrayStr[0]) == "list":
			// list	(for	display	of	list	book	and	status	of	the	book)
			getList()
			fmt.Print("\ncmd : ")
		case strings.ToLower(arrayStr[0]) == "get":
			// get[#space]code_of_book	(show	name	of	book	by	code)
			getBook(arrayStr[1])
		case strings.ToLower(arrayStr[0]) == "clear":
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			fmt.Print("\ncmd : ")
		case strings.ToLower(arrayStr[0]) == "exit":
			os.Exit(1)
		default:
			fmt.Println("command undefined!")
		}
	}
}

func addBook(m MyBoxItem) {
	arr = append(arr, m)
	var data = fmt.Sprintf("new book inserted! Title : %s", m.NamOfBook)
	messages <- data
}

func getList() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s\t", "code_of_book", "name_of_book", "rented")
	fmt.Fprintf(w, "\n%s\t%s\t%s\t", "----", "----", "----")

	for _, p := range arr {
		var isRented string
		if p.IsRented == false {
			isRented = "No"
		} else {
			isRented = "Yes"
		}
		fmt.Fprintf(w, "\n%s\t%s\t%s\t", p.CodeOfBook, p.NamOfBook, isRented)
	}
}

func getBookByCode(m MyBoxItem) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s\t", "code_of_book", "name_of_book", "rented")
	fmt.Fprintf(w, "\n%s\t%s\t%s\t", "----", "----", "----")

	var isRented string
	if m.IsRented == false {
		isRented = "No"
	} else {
		isRented = "Yes"
	}
	fmt.Fprintf(w, "\n%s\t%s\t%s\t", m.CodeOfBook, m.NamOfBook, isRented)
}

func getBook(code string) {
	for _, k := range arr {
		if k.CodeOfBook == code {
			getBookByCode(k)
			fmt.Print("\ncmd : ")
		}
	}
}

func rentBook(code string) {
	for i, k := range arr {
		if k.CodeOfBook == code {
			arr[i].IsRented = true
		}
	}
	var data = fmt.Sprintf("successfuly update code book : %s", code)
	messages <- data
}

func listRentedBook() {
	var tmpArr []MyBoxItem

	for _, k := range arr {
		if k.IsRented == true {
			tmpArr = append(tmpArr, k)
		}
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s\t", "code_of_book", "name_of_book", "rented")
	fmt.Fprintf(w, "\n%s\t%s\t%s\t", "----", "----", "----")

	for _, p := range tmpArr {
		var isRented string
		if p.IsRented == false {
			isRented = "No"
		} else {
			isRented = "Yes"
		}
		fmt.Fprintf(w, "\n%s\t%s\t%s\t", p.CodeOfBook, p.NamOfBook, isRented)
	}
}

func returnBook(code string) {
	for i, k := range arr {
		if k.CodeOfBook == code {
			arr[i].IsRented = false
		}
	}
	var data = fmt.Sprintf("successfuly returned book : %s", code)
	messages <- data
}
