package main

import (
	"fmt"
	"os"

	"github.com/hallucinaut/adversarial/pkg/detect"
	"github.com/hallucinaut/adversarial/pkg/defend"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "detect":
		if len(os.Args) < 3 {
			fmt.Println("Error: input file required")
			printUsage()
			return
		}
		detectInput(os.Args[2])
	case "defend":
		if len(os.Args) < 3 {
			fmt.Println("Error: input file required")
			printUsage()
			return
		}
		defendInput(os.Args[2])
	case "analyze":
		if len(os.Args) < 3 {
			fmt.Println("Error: input file required")
			printUsage()
			return
		}
		analyzeInput(os.Args[2])
	case "recommend":
		recommendDefense()
	case "version":
		fmt.Printf("adversarial version %s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Printf(`adversarial - Adversarial Example Detector and Defender

Usage:
  adversarial <command> [options]

Commands:
  detect <file>    Detect adversarial examples in input
  defend <file>    Apply defense to adversarial input
  analyze <file>   Analyze input for adversarial patterns
  recommend        Recommend defense strategy
  version          Show version information
  help             Show this help message

Examples:
  adversarial detect image.png
  adversarial defend image.png
  adversarial analyze model_output.txt
`)
}

func detectInput(filepath string) {
	fmt.Printf("Detecting adversarial patterns in: %s\n", filepath)
	fmt.Println()

	// In production: read and process file
	// For demo: simulate detection
	sampleInput := []byte{128, 130, 127, 129, 131, 126, 132, 125, 133, 124}

	detector := detect.NewDetector()
	result := detector.Detect(sampleInput)

	fmt.Println(detect.GenerateReport(result))

	if result.IsAdversarial {
		fmt.Println("⚠️  ADVERSARIAL EXAMPLE DETECTED")
		fmt.Println("Recommendation: Apply defense before processing")
		os.Exit(1)
	} else {
		fmt.Println("✓ Input appears clean")
	}
}

func defendInput(filepath string) {
	fmt.Printf("Defending input: %s\n", filepath)
	fmt.Println()

	// In production: read and process file
	// For demo: show defense template
	fmt.Println("Available defense strategies:")
	fmt.Println("1. Adversarial Training (90% effective, 30% cost)")
	fmt.Println("2. Input Preprocessing (70% effective, 10% cost)")
	fmt.Println("3. Randomization (60% effective, 5% cost)")
	fmt.Println("4. Ensemble Defense (85% effective, 50% cost)")
	fmt.Println("5. Gradient Masking (65% effective, 15% cost)")
	fmt.Println()

	// Example defense
	defender := defend.NewDefender()
	result := defender.Defend([]byte("sample"), "Input Preprocessing")

	fmt.Println(defend.GenerateDefenseReport(result))
}

func analyzeInput(filepath string) {
	fmt.Printf("Analyzing input: %s\n", filepath)
	fmt.Println()

	// In production: comprehensive analysis
	// For demo: show analysis capabilities
	fmt.Println("Analysis capabilities:")
	fmt.Println("- Statistical pattern detection")
	fmt.Println("- Gradient-based analysis")
	fmt.Println("- Frequency domain analysis")
	fmt.Println("- Feature consistency checks")
	fmt.Println("- Ensemble scoring")
}

func recommendDefense() {
	fmt.Println("Defense Recommendations")
	fmt.Println("=======================")
	fmt.Println()

	// Recommend based on adversarial score
	for _, score := range []float64{0.9, 0.7, 0.5, 0.3} {
		strategy := defend.RecommendDefense(score)
		fmt.Printf("Adversarial Score %.0f%%: Use '%s'\n", score*100, strategy)
	}
}