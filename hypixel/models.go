package hypixel

// KeyInfoResponse is the bare response of the API containing maybe the record, maybe a cause and a bool determine if
// the request was successful.
type KeyInfoResponse struct {
	Record  *KeyInfo `json:"record"`
	Cause   string   `json:"cause"`
	Success bool     `json:"success"`
}

// KeyInfo contains the information of a /key request.
type KeyInfo struct {
	Key              string `json:"key"`
	OwnerUUID        string `json:"ownerUuid"`
	TotalQueries     int32  `json:"totalQueries"`
	QueriesInPastMin int32  `json:"queriesInPastMin"`
}

// GuildIdResponse is the bare response of the API containing maybe a guild id, maybe a cause and a bool determine if
// the request was successful.
type GuildIdResponse struct {
	Guild	string   `json:"guild"`
	Cause   string   `json:"cause"`
	Success bool     `json:"success"`
}

// FriendsResponse is the bare response of the API containing maybe an array of friends, maybe a cause and a bool
// determine if the request was successful.
type FriendsResponse struct {
	Records []*Friend `json:"records"`
	Cause   string    `json:"cause"`
	Success bool      `json:"success"`
}

// Friend contains the information of a /friends records entry.
type Friend struct {
	ID 		     string `json:"_id"`
	Receiver     string `json:"receiver"`
	Sender       string `json:"sender"`
	UUIDReceiver string `json:"uuidReceiver"`
	UUIDSender   string `json:"uuidSender"`
	Started      int64  `json:"started"`
}

type Player struct {
	ID                         string                            `json:"_id"`
	Achievements               map[string]int32                  `json:"achievements"`
	OneTimeAchievements        []string                          `json:"achievementsOneTime"`
	Clock                      bool                              `json:"clock"`
	DisplayName                string                            `json:"displayname"`
	EulaCoins                  bool                              `json:"eulaCoins"`
	FirstLogin                 int64                             `json:"firstLogin"`
	Fly                        bool                              `json:"fly"`
	// ToDo: friendRequests (I currently don't know the format)
	// ToDo: friendRequestsUuid
	Gadget                     string                            `json:"gadget"`
	Karma                      int32                             `json:"karma"`
	KnownAliases               []string                          `json:"knownAliases"`
	KnownAliasesLower          []string                          `json:"knownAliasesLower"`
	LastLogin                  int64                             `json:"lastLogin"`
	MostRecentMinecraftVersion int32                             `json:"mostRecentMinecraftVersion"`
	MostRecentlyThanked        string                            `json:"mostRecentlyThanked"`
	MostRecentlyThankedUUID    string                            `json:"mostRecentlyThankedUuid"`
	NetworkExp                 int32                             `json:"networkExp"`
	NetworkLevel               int32                             `json:"networkLevel"`
	NewClock                   string                            `json:"newClock"`
	NewPackageRank             string                            `json:"newPackageRank"`
	PackageRank                string                            `json:"packageRank"`
	Packages                   []string                          `json:"packages"`
	PlayerName                 string                            `json:"playername"`
	Quests                     []*PlayerQuest                    `json:"quests"`
	Settings                   map[string]interface{}            `json:"settings"`
	SpecSpeed                  int32                             `json:"spec_speed"`
	Stats                      map[string]map[string]interface{} `json:"stats"`
	ThanksSent                 int32                             `json:"thanksSent"`
	TimePlaying                int64                             `json:"timePlaying"`
	TournamentTokens           int32                             `json:"tournamentTokens"`
	UUID                       string                            `json:"uuid"`
	VanityMeta                 PlayerVanityMeta                  `json:"vanityMeta"`
	VanityTokens               int32                             `json:"vanityTokens"`
	WebsiteSet                 bool                              `json:"websiteSet"`
	LastEugeneMessage          int64                             `json:"lastEugeneMessage"`
	LastSurvey                 int64                             `json:"last_survey"`
	Eugene                     PlayerEugene                      `json:"eugene"`
	Voting                     map[string]int64                  `json:"voting"`
	McVersionRp                string                            `json:"mcVersionRp"`
	MostRecentlyTippedUuid     string                            `json:"mostRecentlyTippedUuid"`
	PetConsumables             map[string]int32                  `json:"petConsumables"`
	HousingMeta                PlayerHousingMeta                 `json:"housingMeta"`
	PetStats                   map[string]PlayerPetStats         `json:"petStats"`
	PetJourneyTime             int64                             `json:"petJourneyTimestamp"`
	CurrentGadget              string                            `json:"currentGadget"`
	ParkourCompletions         map[string][]PlayerParkourTime    `json:"parkourCompletions"`
	Language                   string                            `json:"language"`
	// ToDo: Rewards
	// ToDo: Presents
	SpecialtyCooldowns         map[string]bool                   `json:"specialtyCooldowns"`
	Outfit                     PlayerOutfit                      `json:"outfit"`
	MostRecentGameType         string                            `json:"mostRecentGameType"`
	Channel                    string                            `json:"channel"`
	Stoggle                    bool                              `json:"stoggle"`
	BoxGiveaways               map[string]bool                   `json:"boxGiveaways"`
	ParticleQuality            string                            `json:"particleQuality"`
	CustomFilter               string                            `json:"customFilter"`
	Vanished                   bool                              `json:"vanished"`
	PetUpdate                  int64                             `json:"petUpdate"`
	Transformation             string                            `json:"transformation"`
	Particle                   string                            `json:"particle"`
	Rank                       string                            `json:"rank"`
	MysteryKeys                int32                             `json:"mysteryKeys"`
	Silence                    bool                              `json:"silence"`
	SpecNightVision            bool                              `json:"spec_night_vision"`
	UserPermissions            []string                          `json:"userPermissions"`
	CurrentPet                 string                            `json:"currentPet"`
}

type PlayerQuest struct {
	Completions []*PlayerQuestCompletion `json:"completions"`
	Active      PlayerQuestActive        `json:"active"`
}

type PlayerQuestCompletion struct {
	Time int64 `json:"time"`
}

type PlayerQuestActive struct {
	Started    int64            `json:"started"`
	Objectives map[string]int32 `json:"objectives"`
}

type PlayerVanityMeta struct {
	Packages []string `json:"packages"`
}

type PlayerEugene struct {
	DailyTwoKExp  int64 `json:"dailyTwoKExp"`
	WeeklyBooster int64 `json:"weekly_booster"`
}

type PlayerHousingMeta struct {
	AllowedBlocks  []string          `json:"allowedBlocks"`
	Packages       []string          `json:"packages"`
	TutorialStep   string            `json:"tutorialStep"`
	PlayerSettings map[string]string `json:"playerSettings"`
	GivenCookiesA  []string          `json:"given_cookies_a"`
	GivenCookiesB  []string          `json:"given_cookies_b"`
	GivenCookiesC  []string          `json:"given_cookies_c"`
	GivenCookiesD  []string          `json:"given_cookies_d"`
}

type PlayerPetStats struct {
	Name     string              `json:"name"`
	Hunger   PlayerPetStatsValue `json:"HUNGER"`
	Thirst   PlayerPetStatsValue `json:"THIRST"`
	Exercise PlayerPetStatsValue `json:"EXERCISE"`
}

type PlayerPetStatsValue struct {
	Time  int64 `json:"timestamp"`
	Value int32 `json:"value"`
}

type PlayerParkourTime struct {
	TimeStart int64 `json:"timeStart"`
	TimeTook  int64 `json:"timeTook"`
}

type PlayerOutfit struct {
	Boots string `json:"BOOTS"`
}
