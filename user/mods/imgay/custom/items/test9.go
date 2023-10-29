package items

import "MT-GO/database"

var Test9 = map[string]*database.CustomItemAPI{
	"test9": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5b432d215acfc4771e1c6624",
			HandbookPrice:    50000,
			ModifierType:     "clone",
			AddToTrader: map[string]*database.CustomItemAddToTrader{
				"58330581ace78e27b8b10cee": {
					LoyaltyLevel: 1,
					BarterScheme: map[string]float32{
						"5449016a4bdc2d6f028b456f": 50000,
					},
					AmountInStock: 999,
				},
			},
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "Sombrero",
				ShortName:   "Sombrero",
				Description: "A traditional and iconic type of wide-brimmed hat that is commonly associated with Mexican culture",
			},
		},
	},
}
