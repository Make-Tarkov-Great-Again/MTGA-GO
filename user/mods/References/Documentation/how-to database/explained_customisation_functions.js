class SuckMeOffAKIAPIs {

    /* I just did the neighborly thing and snagged the way they do it as reference for ya
    I put bundles you can test in the bundles/Clothing folder pimp, have fun! */


    static AddTop(db, OutfitID, TopBundlePath, HandsBundlePath, handsBaseID) {
        // add top
        let NewTop = JsonUtil.clone(DatabaseServer.tables.templates.customization["5d28adcb86f77429242fc893"]); //Loads reshala top as template
        /*let reshalaTop = {
            "5d28adcb86f77429242fc893": {
                "_id": "5d28adcb86f77429242fc893",
                "_name": "Wild_Dealmaker_body",
                "_parent": "5cc0868e14c02e000c6bea68",
                "_type": "Item",
                "_props": {
                    "Name": "wild_body_1",
                    "ShortName": "wild_body_1",
                    "Description": "wild_body_1",
                    "Side": [
                        "Savage"
                    ],
                    "BodyPart": "Body",
                    "Prefab": {
                        "path": "assets/content/characters/character/prefabs/wild_dealmaker_body.bundle",
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
        }*/

        NewTop._id = OutfitID; // Changes _id to outfit ID
        NewTop._name = OutfitID; // Changes _name to outfit ID
        NewTop._props.Prefab.path = TopBundlePath; // Changes prefab path to top bundle path
        DatabaseServer.tables.templates.customization[OutfitID] = NewTop; // Adds the above changed template to customisation cache

        // add hands
        let NewHands = JsonUtil.clone(DatabaseServer.tables.templates.customization[handsBaseID]); // Loads hands as template (ID depends on input, given example is usec ironsight)
        /*let usecHands = {
            "5d1f5b5386f7744bcc048757": {
                "_id": "5d1f5b5386f7744bcc048757",
                "_name": "usec_hands_pcuironsight",
                "_parent": "5cc086a314c02e000c6bea69",
                "_type": "Item",
                "_props": {
                    "Name": "DefaultUsecHands",
                    "ShortName": "DefaultUsecHands",
                    "Description": "DefaultUsecHands",
                    "Side": [
                        "Usec"
                    ],
                    "BodyPart": "Hands",
                    "Prefab": {
                        "path": "assets/content/hands/hands_usec_orc_pcu/hands_usec_orc_pcu.skin.bundle",
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
            },
        }*/

        NewHands._id = `${OutfitID}Hands`; // Changes _id to outfit ID + "Hands" at the end
        NewHands._name = `${OutfitID}Hands`; // Changes _name to outfit ID + "Hands" at the end
        NewHands._props.Prefab.path = HandsBundlePath; // Changes prefab path to hands bundle path
        DatabaseServer.tables.templates.customization[`${OutfitID}Hands`] = NewHands; // Adds the above changed template to customisation cache

        // add suite
        let NewSuite = JsonUtil.clone(DatabaseServer.tables.templates.customization["5d1f623e86f7744bce0ef705"]); // Loads usec kit upper pcuironsight as template
        /*let usecKit = {
            "5d1f623e86f7744bce0ef705": {
                "_id": "5d1f623e86f7744bce0ef705",
                "_name": " usec_kit_upper_pcuironsight",
                "_parent": "5cd944ca1388ce03a44dc2a4",
                "_type": "Item",
                "_props": {
                    "Name": "DefaultUsecUpperSuite",
                    "ShortName": "DefaultUsecUpperSuite",
                    "Description": "DefaultUsecUpperSuite",
                    "Side": [
                        "Usec"
                    ],
                    "AvailableAsDefault": false,
                    "Body": "5d1f56a686f7744bce0ee9eb",
                    "Hands": "5d1f5b5386f7744bcc048757"
                },
                "_proto": "5cde9ec17d6c8b04723cf479"
            },
        }*/

        NewSuite._id = `${OutfitID}Suite`; // Changes _id to the outfitID + "Suite" at the end
        NewSuite._name = `${OutfitID}Suite`; // Changes _name to the outfitID + "Suite" at the end
        NewSuite._props.Body = OutfitID; // Sets the body in suite to the one we made
        NewSuite._props.Hands = `${OutfitID}Hands`; // Sets the hands in suite to the ones we made
        NewSuite._props.Side = ["Usec", "Bear", "Savage"]; // Sets sides for suite
        DatabaseServer.tables.templates.customization[`${OutfitID}Suite`] = NewSuite; // Adds the above changed template to customisation cache

        // locale
        for (const localeID in DatabaseServer.tables.locales.global) {
            // en placeholder
            DatabaseServer.tables.locales.global[localeID].templates[`${OutfitID}Suite`] = JsonUtil.deserialize(VFS.readFile(`${db}locales/en.json`))[OutfitID];

            // actual locale
            if (VFS.exists(`${db}locales/${localeID}.json`)) { // This isn't needed, since we make locales inside the file itself
                DatabaseServer.tables.locales.global[localeID].templates[`${OutfitID}Suite`] = JsonUtil.deserialize(VFS.readFile(`${db}locales/${localeID}.json`))[OutfitID]; // Adds locale to locale_en.templates
            }
            /*let localeExample = {
                "5d1f623e86f7744bce0ef705": {
                    "Description": "",
                    "Name": "USEC PCU Ironsight",
                    "ShortName": ""
                },
            }*/
        }

        // add suite to the ragman
        DatabaseServer.tables.traders["5ac3b934156ae10c4430e83c"].suits.push({ // Pushes below template to a trader
            "_id": OutfitID,
            "tid": "5ac3b934156ae10c4430e83c", //same as trader id
            "suiteId": `${OutfitID}Suite`,
            "isActive": true,
            "requirements": {
                "loyaltyLevel": 0,
                "profileLevel": 0,
                "standing": 0,
                "skillRequirements": [],
                "questRequirements": [],
                "itemRequirements": [
                    {
                        "count": 0,
                        "_tpl": "5449016a4bdc2d6f028b456f"
                    }
                ]
            }
        });
    }

    static AddBottom(db, OutfitID, BundlePath) {
        // basically same shit but bottom. Doesn't have hands
        // add Bottom
        let NewBottom = JsonUtil.clone(DatabaseServer.tables.templates.customization["5d5e7f4986f7746956659f8a"]);
        /*let bearPants = {
            "5d5e7f4986f7746956659f8a": {
                "_id": "5d5e7f4986f7746956659f8a",
                "_name": "Pants_security_Gorka4",
                "_parent": "5cc0869814c02e000a4cad94",
                "_type": "Item",
                "_props": {
                    "Name": "Pants_security_Gorka4",
                    "ShortName": "Pants_security_Gorka4",
                    "Description": "Pants_security_Gorka4",
                    "Side": [
                        "Savage"
                    ],
                    "BodyPart": "Feet",
                    "Prefab": {
                        "path": "assets/content/characters/character/prefabs/pants_security_gorka4.bundle",
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
            },
        }*/

        NewBottom._id = OutfitID;
        NewBottom._name = OutfitID;
        NewBottom._props.Prefab.path = BundlePath;
        DatabaseServer.tables.templates.customization[OutfitID] = NewBottom;

        // add suite
        let NewSuite = JsonUtil.clone(DatabaseServer.tables.templates.customization["5cd946231388ce000d572fe3"]);
        /*let bearLower = {
            "5cd946231388ce000d572fe3": {
                "_id": "5cd946231388ce000d572fe3",
                "_name": "DefaultBearLowerSuite",
                "_parent": "5cd944d01388ce000a659df9",
                "_type": "Item",
                "_props": {
                    "Name": "DefaultBearLowerSuite",
                    "ShortName": "DefaultBearLowerSuite",
                    "Description": "DefaultBearLowerSuite",
                    "Side": [
                        "Bear"
                    ],
                    "AvailableAsDefault": true,
                    "Feet": "5cc085bb14c02e000e67a5c5"
                }
            },
        }*/

        NewSuite._id = `${OutfitID}Suite`;
        NewSuite._name = `${OutfitID}Suite`;
        NewSuite._props.Feet = OutfitID;
        NewSuite._props.Side = ["Usec", "Bear", "Savage"];
        DatabaseServer.tables.templates.customization[`${OutfitID}Suite`] = NewSuite;

        // locale
        for (const localeID in DatabaseServer.tables.locales.global) {
            // en placeholder
            DatabaseServer.tables.locales.global[localeID].templates[`${OutfitID}Suite`] = JsonUtil.deserialize(VFS.readFile(`${db}locales/en.json`))[OutfitID];

            // actual locale
            if (VFS.exists(`${db}locales/${localeID}.json`)) {
                DatabaseServer.tables.locales.global[localeID].templates[`${OutfitID}Suite`] = JsonUtil.deserialize(VFS.readFile(`${db}locales/${localeID}.json`))[OutfitID];
            }
        }

        // add suite to the ragman
        DatabaseServer.tables.traders["5ac3b934156ae10c4430e83c"].suits.push({
            "_id": OutfitID,
            "tid": "5ac3b934156ae10c4430e83c",
            "suiteId": `${OutfitID}Suite`,
            "isActive": true,
            "requirements": {
                "loyaltyLevel": 0,
                "profileLevel": 0,
                "standing": 0,
                "skillRequirements": [],
                "questRequirements": [],
                "itemRequirements": [
                    {
                        "count": 0,
                        "_tpl": "5449016a4bdc2d6f028b456f"
                    }
                ]
            }
        });
    }
}