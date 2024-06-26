package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetWatchStat struct {
	response.ResponseWork

	NextKey  string         `json:"next_key"`
	StatInfo *power.HashMap `json:"stat_info"`
}
