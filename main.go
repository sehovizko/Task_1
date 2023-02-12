package main

import (
	"Task_1/dbOperations"
	"Task_1/entity"
	supa "github.com/nedpals/supabase-go"
)

func main() {
	supaBaseUrl := "https://yidsnxqhpectuanswypc.supabase.co"
	supaBaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InlpZHNueHFocGVjdHVhbnN3eXBjIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTY3NjE1MjExMSwiZXhwIjoxOTkxNzI4MTExfQ.m3FAlEegaAFqajDn2zUiiZx6cTQ-_5qNLznm4bO909U"
	supaBaseClient := supa.CreateClient(supaBaseUrl, supaBaseKey)

	data := &[]entity.Student{
		{
			Name:    "Ali",
			Surname: "Yilmaz",
			Number:  1,
		},
		{
			Name:    "Veli",
			Surname: "Demir",
			Number:  2,
		},
		{
			Name:    "Fatma",
			Surname: "Ozgur",
			Number:  3,
		},
		{
			Name:    "Max",
			Surname: "Ayden",
			Number:  4,
		},
		{
			Name:    "Naz",
			Surname: "Kokcu",
			Number:  5,
		},
	}
	//Inserting some data
	DbOperations.InsertSomeData(supaBaseClient, data)

	//Getting all data with different columns
	DbOperations.GetAllColumns(supaBaseClient)
	DbOperations.GetNames(supaBaseClient)
	DbOperations.GetNumbers(supaBaseClient)
}
