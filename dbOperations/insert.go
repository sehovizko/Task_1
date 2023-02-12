package DbOperations

import (
	"Task_1/dto"
	"fmt"
	"github.com/nedpals/supabase-go"
)

func InsertSomeData(client *supabase.Client, data *[]dto.StudentForCreationDto) {

	var results []dto.StudentForCreationDto
	err := client.DB.From("Students").Insert(data).Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println("##INSERTED DATA##")
	fmt.Println(results) // Inserted rows

}
