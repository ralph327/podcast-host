package test

import (
	"codenex.us/ralph/podcast-host"
	_ "codenex.us/ralph/podcast-host/mock"
	"codenex.us/ralph/podcast-host/system"
	"fmt"
	"github.com/solher/arangolite"
	"testing"
)

func TestCreateUser(t *testing.T) {
	var db_s system.UserService

	t.Log("In Create")

	db_s.DB = arangolite.New().
		LoggerOptions(false, false, false).
		Connect("http://devpod.codenex.us:8529", "podcast-host", "devpod", "**DB!!podhost")

	user := new(podhost.User)

	user.FirstName = "Ralph"
	user.LastName = "Martinez"
	user.Active = true

	if err := db_s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUser(t *testing.T) {
	var db_s system.UserService

	t.Log("In Delete")

	db_s.DB = arangolite.New().
		LoggerOptions(false, false, false).
		Connect("http://devpod.codenex.us:8529", "podcast-host", "devpod", "**DB!!podhost")

	if err := db_s.DeleteUser("239003"); err != nil {
		t.Fatal(err)
	}
}

func TestUser(t *testing.T) {
	var db_s system.UserService

	t.Log("In User")

	db_s.DB = arangolite.New().
		LoggerOptions(false, false, false).
		Connect("http://devpod.codenex.us:8529", "podcast-host", "devpod", "**DB!!podhost")

	u, err := db_s.User("239003")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("|Key: %s| |FirstName: %s| |LastName: %s| \n", *u.Key, u.FirstName, u.LastName)
}
