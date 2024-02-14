package reader

import (
	"testing"
)

func TestGetRss(t *testing.T) {
	feed, err := GetRss("https://habr.com/ru/rss/best/daily/?fl=ru")
	if err != nil {
		t.Fatal(err)
	}
	if feed == nil || len(feed.Chanel.Items) == 0 {
		t.Fatal("данные не рскодированы")
	}
	t.Logf("получено %d новостей", len(feed.Chanel.Items))
}
