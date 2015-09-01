package steamuser

func getMockExistingAPIKeyPage() string {
	return `
        <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
        <html>
        <head>
        <title>Steam</title>
        </head>
        <body>
            <!-- main body -->
            <div id="BG_bottom">
                <div id="bodyContents_ex">

                        <h2>Your Steam Web API Key</h2>
                        <p>Key: 01234567890123456789012345678901</p>
                        <p>Domain name: motus.com</p>
                        <form class="smallForm" id="editForm" name="editForm" method="POST" action="http://fakeurl.com/dev/revokekey" onSubmit="return verifyRevoke()">
                        <input type="submit" name="Revoke" value="">
                        <input type="hidden" name="sessionid" value="fakesession">
                        </form>


                </div>
            </div>
        </body>
        </html>
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

func getMockEmptyAPIKeyPage() string {
	return `
        <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
        <html>
        <head>
        <title>Steam</title>
        </head>
        <body>
            <!-- main body -->
            <div id="BG_bottom">
                <div id="bodyContents_ex">

                        <h2>Your Steam Web API Key</h2>
                        <p>Key: </p>
                        <p>Domain name: motus.com</p>
                        <form class="smallForm" id="editForm" name="editForm" method="POST" action="http://fakeurl.com/dev/revokekey" onSubmit="return verifyRevoke()">
                        <input type="submit" name="Revoke" value="">
                        <input type="hidden" name="sessionid" value="fakesession">
                        </form>


                </div>
            </div>
        </body>
        </html>
    `
}
func getMockAccessDeniedgetAPIKey() string {

	return `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
</head>
<body>
<div id="mainContents">
<h2>Access Denied</h2>
</div>
</body>
</html>
`
}
