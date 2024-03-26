package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
	"github.com/px94/PowerWeChat/v3/src/work/user/request"
)

type ResponseTagList struct {
	response.ResponseWork

	TagName string                `json:"tagname"`
	TagList []*request.RequestTag `json:"taglist"`
}
