package checkout

import "fmt"

type Item struct {
	SKU              string
	UnitPrice        int
	MultiBuyQuantity int // number of units to trigger SpecialPrice for associated SKU
	SpecialPrice     int
}

type SKUHandler struct {
	stockList map[string]Item
	scanned   map[string]int
}

func New(stockList map[string]Item) *SKUHandler {
	return &SKUHandler{
		stockList: stockList,
		scanned:   make(map[string]int),
	}
}

func (sh *SKUHandler) Scan(items ...string) []string {
	unrecognizedSKUs := make(map[string]int)

	for _, item := range items {
		_, exists := sh.stockList[item]
		if !exists {
			unrecognizedSKUs[item]++
		} else {
			sh.scanned[item]++
		}
	}

	// tally up unrecognized SKUs instead of printing the same SKU multiple times
	// e.g. unrecognizedSKUs = "Z x5, Y x3, X"
	// rather than unrecognized SKUs = "Z Z Z Y Y Y Z Z X"
	var result []string
	for sku, count := range unrecognizedSKUs {
		if count > 1 {
			result = append(result, fmt.Sprintf("%s: x%d", sku, count))
		} else {
			result = append(result, sku)
		}
	}

	return result
}

func (sh *SKUHandler) CalculateTotalPrice() int {
	total := 0

	for sku, count := range sh.scanned {
		item := sh.stockList[sku]

		// discount if SKU has a multibuy price
		if item.MultiBuyQuantity > 0 && count >= item.MultiBuyQuantity {
			// calculate how many times the multibuy pricing can be applied
			offers := count / item.MultiBuyQuantity

			// add price for multibuy offers to total
			total += offers * item.SpecialPrice

			// update count to reflect remaining items after applying multibuy offers
			count %= item.MultiBuyQuantity
		}

		// add remaining items at regular price to total
		total += count * item.UnitPrice
	}

	return total
}
