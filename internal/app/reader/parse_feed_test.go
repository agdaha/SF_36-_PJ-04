package reader

import (
	"testing"
)

func TestParseRss(t *testing.T) {
	posts, err := ParseRss("https://habr.com/ru/rss/best/daily/?fl=ru")
	if err != nil {
		t.Fatal(err)
	}
	if posts == nil || len(posts) == 0 {
		t.Fatal("данные не рскодированы")
	}
	t.Logf("получено %d новостей", len(posts))

}
