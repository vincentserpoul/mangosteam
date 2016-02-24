package tradeoffer

func cancelMockSteamTradeOffer() string {

	return `
        {
            "tradeofferid":"123456"
		}
		`
}

func cancelWeird500Success16SteamTradeOffer() string {

	return `
        {
            "success":16
        }
        `
}

func cancelWeird500Success11MockSteamTradeOffer() string {

	return `
        {
            "success":11
        }
        `
}

func cancel500MockSteamTradeOffer() string {

	return `
        {
            "test":11
        }
        `
}
