package npcs

// Power describe the attack and defence of an NPC
type Power struct {
	Attack  int
	Defence int
}

// Location describes where in the virtual world an NPC exists
type Location struct {
	X float64
	Y float64
	Z float64
}

// NonPlayerCharacter represents metadata for an in-game creature
type NonPlayerCharacter struct {
	Name  string
	Speed int
	HP    int
	Power Power
	Loc   Location
}
