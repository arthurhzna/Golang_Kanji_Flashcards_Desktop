package main

import (
    "fmt"
    "log"

    "github.com/xuri/excelize/v2"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

type Kosakata struct {
    Kanji, Hiragana, Bacaan, Arti string
}

func main() {
    // Load Excel
    f, err := excelize.OpenFile("data.xlsx")
    if err != nil {
        log.Fatal(err)
    }
    rows, _ := f.GetRows("Sheet1")

    var data []Kosakata
    for i, row := range rows {
        if i == 0 {
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

    a := app.New()
    title := "Flashcard"
    w := a.NewWindow(title)

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
    // exitBtn removed

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
        ), 
        container.NewGridWithColumns(2, prevBtn, nextBtn), 
        nil, 
        nil, 
        nil, 
    )

    w.SetContent(content)
    w.Resize(fyne.NewSize(300, 90))
    w.ShowAndRun()
}