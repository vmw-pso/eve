package sde

type TypeIDs struct {
	TypeIDs map[int]Type
}

type Type struct {
	GroupID        int         `yaml:"groupID" json:"groupId"`
	Mass           float64     `yaml:"mass" json:"mass"`
	Name           Name        `yaml:"name" json:"name"`
	PortionSize    int         `yaml:"portionSize" json:"portionSize"`
	Published      bool        `yaml:"published" json:"published"`
	Radius         float64     `yaml:"radius,omitempty" json:"radius,omitempty"`
	Volume         float64     `yaml:"volume,omitempty" json:"volume,omitempty"`
	Description    Description `yaml:"description,omitempty" json:"description,omitempty"`
	GraphicID      int         `yaml:"graphicID,omitempty" json:"graphicId,omitempty"`
	SoundID        int         `yaml:"soundID,omitempty" json:"soundId,omitempty"`
	IconID         int         `yaml:"iconID,omitempty" json:"iconId,omitempty"`
	RaceID         int         `yaml:"raceID,omitempty" json:"raceId,omitempty"`
	SofFactionName string      `yaml:"sofFactionName,omitempty" json:"sofFactionName,omitempty"`
	BasePrice      float64     `yaml:"basePrice,omitempty" json:"basePrice,omitempty"`
	MarketGroupID  int         `yaml:"marketGroupID,omitempty" json:"marketGroupId,omitempty"`
	Capacity       float64     `yaml:"capacity,omitempty" json:"capacity,omitempty"`
}

type Name struct {
	German   string `yaml:"de" json:"de"`
	English  string `yaml:"en" json:"en"`
	French   string `yaml:"fr" json:"fr"`
	Japanese string `yaml:"ja" json:"ja"`
	Russian  string `yaml:"ru" json:"ru"`
	Chinese  string `yaml:"zh" json:"zh"`
}

type Description struct {
	German   string `yaml:"de" json:"de"`
	English  string `yaml:"en" json:"en"`
	French   string `yaml:"fr" json:"fr"`
	Japanese string `yaml:"ja" json:"ja"`
	Russian  string `yaml:"ru" json:"ru"`
	Chinese  string `yaml:"zh" json:"zh"`
}
