package gx

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	c := New("lainchan.org")
	c.UserAgent = "gx v0.0.0/test"

	thread, err := c.GetThread("λ", 31169)
	if err != nil {
		t.Error(err)
	}

	if !strings.HasPrefix(thread.Posts[0].Comment, "This is the") {
		t.Error("Body does not match.")
	}

}
