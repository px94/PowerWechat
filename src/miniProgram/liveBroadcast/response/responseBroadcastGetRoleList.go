package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetRoleList struct {
	response.ResponseMiniProgram

	Total int              `json:"total,omitempty"`
	List  []*power.HashMap `json:"list,omitempty"`
}
