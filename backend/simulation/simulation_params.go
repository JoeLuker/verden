// File: simulation/simulation_params.go
package simulation

type SimulationParams struct {
    Economic *EconomicFactors
    Skills   *CharacterSkills
}

// Add a Validate method in the SimulationParams struct.
func (sp *SimulationParams) Validate() error {
    // Add validation logic here
    // Example: Check if any parameter is out of an expected range
    return nil
}
