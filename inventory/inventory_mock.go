package inventory

// GetMockOKProfilesInventory mocks /profiles/1234567890/inventory/json/730/2
func GetMockOKProfilesInventory() string {
	return `
		{
		  "success": true,
		  "rgInventory": {
		    "8742038": {
		      "id": "8742038",
		      "classid": "77838",
		      "instanceid": "0",
		      "amount": "1",
		      "pos": 1
		    },
		    "172795": {
		      "id": "172795",
		      "classid": "2107773",
		      "instanceid": "0",
		      "amount": "1",
		      "pos": 2
		    }
		  },
		  "rgCurrency": [],
		  "rgDescriptions": {
		    "77838_0": {
		      "appid": "730",
		      "classid": "77838",
		      "instanceid": "0",
		      "icon_url": "-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9Q1LO5kNoBhSQl-fEv2o1t3QXFR6a1wE4uOkKlFm0qvJd2gSvYS3x9nbwfXyZrqBxDkCvZYmjurEpomlilL6ux07YtuiRwA",
		      "icon_drag_url": "",
		      "name": "5 Year Veteran Coin",
		      "market_hash_name": "5 Year Veteran Coin",
		      "market_name": "5 Year Veteran Coin",
		      "name_color": "D2D2D2",
		      "background_color": "",
		      "type": "Extraordinary Collectible",
		      "tradable": 0,
		      "marketable": 0,
		      "commodity": 0,
		      "market_tradable_restriction": "7",
		      "descriptions": [
		        {
		          "type": "html",
		          "value": "Has been a member of the Counter-Strike community for over 5 years."
		        },
		        {
		          "type": "html",
		          "value": " "
		        },
		        {
		          "type": "html",
		          "value": ""
		        }
		      ],
		      "actions": [
		        {
		          "name": "Inspect in Game...",
		          "link": "steam://rungame/730/1202255/+csgo_econ_action_preview%20S%owner_steamid%A%assetid%D437851"
		        }
		      ],
		      "market_actions": [
		        {
		          "name": "Inspect in Game...",
		          "link": "steam://rungame/730/1202255/+csgo_econ_action_preview%20M%listingid%A%assetid%D437851"
		        }
		      ],
		      "tags": [
		        {
		          "internal_name": "CSGO_Type_Collectible",
		          "name": "Collectible",
		          "category": "Type",
		          "category_name": "Type"
		        },
		        {
		          "internal_name": "normal",
		          "name": "Normal",
		          "category": "Quality",
		          "category_name": "Category"
		        },
		        {
		          "internal_name": "Rarity_Ancient",
		          "name": "Extraordinary",
		          "category": "Rarity",
		          "color": "eb4b4b",
		          "category_name": "Quality"
		        }
		      ]
		    },
		    "2107773_0": {
		      "appid": "730",
		      "classid": "2107773",
		      "instanceid": "0",
		      "icon_url": "-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgporrf0e1Y07ODHTjBN_8-JmYWPnuL5feuJwjlVscQhj7rH9tzw2wXmqkNlYG-hJNWSegc9Zl-E_QK9xbjr08Si_MOejgzGL-s",
		      "icon_url_large": "-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgporrf0e1Y07ODHTjBN_8-JmYWPnuL5DLfQhGxUppUp3rvFrI2ljQeyqEM5YGjxdYaccQBsYQvX8lC9xum7gsW_uMvNnyB9-n51RxxVmwc",
		      "icon_drag_url": "",
		      "name": "XM1014 | Blue Spruce",
		      "market_hash_name": "XM1014 | Blue Spruce (Field-Tested)",
		      "market_name": "XM1014 | Blue Spruce (Field-Tested)",
		      "name_color": "D2D2D2",
		      "background_color": "",
		      "type": "Consumer Grade Shotgun",
		      "tradable": 1,
		      "marketable": 1,
		      "commodity": 0,
		      "market_tradable_restriction": "7",
		      "descriptions": [
		        {
		          "type": "html",
		          "value": "Exterior: Field-Tested"
		        },
		        {
		          "type": "html",
		          "value": " "
		        },
		        {
		          "type": "html",
		          "value": "The XM1014 is a powerful fully automatic shotgun that justifies its heftier price tag with the ability to paint a room with lead fast. It has individual parts spray-painted solid colors in a moss color scheme."
		        },
		        {
		          "type": "html",
		          "value": " "
		        },
		        {
		          "type": "html",
		          "value": "The Lake Collection",
		          "color": "9da1a9",
		          "app_data": {
		            "def_index": "65535",
		            "is_itemset_name": 1
		          }
		        },
		        {
		          "type": "html",
		          "value": " "
		        }
		      ],
		      "actions": [
		        {
		          "name": "Inspect in Game...",
		          "link": "steam://rungame/730/1202255/+csgo_econ_action_preview%20S%owner_steamid%A%assetid%D1475"
		        }
		      ],
		      "market_actions": [
		        {
		          "name": "Inspect in Game...",
		          "link": "steam://rungame/730/1202255/+csgo_econ_action_preview%20M%listingid%A%assetid%D1475"
		        }
		      ],
		      "tags": [
		        {
		          "internal_name": "CSGO_Type_Shotgun",
		          "name": "Shotgun",
		          "category": "Type",
		          "category_name": "Type"
		        },
		        {
		          "internal_name": "weapon_xm1014",
		          "name": "XM1014",
		          "category": "Weapon",
		          "category_name": "Weapon"
		        },
		        {
		          "internal_name": "set_lake",
		          "name": "The Lake Collection",
		          "category": "ItemSet",
		          "category_name": "Collection"
		        },
		        {
		          "internal_name": "normal",
		          "name": "Normal",
		          "category": "Quality",
		          "category_name": "Category"
		        },
		        {
		          "internal_name": "Rarity_Common_Weapon",
		          "name": "Consumer Grade",
		          "category": "Rarity",
		          "color": "b0c3d9",
		          "category_name": "Quality"
		        },
		        {
		          "internal_name": "WearCategory2",
		          "name": "Field-Tested",
		          "category": "Exterior",
		          "category_name": "Exterior"
		        }
		      ]
		    }
		  },
		  "more": false,
		  "more_start": false
		}
		`
}

// GetMockOKGetPlayerItems steam API mock response gives you the link between player item_id and def_index
// https://api.steampowered.com/IEconItems_XXXX/GetPlayerItems/v1/?key=XXXXXX&format=json&steamid=XXXXX
func GetMockOKGetPlayerItems() string {
	return `
		{
			"result": {
				"status": 1,
				"num_backpack_slots": 840,
				"items": [
					{
						"id": 1234567894,
						"original_id": 1234567894,
						"defindex": 5470,
						"level": 1,
						"quality": 4,
						"inventory": 19,
						"quantity": 1,
						"equipped": [
							{
								"class": 23,
								"slot": 6
							}
						]
						
					},
					{
						"id": 1234567895,
						"original_id": 1234567895,
						"defindex": 15001,
						"level": 1,
						"quality": 4,
						"inventory": 30,
						"quantity": 1,
						"flag_cannot_trade": true,
						"attributes": [
							{
								"defindex": 153,
								"value": 1065353216,
								"float_value": 1
							},
							{
								"defindex": 16,
								"value": 1,
								"float_value": 1.4012984643248171e-045
							}
						]
						
					},
					{
						"id": 1234567896,
						"original_id": 1234567896,
						"defindex": 10068,
						"level": 1,
						"quality": 4,
						"inventory": 5,
						"quantity": 1,
						"flag_cannot_trade": true,
						"attributes": [
							{
								"defindex": 153,
								"value": 1065353216,
								"float_value": 1
							}
						]
						
					},
					{
						"id": 1234567897,
						"original_id": 1234567897,
						"defindex": 5508,
						"level": 1,
						"quality": 4,
						"inventory": 25,
						"quantity": 1,
						"equipped": [
							{
								"class": 4,
								"slot": 4
							}
						]
						
					},
					{
						"id": 1234567898,
						"original_id": 1234567898,
						"defindex": 7480,
						"level": 1,
						"quality": 4,
						"inventory": 142,
						"quantity": 1,
						"equipped": [
							{
								"class": 15,
								"slot": 0
							}
						]
						,
						"flag_cannot_trade": true,
						"attributes": [
							{
								"defindex": 153,
								"value": 1,
								"float_value": 1.4012984643248171e-045
							},
							{
								"defindex": 213,
								"value": 1,
								"float_value": 1.4012984643248171e-045
							}
						]
						
					}
				]
				
			}
		}
	`
}

// GetMockOKGetAssetPrices gives the match between def_index and class_id
// https://api.steampowered.com/ISteamEconomy/GetAssetPrices/v1/?key=XXXXX&format=json&appid=XXX&currency=usd
func GetMockOKGetAssetPrices() string {
	return `
		{
			"result": {
				"success": true,
				"assets": [
					{
						"prices": {
							"USD": 0
						},
						"name": "4004",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4004"
							}
						]
						,
						"classid": "57939754"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4008",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4008"
							}
						]
						,
						"classid": "57939594"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4009",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4009"
							}
						]
						,
						"classid": "57939591"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4010",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4010"
							}
						]
						,
						"classid": "57939593"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4049",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4049"
							}
						]
						,
						"classid": "93966736"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4097",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4097"
							}
						]
						,
						"classid": "147888890"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4110",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4110"
							}
						]
						,
						"classid": "57939654"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "20886",
						"date": "9/2/2015",
						"class": [
							{
								"name": "def_index",
								"value": "20886"
							}
						]
						,
						"classid": "1218050854"
					}
				]		
			}
		}
	`
}

// GetMockOKGetAssetClassInfo steam API gives you the link between class_id and market_hash_name
// http://api.steampowered.com/ISteamEconomy/GetAssetClassInfo/v0001?key=XXX&format=json&language=en&appid=XXX&class_count=2&classid0=123456789
func GetMockOKGetAssetClassInfo() string {
	return `
		{
			"result": {
				"123456789": {
					"icon_url": "W_I_5GLm4wPcv9jJQ7z7tz_l_0sEIYUhRfbF4arNQkgGQGKd3kMuVpMgCwRZrhSfeEqb1qNMeO7lDgsvJYj2VkHyNb-A-UWkTe9Xc8Rgd2sbj9_ugkgSUXffBrFHXNQrvM7K0Ay7XgXDLWdun9gFgPqagJWGCPPO6UywK3ID03w",
					"icon_url_large": "W_I_5GLm4wPcv9jJQ7z7tz_l_0sEIYUhRfbF4arNQkgGQGKd3kMuVpMgCwRZrhSfeEqb1qNMeO7lDgsvJYj2VkHyNb-A-UWkTe9Xc8RgBmMYzo69mB0TByTSDb8RDYMpupzD1APoW1HCcWFun4wGivufgpfQUqHSrESyJVJuk7o-hPMuyZ4",
					"icon_drag_url": "",
					"name": "Ye Olde Pipe",
					"market_hash_name": "Ye Olde Pipe",
					"market_name": "Ye Olde Pipe",
					"name_color": "D2D2D2",
					"background_color": "",
					"type": "Common Pipe",
					"tradable": "1",
					"marketable": "1",
					"commodity": "0",
					"market_tradable_restriction": "7",
					"market_marketable_restriction": "7",
					"fraudwarnings": "",
					"descriptions": {
						"0": {
							"type": "html",
							"value": "Used By: Kunkka",
							"app_data": ""
						},
						"1": {
							"type": "html",
							"value": " ",
							"app_data": ""
						},
						"2": {
							"type": "html",
							"value": "Armaments of Leviathan",
							"color": "9da1a9",
							"app_data": {
								"def_index": "20267",
								"is_itemset_name": "1"
							}
						},
						"3": {
							"type": "html",
							"value": "Admiral's Foraged Cap",
							"color": "6c7075",
							"app_data": {
								"def_index": "5463"
							}
						},
						"4": {
							"type": "html",
							"value": "Admiral's Stash",
							"color": "6c7075",
							"app_data": {
								"def_index": "5464"
							}
						},
						"5": {
							"type": "html",
							"value": "Claddish Gauntlets",
							"color": "6c7075",
							"app_data": {
								"def_index": "5465"
							}
						},
						"6": {
							"type": "html",
							"value": "Claddish Guard",
							"color": "6c7075",
							"app_data": {
								"def_index": "5466"
							}
						},
						"7": {
							"type": "html",
							"value": "Claddish Hightops",
							"color": "6c7075",
							"app_data": {
								"def_index": "5467"
							}
						},
						"8": {
							"type": "html",
							"value": "Neptunian Sabre",
							"color": "6c7075",
							"app_data": {
								"def_index": "5468"
							}
						},
						"9": {
							"type": "html",
							"value": "Admiral's Salty Shawl",
							"color": "6c7075",
							"app_data": {
								"def_index": "5469"
							}
						},
						"10": {
							"type": "html",
							"value": "Ye Olde Pipe",
							"color": "6c7075",
							"app_data": {
								"def_index": "5470"
							}
						},
						"11": {
							"type": "html",
							"value": "An old pipe Kunkka plucked from a tidepool after being shipwrecked, the old seadog claims it brings him fortune and good luck, and thus never ever takes it from his scurvy-ridden mouth. \r\n\t",
							"app_data": ""
						}
					},
					"tags": {
						"0": {
							"internal_name": "unique",
							"name": "Standard",
							"category": "Quality",
							"color": "D2D2D2",
							"category_name": "Quality"
						},
						"1": {
							"internal_name": "Rarity_Common",
							"name": "Common",
							"category": "Rarity",
							"color": "b0c3d9",
							"category_name": "Rarity"
						},
						"2": {
							"internal_name": "wearable",
							"name": "Wearable",
							"category": "Type",
							"category_name": "Type"
						},
						"3": {
							"internal_name": "neck",
							"name": "Neck",
							"category": "Slot",
							"category_name": "Slot"
						},
						"4": {
							"internal_name": "npc_dota_hero_kunkka",
							"name": "Kunkka",
							"category": "Hero",
							"category_name": "Hero"
						}
					},
					"classid": "123456789"
				},
				"success": true
			}
		}
	`
}
