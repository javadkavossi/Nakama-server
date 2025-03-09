package main

import (
	"context"
	"database/sql"
	"encoding/json"

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
