package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInSchedulist struct {
	response.ResponseWork

	ScheduleList []*power.HashMap `json:"schedule_list"`
}
