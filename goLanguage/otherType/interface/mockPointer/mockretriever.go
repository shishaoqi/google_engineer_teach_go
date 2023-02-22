package mockPointer

type Retriever struct {
	Contents string
}

// 指针接收者
func (r *Retriever) Get(url string) string {
	return r.Contents
}

// 指针接收者
func (r *Retriever) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]

	return "ok"
}

/*func (r *Retriever) String() string {
	return fmt.Sprintf("toString prinf Retriever: {Contents=%s}", r.Contents)
}*/
