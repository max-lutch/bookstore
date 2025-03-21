package bookstore_test

import (
	"bookstore"
	"bookstore/creditcard"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	want := "Autobiography"
	err := b.SetCategory(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.Category()
	if !cmp.Equal(want, got,
		cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf("want category %q, got %q", want, got)
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory("bogus")
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}

func TestNew(t *testing.T) {
	t.Parallel()
	want := "123456789"
	cc, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}
	got := cc.Number
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
