package configs

type ConfigDataESP struct {
	Enabled          bool                     `json:"enabled"`
	AllyBoundingBox  ConfigDataESPBoundingBox `json:"allyBox"`
	EnemyBoundingBox ConfigDataESPBoundingBox `json:"enemyBox"`
	DrawSnapLines    bool                     `json:"snapLines"`
}

type ConfigDataESPBoundingBox struct {
	Enabled               bool    `json:"enabled"`
	DrawBox               bool    `json:"box"`
	Layout                int     `json:"layout"` // 2d, 2d corners, 3d
	Outline               bool    `json:"outline"`
	OutlineColor          string  `json:"outlineColor"`
	Color                 string  `json:"color"`
	ColorAlpha            float32 `json:"alpha"`
	IsColorHealthBasedBox bool    `json:"healthBased"`
	FullfillBox           bool    `json:"fullFillBox"`
	FullfillBoxColor      string  `json:"fullFillBoxColor"`
	FullfillBoxColorAlpha float32 `json:"fullFillBoxColorAlpha"`
	DrawName              bool    `json:"name"`
	DrawHealth            bool    `json:"health"`
	Thickness             float32 `json:"thickness"`
}

type ConfigDataGlow struct {
	Enabled       bool    `json:"enabled"`
	BaseColor     string  `json:"glowBaseColor"`
	Alpha         float32 `json:"glowAlpha"`
	IsHealthBased bool    `json:"healthBased"`
}

type ConfigDataTrigger struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
	Delay   uint   `json:"delayInMs"`
}

type ConfigDataAutoWeapons struct {
	Enabled bool `json:"enabled"`
	Delay   uint `json:"delayInMs"`
}

type ConfigDataAimbot struct {
	Enabled      bool    `json:"enabled"`
	Key          string  `json:"key"`
	Fov          float64 `json:"fov"`
	Smooth       float64 `json:"smooth"`
	NonStickyAim bool    `json:"nonStickyAim"`
}

type ConfigData struct {
	ReloadKey   string                `json:"reloadKey"`
	StopKey     string                `json:"stopKey"`
	Radar       bool                  `json:"radar"`
	EngineChams bool                  `json:"engineChams"`
	Bunnyhop    bool                  `json:"bhop"`
	Glow        ConfigDataGlow        `json:"glow"`
	Triggerbot  ConfigDataTrigger     `json:"trigger"`
	AutoWeapons ConfigDataAutoWeapons `json:"autoWeapons"`
	Aimbot      ConfigDataAimbot      `json:"aimbot"`
	ESP         ConfigDataESP         `json:"esp"`
}
