package structs

type Bots struct {
	BotTypes      map[string]*BotType
	BotAppearance map[string]*BotAppearance
	BotNames      *BotNames
}

type BotNames struct {
	BossGluhar       []string `json:"bossGluhar,omitempty"`
	BossZryachiy     []string `json:"bossZryachiy,omitempty"`
	FollowerZryachiy []string `json:"followerZryachiy,omitempty"`
	GeneralFollower  []string `json:"generalFollower,omitempty"`
	BossKilla        []string `json:"bossKilla,omitempty"`
	BossBully        []string `json:"bossBully,omitempty"`
	FollowerBully    []string `json:"followerBully,omitempty"`
	BossKojaniy      []string `json:"bossKojaniy,omitempty"`
	FollowerKojaniy  []string `json:"followerKojaniy,omitempty"`
	BossSanitar      []string `json:"bossSanitar,omitempty"`
	FollowerSanitar  []string `json:"followerSanitar,omitempty"`
	BossTagilla      []string `json:"bossTagilla,omitempty"`
	FollowerTagilla  []string `json:"followerTagilla,omitempty"`
	FollowerBigPipe  []string `json:"followerBigPipe,omitempty"`
	FollowerBirdEye  []string `json:"followerBirdEye,omitempty"`
	BossKnight       []string `json:"bossKnight,omitempty"`
	Gifter           []string `json:"gifter,omitempty"`
	Sectantpriest    []string `json:"sectantpriest,omitempty"`
	Sectantwarrior   []string `json:"sectantwarrior,omitempty"`
	Normal           []string `json:"normal,omitempty"`
	Scav             []string `json:"scav,omitempty"`
}

type BotAppearance struct {
	Voice []string
	Body  []string
	Head  []string
	Hands []string
	Feet  []string
}

type BotType struct {
	Difficulties *BotDifficulties `json:"difficulties,omitempty"`
	Health       *BotHealth       `json:"health,omitempty"`
	Loadout      *BotLoadout      `json:"loadout,omitempty"`
}

type BotDifficulties struct {
	Easy       *BotDifficulty `json:"easy,omitempty"`
	Normal     *BotDifficulty `json:"normal,omitempty"`
	Hard       *BotDifficulty `json:"hard,omitempty"`
	Impossible *BotDifficulty `json:"impossible,omitempty"`
}

type BotDifficulty struct {
	Aiming     Aiming     `json:"Aiming,omitempty"`
	Boss       Boss       `json:"Boss,omitempty"`
	Change     Change     `json:"Change,omitempty"`
	Core       BotCore    `json:"Core,omitempty"`
	Cover      Cover      `json:"Cover,omitempty"`
	Grenade    Grenade    `json:"Grenade,omitempty"`
	Hearing    Hearing    `json:"Hearing,omitempty"`
	Lay        Lay        `json:"Lay,omitempty"`
	Look       Look       `json:"Look,omitempty"`
	Mind       Mind       `json:"Mind,omitempty"`
	Move       Move       `json:"Move,omitempty"`
	Patrol     Patrol     `json:"Patrol,omitempty"`
	Scattering Scattering `json:"Scattering,omitempty"`
	Shoot      Shoot      `json:"Shoot,omitempty"`
}
type Aiming struct {
	AimingType                 float32 `json:"AIMING_TYPE,omitempty"`
	AnytimeLightWhenAim100     float32 `json:"ANYTIME_LIGHT_WHEN_AIM_100,omitempty"`
	AnyPartShootTime           float32 `json:"ANY_PART_SHOOT_TIME,omitempty"`
	BaseHitAffectionDelaySec   float32 `json:"BASE_HIT_AFFECTION_DELAY_SEC,omitempty"`
	BaseHitAffectionMaxAng     float32 `json:"BASE_HIT_AFFECTION_MAX_ANG,omitempty"`
	BaseHitAffectionMinAng     float32 `json:"BASE_HIT_AFFECTION_MIN_ANG,omitempty"`
	BaseShief                  float32 `json:"BASE_SHIEF,omitempty"`
	BaseShiefStationaryGrenade float32 `json:"BASE_SHIEF_STATIONARY_GRENADE,omitempty"`
	BetterPrecicingCoef        float32 `json:"BETTER_PRECICING_COEF,omitempty"`
	BottomCoef                 float32 `json:"BOTTOM_COEF,omitempty"`
	BotMoveIfDelta             float32 `json:"BOT_MOVE_IF_DELTA,omitempty"`
	CoefFromCover              float32 `json:"COEF_FROM_COVER,omitempty"`
	CoefIfMove                 float32 `json:"COEF_IF_MOVE,omitempty"`
	DamagePanicTime            float32 `json:"DAMAGE_PANIC_TIME,omitempty"`
	DamageToDiscardAim0100     float32 `json:"DAMAGE_TO_DISCARD_AIM_0_100,omitempty"`
	DangerUpPoint              float32 `json:"DANGER_UP_POINT,omitempty"`
	DistToShootNoOffset        float32 `json:"DIST_TO_SHOOT_NO_OFFSET,omitempty"`
	DistToShootToCenter        float32 `json:"DIST_TO_SHOOT_TO_CENTER,omitempty"`
	FirstContactAddChance100   float32 `json:"FIRST_CONTACT_ADD_CHANCE_100,omitempty"`
	FirstContactAddSec         float32 `json:"FIRST_CONTACT_ADD_SEC,omitempty"`
	HardAim                    float32 `json:"HARD_AIM,omitempty"`
	HardAimChance100           float32 `json:"HARD_AIM_CHANCE_100,omitempty"`
	MaxAimingUpgradeByTime     float32 `json:"MAX_AIMING_UPGRADE_BY_TIME,omitempty"`
	MaxAimPrecicing            float32 `json:"MAX_AIM_PRECICING,omitempty"`
	MaxAimTime                 float32 `json:"MAX_AIM_TIME,omitempty"`
	MaxTimeDiscardAimSec       float32 `json:"MAX_TIME_DISCARD_AIM_SEC,omitempty"`
	MinDamageToGetHitAffets    float32 `json:"MIN_DAMAGE_TO_GET_HIT_AFFETS,omitempty"`
	MinTimeDiscardAimSec       float32 `json:"MIN_TIME_DISCARD_AIM_SEC,omitempty"`
	NextShotMissChance100      float32 `json:"NEXT_SHOT_MISS_CHANCE_100,omitempty"`
	NextShotMissYOffset        float32 `json:"NEXT_SHOT_MISS_Y_OFFSET,omitempty"`
	OffsetRecalAnywayTime      float32 `json:"OFFSET_RECAL_ANYWAY_TIME,omitempty"`
	PanicAccuratyCoef          float32 `json:"PANIC_ACCURATY_COEF,omitempty"`
	PanicCoef                  float32 `json:"PANIC_COEF,omitempty"`
	PanicTime                  float32 `json:"PANIC_TIME,omitempty"`
	RecalcDist                 float32 `json:"RECALC_DIST,omitempty"`
	RecalcMustTime             float32 `json:"RECALC_MUST_TIME,omitempty"`
	RecalcSqrDist              float32 `json:"RECALC_SQR_DIST,omitempty"`
	ScatteringDistModif        float32 `json:"SCATTERING_DIST_MODIF,omitempty"`
	ScatteringDistModifClose   float32 `json:"SCATTERING_DIST_MODIF_CLOSE,omitempty"`
	ScatteringHaveDamageCoef   float32 `json:"SCATTERING_HAVE_DAMAGE_COEF,omitempty"`
	ShootToChangePriority      float32 `json:"SHOOT_TO_CHANGE_PRIORITY,omitempty"`
	ShpereFriendyFireSize      float32 `json:"SHPERE_FRIENDY_FIRE_SIZE,omitempty"`
	TimeCoefIfMove             float32 `json:"TIME_COEF_IF_MOVE,omitempty"`
	WeaponRootOffset           float32 `json:"WEAPON_ROOT_OFFSET,omitempty"`
	XzCoef                     float32 `json:"XZ_COEF,omitempty"`
	XzCoefStationaryGrenade    float32 `json:"XZ_COEF_STATIONARY_GRENADE,omitempty"`
	YBottomOffsetCoef          float32 `json:"Y_BOTTOM_OFFSET_COEF,omitempty"`
	YTopOffsetCoef             float32 `json:"Y_TOP_OFFSET_COEF,omitempty"`
}
type Boss struct {
	BossDistToShoot                float32 `json:"BOSS_DIST_TO_SHOOT,omitempty"`
	BossDistToShootSqrt            float32 `json:"BOSS_DIST_TO_SHOOT_SQRT,omitempty"`
	BossDistToWarning              float32 `json:"BOSS_DIST_TO_WARNING,omitempty"`
	BossDistToWarningOut           float32 `json:"BOSS_DIST_TO_WARNING_OUT,omitempty"`
	BossDistToWarningOutSqrt       float32 `json:"BOSS_DIST_TO_WARNING_OUT_SQRT,omitempty"`
	BossDistToWarningSqrt          float32 `json:"BOSS_DIST_TO_WARNING_SQRT,omitempty"`
	ChanceToSendGrenade100         float32 `json:"CHANCE_TO_SEND_GRENADE_100,omitempty"`
	ChanceUseReservePatrol100      float32 `json:"CHANCE_USE_RESERVE_PATROL_100,omitempty"`
	CoverToSend                    bool    `json:"COVER_TO_SEND,omitempty"`
	DeltaSearchTime                float32 `json:"DELTA_SEARCH_TIME,omitempty"`
	KillaAfterGrenadeSuppressDelay float32 `json:"KILLA_AFTER_GRENADE_SUPPRESS_DELAY,omitempty"`
	KillaBulletToReload            float32 `json:"KILLA_BULLET_TO_RELOAD,omitempty"`
	KillaCloseattackDelay          float32 `json:"KILLA_CLOSEATTACK_DELAY,omitempty"`
	KillaCloseattackTimes          float32 `json:"KILLA_CLOSEATTACK_TIMES,omitempty"`
	KillaCloseAttackDist           float32 `json:"KILLA_CLOSE_ATTACK_DIST,omitempty"`
	KillaContutionTime             float32 `json:"KILLA_CONTUTION_TIME,omitempty"`
	KillaDefDistSqrt               float32 `json:"KILLA_DEF_DIST_SQRT,omitempty"`
	KillaDistToGoToSuppress        float32 `json:"KILLA_DIST_TO_GO_TO_SUPPRESS,omitempty"`
	KillaDitanceToBeEnemyBoss      float32 `json:"KILLA_DITANCE_TO_BE_ENEMY_BOSS,omitempty"`
	KillaEnemiesToAttack           float32 `json:"KILLA_ENEMIES_TO_ATTACK,omitempty"`
	KillaHoldDelay                 float32 `json:"KILLA_HOLD_DELAY,omitempty"`
	KillaLargeAttackDist           float32 `json:"KILLA_LARGE_ATTACK_DIST,omitempty"`
	KillaMiddleAttackDist          float32 `json:"KILLA_MIDDLE_ATTACK_DIST,omitempty"`
	KillaOneIsClose                float32 `json:"KILLA_ONE_IS_CLOSE,omitempty"`
	KillaSearchMeters              float32 `json:"KILLA_SEARCH_METERS,omitempty"`
	KillaSearchSecStopAfterComing  float32 `json:"KILLA_SEARCH_SEC_STOP_AFTER_COMING,omitempty"`
	KillaStartSearchSec            float32 `json:"KILLA_START_SEARCH_SEC,omitempty"`
	KillaTriggerDownDelay          float32 `json:"KILLA_TRIGGER_DOWN_DELAY,omitempty"`
	KillaWaitInCoverCoef           float32 `json:"KILLA_WAIT_IN_COVER_COEF,omitempty"`
	KillaYDeltaToBeEnemyBoss       float32 `json:"KILLA_Y_DELTA_TO_BE_ENEMY_BOSS,omitempty"`
	MaxDistCoverBoss               float32 `json:"MAX_DIST_COVER_BOSS,omitempty"`
	MaxDistCoverBossSqrt           float32 `json:"MAX_DIST_COVER_BOSS_SQRT,omitempty"`
	MaxDistDeciderToSend           float32 `json:"MAX_DIST_DECIDER_TO_SEND,omitempty"`
	MaxDistDeciderToSendSqrt       float32 `json:"MAX_DIST_DECIDER_TO_SEND_SQRT,omitempty"`
	PersonsSend                    float32 `json:"PERSONS_SEND,omitempty"`
	ShallWarn                      bool    `json:"SHALL_WARN,omitempty"`
	TimeAfterLose                  float32 `json:"TIME_AFTER_LOSE,omitempty"`
	TimeAfterLoseDelta             float32 `json:"TIME_AFTER_LOSE_DELTA,omitempty"`
	WaitNoAttackSavage             float32 `json:"WAIT_NO_ATTACK_SAVAGE,omitempty"`
}
type Change struct {
	FlashAccuraty   float32 `json:"FLASH_ACCURATY,omitempty"`
	FlashGainSight  float32 `json:"FLASH_GAIN_SIGHT,omitempty"`
	FlashHearing    float32 `json:"FLASH_HEARING,omitempty"`
	FlashLayChance  float32 `json:"FLASH_LAY_CHANCE,omitempty"`
	FlashPrecicing  float32 `json:"FLASH_PRECICING,omitempty"`
	FlashScattering float32 `json:"FLASH_SCATTERING,omitempty"`
	FlashVisionDist float32 `json:"FLASH_VISION_DIST,omitempty"`
	SmokeAccuraty   float32 `json:"SMOKE_ACCURATY,omitempty"`
	SmokeGainSight  float32 `json:"SMOKE_GAIN_SIGHT,omitempty"`
	SmokeHearing    float32 `json:"SMOKE_HEARING,omitempty"`
	SmokeLayChance  float32 `json:"SMOKE_LAY_CHANCE,omitempty"`
	SmokePrecicing  float32 `json:"SMOKE_PRECICING,omitempty"`
	SmokeScattering float32 `json:"SMOKE_SCATTERING,omitempty"`
	SmokeVisionDist float32 `json:"SMOKE_VISION_DIST,omitempty"`
	StunHearing     float32 `json:"STUN_HEARING,omitempty"`
}
type BotCore struct {
	AccuratySpeed              float32 `json:"AccuratySpeed,omitempty"`
	AimingType                 string  `json:"AimingType,omitempty"`
	CanGrenade                 bool    `json:"CanGrenade,omitempty"`
	CanRun                     bool    `json:"CanRun,omitempty"`
	DamageCoeff                float32 `json:"DamageCoeff,omitempty"`
	GainSightCoef              float32 `json:"GainSightCoef,omitempty"`
	HearingSense               float32 `json:"HearingSense,omitempty"`
	PistolFireDistancePref     float32 `json:"PistolFireDistancePref,omitempty"`
	RifleFireDistancePref      float32 `json:"RifleFireDistancePref,omitempty"`
	ScatteringClosePerMeter    float32 `json:"ScatteringClosePerMeter,omitempty"`
	ScatteringPerMeter         float32 `json:"ScatteringPerMeter,omitempty"`
	ShotgunFireDistancePref    float32 `json:"ShotgunFireDistancePref,omitempty"`
	VisibleAngle               float32 `json:"VisibleAngle,omitempty"`
	VisibleDistance            float32 `json:"VisibleDistance,omitempty"`
	WaitInCoverBetweenShotsSec float32 `json:"WaitInCoverBetweenShotsSec,omitempty"`
}
type Cover struct {
	ChangeRunToCoverSec          float32 `json:"CHANGE_RUN_TO_COVER_SEC,omitempty"`
	ChangeRunToCoverSecGreande   float32 `json:"CHANGE_RUN_TO_COVER_SEC_GREANDE,omitempty"`
	CheckCoverEnemyLook          bool    `json:"CHECK_COVER_ENEMY_LOOK,omitempty"`
	CloseDistPointSqrt           float32 `json:"CLOSE_DIST_POINT_SQRT,omitempty"`
	DeltaSeenFromCoveLastPos     float32 `json:"DELTA_SEEN_FROM_COVE_LAST_POS,omitempty"`
	DependsYDistToBot            bool    `json:"DEPENDS_Y_DIST_TO_BOT,omitempty"`
	DistCantChangeWay            float32 `json:"DIST_CANT_CHANGE_WAY,omitempty"`
	DistCantChangeWaySqr         float32 `json:"DIST_CANT_CHANGE_WAY_SQR,omitempty"`
	DistCheckSfety               float32 `json:"DIST_CHECK_SFETY,omitempty"`
	DogFightAfterLeave           float32 `json:"DOG_FIGHT_AFTER_LEAVE,omitempty"`
	EnemyDistToGoOut             float32 `json:"ENEMY_DIST_TO_GO_OUT,omitempty"`
	GoodDistToPointCoef          float32 `json:"GOOD_DIST_TO_POINT_COEF,omitempty"`
	HideToCoverTime              float32 `json:"HIDE_TO_COVER_TIME,omitempty"`
	HitsToLeaveCover             float32 `json:"HITS_TO_LEAVE_COVER,omitempty"`
	HitsToLeaveCoverUnknown      float32 `json:"HITS_TO_LEAVE_COVER_UNKNOWN,omitempty"`
	LookLastEnemyPosLookaround   float32 `json:"LOOK_LAST_ENEMY_POS_LOOKAROUND,omitempty"`
	LookLastEnemyPosMoving       float32 `json:"LOOK_LAST_ENEMY_POS_MOVING,omitempty"`
	LookToHitPointIfLastEnemy    float32 `json:"LOOK_TO_HIT_POINT_IF_LAST_ENEMY,omitempty"`
	MaxDistOfCover               float32 `json:"MAX_DIST_OF_COVER,omitempty"`
	MaxDistOfCoverSqr            float32 `json:"MAX_DIST_OF_COVER_SQR,omitempty"`
	MaxSpottedTimeSec            float32 `json:"MAX_SPOTTED_TIME_SEC,omitempty"`
	MinDefenceLevel              float32 `json:"MIN_DEFENCE_LEVEL,omitempty"`
	MinDistToEnemy               float32 `json:"MIN_DIST_TO_ENEMY,omitempty"`
	MoveToCoverWhenTarget        bool    `json:"MOVE_TO_COVER_WHEN_TARGET,omitempty"`
	NotLookAtWallIsDanger        bool    `json:"NOT_LOOK_AT_WALL_IS_DANGER,omitempty"`
	OffsetLookAlongWallAng       float32 `json:"OFFSET_LOOK_ALONG_WALL_ANG,omitempty"`
	ReturnToAttackAfterAmbushMax float32 `json:"RETURN_TO_ATTACK_AFTER_AMBUSH_MAX,omitempty"`
	ReturnToAttackAfterAmbushMin float32 `json:"RETURN_TO_ATTACK_AFTER_AMBUSH_MIN,omitempty"`
	RunCoverIfCanAndNoEnemies    bool    `json:"RUN_COVER_IF_CAN_AND_NO_ENEMIES,omitempty"`
	RunIfFar                     float32 `json:"RUN_IF_FAR,omitempty"`
	RunIfFarSqrt                 float32 `json:"RUN_IF_FAR_SQRT,omitempty"`
	ShootNearSecPeriod           float32 `json:"SHOOT_NEAR_SEC_PERIOD,omitempty"`
	ShootNearToLeave             float32 `json:"SHOOT_NEAR_TO_LEAVE,omitempty"`
	SoundToGetSpotted            float32 `json:"SOUND_TO_GET_SPOTTED,omitempty"`
	SpottedCoversRadius          float32 `json:"SPOTTED_COVERS_RADIUS,omitempty"`
	SpottedGrenadeRadius         float32 `json:"SPOTTED_GRENADE_RADIUS,omitempty"`
	SpottedGrenadeTime           float32 `json:"SPOTTED_GRENADE_TIME,omitempty"`
	StationaryWeaponMaxDistToUse float32 `json:"STATIONARY_WEAPON_MAX_DIST_TO_USE,omitempty"`
	StationaryWeaponNoEnemyGetup float32 `json:"STATIONARY_WEAPON_NO_ENEMY_GETUP,omitempty"`
	StayIfFar                    float32 `json:"STAY_IF_FAR,omitempty"`
	StayIfFarSqrt                float32 `json:"STAY_IF_FAR_SQRT,omitempty"`
	TimeCheckSafe                float32 `json:"TIME_CHECK_SAFE,omitempty"`
	TimeToMoveToCover            float32 `json:"TIME_TO_MOVE_TO_COVER,omitempty"`
	WaitIntCoverFindingEnemy     float32 `json:"WAIT_INT_COVER_FINDING_ENEMY,omitempty"`
}
type Grenade struct {
	AddGrenadeAsDanger             float32 `json:"ADD_GRENADE_AS_DANGER,omitempty"`
	AddGrenadeAsDangerSqr          float32 `json:"ADD_GRENADE_AS_DANGER_SQR,omitempty"`
	AmbushIfSmokeInZone100         float32 `json:"AMBUSH_IF_SMOKE_IN_ZONE_100,omitempty"`
	AmbushIfSmokeReturnToAttackSec float32 `json:"AMBUSH_IF_SMOKE_RETURN_TO_ATTACK_SEC,omitempty"`
	AngType                        float32 `json:"ANG_TYPE,omitempty"`
	BewareType                     float32 `json:"BEWARE_TYPE,omitempty"`
	BeAttentionCoef                float32 `json:"BE_ATTENTION_COEF,omitempty"`
	CanThrowStraightContact        bool    `json:"CAN_THROW_STRAIGHT_CONTACT,omitempty"`
	ChanceRunFlashed100            float32 `json:"CHANCE_RUN_FLASHED_100,omitempty"`
	ChanceToNotifyEnemyGr100       float32 `json:"CHANCE_TO_NOTIFY_ENEMY_GR_100,omitempty"`
	CheatStartGrenadePlace         bool    `json:"CHEAT_START_GRENADE_PLACE,omitempty"`
	CloseToSmokeTimeDelta          float32 `json:"CLOSE_TO_SMOKE_TIME_DELTA,omitempty"`
	CloseToSmokeToShoot            float32 `json:"CLOSE_TO_SMOKE_TO_SHOOT,omitempty"`
	CloseToSmokeToShootSqrt        float32 `json:"CLOSE_TO_SMOKE_TO_SHOOT_SQRT,omitempty"`
	DamageGrenadeSuppressDelta     float32 `json:"DAMAGE_GRENADE_SUPPRESS_DELTA,omitempty"`
	DeltaGrenadeStartTime          float32 `json:"DELTA_GRENADE_START_TIME,omitempty"`
	DeltaNextAttempt               float32 `json:"DELTA_NEXT_ATTEMPT,omitempty"`
	DeltaNextAttemptFromCover      float32 `json:"DELTA_NEXT_ATTEMPT_FROM_COVER,omitempty"`
	FlashGrenadeTimeCoef           float32 `json:"FLASH_GRENADE_TIME_COEF,omitempty"`
	GrenadePerMeter                float32 `json:"GrenadePerMeter,omitempty"`
	GrenadePrecision               float32 `json:"GrenadePrecision,omitempty"`
	MaxFlashedDistToShoot          float32 `json:"MAX_FLASHED_DIST_TO_SHOOT,omitempty"`
	MaxFlashedDistToShootSqrt      float32 `json:"MAX_FLASHED_DIST_TO_SHOOT_SQRT,omitempty"`
	MaxThrowPower                  float32 `json:"MAX_THROW_POWER,omitempty"`
	MinDistNotToThrow              float32 `json:"MIN_DIST_NOT_TO_THROW,omitempty"`
	MinDistNotToThrowSqr           float32 `json:"MIN_DIST_NOT_TO_THROW_SQR,omitempty"`
	MinThrowGrenadeDist            float32 `json:"MIN_THROW_GRENADE_DIST,omitempty"`
	MinThrowGrenadeDistSqrt        float32 `json:"MIN_THROW_GRENADE_DIST_SQRT,omitempty"`
	NearDeltaThrowTimeSec          float32 `json:"NEAR_DELTA_THROW_TIME_SEC,omitempty"`
	NoRunFromAiGrenades            bool    `json:"NO_RUN_FROM_AI_GRENADES,omitempty"`
	RequestDistMustThrow           float32 `json:"REQUEST_DIST_MUST_THROW,omitempty"`
	RequestDistMustThrowSqrt       float32 `json:"REQUEST_DIST_MUST_THROW_SQRT,omitempty"`
	RunAway                        float32 `json:"RUN_AWAY,omitempty"`
	RunAwaySqr                     float32 `json:"RUN_AWAY_SQR,omitempty"`
	ShootToSmokeChance100          float32 `json:"SHOOT_TO_SMOKE_CHANCE_100,omitempty"`
	SizeSpottedCoef                float32 `json:"SIZE_SPOTTED_COEF,omitempty"`
	SmokeCheckDelta                float32 `json:"SMOKE_CHECK_DELTA,omitempty"`
	SmokeSuppressDelta             float32 `json:"SMOKE_SUPPRESS_DELTA,omitempty"`
	StopWhenThrowGrenade           bool    `json:"STOP_WHEN_THROW_GRENADE,omitempty"`
	StraightContactDeltaSec        float32 `json:"STRAIGHT_CONTACT_DELTA_SEC,omitempty"`
	StunSuppressDelta              float32 `json:"STUN_SUPPRESS_DELTA,omitempty"`
	TimeShootToFlash               float32 `json:"TIME_SHOOT_TO_FLASH,omitempty"`
	WaitTimeTurnAway               float32 `json:"WAIT_TIME_TURN_AWAY,omitempty"`
}
type Hearing struct {
	BotClosePanicDist         float32 `json:"BOT_CLOSE_PANIC_DIST,omitempty"`
	ChanceToHearSimpleSound01 float32 `json:"CHANCE_TO_HEAR_SIMPLE_SOUND_0_1,omitempty"`
	CloseDist                 float32 `json:"CLOSE_DIST,omitempty"`
	DeadBodySoundRad          float32 `json:"DEAD_BODY_SOUND_RAD,omitempty"`
	DispersionCoef            float32 `json:"DISPERSION_COEF,omitempty"`
	DistPlaceToFindPoint      float32 `json:"DIST_PLACE_TO_FIND_POINT,omitempty"`
	FarDist                   float32 `json:"FAR_DIST,omitempty"`
	HearDelayWhenHaveSmt      float32 `json:"HEAR_DELAY_WHEN_HAVE_SMT,omitempty"`
	HearDelayWhenPeace        float32 `json:"HEAR_DELAY_WHEN_PEACE,omitempty"`
	LookOnlyDanger            bool    `json:"LOOK_ONLY_DANGER,omitempty"`
	LookOnlyDangerDelta       float32 `json:"LOOK_ONLY_DANGER_DELTA,omitempty"`
	ResetTimerDist            float32 `json:"RESET_TIMER_DIST,omitempty"`
	SoundDirDeefree           float32 `json:"SOUND_DIR_DEEFREE,omitempty"`
}
type Lay struct {
	AttackLayChance            float32 `json:"ATTACK_LAY_CHANCE,omitempty"`
	CheckShootWhenLaying       bool    `json:"CHECK_SHOOT_WHEN_LAYING,omitempty"`
	ClearPointsOfScareSec      float32 `json:"CLEAR_POINTS_OF_SCARE_SEC,omitempty"`
	DamageTimeToGetup          float32 `json:"DAMAGE_TIME_TO_GETUP,omitempty"`
	DeltaAfterGetup            float32 `json:"DELTA_AFTER_GETUP,omitempty"`
	DeltaGetup                 float32 `json:"DELTA_GETUP,omitempty"`
	DeltaLayCheck              float32 `json:"DELTA_LAY_CHECK,omitempty"`
	DeltaWantLayCheclSec       float32 `json:"DELTA_WANT_LAY_CHECL_SEC,omitempty"`
	DistEnemyCanLay            float32 `json:"DIST_ENEMY_CAN_LAY,omitempty"`
	DistEnemyCanLaySqrt        float32 `json:"DIST_ENEMY_CAN_LAY_SQRT,omitempty"`
	DistEnemyGetupLay          float32 `json:"DIST_ENEMY_GETUP_LAY,omitempty"`
	DistEnemyGetupLaySqrt      float32 `json:"DIST_ENEMY_GETUP_LAY_SQRT,omitempty"`
	DistEnemyNullDangerLay     float32 `json:"DIST_ENEMY_NULL_DANGER_LAY,omitempty"`
	DistEnemyNullDangerLaySqrt float32 `json:"DIST_ENEMY_NULL_DANGER_LAY_SQRT,omitempty"`
	DistGrassTerrainSqrt       float32 `json:"DIST_GRASS_TERRAIN_SQRT,omitempty"`
	DistToCoverToLay           float32 `json:"DIST_TO_COVER_TO_LAY,omitempty"`
	DistToCoverToLaySqrt       float32 `json:"DIST_TO_COVER_TO_LAY_SQRT,omitempty"`
	LayAim                     float32 `json:"LAY_AIM,omitempty"`
	LayChanceDanger            float32 `json:"LAY_CHANCE_DANGER,omitempty"`
	MaxCanLayDist              float32 `json:"MAX_CAN_LAY_DIST,omitempty"`
	MaxCanLayDistSqrt          float32 `json:"MAX_CAN_LAY_DIST_SQRT,omitempty"`
	MaxLayTime                 float32 `json:"MAX_LAY_TIME,omitempty"`
	MinCanLayDist              float32 `json:"MIN_CAN_LAY_DIST,omitempty"`
	MinCanLayDistSqrt          float32 `json:"MIN_CAN_LAY_DIST_SQRT,omitempty"`
}
type Look struct {
	BodyDeltaTimeSearchSec        float32 `json:"BODY_DELTA_TIME_SEARCH_SEC,omitempty"`
	CanLookToWall                 bool    `json:"CAN_LOOK_TO_WALL,omitempty"`
	ComeToBodyDist                float32 `json:"COME_TO_BODY_DIST,omitempty"`
	CloseDeltaTimeSec             float32 `json:"CloseDeltaTimeSec,omitempty"`
	DistCheckWall                 float32 `json:"DIST_CHECK_WALL,omitempty"`
	DistNotToIgnoreWall           float32 `json:"DIST_NOT_TO_IGNORE_WALL,omitempty"`
	EnemyLightAdd                 float32 `json:"ENEMY_LIGHT_ADD,omitempty"`
	EnemyLightStartDist           float32 `json:"ENEMY_LIGHT_START_DIST,omitempty"`
	FarDistance                   float32 `json:"FAR_DISTANCE,omitempty"`
	FarDeltaTimeSec               float32 `json:"FarDeltaTimeSec,omitempty"`
	GoalToFullDissapear           float32 `json:"GOAL_TO_FULL_DISSAPEAR,omitempty"`
	GoalToFullDissapearShoot      float32 `json:"GOAL_TO_FULL_DISSAPEAR_SHOOT,omitempty"`
	LookAroundDelta               float32 `json:"LOOK_AROUND_DELTA,omitempty"`
	LookLastPosenemyIfNoDangerSec float32 `json:"LOOK_LAST_POSENEMY_IF_NO_DANGER_SEC,omitempty"`
	LightOnVisionDistance         float32 `json:"LightOnVisionDistance,omitempty"`
	MarksmanVisibleDistCoef       float32 `json:"MARKSMAN_VISIBLE_DIST_COEF,omitempty"`
	MaxVisionGrassMeters          float32 `json:"MAX_VISION_GRASS_METERS,omitempty"`
	MaxVisionGrassMetersFlare     float32 `json:"MAX_VISION_GRASS_METERS_FLARE,omitempty"`
	MaxVisionGrassMetersFlareOpt  float32 `json:"MAX_VISION_GRASS_METERS_FLARE_OPT,omitempty"`
	MaxVisionGrassMetersOpt       float32 `json:"MAX_VISION_GRASS_METERS_OPT,omitempty"`
	MiddleDist                    float32 `json:"MIDDLE_DIST,omitempty"`
	MinLookAroudTime              float32 `json:"MIN_LOOK_AROUD_TIME,omitempty"`
	MiddleDeltaTimeSec            float32 `json:"MiddleDeltaTimeSec,omitempty"`
	OldTimePoint                  float32 `json:"OLD_TIME_POINT,omitempty"`
	OptimizeToOnlyBody            bool    `json:"OPTIMIZE_TO_ONLY_BODY,omitempty"`
	PosibleVisionSpace            float32 `json:"POSIBLE_VISION_SPACE,omitempty"`
	VisibleDisnaceWithLight       float32 `json:"VISIBLE_DISNACE_WITH_LIGHT,omitempty"`
	WaitNewSensor                 float32 `json:"WAIT_NEW_SENSOR,omitempty"`
	WaitNewLookSensor             float32 `json:"WAIT_NEW__LOOK_SENSOR,omitempty"`
}
type Mind struct {
	AiPowerCoef                           float32 `json:"AI_POWER_COEF,omitempty"`
	AmbushWhenUnderFire                   bool    `json:"AMBUSH_WHEN_UNDER_FIRE,omitempty"`
	AmbushWhenUnderFireTimeResist         float32 `json:"AMBUSH_WHEN_UNDER_FIRE_TIME_RESIST,omitempty"`
	AttackEnemyIfProtectDeltaLastTimeSeen float32 `json:"ATTACK_ENEMY_IF_PROTECT_DELTA_LAST_TIME_SEEN,omitempty"`
	AttackImmediatlyChance0100            float32 `json:"ATTACK_IMMEDIATLY_CHANCE_0_100,omitempty"`
	BulletFeelCloseSdist                  float32 `json:"BULLET_FEEL_CLOSE_SDIST,omitempty"`
	BulletFeelDist                        float32 `json:"BULLET_FEEL_DIST,omitempty"`
	CanPanicIsProtect                     bool    `json:"CAN_PANIC_IS_PROTECT,omitempty"`
	CanReceivePlayerRequestsBear          bool    `json:"CAN_RECEIVE_PLAYER_REQUESTS_BEAR,omitempty"`
	CanReceivePlayerRequestsSavage        bool    `json:"CAN_RECEIVE_PLAYER_REQUESTS_SAVAGE,omitempty"`
	CanReceivePlayerRequestsUsec          bool    `json:"CAN_RECEIVE_PLAYER_REQUESTS_USEC,omitempty"`
	CanStandBy                            bool    `json:"CAN_STAND_BY,omitempty"`
	CanTakeItems                          bool    `json:"CAN_TAKE_ITEMS,omitempty"`
	CanThrowRequests                      bool    `json:"CAN_THROW_REQUESTS,omitempty"`
	CanUseMeds                            bool    `json:"CAN_USE_MEDS,omitempty"`
	ChanceFuckYouOnContact100             float32 `json:"CHANCE_FUCK_YOU_ON_CONTACT_100,omitempty"`
	ChanceShootWhenWarnPlayer100          float32 `json:"CHANCE_SHOOT_WHEN_WARN_PLAYER_100,omitempty"`
	ChanceToRunCauseDamage0100            float32 `json:"CHANCE_TO_RUN_CAUSE_DAMAGE_0_100,omitempty"`
	ChanceToStayWhenWarnPlayer100         float32 `json:"CHANCE_TO_STAY_WHEN_WARN_PLAYER_100,omitempty"`
	CoverDistCoef                         float32 `json:"COVER_DIST_COEF,omitempty"`
	CoverSecondsAfterLoseVision           float32 `json:"COVER_SECONDS_AFTER_LOSE_VISION,omitempty"`
	CoverSelfAlwaysIfDamaged              bool    `json:"COVER_SELF_ALWAYS_IF_DAMAGED,omitempty"`
	DamageReductionTimeSec                float32 `json:"DAMAGE_REDUCTION_TIME_SEC,omitempty"`
	DangerPointChooseCoef                 float32 `json:"DANGER_POINT_CHOOSE_COEF,omitempty"`
	DistToEnemyYoCanHeal                  float32 `json:"DIST_TO_ENEMY_YO_CAN_HEAL,omitempty"`
	DistToFoundSqrt                       float32 `json:"DIST_TO_FOUND_SQRT,omitempty"`
	DistToStopRunEnemy                    float32 `json:"DIST_TO_STOP_RUN_ENEMY,omitempty"`
	DogFightIn                            float32 `json:"DOG_FIGHT_IN,omitempty"`
	DogFightOut                           float32 `json:"DOG_FIGHT_OUT,omitempty"`
	EnemyLookAtMeAng                      float32 `json:"ENEMY_LOOK_AT_ME_ANG,omitempty"`
	FindCoverToGetPositionWithShoot       float32 `json:"FIND_COVER_TO_GET_POSITION_WITH_SHOOT,omitempty"`
	FriendAgrKill                         float32 `json:"FRIEND_AGR_KILL,omitempty"`
	FriendDeadAgrLow                      float32 `json:"FRIEND_DEAD_AGR_LOW,omitempty"`
	GroupAnyPhraseDelay                   float32 `json:"GROUP_ANY_PHRASE_DELAY,omitempty"`
	GroupExactlyPhraseDelay               float32 `json:"GROUP_EXACTLY_PHRASE_DELAY,omitempty"`
	HealDelaySec                          float32 `json:"HEAL_DELAY_SEC,omitempty"`
	HitDelayWhenHaveSmt                   float32 `json:"HIT_DELAY_WHEN_HAVE_SMT,omitempty"`
	HitDelayWhenPeace                     float32 `json:"HIT_DELAY_WHEN_PEACE,omitempty"`
	HitPointDetection                     float32 `json:"HIT_POINT_DETECTION,omitempty"`
	HoldIfProtectDeltaLastTimeSeen        float32 `json:"HOLD_IF_PROTECT_DELTA_LAST_TIME_SEEN,omitempty"`
	HowWorkOverDeadBody                   float32 `json:"HOW_WORK_OVER_DEAD_BODY,omitempty"`
	LastseenPointChooseCoef               float32 `json:"LASTSEEN_POINT_CHOOSE_COEF,omitempty"`
	LastEnemyLookTo                       float32 `json:"LAST_ENEMY_LOOK_TO,omitempty"`
	MaxAggroBotDist                       float32 `json:"MAX_AGGRO_BOT_DIST,omitempty"`
	MaxAggroBotDistSqr                    float32 `json:"MAX_AGGRO_BOT_DIST_SQR,omitempty"`
	MaxShootsTime                         float32 `json:"MAX_SHOOTS_TIME,omitempty"`
	MaxStartAggresionCoef                 float32 `json:"MAX_START_AGGRESION_COEF,omitempty"`
	MinDamageScare                        float32 `json:"MIN_DAMAGE_SCARE,omitempty"`
	MinShootsTime                         float32 `json:"MIN_SHOOTS_TIME,omitempty"`
	MinStartAggresionCoef                 float32 `json:"MIN_START_AGGRESION_COEF,omitempty"`
	NoRunAwayForSafe                      bool    `json:"NO_RUN_AWAY_FOR_SAFE,omitempty"`
	PartPercentToHeal                     float32 `json:"PART_PERCENT_TO_HEAL,omitempty"`
	PistolShotgunAmbushDist               float32 `json:"PISTOL_SHOTGUN_AMBUSH_DIST,omitempty"`
	ProtectDeltaHealSec                   float32 `json:"PROTECT_DELTA_HEAL_SEC,omitempty"`
	ProtectTimeReal                       bool    `json:"PROTECT_TIME_REAL,omitempty"`
	SecToMoreDistToRun                    float32 `json:"SEC_TO_MORE_DIST_TO_RUN,omitempty"`
	ShootInsteadDogFight                  float32 `json:"SHOOT_INSTEAD_DOG_FIGHT,omitempty"`
	SimplePointChooseCoef                 float32 `json:"SIMPLE_POINT_CHOOSE_COEF,omitempty"`
	StandartAmbushDist                    float32 `json:"STANDART_AMBUSH_DIST,omitempty"`
	SuspetionPointChanceAdd100            float32 `json:"SUSPETION_POINT_CHANCE_ADD100,omitempty"`
	TalkWithQuery                         bool    `json:"TALK_WITH_QUERY,omitempty"`
	TimeLeaveMap                          float32 `json:"TIME_LEAVE_MAP,omitempty"`
	TimeToFindEnemy                       float32 `json:"TIME_TO_FIND_ENEMY,omitempty"`
	TimeToForgorAboutEnemySec             float32 `json:"TIME_TO_FORGOR_ABOUT_ENEMY_SEC,omitempty"`
	TimeToRunToCoverCauseShootSec         float32 `json:"TIME_TO_RUN_TO_COVER_CAUSE_SHOOT_SEC,omitempty"`
	WillPersueAxeman                      bool    `json:"WILL_PERSUE_AXEMAN,omitempty"`
}
type Move struct {
	BasestartSlowDist       float32 `json:"BASESTART_SLOW_DIST,omitempty"`
	BaseRotateSpeed         float32 `json:"BASE_ROTATE_SPEED,omitempty"`
	BaseSqrtStartSerach     float32 `json:"BASE_SQRT_START_SERACH,omitempty"`
	BaseStartSerach         float32 `json:"BASE_START_SERACH,omitempty"`
	ChanceToRunIfNoAmmo0100 float32 `json:"CHANCE_TO_RUN_IF_NO_AMMO_0_100,omitempty"`
	DeltaLastSeenEnemy      float32 `json:"DELTA_LAST_SEEN_ENEMY,omitempty"`
	DistToCanChangeWay      float32 `json:"DIST_TO_CAN_CHANGE_WAY,omitempty"`
	DistToCanChangeWaySqr   float32 `json:"DIST_TO_CAN_CHANGE_WAY_SQR,omitempty"`
	DistToStartRaycast      float32 `json:"DIST_TO_START_RAYCAST,omitempty"`
	DistToStartRaycastSqr   float32 `json:"DIST_TO_START_RAYCAST_SQR,omitempty"`
	FarDist                 float32 `json:"FAR_DIST,omitempty"`
	FarDistSqr              float32 `json:"FAR_DIST_SQR,omitempty"`
	ReachDist               float32 `json:"REACH_DIST,omitempty"`
	ReachDistCover          float32 `json:"REACH_DIST_COVER,omitempty"`
	ReachDistRun            float32 `json:"REACH_DIST_RUN,omitempty"`
	RunIfCantShoot          bool    `json:"RUN_IF_CANT_SHOOT,omitempty"`
	RunIfGaolFarThen        float32 `json:"RUN_IF_GAOL_FAR_THEN,omitempty"`
	RunToCoverMin           float32 `json:"RUN_TO_COVER_MIN,omitempty"`
	SecToChangeToRun        float32 `json:"SEC_TO_CHANGE_TO_RUN,omitempty"`
	SlowCoef                float32 `json:"SLOW_COEF,omitempty"`
	StartSlowDist           float32 `json:"START_SLOW_DIST,omitempty"`
	UpdateTimeRecalWay      float32 `json:"UPDATE_TIME_RECAL_WAY,omitempty"`
	YApproximation          float32 `json:"Y_APPROXIMATION,omitempty"`
}
type Patrol struct {
	CanChooseReserv                       bool    `json:"CAN_CHOOSE_RESERV,omitempty"`
	CanFriendlyTilt                       bool    `json:"CAN_FRIENDLY_TILT,omitempty"`
	CanHardAim                            bool    `json:"CAN_HARD_AIM,omitempty"`
	CanLookToDeadbodies                   bool    `json:"CAN_LOOK_TO_DEADBODIES,omitempty"`
	CanWatchSecondWeapon                  bool    `json:"CAN_WATCH_SECOND_WEAPON,omitempty"`
	ChanceToChangeWay0100                 float32 `json:"CHANCE_TO_CHANGE_WAY_0_100,omitempty"`
	ChanceToCutWay0100                    float32 `json:"CHANCE_TO_CUT_WAY_0_100,omitempty"`
	ChanceToShootDeadbody                 float32 `json:"CHANCE_TO_SHOOT_DEADBODY,omitempty"`
	ChangeWayTime                         float32 `json:"CHANGE_WAY_TIME,omitempty"`
	CloseToSelectReservWay                float32 `json:"CLOSE_TO_SELECT_RESERV_WAY,omitempty"`
	CutWayMax01                           float32 `json:"CUT_WAY_MAX_0_1,omitempty"`
	CutWayMin01                           float32 `json:"CUT_WAY_MIN_0_1,omitempty"`
	DeadBodyLookPeriod                    float32 `json:"DEAD_BODY_LOOK_PERIOD,omitempty"`
	FriendSearchSec                       float32 `json:"FRIEND_SEARCH_SEC,omitempty"`
	LookTimeBase                          float32 `json:"LOOK_TIME_BASE,omitempty"`
	MaxYdistToStartWarnRequestToRequester float32 `json:"MAX_YDIST_TO_START_WARN_REQUEST_TO_REQUESTER,omitempty"`
	MinDistToCloseTalk                    float32 `json:"MIN_DIST_TO_CLOSE_TALK,omitempty"`
	MinDistToCloseTalkSqr                 float32 `json:"MIN_DIST_TO_CLOSE_TALK_SQR,omitempty"`
	MinTalkDelay                          float32 `json:"MIN_TALK_DELAY,omitempty"`
	ReserveOutTime                        float32 `json:"RESERVE_OUT_TIME,omitempty"`
	ReserveTimeStay                       float32 `json:"RESERVE_TIME_STAY,omitempty"`
	SuspetionPlaceLifetime                float32 `json:"SUSPETION_PLACE_LIFETIME,omitempty"`
	TalkDelay                             float32 `json:"TALK_DELAY,omitempty"`
	TalkDelayBig                          float32 `json:"TALK_DELAY_BIG,omitempty"`
	TryChooseReservWayOnStart             bool    `json:"TRY_CHOOSE_RESERV_WAY_ON_START,omitempty"`
	VisionDistCoefPeace                   float32 `json:"VISION_DIST_COEF_PEACE,omitempty"`
}
type Scattering struct {
	AmplitudeFactor                float32 `json:"AMPLITUDE_FACTOR,omitempty"`
	AmplitudeSpeed                 float32 `json:"AMPLITUDE_SPEED,omitempty"`
	BloodFall                      float32 `json:"BloodFall,omitempty"`
	Caution                        float32 `json:"Caution,omitempty"`
	DistFromOldPointToNotAim       float32 `json:"DIST_FROM_OLD_POINT_TO_NOT_AIM,omitempty"`
	DistFromOldPointToNotAimSqrt   float32 `json:"DIST_FROM_OLD_POINT_TO_NOT_AIM_SQRT,omitempty"`
	DistNotToShoot                 float32 `json:"DIST_NOT_TO_SHOOT,omitempty"`
	FromShot                       float32 `json:"FromShot,omitempty"`
	HandDamageAccuracySpeed        float32 `json:"HandDamageAccuracySpeed,omitempty"`
	HandDamageScatteringMinMax     float32 `json:"HandDamageScatteringMinMax,omitempty"`
	LayFactor                      float32 `json:"LayFactor,omitempty"`
	MaxScatter                     float32 `json:"MaxScatter,omitempty"`
	MinScatter                     float32 `json:"MinScatter,omitempty"`
	MovingSlowCoef                 float32 `json:"MovingSlowCoef,omitempty"`
	PoseChnageCoef                 float32 `json:"PoseChnageCoef,omitempty"`
	RecoilControlCoefShootDone     float32 `json:"RecoilControlCoefShootDone,omitempty"`
	RecoilControlCoefShootDoneAuto float32 `json:"RecoilControlCoefShootDoneAuto,omitempty"`
	RecoilYCoef                    float32 `json:"RecoilYCoef,omitempty"`
	RecoilYCoefSppedDown           float32 `json:"RecoilYCoefSppedDown,omitempty"`
	RecoilYMax                     float32 `json:"RecoilYMax,omitempty"`
	SpeedDown                      float32 `json:"SpeedDown,omitempty"`
	SpeedUp                        float32 `json:"SpeedUp,omitempty"`
	SpeedUpAim                     float32 `json:"SpeedUpAim,omitempty"`
	ToCaution                      float32 `json:"ToCaution,omitempty"`
	ToLowBotAngularSpeed           float32 `json:"ToLowBotAngularSpeed,omitempty"`
	ToLowBotSpeed                  float32 `json:"ToLowBotSpeed,omitempty"`
	ToSlowBotSpeed                 float32 `json:"ToSlowBotSpeed,omitempty"`
	ToStopBotAngularSpeed          float32 `json:"ToStopBotAngularSpeed,omitempty"`
	ToUpBotSpeed                   float32 `json:"ToUpBotSpeed,omitempty"`
	TracerCoef                     float32 `json:"TracerCoef,omitempty"`
	WorkingScatter                 float32 `json:"WorkingScatter,omitempty"`
}
type Shoot struct {
	AutomaticFireScatteringCoef      float32 `json:"AUTOMATIC_FIRE_SCATTERING_COEF,omitempty"`
	BaseAutomaticTime                float32 `json:"BASE_AUTOMATIC_TIME,omitempty"`
	CanShootsTimeToAmbush            float32 `json:"CAN_SHOOTS_TIME_TO_AMBUSH,omitempty"`
	ChanceToChangeToAutomaticFire100 float32 `json:"CHANCE_TO_CHANGE_TO_AUTOMATIC_FIRE_100,omitempty"`
	ChanceToChangeWeapon             float32 `json:"CHANCE_TO_CHANGE_WEAPON,omitempty"`
	ChanceToChangeWeaponWithHelmet   float32 `json:"CHANCE_TO_CHANGE_WEAPON_WITH_HELMET,omitempty"`
	FarDistEnemy                     float32 `json:"FAR_DIST_ENEMY,omitempty"`
	FarDistEnemySqr                  float32 `json:"FAR_DIST_ENEMY_SQR,omitempty"`
	FarDistToChangeWeapon            float32 `json:"FAR_DIST_TO_CHANGE_WEAPON,omitempty"`
	FingerHoldSingleShot             float32 `json:"FINGER_HOLD_SINGLE_SHOT,omitempty"`
	FingerHoldStationaryGrenade      float32 `json:"FINGER_HOLD_STATIONARY_GRENADE,omitempty"`
	HorizontRecoilCoef               float32 `json:"HORIZONT_RECOIL_COEF,omitempty"`
	LowDistToChangeWeapon            float32 `json:"LOW_DIST_TO_CHANGE_WEAPON,omitempty"`
	MarksmanDistSekCoef              float32 `json:"MARKSMAN_DIST_SEK_COEF,omitempty"`
	MaxDistCoef                      float32 `json:"MAX_DIST_COEF,omitempty"`
	MaxRecoilPerMeter                float32 `json:"MAX_RECOIL_PER_METER,omitempty"`
	NotToSeeEnemyToWantReloadPercent float32 `json:"NOT_TO_SEE_ENEMY_TO_WANT_RELOAD_PERCENT,omitempty"`
	NotToSeeEnemyToWantReloadSec     float32 `json:"NOT_TO_SEE_ENEMY_TO_WANT_RELOAD_SEC,omitempty"`
	RecoilDeltaPress                 float32 `json:"RECOIL_DELTA_PRESS,omitempty"`
	RecoilPerMeter                   float32 `json:"RECOIL_PER_METER,omitempty"`
	RecoilTimeNormalize              float32 `json:"RECOIL_TIME_NORMALIZE,omitempty"`
	ReloadPecnetNoEnemy              float32 `json:"RELOAD_PECNET_NO_ENEMY,omitempty"`
	RepairMalfunctionImmediateChance float32 `json:"REPAIR_MALFUNCTION_IMMEDIATE_CHANCE,omitempty"`
	RunDistNoAmmo                    float32 `json:"RUN_DIST_NO_AMMO,omitempty"`
	RunDistNoAmmoSqrt                float32 `json:"RUN_DIST_NO_AMMO_SQRT,omitempty"`
	ShootFromCover                   float32 `json:"SHOOT_FROM_COVER,omitempty"`
	SuppressByShootTime              float32 `json:"SUPPRESS_BY_SHOOT_TIME,omitempty"`
	SuppressTriggersDown             float32 `json:"SUPPRESS_TRIGGERS_DOWN,omitempty"`
	ValidateMalfunctionChance        float32 `json:"VALIDATE_MALFUNCTION_CHANCE,omitempty"`
	WaitNextSingleShot               float32 `json:"WAIT_NEXT_SINGLE_SHOT,omitempty"`
	WaitNextSingleShotLongMax        float32 `json:"WAIT_NEXT_SINGLE_SHOT_LONG_MAX,omitempty"`
	WaitNextSingleShotLongMin        float32 `json:"WAIT_NEXT_SINGLE_SHOT_LONG_MIN,omitempty"`
	WaitNextStationaryGrenade        float32 `json:"WAIT_NEXT_STATIONARY_GRENADE,omitempty"`
}

type BotHealth struct {
	BodyParts BodyParts `json:"BodyParts,omitempty"`
}

type BodyParts struct {
	Head     BodyPartHealth `json:"Head,omitempty"`
	Chest    BodyPartHealth `json:"Chest,omitempty"`
	Stomach  BodyPartHealth `json:"Stomach,omitempty"`
	LeftArm  BodyPartHealth `json:"LeftArm,omitempty"`
	RightArm BodyPartHealth `json:"RightArm,omitempty"`
	LeftLeg  BodyPartHealth `json:"LeftLeg,omitempty"`
	RightLeg BodyPartHealth `json:"RightLeg,omitempty"`
}

type BotLoadout struct {
	Earpiece        []string `json:"earpiece,omitempty"`
	Headerwear      []string `json:"headerwear,omitempty"`
	Facecover       []string `json:"facecover,omitempty"`
	BodyArmor       []string `json:"bodyArmor,omitempty"`
	Vest            []string `json:"vest,omitempty"`
	Backpack        []string `json:"backpack,omitempty"`
	PrimaryWeapon   []string `json:"primaryWeapon,omitempty"`
	SecondaryWeapon []string `json:"secondaryWeapon,omitempty"`
	Holster         []string `json:"holster,omitempty"`
	Melee           []string `json:"melee,omitempty"`
	Pocket          []string `json:"pocket,omitempty"`
}
