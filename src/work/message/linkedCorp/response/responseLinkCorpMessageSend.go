package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpMessageSend struct {
	response.ResponseWork

	InvalidUser  []string `json:"invaliduser"`
	InvalidParty []string `json:"invalidparty"`
	InvalidTag   []string `json:"invalidtag"`
}
