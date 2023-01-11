package entities

type Players struct {
	All []Player `json:"players"`
}
type Player struct {
	Name string `json:"name"`

	Coordinate map[string]int `json:"coordinate"`

	AttackDamage int `json:"attack-damage"`
	Health       int `json:"health"`
}
