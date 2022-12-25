package models

type Book struct {
	ID         int    `json:'id'`
	BookName   string `json:'bookname'`
	BookType   string `json:'booktype'`
	Author     string `json:'author'`
	Popularity int    `json:'popularity'`
	TotalBook  int    `json:'totalbook'`
	AgeLimit   int    `json:'agelimit'`
}

type User struct {
	ID        int    `json:'id'`
	Name      string `json:'name'`
	Email     string `json:'email'`
	Password  string `json:'password'`
	Birthyear string `json:'birthyear'`
}

type LoginUser struct {
	Email    string `json:'email'`
	Password string `json:'password'`
}

type LoginUserInformation struct {
	Email    string `json:'email'`
	Password string `json:'password'`
	Age      int    `json:'age'`
}
