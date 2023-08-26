package mssqlafterconnect

import (
	"context"
	"database/sql"
)

func AfterConnectInputs(ctx context.Context, db *sql.DB) (err error) {
	_, err = db.PrepareContext(ctx, "addInput", `
		INSERT INTO inputs(name, type_inputs)
		VALUES (?, ?)
		OUTPUT INSERTED.id;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "getInputByID", `
		SELECT id, name, type_inputs
		FROM inputs
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "updateInput", `
		UPDATE inputs
		SET name = ?, type_inputs = ?
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "deleteInput", `
		DELETE FROM inputs
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	return nil
}
