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
