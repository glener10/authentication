package promote_user_admin_controller

import (
	"testing"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	"github.com/glener10/authentication/tests"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}
