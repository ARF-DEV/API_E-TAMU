package seeders

import "database/sql"

func MigrateDB(db *sql.DB) {

	MigrateUser(db)
	MigrateVisit(db)
}
