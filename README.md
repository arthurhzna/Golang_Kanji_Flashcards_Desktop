# Japanese Flashcard Application

A simple flashcard application for learning Japanese vocabulary with a GUI interface using Fyne.

## Features

- 📚 Display Japanese vocabulary (Kanji, Hiragana, Reading, Meaning)
- ⏭️ Navigate forward/backward between cards
- 👁️ Toggle to hide/show readings and meanings
- 📱 Simple and user-friendly GUI interface
- 📊 Read data from Excel files (.xlsx)

## Prerequisites

### Windows
1. **Install GCC Compiler**
   - Download and install TDM-GCC from: https://jmeubank.github.io/tdm-gcc/
   - Or install MinGW-w64

2. **Install Go**
   - Download from: https://golang.org/dl/
   - Minimum version: Go 1.16+

## Installation

1. **Clone or download this project**
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Prepare data.xlsx file** with the format:
   | Kanji | Hiragana | Reading | Meaning |
   |-------|----------|---------|---------|
   | 私 | わたし | watashi | I/me |
   | 本 | ほん | hon | book |

## How to Run

### Run directly
```bash
go run main.go
```

### Compile to executable
```bash
# For Windows
go build -o flashcard.exe main.go

# For other systems
go build -o flashcard main.go
```

After compilation, you can simply double-click the `flashcard.exe` file to run it.

### Cross-compile for other OS
```bash
# For Linux from Windows
set GOOS=linux
set GOARCH=amd64
go build -o flashcard-linux main.go

# For macOS from Windows  
set GOOS=darwin
set GOARCH=amd64
go build -o flashcard-mac main.go
```

## Usage

1. Make sure `data.xlsx` file is in the same folder as the executable
2. Run the application
3. Use `<` and `>` buttons for navigation
4. Click "Hide" button to hide reading or meaning
5. Study vocabulary by hiding answers first

## Excel File Format

The `data.xlsx` file must have the following structure:
- **Sheet1** with columns:
  - Column A: Kanji
  - Column B: Hiragana  
  - Column C: Reading (romaji)
  - Column D: Meaning (Indonesian/English)
- First row as header (will be ignored)

## Troubleshooting

### Error "gcc not found"
- Install TDM-GCC or MinGW-w64
- Make sure gcc is in PATH

### Error "cannot find data.xlsx"
- Make sure data.xlsx file is in the same folder as the executable
- Check the format and content of the Excel file

### GUI doesn't appear
- Make sure all Fyne dependencies are installed
- Try running: `go mod tidy`

## Dependencies

- `github.com/xuri/excelize/v2` - for reading Excel files
- `fyne.io/fyne/v2` - for GUI framework

## Contributing

Please create issues or pull requests for improvements and new features.

## License

This project is free to use for educational purposes.
