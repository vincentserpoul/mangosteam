package tradeoffer

func createMockEmptyConfirmation() string {
	return `
    <!DOCTYPE html>
    <html class=" responsive">
    <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width,initial-scale=1">
        <meta name="theme-color" content="#171a21">
        <title>Steam Community :: Confirmations</title>
    </head>
    <body class=" responsive_page">
        <div class="responsive_page_template_content">

                <div id="mobileconf_empty" class="mobileconf_done mobileconf_header">
                    <div>Nothing to confirm</div>
                    <div>You don't have anything to confirm right now.</div>
                </div>

                <div id="mobileconf_details" style="display: none">
                </div>

            <div id="mobileconf_buttons" style="display: none">
                <div>
                <div class="mobileconf_button mobileconf_button_cancel">
                </div><div class="mobileconf_button mobileconf_button_accept">
                </div>
                </div>
            </div>

            <div id="mobileconf_throbber" style="display: none">
                <div style="text-align:center; margin: auto;">
                    <img src="https://steamcommunity-a.akamaihd.net/public/images/login/throbber.gif" alt="Loading">
                </div>
            </div>
        </div>  <!-- responsive_page_legacy_content -->
    </body>
    </html>`
}

func createMockMultiConfirmation() string {
	return `
    <!DOCTYPE html>
    <html class=" responsive">
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width,initial-scale=1">
        <meta name="theme-color" content="#171a21">
        <title>Steam Community :: Confirmations</title>
    </head>
    <body class=" responsive_page">
        <div class="responsive_page_template_content">
            <div id="mobileconf_list">
                <div class="mobileconf_list_entry" id="conf123456" data-confid="123456" data-key="123456789" data-cancel="Cancel" data-accept="Send Offer">
                    <div class="mobileconf_list_entry_content">
                        <div class="mobileconf_list_entry_description">
                        <div>Trade with Bob</div>
                        <div>You will receive nothing</div>
                        <div>Just now</div>
                    </div>
                </div>
                <div class="mobileconf_list_entry" id="conf1234567" data-confid="1234567890" data-key="1234567890" data-cancel="Cancel" data-accept="Send Offer">
                    <div class="mobileconf_list_entry_content">
                        <div class="mobileconf_list_entry_description">
                        <div>Trade with Rob</div>
                        <div>You will receive nothing</div>
                        <div>Just now</div>
                    </div>
                </div>
            </div>
            <div id="mobileconf_done" class="mobileconf_done mobileconf_header" style="display: none">
            <div>All done</div>
            <div>You're all done, there's nothing left to confirm.</div>
        </div>
    </body>
    </html>`
}
