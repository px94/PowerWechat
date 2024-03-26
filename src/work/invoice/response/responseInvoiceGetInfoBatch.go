package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseInvoiceGetInfoBatch struct {
	response.ResponseWork

	ItemList []*power.HashMap `json:"item_list"`
}
