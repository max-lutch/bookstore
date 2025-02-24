package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := bookstore.Book{ID: 1, Title: "For the Love of Go"}
	got := bookstore.GetBook(catalog, 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
