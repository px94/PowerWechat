package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGoodsAudit struct {
	response.ResponseMiniProgram

	AuditID int `json:"auditId"`
}
