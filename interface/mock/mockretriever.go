package mock

import "fmt"

type Retriever struct {
	Contents string
}
// 大家都改成 指针接收者
/*func (r Retriever) Get(url string) string {
	return r.Contents
}*/

func (r *Retriever) Get(url string) string {
	return r.Contents
}

// 第一版本 值接收者
/*func (r Retriever) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]

	return "ok"
}*/

func (r *Retriever) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]

	return "ok"
}

func (r *Retriever) String() string {
	return fmt.Sprintf("toString prinf Retriever: {Contents=%s}", r.Contents)
}
