# adversarial - Adversarial Example Detector

[![Go](https://img.shields.io/badge/Go-1.21-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

**Detect and defend against adversarial attacks on machine learning models.**

Identify adversarial examples designed to fool ML models and apply defenses to protect your AI systems.

## 🚀 Features

- **Multi-Method Detection**: Statistical, gradient, frequency, and feature-based analysis
- **Adversarial Pattern Recognition**: Detect common attack patterns (FGSM, PGD, CW)
- **Defense Mechanisms**: Apply multiple defense strategies
- **Ensemble Scoring**: Combine multiple detection methods
- **Real-time Analysis**: Fast detection suitable for production
- **Customizable Thresholds**: Adjust sensitivity to your needs

## 📦 Installation

### Build from Source

```bash
git clone https://github.com/hallucinaut/adversarial.git
cd adversarial
go build -o adversarial ./cmd/adversarial
sudo mv adversarial /usr/local/bin/
```

### Install via Go

```bash
go install github.com/hallucinaut/adversarial/cmd/adversarial@latest
```

## 🎯 Usage

### Detect Adversarial Examples

```bash
# Detect adversarial patterns
adversarial detect image.png

# Analyze for patterns
adversarial analyze model_output.txt
```

### Apply Defenses

```bash
# Defend adversarial input
adversarial defend input.png

# Get defense recommendations
adversarial recommend
```

### Programmatic Usage

```go
package main

import (
    "fmt"
    "github.com/hallucinaut/adversarial/pkg/detect"
    "github.com/hallucinaut/adversarial/pkg/defend"
)

func main() {
    // Create detector
    detector := detect.NewDetector()

    // Detect adversarial examples
    result := detector.Detect(input)

    fmt.Printf("Is Adversarial: %v\n", result.IsAdversarial)
    fmt.Printf("Confidence: %.0f%%\n", result.Score*100)

    // Apply defense
    defender := defend.NewDefender()
    defense := defender.Defend(input, "Input Preprocessing")

    fmt.Printf("Defense Success: %v\n", defense.Success)
    fmt.Printf("Effectiveness: %.0f%%\n", defense.ImprovedScore*100)
}
```

## 🔍 Detection Methods

### Statistical Analysis

Detects unusual statistical patterns in input data:
- Mean intensity anomalies
- Standard deviation deviations
- Entropy irregularities
- Noise level analysis

### Gradient Analysis

Analyzes gradient patterns to detect perturbations:
- Gradient magnitude checks
- Directional consistency
- High-frequency gradients

### Frequency Analysis

Examines frequency domain characteristics:
- Fourier transform analysis
- Spectral pattern detection
- Frequency domain anomalies

### Feature Consistency

Verifies feature consistency:
- Multi-feature correlation
- Distribution consistency
- Feature relationship checks

## 🛡️ Defense Strategies

| Strategy | Effectiveness | Cost | Use Case |
|----------|--------------|------|----------|
| Adversarial Training | 90% | 30% | High-security environments |
| Input Preprocessing | 70% | 10% | Real-time defense |
| Randomization | 60% | 5% | Low-latency requirements |
| Ensemble Defense | 85% | 50% | Maximum security |
| Gradient Masking | 65% | 15% | Attack prevention |

## 📊 Detection Scores

| Score | Status | Action |
|-------|--------|--------|
| 0.0-0.3 | Clean | No action needed |
| 0.3-0.6 | Suspicious | Monitor closely |
| 0.6-0.8 | Likely Adversarial | Apply defense |
| 0.8-1.0 | High Confidence | Block and investigate |

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -v ./pkg/detect -run TestDetectAdversarial
```

## 📋 Example Output

```
Detecting adversarial patterns in: image.png

=== Adversarial Detection Report ===

Is Adversarial: yes
Confidence Score: 78%
Detection Method: ensemble

Detected Patterns:
[1] Gradient Anomaly
    Severity: HIGH
    Confidence: 75%
    Recommendation: Use gradient regularization in defense

[2] Statistical Anomaly
    Severity: MEDIUM
    Confidence: 70%
    Recommendation: Apply denoising or smoothing

⚠️  ADVERSARIAL EXAMPLE DETECTED
Recommendation: Apply defense before processing
```

## 🔒 Security Use Cases

- **AI System Protection**: Protect ML models from adversarial attacks
- **Image Classification Security**: Detect manipulated images
- **Autonomous Vehicle Safety**: Ensure sensor data integrity
- **Medical AI Security**: Protect diagnostic systems
- **Financial AI**: Prevent fraud detection manipulation

## 🛡️ Best Practices

1. **Enable defense by default** for critical systems
2. **Use ensemble detection** for higher accuracy
3. **Monitor detection scores** over time
4. **Regularly update detection models**
5. **Combine with other security measures**

## 🏗️ Architecture

```
adversarial/
├── cmd/
│   └── adversarial/
│       └── main.go          # CLI entry point
├── pkg/
│   ├── detect/
│   │   ├── detect.go        # Detection logic
│   │   └── detect_test.go   # Unit tests
│   └── defend/
│       ├── defend.go        # Defense mechanisms
│       └── defend_test.go   # Unit tests
└── README.md
```

## 📄 License

MIT License

## 🙏 Acknowledgments

- Adversarial machine learning research community
- Defensive AI researchers
- ML security practitioners

## 🔗 Resources

- [Adversarial Machine Learning](https://adversarial-ml-guide.github.io/)
- [TensorFlow Attacks](https://github.com/tensorflow/tf-attacks)
- [Adversarial Robustness Toolbox](https://github.com/Trusted-AI/adversarial-robustness-toolbox)

---

**Built with ❤️ by [hallucinaut](https://github.com/hallucinaut)**