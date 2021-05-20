package post

type Post struct {
	Url string
	From map[string]string
}

func (p Post)Post(url string, from map[string]string) string {
	return p.Url
}
