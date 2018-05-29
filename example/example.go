package main

import (
	"github.com/wilhelmguo/golang-to-typescript/typescriptify"
)

type Address struct {
	// Used in html
	City    string  `json:"city"`
	Number  float64 `json:"number"`
	Country string  `json:"country,omitempty"`
}

type PersonalInfo struct {
	Hobbies []string `json:"hobby"`
	PetName string   `json:"pet_name"`
}

type Person struct {
	Name         string       `json:"name"`
	PersonalInfo PersonalInfo `json:"personal_info"`
	Nicknames    []string     `json:"nicknames"`
	Addresses    []Address    `json:"addresses"`
}

func main() {
	converter := typescriptify.New()
	converter.CreateConstructor = true
	converter.CreateEmptyObject = true
	converter.Indent = "    "
	converter.BackupExtension = ""

	converter.Add(Person{})

	err := converter.ConvertToFile("./example_output.ts")
	if err != nil {
		panic(err.Error())
	}
}
