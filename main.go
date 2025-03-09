package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializers runtime.Initializer) error {
    initStart := time.Now()
    err := initializers.RegisterRpc("healthcheck", RpcHealthcheck)
    if err != nil {
        logger.Error("ثبت RPC healthcheck با مشکل مواجه شد: %s", err)
        return err
    }
    logger.Info("ماژول با موفقیت در %s بارگذاری شد", time.Since(initStart).Milliseconds())
    return nil
}


