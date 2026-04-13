package pkg

import "errors"

var (
	// ErrDBConnection occurs when some error happen in database connection.
	ErrDBConnection = errors.New("database connection")
	// ErrDBMigration occurs when some error happen in database migration.
	ErrDBMigration = errors.New("database migration")
	// ErrDBTransaction occurs when some error happen in database transaction.
	ErrDBTransaction = errors.New("database transaction")
	// ErrDBInsert occurs when some error happen in database insert.
	ErrDBInsert = errors.New("database insert")
	// ErrDBSelect occurs when some error happen in database select.
	ErrDBSelect = errors.New("database select")
	// ErrDBScan occurs when some error happen in database scan.
	ErrDBScan = errors.New("database scan")
	// ErrDBUpdate occurs when some error happen in database update.
	ErrDBUpdate = errors.New("database update")
	// ErrDBDelete occurs when some error happen in database delete.
	ErrDBDelete = errors.New("database update")
)
