// File: simulation/character_skills.go
package simulation

type CharacterSkills struct {
    CraftingAbility float64
    Negotiation     float64
    // Add other skills as needed
}

func NewCharacterSkills() *CharacterSkills {
    return &CharacterSkills{
        CraftingAbility: 80,  // Default value, adjust as needed
        Negotiation:     60,  // Default value, adjust as needed
    }
}

func (c *CharacterSkills) Improve() {
    // Implement logic to improve character skills
    // Example: c.CraftingAbility += 5
}
