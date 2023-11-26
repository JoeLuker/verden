// File: simulation/simulation_params.go
package simulation

import "errors"

type SimulationParams struct {
    Economic *EconomicFactors
    Skills   *CharacterSkills
}

func (sp *SimulationParams) Validate() error {
    // Validation logic for EconomicFactors
    if sp.Economic == nil || !validateEconomicFactors(sp.Economic) {
        return errors.New("invalid or missing EconomicFactors")
    }

    // Validation logic for CharacterSkills
    if sp.Skills == nil || !validateCharacterSkills(sp.Skills) {
        return errors.New("invalid or missing CharacterSkills")
    }

    return nil
}

func validateEconomicFactors(e *EconomicFactors) bool {
    // Example: Validate if factors are within realistic ranges
    return e.ResourceAvailability >= 0 && e.ResourceAvailability <= 100 &&
           e.MarketDemand >= 0 && e.MarketDemand <= 100 &&
           e.PoliticalStability >= 0 && e.PoliticalStability <= 100 &&
           e.TradeRelations >= 0 && e.TradeRelations <= 100
}

func validateCharacterSkills(s *CharacterSkills) bool {
    // Example: Validate if skills are within realistic ranges
    return s.CraftingAbility >= 0 && s.CraftingAbility <= 100 &&
           s.Negotiation >= 0 && s.Negotiation <= 100 &&
           s.CombatAbility >= 0 && s.CombatAbility <= 100 &&
           s.MagicUse >= 0 && s.MagicUse <= 100 &&
           s.Diplomacy >= 0 && s.Diplomacy <= 100
}
