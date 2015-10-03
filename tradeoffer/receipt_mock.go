package tradeoffer

func getMockOKReceipt() string {
	return `
        <!DOCTYPE html>
        <html>
            <head>
                <title>Échange effectué</title>
            </head>
            <body>
                <script type="text/javascript">
                var g_bIsTrading = true;
                var g_rgAppContextData = {"570":{"appid":570,"name":"Dota 2","icon":"http:\/\/cdn.akamai.steamstatic.com\/steamcommunity\/public\/images\/apps\/570\/0bbb630d63262dd66d2fdd0f7d37e8661a410075.jpg","link":"http:\/\/steamcommunity.com\/app\/570","asset_count":152,"inventory_logo":"http:\/\/cdn.akamai.steamstatic.com\/steamcommunity\/public\/images\/apps\/570\/910ef16cd7bf6c6986e78b3ad4eee7eaa5d26cc0.png","trade_permissions":"FULL"},"730":{"appid":730,"name":"Counter-Strike: Global Offensive","icon":"http:\/\/cdn.akamai.steamstatic.com\/steamcommunity\/public\/images\/apps\/730\/69f7ebe2735c366c65c0b33dae00e12dc40edbe4.jpg","link":"http:\/\/steamcommunity.com\/app\/730","asset_count":7,"inventory_logo":"http:\/\/cdn.akamai.steamstatic.com\/steamcommunity\/public\/images\/apps\/730\/3ab6e87a04994b900881f694284a75150e640536.png","trade_permissions":"FULL"},"753":{"appid":753,"name":"Steam","icon":"http:\/\/cdn.akamai.steamstatic.com\/steamcommunity\/public\/images\/apps\/753\/135dc1ac1cd9763dfc8ad52f4e880d2ac058a36c.jpg","link":"http:\/\/steamcommunity.com\/app\/753","asset_count":6,"inventory_logo":"http:\/\/cdn.akamai.steamstatic.com\/steamcommunity\/public\/images\/apps\/753\/db8ca9e130b7b37685ab2229bf5a288aefc3f0fa.png","trade_permissions":"FULL"}};
                var g_rgWalletInfo = {"wallet_currency":13,"wallet_country":"SG","wallet_fee":1,"wallet_fee_minimum":1,"wallet_fee_percent":"0.05","wallet_publisher_fee_percent_default":"0.10","wallet_fee_base":0,"wallet_other_currency":3,"wallet_other_country":"FR","wallet_fee_base_for_other":0,"wallet_conversion_rate":"0.618118","wallet_inverse_conversion_rate":"1.617814","wallet_balance":42436,"wallet_delayed_balance":0,"wallet_max_balance":70000,"wallet_trade_max_balance":56000,"success":true};
                UserYou.LoadContexts( g_rgAppContextData );
                var oItem;
                    oItem = {"id":"1234","owner":"12345678999","classid":"1234567","instanceid":"0","icon_url":"-9a81dl","icon_url_large":"-9a81dlWMRkL5","icon_drag_url":"","name":"MOllusk 12 | 345","market_hash_name":"MOllusk 12 | 345","market_name":"MOllusk 12 | 345","name_color":"D2D2D2","background_color":"","type":"Mesh","tradable":1,"marketable":1,"commodity":0,"market_tradable_restriction":"7","descriptions":[{"type":"html","value":"dfdsfsdf"},{"type":"html","value":" "},{"type":"html","value":"fsefesf"},{"type":"html","value":" "},{"type":"html","value":"Collection Bank","color":"34324324","app_data":{"def_index":"432432","is_itemset_name":1}},{"type":"html","value":" ","app_data":{"def_index":"3432432"}},{"type":"html","value":" "}],"owner_descriptions":"","tags":[{"internal_name":"fsefes","name":"PM","category":"Type","category_name":"Type"},{"internal_name":"efsefff","name":"esfs","category":"esfsef","category_name":"seffesf"},{"internal_name":"sefgggse","name":"efse","category":"fesf","category_name":"esfsef"},{"internal_name":"fsegge","name":"efsef","category":"fsefes","category_name":"fesfse"},{"internal_name":"fesfgggsgses","name":"esfesf","category":"34324","color":"b0c3d9","category_name":"3243242"}],"pos":1,"appid":730,"contextid":2};
                    oItem.appid = 730;
                    oItem.contextid = 2;
                    oItem.amount = 1;
                    oItem.is_stackable = oItem.amount > 1;
                    BuildHover( 'item0', oItem, UserYou );
                    $('item0').show();
                    oItem = {"id":"1235","owner":"12345678999","classid":"1234568","instanceid":"0","icon_url":"-9a81dlVdw8","icon_url_large":"-9a81dEbQ","icon_drag_url":"","name":"AKlove 12 | 3","market_hash_name":"AKlove 12 | 3","market_name":"AKlove 12 | 3","name_color":"32323DD","background_color":"","type":"AKlove 12 | 3","tradable":1,"marketable":1,"commodity":0,"market_tradable_restriction":"7","descriptions":[{"type":"wdaw","value":"wdawd"},{"type":"html","value":" "},{"type":"html","value":""},{"type":"html","value":" "}],"owner_descriptions":"","tags":[{"internal_name":"fsfsef","name":"fsfsef","category":"gesgs","category_name":"htdthr"}],"pos":2,"appid":730,"contextid":2};
                    oItem.appid = 730;
                    oItem.contextid = 2;
                    oItem.amount = 1;
                    oItem.is_stackable = oItem.amount > 1;
                    BuildHover( 'item1', oItem, UserYou );
                    $('item1').show();
                    oItem = {"id":"1236","owner":"12345678999","classid":"1234569","instanceid":"0","icon_url":"-9a8-E_","icon_url_large":"-9a81dlWc","icon_drag_url":"","name":"TG43 | Elpiero","market_hash_name":"TG43 | Elpiero","market_name":"TG43 | Elpiero","name_color":"rwerwr","background_color":"","type":"TG43 | Elpiero","tradable":1,"marketable":1,"commodity":0,"market_tradable_restriction":"7","descriptions":[{"type":"html","value":"awdawdwadawdawd"},{"type":"html","value":" "},{"type":"html","value":"wadawdawdwa"},{"type":"html","value":" "},{"type":"html","value":"awdawdawda","color":"dawawd","app_data":{"def_index":"awdawd","is_itemset_name":1}},{"type":"html","value":" ","app_data":{"def_index":"dwada"}},{"type":"html","value":" "}],"owner_descriptions":"","tags":[{"internal_name":"wdaawd","name":"dwdawdawd","category":"wdawwa","category_name":"dawawd"},{"internal_name":"dwdwd","name":"dwdwdwd","category":"daww","category_name":"dawdaw"}],"pos":3,"appid":730,"contextid":2};
                    oItem.appid = 730;
                    oItem.contextid = 2;
                    oItem.amount = 1;
                    oItem.is_stackable = oItem.amount > 1;
                    BuildHover( 'item2', oItem, UserYou );
                    $('item2').show();
                    oItem = {"id":"1237","owner":"12345678999","classid":"12345670","instanceid":"7855747","icon_url":"-bgvvnWI1RoN","icon_url_large":"-9qBUup_Omyd","icon_drag_url":"","name":"Mallo | 441","market_hash_name":"Mallo | 441","market_name":"Mallo | 441","name_color":"FRGT","background_color":"","type":"Mallo | 441","tradable":1,"marketable":1,"commodity":0,"market_tradable_restriction":"7","descriptions":[{"type":"html","value":"fsesefse"},{"type":"html","value":" "}],"owner_descriptions":"","tags":[{"internal_name":"gdfgdfg","name":"dfgdfgd","category":"dfgdfgdf","category_name":"gdfgdf"},{"internal_name":"gfgdfgdfgdfg","name":"fgdfgdfgd","category":"dfgdfgdf","category_name":"gdfgdfgdf"}],"pos":4,"appid":730,"contextid":2};
                    oItem.appid = 730;
                    oItem.contextid = 2;
                    oItem.amount = 1;
                    oItem.is_stackable = oItem.amount > 1;
                    BuildHover( 'item3', oItem, UserYou );
                    $('item3').show();
                </script>
            </body>
        </html>
    `
}

func getMockKOReceipt() string {
	return `
                    oItem = {"id":"1234";
    `
}
