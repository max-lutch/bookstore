package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}
	for _, cat := range cats {
		err := b.SetCategory(cat)
		if err != nil {
			t.Fatal(err)
		}
		got := b.Category()
		if cat != got {
			t.Errorf("want category %q, got %q", cat, got)
		}
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory(999)
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	func TestSetPriceCents(t *testing.T) {
		t.Parallel()
		b := bookstore.Book{
			Title:      "For the Love of Go",
			PriceCents: 4000,
		}
		want := 3000
		err := b.SetPriceCents(want)
		if err != nil {
			t.Fatal(err)
		}
		got := b.PriceCents
		if want != got {
			t.Errorf("want updated price %d, got %d", want, got)
		}
	}
