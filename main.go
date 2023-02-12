package main

import (
	"Task_1/dbOperations"
	"Task_1/dto"
	supa "github.com/nedpals/supabase-go"
)

func main() {
	supaBaseUrl := "https://yidsnxqhpectuanswypc.supabase.co"
	supaBaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InlpZHNueHFocGVjdHVhbnN3eXBjIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTY3NjE1MjExMSwiZXhwIjoxOTkxNzI4MTExfQ.m3FAlEegaAFqajDn2zUiiZx6cTQ-_5qNLznm4bO909U"
	supaBaseClient := supa.CreateClient(supaBaseUrl, supaBaseKey)

	data := &[]dto.StudentForCreationDto{
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
			Name:    "Ali",
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

	//###DEV INSIGHTS###
	//I used the given library for supabase database operations. (github.com/nedpals/supabase-go)
	//It has a lots of implementations on supabase and its features but it is not well documented.
	//I found the features of it by trying myself. Maybe we can create our own repo methods
	//for postqres by connecting to database of supabase with given credentials.

	//Getting all data with different columns
	DbOperations.GetAllData(supaBaseClient)
	//Getting data with id
	DbOperations.FindById(supaBaseClient, "1")
	DbOperations.FindById(supaBaseClient, "45")
	DbOperations.FindById(supaBaseClient, "40")
	//Getting data with given column and value
	DbOperations.FindByGivenColumn(supaBaseClient, "Name", "Ali")
	DbOperations.FindByGivenColumn(supaBaseClient, "Surname", "Ayden")
	DbOperations.FindByGivenColumn(supaBaseClient, "Number", "3")

}
