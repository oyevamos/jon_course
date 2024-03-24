package main

import (
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
	Meta struct {
		Visists int
	}
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "John Smith",
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
