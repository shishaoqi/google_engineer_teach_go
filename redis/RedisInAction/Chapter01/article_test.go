package main

import (
	"shishaoGo/redis/RedisInAction/Chapter01/model"
	"testing"
	"shishaoGo/redis/RedisInAction/redisConn"

)

func Test(t *testing.T) {
	conn := redisConn.ConnectRedis()
	article := model.NewArticleRepo(conn)

	articleId := article.PostArticle("username", "A title", "http://www.google.com")
	t.Log("We posted a new article with id:", article)
	assertStringResult(t, "1", articleId)
}

func assertStringResult(t *testing.T, want, get string) {
	t.Helper()
	if want != get {
		t.Errorf("want get %v, actual get %v\n", want, get)
	}
}