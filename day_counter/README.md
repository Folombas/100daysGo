# 100DaysGo Challenge - Модуль подсчета строк кода

## Обзор

Модуль подсчета строк кода является специализированной частью челленджа 100DaysGo, позволяющей отслеживать и проверять ваши ежедневные вклады в программирование. Он предоставляет подробную информацию о прогрессе в обучении, подсчитывая реальные строки кода Go в указанных каталогах.

## Features

### 📊 Code Line Tracking
- Counts actual Go source lines (`.go` files only)
- Excludes empty lines and comments from count
- Provides detailed breakdown of daily coding efforts
- Cross-platform compatibility with Windows, Linux, and macOS

### 🎯 Gamification Integration
- Links directly to your 100DaysGo challenge progress
- Displays achievements based on line count thresholds
- Encourages consistent daily coding habits
- Visual feedback with emojis and progress indicators

### 🛠️ Technical Capabilities
- Recursive directory traversal
- Error handling for missing directories
- Real-time statistics calculation
- Formatted output with color-coded results

## Usage

### Basic Operation
1. Run the main program: `go run main.go`
2. When prompted, enter a day number in format `dayXX` (e.g., `day25`)
3. View the line count result for that specific day

### Error Handling
- Invalid input format (missing "day" prefix or insufficient digits)
- Non-existent directories
- File reading errors
- Clear error messages with guidance

## Implementation Details

### Functionality
…
## Requirements

- Go 1.21 or higher
- Valid directory structure with daily code repositories
- Properly formatted day directories (e.g., `day01`, `day25`)

## Best Practices

1. Always maintain organized daily folders
2. Commit your code before checking line counts
3. Use descriptive names for your Go files
4. Include meaningful comments in your code
5. Regularly verify your progress through the counter

## Troubleshooting
…
## Author

Gosha (Goша) - Programming Journey

## License

This module is part of the 100DaysGo challenge and follows the same licensing terms as the main project.
