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
