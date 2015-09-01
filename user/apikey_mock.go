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
                        <p>Key: 1234567890123456789012</p>
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
