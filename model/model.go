package model

// NamedResource is a pair of Name of a resource and a URL where that resource can be fetched from.
type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Pokemon contains data for a single Pokemon
type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	// The base experience gained for defeating this Pokémon
	BaseExperience int `json:"base_experience"`

	// The height of this Pokémon in decimetres
	Height int `json:"height"`

	// Set for exactly one Pokémon used as the default for each species
	IsDefault bool `json:"is_default"`

	// Order for sorting. Almost national order, except families are grouped together
	Order int `json:"order"`

	// The weight of this Pokémon in hectograms
	Weight int `json:"weight"`

	// A list of abilities this Pokémon could potentially have
	Abilities []Ability `json:"abilities"`

	// A list of details showing the abilities this pokémon had in previous generations
	PastAbilities []AbilityPast `json:"past_abilities"`

	// A list of forms this Pokémon can take on
	Forms []NamedResource `json:"forms"`

	// A list of game indices relevant to Pokémon item by generation
	GameIndices []GameIndex `json:"game_indices"`

	// A list of items this Pokémon may be holding when encountered
	HeldItems []HeldItem `json:"held_items"`

	// A link to a list of location areas, as well as encounter details pertaining to specific versions
	LocationAreaEncounters string `json:"location_area_encounters"`

	// A list of moves along with learn methods and level details pertaining to specific version groups
	Moves []Move `json:"moves"`

	// A list of details showing types this Pokémon has
	Types []Type `json:"types"`

	// A list of details showing types this Pokémon had in previous generations
	PastTypes []TypePast `json:"past_types"`

	// A set of sprites used to depict this Pokémon in the game
	Sprites Sprites `json:"sprites"`

	//A set of cries used to depict this Pokémon in the game
	Cries Cries `json:"cries"`

	// The species this Pokémon belongs to
	Species NamedResource `json:"species"`

	// A list of base stat values for this Pokémon
	Stats []Stat `json:"stats"`
}

type Ability struct {
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
	Ability  NamedResource `json:"ability"`
}

type AbilityPast struct {
	Generation NamedResource `json:"generation"`
	Abilities  []Ability     `json:"abilities"`
}

type Type struct {
	Slot int           `json:"slot"`
	Type NamedResource `json:"type"`
}

type TypePast struct {
	Generation NamedResource `json:"generation"`
	Types      []Type        `json:"types"`
}

type GameIndex struct {
	GameIndex int           `json:"game_index"`
	Version   NamedResource `json:"version"`
}

type HeldItem struct {
	Item   NamedResource `json:"item"`
	Rarity int           `json:"rarity"`
}

type Move struct {
	Move                NamedResource `json:"move"`
	VersionGroupDetails []MoveVersion `json:"version_group_details"`
}

type MoveVersion struct {
	MoveLearnMethod NamedResource `json:"move_learn_method"`
	VersionGroup    NamedResource `json:"version_group"`
	LevelLearnedAt  int           `json:"level_learned_at"`
	Order           int           `json:"order"`
}

type Stat struct {
	Stat     NamedResource `json:"stat"`
	Effort   int           `json:"effort"`
	BaseStat int           `json:"base_stat"`
}

type Sprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// A Generation is a grouping of the Pokémon games that separates them based on the Pokémon they include.
type Generation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	// A list of abilities that were introduced in this generation
	Abilities []NamedResource `json:"abilities"`

	//The name of this resource listed in different languages
	Names []Name `json:"names"`

	// The main region traveled in this generation
	MainRegion NamedResource `json:"main_region"`

	// A list of moves that were introduced in this generation
	Moves []NamedResource `json:"moves"`

	// A list of Pokémon species that were introduced in this generation
	PokemonSpecies []NamedResource `json:"pokemon_species"`

	// A list of types that were introduced in this generation
	Types []NamedResource `json:"types"`

	// A list of version groups that were introduced in this generation
	VersionGroups []NamedResource `json:"version_groups"`
}

type Name struct {
	Name     string        `json:"name"`
	Language NamedResource `json:"language"`
}

// A PokemonForm represents different visual forms as a Pokémon.
// This type does not contain all the fields from the API and is just for demo purposes
type PokemonForm struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	//The order in which forms should be sorted within all forms
	Order int `json:"order"`

	// The order in which forms should be sorted within a species' forms
	FormOrder int `json:"form_order"`

	// True for exactly one form used as the default for each Pokémon.
	IsDefault bool `json:"is_default"`

	// Whether this form can only happen during battle
	IsBattleOnly bool `json:"is_battle_only"`

	// Whether this form requires mega evolution
	IsMega bool `json:"is_mega"`
}
