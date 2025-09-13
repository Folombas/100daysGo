package vibe

import (
	"strings"
	"time"
)

// VibeCodingPrinciples demonstrates the core principles of vibe coding
type VibeCodingPrinciples struct {
	NaturalFlow        bool
	IntuitiveAPIs      bool
	ContextualClarity  bool
	EmotionalResonance bool
}

// AdvancedVibeExamples provides more complex examples of vibe coding
type AdvancedVibeExamples struct {
	Title       string
	Description string
	BeforeCode  string
	AfterCode   string
	Improvement string
}

// GetAdvancedExamples returns advanced vibe coding examples
func GetAdvancedExamples() []AdvancedVibeExamples {
	return []AdvancedVibeExamples{
		{
			Title:       "Database Operations",
			Description: "Making database operations feel natural",
			BeforeCode: `func createUser(userData map[string]interface{}) error {
    query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
    return db.Exec(query, userData["name"], userData["email"], userData["age"])
}`,
			AfterCode: `func (u *User) Create() error {
    if !u.IsValid() {
        return ErrInvalidUser
    }
    
    if u.AlreadyExists() {
        return ErrUserExists
    }
    
    return u.db.Store(u)
}`,
			Improvement: "Method belongs to User, clear validation, descriptive errors",
		},
		{
			Title:       "Configuration Management",
			Description: "Intuitive configuration setup",
			BeforeCode: `func setupConfig() *Config {
    config := &Config{}
    config.SetHost("localhost")
    config.SetPort(8080)
    config.SetSSL(true)
    config.SetTimeout(30)
    return config
}`,
			AfterCode: `func NewServerConfig() *Config {
    return NewConfig().
        WithHost("localhost").
        WithPort(8080).
        EnableSSL().
        WithTimeout(30 * time.Second)
}`,
			Improvement: "Fluent interface, descriptive factory method, natural chaining",
		},
		{
			Title:       "Error Handling",
			Description: "Making error handling feel natural",
			BeforeCode: `func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    data, err := ioutil.ReadAll(file)
    if err != nil {
        return err
    }
    
    return processData(data)
}`,
			AfterCode: `func (p *Processor) ProcessFile(filename string) error {
    return p.
        OpenFile(filename).
        ReadContent().
        ProcessData().
        HandleErrors()
}`,
			Improvement: "Method chaining, encapsulated error handling, clear flow",
		},
	}
}

// VibeAssessment provides tools for assessing code vibe
type VibeAssessment struct {
	Readability     int
	Maintainability int
	Testability     int
	Documentation   int
	OverallVibe     int
}

// AssessCodeVibe evaluates code for vibe quality
func AssessCodeVibe(code string) VibeAssessment {
	assessment := VibeAssessment{}

	// Readability assessment
	if strings.Contains(code, "func") {
		assessment.Readability += 20
	}
	if strings.Contains(code, "if") && strings.Contains(code, "else") {
		assessment.Readability += 15
	}
	if strings.Contains(code, "//") {
		assessment.Readability += 10
	}
	if strings.Contains(code, "return") {
		assessment.Readability += 10
	}

	// Maintainability assessment
	if strings.Contains(code, "error") {
		assessment.Maintainability += 20
	}
	if strings.Contains(code, "defer") {
		assessment.Maintainability += 15
	}
	if strings.Contains(code, "interface") {
		assessment.Maintainability += 15
	}
	if strings.Contains(code, "struct") {
		assessment.Maintainability += 10
	}

	// Testability assessment
	if strings.Contains(code, "func") && strings.Contains(code, "return") {
		assessment.Testability += 25
	}
	if strings.Contains(code, "interface") {
		assessment.Testability += 20
	}
	if strings.Contains(code, "error") {
		assessment.Testability += 15
	}

	// Documentation assessment
	commentLines := strings.Count(code, "//")
	if commentLines > 0 {
		assessment.Documentation = min(commentLines*5, 50)
	}

	// Calculate overall vibe
	assessment.OverallVibe = (assessment.Readability +
		assessment.Maintainability +
		assessment.Testability +
		assessment.Documentation) / 4

	return assessment
}

// VibeCodingTips provides practical tips for vibe coding
type VibeCodingTips struct {
	Category string
	Tip      string
	Example  string
}

// GetVibeCodingTips returns practical tips for better vibe coding
func GetVibeCodingTips() []VibeCodingTips {
	return []VibeCodingTips{
		{
			Category: "Naming",
			Tip:      "Use verbs for methods, nouns for variables",
			Example:  "user.Validate() instead of user.Validation()",
		},
		{
			Category: "Flow",
			Tip:      "Make code read like a story",
			Example:  "if user.IsValid() { user.Process() }",
		},
		{
			Category: "APIs",
			Tip:      "Design for the happy path first",
			Example:  "config.WithHost().WithPort().Build()",
		},
		{
			Category: "Context",
			Tip:      "Let the context tell the story",
			Example:  "project.WithDeadline().Requires().MustBe()",
		},
		{
			Category: "Errors",
			Tip:      "Make errors descriptive and actionable",
			Example:  "ErrInvalidUser instead of ErrValidation",
		},
		{
			Category: "Comments",
			Tip:      "Explain why, not what",
			Example:  "// Retry on network timeout to handle flaky connections",
		},
	}
}

// VibeCodingExercises provides hands-on exercises
type VibeCodingExercises struct {
	Title       string
	Description string
	StarterCode string
	Hint        string
	Solution    string
	Difficulty  string
}

// GetVibeCodingExercises returns practical exercises
func GetVibeCodingExercises() []VibeCodingExercises {
	return []VibeCodingExercises{
		{
			Title:       "Refactor User Management",
			Description: "Transform procedural user management into vibe-coded methods",
			StarterCode: `func createUser(name, email string, age int) error {
    if name == "" || email == "" || age < 0 {
        return errors.New("invalid input")
    }
    
    user := map[string]interface{}{
        "name": name,
        "email": email,
        "age": age,
    }
    
    return db.Insert("users", user)
}`,
			Hint: "Create a User struct with methods that feel natural",
			Solution: `type User struct {
    Name  string
    Email string
    Age   int
    db    Database
}

func (u *User) Create() error {
    if !u.IsValid() {
        return ErrInvalidUser
    }
    
    if u.AlreadyExists() {
        return ErrUserExists
    }
    
    return u.db.Store(u)
}

func (u *User) IsValid() bool {
    return u.Name != "" && u.Email != "" && u.Age >= 0
}

func (u *User) AlreadyExists() bool {
    // Implementation details...
    return false
}`,
			Difficulty: "easy",
		},
		{
			Title:       "Design Fluent Configuration",
			Description: "Create a configuration API that flows naturally",
			StarterCode: `func setupServer() {
    config := &ServerConfig{}
    config.Host = "localhost"
    config.Port = 8080
    config.SSL = true
    config.Timeout = 30
    
    server := NewServer(config)
    server.Start()
}`,
			Hint: "Use method chaining and builder pattern",
			Solution: `func NewServer() *Server {
    return NewServer().
        WithHost("localhost").
        WithPort(8080).
        EnableSSL().
        WithTimeout(30 * time.Second).
        Start()
}

type Server struct {
    config *Config
}

func (s *Server) WithHost(host string) *Server {
    s.config.Host = host
    return s
}

func (s *Server) WithPort(port int) *Server {
    s.config.Port = port
    return s
}

func (s *Server) EnableSSL() *Server {
    s.config.SSL = true
    return s
}

func (s *Server) WithTimeout(timeout time.Duration) *Server {
    s.config.Timeout = timeout
    return s
}

func (s *Server) Start() *Server {
    // Start server logic
    return s
}`,
			Difficulty: "medium",
		},
		{
			Title:       "Create Intuitive Data Pipeline",
			Description: "Design a data processing pipeline that reads naturally",
			StarterCode: `func processData(input []string) []string {
    var result []string
    
    for _, item := range input {
        if len(item) > 0 {
            processed := strings.ToUpper(item)
            if len(processed) > 3 {
                result = append(result, processed)
            }
        }
    }
    
    return result
}`,
			Hint: "Think about how you would describe the process step by step",
			Solution: `type DataPipeline struct {
    data []string
}

func NewPipeline(input []string) *DataPipeline {
    return &DataPipeline{data: input}
}

func (p *DataPipeline) FilterEmpty() *DataPipeline {
    var filtered []string
    for _, item := range p.data {
        if item != "" {
            filtered = append(filtered, item)
        }
    }
    p.data = filtered
    return p
}

func (p *DataPipeline) TransformToUpper() *DataPipeline {
    for i, item := range p.data {
        p.data[i] = strings.ToUpper(item)
    }
    return p
}

func (p *DataPipeline) FilterByLength(minLength int) *DataPipeline {
    var filtered []string
    for _, item := range p.data {
        if len(item) >= minLength {
            filtered = append(filtered, item)
        }
    }
    p.data = filtered
    return p
}

func (p *DataPipeline) Result() []string {
    return p.data
}

// Usage:
result := NewPipeline(input).
    FilterEmpty().
    TransformToUpper().
    FilterByLength(3).
    Result()`,
			Difficulty: "hard",
		},
	}
}

// VibeCodingBestPractices provides best practices for vibe coding
type VibeCodingBestPractices struct {
	Practice string
	Reason   string
	Example  string
}

// GetVibeCodingBestPractices returns best practices
func GetVibeCodingBestPractices() []VibeCodingBestPractices {
	return []VibeCodingBestPractices{
		{
			Practice: "Use descriptive method names",
			Reason:   "Method names should clearly indicate what they do",
			Example:  "user.CanReceiveNewsletter() instead of user.Check()",
		},
		{
			Practice: "Design for the happy path",
			Reason:   "Make the common case simple and elegant",
			Example:  "config.WithHost().WithPort().Build()",
		},
		{
			Practice: "Use method chaining for related operations",
			Reason:   "Creates a fluent interface that feels natural",
			Example:  "query.Where().OrderBy().Limit().Execute()",
		},
		{
			Practice: "Make errors descriptive and actionable",
			Reason:   "Help developers understand what went wrong",
			Example:  "ErrInvalidEmailFormat instead of ErrValidation",
		},
		{
			Practice: "Use context to provide clarity",
			Reason:   "Let the surrounding code explain the intent",
			Example:  "project.WithDeadline().Requires().MustBe()",
		},
		{
			Practice: "Prefer composition over inheritance",
			Reason:   "Creates more flexible and maintainable code",
			Example:  "type User struct { Name string; Email string }",
		},
	}
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// VibeCodingWorkshop provides a complete workshop experience
type VibeCodingWorkshop struct {
	Title            string
	Description      string
	Duration         time.Duration
	Prerequisites    []string
	LearningOutcomes []string
	Exercises        []VibeCodingExercises
}

// GetVibeCodingWorkshop returns a complete workshop
func GetVibeCodingWorkshop() VibeCodingWorkshop {
	return VibeCodingWorkshop{
		Title:       "The Art of Vibe Coding in Go",
		Description: "Learn to write code that feels natural, intuitive, and flows like poetry",
		Duration:    2 * time.Hour,
		Prerequisites: []string{
			"Basic Go knowledge",
			"Understanding of structs and methods",
			"Familiarity with interfaces",
		},
		LearningOutcomes: []string{
			"Write code that reads like natural language",
			"Design intuitive APIs that feel natural to use",
			"Create code with contextual clarity",
			"Develop emotional resonance in your code",
			"Apply vibe coding principles to real projects",
		},
		Exercises: GetVibeCodingExercises(),
	}
}
