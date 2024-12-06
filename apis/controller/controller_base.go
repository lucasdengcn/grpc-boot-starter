package controller

import (
	"context"
	"fmt"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/infra/db"
)

// ControllerBase define
type ControllerBase struct{}

func (c *ControllerBase) deferTxCallback(ctx context.Context, r any) error {
	if r != nil {
		db.RollbackTx(ctx)
		if err, ok := r.(error); ok {
			logging.Error(ctx).Msgf("Recover Tx Err: %v", r)
			return err
		} else {
			logging.Error(ctx).Msgf("Recover Tx Err: %v", r)
			return fmt.Errorf("%v", r)
		}
	} else {
		db.CommitTx(ctx)
		return nil
	}
}
