package helpers

import "database/sql"

func CommitOrRollback(transaction *sql.Tx) {

	err := recover()
	if err != nil {
		errorRollback := transaction.Rollback()
		PanicOnError(errorRollback)

		panic(err)
	}

	errorCommit := transaction.Commit()
	PanicOnError(errorCommit)

}
