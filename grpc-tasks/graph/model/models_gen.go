// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ExistingUser struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastnamwe string `json:"lastnamwe"`
}

type ID struct {
	ID string `json:"id"`
}

type IDInput struct {
	ID string `json:"id"`
}

type NewUser struct {
	Firstname string `json:"firstname"`
	Lastnamwe string `json:"lastnamwe"`
}

type Status struct {
	Status int `json:"Status"`
}

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastnamwe string `json:"lastnamwe"`
}
