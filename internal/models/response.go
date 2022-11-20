package models

type ResponseMeta struct {
	Count int64 `json:"count"`
}

type ResponseWithPagination struct {
	Data interface{}  `json:"data"`
	Meta ResponseMeta `json:"_meta"`
}
