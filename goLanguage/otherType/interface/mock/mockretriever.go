/**
 * 与 realretriever 相对应，mock 表示这是个模拟的（用于测试什么的），不是像 realtriever 是真实运行的(现实中的代码)
 */
package mock

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

// 第一版本的 值接收者
/*func (r Retriever) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]

	return "ok"
}*/

// 指针接收者
func (r *Retriever) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]

	return "ok"
}

/*func (r *Retriever) String() string {
	return fmt.Sprintf("toString prinf Retriever: {Contents=%s}", r.Contents)
}*/
