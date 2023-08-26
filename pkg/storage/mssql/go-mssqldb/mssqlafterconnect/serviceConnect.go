package mssqlafterconnect

import (
	"context"
	"database/sql"
)

func AfterConnectService(ctx context.Context, db *sql.DB) (err error) {
	_, err = db.PrepareContext(ctx, "addService", `
		INSERT INTO services(name)
		VALUES (?)
		OUTPUT INSERTED.id;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "getServiceByID", `
		SELECT id, name
		FROM services
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "updateService", `
		UPDATE services
		SET name = ?
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "deleteService", `
		DELETE FROM services
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	return nil
}
