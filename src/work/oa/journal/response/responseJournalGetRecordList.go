package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseJournalGetRecordList struct {
	response.ResponseWork

	JournalUUIDList []string `json:"journaluuid_list"`
	NextCursor      int      `json:"next_cursor"`
	EndFlag         bool     `json:"endflag"`
}
