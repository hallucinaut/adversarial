// Package defend provides adversarial defense mechanisms.
package defend

import (
	"math"
)

// DefenseStrategy represents a defense strategy.
type DefenseStrategy struct {
	Name        string
	Description string
	Effectiveness float64
	PerformanceCost float64
}

// DefenseResult contains defense results.
type DefenseResult struct {
	Success       bool
	DefenseUsed   string
	ImprovedScore float64
	Cost          float64
}

// Defender applies adversarial defenses.
type Defender struct {
	strategies []DefenseStrategy
}

// NewDefender creates a new defender.
func NewDefender() *Defender {
	return &Defender{
		strategies: []DefenseStrategy{
			{Name: "Adversarial Training", Description: "Train on adversarial examples", Effectiveness: 0.9, PerformanceCost: 0.3},
			{Name: "Input Preprocessing", Description: "Clean input before processing", Effectiveness: 0.7, PerformanceCost: 0.1},
			{Name: "Randomization", Description: "Add randomness to input", Effectiveness: 0.6, PerformanceCost: 0.05},
			{Name: "Ensemble Defense", Description: "Use multiple models", Effectiveness: 0.85, PerformanceCost: 0.5},
			{Name: "Gradient Masking", Description: "Hide gradient information", Effectiveness: 0.65, PerformanceCost: 0.15},
		},
	}
}

// Defend applies defense strategy to input.
func (d *Defender) Defend(input []byte, strategy string) *DefenseResult {
	for _, strat := range d.strategies {
		if strat.Name == strategy {
			// Apply defense
			defended := d.applyDefense(input, strat)
			return &DefenseResult{
				Success:      true,
				DefenseUsed:  strat.Name,
				ImprovedScore: strat.Effectiveness,
				Cost:         strat.PerformanceCost,
			}
		}
	}

	return &DefenseResult{
		Success: false,
	}
}

// applyDefense applies a specific defense technique.
func (d *Defender) applyDefense(input []byte, strategy DefenseStrategy) []byte {
	switch strategy.Name {
	case "Input Preprocessing":
		return d.preprocessInput(input)
	case "Randomization":
		return d.addRandomization(input)
	case "Gradient Masking":
		return d.applyGradientMasking(input)
	default:
		return input
	}
}

// preprocessInput cleans input.
func (d *Defender) preprocessInput(input []byte) []byte {
	// Apply smoothing filter
	filtered := make([]byte, len(input))
	windowSize := 3

	for i := 0; i < len(input); i++ {
		sum := 0.0
		count := 0
		for j := i - windowSize/2; j <= i+windowSize/2; j++ {
			if j >= 0 && j < len(input) {
				sum += float64(input[j])
				count++
			}
		}
		filtered[i] = byte(sum / float64(count))
	}

	return filtered
}

// addRandomization adds noise for defense.
func (d *Defender) addRandomization(input []byte) []byte {
	defended := make([]byte, len(input))

	// Add small random noise
	for i, b := range input {
		noise := float64(b%7-3) // Small random value between -3 and 3
		defended[i] = byte(math.Max(0, math.Min(255, float64(b)+noise)))
	}

	return defended
}

// applyGradientMasking applies gradient masking.
func (d *Defender) applyGradientMasking(input []byte) []byte {
	defended := make([]byte, len(input))

	// Smooth gradients
	for i := 0; i < len(input); i++ {
		if i == 0 || i == len(input)-1 {
			defended[i] = input[i]
		} else {
			defended[i] = byte((float64(input[i-1]) + 2*float64(input[i]) + float64(input[i+1])) / 4)
		}
	}

	return defended
}

// EnsembleDefense combines multiple defenses.
func EnsembleDefense(input []byte, strategies []string) *DefenseResult {
	defender := NewDefender()
	result := &DefenseResult{
		Success: true,
	}

	for _, strategy := range strategies {
		def := defender.Defend(input, strategy)
		if def.Success {
			result.ImprovedScore += def.ImprovedScore
			result.Cost += def.Cost
		}
	}

	if len(strategies) > 0 {
		result.ImprovedScore /= float64(len(strategies))
		result.Cost /= float64(len(strategies))
	}

	return result
}

// CalculateDefenseScore calculates defense score.
func CalculateDefenseScore(defenseResults []*DefenseResult) float64 {
	if len(defenseResults) == 0 {
		return 0
	}

	score := 0.0
	for _, result := range defenseResults {
		if result.Success {
			score += result.ImprovedScore
		}
	}

	return score / float64(len(defenseResults))
}

// RecommendDefense recommends best defense strategy.
func RecommendDefense(adversarialScore float64) string {
	if adversarialScore > 0.8 {
		return "Adversarial Training"
	} else if adversarialScore > 0.6 {
		return "Input Preprocessing"
	} else if adversarialScore > 0.4 {
		return "Randomization"
	}

	return "Ensemble Defense"
}

// GenerateDefenseReport generates defense report.
func GenerateDefenseReport(result *DefenseResult) string {
	var report string

	report += "=== Defense Report ===\n\n"
	report += "Success: " + boolToString(result.Success) + "\n"
	report += "Defense Used: " + result.DefenseUsed + "\n"
	report += "Effectiveness: " + string(rune(int(result.ImprovedScore*100)+48)) + "%\n"
	report += "Performance Cost: " + string(rune(int(result.Cost*100)+48)) + "%\n"

	return report
}

// boolToString converts bool to string.
func boolToString(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}