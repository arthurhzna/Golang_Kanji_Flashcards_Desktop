package main

import (
    "fmt"

    "github.com/xuri/excelize/v2"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
)

type Kosakata struct {
    Kanji, Hiragana, Bacaan, Arti string
}

func loadExcelData(filePath string) ([]Kosakata, error) {
    f, err := excelize.OpenFile(filePath)
    if err != nil {
        return nil, err
    }
    rows, _ := f.GetRows("Sheet1")

    var data []Kosakata
    for i, row := range rows {
        if i == 0 { // skip header
            continue
        }
        if len(row) < 4 {
            continue
        }
        data = append(data, Kosakata{
            Kanji:    row[0],
            Hiragana: row[1],
            Bacaan:   row[2],
            Arti:     row[3],
        })
    }
    return data, nil
}

func createFlashcardUI(w fyne.Window, data []Kosakata) {
    index := 0
    showBacaan := true
    showArti := true

    kanjiLabel := widget.NewLabelWithStyle(data[index].Kanji, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
    hiraLabel := widget.NewLabel(fmt.Sprintf("%s (%s)", data[index].Hiragana, data[index].Bacaan))
    artiLabel := widget.NewLabel(data[index].Arti)

    updateLabels := func() {
        kanjiLabel.SetText(data[index].Kanji)
        if showBacaan {
            hiraLabel.SetText(fmt.Sprintf("%s (%s)", data[index].Hiragana, data[index].Bacaan))
        } else {
            hiraLabel.SetText("•••")
        }
        if showArti {
            artiLabel.SetText(data[index].Arti)
        } else {
            artiLabel.SetText("•••")
        }
    }

    nextBtn := widget.NewButton(">", func() {
        if index < len(data)-1 {
            index++
            updateLabels()
        }
    })
    prevBtn := widget.NewButton("<", func() {
        if index > 0 {
            index--
            updateLabels()
        }
    })

    toggleBacaanBtn := widget.NewButton("Hide", func() {
        showBacaan = !showBacaan
        updateLabels()
    })
    toggleArtiBtn := widget.NewButton("Hide", func() {
        showArti = !showArti
        updateLabels()
    })

    toggleBacaanBtn.Resize(fyne.NewSize(15, 15))
    toggleArtiBtn.Resize(fyne.NewSize(15, 15))

    content := container.NewBorder(
        container.NewVBox(
            kanjiLabel,
            container.NewHBox(hiraLabel, toggleBacaanBtn),
            container.NewHBox(artiLabel, toggleArtiBtn),
        ), // top
        container.NewGridWithColumns(2, prevBtn, nextBtn), // bottom
        nil, // left
        nil, // right
        nil, // center
    )

    w.SetContent(content)
    w.Resize(fyne.NewSize(300, 90))
}

func main() {
    a := app.New()
    title := "Flashcard - Select Excel file"
    w := a.NewWindow(title)

    // Create a welcome message
    welcomeLabel := widget.NewLabel("Select Excel file to load flashcard")
    welcomeLabel.Alignment = fyne.TextAlignCenter

    selectButton := widget.NewButton("Select Excel file", func() {
        // Create file dialog
        fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
            if err != nil {
                dialog.ShowError(err, w)
                return
            }
            if reader == nil {
                return // User cancelled
            }
            defer reader.Close()

            filePath := reader.URI().Path()
            
            // Load data from selected file
            data, loadErr := loadExcelData(filePath)
            if loadErr != nil {
                dialog.ShowError(fmt.Errorf("Error loading Excel file: %v", loadErr), w)
                return
            }

            if len(data) == 0 {
                dialog.ShowError(fmt.Errorf("Data format does not match"), w)
                return
            }

            // Change window title and create flashcard UI
            w.SetTitle("Flashcard")
            createFlashcardUI(w, data)

        }, w)

        // Set file filter for Excel files
        fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx", ".xls"}))
        
        // Set dialog size to be larger
        fileDialog.Resize(fyne.NewSize(800, 600))
        fileDialog.Show()
    })

    // Create initial content
    content := container.NewVBox(
        welcomeLabel,
        widget.NewSeparator(),
        selectButton,
    )

    w.SetContent(content)
    // Increase main window size
    w.Resize(fyne.NewSize(500, 250))
    w.ShowAndRun()
}