package afterconnect

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func AfterConnectInputs(ctx context.Context, conn *pgx.Conn) (err error) {

	_, err = conn.Prepare(ctx, "addInput", `
		INSERT INTO inputs(name, type_inputs)
		VALUES ($1, $2)
		RETURNING id;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "getInputByID", `
		SELECT id, name, type_inputs
		FROM inputs
		WHERE id = $1;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "updateInput", `
		UPDATE inputs
		SET name = $1, type_inputs = $2
		WHERE id = $3;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "deleteInput", `
		DELETE FROM inputs
		WHERE id = $1;
	`)
	if err != nil {
		return err
	}

	return nil
}
