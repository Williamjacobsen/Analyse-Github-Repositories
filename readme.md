# Github Repository Codebase Analyzer

This Go project analyzes the contents of a GitHub repository by concurrently crawling directories, identifying files with specific extensions, counting lines of code in those files, and saving the results in a structured JSON format.

## Features

- Concurrent traversal of GitHub repository directories
- Configurable file extension filtering (`fileExtensions.txt`)
- Line counting for selected file types
- JSON output for file line analysis
- Logging for time measurements and progress tracking

## How It Works

1. **Fetch Directory Tree**: The program starts at a GitHub repo URL and retrieves all directories and files using concurrent workers.
2. **Filter Files**: Only files with extensions listed in `fileExtensions.txt` are analyzed.
3. **Count Lines**: Each valid file is fetched and its line count is calculated.
4. **Output**: A JSON file (`analysis_result.json`) is created containing per-file line counts and the total.

## Configuration

Set the following constants in `main.go`:

```go
const (
	URL                  = "https://github.com/your/repo/tree/main"
	BRANCH               = "main"
	LOGGING              = true
	SAVE_RESULTS_TO_FILE = true
	NUMBER_OF_WORKERS    = 10
)
```

Create a **fileExtensions.txt** with each extension on a new line:

```go
.go
.py
.java
```

## Output

The result is saved as **analysis_result.json**:

```json
{
  "files and their lines": {
    "https://raw.githubusercontent.com/your/repo/file1.java": 856,
    "https://raw.githubusercontent.com/your/repo/file2.java": 428,
    "https://raw.githubusercontent.com/your/repo/file3.py": 643
  },
  "lines per language": [
    {
      "extension": ".java",
      "lines": 1284
    },
    {
      "extension": ".py",
      "lines": 643
    }
  ],
  "total amount of lines": 1927
}
```

## Requirements

- Go 1.18+
- Internet access to reach GitHub URLs
- **fileExtensions.txt** in the root directory

## Run

```bash
go run .
```

Or

```bash
main.exe
```
