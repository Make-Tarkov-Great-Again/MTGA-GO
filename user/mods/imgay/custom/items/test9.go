package items

import "MT-GO/database"

var Test9 = map[string]*database.CustomItemAPI{
	"test9": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5b432d215acfc4771e1c6624",
			HandbookPrice:    50000,
			ModifierType:     "clone",
			AddToTrader: map[string][]*database.CustomItemAddToTrader{
				"Ragman": {
					{
						LoyaltyLevel: 0,
						BarterScheme: map[string]float32{
							"5449016a4bdc2d6f028b456f": 50000,
						},
						AmountInStock: 999,
					},
				},
			},
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"path": "test9.bundle",
				"rcid": "",
			},
			"Slots": []*database.Slot{
				{
					Name:   "mod_nvg",
					ID:     "sombrero_nvg",
					Parent: "test9",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"5ea058e01dbce517f324b3e2",
									"5c0558060db834001b735271",
									"5a16b8a9fcdbcb00165aa6ca",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
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
