exports.mod = (mod_info) => {

    logger.logInfo("[FLEALOCK] Calling Punisher...");

    //ragfair.js mod
    ragfair_f = require("./FleaLock/ragfair.js");
    global._database.gameplayConfig = fileIO.readParsed("./user/mods/FireEqual-StructureOverhaul-1.0.0/src/FleaLock/gameplay.json")

    logger.logSuccess("[FLEALOCK] Punisher defending Flea Market!");

    logger.logInfo("[LANGUAGES] Loading French and German language packs...");

    // Loading French and German locales into memory
    // This doesn't work yet because v16 broke languages
    const cacheLoad = function (filepath) { return global.fileIO.readParsed(filepath) }
    let ModFolderName = `FireEqual-StructureOverhaul-1.0.0`;
    _database.locales.global['fr'] = cacheLoad(`user/mods/${ModFolderName}/src/Languages/Altered/fr/locale.json`);
    _database.locales.menu['fr'] = cacheLoad(`user/mods/${ModFolderName}/src/Languages/Altered/fr/menu.json`);
    _database.languages['fr'] = cacheLoad(`user/mods/${ModFolderName}/src/Languages/Altered/fr/fr.json`);
    _database.locales.global['ge'] = cacheLoad(`user/mods/${ModFolderName}/src/Languages/Altered/ge/locale.json`);
    _database.locales.menu['ge'] = cacheLoad(`user/mods/${ModFolderName}/src/Languages/Altered/ge/menu.json`);
    _database.languages['ge'] = cacheLoad(`user/mods/${ModFolderName}/src/Languages/Altered/ge/ge.json`);
    logger.logSuccess("[LANGUAGES] French and German language packs loaded!");



    logger.logInfo("[INVENTORY] Optimizing Inventory");
    //Inventory.json fix so that it loads parentIDs instead of multiple itemID's
    global.db.items.Inventory = "user/mods/FireEqual-StructureOverhaul-1.0.0/src/DeveloperProfile/Inventory.json"
    logger.logSuccess("[INVENTORY] Optimized Inventory");

    

    //developer profile overhaul attempt (with the help of `life`)
    let profileDir = internal.path.resolve(__dirname, "DeveloperProfile");
    let files = fileIO.readDir(profileDir);
    let profile = {}
    for (let index in files) {
        let file = files[index]
        let fileName = file.split('.')[0]
        let fullPath = internal.path.resolve(profileDir, file)
        profile[fileName] = fullPath
    }
    global.db.profile["Developer"] = profile;
    global.db.profile["Altered Escape Developer"] = global.db.profile["Developer"];
    delete global.db.profile["Developer"];

    let globals = fileIO.readParsed(db.base.globals);
    let king_globals = internal.path.resolve(__dirname, "globals");
    let tweaks = globals.data;
    tweaks.ItemPresets = [];
    db.base.globals = king_globals;
    fileIO.write(king_globals, globals);
  
  /*
  Below are the search paths if they need to be changed in the future (which they will):
      global_presets = fileIO.readParsed("./user/mods/FireEqual-StructureOverhaul-1.0.0/src/Presets/Presets.json");
      this.presets = global_presets = fileIO.readParsed("./user/mods/FireEqual-StructureOverhaul-1.0.0/src/Presets/Presets.json");
  */
  
    logger.logInfo("[PRESETS] Giving Mechanic some cigs...");
    preset_f = require("./Presets/preset.js");
    preset_f.handler.initialize();
    //preset_f.handler.initialize(); is called on the new preset_f to repopulate data (that would be missing without);
    helper_f.getPreset = require("./Presets/helper.js").getPreset;
    //`classes/location.js` uses `helper_f.getPreset` in `function _GenerateContainerLoot(_items);
    helper_f.getItem = require("./Presets/helper.js").getItem;
    //trying to get this thing to read the cache?
    location_f.FindIfItemIsAPreset = require("./Presets/location.js").FindIfItemIsAPreset;
    //function FindIfItemIsAPreset tries to load from `global._database.globals.ItemPreset`
    move_f.addItem = require("./Presets/move.js").addItem;
  
    logger.logSuccess("[PRESETS] Mechanic is Pimpin' my Mosin!");


    logger.logSuccess("[SUITS] Locating Drip...");

    //customization & trader.js
    //customization_f = require("./MakeSuitsWorkAgain/customization.js");
    trader_f = require("./MakeSuitsWorkAgain/trader.js");

    logger.logSuccess("[SUITS] Drip Located!");
}