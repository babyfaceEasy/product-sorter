package sorter

import "github.com/babyfaceeasy/product-sorter/internal/models"

type Sorter interface {
	Sort(products []models.Product) []models.Product
}
