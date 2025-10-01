module example.com/demo-app

go 1.21

require (
	github.com/pkg/errors v0.9.1
	// Эти зависимости реально используются
	golang.org/x/text v0.3.7
)

replace local/module => ./local
