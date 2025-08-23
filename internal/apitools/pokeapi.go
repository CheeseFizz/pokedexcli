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
	Language *Language
}

type Description struct {
	Description string
	Language    *Language
}

type FlavorText struct {
	Flavor_text string
	Language    *Language
	Version     *Version
}

type Version struct {
	Id            int
	Name          string
	Names         []Name
	Version_group *VersionGroup
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

type GenerationGameIndex struct {
	Game_index int
	Generation Generation
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
	Id           int
	Name         string
	Region       *Region
	Names        []Name
	Game_indices []GenerationGameIndex
	Areas        []LocationArea
}

type Encounter struct {
	Min_level        int
	Max_level        int
	Condition_values []EncounterConditionValue
	Chance           int
	Method           *EncounterMethod
}

type EncounterCondition struct {
	Id     int
	Name   string
	Names  []Name
	Values []EncounterConditionValue
}

type EncounterConditionValue struct {
	Id        int
	Name      string
	Names     []Name
	Condition *EncounterCondition
}

type EncounterMethod struct {
	Id    int
	Name  string
	Order int
	Names []Name
}

type EncounterMethodRate struct {
	Encounter_method *EncounterMethod
	Version_details  []EncounterVersionDetails
}

type EncounterVersionDetails struct {
	Rate    int
	Version *Version
}

type LocationArea struct {
	Id                     int
	Name                   string
	Game_index             int
	Encounter_method_rates []EncounterMethodRate
	Location               *Location
	Names                  []Name
	Pokemon_encounters     []PokemonEncounter
}

type Pokedex struct {
	Id              int
	Name            string
	Is_main_series  bool
	Descriptions    []Description
	Names           []Name
	Pokemon_entries []PokemonEntry
	Region          *Region
	Version_groups  []VersionGroup
}

type Pokemon struct {
	Id                       int
	Name                     string
	Base_experience          int
	Height                   int
	Is_default               bool
	Order                    int
	Weight                   int
	Abilities                []PokemonAbility
	Forms                    []PokemonForm
	Game_indices             []VersionGameIndex
	Held_items               []PokemonHeldItem
	Location_area_encounters string
	Moves                    []PokemonMove
	Past_types               []PokemonTypePast
	Past_abilities           []PokemonAbilityPast
	Sprites                  PokemonSprites
	Cries                    PokemonCries
	Species                  PokemonSpecies
	Stats                    []PokemonStat
	Types                    []PokemonType
}

type PokemonAbility struct {
	Is_hidden bool
	Slot      int
	Ability   *Ability
}

type PokemonType struct {
	Slot int
	Type *Type
}

type PokemonFormType struct {
	Slot int
	Type *Type
}

type PokemonTypePast struct {
	Generation *Generation
	Types      []PokemonType
}

type PokemonAbilityPast struct {
	Generation *Generation
	Abilities  []PokemonAbility
}

type PokemonHeldItem struct {
	Item            *Item
	Version_details []PokemonHeldItemVersion
}

type PokemonHeldItemVersion struct {
	Version *Version
	Rarity  int
}

type PokemonMove struct {
	Move                  *Move
	Version_group_details []PokemonMoveVersion
}

type PokemonMoveVersion struct {
	Move_learn_method *MoveLearnMethod
	Version_group     *VersionGroup
	Level_learned_at  int
	Order             int
}

type PokemonStat struct {
	Stat      *Stat
	Effort    int
	Base_stat int
}

type PokemonSprites struct {
	Front_default       int
	Front_shinty        int
	Front_female        int
	Front_shinty_female int
	Back_default        int
	Back_shinty         int
	Back_female         int
	Back_shinty_female  int
}

type PokemonCries struct {
	Latest string
	Legacy string
}

type LocationAreaEncounter struct {
	Location_area   *LocationArea
	Version_details *VersionEncounterDetail
}

type PokemonColor struct {
	Id              int
	Name            string
	Names           []Name
	Pokemon_species []PokemonSpecies
}

type PokemonForm struct {
	Id             int
	Name           string
	Names          []Name
	Order          int
	Form_order     int
	Is_default     bool
	Is_battle_only bool
	Is_mega        bool
	Form_name      string
	Pokemon        *Pokemon
	Types          []PokemonFormType
	Sprites        *PokemonFormSprites
	Version_group  *VersionGroup
	Form_names     []Name
}

type PokemonFormSprites struct {
	Front_default int
	Front_shinty  int
	Back_default  int
	Back_shinty   int
}

type PokemonHabitat struct {
	Id              int
	Name            string
	Names           []Name
	Pokemon_species []PokemonSpecies
}

type PokemonShape struct {
	Id              int
	Name            string
	Names           []Name
	Awesome_names   []AwesomeName
	Pokemon_species []PokemonSpecies
}

type AwesomeName struct {
	Awesome_name string
	Language     *Language
}

type PokemonEntry struct {
	Entry_number    int
	Pokemon_species *PokemonSpecies
}

type PokemonSpecies struct {
	Id                     int
	Name                   string
	Order                  int
	Gender_rate            int
	Capture_rate           int
	Base_happiness         int
	Is_baby                bool
	Is_Legendary           bool
	Is_mythical            bool
	Hatch_counter          int
	Has_gender_differences bool
	Forms_switchable       bool
	Growth_rate            *GrowthRate
	Pokedex_numbers        []PokemonSpeciesDexEntry
	Egg_groups             []EggGroup
	Color                  *PokemonColor
	Shape                  *PokemonShape
	Evolves_from_species   *PokemonSpecies
	Evolution_chain        *EvolutionChain
	Habitat                *PokemonHabitat
	Generation             *Generation
	Names                  []Name
	Pal_park_encounters    []PalParkEncounterArea
	Flavor_text_entries    []FlavorText
	Form_descriptions      []Description
	Genera                 []Genus
	Varieties              []PokemonSpeciesVariety
}

type PokemonSpeciesVariety struct {
	Is_default bool
	Pokemon    Pokemon
}

type Stat struct {
	Id                int
	Name              string
	Game_index        int
	Is_battle_only    bool
	Affecting_moves   *MoveStatAffectSets
	Affecting_natures *NatureStatAffectSets
	Characteristics   []Characteristic
	Move_damage_class *MoveDamageClass
	Names             []Name
}

type MoveStatAffectSets struct {
	Increase *MoveStatAffect
	Decrease *MoveStatAffect
}

type MoveStatAffect struct {
	Change int
	Move   *Move
}

type NatureStatAffectSets struct {
	Increase []Nature
	Decrease []Nature
}

type Genus struct {
	Genus    string
	Language *Language
}

type GrowthRate struct {
	Id              int
	Name            string
	Formula         string
	Descriptions    []Description
	Levels          []GrowthRateExperienceLevel
	Pokemon_species []PokemonSpecies
}

type GrowthRateExperienceLevel struct {
	Level      int
	Experience int
}

type PokemonSpeciesDexEntry struct {
	Entry_number int
	Pokedex      *Pokedex
}

type PalParkEncounterArea struct {
	Base_score int
	Rate       int
	Area       *PalParkArea
}

type PalParkArea struct {
	Id                 int
	Name               string
	Names              []Name
	Pokemon_encounters []PalParkEncounterSpecies
}

type PalParkEncounterSpecies struct {
	Base_score      int
	Rate            int
	Pokemon_species *PokemonSpecies
}

type PokemonEncounter struct {
	Pokemon         *Pokemon
	Version_details []VersionEncounterDetail
}

type VersionEncounterDetail struct {
	Version           *Version
	Max_chance        int
	Encounter_details []Encounter
}

type VersionGameIndex struct {
	Game_index int
	Version    *Version
}

type VersionGroupFlavorText struct {
	Text          string
	Language      *Language
	Version_group *VersionGroup
}

type Type struct {
	Id                    int
	Name                  string
	Damage_relations      *TypeRelations
	Past_damage_relations []TypeRelationsPast
	Game_indices          []GenerationGameIndex
	Generation            *Generation
	Move_damage_class     *MoveDamageClass
	Names                 []Name
	Pokemon               *TypePokemon
	Moves                 []Move
}

type TypePokemon struct {
	Slot    int
	Pokemon Pokemon
}

type TypeRelations struct {
	No_damage_to       *Type
	Half_damage_to     *Type
	Double_damage_to   *Type
	No_damage_from     *Type
	Half_damage_from   *Type
	Double_damage_from *Type
}

type TypeRelationsPast struct {
	Generation       *Generation
	Damage_relations *TypeRelations
}

type Move struct {
	Id                   int
	Name                 string
	Accuracy             int
	Effect_chance        int
	PP                   int
	Priority             int
	Power                int
	Contest_combos       *ContestComboSets
	Contest_type         *ContestType
	Contest_effect       *ContestEffect
	Damage_class         *MoveDamageClass
	Effect_entries       []VerboseEffect
	Effect_changes       []AbilityEffectChange
	Learned_by_pokemon   []Pokemon
	Flavor_text_entries  []MoveFlavorText
	Generation           *Generation
	Machines             []MachineVersionDetail
	Meta                 *MoveMetaData
	Names                []Name
	Past_values          []PastMoveStatValues
	Stat_changes         []MoveStatChange
	Super_contest_effect *SuperContestEffect
	Target               *MoveTarget
	Type                 *Type
}

type MoveDamageClass struct {
	Id           int
	Name         string
	Descriptions []Description
	Moves        []Move
	Names        []Name
}

type MoveTarget struct {
	Id           int
	Name         string
	Descriptions []Description
	Moves        []Move
	Names        []Name
}

type MoveLearnMethod struct {
	Id             int
	Name           string
	Descriptions   []Description
	Names          []Name
	Version_groups []VersionGroup
}

type MoveFlavorText struct {
	Flavor_text   string
	Language      *Language
	Version_group *VersionGroup
}

type MoveMetaData struct {
	Ailment        *MoveAilment
	Category       *MoveCategory
	Min_hits       int
	Max_hits       int
	Min_turns      int
	Max_turns      int
	Drain          int
	Healing        int
	Crit_rate      int
	Ailment_chance int
	Flinch_chance  int
	Stat_chance    int
}

type MoveStatChange struct {
	Change int
	Stat   *Stat
}

type PastMoveStatValues struct {
	Accuracy       int
	Effect_chance  int
	Power          int
	PP             int
	Effect_entries []VerboseEffect
	Type           *Type
	Version_group  *VersionGroup
}

type MoveAilment struct {
	Id    int
	Name  string
	Moves []Move
	Names []Name
}

type MoveBattleStyle struct {
	Id    int
	Name  string
	Names []Name
}

type MoveCategory struct {
	Id           int
	Name         string
	Moves        []Move
	Descriptions []Description
}

type ContestComboSets struct {
	Normal *ContestComboDetail
	Super  *ContestComboDetail
}

type ContestComboDetail struct {
	Use_before []Move
	Use_after  []Move
}

type Ability struct {
	Id                  int
	Name                string
	Is_main_series      bool
	Generation          Generation
	Names               []Name
	Effect_entries      []VerboseEffect
	Effect_changes      []AbilityEffectChange
	Flavor_text_entries []AbilityFlavorText
	Pokemon             []AbilityPokemon
}

type AbilityEffectChange struct {
	Effect_entries []Effect
	Version_group  VersionGroup
}

type AbilityFlavorText struct {
	Flavor_text   string
	Language      *Language
	Version_group *VersionGroup
}

type AbilityPokemon struct {
	Is_hidden bool
	Slot      int
	Pokemon   *Pokemon
}

type Characteristic struct {
	Id              int
	Gene_modulo     int
	Possible_values []int
	Highest_stat    *Stat
	Descriptions    []Description
}

type EggGroup struct {
	Id              int
	Name            string
	Names           []Name
	Pokemon_species []PokemonSpecies
}

type Gender struct {
	Id                      int
	Name                    string
	Pokemon_species_details []PokemonSpeciesGender
	Required_for_evolution  []PokemonSpecies
}

type PokemonSpeciesGender struct {
	Rate            int
	Pokemon_species *PokemonSpecies
}

type Nature struct {
	Id                            int
	Name                          string
	Decreased_stat                *Stat
	Increased_stat                *Stat
	Hates_flavor                  *BerryFlavor
	Likes_flavor                  *BerryFlavor
	Pokeathlon_stat_changes       []NatureStatChange
	Move_battle_style_preferences []MoveBattleStylePreference
	Names                         []Name
}

type NatureStatChange struct {
	Max_change      int
	Pokeathlon_stat *PokeathlonStat
}

type MoveBattleStylePreference struct {
	Low_hp_preference  int
	High_hp_preference int
	Move_battle_style  *MoveBattleStyle
}

type PokeathlonStat struct {
	Id                int
	Name              string
	Names             []Name
	Affecting_natures *NaturePokeathlonStatAffectSets
}

type NaturePokeathlonStatAffectSets struct {
	Increase []NaturePokeathlonStatAffect
	Decrease []NaturePokeathlonStatAffect
}

type NaturePokeathlonStatAffect struct {
	Max_change int
	Nature     *Nature
}

type Effect struct {
	Effect   string
	Language *Language
}

type VerboseEffect struct {
	Effect       string
	Short_effect string
	Language     *Language
}

type ContestType struct {
	Id           int
	Name         string
	Berry_flavor *BerryFlavor
	Names        []ContestName
}

type ContestName struct {
	Name     string
	Color    string
	Language *Language
}

type ContestEffect struct {
	Id                  int
	Appeal              int
	Jam                 int
	Effect_entries      []Effect
	Flavor_text_entries []FlavorText
}

type SuperContestEffect struct {
	Id                  int
	Appeal              int
	Flavor_text_entries []FlavorText
	Moves               []Move
}

type Machine struct {
	Id            int
	Item          *Item
	Move          *Move
	Version_group *VersionGroup
}

type MachineVersionDetail struct {
	Machine       *Machine
	Version_group *VersionGroup
}

type Item struct {
	Id                  int
	Name                string
	Cost                int
	Fling_power         int
	Fling_effect        *ItemFlingEffect
	Attributes          []ItemAttribute
	Category            *ItemCategory
	Effect_entries      []VerboseEffect
	Flavor_text_entries []VersionGroupFlavorText
	Game_indices        []GenerationGameIndex
	Names               []Name
	Sprites             *ItemSprites
	Held_by_pokemon     []ItemHolderPokemon
	Baby_trigger_for    *EvolutionChain
	Machines            []MachineVersionDetail
}

type ItemSprites struct {
	Default string
}

type ItemHolderPokemon struct {
	Pokemon         *Pokemon
	Version_details []ItemHolderPokemonVersionDetail
}

type ItemHolderPokemonVersionDetail struct {
	Rarity  int
	Version *Version
}

type ItemAttribute struct {
	Id           int
	Name         string
	Items        []Item
	Names        []Name
	Descriptions []Description
}

type ItemCategory struct {
	Id     int
	Name   string
	Items  []Item
	Names  []Name
	Pocket *ItemPocket
}

type ItemFlingEffect struct {
	Id             int
	Name           string
	Items          []Item
	Effect_entries []Effect
}

type ItemPocket struct {
	Id         int
	Name       string
	Names      []Name
	Categories []ItemCategory
}

type Berry struct {
	Id                 int
	Name               string
	Growth_time        int
	Max_harvest        int
	Natural_gift_power int
	Size               int
	Smoothness         int
	Soil_dryness       int
	Firmness           *BerryFirmness
	Flavors            []BerryFlavorMap
	Item               *Item
	Natural_gift_type  *Type
}

type BerryFlavorMap struct {
	Potency int
	Flavor  *BerryFlavor
}

type BerryFirmness struct {
	Id      int
	Name    string
	Names   []Name
	Berries []Berry
}

type BerryFlavor struct {
	Id           int
	Name         string
	Names        []Name
	Berries      []FlavorBerryMap
	Contest_type *ContestType
}

type FlavorBerryMap struct {
	Potency int
	Berry   *Berry
}

type EvolutionChain struct {
	Id                int
	Baby_trigger_item *Item
	Chain             *ChainLink
}

type ChainLink struct {
	Is_baby           bool
	Species           *PokemonSpecies
	Evolution_details []EvolutionDetail
	Evolves_to        []ChainLink
}

type EvolutionDetail struct {
	Item                    *Item
	Trigger                 *EvolutionTrigger
	Gender                  int
	Held_item               *Item
	Known_move              *Move
	Known_move_type         *Type
	Location                *Location
	Min_level               int
	Min_happiness           int
	Min_beauty              int
	Min_affection           int
	Needs_overworld_rain    bool
	Party_species           *PokemonSpecies
	Party_type              *Type
	Relative_physical_stats int
	Time_of_day             string
	Trade_species           *PokemonSpecies
	Turn_upside_down        bool
}

type EvolutionTrigger struct {
	Id              int
	Name            string
	Names           []Name
	Pokemon_species *PokemonSpecies
}
