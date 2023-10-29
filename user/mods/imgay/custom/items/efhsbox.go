package items

import "MT-GO/database"

var SmallBox = map[string]*database.CustomItemAPI{
	"efh_small_box": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5b7c710788a4506dec015957",
			HandbookPrice:    12000,
			ModifierType:     "clone",
			AddToTrader: map[string][]*database.CustomItemAddToTrader{
				"Therapist": {
					{
						LoyaltyLevel: 1,
						BarterScheme: map[string]float32{
							"5449016a4bdc2d6f028b456f": 12000,
						},
						AmountInStock: 5,
					},
				},
			},
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "[EFH] Small box",
				ShortName:   "S-Box",
				Description: "A small cardboard box, cant hold much, but very useful for sorting and organization\n\n MAX: 60kg \n Cannot hold large items \n Cannot hold most cases",
			},
			"ru": {
				Name:        "[EFH] Маленькая коробка",
				ShortName:   "М-Короб",
				Description: "Маленькая картонная коробка, не может вмещать много, но очень полезна для сортировки и организации\n\n МАКС: 60 кг \n Не может вмещать большие предметы \n Не может вмещать большинство контейнеров",
			},
		},
		Overrides: map[string]any{
			"Prefab": database.Prefab{
				Path: "assets/content/items/storage/box/box.bundle",
			},
			"Grids": []database.Grid{
				{
					Name:   "efh_small_box",
					ID:     "efh_small_box",
					Parent: "efh_small_box",
					Props: database.GridProps{
						Filters: []database.GridFilters{
							{
								Filter: []string{
									"5448eb774bdc2d0a728b4567",
									"543be5f84bdc2dd4348b456a",
									"5645bcb74bdc2ded0b8b4578",
									"5448fe124bdc2da5018b4567",
									"5448e53e4bdc2d60728b4567",
									"5448e5284bdc2dcb718b4567",
									"60b0f6c058e0b0481a09ad11",
									"5aafbde786f774389d0cbc0f",
									"619cbf9e0a7c3a1a2731940a",
									"619cbf7d23893217ec30b689",
									"59fafd4b86f7745ca07e1232",
									"62a09d3bcf4a99369e262447",
									"5d235bb686f77443f4331278",
									"5c093e3486f77430cb02e593",
									"590c60fc86f77412b13fddcf",
									"5783c43d2459774bbe137486",
									"5422acb9af1c889c16000029",
									"543be6674bdc2df1348b4569",
									"5448ecbe4bdc2d60728b4568",
									"543be5e94bdc2df1348b4568",
									"5447e1d04bdc2dff2f8b4567",
									"567849dd4bdc2d150f8b456e",
									"543be5664bdc2dd4348b4569",
									"5f4fbaaca5573a5ac31db429",
									"61605ddea09d851a0a0c1bbc",
									"616eb7aea207f41933308f46",
									"5991b51486f77447b112d44f",
									"5ac78a9b86f7741cca0bbd8d",
									"5b4391a586f7745321235ab2",
									"5af056f186f7746da511291f",
									"544fb5454bdc2df8738b456a",
									"5661632d4bdc2d903d8b456b",
									"543be6564bdc2df4348b4568",
								},
								ExcludedFilter: []string{
									"5b7c710788a4506dec015957",
									"5c0a840b86f7742ffa4f2482",
									"5b6d9ce188a4501afc1b2b25",
									"5aafbde786f774389d0cbc0f",
									"5c093db286f7740a1b2617e3",
									"59fb023c86f7746d0d4b423c",
									"5d1b36a186f7742523398433",
									"59fb042886f7746c5005a7b2",
									"5aafbcd986f7745e590fff23",
									"5e2af55f86f7746d4159f07c",
									"5c127c4486f7745625356c13",
									"efh_small_box",
									"efh_box_big",
								},
							},
						},
						CellsH:         7,
						CellsV:         7,
						MinCount:       0,
						MaxCount:       0,
						MaxWeight:      60,
						IsSortingTable: true,
					},
				},
			},
		},
	},
}

var BigBox = map[string]*database.CustomItemAPI{
	"efh_big_box": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5b7c710788a4506dec015957",
			HandbookPrice:    20000,
			ModifierType:     "clone",
			AddToTrader: map[string][]*database.CustomItemAddToTrader{
				"Therapist": {
					{
						LoyaltyLevel: 0,
						BarterScheme: map[string]float32{
							"5449016a4bdc2d6f028b456f": 20000,
						},
						AmountInStock: 5,
					},
					{
						LoyaltyLevel: 0,
						BarterScheme: map[string]float32{
							"efh_ration_card": 1,
						},
						AmountInStock: 1,
					},
				},
			},
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "[EFH] Big box",
				ShortName:   "B-Box",
				Description: "A Big cardboard box, cant hold much, but very useful for sorting and organization\n\n MAX: 60kg \n Cannot hold large items \n Cannot hold most cases",
			},
			"ru": {
				Name:        "[EFH] Маленькая коробка",
				ShortName:   "М-Короб",
				Description: "Маленькая картонная коробка, не может вмещать много, но очень полезна для сортировки и организации\n\n МАКС: 60 кг \n Не может вмещать большие предметы \n Не может вмещать большинство контейнеров",
			},
		},
		Overrides: map[string]any{
			"Prefab": database.Prefab{
				Path: "assets/content/items/storage/box/box.bundle",
			},
			"Grids": []database.Grid{
				{
					Name:   "efh_big_box",
					ID:     "efh_big_box",
					Parent: "efh_big_box",
					Props: database.GridProps{
						Filters: []database.GridFilters{
							{
								Filter: []string{
									"5448eb774bdc2d0a728b4567",
									"543be5f84bdc2dd4348b456a",
									"5645bcb74bdc2ded0b8b4578",
									"5448fe124bdc2da5018b4567",
									"5448e53e4bdc2d60728b4567",
									"5448e5284bdc2dcb718b4567",
									"60b0f6c058e0b0481a09ad11",
									"5aafbde786f774389d0cbc0f",
									"619cbf9e0a7c3a1a2731940a",
									"619cbf7d23893217ec30b689",
									"59fafd4b86f7745ca07e1232",
									"62a09d3bcf4a99369e262447",
									"5d235bb686f77443f4331278",
									"5c093e3486f77430cb02e593",
									"590c60fc86f77412b13fddcf",
									"5783c43d2459774bbe137486",
									"5422acb9af1c889c16000029",
									"543be6674bdc2df1348b4569",
									"5448ecbe4bdc2d60728b4568",
									"543be5e94bdc2df1348b4568",
									"5447e1d04bdc2dff2f8b4567",
									"567849dd4bdc2d150f8b456e",
									"543be5664bdc2dd4348b4569",
									"5f4fbaaca5573a5ac31db429",
									"61605ddea09d851a0a0c1bbc",
									"616eb7aea207f41933308f46",
									"5991b51486f77447b112d44f",
									"5ac78a9b86f7741cca0bbd8d",
									"5b4391a586f7745321235ab2",
									"5af056f186f7746da511291f",
									"544fb5454bdc2df8738b456a",
									"5661632d4bdc2d903d8b456b",
									"543be6564bdc2df4348b4568",
								},
								ExcludedFilter: []string{
									"5b7c710788a4506dec015957",
									"5c0a840b86f7742ffa4f2482",
									"5b6d9ce188a4501afc1b2b25",
									"5aafbde786f774389d0cbc0f",
									"5c093db286f7740a1b2617e3",
									"59fb023c86f7746d0d4b423c",
									"5d1b36a186f7742523398433",
									"59fb042886f7746c5005a7b2",
									"5aafbcd986f7745e590fff23",
									"5e2af55f86f7746d4159f07c",
									"5c127c4486f7745625356c13",
									"efh_small_box",
									"efh_box_big",
								},
							},
						},
						CellsH:         12,
						CellsV:         14,
						MinCount:       0,
						MaxCount:       0,
						MaxWeight:      100,
						IsSortingTable: true,
					},
				},
			},
		},
	},
}

var RationCard = map[string]*database.CustomItemAPI{
	"efh_ration_card": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "59faff1d86f7746c51718c9c",
			HandbookPrice:    20000,
			ModifierType:     "clone",
			AddToTrader: map[string][]*database.CustomItemAddToTrader{
				"Therapist": {
					{
						LoyaltyLevel: 0,
						BarterScheme: map[string]float32{
							"619cc01e0a7c3a1a2731940c": 1,
							"59e361e886f774176c10a2a5": 2,
							"5449016a4bdc2d6f028b456f": 1500,
						},
						AmountInStock: 5,
					},
					{
						LoyaltyLevel: 2,
						BarterScheme: map[string]float32{
							"59e361e886f774176c10a2a5": 1,
							"57347b8b24597737dd42e192": 1,
						},
						AmountInStock: 1,
					},
				},
			},
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "[EFH] Ration Card",
				ShortName:   "Ration Card",
				Description: "The UN Ration Card, issued by the United Nations Tarkov, is a vital lifeline for civilian's living in the war-torn Norvinsk region. This ration card represents a lifeline of hope and sustenance in a region plagued by scarcity and instability.\nEach UN Ration Card provides the bearer with a month's worth of essential rations, including food, clean water, and basic necessities. In a region where access to such resources has become a luxury, these ration cards are symbolic of a commitment to ensuring the survival of the most vulnerable population.\nHowever, as the economy in Norvinsk Region deteriorates and traditional currency loses its value, the ration card has taken on a new and unexpected role. Desperate times have led some individuals to use these ration cards as a form of alternative currency.\n In a place where trust and security are scarce, the UN Ration Card has become a valuable and sought-after commodity in the underground economy. \n This unlikely transformation of the ration card underscores the extraordinary challenges faced by the people of Norvinsk, as they adapt to a world where survival often hinges on resourcefulness, resilience, and the ability to find value in the unlikeliest of places. While originally intended as a means of providing sustenance, the UN Ration Card now serves as both a symbol of hope and a unit of exchange, embodying the resourcefulness and determination of those who hold it.",
			},
			"ru": {
				Name:        "[EFH] Продовольственная карта",
				ShortName:   "Продовольственная карта",
				Description: "Продовольственная карта ООН, выдаваемая Организацией Объединенных Наций в Таркове, представляет собой важную связь для жителей живущих в разрушенном войной регионе Норвинск. Эта продовольственная карта символизирует линию жизни надежды и поддержки в регионе, страдающем от нехватки и нестабильности.\nКаждая продовольственная карта ООН предоставляет ее обладателю месяц запаса жизненно важных продуктов, включая пищу, чистую воду и необходимые товары. В регионе, где доступ к таким ресурсам стал редкостью, эти продовольственные карты символизируют обязательство обеспечить выживание наиболее уязвимого населения.\nОднако по мере ухудшения экономики в регионе Норвинск и потери традиционной валюты, продовольственная карта приобрела новую и неожиданную роль. Отчаянные времена побудили некоторых людей использовать эти продовольственные карты в качестве альтернативной валюты.\n В месте, где доверие и безопасность редкость, продовольственная карта ООН стала ценным и востребованным товаром в подпольной экономике. \n Это неожиданное изменение назначения продовольственной карты подчеркивает невероятные вызовы, с которыми сталкиваются жители Норвинска, адаптируясь к миру, где выживание часто зависит от находчивости, устойчивости и способности находить ценность в самых неожиданных местах. В то время как изначально она предназначалась для обеспечения пропитания, продовольственная карта ООН теперь служит и символом надежды, и средством обмена, воплощая находчивость и решимость тех, кто ее держит.",
			},
		},
		Overrides: map[string]any{
			"Prefab": database.Prefab{
				Path: "assets/content/items/barter/ration-card/item_barter_valuable_bitcoin.bundle",
			},
		},
	},
}

var HQImprovisedArmorRig = map[string]*database.CustomItemAPI{
	"efh_improvised_armor_high_quality": {
		API: "item",
		Parameters: database.CustomItemParams{
			ReferenceItemTPL: "5d5d646386f7742797261fd9",
			HandbookPrice:    55420,
			ModifierType:     "clone",
			AddToTrader: map[string][]*database.CustomItemAddToTrader{
				"Ragman": {
					{
						LoyaltyLevel: 0,
						BarterScheme: map[string]float32{
							"5fd4c4fa16cac650092f6771": 1,
							"591094e086f7747caa7bb2ef": 1,
							"61bf83814088ec1a363d7097": 1,
							"59e35de086f7741778269d84": 1,
							"5449016a4bdc2d6f028b456f": 9000,
						},
						AmountInStock: 999,
					},
					{
						LoyaltyLevel: 3,
						BarterScheme: map[string]float32{
							"5449016a4bdc2d6f028b456f": 67000,
						},
						AmountInStock: 2,
					},
				},
			},
		},
		Locale: map[string]*database.CustomItemLocale{
			"en": {
				Name:        "[EFH] Improvised Armored Carrier",
				ShortName:   "IAC",
				Description: "This Improvised Armored Carrier is a testament to craftsmanship and innovation in a world defined by scarcity. Its foundation is built upon sturdy cargo tarp, repurposed from the shelves of \"IDEA,\". This trap, known for its flexibility and durability, forms the resilient core of the carrier to hold your magazines, and whatever else, you could need to cary on the battlefield.\nFor added protection, this variant features Level 5A armored plates salvaged from somewhere within the Tarkov region. Meticulously integrated into the structure, these plates are bound together through a combination of expert sewing, precise riveting, and cold-welding techniques. The result is a meticulously crafted carrier that offers dependable shielding and versatility, capable of withstanding even some of the strongest bullets.\nAs you finish examining the object, you ponder who made was the master improvisor who put this thing together.",
			},
			"ru": {
				Name:        "[EFH] Самодельный Бронированный Несущий",
				ShortName:   "СБН",
				Description: "Этот Самодельный Бронированный Несущий - свидетельство искусства и инноваций в мире, определенном нехваткой. Его основа построена на прочном грузовом тенте, переделанном с полок магазина \"IDEA\". Этот тент, известный своей гибкостью и прочностью, формирует устойчивое ядро несущего, предназначенного для хранения ваших магазинов и всего, что вам может понадобиться на поле боя.\nДля дополнительной защиты эта модификация включает в себя броневые пластины уровня 5А, извлеченные откуда-то в регионе Таркова. Эти пластины тщательно интегрированы в структуру с помощью опытного шитья, точного заклепывания и методов холодной сварки. Результатом является метикулезно сделанный несущий, обеспечивающий надежную защиту и универсальность, способный выдерживать даже некоторые из самых сильных пуль.\nЗавершая осмотр объекта, вы размышляете о том, кто был мастером-импровизатором, создавшим это удивительное изделие.",
			},
		},
		Overrides: map[string]any{
			"MaxDurability": 65,
			"Weight":        13.1,
			"armorClass":    5,
			"armorZone": []string{
				"Chest",
				"Wallet",
				"Stomach",
			},
			"Prefab": database.Prefab{
				Path: "assets/content/items/barter/ration-card/item_barter_valuable_bitcoin.bundle",
			},
		},
	},
}
