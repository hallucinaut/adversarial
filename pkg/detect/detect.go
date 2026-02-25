// Package detect provides adversarial example detection capabilities.
package detect

import (
	"math"
)

// AdversarialPattern represents a detected adversarial pattern.
type AdversarialPattern struct {
	Type         string
	Description  string
	Severity     string
	Confidence   float64
	Evidence     string
	Recommendation string
}

// DetectionResult contains detection results.
type DetectionResult struct {
	IsAdversarial bool
	Score         float64
	Patterns      []AdversarialPattern
	Method        string
}

// Detector detects adversarial examples.
type Detector struct {
	methods []DetectionMethod
}

// DetectionMethod defines a detection approach.
type DetectionMethod struct {
	Name        string
	Description string
	Threshold   float64
}

// Feature represents a feature used for detection.
type Feature struct {
	Name    string
	Value   float64
	Weight  float64
}

// ImagePatch represents an image patch for analysis.
type ImagePatch struct {
	X, Y, Width, Height int
	MeanBrightness      float64
	Entropy             float64
	EdgeDensity         float64
	NoiseLevel          float64
}

// NewDetector creates a new adversarial detector.
func NewDetector() *Detector {
	return &Detector{
		methods: []DetectionMethod{
			{Name: "Statistical Analysis", Description: "Statistical pattern detection", Threshold: 0.7},
			{Name: "Gradient Analysis", Description: "Gradient-based detection", Threshold: 0.6},
			{Name: "Frequency Analysis", Description: "Frequency domain analysis", Threshold: 0.65},
			{Name: "Feature Consistency", Description: "Feature consistency check", Threshold: 0.75},
		},
	}
}

// Detect analyzes input for adversarial characteristics.
func (d *Detector) Detect(input []byte) *DetectionResult {
	result := &DetectionResult{
		Method: "ensemble",
	}

	features := d.extractFeatures(input)
	result.Score = d.analyzeFeatures(features)
	result.IsAdversarial = result.Score > 0.7

	for _, method := range d.methods {
		pattern := d.checkMethod(method, features, input)
		if pattern.Confidence > method.Threshold {
			result.Patterns = append(result.Patterns, pattern)
		}
	}

	return result
}

// extractFeatures extracts features from input.
func (d *Detector) extractFeatures(input []byte) []Feature {
	features := make([]Feature, 0, 10)

	// Extract basic statistical features
	mean := d.calculateMean(input)
	stdDev := d.calculateStdDev(input, mean)
	entropy := d.calculateEntropy(input)
	gradient := d.calculateGradient(input)

	features = append(features,
		Feature{Name: "mean_intensity", Value: mean, Weight: 0.3},
		Feature{Name: "std_deviation", Value: stdDev, Weight: 0.25},
		Feature{Name: "entropy", Value: entropy, Weight: 0.2},
		Feature{Name: "gradient_magnitude", Value: gradient, Weight: 0.15},
	)

	// Add more features
	features = append(features,
		Feature{Name: "noise_level", Value: d.calculateNoiseLevel(input), Weight: 0.1},
	)

	return features
}

// calculateMean calculates mean value.
func (d *Detector) calculateMean(input []byte) float64 {
	if len(input) == 0 {
		return 0
	}

	sum := 0.0
	for _, b := range input {
		sum += float64(b)
	}

	return sum / float64(len(input))
}

// calculateStdDev calculates standard deviation.
func (d *Detector) calculateStdDev(input []byte, mean float64) float64 {
	if len(input) == 0 {
		return 0
	}

	sum := 0.0
	for _, b := range input {
		sum += math.Pow(float64(b)-mean, 2)
	}

	return math.Sqrt(sum / float64(len(input)))
}

// calculateEntropy calculates entropy.
func (d *Detector) calculateEntropy(input []byte) float64 {
	// Simplified entropy calculation
	counts := make(map[byte]int)
	for _, b := range input {
		counts[b]++
	}

	entropy := 0.0
	total := float64(len(input))
	for _, count := range counts {
		p := float64(count) / total
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}

	return entropy
}

// calculateGradient calculates gradient magnitude.
func (d *Detector) calculateGradient(input []byte) float64 {
	if len(input) < 2 {
		return 0
	}

	sum := 0.0
	for i := 1; i < len(input); i++ {
		diff := float64(input[i]) - float64(input[i-1])
		sum += math.Abs(diff)
	}

	return sum / float64(len(input)-1)
}

// calculateNoiseLevel calculates noise level.
func (d *Detector) calculateNoiseLevel(input []byte) float64 {
	if len(input) < 3 {
		return 0
	}

	noisyCount := 0
	for i := 1; i < len(input)-1; i++ {
		diff1 := math.Abs(float64(input[i]) - float64(input[i-1]))
		diff2 := math.Abs(float64(input[i+1]) - float64(input[i]))
		if diff1 > 10 || diff2 > 10 {
			noisyCount++
		}
	}

	return float64(noisyCount) / float64(len(input)-2)
}

// analyzeFeatures analyzes extracted features.
func (d *Detector) analyzeFeatures(features []Feature) float64 {
	score := 0.0
	totalWeight := 0.0

	for _, f := range features {
		// Anomalous feature values increase adversarial score
		anomaly := d.calculateAnomalyScore(f)
		score += anomaly * f.Weight
		totalWeight += f.Weight
	}

	if totalWeight > 0 {
		score /= totalWeight
	}

	return math.Min(score, 1.0)
}

// calculateAnomalyScore calculates feature anomaly score.
func (d *Detector) calculateAnomalyScore(f Feature) float64 {
	// Define expected ranges for normal features
	expectedRanges := map[string]struct{ min, max float64 }{
		"mean_intensity":    {0.3, 0.7},
		"std_deviation":     {0.1, 0.5},
		"entropy":           {0.4, 0.9},
		"gradient_magnitude": {0.05, 0.3},
		"noise_level":       {0.0, 0.2},
	}

	expectedRange, exists := expectedRanges[f.Name]
	if !exists {
		return 0.5
	}

	// Calculate how far outside the range the value is
	if f.Value < expectedRange.min {
		return (expectedRange.min - f.Value) / expectedRange.min * 2
	} else if f.Value > expectedRange.max {
		return (f.Value - expectedRange.max) / expectedRange.max * 2
	}

	return 0.0
}

// checkMethod checks for adversarial patterns using a specific method.
func (d *Detector) checkMethod(method DetectionMethod, features []Feature, input []byte) AdversarialPattern {
	switch method.Name {
	case "Statistical Analysis":
		return d.checkStatisticalAnalysis(features, input)
	case "Gradient Analysis":
		return d.checkGradientAnalysis(input)
	case "Frequency Analysis":
		return d.checkFrequencyAnalysis(input)
	case "Feature Consistency":
		return d.checkFeatureConsistency(features)
	default:
		return AdversarialPattern{}
	}
}

// checkStatisticalAnalysis performs statistical analysis.
func (d *Detector) checkStatisticalAnalysis(features []Feature, input []byte) AdversarialPattern {
	stdDevFeature := d.getFeature(features, "std_deviation")
	noiseFeature := d.getFeature(features, "noise_level")

	if stdDevFeature.Value > 0.5 && noiseFeature.Value > 0.1 {
		return AdversarialPattern{
			Type:         "Statistical Anomaly",
			Description:  "Unusual statistical patterns detected",
			Severity:     "MEDIUM",
			Confidence:   0.7,
			Evidence:     "High std dev and noise level",
			Recommendation: "Apply denoising or smoothing",
		}
	}

	return AdversarialPattern{}
}

// checkGradientAnalysis performs gradient analysis.
func (d *Detector) checkGradientAnalysis(input []byte) AdversarialPattern {
	gradient := d.calculateGradient(input)

	if gradient > 0.25 {
		return AdversarialPattern{
			Type:         "Gradient Anomaly",
			Description:  "High gradient magnitude suggests adversarial perturbation",
			Severity:     "HIGH",
			Confidence:   0.75,
			Evidence:     "Elevated gradient magnitude",
			Recommendation: "Use gradient regularization in defense",
		}
	}

	return AdversarialPattern{}
}

// checkFrequencyAnalysis performs frequency analysis.
func (d *Detector) checkFrequencyAnalysis(input []byte) AdversarialPattern {
	// Simplified frequency analysis
	entropy := d.calculateEntropy(input)

	if entropy > 0.85 {
		return AdversarialPattern{
			Type:         "Frequency Anomaly",
			Description:  "High entropy in frequency domain",
			Severity:     "MEDIUM",
			Confidence:   0.65,
			Evidence:     "Unusual frequency patterns",
			Recommendation: "Apply frequency-domain filtering",
		}
	}

	return AdversarialPattern{}
}

// checkFeatureConsistency performs feature consistency check.
func (d *Detector) checkFeatureConsistency(features []Feature) AdversarialPattern {
	entropy := d.getFeature(features, "entropy")

	if entropy.Value > 0.9 {
		return AdversarialPattern{
			Type:         "Feature Inconsistency",
			Description:  "Feature distribution inconsistency detected",
			Severity:     "LOW",
			Confidence:   0.6,
			Evidence:     "Abnormal entropy value",
			Recommendation: "Verify input integrity",
		}
	}

	return AdversarialPattern{}
}

// getFeature retrieves a feature by name.
func (d *Detector) getFeature(features []Feature, name string) Feature {
	for _, f := range features {
		if f.Name == name {
			return f
		}
	}
	return Feature{Name: name, Value: 0}
}

// EnsembleScore combines multiple detection methods.
func EnsembleScore(results []*DetectionResult) float64 {
	if len(results) == 0 {
		return 0
	}

	score := 0.0
	for _, result := range results {
		score += result.Score
	}

	return score / float64(len(results))
}

// GenerateReport generates detection report.
func GenerateReport(result *DetectionResult) string {
	var report string

	report += "=== Adversarial Detection Report ===\n\n"
	report += "Is Adversarial: " + boolToString(result.IsAdversarial) + "\n"
	report += "Confidence Score: " + string(rune(int(result.Score*100)+48)) + "%\n"
	report += "Detection Method: " + result.Method + "\n\n"

	if len(result.Patterns) > 0 {
		report += "Detected Patterns:\n"
		for i, pattern := range result.Patterns {
			report += "[" + string(rune(i+49)) + "] " + pattern.Type + "\n"
			report += "    Severity: " + pattern.Severity + "\n"
			report += "    Confidence: " + string(rune(int(pattern.Confidence*100)+48)) + "%\n"
			report += "    Recommendation: " + pattern.Recommendation + "\n\n"
		}
	}

	return report
}

// boolToString converts bool to string.
func boolToString(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}