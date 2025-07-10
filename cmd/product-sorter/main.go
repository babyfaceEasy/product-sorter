package main

import (
	"encoding/json"
	"fmt"

	"github.com/babyfaceeasy/product-sorter/internal/logger"
	"github.com/babyfaceeasy/product-sorter/internal/models"
	"github.com/babyfaceeasy/product-sorter/internal/sorter"
	"github.com/babyfaceeasy/product-sorter/internal/utils"
	"go.uber.org/zap"
)

func main() {
	// initialize production logger
	logger.Init(true)
	log := logger.L()
	defer log.Sync()

	products := utils.SampleProducts()

	log.Info("Loaded products", zap.Int("count", len(products)))

	// sort by price
	priceSorter := sorter.ByPrice{}
	sortedByPrice := priceSorter.Sort(products)
	log.Info("Sorted products by price")
	printProducts(sortedByPrice)

	// sort by sales/view ratio
	ratioSorter := sorter.BySalesViewRatio{}
	sortedByRatio := ratioSorter.Sort(products)
	log.Info("Sorted products by sales/view ratio")
	printProducts(sortedByRatio)
}

func printProducts(products []models.Product) {
	out, _ := json.MarshalIndent(products, "", "  ")
	fmt.Println(string(out))
}
