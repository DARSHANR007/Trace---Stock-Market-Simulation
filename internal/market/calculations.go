package market

import "math"

// Basic stock market calculation helpers.
// These functions are connected to any others , i still have to check and get the calculations right

// RoundTo rounds a value to the requested number of decimal places.
func RoundTo(value float64, places int) float64 {
	if places < 0 {
		return value
	}

	factor := math.Pow(10, float64(places))
	return math.Round(value*factor) / factor
}

// TotalValue returns the value of a position at a given price.
func TotalValue(quantity, price float64) float64 {
	if quantity <= 0 || price <= 0 {
		return 0
	}
	return quantity * price
}

// TotalCost returns the money spent for a position.
func TotalCost(quantity, price float64) float64 {
	return TotalValue(quantity, price)
}

// AveragePrice returns the average cost per share.
func AveragePrice(totalCost, quantity float64) float64 {
	if quantity <= 0 || totalCost <= 0 {
		return 0
	}
	return totalCost / quantity
}

// WeightedAveragePrice returns the weighted average price for multiple buys.
func WeightedAveragePrice(prices []float64, quantities []float64) float64 {
	if len(prices) == 0 || len(prices) != len(quantities) {
		return 0
	}

	var totalQuantity float64
	var totalCost float64
	for i := range prices {
		if prices[i] <= 0 || quantities[i] <= 0 {
			continue
		}
		totalQuantity += quantities[i]
		totalCost += prices[i] * quantities[i]
	}

	return AveragePrice(totalCost, totalQuantity)
}

// ProfitLoss returns the absolute and percentage profit or loss.
// A positive result means profit; a negative result means loss.
func ProfitLoss(quantity, buyPrice, sellPrice float64) (float64, float64) {
	if quantity <= 0 || buyPrice <= 0 || sellPrice <= 0 {
		return 0, 0
	}

	buyValue := quantity * buyPrice
	sellValue := quantity * sellPrice
	profitLoss := sellValue - buyValue
	profitLossPercent := 0.0
	if buyValue != 0 {
		profitLossPercent = (profitLoss / buyValue) * 100
	}

	return profitLoss, profitLossPercent
}

// GainLossValue returns the gain or loss between two prices for a given quantity.
func GainLossValue(quantity, entryPrice, exitPrice float64) float64 {
	if quantity <= 0 || entryPrice <= 0 || exitPrice <= 0 {
		return 0
	}
	return quantity * (exitPrice - entryPrice)
}

// GainLossPercent returns the percentage gain or loss between two prices.
func GainLossPercent(entryPrice, exitPrice float64) float64 {
	if entryPrice <= 0 || exitPrice <= 0 {
		return 0
	}
	return ((exitPrice - entryPrice) / entryPrice) * 100
}

// ReturnPercentage returns the percentage change from the initial value to the final value.
func ReturnPercentage(initialValue, finalValue float64) float64 {
	if initialValue <= 0 || finalValue < 0 {
		return 0
	}
	return ((finalValue - initialValue) / initialValue) * 100
}

// BreakEvenPrice returns the price at which a position breaks even.
func BreakEvenPrice(totalCost, quantity float64) float64 {
	return AveragePrice(totalCost, quantity)
}

// BreakEvenPriceWithFees returns the break-even price when fees are included.
func BreakEvenPriceWithFees(totalCost, quantity, fees float64) float64 {
	if quantity <= 0 || totalCost < 0 || fees < 0 {
		return 0
	}
	return (totalCost + fees) / quantity
}

// NetProfitLoss returns profit or loss after fees.
func NetProfitLoss(quantity, buyPrice, sellPrice, fees float64) float64 {
	if quantity <= 0 || buyPrice <= 0 || sellPrice <= 0 || fees < 0 {
		return 0
	}
	return (quantity * sellPrice) - (quantity * buyPrice) - fees
}

// QuantityAffordable returns how many shares can be bought with the available capital.
func QuantityAffordable(capital, price float64) float64 {
	if capital <= 0 || price <= 0 {
		return 0
	}
	return capital / price
}

// PriceDifference returns the difference between two prices.
func PriceDifference(currentPrice, previousPrice float64) float64 {
	if currentPrice <= 0 || previousPrice <= 0 {
		return 0
	}
	return currentPrice - previousPrice
}

// PriceDifferencePercent returns the percentage difference between two prices.
func PriceDifferencePercent(currentPrice, previousPrice float64) float64 {
	if currentPrice <= 0 || previousPrice <= 0 {
		return 0
	}
	return ((currentPrice - previousPrice) / previousPrice) * 100
}

// MarketCapitalization returns the company's market value.
func MarketCapitalization(sharePrice, outstandingShares float64) float64 {
	if sharePrice <= 0 || outstandingShares <= 0 {
		return 0
	}
	return sharePrice * outstandingShares
}

// DividendYield returns the dividend yield percentage.
func DividendYield(annualDividend, sharePrice float64) float64 {
	if annualDividend < 0 || sharePrice <= 0 {
		return 0
	}
	return (annualDividend / sharePrice) * 100
}

// SimpleMovingAverage returns the average of the provided values.
func SimpleMovingAverage(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	var total float64
	for _, value := range values {
		if value <= 0 {
			continue
		}
		total += value
	}

	return total / float64(len(values))
}
