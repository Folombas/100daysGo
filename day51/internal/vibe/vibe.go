package vibe

import (
	"fmt"
	"strings"
	"time"
)

// VibeCoder represents someone who writes code with natural flow and intuition
type VibeCoder struct {
	Name     string
	Style    string
	Approach string
}

// CodeFlow represents the natural flow of code
type CodeFlow struct {
	Readable     bool
	Intuitive    bool
	Maintainable bool
	Elegant      bool
}

// VibeCheck represents the assessment of code's vibe
type VibeCheck struct {
	Score       int
	Feedback    string
	Suggestions []string
}

// DemonstrateNaturalFlow shows how code can read like natural language
func DemonstrateNaturalFlow() {
	fmt.Println("   📖 Code that reads like prose:")
	fmt.Println("   ------------------------------")

	// Example: Processing user data
	user := User{Name: "Alice", Email: "alice@example.com", Age: 28}

	// Vibe coding: Natural flow
	if user.IsValid() {
		if user.CanReceiveNewsletter() {
			user.SubscribeToNewsletter()
			fmt.Printf("   ✅ %s is now subscribed to our newsletter!\n", user.Name)
		}
	}

	fmt.Println()
	fmt.Println("   🔄 Compare with non-vibe coding:")
	fmt.Println("   if validateUser(user) && checkNewsletterEligibility(user) {")
	fmt.Println("       subscribeUserToNewsletter(user)")
	fmt.Println("   }")
}

// DemonstrateIntuitiveAPIs shows APIs that feel natural to use
func DemonstrateIntuitiveAPIs() {
	fmt.Println("   🎯 APIs that feel natural:")
	fmt.Println("   -------------------------")

	// Create a vibe coder
	coder := NewVibeCoder("Alex", "Zen", "Flow")

	// Natural method chaining
	result := coder.
		StartCoding().
		WithStyle("elegant").
		FocusOn("readability").
		Deliver()

	fmt.Printf("   🎨 %s\n", result)

	fmt.Println()
	fmt.Println("   🔄 Compare with rigid APIs:")
	fmt.Println("   coder.StartCoding()")
	fmt.Println("   coder.SetStyle(\"elegant\")")
	fmt.Println("   coder.SetFocus(\"readability\")")
	fmt.Println("   result := coder.Deliver()")
}

// DemonstrateContextualClarity shows code that explains itself
func DemonstrateContextualClarity() {
	fmt.Println("   💡 Code that explains itself:")
	fmt.Println("   -----------------------------")

	// Vibe coding: Context makes everything clear
	project := NewProject("VibeCodingApp")

	// The context tells the story
	project.
		WithDeadline(time.Now().Add(7 * 24 * time.Hour)).
		Requires("clean architecture").
		MustBe("maintainable").
		ShouldFeel("intuitive")

	fmt.Printf("   📋 Project: %s\n", project.Name)
	fmt.Printf("   ⏰ Deadline: %s\n", project.Deadline.Format("Jan 2, 2006"))
	fmt.Printf("   🎯 Requirements: %s\n", strings.Join(project.Requirements, ", "))

	fmt.Println()
	fmt.Println("   🔄 Compare with unclear context:")
	fmt.Println("   project := Project{Name: \"VibeCodingApp\", Deadline: time.Now().Add(7*24*time.Hour)}")
	fmt.Println("   project.AddRequirement(\"clean architecture\")")
	fmt.Println("   project.AddRequirement(\"maintainable\")")
	fmt.Println("   project.AddRequirement(\"intuitive\")")
}

// DemonstrateEmotionalResonance shows code that feels right
func DemonstrateEmotionalResonance() {
	fmt.Println("   ❤️ Code that feels right:")
	fmt.Println("   ------------------------")

	// Vibe coding: Emotional connection through code
	team := NewTeam("VibeCoders")

	// This feels good to write and read
	team.
		Welcome("new developers").
		Encourage("creative thinking").
		Celebrate("small wins").
		Build("amazing software")

	fmt.Printf("   👥 Team: %s\n", team.Name)
	fmt.Printf("   🌟 Culture: %s\n", strings.Join(team.Culture, " → "))

	fmt.Println()
	fmt.Println("   🔄 Compare with emotionless code:")
	fmt.Println("   team := Team{Name: \"VibeCoders\"}")
	fmt.Println("   team.AddMember(\"new developers\")")
	fmt.Println("   team.SetPolicy(\"creative thinking\")")
	fmt.Println("   team.SetPolicy(\"celebrate wins\")")
	fmt.Println("   team.SetGoal(\"build software\")")
}

// User represents a user with vibe-coded methods
type User struct {
	Name  string
	Email string
	Age   int
}

// IsValid checks if user data is valid (natural method name)
func (u User) IsValid() bool {
	return u.Name != "" && strings.Contains(u.Email, "@") && u.Age > 0
}

// CanReceiveNewsletter checks if user can receive newsletters (clear intent)
func (u User) CanReceiveNewsletter() bool {
	return u.Age >= 18 && u.Email != ""
}

// SubscribeToNewsletter subscribes user to newsletter (action-oriented)
func (u *User) SubscribeToNewsletter() {
	// In real implementation, this would update database
	fmt.Printf("   📧 Newsletter subscription activated for %s\n", u.Name)
}

// VibeCoder methods for natural API design
func NewVibeCoder(name, style, approach string) *VibeCoder {
	return &VibeCoder{
		Name:     name,
		Style:    style,
		Approach: approach,
	}
}

func (v *VibeCoder) StartCoding() *VibeCoder {
	fmt.Printf("   🚀 %s starts coding with %s style\n", v.Name, v.Style)
	return v
}

func (v *VibeCoder) WithStyle(style string) *VibeCoder {
	v.Style = style
	fmt.Printf("   🎨 Style set to: %s\n", style)
	return v
}

func (v *VibeCoder) FocusOn(focus string) *VibeCoder {
	fmt.Printf("   🎯 Focusing on: %s\n", focus)
	return v
}

func (v *VibeCoder) Deliver() string {
	return fmt.Sprintf("Code delivered with %s approach and %s style", v.Approach, v.Style)
}

// Project represents a project with contextual clarity
type Project struct {
	Name         string
	Deadline     time.Time
	Requirements []string
}

func NewProject(name string) *Project {
	return &Project{
		Name:         name,
		Requirements: make([]string, 0),
	}
}

func (p *Project) WithDeadline(deadline time.Time) *Project {
	p.Deadline = deadline
	return p
}

func (p *Project) Requires(requirement string) *Project {
	p.Requirements = append(p.Requirements, requirement)
	return p
}

func (p *Project) MustBe(requirement string) *Project {
	p.Requirements = append(p.Requirements, requirement)
	return p
}

func (p *Project) ShouldFeel(requirement string) *Project {
	p.Requirements = append(p.Requirements, requirement)
	return p
}

// Team represents a team with emotional resonance
type Team struct {
	Name    string
	Culture []string
}

func NewTeam(name string) *Team {
	return &Team{
		Name:    name,
		Culture: make([]string, 0),
	}
}

func (t *Team) Welcome(who string) *Team {
	t.Culture = append(t.Culture, fmt.Sprintf("Welcome %s", who))
	return t
}

func (t *Team) Encourage(what string) *Team {
	t.Culture = append(t.Culture, fmt.Sprintf("Encourage %s", what))
	return t
}

func (t *Team) Celebrate(what string) *Team {
	t.Culture = append(t.Culture, fmt.Sprintf("Celebrate %s", what))
	return t
}

func (t *Team) Build(what string) *Team {
	t.Culture = append(t.Culture, fmt.Sprintf("Build %s", what))
	return t
}

// CheckVibe analyzes code for its vibe quality
func CheckVibe(code string) VibeCheck {
	score := 0
	feedback := ""
	suggestions := []string{}

	// Simple vibe analysis
	if strings.Contains(code, "if") && strings.Contains(code, "then") {
		score += 20
		feedback += "Good conditional flow. "
	}

	if strings.Contains(code, "func") && strings.Contains(code, "return") {
		score += 20
		feedback += "Clear function structure. "
	}

	if strings.Contains(code, "error") {
		score += 15
		feedback += "Error handling present. "
	} else {
		suggestions = append(suggestions, "Consider adding error handling")
	}

	if len(code) < 100 {
		score += 10
		feedback += "Concise code. "
	}

	if strings.Contains(code, "//") {
		score += 10
		feedback += "Good documentation. "
	} else {
		suggestions = append(suggestions, "Add comments for clarity")
	}

	if score < 50 {
		feedback += "Code could benefit from more vibe!"
		suggestions = append(suggestions, "Focus on readability and natural flow")
		suggestions = append(suggestions, "Use descriptive variable and function names")
	}

	return VibeCheck{
		Score:       score,
		Feedback:    feedback,
		Suggestions: suggestions,
	}
}
