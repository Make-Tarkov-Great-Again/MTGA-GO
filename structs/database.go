package structs

type Database struct {
	Core *Core
	//Connections *ConnectionStruct
	Items     map[string]*DatabaseItem
	Locales   *Locale
	Languages map[string]string
	Handbook  *Handbook
	Traders   map[string]*Trader
	Flea      *Flea
	Quests    map[string]*Quest
	Hideout   *Hideout

	Locations     *Locations
	Weather       *Weather
	Customization map[string]*Customization
	Editions      map[string]*Edition
	Bot           *Bots
	Profiles      map[string]*Profile
	//bundles  []map[string]interface{}
}

type Edition struct {
	Bear    *PlayerTemplate `json:"bear"`
	Usec    *PlayerTemplate `json:"usec"`
	Storage struct {
		Bear []string `json:"bear"`
		Usec []string `json:"usec"`
	} `json:"storage"`
}

type Core struct {
	PlayerTemplate    PlayerTemplate
	PlayerScav        PlayerTemplate
	ClientSettings    ClientSettings
	ServerConfig      ServerConfig
	Globals           Globals
	GlobalBotSettings GlobalBotSettings
	//gameplay        map[string]interface{}
	//blacklist       []interface{}
	MatchMetrics MatchMetrics
}

type GlobalBotSettings struct {
	SavageKillDist                           int16   `json:"SAVAGE_KILL_DIST"`
	SoundDoorBreachMeters                    float32 `json:"SOUND_DOOR_BREACH_METERS"`
	SoundDoorOpenMeters                      float32 `json:"SOUND_DOOR_OPEN_METERS"`
	StepNoiseDelta                           float32 `json:"STEP_NOISE_DELTA"`
	JumpNoiseDelta                           float32 `json:"JUMP_NOISE_DELTA"`
	GunshotSpread                            int16   `json:"GUNSHOT_SPREAD"`
	GunshotSpreadSilence                     int16   `json:"GUNSHOT_SPREAD_SILENCE"`
	BaseWalkSperead2                         float32 `json:"BASE_WALK_SPEREAD2"`
	MoveSpeedCoefMax                         float32 `json:"MOVE_SPEED_COEF_MAX"`
	SpeedServSoundCoefA                      float32 `json:"SPEED_SERV_SOUND_COEF_A"`
	SpeedServSoundCoefB                      float32 `json:"SPEED_SERV_SOUND_COEF_B"`
	G                                        float32 `json:"G"`
	StayCoef                                 float32 `json:"STAY_COEF"`
	SitCoef                                  float32 `json:"SIT_COEF"`
	LayCoef                                  float32 `json:"LAY_COEF"`
	MaxIterations                            int16   `json:"MAX_ITERATIONS"`
	StartDistToCov                           float32 `json:"START_DIST_TO_COV"`
	MaxDistToCov                             float32 `json:"MAX_DIST_TO_COV"`
	StayHeight                               float32 `json:"STAY_HEIGHT"`
	ClosePoints                              float32 `json:"CLOSE_POINTS"`
	CountTurns                               int16   `json:"COUNT_TURNS"`
	SimplePointLifeTimeSec                   float32 `json:"SIMPLE_POINT_LIFE_TIME_SEC"`
	DangerPointLifeTimeSec                   float32 `json:"DANGER_POINT_LIFE_TIME_SEC"`
	DangerPower                              float32 `json:"DANGER_POWER"`
	CoverDistClose                           float32 `json:"COVER_DIST_CLOSE"`
	GoodDistToPoint                          float32 `json:"GOOD_DIST_TO_POINT"`
	CoverToofarFromBoss                      float32 `json:"COVER_TOOFAR_FROM_BOSS"`
	CoverToofarFromBossSqrt                  float32 `json:"COVER_TOOFAR_FROM_BOSS_SQRT"`
	MaxYDiffToProtect                        float32 `json:"MAX_Y_DIFF_TO_PROTECT"`
	FlarePower                               float32 `json:"FLARE_POWER"`
	MoveCoef                                 float32 `json:"MOVE_COEF"`
	PronePose                                float32 `json:"PRONE_POSE"`
	LowerPose                                float32 `json:"LOWER_POSE"`
	MaxPose                                  float32 `json:"MAX_POSE"`
	FlareTime                                float32 `json:"FLARE_TIME"`
	MaxRequestsPerGroup                      int16   `json:"MAX_REQUESTS__PER_GROUP"`
	UpdateGoalTimerSec                       float32 `json:"UPDATE_GOAL_TIMER_SEC"`
	DistNotToGroup                           float32 `json:"DIST_NOT_TO_GROUP"`
	DistNotToGroupSqr                        float32 `json:"DIST_NOT_TO_GROUP_SQR"`
	LastSeenPosLifetime                      float32 `json:"LAST_SEEN_POS_LIFETIME"`
	DeltaGrenadeStartTime                    float32 `json:"DELTA_GRENADE_START_TIME"`
	DeltaGrenadeEndTime                      float32 `json:"DELTA_GRENADE_END_TIME"`
	DeltaGrenadeRunDist                      int16   `json:"DELTA_GRENADE_RUN_DIST"`
	DeltaGrenadeRunDistSqrt                  float32 `json:"DELTA_GRENADE_RUN_DIST_SQRT"`
	PatrolMinLightDist                       float32 `json:"PATROL_MIN_LIGHT_DIST"`
	HoldMinLightDist                         float32 `json:"HOLD_MIN_LIGHT_DIST"`
	StandartBotPauseDoor                     float32 `json:"STANDART_BOT_PAUSE_DOOR"`
	ArmorClassCoef                           float32 `json:"ARMOR_CLASS_COEF"`
	ShotgunPower                             float32 `json:"SHOTGUN_POWER"`
	RiflePower                               float32 `json:"RIFLE_POWER"`
	PistolPower                              float32 `json:"PISTOL_POWER"`
	SmgPower                                 float32 `json:"SMG_POWER"`
	SnipePower                               float32 `json:"SNIPE_POWER"`
	GestusPeriodSec                          float32 `json:"GESTUS_PERIOD_SEC"`
	GestusAimingDelay                        float32 `json:"GESTUS_AIMING_DELAY"`
	GestusRequestLifetime                    float32 `json:"GESTUS_REQUEST_LIFETIME"`
	GestusFirstStageMaxTime                  float32 `json:"GESTUS_FIRST_STAGE_MAX_TIME"`
	GestusSecondStageMaxTime                 float32 `json:"GESTUS_SECOND_STAGE_MAX_TIME"`
	GestusMaxAnswers                         int16   `json:"GESTUS_MAX_ANSWERS"`
	GestusFuckToShoot                        int16   `json:"GESTUS_FUCK_TO_SHOOT"`
	GestusDistAnswers                        float32 `json:"GESTUS_DIST_ANSWERS"`
	GestusDistAnswersSqrt                    float32 `json:"GESTUS_DIST_ANSWERS_SQRT"`
	GestusAnywayChance                       float32 `json:"GESTUS_ANYWAY_CHANCE"`
	TalkDelay                                float32 `json:"TALK_DELAY"`
	CanShootToHead                           bool    `json:"CAN_SHOOT_TO_HEAD"`
	CanTilt                                  bool    `json:"CAN_TILT"`
	TiltChance                               float32 `json:"TILT_CHANCE"`
	MinBlockDist                             float32 `json:"MIN_BLOCK_DIST"`
	MinBlockTime                             float32 `json:"MIN_BLOCK_TIME"`
	CoverSecondsAfterLoseVision              float32 `json:"COVER_SECONDS_AFTER_LOSE_VISION"`
	MinArgCoef                               float32 `json:"MIN_ARG_COEF"`
	MaxArgCoef                               float32 `json:"MAX_ARG_COEF"`
	DeadAgrDist                              float32 `json:"DEAD_AGR_DIST"`
	MaxDangerCareDistSqrt                    float32 `json:"MAX_DANGER_CARE_DIST_SQRT"`
	MaxDangerCareDist                        float32 `json:"MAX_DANGER_CARE_DIST"`
	MinMaxPersonSearch                       int16   `json:"MIN_MAX_PERSON_SEARCH"`
	PercentPersonSearch                      float32 `json:"PERCENT_PERSON_SEARCH"`
	LookAnysideByWallSecOfEnemy              float32 `json:"LOOK_ANYSIDE_BY_WALL_SEC_OF_ENEMY"`
	CloseToWallRotateByWallSqrt              float32 `json:"CLOSE_TO_WALL_ROTATE_BY_WALL_SQRT"`
	ShootToChangeRndPartMin                  int16   `json:"SHOOT_TO_CHANGE_RND_PART_MIN"`
	ShootToChangeRndPartMax                  int16   `json:"SHOOT_TO_CHANGE_RND_PART_MAX"`
	ShootToChangeRndPartDelta                float32 `json:"SHOOT_TO_CHANGE_RND_PART_DELTA"`
	FormulCoefDeltaDist                      float32 `json:"FORMUL_COEF_DELTA_DIST"`
	FormulCoefDeltaShoot                     float32 `json:"FORMUL_COEF_DELTA_SHOOT"`
	FormulCoefDeltaFriendCover               float32 `json:"FORMUL_COEF_DELTA_FRIEND_COVER"`
	SuspetionPointDistCheck                  float32 `json:"SUSPETION_POINT_DIST_CHECK"`
	MaxBaseRequestsPerPlayer                 int16   `json:"MAX_BASE_REQUESTS_PER_PLAYER"`
	MaxHoldRequestsPerPlayer                 int16   `json:"MAX_HOLD_REQUESTS_PER_PLAYER"`
	MaxGoToRequestsPerPlayer                 int16   `json:"MAX_GO_TO_REQUESTS_PER_PLAYER"`
	MaxComeWithMeRequestsPerPlayer           int16   `json:"MAX_COME_WITH_ME_REQUESTS_PER_PLAYER"`
	CorePointMaxValue                        float32 `json:"CORE_POINT_MAX_VALUE"`
	CorePointsMax                            int16   `json:"CORE_POINTS_MAX"`
	CorePointsMin                            int16   `json:"CORE_POINTS_MIN"`
	BornPoistsFreeOnlyFarestBot              bool    `json:"BORN_POISTS_FREE_ONLY_FAREST_BOT"`
	BornPoinstsFreeOnlyFarestPlayer          bool    `json:"BORN_POINSTS_FREE_ONLY_FAREST_PLAYER"`
	ScavGroupsTogether                       bool    `json:"SCAV_GROUPS_TOGETHER"`
	LayDownAngShoot                          float32 `json:"LAY_DOWN_ANG_SHOOT"`
	HoldRequestTimeSec                       float32 `json:"HOLD_REQUEST_TIME_SEC"`
	TriggersDownToRunWhenMove                int16   `json:"TRIGGERS_DOWN_TO_RUN_WHEN_MOVE"`
	MinDistToRunWhileAttackMoving            float32 `json:"MIN_DIST_TO_RUN_WHILE_ATTACK_MOVING"`
	MinDistToRunWhileAttackMovingOtherEnemis float32 `json:"MIN_DIST_TO_RUN_WHILE_ATTACK_MOVING_OTHER_ENEMIS"`
	MinDistToStopRun                         float32 `json:"MIN_DIST_TO_STOP_RUN"`
	JumpSpreadDist                           float32 `json:"JUMP_SPREAD_DIST"`
	LookTimesToKill                          int16   `json:"LOOK_TIMES_TO_KILL"`
	ComeInsideTimes                          int16   `json:"COME_INSIDE_TIMES"`
	TotalTimeKill                            float32 `json:"TOTAL_TIME_KILL"`
	TotalTimeKillAfterWarn                   float32 `json:"TOTAL_TIME_KILL_AFTER_WARN"`
	MovingAimCoef                            float32 `json:"MOVING_AIM_COEF"`
	VerticalDistToIgnoreSound                float32 `json:"VERTICAL_DIST_TO_IGNORE_SOUND"`
	DefenceLevelShift                        float32 `json:"DEFENCE_LEVEL_SHIFT"`
	MinDistCloseDef                          float32 `json:"MIN_DIST_CLOSE_DEF"`
	UseIDPriorWhoGo                          bool    `json:"USE_ID_PRIOR_WHO_GO"`
	StartActiveFollowPlayerEvent             bool    `json:"START_ACTIVE_FOLLOW_PLAYER_EVENT"`
	StartActiveForceAttackPlayerEvent        bool    `json:"START_ACTIVE_FORCE_ATTACK_PLAYER_EVENT"`
	SmokeGrenadeRadiusCoef                   float32 `json:"SMOKE_GRENADE_RADIUS_COEF"`
	GrenadePrecision                         int16   `json:"GRENADE_PRECISION"`
	MaxWarnsBeforeKill                       int16   `json:"MAX_WARNS_BEFORE_KILL"`
	CareEnemyOnlyTime                        float32 `json:"CARE_ENEMY_ONLY_TIME"`
	MiddlePointCoef                          float32 `json:"MIDDLE_POINT_COEF"`
	MainTacticOnlyAttack                     bool    `json:"MAIN_TACTIC_ONLY_ATTACK"`
	LastDamageActive                         float32 `json:"LAST_DAMAGE_ACTIVE"`
	ShallDieIfNotInited                      bool    `json:"SHALL_DIE_IF_NOT_INITED"`
	CheckBotInitTimeSec                      float32 `json:"CHECK_BOT_INIT_TIME_SEC"`
	WeaponRootYOffset                        float32 `json:"WEAPON_ROOT_Y_OFFSET"`
	DeltaSupressDistanceSqrt                 float32 `json:"DELTA_SUPRESS_DISTANCE_SQRT"`
	DeltaSupressDistance                     float32 `json:"DELTA_SUPRESS_DISTANCE"`
	WaveCoefLow                              float32 `json:"WAVE_COEF_LOW"`
	WaveCoefMid                              float32 `json:"WAVE_COEF_MID"`
	WaveCoefHigh                             float32 `json:"WAVE_COEF_HIGH"`
	WaveCoefHorde                            float32 `json:"WAVE_COEF_HORDE"`
	WaveOnlyAsOnline                         bool    `json:"WAVE_ONLY_AS_ONLINE"`
	LocalBotsCount                           int16   `json:"LOCAL_BOTS_COUNT"`
	AxeManKillsEnd                           int16   `json:"AXE_MAN_KILLS_END"`
}

type Locales struct {
	Locales   Locale
	Languages map[string]string
}

type Locale struct {
	CH   LocaleData `json:"ch"`
	CZ   LocaleData `json:"cz"`
	EN   LocaleData `json:"en"`
	FR   LocaleData `json:"fr"`
	GE   LocaleData `json:"ge"`
	HU   LocaleData `json:"hu"`
	IT   LocaleData `json:"it"`
	JP   LocaleData `json:"jp"`
	KR   LocaleData `json:"kr"`
	PL   LocaleData `json:"pl"`
	PO   LocaleData `json:"po"`
	SK   LocaleData `json:"sk"`
	ES   LocaleData `json:"es"`
	ESMX LocaleData `json:"es-mx"`
	TU   LocaleData `json:"tu"`
	RU   LocaleData `json:"ru"`
}

type LocaleData struct {
	Locale map[string]string
	Menu   LocaleMenu
}

type LocaleMenu struct {
	Menu map[string]string `json:"menu"`
}

type Handbook struct {
	Categories []HandbookCategory `json:"Categories"`
	Items      []HandbookItem     `json:"Items"`
}

type HandbookCategory struct {
	Id       string `json:"Id"`
	ParentId string `json:"ParentId"`
	Icon     string `json:"Icon"`
	Color    string `json:"Color"`
	Order    string `json:"Order"`
}

type HandbookItem struct {
	Id       string `json:"Id"`
	ParentId string `json:"ParentId"`
	Price    int    `json:"Price"`
}

type Weather struct {
	WeatherInfo struct {
		Timestamp     int64   `json:"timestamp"`
		Cloud         float32 `json:"cloud"`
		WindSpeed     int     `json:"wind_speed"`
		WindDirection int     `json:"wind_direction"`
		WindGustiness float32 `json:"wind_gustiness"`
		Rain          int     `json:"rain"`
		RainIntensity int     `json:"rain_intensity"`
		Fog           float32 `json:"fog"`
		Temperature   int     `json:"temp"`
		Pressure      int     `json:"pressure"`
		Date          string  `json:"date"`
		Time          string  `json:"time"`
	} `json:"weather"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Acceleration int    `json:"acceleration"`
}

type Customization struct {
	ID     string `json:"_id,omitempty"`
	Name   string `json:"_name,omitempty"`
	Parent string `json:"_parent,omitempty"`
	Type   string `json:"_type,omitempty"`
	Proto  string `json:"_proto,omitempty"`
	Props  struct {
		Name                string      `json:"Name,omitempty"`
		ShortName           string      `json:"ShortName,omitempty"`
		Description         string      `json:"Description,omitempty"`
		Side                []string    `json:"Side,omitempty"`
		BodyPart            string      `json:"BodyPart,omitempty"`
		Prefab              interface{} `json:"Prefab,omitempty"`
		WatchPrefab         interface{} `json:"WatchPrefab,omitempty"`
		IntegratedArmorVest bool        `json:"IntegratedArmorVest,omitempty"`
		WatchPosition       XYZ         `json:"WatchPosition,omitempty"`
		WatchRotation       XYZ         `json:"WatchRotation,omitempty"`
		AvailableAsDefault  bool        `json:"AvailableAsDefault,omitempty"`
	} `json:"_props,omitempty"`
}
