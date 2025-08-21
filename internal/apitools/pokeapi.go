package apitools

// All the types... too many types...

type NamedApiResource struct {
	Name string
	Url  string
}

type NamedApiResourceList struct {
	Count    int
	Next     string
	Previous string
	Results  []NamedApiResource
}

type Language struct {
	Id       int
	Name     string
	Official bool
	Iso639   string
	Iso3166  string
	Names    []Name
}

type Name struct {
	Name     string
	Language Language
}

type Version struct {
	Id            int
	Name          string
	Names         []Name
	Version_group VersionGroup
}

type VersionGroup struct {
	Id                 int
	Name               string
	Order              int
	Generation         *Generation
	Move_learn_methods []MoveLearnMethod
	Pokedexes          []Pokedex
	Regions            []Region
	Versions           []Version
}

type Generation struct {
	Id              int
	Name            string
	Abilities       []Ability
	Names           []Name
	Main_region     *Region
	Moves           []Move
	Pokemon_species []PokemonSpecies
	Types           []Type
	Version_groups  []VersionGroup
}

type Region struct {
	Id              int
	Locations       []Location
	Name            string
	Names           []Name
	Main_generation *Generation
	Pokedexes       []Pokedex
	Version_groups  []VersionGroup
}

type Location struct {
}

type EncounterMethod struct {
	Id    int
	Name  string
	Order int
	Names []Name
}

type EncounterMethodRate struct {
	Encounter_method EncounterMethod
	Version_details  EncounterVersionDetails
}

type EncounterVersionDetails struct {
	Rate    int
	Version Version
}

type LocationArea struct {
	Id                     int
	Name                   string
	Game_index             int
	Encounter_method_rates []EncounterMethodRate
	Location               Location
	Names                  []Name
	Pokemon_encounters     []PokemonEncounter
}

type Pokedex struct {
}

type PokemonSpecies struct {
}

type PokemonEncounter struct {
}

type Type struct {
}

type Move struct {
}

type MoveLearnMethod struct {
}

type Ability struct {
}
