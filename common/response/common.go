package response

import "github.com/gorilla/schema"

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

var Decoder = schema.NewDecoder()

func init() {
	Decoder.SetAliasTag("json")
}
