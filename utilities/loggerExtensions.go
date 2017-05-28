package utilities

import (
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"go.uber.org/zap"
)

//LogError wrapper for zap.SugaredLogger.Error
func LogError(err error, action string, elapsedTime time.Duration, correlationID *uuid.UUID, logger *zap.Logger) {
	logger.Error(err.Error(),
		zap.String("Action", action),
		zap.String("CorrelationID", correlationID.String()),
		zap.String("elapsed time", elapsedTime.String()))
}

//LogInfo wrapper for zap.SugaredLogger.Info
func LogInfo(action string, elapsedTime time.Duration, correlationID *uuid.UUID, logger *zap.Logger) {
	logger.Info("Success",
		zap.String("Action", action),
		zap.String("CorrelationID", correlationID.String()),
		zap.String("elapsed time", elapsedTime.String()))
}
