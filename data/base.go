package data

import (
	"regifted/util/mylog"
)

const LOGGER_NAME = "reader"
const LOGGER_SEVERITY_LEVEL = mylog.SEV_TRACE

var logger mylog.Logger

func InitLogger() {
	logger = mylog.CreateLogger(LOGGER_NAME)
	logger.SetSeverityThresh(LOGGER_SEVERITY_LEVEL)
}
