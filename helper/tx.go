package helper

import (
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {

		err := tx.Rollback().Error
		PanicIfError(err)
		panic(err)

	} else {
		errCommit := tx.Commit().Error
		PanicIfError(errCommit)

	}

}
