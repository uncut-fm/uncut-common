package entgo

import (
	"fmt"
	"github.com/uncut-fm/uncut-auth-api/ent"
)

func CommitOrRollback(tx *ent.Tx, err error) error {
	if err != nil {
		return Rollback(tx, err)
	}

	return tx.Commit()
}

// Rollback calls to tx.Rollback and wraps the given error
// with the Rollback error if occurred.
func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
