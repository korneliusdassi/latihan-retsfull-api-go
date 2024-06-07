package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			return
		}
		panic(err)
	} else {
		errorCallback := tx.Commit()
		PanicIfError(errorCallback)
	}
}
