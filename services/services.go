package services

import (
	"github.com/SunspotsInys/thedoor/logs"
	"github.com/SunspotsInys/thedoor/utils"
)

var (
	logger = logs.Logger
	idGen  = utils.GetSnowflakeInstance()
)
