package mssqlafterconnect

import (
	"context"
	"database/sql"
)

func AfterConnectServiceInputs(ctx context.Context, db *sql.DB) (err error) {
	_, err = db.PrepareContext(ctx, "getServiceInputsByServiceID", `
		SELECT id, service_id, required
		FROM service_inputs
		WHERE service_id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "getInputIDsByServiceID", `
		SELECT input_id
		FROM service_inputs
		WHERE service_id = ?;
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

	_, err = db.PrepareContext(ctx, "getServiceByID", `
		SELECT id, name
		FROM services
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "updateServiceInput", `
		UPDATE service_inputs
		SET input_id = ?, service_id = ?, required = ?
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = db.PrepareContext(ctx, "deleteServiceInput", `
		DELETE FROM service_inputs
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	return nil
}
