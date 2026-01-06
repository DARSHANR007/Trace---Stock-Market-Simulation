package market

import "database/sql"

var InstrumentDB *sql.DB

func SetInstrumentDB(db *sql.DB) {
	InstrumentDB = db
}
