package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetDepartmentList struct {
	response.ResponseWork

	department_list []*power.HashMap `json:"department_list"`
}
