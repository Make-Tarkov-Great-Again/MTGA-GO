package items

import "MT-GO/database"

var Famas = map[string]*database.CustomItemAPI{
	"weapon_famas_556x45": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5447a9cd4bdc2dbd208b4567",
			HandbookPrice:    34230,
			ModifierType:     "clone",
			AddToTrader: map[string]*database.CustomItemAddToTrader{
				"Peacekeeper": {
					LoyaltyLevel: 0,
					BarterScheme: map[string]float32{
						"5449016a4bdc2d6f028b456f": 99999,
					},
					AmountInStock: 0,
				},
			},
			ItemPresets: map[string]map[string]any{},
		},
		Overrides: map[string]any{
			"Slots": []map[string]any{
				{
					"_name":   "mod_muzzle",
					"_id":     "weapon_famas_556x45_barrel",
					"_parent": "weapon_famas_556x45",
					"_props": map[string]any{
						"filters": []map[string]any{
							{
								"Shift": 0,
								"Filter": []string{
									"barrel_famas_488mm",
								},
							},
						},
					},
					"_required":              true,
					"_mergeSlotWithChildren": false,
					"_proto":                 "55d30c4c4bdc2db4468b457e",
				},
			},

			"CameraRecoil":   0.12,
			"CameraSnap":     3.5,
			"ReloadMode":     "ExternalMagazine",
			"CenterOfImpact": 0.01,
			"AimPlane":       0.16,
			"DeviationCurve": 1.35,
			"DeviationMax":   23,
		},
	},
}
