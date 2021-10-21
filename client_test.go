package gx

import "testing"

func Test(t *testing.T) {
	c := New("lainchan.org")
	c.UserAgent = "gx v0.0.0/test"

	catalog, err := c.GetCatalog("λ")
	if err != nil {
		t.Fatalf(err.Error())
	}

	_ = catalog.Pages[0].Threads[0]

}
