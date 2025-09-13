package main

import (
	"fmt"
	"log"
	"net/http"

	"vibe-coding-go/internal/vibe"
	"vibe-coding-go/internal/web"
)

// Lesson Title: "The Zen of Go: Vibe Coding - Where Code Meets Intuition"
// Day 51: Understanding the art of writing code that feels natural and flows

func main() {
	fmt.Println("🎯 Day 51: The Zen of Go - Vibe Coding")
	fmt.Println("=====================================")
	fmt.Println()

	// Demonstrate vibe coding concepts
	demonstrateVibeCoding()

	// Start the web interface
	startWebInterface()
}

func demonstrateVibeCoding() {
	fmt.Println("🌟 Vibe Coding Demo - The Art of Intuitive Code")
	fmt.Println("===============================================")
	fmt.Println()

	// Example 1: Natural Flow
	fmt.Println("1. Natural Flow - Code that reads like prose:")
	vibe.DemonstrateNaturalFlow()

	fmt.Println()

	// Example 2: Intuitive APIs
	fmt.Println("2. Intuitive APIs - Methods that feel natural:")
	vibe.DemonstrateIntuitiveAPIs()

	fmt.Println()

	// Example 3: Contextual Clarity
	fmt.Println("3. Contextual Clarity - Code that explains itself:")
	vibe.DemonstrateContextualClarity()

	fmt.Println()

	// Example 4: Emotional Resonance
	fmt.Println("4. Emotional Resonance - Code that feels right:")
	vibe.DemonstrateEmotionalResonance()

	fmt.Println()
}

func startWebInterface() {
	fmt.Println("🌐 Starting Web Interface...")
	fmt.Println("Visit: http://localhost:8080")
	fmt.Println()

	// Setup routes
	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/examples", web.ExamplesHandler)
	http.HandleFunc("/exercises", web.ExercisesHandler)
	http.HandleFunc("/api/vibe-check", web.VibeCheckHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
