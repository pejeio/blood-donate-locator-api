package types

type ResponseMeta struct {
	Count int64 `json:"count"`
}

type Response struct {
	Data interface{} `json:"data"`
}
type ResponseWithPagination struct {
	Response
	Meta ResponseMeta `json:"_meta"`
}
