# ASCII Art Generator

A simple Go command-line tool that converts text strings into stylized ASCII art banners. It supports multiple fonts/styles by reading character maps from template files and includes smart argument validation.

## Features

- **Multiple Banner Styles**: Supports `standard`, `shadow`, and `thinkertoy` layouts.
- **Flexible Extensions**: Accepts banner names both with or without the `.txt` extension (e.g., `standard` or `standard.txt`).
- **Robust Argument Parsing**: 
  - Automatically falls back to `standard.txt` if an invalid banner is provided.
  - Warns the user and safely ignores any extra or ungrouped arguments trailing after the banner name.
- **Newline Support**: Properly interprets literal `\n` characters in your input string to print multi-line ASCII art.

## Requirements

- [Go](https://go.dev) (version 1.16 or higher recommended)
- Banner template files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) placed in the root directory of the project.

## Installation

1. Clone or download this repository to your local machine.
2. Ensure your banner template files are inside the project folder:
   ```bash
   .
   ├── main.go
   ├── standard.txt
   ├── shadow.txt
   └── thinkertoy.txt
   ```

## Usage

Run the program via the terminal by passing your text string and the desired banner style as arguments:

```bash
go run . "<text>" <banner>
```

### Examples

**Basic Usage:**
```bash
go run . "Hello" standard
```

**Using literal newlines:**
```bash
go run . "Hello\nWorld" shadow
```

### Argument Edge Cases & Warnings

- **Missing/Invalid Banner**: If the banner is omitted or typed incorrectly, the program defaults to `standard.txt` and prints a warning:
  ```bash
  $ go run . "Hello" wrong_banner
  Error: Invalid or missing banner. Only standard.txt, shadow.txt, and thinkertoy.txt are allowed. Using standard.txt instead.
  ```

- **Extra Arguments**: If you pass additional arguments after the banner, the program processes your requested text and prints a warning while ignoring the rest:
  ```bash
  $ go run . "Hello" standard extra_arg1 extra_arg2
  Warning: Extra arguments found after the banner. Ignoring them.
  [ASCII Art for "Hello" prints here]
  ```

## How It Works

1. **Validation**: The program scans `os.Args` to isolate your text input and match a valid banner template.
2. **File Reading**: It reads the selected `.txt` file, where each character block consists of 9 vertical lines (8 lines of art + 1 empty spacing line).
3. **Calculation**: It computes the exact starting line for each character using its ASCII value offset (`(char - 32) * 9 + 1`).
4. **Rendering**: It prints the text layer-by-layer horizontally to output the complete block graphic safely.
