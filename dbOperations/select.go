package DbOperations

import (
	"Task_1/entity"
	json "encoding/json"
	"fmt"
	"github.com/nedpals/supabase-go"
)

// I used nedpals/supabase-go library for db operations.
// But I think the library is not efficient for db operations.
// We can create our own repo methods by connecting to supabase's postqres.
// The library does not have a good documentation and I could not find
// features like filtering etc.

func GetAllColumns(client *supabase.Client) {
	fmt.Println("##GETTING ALL COLUMNS##")
	var results []entity.Student
	err := client.DB.From("Students").Select("*").Execute(&results)
	if err != nil {
		panic(err)
	}

	for _, student := range results {
		stdJson, err := json.Marshal(&student)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(stdJson))
	}
}
func GetNames(client *supabase.Client) {
	fmt.Println("##GETTING NAME COLUMN##")
	var results []entity.Student
	err := client.DB.From("Students").Select("Id,Name").Execute(&results)
	if err != nil {
		panic(err)
	}

	for _, student := range results {
		stdJson, err := json.Marshal(&student)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(stdJson))
	}
}
func GetNumbers(client *supabase.Client) {
	fmt.Println("##GETTING NUMBER COLUMN##")
	var results []entity.Student
	err := client.DB.From("Students").Select("Id,Number").Execute(&results)
	if err != nil {
		panic(err)
	}

	for _, student := range results {
		stdJson, err := json.Marshal(&student)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(stdJson))
	}
}
