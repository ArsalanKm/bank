package util

const (
	USD = "USD"
	EUR = "EUR"
	IRR = "IRR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, IRR:
		return true
	}
	return false
}
