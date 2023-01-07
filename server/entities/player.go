package entities

type Player struct {
	Name string `json:"name"`

	Coordinate map[string]int `json:"coordinate"`

	AttackDamage int `json:"attack-damage"`
	Health       int `json:"health"`
}
