package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseNearbyPoiAdd struct {
	response.ResponseMiniProgram
	Data []*power.HashMap `json:"data"`
}
