package checkout

type Checkout interface {
	Scan(item string)
	CalculateTotal() int
}

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
	var unrecognizedSKUList []string

	for _, item := range items {
		_, exists := sh.stockList[item]
		if !exists {
			unrecognizedSKUList = append(unrecognizedSKUList, item)
			continue
		}
		sh.scanned[item]++
	}

	return unrecognizedSKUList
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
