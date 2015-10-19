package market

func getMockMarket() string {
	return `{"success":true,"lowest_price":"$1.04","volume":"166","median_price":"$1.04"}`
}

func getBadFloatMockMarket() string {
	return `{"success":true,"lowest_price":"$1..04","volume":"1","median_price":"$1.04"}`
}

func getFailedMockMarket() string {
	return `{"success":false}`
}

func getMalformedMockMarket() string {
	return `{success":false}`
}
