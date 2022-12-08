package db

import (
	"wellpict-be/db/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations = []*gormigrate.Migration{

	migrations.CreateUsersTable,
	migrations.CreateAdminsTable,
}
