package sorter

import (
	"testing"

	"github.com/babyfaceeasy/product-sorter/internal/utils"
)

func TestByPrice(t *testing.T) {
	products := utils.SampleProducts()
	s := ByPrice{}
	sorted := s.Sort(products)

	if sorted[0].Name != "Coffee Table" {
		t.Errorf("Expected 'Coffee Table' to be first, got %s", sorted[0].Name)
	}
}

func TestBySalesViewRatio(t *testing.T) {
	products := utils.SampleProducts()
	s := BySalesViewRatio{}
	sorted := s.Sort(products)

	if sorted[0].Name != "Zebra Table" {
		t.Errorf("Expected 'Zebra Table' to be first, got %s", sorted[0].Name)
	}
}
