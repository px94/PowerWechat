package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUserSimpleList struct {
	response.ResponseWork

	UserList []*power.HashMap `json:"userlist"`
}
