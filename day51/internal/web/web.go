package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"vibe-coding-go/internal/vibe"
)

// PageData represents data for web pages
type PageData struct {
	Title       string
	CurrentTime string
	Examples    []Example
	Exercises   []Exercise
}

// Example represents a code example
type Example struct {
	Title       string
	Description string
	Code        string
	VibeScore   int
	Category    string
}

// Exercise represents a coding exercise
type Exercise struct {
	Title       string
	Description string
	Difficulty  string
	Hint        string
	Solution    string
}

// VibeCheckRequest represents a vibe check request
type VibeCheckRequest struct {
	Code string `json:"code"`
}

// VibeCheckResponse represents a vibe check response
type VibeCheckResponse struct {
	Score       int      `json:"score"`
	Feedback    string   `json:"feedback"`
	Suggestions []string `json:"suggestions"`
}

// HomeHandler handles the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:       "The Zen of Go: Vibe Coding",
		CurrentTime: time.Now().Format("January 2, 2006 15:04:05"),
		Examples:    getExamples(),
		Exercises:   getExercises(),
	}

	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            color: #333;
        }
        
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
        }
        
        .header {
            text-align: center;
            color: white;
            margin-bottom: 40px;
        }
        
        .header h1 {
            font-size: 3rem;
            margin-bottom: 10px;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
        }
        
        .header p {
            font-size: 1.2rem;
            opacity: 0.9;
        }
        
        .nav {
            display: flex;
            justify-content: center;
            gap: 20px;
            margin-bottom: 40px;
        }
        
        .nav a {
            background: rgba(255,255,255,0.2);
            color: white;
            padding: 12px 24px;
            text-decoration: none;
            border-radius: 25px;
            transition: all 0.3s ease;
            backdrop-filter: blur(10px);
        }
        
        .nav a:hover {
            background: rgba(255,255,255,0.3);
            transform: translateY(-2px);
        }
        
        .card {
            background: white;
            border-radius: 15px;
            padding: 30px;
            margin-bottom: 30px;
            box-shadow: 0 10px 30px rgba(0,0,0,0.1);
            transition: transform 0.3s ease;
        }
        
        .card:hover {
            transform: translateY(-5px);
        }
        
        .card h2 {
            color: #667eea;
            margin-bottom: 15px;
            font-size: 1.8rem;
        }
        
        .card h3 {
            color: #764ba2;
            margin-bottom: 10px;
            font-size: 1.3rem;
        }
        
        .code-block {
            background: #f8f9fa;
            border: 1px solid #e9ecef;
            border-radius: 8px;
            padding: 20px;
            margin: 15px 0;
            font-family: 'Courier New', monospace;
            overflow-x: auto;
            position: relative;
        }
        
        .code-block::before {
            content: "💻";
            position: absolute;
            top: 10px;
            right: 15px;
            font-size: 1.2rem;
        }
        
        .vibe-score {
            display: inline-block;
            background: linear-gradient(45deg, #ff6b6b, #feca57);
            color: white;
            padding: 5px 15px;
            border-radius: 20px;
            font-weight: bold;
            margin-left: 10px;
        }
        
        .example-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 20px;
            margin-top: 20px;
        }
        
        .example-card {
            background: #f8f9fa;
            border-radius: 10px;
            padding: 20px;
            border-left: 4px solid #667eea;
        }
        
        .exercise-card {
            background: #fff3cd;
            border-radius: 10px;
            padding: 20px;
            border-left: 4px solid #ffc107;
            margin-bottom: 15px;
        }
        
        .difficulty {
            display: inline-block;
            padding: 4px 12px;
            border-radius: 15px;
            font-size: 0.8rem;
            font-weight: bold;
            margin-left: 10px;
        }
        
        .easy { background: #d4edda; color: #155724; }
        .medium { background: #fff3cd; color: #856404; }
        .hard { background: #f8d7da; color: #721c24; }
        
        .footer {
            text-align: center;
            color: white;
            margin-top: 40px;
            opacity: 0.8;
        }
        
        .vibe-checker {
            background: linear-gradient(45deg, #667eea, #764ba2);
            color: white;
            border-radius: 15px;
            padding: 30px;
            margin: 20px 0;
        }
        
        .vibe-checker textarea {
            width: 100%;
            height: 150px;
            border: none;
            border-radius: 8px;
            padding: 15px;
            font-family: 'Courier New', monospace;
            resize: vertical;
            margin: 15px 0;
        }
        
        .vibe-checker button {
            background: rgba(255,255,255,0.2);
            color: white;
            border: none;
            padding: 12px 24px;
            border-radius: 25px;
            cursor: pointer;
            transition: all 0.3s ease;
        }
        
        .vibe-checker button:hover {
            background: rgba(255,255,255,0.3);
        }
        
        .result {
            background: rgba(255,255,255,0.1);
            border-radius: 8px;
            padding: 20px;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🎯 The Zen of Go: Vibe Coding</h1>
            <p>Day 51: Where Code Meets Intuition</p>
            <p>Current Time: {{.CurrentTime}}</p>
        </div>
        
        <div class="nav">
            <a href="/">🏠 Home</a>
            <a href="/examples">💡 Examples</a>
            <a href="/exercises">🏋️ Exercises</a>
        </div>
        
        <div class="card">
            <h2>🌟 What is Vibe Coding?</h2>
            <p>Vibe coding is the art of writing code that feels natural, intuitive, and flows like poetry. It's about creating software that not only works correctly but also feels right to read, write, and maintain.</p>
            
            <h3>🎨 The Four Pillars of Vibe Coding:</h3>
            <ul>
                <li><strong>Natural Flow:</strong> Code that reads like prose</li>
                <li><strong>Intuitive APIs:</strong> Methods that feel natural to use</li>
                <li><strong>Contextual Clarity:</strong> Code that explains itself</li>
                <li><strong>Emotional Resonance:</strong> Code that feels right</li>
            </ul>
        </div>
        
        <div class="vibe-checker">
            <h2>🔍 Vibe Check Your Code</h2>
            <p>Paste your Go code below to get a vibe analysis:</p>
            <textarea id="codeInput" placeholder="func main() {&#10;    fmt.Println(&quot;Hello, Vibe Coding!&quot;)&#10;}"></textarea>
            <button onclick="checkVibe()">✨ Check Vibe</button>
            <div id="result" class="result" style="display: none;"></div>
        </div>
        
        <div class="card">
            <h2>💡 Quick Examples</h2>
            <div class="example-grid">
                {{range .Examples}}
                <div class="example-card">
                    <h3>{{.Title}} <span class="vibe-score">{{.VibeScore}}/100</span></h3>
                    <p>{{.Description}}</p>
                    <div class="code-block">{{.Code}}</div>
                </div>
                {{end}}
            </div>
        </div>
        
        <div class="card">
            <h2>🏋️ Practice Exercises</h2>
            {{range .Exercises}}
            <div class="exercise-card">
                <h3>{{.Title}} <span class="difficulty {{.Difficulty}}">{{.Difficulty}}</span></h3>
                <p>{{.Description}}</p>
                <details>
                    <summary>💡 Hint</summary>
                    <p>{{.Hint}}</p>
                </details>
                <details>
                    <summary>✅ Solution</summary>
                    <div class="code-block">{{.Solution}}</div>
                </details>
            </div>
            {{end}}
        </div>
        
        <div class="footer">
            <p>🎓 Day 51 of 100 Days of Go Programming | The Zen of Go: Vibe Coding</p>
        </div>
    </div>
    
    <script>
        async function checkVibe() {
            const code = document.getElementById('codeInput').value;
            const resultDiv = document.getElementById('result');
            
            if (!code.trim()) {
                alert('Please enter some code to check!');
                return;
            }
            
            try {
                const response = await fetch('/api/vibe-check', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ code: code })
                });
                
                const data = await response.json();
                
                resultDiv.innerHTML = 
                    '<h3>🎯 Vibe Score: ' + data.score + '/100</h3>' +
                    '<p><strong>Feedback:</strong> ' + data.feedback + '</p>' +
                    (data.suggestions.length > 0 ? 
                        '<h4>💡 Suggestions:</h4>' +
                        '<ul>' +
                        data.suggestions.map(s => '<li>' + s + '</li>').join('') +
                        '</ul>' : '');
                resultDiv.style.display = 'block';
            } catch (error) {
                resultDiv.innerHTML = '<p style="color: #ff6b6b;">Error checking vibe. Please try again.</p>';
                resultDiv.style.display = 'block';
            }
        }
    </script>
</body>
</html>`

	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ExamplesHandler handles the examples page
func ExamplesHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to home with examples section
	http.Redirect(w, r, "/#examples", http.StatusSeeOther)
}

// ExercisesHandler handles the exercises page
func ExercisesHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to home with exercises section
	http.Redirect(w, r, "/#exercises", http.StatusSeeOther)
}

// VibeCheckHandler handles vibe check API requests
func VibeCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req VibeCheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Perform vibe check
	check := vibe.CheckVibe(req.Code)

	response := VibeCheckResponse{
		Score:       check.Score,
		Feedback:    check.Feedback,
		Suggestions: check.Suggestions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getExamples returns example code snippets
func getExamples() []Example {
	return []Example{
		{
			Title:       "Natural Flow",
			Description: "Code that reads like natural language",
			Code: `if user.IsValid() {
    if user.CanReceiveNewsletter() {
        user.SubscribeToNewsletter()
        fmt.Printf("✅ %s subscribed!\n", user.Name)
    }
}`,
			VibeScore: 95,
			Category:  "Flow",
		},
		{
			Title:       "Intuitive APIs",
			Description: "Method chaining that feels natural",
			Code: `result := coder.
    StartCoding().
    WithStyle("elegant").
    FocusOn("readability").
    Deliver()`,
			VibeScore: 90,
			Category:  "API Design",
		},
		{
			Title:       "Contextual Clarity",
			Description: "Code that explains itself through context",
			Code: `project.
    WithDeadline(time.Now().Add(7 * 24 * time.Hour)).
    Requires("clean architecture").
    MustBe("maintainable")`,
			VibeScore: 88,
			Category:  "Clarity",
		},
		{
			Title:       "Emotional Resonance",
			Description: "Code that feels right and connects emotionally",
			Code: `team.
    Welcome("new developers").
    Encourage("creative thinking").
    Celebrate("small wins")`,
			VibeScore: 92,
			Category:  "Resonance",
		},
	}
}

// getExercises returns practice exercises
func getExercises() []Exercise {
	return []Exercise{
		{
			Title:       "Refactor for Natural Flow",
			Description: "Take this rigid code and make it flow naturally",
			Difficulty:  "easy",
			Hint:        "Think about how you would describe the process in natural language",
			Solution: `// Before: Rigid and procedural
func processUser(user User) {
    if validateUser(user) {
        if checkEligibility(user) {
            subscribeUser(user)
        }
    }
}

// After: Natural flow
func (u *User) Process() {
    if u.IsValid() {
        if u.CanSubscribe() {
            u.Subscribe()
        }
    }
}`,
		},
		{
			Title:       "Design Intuitive API",
			Description: "Create a configuration API that feels natural to use",
			Difficulty:  "medium",
			Hint:        "Use method chaining and descriptive method names",
			Solution: `type Config struct {
    host string
    port int
    ssl  bool
}

func NewConfig() *Config {
    return &Config{}
}

func (c *Config) WithHost(host string) *Config {
    c.host = host
    return c
}

func (c *Config) WithPort(port int) *Config {
    c.port = port
    return c
}

func (c *Config) EnableSSL() *Config {
    c.ssl = true
    return c
}

// Usage:
config := NewConfig().
    WithHost("localhost").
    WithPort(8080).
    EnableSSL()`,
		},
		{
			Title:       "Add Contextual Clarity",
			Description: "Make this database operation more contextually clear",
			Difficulty:  "hard",
			Hint:        "Think about what the code is trying to achieve and make that clear",
			Solution: `// Before: Unclear context
func saveUserData(userData map[string]interface{}) error {
    return db.Insert("users", userData)
}

// After: Contextually clear
func (u *User) SaveToDatabase() error {
    return u.db.StoreUser(u.toMap())
}

// Or even better with context:
func (u *User) Register() error {
    if !u.IsValid() {
        return errors.New("invalid user data")
    }
    
    if u.AlreadyExists() {
        return errors.New("user already registered")
    }
    
    return u.db.StoreUser(u.toMap())
}`,
		},
	}
}
