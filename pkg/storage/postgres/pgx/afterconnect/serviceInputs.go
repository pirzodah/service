package afterconnect

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func AfterConnectServiceInputs(ctx context.Context, conn *pgx.Conn) (err error) {

	_, err = conn.Prepare(ctx, "addServiceInput", `
		INSERT INTO service_inputs(input_id, service_id, required)
		VALUES ($1, $2, $3);
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "getServiceInputsByServiceID", `
		SELECT id, service_id, required
		FROM service_inputs
		WHERE service_id = $1;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "getInputIDsByServiceID", `
		SELECT input_id
		FROM service_inputs
		WHERE service_id = $1;
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

	_, err = conn.Prepare(ctx, "getServiceByID", `
		SELECT id, name
		FROM services
		WHERE id = $1;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "updateServiceInput", `
		UPDATE service_inputs
		SET input_id = $1, service_id = $2, required = $3
		WHERE id = $4;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "deleteServiceInput", `
		DELETE FROM service_inputs
		WHERE id = $1;
	`)
	if err != nil {
		return err
	}

	return nil
}
