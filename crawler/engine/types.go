package engine

// Request type
type Request struct {
	Url	string
	ParserFunc func([]byte) ParserResult
}

// ParserResult type
type ParserResult struct {
	Requests []Request
	Items []interface{}
}

// NilParser func
func NilParser([]byte) ParserResult {
	return ParserResult{}
}