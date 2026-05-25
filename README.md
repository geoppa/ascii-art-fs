# ASCII Art Generator

A simple Go program that converts text input into ASCII art using specific banner styles.

## Features
- Supports three banner styles: `standard`, `shadow`, and `thinkertoy`.
- Handles newline characters (`\n`) within the input string.
- Validates command-line arguments and provides helpful error messages.

## Usage

Run the program using `go run .` followed by your text and the desired banner name.

```bash
go run . "Hello" standard
```

### Banner Options
You can specify the banner with or without the `.txt` extension:
- `standard`
- `shadow`
- `thinkertoy`

### Examples

**Standard Style:**
```bash
go run . "Go" standard
```

**Shadow Style:**
```bash
go run . "Hello" shadow.txt
```

## Behavior Notes
- **Multi-word input:** If you provide multiple words without quotes (e.g., `go run . Hello World standard`), the program will notify you and only process the first word (`Hello`).
- **Missing Banner:** If an invalid or no banner is provided, the program defaults to `standard.txt`.
- **Empty Input:** Providing an empty string will result in no output.

## Installation
Ensure you have [Go](https://go.dev) installed, then clone this repository and navigate to the project folder.

```bash
git clone <repository-url>
cd <project-folder>
```

*Note: This program requires the corresponding `.txt` banner files to be present in the root directory.*
