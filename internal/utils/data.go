package utils

import (
	"time"

	"github.com/babyfaceeasy/product-sorter/internal/models"
)

func SampleProducts() []models.Product {
	layout := "2006-01-02"
	return []models.Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: mustParse(layout, "2019-01-04"), SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: mustParse(layout, "2012-01-04"), SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: mustParse(layout, "2014-05-28"), SalesCount: 1048, ViewsCount: 20123},
	}
}

func mustParse(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}
