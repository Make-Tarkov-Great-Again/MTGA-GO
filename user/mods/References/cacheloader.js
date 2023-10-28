exports.mod = (mod_data) => {

    // Getting data from mod.config.json
    let ModFolderName = `${mod_data.author}-${mod_data.name}-${mod_data.version}`;
    let ModFileNames = mod_data.filenames;
    let PathResolver = global.internal.path.resolve;

    // Checking for no FileNames
    if (ModFileNames.length == 0) {
        logger.logWarning(`[WARNING] ${mod_data.name} has no filenames in mod.config.json`)
        return
    }

    // Loading our cachefiles
    const cacheLoad = function (filepath) { return global.fileIO.readParsed(PathResolver(filepath)) }
    let items = cacheLoad('user/cache/items.json');
    let assort_ragfair = cacheLoad('user/cache/assort_ragfair.json');
    let locale_en = cacheLoad('user/cache/locale_en.json');
    let locale_fr = cacheLoad(`user/mods/${ModFolderName}/src/Languages/Original/fr/locale.json`);
    let locale_ge = cacheLoad(`user/mods/${ModFolderName}/src/Languages/Original/ge/locale.json`);
    let locale_ru = cacheLoad('user/cache/locale_ru.json');
    let templates = cacheLoad('user/cache/templates.json');
    let assortPrapor = cacheLoad('user/cache/assort_54cb50c76803fa8b248b4571.json');
    let assortTherapist = cacheLoad('user/cache/assort_54cb57776803fa99248b456e.json');
    let assortFence = cacheLoad('user/cache/assort_579dc571d53a0658a154fbec.json');
    let assortSkier = cacheLoad('user/cache/assort_58330581ace78e27b8b10cee.json');
    let assortPeacekeeper = cacheLoad('user/cache/assort_5935c25fb3acc3127c3d8cd9.json');
    let assortMechanic = cacheLoad('user/cache/assort_5a7c2eca46aef81a7ca2145d.json');
    let assortRagman = cacheLoad('user/cache/assort_5ac3b934156ae10c4430e83c.json');
    let assortJaeger = cacheLoad('user/cache/assort_5c0647fdd443bc2504c2d371.json');
    let suitsGlobal = cacheLoad('user/cache/customization.json');
    let suitsRagman = cacheLoad('user/cache/customization_5ac3b934156ae10c4430e83c.json');

    // Trader function
    function pushTrader(assort, curAddTrader, curItem, curTrader, traderName, modifierFlag) {
        assort.data.barter_scheme[curItem] = curAddTrader.barter_scheme;
        assort.data.loyal_level_items[curItem] = curAddTrader.loyality;
        switch (modifierFlag) {
            case "edit":
                let foundInAssort = false;
                for (let curAssortItem in assort.data.items) {
                    if (curAssortItem.hasOwnProperty(curItem)) {
                        foundInAssort = true;
                        break;
                    }
                }
                if (!foundInAssort) {
                    assort.data.items.push(curTrader);
                }
                logger.logInfo(`[ADDTOTRADER] Added or changed ${curItem} in ${traderName}'s assort.`);
                break;
            default:
                assort.data.items.push(curTrader);
                logger.logInfo(`[ADDTOTRADER] ${curItem} was added to ${traderName}'s assort.`);
        }
        return assort;
    }

    // Assemble top/bottom part of suit function
    function topBottomAssemble(topBottomID, topBottomPath, handsPath, loyalty, level, price, name) {

        // Initialising a new variable to assemble locales
        let newLocale = {
            "Description": "",
            "Name": name,
            "ShortName": ""
        }

        // Checking if item has hands path or it's a bottom part
        switch (handsPath) {
            case false:
                // Item doesn't have hands path - thus it's bottom
                // Initialising a new variable to assemble the bottom
                let newBottom = {
                    "_id": topBottomID,
                    "_name": topBottomID,
                    "_parent": "5cc0869814c02e000a4cad94",
                    "_type": "Item",
                    "_props": {
                        "Name": topBottomID,
                        "ShortName": topBottomID,
                        "Description": topBottomID,
                        "Side": [
                            "Usec",
                            "Bear",
                            "Savage"
                        ],
                        "BodyPart": "Feet",
                        "Prefab": {
                            "path": topBottomPath,
                            "rcid": ""
                        },
                        "WatchPrefab": {
                            "path": "",
                            "rcid": ""
                        },
                        "IntegratedArmorVest": false,
                        "WatchPosition": {
                            "x": 0,
                            "y": 0,
                            "z": 0
                        },
                        "WatchRotation": {
                            "x": 0,
                            "y": 0,
                            "z": 0
                        }
                    },
                    "_proto": "5cdea3c47d6c8b0475341734"
                }

                // Initialising a new variable to assemble the suit
                let newSuit = {
                    "_id": `${topBottomID}Suit`,
                    "_name": `${topBottomID}Suit`,
                    "_parent": "5cd944d01388ce000a659df9",
                    "_type": "Item",
                    "_props": {
                        "Name": `${topBottomID}Suit`,
                        "ShortName": `${topBottomID}Suit`,
                        "Description": `${topBottomID}Suit`,
                        "Side": [
                            "Usec",
                            "Bear",
                            "Savage"
                        ],
                        "AvailableAsDefault": true,
                        "Feet": topBottomID
                    },
                    "_proto": "5cc085bb14c02e000e67a5c5"
                }

                // Initialising a new variable to assemble trader info
                let newTrader = {
                    "_id": topBottomID,
                    "tid": "5ac3b934156ae10c4430e83c",
                    "suiteId": `${topBottomID}Suit`,
                    "isActive": true,
                    "requirements": {
                        "loyaltyLevel": loyalty,
                        "profileLevel": level,
                        "standing": 0,
                        "skillRequirements": [],
                        "questRequirements": [],
                        "itemRequirements": [
                            {
                                "count": price,
                                "_tpl": "5449016a4bdc2d6f028b456f",
                                "onlyFunctional": true
                            }
                        ]
                    }
                }

                // Returning an array with all assembled stuff
                let fullBottomSuit = [newBottom, newSuit, newTrader, newLocale];
                return fullBottomSuit;

            default:
                // Item has hands path - thus it's top
                // Initialising a new variable to assemble the top
                let newTop = {
                    "_id": topBottomID,
                    "_name": topBottomID,
                    "_parent": "5cc0868e14c02e000c6bea68",
                    "_type": "Item",
                    "_props": {
                        "Name": topBottomID,
                        "ShortName": topBottomID,
                        "Description": topBottomID,
                        "Side": [
                            "Usec",
                            "Bear",
                            "Savage"
                        ],
                        "BodyPart": "Body",
                        "Prefab": {
                            "path": topBottomPath,
                            "rcid": ""
                        },
                        "WatchPrefab": {
                            "path": "",
                            "rcid": ""
                        },
                        "IntegratedArmorVest": false,
                        "WatchPosition": {
                            "x": 0,
                            "y": 0,
                            "z": 0
                        },
                        "WatchRotation": {
                            "x": 0,
                            "y": 0,
                            "z": 0
                        }
                    },
                    "_proto": "5cde9f337d6c8b0474535da8"
                }

                // Initialising a new variable to assemble the hands
                let newHands = {
                    "_id": `${topBottomID}Hands`,
                    "_name": `${topBottomID}Hands`,
                    "_parent": "5cc086a314c02e000c6bea69",
                    "_type": "Item",
                    "_props": {
                        "Name": `${topBottomID}Hands`,
                        "ShortName": `${topBottomID}Hands`,
                        "Description": `${topBottomID}Hands`,
                        "Side": [
                            "Usec",
                            "Bear",
                            "Savage"
                        ],
                        "BodyPart": "Hands",
                        "Prefab": {
                            "path": handsPath,
                            "rcid": ""
                        },
                        "WatchPrefab": {
                            "path": "",
                            "rcid": ""
                        },
                        "IntegratedArmorVest": false,
                        "WatchPosition": {
                            "x": 0,
                            "y": 0,
                            "z": 0
                        },
                        "WatchRotation": {
                            "x": 0,
                            "y": 0,
                            "z": 0
                        }
                    },
                    "_proto": "5cde95fa7d6c8b04737c2d13"
                }

                // Initialising a new variable to assemble the suit
                let newSuitTop = {
                    "_id": `${topBottomID}Suit`,
                    "_name": `${topBottomID}Suit`,
                    "_parent": "5cd944ca1388ce03a44dc2a4",
                    "_type": "Item",
                    "_props": {
                        "Name": `${topBottomID}Suit`,
                        "ShortName": `${topBottomID}Suit`,
                        "Description": `${topBottomID}Suit`,
                        "Side": [
                            "Usec",
                            "Bear",
                            "Savage"
                        ],
                        "AvailableAsDefault": true,
                        "Body": topBottomID,
                        "Hands": `${topBottomID}Hands`
                    },
                    "_proto": "5cde9ec17d6c8b04723cf479"
                }

                // Initialising a new variable to assemble trader info
                let newTraderTop = {
                    "_id": topBottomID,
                    "tid": "5ac3b934156ae10c4430e83c",
                    "suiteId": `${topBottomID}Suit`,
                    "isActive": true,
                    "requirements": {
                        "loyaltyLevel": loyalty,
                        "profileLevel": level,
                        "standing": 0,
                        "skillRequirements": [],
                        "questRequirements": [],
                        "itemRequirements": [
                            {
                                "count": price,
                                "_tpl": "5449016a4bdc2d6f028b456f",
                                "onlyFunctional": true
                            }
                        ]
                    }
                }

                // Returning an array with all assembled stuff
                let fullTopSuit = [newTop, newHands, newSuitTop, newTraderTop, newLocale];
                return fullTopSuit;
        }

    }

    // Haystack search function
    const findID = (haystack, needle) => {
        return needle.some(a => haystack.includes(a));
    };

    // Temporary variables
    let tempData = {};
    let progress = 0;

    logger.logInfo(`__________________________________________________________________________________`);
    // Parsing database folder
    for (let file of ModFileNames) {
        // Parsing data in our database folder
        let fileData = global.fileIO.readParsed(PathResolver(`user/mods/${ModFolderName}/database/${file}.json`));
        // Saving read data to our temporary variable
        tempData[file] = fileData;
        logger.logInfo(`[DATABASE] ${file} database is loaded`);
    }

    for (let curFile in tempData) {
        // Parsing each database for separate items
        logger.logInfo(`__________________________________________________________________________________`);
        logger.logSuccess(`[PROGRESS] ${progress*100/Object.keys(tempData).length}%...`);
        logger.logInfo(`[DATABASE] Parsing ${curFile} database`);
        progress += 1;

        for (let curItem in tempData[curFile]) {
            logger.logInfo(`_________________________________________`);
            // Setting up variables, flags, checking for existing info
            let curAddTrader = false;
            let curMods = false;
            let modifierFlag = tempData[curFile][curItem].modifier;
            let curBasedOn = false;
            let curItemInfo = false;
            let curTemplate = false;
            let curTranslation = false;
            let curTrader = {
                "_id": curItem,
                "_tpl": curItem,
                "parentId": "hideout",
                "slotId": "hideout",
                "upd": {
                    "UnlimitedCount": true,
                    "StackObjectsCount": 99999999
                }
            }

            // Checking for modifier flags
            switch (modifierFlag) {
                case "useBase":
                    // Current item isn't an edit file, but it requests to useBase
                    // Using and modifying base item info in accordance with current item info
                    logger.logInfo(`[BASE ITEM] Started loading ${curItem}`);
                    curItemInfo = tempData[curFile][curItem].item;
                    curBasedOn = tempData[curFile][curItem].basedOn;
                    curBaseInfo = JSON.parse(JSON.stringify(items.data[curBasedOn]));
                    curBaseInfo._id = curItemInfo._id;
                    curBaseInfo._name = curItemInfo._name;
                    for (let curItemChange in curItemInfo._props) {
                        curBaseInfo._props[curItemChange] = curItemInfo._props[curItemChange]
                    }
                    curItemInfo = curBaseInfo;
                    logger.logInfo(`[BASE ITEM] Item info copied in accordance with given base.`);

                    // Assembling current template
                    curTemplate = {
                        "Id": curItem,
                        "ParentId": false,
                        "Price": curItemInfo._props.CreditsPrice
                    }
                    // Searching for based on item in templates to find ParentId
                    for (let baseParentSearch in templates.data.Items) {
                        if (templates.data.Items[baseParentSearch].Id === curBasedOn) {
                            curTemplate.ParentId = templates.data.Items[baseParentSearch].ParentId;
                            break;
                        }
                    }
                    if (curTemplate.ParentId === false) {
                        logger.logWarning(`[WARNING] Can't find category for base item, using tushonka category.`);
                        curTemplate.ParentId = "5b47574386f77428ca22b336";
                    }

                    curTranslation = tempData[curFile][curItem].translation;
                    // Checking if we need to add to trader
                    if (tempData[curFile][curItem].hasOwnProperty("addToTrader")) {
                        curAddTrader = tempData[curFile][curItem].addToTrader;
                        curAddTrader.barter_scheme = [
                            [{ "count": curTemplate.Price, "_tpl": curAddTrader.currency }]
                        ];
                        logger.logInfo(`[ADDTOTRADER] Item ${curItem} has add to trader request.`)
                    }
                    break;

                case "edit":
                    // Checking if we need to edit any item properties
                    logger.logInfo(`[CHANGE] Started loading changes for ${curItem}.`);

                    if (tempData[curFile][curItem].hasOwnProperty("item")) {
                        curItemInfo = tempData[curFile][curItem].item;
                        curBasedOn = curItem;
                        logger.logInfo(`[CHANGE] ${curItem} has item properties.`);

                        // Checking if we need to edit the template
                        switch (curItemInfo.hasOwnProperty("CreditsPrice")) {
                            case true:
                                curTemplate = {
                                    "Price": curItemInfo.CreditsPrice
                                }
                                if (tempData[curFile][curItem].hasOwnProperty("addToTrader")) {
                                    curAddTrader = tempData[curFile][curItem].addToTrader;
                                    curAddTrader.barter_scheme = [
                                        [{ "count": curTemplate.Price, "_tpl": curAddTrader.currency }]
                                    ];
                                    logger.logInfo(`[ADDTOTRADER] Item ${curItem} has add to trader request.`)
                                }
                                break;
                            case false:
                                logger.logInfo(`[CHANGE] ${curItem} doesn't have a credits price, template won't be changed.`);
                                if (tempData[curFile][curItem].hasOwnProperty("addToTrader")) {
                                    curAddTrader = tempData[curFile][curItem].addToTrader;
                                    curAddTrader.barter_scheme = [
                                        [{ "count": items.data[curItem]._props.CreditsPrice, "_tpl": curAddTrader.currency }]
                                    ];
                                    logger.logInfo(`[ADDTOTRADER] Item ${curItem} has add to trader request.`)
                                }
                        }
                    }

                    // Checking if we need to edit translations
                    if (tempData[curFile][curItem].hasOwnProperty("translation")) {
                        curTranslation = tempData[curFile][curItem].translation;
                        logger.logInfo(`[CHANGE] ${curItem} has localisation.`);
                    }
                    break;

                case "suit":
                    // Current item is a suit
                    switch (tempData[curFile][curItem].hasOwnProperty("bottom")) {
                        case true:
                            logger.logInfo(`[SUIT] ${curItem} is the bottom part of a suit.`);
                            let fullNewBottom = topBottomAssemble(curItem, tempData[curFile][curItem].bottom, false, tempData[curFile][curItem].loyalty, tempData[curFile][curItem].level, tempData[curFile][curItem].price, tempData[curFile][curItem].name);
                            logger.logInfo(`[SUIT] ${curItem} templates are complete.`);

                            // Adding the above generated templates to customization, Ragman and locales
                            suitsGlobal.data[curItem] = fullNewBottom[0];
                            suitsGlobal.data[`${curItem}Suit`] = fullNewBottom[1];
                            suitsRagman.push(fullNewBottom[2]);
                            locale_en.templates[`${curItem}Suit`] = fullNewBottom[3];
                            logger.logInfo(`[SUIT] ${curItem} was successfully added to the game.`);
                            continue;

                        case false:
                            logger.logInfo(`[SUIT] ${curItem} is the top part of a suit.`);
                            let fullNewTop = topBottomAssemble(curItem, tempData[curFile][curItem].top, tempData[curFile][curItem].hands, tempData[curFile][curItem].loyalty, tempData[curFile][curItem].level, tempData[curFile][curItem].price, tempData[curFile][curItem].name);
                            logger.logInfo(`[SUIT] ${curItem} templates are complete.`);

                            // Adding the above generated templates to customization, Ragman and locales
                            suitsGlobal.data[curItem] = fullNewTop[0];
                            suitsGlobal.data[`${curItem}Hands`] = fullNewTop[1];
                            suitsGlobal.data[`${curItem}Suit`] = fullNewTop[2];
                            suitsRagman.push(fullNewTop[3]);
                            locale_en.templates[`${curItem}Suit`] = fullNewTop[4];
                            logger.logInfo(`[SUIT] ${curItem} was successfully added to the game.`);

                            continue;
                    }

                case "none":
                    // Current item isn't an edit file, neither does it request to useBase
                    logger.logInfo(`[ITEM] Started loading ${curItem}`);
                    curItemInfo = tempData[curFile][curItem].item;
                    curBasedOn = tempData[curFile][curItem].basedOn;

                    // Assembling current template
                    curTemplate = {
                        "Id": curItem,
                        "ParentId": false,
                        "Price": curItemInfo._props.CreditsPrice
                    }
                    // Searching for based on item in templates to find ParentId
                    for (let baseParentSearch in templates.data.Items) {
                        if (templates.data.Items[baseParentSearch].Id === curBasedOn) {
                            curTemplate.ParentId = templates.data.Items[baseParentSearch].ParentId;
                            break;
                        }
                    }
                    if (curTemplate.ParentId === false) {
                        logger.logWarning(`[WARNING] Item's based on is incorrect! Using tushonka category.`);
                        curTemplate.ParentId = "5b47574386f77428ca22b336";
                    }

                    curTranslation = tempData[curFile][curItem].translation;
                    // Checking if we need to add to trader
                    if (tempData[curFile][curItem].hasOwnProperty("addToTrader")) {
                        curAddTrader = tempData[curFile][curItem].addToTrader;
                        curAddTrader.barter_scheme = [
                            [{ "count": curTemplate.Price, "_tpl": curAddTrader.currency }]
                        ];
                        logger.logInfo(`[ADDTOTRADER] Item ${curItem} has add to trader request.`)
                    }
                    break;
            }

            // Checking if item has to be added to mod slots
            if (tempData[curFile][curItem].hasOwnProperty("addMods")) {
                curMods = tempData[curFile][curItem].addMods;
                logger.logInfo(`[ATTACHMENT] Item ${curItem} has add to mod slots request.`)
            }


            // Ragfair config. If file is editing existing info: the item isn't added to the ragfair a 2nd time
            switch (modifierFlag) {
                case "edit":
                    assort_ragfair.data.barter_scheme[curTrader._id] = [
                        [{ "count": curTemplate.Price, "_tpl": "5449016a4bdc2d6f028b456f" }]
                    ];
                    logger.logInfo(`[CHANGE] Item ragfair price was edited.`);
                    break;
                default:
                    assort_ragfair.data.items.push(curTrader);
                    assort_ragfair.data.loyal_level_items[curTrader._id] = 1;
                    assort_ragfair.data.barter_scheme[curTrader._id] = [
                        [{ "count": curTemplate.Price, "_tpl": "5449016a4bdc2d6f028b456f" }]
                    ];
                    logger.logInfo(`[ITEM] Item was added to ragfair.`)
            }

            // Items config. If file is editing item info: the corresponding item properties will be changed
            switch (modifierFlag) {
                case "edit":
                    if (curItemInfo) {
                        for (let curItemChange in curItemInfo) {
                            items.data[curItem]._props[curItemChange] = curItemInfo[curItemChange]
                        }
                        logger.logInfo(`[CHANGE] Item properties were edited.`);
                    }
                    break;
                default:
                    items.data[curItem] = curItemInfo;
                    logger.logInfo(`[ITEM] Item properties were added.`);
            }

            // Locales config. If file is editing item info: localisation is overwritten.
            if (curTranslation) {
                locale_en.templates[curItem] = curTranslation.english;
                locale_fr.templates[curItem] = curTranslation.french;
                locale_ge.templates[curItem] = curTranslation.german;
                locale_ru.templates[curItem] = curTranslation.russian;
                logger.logInfo(`[ITEM] Localisation was updated.`)
            }

            // Templates config. If file is editing item info: only price is changed.
            switch (modifierFlag) {
                case "edit":
                    for (let curTemplateObject in templates.data.Items) {
                        if (curTemplateObject.hasOwnProperty(curItem)) {
                            curTemplateObject.Price = curTemplate.Price;
                            break;
                        }
                    }
                    logger.logInfo(`[CHANGE] Item template price was changed.`);
                    break;
                default:
                    templates.data.Items.push(curTemplate);
                    logger.logInfo(`[ITEM] Item template was added.`);
            }

            // Adding item to item mods.
            if (curMods) {
                for (let curMod in curMods) {
                    switch (curMod) {
                        case "allCompatible":
                            // Searching for all items compatible with given items
                            let compatibleIDs = curMods[curMod];
                            for (let curModSearch in items.data) {
                                if (curModSearch != curItem) {
                                    let curIDSlots = items.data[curModSearch]._props.Slots;
                                    // Cycle below checks all mod slots for given item IDs
                                    for (let checkMod in curIDSlots) {
                                        let curSlot = curIDSlots[checkMod]._props.filters[0].Filter;
                                        if (findID(curSlot, compatibleIDs)) {
                                            curIDSlots[checkMod]._props.filters[0].Filter.push(curItem);
                                            logger.logInfo(`[ATTACHMENT] Added ${curItem} to ${curModSearch}.`);
                                        }
                                    }
                                }
                            }
                            break;
                        default:
                            // Adding mods to specific slots in specific weapons
                            let curModIDs = curMods[curMod];
                            for (let curModID in curModIDs) {
                                switch (curModID != curItem) {
                                    case true:
                                        let curIDSlots = items.data[curModIDs[curModID]]._props.Slots;
                                        for (let checkMod in curIDSlots) {
                                            if (curIDSlots[checkMod]._name == curMod) {
                                                curIDSlots[checkMod]._props.filters[0].Filter.push(curItem);
                                                logger.logInfo(`[ATTACHMENT] Added ${curItem} to ${curModIDs[curModID]}.`);
                                                break;
                                            }
                                        }
                                        break;
                                    case false:
                                        logger.logWarning(`[WARNING] ${curItem} had a request to add to itself. Denied.`)
                                }
                            }
                    }

                }
            }


            // Trader Assort config. If file is editing - the code checks for existing assort and behaves accordingly.
            for (let curAdd in curAddTrader.traderName) {
                switch (curAddTrader.traderName[curAdd]) {
                    case "Prapor":
                        assortPrapot = pushTrader(assortPrapor, curAddTrader, curItem, curTrader, "Prapor", modifierFlag);
                        break;
                    case "Therapist":
                        assortTherapist = pushTrader(assortTherapist, curAddTrader, curItem, curTrader, "Therapist", modifierFlag);
                        break;
                    case "Fence":
                        assortFence = pushTrader(assortFence, curAddTrader, curItem, curTrader, "Fence", modifierFlag);
                        break;
                    case "Skier":
                        assortSkier = pushTrader(assortSkier, curAddTrader, curItem, curTrader, "Skier", modifierFlag);
                        break;
                    case "Peacekeeper":
                        assortPeacekeeper = pushTrader(assortPeacekeeper, curAddTrader, curItem, curTrader, "Peacekeeper", modifierFlag);
                        break;
                    case "Mechanic":
                        assortMechanic = pushTrader(assortMechanic, curAddTrader, curItem, curTrader, "Mechanic", modifierFlag);
                        break;
                    case "Ragman":
                        assortRagman = pushTrader(assortRagman, curAddTrader, curItem, curTrader, "Ragman", modifierFlag);
                        break;
                    case "Jaeger":
                        assortJaeger = pushTrader(assortJaeger, curAddTrader, curItem, curTrader, "Jaeger", modifierFlag);
                        break;
                }
            }

            // Logging info about current item
            switch (modifierFlag) {
                case "edit":
                    logger.logInfo(`[CHANGE] ${curItem} has been changed.`);
                    break;
                case "useBase":
                    logger.logInfo(`[BASE ITEM] ${curItem} has been loaded using a base.`);
                    break;
                case "none":
                    logger.logInfo(`[ITEM] ${curItem} has been loaded.`);
            }

        }
    }

    // Writing to file
    fileIO.write(PathResolver('user/cache/items.json'), items, true);
    fileIO.write(PathResolver('user/cache/assort_ragfair.json'), assort_ragfair, true);
    fileIO.write(PathResolver('user/cache/locale_en.json'), locale_en, true);
    fileIO.write(PathResolver(`user/mods/${ModFolderName}/src/Languages/Altered/fr/locale.json`), locale_fr, true);
    fileIO.write(PathResolver(`user/mods/${ModFolderName}/src/Languages/Altered/ge/locale.json`), locale_ge, true);
    fileIO.write(PathResolver('user/cache/locale_ru.json'), locale_ru, true);
    fileIO.write(PathResolver('user/cache/templates.json'), templates, true);
    fileIO.write(PathResolver('user/cache/assort_54cb50c76803fa8b248b4571.json'), assortPrapor, true);
    fileIO.write(PathResolver('user/cache/assort_54cb57776803fa99248b456e.json'), assortTherapist, true);
    fileIO.write(PathResolver('user/cache/assort_579dc571d53a0658a154fbec.json'), assortFence, true);
    fileIO.write(PathResolver('user/cache/assort_58330581ace78e27b8b10cee.json'), assortSkier, true);
    fileIO.write(PathResolver('user/cache/assort_5935c25fb3acc3127c3d8cd9.json'), assortPeacekeeper, true);
    fileIO.write(PathResolver('user/cache/assort_5a7c2eca46aef81a7ca2145d.json'), assortMechanic, true);
    fileIO.write(PathResolver('user/cache/assort_5ac3b934156ae10c4430e83c.json'), assortRagman, true);
    fileIO.write(PathResolver('user/cache/assort_5c0647fdd443bc2504c2d371.json'), assortJaeger, true);
    fileIO.write(PathResolver('user/cache/customization.json'), suitsGlobal, true);
    fileIO.write(PathResolver('user/cache/customization_5ac3b934156ae10c4430e83c.json'), suitsRagman, true);

    // Logging success on modload
    logger.logInfo(`_________________________________________`);
    logger.logSuccess(`[MOD] ${mod_data.name} Applied`);
}