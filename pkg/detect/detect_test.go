package detect

import (
	"testing"
)

func TestNewDetector(t *testing.T) {
	detector := NewDetector()
	if detector == nil {
		t.Fatal("Expected detector to be created")
	}
	if detector.methods == nil {
		t.Error("Expected methods slice to be initialized")
	}
}

func TestDetect(t *testing.T) {
	detector := NewDetector()
	input := []byte("test input")

	result := detector.Detect(input)
	if result == nil {
		t.Fatal("Expected detection result")
	}
}

func TestExtractFeatures(t *testing.T) {
	detector := NewDetector()
	input := []byte("test input")

	features := detector.extractFeatures(input)
	if features == nil {
		t.Error("Expected features to be extracted")
	}
}

func TestCalculateMean(t *testing.T) {
	detector := NewDetector()
	input := []byte("test input")

	mean := detector.calculateMean(input)
	if mean < 0 {
		t.Errorf("Expected non-negative mean, got %f", mean)
	}
}

func TestCalculateStdDev(t *testing.T) {
	detector := NewDetector()
	input := []byte("test input")
	mean := float64(0)

	stdDev := detector.calculateStdDev(input, mean)
	if stdDev < 0 {
		t.Errorf("Expected non-negative stddev, got %f", stdDev)
	}
}

func TestCalculateEntropy(t *testing.T) {
	detector := NewDetector()
	input := []byte("test input")

	entropy := detector.calculateEntropy(input)
	if entropy < 0 {
		t.Errorf("Expected non-negative entropy, got %f", entropy)
	}
}

func TestCalculateGradient(t *testing.T) {
	detector := NewDetector()
	input := []byte("test input")

	gradient := detector.calculateGradient(input)
	if gradient < 0 {
		t.Errorf("Expected non-negative gradient, got %f", gradient)
	}
}

func TestCalculateNoiseLevel(t *testing.T) {
	detector := NewDetector()
	input := []byte("test input")

	noiseLevel := detector.calculateNoiseLevel(input)
	if noiseLevel < 0 {
		t.Errorf("Expected non-negative noise level, got %f", noiseLevel)
	}
}

func TestAnalyzeFeatures(t *testing.T) {
	detector := NewDetector()
	features := []Feature{}

	score := detector.analyzeFeatures(features)
	if score < 0 {
		t.Errorf("Expected non-negative score, got %f", score)
	}
}

func TestCalculateAnomalyScore(t *testing.T) {
	detector := NewDetector()
	feature := Feature{Name: "test", Value: 0.5}

	score := detector.calculateAnomalyScore(feature)
	if score < 0 {
		t.Errorf("Expected non-negative score, got %f", score)
	}
}

func TestCheckMethod(t *testing.T) {
	detector := NewDetector()
	method := DetectionMethod{Name: "Statistical Analysis"}
	features := []Feature{}
	input := []byte("test")

	pattern := detector.checkMethod(method, features, input)
	// Pattern may be empty, just verify it doesn't panic
	_ = pattern
}

func TestCheckStatisticalAnalysis(t *testing.T) {
	detector := NewDetector()
	features := []Feature{}
	input := []byte("test")

	pattern := detector.checkStatisticalAnalysis(features, input)
	// Pattern may be empty, just verify it doesn't panic
	_ = pattern
}

func TestCheckGradientAnalysis(t *testing.T) {
	detector := NewDetector()
	input := []byte("test")

	pattern := detector.checkGradientAnalysis(input)
	if pattern.Type == "" {
		t.Error("Expected pattern to be returned")
	}
}

func TestCheckFrequencyAnalysis(t *testing.T) {
	detector := NewDetector()
	input := []byte("test")

	pattern := detector.checkFrequencyAnalysis(input)
	if pattern.Type == "" {
		t.Error("Expected pattern to be returned")
	}
}

func TestCheckFeatureConsistency(t *testing.T) {
	detector := NewDetector()
	features := []Feature{}

	pattern := detector.checkFeatureConsistency(features)
	// Pattern may be empty, just verify it doesn't panic
	_ = pattern
}

func TestGetFeature(t *testing.T) {
	detector := NewDetector()
	features := []Feature{{Name: "test", Value: 0.5}}

	feature := detector.getFeature(features, "test")
	if feature.Name != "test" {
		t.Errorf("Expected feature name 'test', got '%s'", feature.Name)
	}
}

func TestEnsembleScore(t *testing.T) {
	results := []*DetectionResult{}

	score := EnsembleScore(results)
	if score < 0 {
		t.Errorf("Expected non-negative score, got %f", score)
	}
}

func TestGenerateReport(t *testing.T) {
	result := &DetectionResult{}

	report := GenerateReport(result)
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

func TestDetectionMethod(t *testing.T) {
	method := DetectionMethod{
		Name:        "Test Method",
		Description: "Test description",
		Threshold:   0.5,
	}

	if method.Name != "Test Method" {
		t.Errorf("Expected name 'Test Method', got '%s'", method.Name)
	}
}

func TestAdversarialPattern(t *testing.T) {
	pattern := AdversarialPattern{
		Type:         "test",
		Description:  "Test description",
		Severity:     "high",
		Confidence:   0.95,
		Evidence:     "evidence",
		Recommendation: "recommendation",
	}

	if pattern.Type != "test" {
		t.Errorf("Expected type 'test', got '%s'", pattern.Type)
	}
}

func TestDetectionResult(t *testing.T) {
	result := DetectionResult{
		IsAdversarial: true,
		Score:         0.85,
		Patterns:      []AdversarialPattern{{Type: "test"}},
		Method:        "Statistical Analysis",
	}

	if !result.IsAdversarial {
		t.Error("Expected IsAdversarial to be true")
	}
}

func TestFeature(t *testing.T) {
	feature := Feature{
		Name:  "test_feature",
		Value: 0.5,
		Weight: 1.0,
	}

	if feature.Name != "test_feature" {
		t.Errorf("Expected name 'test_feature', got '%s'", feature.Name)
	}
}

func TestGetDetectionResult(t *testing.T) {
	result := DetectionResult{
		IsAdversarial: true,
	}

	if !result.IsAdversarial {
		t.Error("Expected IsAdversarial to be true")
	}
}

func TestGetDetectionMethod(t *testing.T) {
	method := DetectionMethod{
		Name: "Test",
	}

	if method.Name != "Test" {
		t.Errorf("Expected name 'Test', got '%s'", method.Name)
	}
}

func TestGetAdversarialPattern(t *testing.T) {
	pattern := AdversarialPattern{
		Type: "test",
	}

	if pattern.Type != "test" {
		t.Errorf("Expected type 'test', got '%s'", pattern.Type)
	}
}