package configs

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
}
