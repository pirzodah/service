package afterconnect

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func AfterConnectService(ctx context.Context, conn *pgx.Conn) (err error) {
	_, err = conn.Prepare(ctx, "addService", `
		INSERT INTO services(name)
		VALUES ($1)
		RETURNING id;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "getServiceByID", `
		SELECT id, name
		FROM services
		WHERE id = $1;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "updateService", `
		UPDATE services
		SET name = $1
		WHERE id = $2;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "deleteService", `
		DELETE FROM services
		WHERE id = $1;
	`)
	if err != nil {
		return err
	}

	return nil
}
