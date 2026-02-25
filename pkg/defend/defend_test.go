package defend

import (
	"testing"
)

func TestNewDefender(t *testing.T) {
	defender := NewDefender()
	if defender == nil {
		t.Fatal("Expected defender to be created")
	}
	if defender.strategies == nil {
		t.Error("Expected strategies slice to be initialized")
	}
}

func TestDefend(t *testing.T) {
	defender := NewDefender()
	input := []byte("test input")
	strategy := "Test Strategy"

	result := defender.Defend(input, strategy)
	if result == nil {
		t.Fatal("Expected defense result")
	}
}

func TestApplyDefense(t *testing.T) {
	defender := NewDefender()
	input := []byte("test input")
	strategy := DefenseStrategy{Name: "Test"}

	output := defender.applyDefense(input, strategy)
	if output == nil {
		t.Error("Expected output to be non-nil")
	}
}

func TestPreprocessInput(t *testing.T) {
	defender := NewDefender()
	input := []byte("test input")

	output := defender.preprocessInput(input)
	if output == nil {
		t.Error("Expected output to be non-nil")
	}
}

func TestAddRandomization(t *testing.T) {
	defender := NewDefender()
	input := []byte("test input")

	output := defender.addRandomization(input)
	if output == nil {
		t.Error("Expected output to be non-nil")
	}
}

func TestApplyGradientMasking(t *testing.T) {
	defender := NewDefender()
	input := []byte("test input")

	output := defender.applyGradientMasking(input)
	if output == nil {
		t.Error("Expected output to be non-nil")
	}
}

func TestEnsembleDefense(t *testing.T) {
	input := []byte("test input")
	strategies := []string{"Strategy 1", "Strategy 2"}

	result := EnsembleDefense(input, strategies)
	if result == nil {
		t.Error("Expected defense result")
	}
}

func TestCalculateDefenseScore(t *testing.T) {
	results := []*DefenseResult{}

	score := CalculateDefenseScore(results)
	if score < 0 {
		t.Errorf("Expected non-negative score, got %f", score)
	}
}

func TestRecommendDefense(t *testing.T) {
	strategy := RecommendDefense(0.5)
	if strategy == "" {
		t.Error("Expected non-empty strategy recommendation")
	}
}

func TestGenerateDefenseReport(t *testing.T) {
	result := &DefenseResult{}

	report := GenerateDefenseReport(result)
	if report == "" {
		t.Error("Expected non-empty report")
	}
}

func TestBoolToString(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{true, "yes"},
		{false, "no"},
	}

	for _, test := range tests {
		result := boolToString(test.input)
		if result != test.expected {
			t.Errorf("Expected '%s', got '%s'", test.expected, result)
		}
	}
}

func TestDefenseStrategy(t *testing.T) {
	strategy := DefenseStrategy{
		Name:          "Test Strategy",
		Description:   "Test description",
		Effectiveness: 0.9,
	}

	if strategy.Name != "Test Strategy" {
		t.Errorf("Expected name 'Test Strategy', got '%s'", strategy.Name)
	}
}

func TestDefenseResult(t *testing.T) {
	result := DefenseResult{
		Success:       true,
		DefenseUsed:   "Test Strategy",
		ImprovedScore: 0.9,
		Cost:          0.1,
	}

	if !result.Success {
		t.Error("Expected Success to be true")
	}
}

func TestGetDefenseResult(t *testing.T) {
	result := DefenseResult{
		Success: true,
	}

	if !result.Success {
		t.Error("Expected Success to be true")
	}
}

func TestGetDefenseStrategy(t *testing.T) {
	strategy := DefenseStrategy{
		Name: "Test",
	}

	if strategy.Name != "Test" {
		t.Errorf("Expected name 'Test', got '%s'", strategy.Name)
	}
}