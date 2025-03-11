package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HealthcheckResponse struct {
	Success bool `json:"success"`
}

func RpcHealthcheck(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("Healthcheck RPC called ")
	response := &HealthcheckResponse{
		Success: true,
	}
	out, err := json.Marshal(response)
	if err != nil {
		logger.Error("Failed to marshal healthcheck response: %s", err)
		return "", runtime.NewError("failed to marshal healthcheck response", 13)
	}
	return string(out), nil
}

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
