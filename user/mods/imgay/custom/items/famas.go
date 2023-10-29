package items

import "MT-GO/database"

var FAMAS = map[string]*database.CustomItemAPI{
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
						"5449016a4bdc2d6f028b456f": 34230,
					},
					AmountInStock: 0,
				},
			},
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "GIAT FAMAS 5.56x45 Assault Rifle",
				ShortName:   "FAMAS",
				Description: "The FAMAS (Fusil d'Assaut de la Manufacture d'Armes de Saint-Étienne) is a bullpup assault rifle designed and manufactured in France by MAS in 1978. The FAMAS is recognised for its high rate of fire at 1,100 rounds per minute.",
			},
		},
		Overrides: map[string]any{
			"Slots": []*database.Slot{
				{
					Name:   "mod_muzzle",
					ID:     "weapon_famas_556x45_barrel",
					Parent: "weapon_famas_556x45",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"barrel_famas_488mm",
								},
							},
						},
					},
					Required:              true,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
				{
					Name:   "mod_magazine",
					ID:     "weapon_famas_556x45_magazine",
					Parent: "weapon_famas_556x45",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"mag_famas_556x45_25",
									"mag_famas_556x45_30",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
				{
					Name:   "mod_scope",
					ID:     "weapon_famas_556x45_scope",
					Parent: "weapon_famas_556x45",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"upper_famas",
								},
							},
						},
					},
					Required:              true,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
				{
					Name:   "mod_handguard",
					ID:     "weapon_famas_556x45_handguard",
					Parent: "weapon_famas_556x45",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"mount_famas_bottom",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
				{
					Name:   "mod_tactical_002",
					ID:     "weapon_famas_556x45_tac02",
					Parent: "weapon_famas_556x45",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"mount_famas_side",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
			},
			"CameraRecoil":   0.12,
			"CameraSnap":     3.5,
			"ReloadMode":     "ExternalMagazine",
			"CenterOfImpact": 0.01,
			"AimPlane":       0.16,
			"DeviationCurve": 1.35,
			"DeviationMax":   23,
			"weapFireType": []string{
				"single",
				"burst",
				"fullauto",
			},
			"defMagType": "mag_famas_556x45_30",
			"Chambers": []database.Slot{
				{
					Name:   "patron_in_weapon",
					ID:     "weapon_famas_556x45_chamber",
					Parent: "weapon_famas_556x45",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Filter: []string{
									"59e6920f86f77411d82aa167",
									"59e6927d86f77411da468256",
									"54527a984bdc2d4e668b4567",
									"54527ac44bdc2d36668b4567",
									"59e68f6f86f7746c9f75e846",
									"59e6906286f7746c9f75e847",
									"59e690b686f7746c9f75e848",
									"59e6918f86f7746c9f75e849",
									"60194943740c5d77f6705eea",
									"601949593ae8f707c4608daa",
									"5c0d5ae286f7741e46554302",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d4af244bdc2d962f8b4571",
				},
			},
		},
		ItemPresets: map[string]*database.CustomItemPreset{
			"famas_std_presetid": {
				Id:               "famas_rails_presetid",
				Type:             "Preset",
				ChangeWeaponName: false,
				Name:             "FAMAS Default",
				Encyclopedia:     "weapon_famas_556x45",
				Parent:           "famas_std_id",
				Items: []*database.InventoryItem{
					{
						ID:  "famas_std_id",
						TPL: "weapon_famas_556x45",
					},
					{
						ID:       "famas_std_barrelID",
						TPL:      "barrel_famas_488mm",
						ParentID: "famas_std_id",
						SlotID:   "mod_muzzle",
					},
					{
						ID:       "famas_std_muzzleID",
						TPL:      "muzzle_famas_flash_hider",
						ParentID: "famas_std_barrelID",
						SlotID:   "mod_muzzle",
					},
					{
						ID:       "famas_std_upperID",
						TPL:      "upper_famas",
						ParentID: "famas_std_id",
						SlotID:   "mod_scope",
					},
					{
						ID:       "famas_std_magazineID",
						TPL:      "mag_famas_556x45_25",
						ParentID: "famas_std_id",
						SlotID:   "mod_magazine",
					},
				},
			},
			"famas_rails_presetid": {
				Id:               "famas_rails_presetid",
				Type:             "Preset",
				ChangeWeaponName: false,
				Name:             "FAMAS rails",
				Encyclopedia:     "weapon_famas_556x45",
				Parent:           "famas_rails_id",
				Items: []*database.InventoryItem{
					{
						ID:  "famas_rails_id",
						TPL: "weapon_famas_556x45",
					},
					{
						ID:       "famas_rails_barrelID",
						TPL:      "barrel_famas_488mm",
						ParentID: "famas_rails_id",
						SlotID:   "mod_muzzle",
					},
					{
						ID:       "famas_rails_muzzleID",
						TPL:      "muzzle_famas_flash_hider",
						ParentID: "famas_rails_barrelID",
						SlotID:   "mod_muzzle",
					},
					{
						ID:       "famas_rails_upperID",
						TPL:      "upper_famas",
						ParentID: "famas_rails_id",
						SlotID:   "mod_scope",
					},
					{
						ID:       "famas_rails_toprailID",
						TPL:      "mount_famas_optic_rail",
						ParentID: "famas_rails_upperID",
						SlotID:   "mod_mount",
					},
					{
						ID:       "famas_rails_siderailID",
						TPL:      "mount_famas_side",
						ParentID: "famas_rails_id",
						SlotID:   "mod_tactical_002",
					},
					{
						ID:       "famas_rails_bottomrailID",
						TPL:      "mount_famas_bottom",
						ParentID: "famas_rails_id",
						SlotID:   "mod_handguard",
					},
					{
						ID:       "famas_rails_magazineID",
						TPL:      "mag_famas_556x45_30",
						ParentID: "famas_rails_id",
						SlotID:   "mod_magazine",
					},
				},
			},
		},
	},
}
