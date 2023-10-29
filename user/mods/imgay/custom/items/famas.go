package items

import "MT-GO/database"

var FAMAS = map[string]*database.CustomItemAPI{
	"weapon_famas_556x45": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5447a9cd4bdc2dbd208b4567",
			HandbookPrice:    34230,
			ModifierType:     "clone",
			AddToTrader: map[string][]*database.CustomItemAddToTrader{
				"5935c25fb3acc3127c3d8cd9": {
					{
						LoyaltyLevel: 1,
						Set: map[string]any{
							"mod_muzzle": map[string]any{
								"_tpl": "barrel_famas_488mm",
								"attachments": map[string]any{
									"mod_muzzle": map[string]any{
										"_tpl": "muzzle_famas_flash_hider",
									},
								},
							},
							"mod_scope": map[string]any{
								"_tpl": "upper_famas",
							},
							"mod_magazine": map[string]any{
								"_tpl": "mag_famas_556x45_25",
							},
						},
						BarterScheme: map[string]float32{
							"5696686a4bdc2da3298b456a": 784,
						},
						AmountInStock: 999,
					},
					{
						LoyaltyLevel: 1,
						Set: map[string]any{
							"mod_muzzle": map[string]any{
								"_tpl": "barrel_famas_488mm",
								"attachments": map[string]any{
									"mod_muzzle": map[string]any{
										"_tpl": "muzzle_famas_flash_hider",
									},
								},
							},
							"mod_scope": map[string]any{
								"_tpl": "upper_famas",
								"attachments": map[string]any{
									"mod_mount": map[string]any{
										"_tpl": "mount_famas_optic_rail",
									},
								},
							},
							"mod_tactical_002": map[string]any{
								"_tpl": "mount_famas_side",
							},
							"mod_handguard": map[string]any{
								"_tpl": "mount_famas_bottom",
							},
							"mod_magazine": map[string]any{
								"_tpl": "mag_famas_556x45_30",
							},
						},
						BarterScheme: map[string]float32{
							"5696686a4bdc2da3298b456a": 892,
						},
						AmountInStock: 999,
					},
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
			"Prefab": database.Prefab{
				Path: "assets/content/weapons/famas/weapon_mas_famas_556x45_container.bundle",
				Rcid: "",
			},
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
			"CenterOfImpact": 0.01,
			"AimPlane":       0.16,
			"DeviationCurve": 1.35,
			"DeviationMax":   23,
			"weapFireType": []string{
				"single",
				"burst",
				"fullauto",
			},
			"defMagType":                  "mag_famas_556x45_30",
			"isFastReload":                true,
			"RecoilForceUp":               90,
			"RecoilForceBack":             350,
			"Convergence":                 1.5,
			"RecoilAngle":                 90,
			"RecolDispersion":             25,
			"SingleFireRate":              450,
			"CanQueueSecondShot":          true,
			"bFirerate":                   1100,
			"Ergonomics":                  76,
			"Velocity":                    0,
			"bEffDist":                    300,
			"bHearDist":                   80,
			"HipAccuracyRestorationDelay": 0.2,
			"HipAccuracyRestorationSpeed": 7,
			"HipInnaccuracyGain":          0.16,
			"AimSensitivity":              0.65,
			"BurstShotsCount":             3,
			"BaseMalfunctionChance":       0.1595,
			"DurabilityBurnRatio":         1.15,
			"HeatFactorGun":               1,
			"CoolFactorGun":               3.168,
			"CoolFactorGunMods":           1,
			"HeatFactorByShot":            1.235,
			"DoubleActionAccuracyPenalty": 1.5,
			"RecoilPosZMult":              1,
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
	"upper_famas": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "55d355e64bdc2d962f8b4569",
			HandbookPrice:    22640,
			ModifierType:     "clone",
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "Upper part for GIAT FAMAS",
				ShortName:   "FAMAS Upper",
				Description: "Standard upper part manufactured by GIAT for FAMAS assault rifle 5.56x45.",
			},
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"path": "assets/content/items/mods/famas/upper_famas.bundle",
				"rcid": "",
			},
			"Slots": []map[string]any{
				{
					"_name":   "mod_mount",
					"_id":     "upper_famas_mount",
					"_parent": "upper_famas",
					"_props": map[string]any{
						"filters": []map[string]any{
							{
								"Shift": 0,
								"Filter": []string{
									"mount_famas_optic_rail",
								},
							},
						},
					},
					"_required":              false,
					"_mergeSlotWithChildren": false,
					"_proto":                 "55d30c4c4bdc2db4468b457e",
				},
			},
		},
	},
	"mag_famas_556x45_25": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "55d4887d4bdc2d962f8b4570",
			HandbookPrice:    3120,
			ModifierType:     "clone",
			AddToTrader: map[string][]*database.CustomItemAddToTrader{
				"5935c25fb3acc3127c3d8cd9": {
					{
						LoyaltyLevel: 0,
						BarterScheme: map[string]float32{
							"5696686a4bdc2da3298b456a": 26,
						},
						AmountInStock: 0,
					},
				},
			},
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"Path": "assets/content/items/mods/famas/mag_famas_556x45_25.bundle",
				"Rcid": "",
			},
			"Cartridges": []database.Cartridges{
				{
					Name:     "cartridges",
					Id:       "mag_famas_556x45_25_cartridges",
					Parent:   "mag_famas_556x45_25",
					MaxCount: 25,
					Proto:    "5748538b2459770af276a261",
					Props: database.CartridgeFilters{
						Filters: []database.CartridgeFilter{
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
									"5fbe3ffdf8b6a877a729ea82",
									"5fd20ff893a8961fc660a954",
									"619636be6db0f2477964e710",
									"6196364158ef8c428c287d9f",
									"6196365d58ef8c428c287da1",
								},
							},
						},
					},
				},
			},
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "GIAT FAMAS 5.56x45 25-round box magazine",
				ShortName:   "FAMAS 25",
				Description: "25-round standard GIAT FAMAS metal magazine for 5.56x45 NATO cartridges.",
			},
		},
	},
	"mag_famas_556x45_30": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "55d4887d4bdc2d962f8b4570",
			HandbookPrice:    4210,
			ModifierType:     "clone",
			AddToTrader: map[string][]*database.CustomItemAddToTrader{
				"5935c25fb3acc3127c3d8cd9": {
					{
						LoyaltyLevel: 0,
						BarterScheme: map[string]float32{
							"5696686a4bdc2da3298b456a": 39,
						},
						AmountInStock: 0,
					},
				},
			},
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"Path": "assets/content/items/mods/famas/mag_famas_556x45_30.bundle",
				"Rcid": "",
			},
			"Cartridges": []database.Cartridges{
				{
					Name:     "cartridges",
					Id:       "mag_famas_556x45_30_cartridges",
					Parent:   "mag_famas_556x45_30",
					MaxCount: 30,
					Proto:    "5748538b2459770af276a261",
					Props: database.CartridgeFilters{
						Filters: []database.CartridgeFilter{
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
									"5fbe3ffdf8b6a877a729ea82",
									"5fd20ff893a8961fc660a954",
									"619636be6db0f2477964e710",
									"6196364158ef8c428c287d9f",
									"6196365d58ef8c428c287da1",
								},
							},
						},
					},
				},
			},
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "GIAT FAMAS 5.56x45 30-round extended magazine",
				ShortName:   "FAMAS 30",
				Description: "30-round extended GIAT FAMAS metal magazine for 5.56x45 NATO cartridges.",
			},
		},
	},
	"barrel_famas_488mm": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "55d3632e4bdc2d972f8b4569",
			HandbookPrice:    26460,
			ModifierType:     "clone",
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"path": "assets/content/items/mods/famas/barrel_famas_488mm.bundle",
				"rcid": "",
			},
			"Slots": []database.Slot{
				{
					Name:                  "mod_muzzle",
					ID:                    "barrel_famas_488mm_muzzle",
					Parent:                "barrel_famas_488mm",
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"5b3a16655acfc40016387a2a",
									"5c7e5f112e221600106f4ede",
									"5c0fafb6d174af02a96260ba",
									"5cf6937cd7f00c056c53fb39",
									"544a38634bdc2d58388b4568",
									"5cff9e5ed7ad1a09407397d4",
									"5c48a2a42e221602b66d1e07",
									"5f6372e2865db925d54f3869",
									"615d8e2f1cb55961fa0fd9a4",
									"56ea8180d2720bf2698b456a",
									"5d02676dd7ad1a049e54f6dc",
									"56ea6fafd2720b844b8b4593",
									"5943ee5a86f77413872d25ec",
									"609269c3b0e443224b421cc1",
									"5c7fb51d2e2216001219ce11",
									"5ea172e498dacb342978818e",
									"5c6d710d2e22165df16b81e7",
									"612e0e55a112697a4b3a66e7",
									"5d440625a4b9361eec4ae6c5",
									"5cc9b815d7f00c000e2579d6",
									"5a7c147ce899ef00150bd8b8",
									"5c7954d52e221600106f4cc7",
									"59bffc1f86f77435b128b872",
									"5a9fbb84a2750c00137fa685",
									"muzzle_famas_flash_hider",
								},
							},
						},
					},
				},
			},
			"Accuracy":                  0,
			"Recoil":                    -7,
			"Loudness":                  0,
			"EffectiveDistance":         0,
			"Ergonomics":                -6,
			"Velocity":                  3,
			"CenterOfImpact":            0.032,
			"ShotgunDispersion":         1,
			"DurabilityBurnModificator": 1.1,
			"HeatFactor":                0.94,
			"CoolFactor":                0.94,
			"DeviationCurve":            1.35,
			"DeviationMax":              23,
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "488mm barrel for FAMAS 5.56x45",
				ShortName:   "488mm FAMAS 5.56x45",
				Description: "A barrel for GIAT FAMAS 5.56x45 ammo, 488mm long.",
			},
		},
	},
	"muzzle_famas_flash_hider": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "544a38634bdc2d58388b4568",
			ModifierType:     "clone",
			HandbookPrice:    1340,
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"path": "assets/content/items/mods/famas/muzzle_famas_flash_hider.bundle",
				"rcid": "",
			},
			"Accuracy":          0,
			"Recoil":            -7,
			"Loudness":          0,
			"EffectiveDistance": 0,
			"Ergonomics":        -1,
			"Velocity":          0.5,
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "GIAT 5.56x45 flash hider for FAMAS",
				ShortName:   "FAMAS FH",
				Description: "Standard issue flash hider designed for installation on GIAT FAMAS chambered in 5.56x45mm.",
			},
		},
	},
	"mount_famas_optic_rail": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5bfebc530db834001d23eb65",
			ModifierType:     "clone",
			HandbookPrice:    5640,
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"path": "assets/content/items/mods/famas/mount_famas_opticrail.bundle",
				"rcid": "",
			},
			"ConflictingItems": []string{
				"5947db3f86f77447880cf76f",
				"57486e672459770abd687134",
				"576fd4ec2459777f0b518431",
				"5a7c74b3e899ef0014332c29",
				"591ee00d86f774592f7b841e",
				"57acb6222459771ec34b5cb0",
			},
			"Slots": []*database.Slot{
				{
					Name:   "mod_scope",
					ID:     "mount_famas_optic_rail_scope",
					Parent: "mount_famas_optic_rail",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"57ac965c24597706be5f975c",
									"57aca93d2459771f2c7e26db",
									"544a3a774bdc2d3a388b4567",
									"5d2dc3e548f035404a1a4798",
									"57adff4f24597737f373b6e6",
									"5c0517910db83400232ffee5",
									"591c4efa86f7741030027726",
									"570fd79bd2720bc7458b4583",
									"570fd6c2d2720bc6458b457f",
									"558022b54bdc2dac148b458d",
									"5c07dd120db834001c39092d",
									"5c0a2cec0db834001b7ce47d",
									"58491f3324597764bc48fa02",
									"584924ec24597768f12ae244",
									"5b30b0dc5acfc400153b7124",
									"6165ac8c290d254f5e6b2f6c",
									"60a23797a37c940de7062d02",
									"5d2da1e948f035477b1ce2ba",
									"5c0505e00db834001b735073",
									"609a63b6e2ff132951242d09",
									"584984812459776a704a82a6",
									"59f9d81586f7744c7506ee62",
									"570fd721d2720bc5458b4596",
									"57ae0171245977343c27bfcf",
									"5dfe6104585a0c3e995c7b82",
									"544a3d0a4bdc2d1b388b4567",
									"5d1b5e94d7ad1a2b865a96b0",
									"609bab8b455afd752b2e6138",
									"58d39d3d86f77445bb794ae7",
									"616554fe50224f204c1da2aa",
									"5c7d55f52e221644f31bff6a",
									"616584766ef05c2ce828ef57",
									"5b3b6dc75acfc47a8773fb1e",
									"615d8d878004cc50514c3233",
									"5b2389515acfc4771e1be0c0",
									"577d128124597739d65d0e56",
									"618b9643526131765025ab35",
									"618bab21526131765025ab3f",
									"5c86592b2e2216000e69e77c",
									"5a37ca54c4a282000d72296a",
									"5d0a29fed7ad1a002769ad08",
									"5c064c400db834001d23f468",
									"58d2664f86f7747fec5834f6",
									"57c69dd424597774c03b7bbc",
									"5b3b99265acfc4704b4a1afb",
									"5aa66a9be5b5b0214e506e89",
									"5aa66c72e5b5b00016327c93",
									"5c1cdd302e221602b3137250",
									"61714b2467085e45ef140b2c",
									"6171407e50224f204c1da3c5",
									"61713cc4d8e3106d9806c109",
									"5b31163c5acfc400153b71cb",
									"5a33b652c4a28232996e407c",
									"5a33b2c9c4a282000c5a9511",
									"59db7eed86f77461f8380365",
									"5a1ead28fcdbcb001912fa9f",
									"5dff77c759400025ea5150cf",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
				{
					Name:   "mod_sight_rear",
					ID:     "mount_famas_optic_rearsight",
					Parent: "mount_famas_optic",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"5ba26b17d4351e00367f9bdd",
									"5dfa3d7ac41b2312ea33362a",
									"5c1780312e221602b66cc189",
									"5fb6564947ce63734e3fa1da",
									"5bc09a18d4351e003562b68e",
									"5c18b9192e2216398b5a8104",
									"5fc0fa957283c4046c58147e",
									"5894a81786f77427140b8347",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
				{
					Name:   "mod_sight_front",
					ID:     "mount_famas_optic_rail_frontsight",
					Parent: "mount_famas_optic_rail",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"5ba26b01d4351e0085325a51",
									"5dfa3d950dee1b22f862eae0",
									"5c17804b2e2216152006c02f",
									"5fb6567747ce63734e3fa1dc",
									"5bc09a30d4351e00367fb7c8",
									"5c18b90d2e2216152142466b",
									"5fc0fa362770a0045c59c677",
									"5894a73486f77426d259076c",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
				{
					Name:   "mod_tactical_000",
					ID:     "mount_famas_optic_rail_tac00",
					Parent: "mount_famas_optic_rail",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"5649a2464bdc2d91118b45a8",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
			},
			"Weight":            0.15,
			"Width":             2,
			"Height":            1,
			"Accuracy":          0,
			"Recoil":            0,
			"Loudness":          0,
			"EffectiveDistance": 0,
			"Ergonomics":        0,
			"Velocity":          0,
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "GIAT top optic rail for FAMAS",
				ShortName:   "FAMAS top",
				Description: "Custom top mount for installation of optical scopes, collimator sights and other devices and accessories on FAMAS.",
			},
		},
	},
	"mount_famas_side": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5a9d6d00a2750c5c985b5305",
			ModifierType:     "clone",
			HandbookPrice:    2033,
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"path": "assets/content/items/mods/famas/mount_famas_side.bundle",
				"rcid": "",
			},
			"Slots": []*database.Slot{
				{
					Name:   "mod_tactical",
					ID:     "mount_famas_side_tac",
					Parent: "mount_famas_side",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"5a800961159bd4315e3a1657",
									"57fd23e32459772d0805bcf1",
									"544909bb4bdc2d6f028b4577",
									"5c06595c0db834001a66af6c",
									"5cc9c20cd7f00c001336c65d",
									"5d2369418abbc306c62e0c80",
									"5b07dd285acfc4001754240d",
									"56def37dd2720bec348b456a",
									"5a7b483fe899ef0016170d15",
									"61605d88ffa6e502ac5e7eeb",
									"5a5f1ce64f39f90b401987bc",
									"560d657b4bdc2da74d8b4572",
									"5b3a337e5acfc4704b4a19a0",
									"5c5952732e2216398b5abda2",
									"57d17e212459775a1179a0f5",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
			},
			"Accuracy":          0,
			"Recoil":            0,
			"Loudness":          0,
			"EffectiveDistance": 0,
			"Ergonomics":        0,
			"Velocity":          0,
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "GIAT FAMAS side rail",
				ShortName:   "FAMAS side",
				Description: "Short side rail allows you to install additional equipment on the side of a FAMAS assault rifle.",
			},
		},
	},
	"mount_famas_bottom": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5a9d6d00a2750c5c985b5305",
			ModifierType:     "clone",
			HandbookPrice:    3340,
		},
		Overrides: map[string]any{
			"Prefab": map[string]any{
				"path": "assets/content/items/mods/famas/mount_famas_bottom.bundle",
				"rcid": "",
			},
			"Slots": []*database.Slot{
				{
					Name:   "mod_foregrip",
					ID:     "mount_famas_bottom_fg",
					Parent: "mount_famas_bottom",
					Props: database.SlotProps{
						Filters: []database.SlotFilters{
							{
								Shift: 0,
								Filter: []string{
									"5c7fc87d2e221644f31c0298",
									"5cda9bcfd7f00c0c0b53e900",
									"59f8a37386f7747af3328f06",
									"619386379fb0c665d5490dbe",
									"5c87ca002e221600114cb150",
									"588226d124597767ad33f787",
									"588226dd24597767ad33f789",
									"588226e62459776e3e094af7",
									"588226ef24597767af46e39c",
									"59fc48e086f77463b1118392",
									"5fce0cf655375d18a253eff0",
									"5cf4fb76d7f00c065703d3ac",
									"5b057b4f5acfc4771e1bd3e9",
									"5c791e872e2216001219c40a",
									"558032614bdc2de7118b4585",
									"58c157be86f77403c74b2bb6",
									"58c157c886f774032749fb06",
									"5f6340d3ca442212f4047eb2",
									"591af28e86f77414a27a9e1d",
									"5c1cd46f2e22164bef5cfedb",
									"5c1bc4812e22164bef5cfde7",
									"5c1bc5612e221602b5429350",
									"5c1bc5af2e221602b412949b",
									"5c1bc5fb2e221602b1779b32",
									"5c1bc7432e221602b412949d",
									"5c1bc7752e221602b1779b34",
								},
							},
						},
					},
					Required:              false,
					MergeSlotWithChildren: false,
					Proto:                 "55d30c4c4bdc2db4468b457e",
				},
			},
			"Accuracy":          0,
			"Recoil":            0,
			"Loudness":          0,
			"EffectiveDistance": 0,
			"Ergonomics":        0,
			"Velocity":          0,
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "GIAT FAMAS bottom rail",
				ShortName:   "FAMAS bottom",
				Description: "Bottom rail allows you to install different foregrips on a FAMAS assault rifle.",
			},
		},
	},
}
