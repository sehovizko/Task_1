package DbOperations

import (
	"Task_1/entity"
	"fmt"
	"github.com/nedpals/supabase-go"
)

func InsertSomeData(client *supabase.Client, data *[]entity.Student) {

	var results []entity.Student
	err := client.DB.From("Students").Insert(data).Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println("##INSERTED DATA##")
	fmt.Println(results) // Inserted rows

}
