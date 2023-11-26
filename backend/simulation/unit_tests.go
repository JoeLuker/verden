// File: simulation/unit_tests.go
package simulation

import "testing"

func TestNewCharacterSkills(t *testing.T) {
	skills := NewCharacterSkills()
	if skills.CraftingAbility != 80 || skills.Negotiation != 60 {
		t.Errorf("Default values for CharacterSkills are incorrect")
	}
}

func TestImproveSkill(t *testing.T) {
	initialSkill := 50.0
	improvedSkill := improveSkill(initialSkill)
	if improvedSkill <= initialSkill {
		t.Errorf("Skill improvement failed to increase the skill value")
	}
}

func TestUpdateResourceAvailability(t *testing.T) {
	initialAvailability := 100.0
	politicalStability := 70.0
	updatedAvailability := updateResourceAvailability(initialAvailability, politicalStability)
	if updatedAvailability >= initialAvailability {
		t.Errorf("Resource availability update logic is not functioning as expected")
	}
}

func TestValidateSimulationParams(t *testing.T) {
	params := SimulationParams{
		Economic: NewEconomicFactors(),
		Skills:   NewCharacterSkills(),
	}
	err := params.Validate()
	if err != nil {
		t.Errorf("Validation failed for valid SimulationParams")
	}
}

// Additional tests for other functions can be added here.
