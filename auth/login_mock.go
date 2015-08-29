package auth

func getMockKOLoginDologin() string {

	return `
		{
		  "success": false,
		  "requires_twofactor": false,
		  "captcha_needed": false,
		  "captcha_gid": -1,
		  "message": "Informations de connexion incorrectes."
		}
		`
}
func getMockOKLoginDologin() string {

	return `
		{
			"success":true,
			"requires_twofactor":false,
			"login_complete":true,
			"transfer_url":"https:\/\/steamcommunity.com\/login\/transfer",
			"transfer_parameters":
			{
				"steamid":"75894854842678978",
				"token":"2F55FFBBB857485858CC5785966AAA4585585518",
				"auth":"4AF5848554B8547855C8574580852828",
				"remember_login":false,
				"token_secure":"85440F8545A46454D4854848422484248444BFFB"}
			}
		}
		`
}

func getMockEmailauthNeeded() string {

	return `
		{
		  "success": false,
		  "emailauth_needed": true
		}
		`
}

func getMockRespBody() string {

	return `
		{
		  "success": false,,
		  "emailauth_needed: true
		}
		`
}
