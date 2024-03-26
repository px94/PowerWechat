package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseAgentList struct {
	response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}
