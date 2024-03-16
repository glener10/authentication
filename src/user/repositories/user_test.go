package user_repository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/glener10/rotating-pairs-back/src/db"
	postgres_db "github.com/glener10/rotating-pairs-back/src/db/postgres"
	user_dtos "github.com/glener10/rotating-pairs-back/src/user/dtos"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	pg_container, err := postgres_db.UpTestContainerPostgres()
	if err != nil {
		log.Fatalf(err.Error())
	}
	connStr, err := postgres_db.ReturnTestContainerConnectionString(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	db.ConnectDb(*connStr, "file://src/db/migrations")
	exitCode := m.Run()
	err = postgres_db.DownTestContainerPostgres(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	os.Exit(exitCode)
}

func TestCreateUser(t *testing.T) {
	userDto := user_dtos.CreateUserRequest{
		Email:    "roi@roi.com",
		Password: "aasd12y37asd#8",
	}
	user, err := CreateUser(userDto)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(user)
	assert.Equal(t, 2, 2, "The create object need to be equal to expected object")
}
