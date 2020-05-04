package engine

// Request type
type Request struct {
	Url	string
	ParserFunc func([]byte) ParserResult
}

// Item type
type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

// ParserResult type
type ParserResult struct {
	Requests []Request
	Items []Item
}

// NilParser func
func NilParser([]byte) ParserResult {
	return ParserResult{}
}