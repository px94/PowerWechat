package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetUserAllLivingID struct {
	response.ResponseWork

	NextCursor   string   `json:"next_cursor"`
	LivingIDList []string `json:"livingid_list"`
}
