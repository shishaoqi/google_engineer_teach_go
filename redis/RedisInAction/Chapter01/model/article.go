package model

import (
	"github.com/go-redis/redis/v7"
	"shishaoGo/redis/RedisInAction/Chapter01/common"
	"strconv"
	"strings"
	"time"
)

type Article interface {
	ArticleVote(string, string)
	PostArticle(string, string, string) string
}

type ArticleObj struct {
	Conn *redis.Client
}

func NewArticleRepo(conn *redis.Client) *ArticleObj {
	return &ArticleObj{Conn: conn}
}

func (a *ArticleObj)ArticleVote(user, article string) {
	cutoff := time.Now().Unix() - common.OneWeekInSeconds
	if a.Conn.ZScore("time:", article).Val() < float64(cutoff) {
		return
	}

	articleId := strings.Split(article, ":")[1]
	if a.Conn.SAdd("voted:"+articleId, user).Val() != 0 {
		a.Conn.ZIncrBy("score:", common.VoteScore, article)
		a.Conn.HIncrBy(article, "votes", 1)
	}
}

func (a *ArticleObj)PostArticle(user, title, link string) string {
	articleId := strconv.Itoa(int(a.Conn.Incr("article:").Val()))

	voted := "voted:" + articleId
	a.Conn.SAdd(voted, user)
	a.Conn.Expire(voted, common.OneWeekInSeconds*time.Second)

	now := time.Now().Unix()
	article := "article:" + articleId
	a.Conn.HMSet(article, map[string]interface{}{
		"title": title,
		"link": link,
		"poster": user,
		"time": now,
		"votes": 1,
	})

	a.Conn.ZAdd("score", &redis.Z{Score: float64(now + common.VoteScore), Member: article})
	a.Conn.ZAdd("time", &redis.Z{Score: float64(now), Member: article})
	return articleId
}