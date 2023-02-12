package DbOperations

import (
	"Task_1/entity"
	"fmt"
	"github.com/nedpals/supabase-go"
)

func GetAllColumns(client *supabase.Client) {
	fmt.Println("##GETTING ALL COLUMNS##")
	var results []entity.Student
	err := client.DB.From("Students").Select("*").Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Selected rows
}

func GetNames(client *supabase.Client) {
	fmt.Println("##GETTING NAME COLUMN##")
	var results []entity.Student
	err := client.DB.From("Students").Select("Name").Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Selected rows
}

func GetNumbers(client *supabase.Client) {
	fmt.Println("##GETTING NUMBER COLUMN##")
	var results []entity.Student
	err := client.DB.From("Students").Select("Number").Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Selected rows
}
