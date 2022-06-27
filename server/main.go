package main

import (
	"go_todo/db"
	model "go_todo/model/db"
	"log"
)

func main() {

	roles := []model.Role{
		{
			Name:        "role1",
			AccessLevel: 1,
		},
		{
			Name:        "role2",
			AccessLevel: 2,
		},
		{
			Name:        "role3",
			AccessLevel: 3,
		},
	}

	gdb, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Couldn't connect to database %v", err)
	}

	err = gdb.Create(&roles).Error
	if err != nil {
		log.Fatalln(err)
	}

}
