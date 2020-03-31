package mock


// Retriever struct
type Retriever struct {
	Contents string
}

// Get example
func (r *Retriever) Get(url string) string {
	return r.Contents
}

// Post func
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

