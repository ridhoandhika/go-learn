package dto

type ErrorMessage struct {
	Id string `json:"in"`
	En string `json:"en"`
}

// Struct untuk error_schema
type ErrorSchema struct {
	ErrorCode    string       `json:"error_code"`
	ErrorMessage ErrorMessage `json:"error_message"`
}

type BaseResp struct {
	ErrorSchema  ErrorSchema `json:"error_schema"`
	OutputSchema interface{} `json:"output_schema,omitempty"` // OutputSchema fleksibel (bisa berisi apa saja)
}
