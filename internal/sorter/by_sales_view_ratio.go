package sorter

import (
	"slices"
	"github.com/babyfaceeasy/product-sorter/internal/models"
)

type BySalesViewRatio struct{}

func (s BySalesViewRatio) Sort(products []models.Product) []models.Product {
	sorted := make([]models.Product, len(products))
	copy(sorted, products)
	slices.SortFunc(sorted, func(a, b models.Product) int {
		ar := float64(a.SalesCount) / float64(max(1, a.ViewsCount))
		br := float64(b.SalesCount) / float64(max(1, b.ViewsCount))
		if ar > br {
			return -1
		} else if ar < br {
			return 1
		}
		return 0
	})
	return sorted
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
