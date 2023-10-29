package items

import "MT-GO/database"

var smallBox = map[string]*database.CustomItemAPI{
	"efh_small_box": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5b7c710788a4506dec015957",
			HandbookPrice:    12000,
			ModifierType:     "clone",
			AddToTrader: map[string]*database.CustomItemAddToTrader{
				"Therapist": {
					LoyaltyLevel: 0,
					BarterScheme: map[string]float32{
						"5449016a4bdc2d6f028b456f": 12000,
					},
					AmountInStock: 2,
				},
			},
		},
		ItemPresets: map[string]*database.CustomItemPreset{},
		Overrides: map[string]any{
			"Grids": map[string]any{},
		},
	},
}
