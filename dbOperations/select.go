package DbOperations

import (
	"Task_1/entity"
	json "encoding/json"
	"fmt"
	"github.com/nedpals/supabase-go"
)

func GetAllData(client *supabase.Client) {
	fmt.Println("##GETTING ALL COLUMNS##")
	var results []entity.Student
	err := client.DB.From("Students").Select("*").Execute(&results)
	if err != nil {
		panic(err)
	}
	if len(results) == 0 {
		fmt.Println("There is no data")
		return
	}
	for _, student := range results {
		stdJson, err := json.Marshal(&student)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(stdJson))
	}
}
func FindById(client *supabase.Client, id string) {
	fmt.Println("##FINDING DATA FOR ID = ", id, " ##")
	var results []entity.Student
	err := client.DB.From("Students").Select("*").Eq("Id", id).Execute(&results)
	if err != nil {
		panic(err)
	}
	if len(results) == 0 {
		fmt.Println("Student not found")
		return
	}
	for _, student := range results {
		stdJson, err := json.Marshal(&student)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(stdJson))
	}

}
func FindByGivenColumn(client *supabase.Client, column string, value string) {
	fmt.Println("##FINDING DATA FOR COLUMN = ", column, " AND VALUE = ", value, " ##")
	var results []entity.Student
	err := client.DB.From("Students").Select("*").Eq(column, value).Execute(&results)
	if err != nil {
		panic(err)
	}
	if len(results) == 0 {
		fmt.Println("Student not found")
		return
	}
	for _, student := range results {
		stdJson, err := json.Marshal(&student)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(stdJson))
	}

}
