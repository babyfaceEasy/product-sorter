package sorter

import (
	"slices"
	"github.com/babyfaceeasy/product-sorter/internal/models"
)

type ByPrice struct{}

func (s ByPrice) Sort(products []models.Product) []models.Product {
	sorted := make([]models.Product, len(products))
	copy(sorted, products)
	slices.SortFunc(sorted, func(a, b models.Product) int {
		if a.Price < b.Price {
			return -1
		} else if a.Price > b.Price {
			return 1
		}
		return 0
	})
	return sorted
}
