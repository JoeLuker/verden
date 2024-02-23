// File: simulation/character_skills.go
package simulation

type CharacterSkills struct {
	CraftingAbility float64
	Negotiation     float64
	CombatAbility   float64
	MagicUse        float64
	Diplomacy       float64
	// Add other skills as needed
}

func (c *CharacterSkills) Improve() {
	// Example improvements with diminishing returns
	c.CraftingAbility = improveSkill(c.CraftingAbility)
	c.Negotiation = improveSkill(c.Negotiation)
	c.CombatAbility = improveSkill(c.CombatAbility)
	c.MagicUse = improveSkill(c.MagicUse)
	c.Diplomacy = improveSkill(c.Diplomacy)
}

// improveSkill increases the skill value with diminishing returns
func improveSkill(skill float64) float64 {
	return skill + (100-skill)*0.05 // Increment decreases as skill approaches 100
}
