package checkout

import (
	"testing"
)

var testStock = map[string]Item{
	"A": {SKU: "A", UnitPrice: 50, MultiBuyQuantity: 3, SpecialPrice: 130},
	"B": {SKU: "B", UnitPrice: 30, MultiBuyQuantity: 2, SpecialPrice: 45},
	"C": {SKU: "C", UnitPrice: 20},
	"D": {SKU: "D", UnitPrice: 15},
	"E": {SKU: "E", UnitPrice: 50, MultiBuyQuantity: 5, SpecialPrice: 200},
	"F": {SKU: "F", UnitPrice: 25, MultiBuyQuantity: 10, SpecialPrice: 225},
}

func TestScan(t *testing.T) {
	testCases := []struct {
		name                 string
		itemsToScan          []string
		expectedTotal        int
		expectedUnrecognized []string
	}{
		{
			name:                 "All items exist with multibuy, no unrecognized SKUs",
			itemsToScan:          []string{"A", "A", "A", "E", "E", "B", "F", "E", "E", "E", "B"},
			expectedTotal:        400,
			expectedUnrecognized: nil,
		},
		{
			name:                 "Some items exist with some unrecognized SKUs",
			itemsToScan:          []string{"B", "B", "C", "D", "D", "B", "Z", "Y", "Z"},
			expectedTotal:        125,
			expectedUnrecognized: []string{"Z", "Y", "Z"},
		},
		{
			name:                 "All scanned items are unrecognized with expected total of 0",
			itemsToScan:          []string{"Z", "X", "Unexpected item in the bagging area"},
			expectedTotal:        0,
			expectedUnrecognized: []string{"Z", "X", "Unexpected item in the bagging area"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// make a new checkout/scanner/basket so each test case is fresh
			basket := New(testStock)

			unrecognizedSKUs := basket.Scan(tc.itemsToScan...)

			if len(unrecognizedSKUs) != len(tc.expectedUnrecognized) {
				t.Errorf(
					"expected %d unrecognized SKUs, got %d",
					len(tc.expectedUnrecognized),
					len(unrecognizedSKUs))
			}

			total := basket.CalculateTotalPrice()
			if total != tc.expectedTotal {
				t.Errorf("expected a total of %d, got a total of %d", tc.expectedTotal, total)
			}
		})
	}
}