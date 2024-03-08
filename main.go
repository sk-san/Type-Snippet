package main

import (
	"database/sql"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
	"image/color"
	"log"
	"math/rand"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Type snippet")

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	text := canvas.NewText("", blue)
	target_text := canvas.NewText("ok", color.White)
	Datadetail := canvas.NewText("", color.White)

	target_text.TextSize = 20
	Datadetail.TextSize = 20
	text.TextSize = 20
	target_text.Move(fyne.NewPos(50, 50))
	Datadetail.Move(fyne.NewPos(50, 50))

	status_label := widget.NewLabel("")

	textArea := widget.NewEntry()

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	form := &widget.Form{}

	min_DB := 1
	max_DB := Get_tail_ID()

	go func() {
		for {
			text.Text = "" + textArea.Text
			text.Color = green
			fmt.Println(text.Text)
			text.Refresh()

			if target_text.Text == text.Text {
				rand.Seed(time.Now().UnixNano())
				detail, snippet := Get_data_from_DB(rand.Intn(max_DB-min_DB+1) + min_DB)

				Datadetail.Text = detail
				target_text.Text = snippet

				target_text.Refresh()
				Datadetail.Refresh()
			}

			time.Sleep(50 * time.Millisecond)
		}
	}()
	form.Append("", textArea)
	status_label.SetText("All: cmd+a, Delete: any + All")
	myWindow.SetContent(container.NewVBox(Datadetail, target_text, text, status_label, form))
	myWindow.Resize(fyne.NewSize(300, 150))
	myWindow.ShowAndRun()

}

func Get_data_from_DB(id int) (string, string) {

	db, err := sql.Open("sqlite3", "dataset.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM snipets WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var detail string
	var snippet string

	for rows.Next() {
		var id int
		err := rows.Scan(&id, &detail, &snippet)
		if err != nil {
			log.Println(err)
		}
	}
	return detail, snippet
}
func Get_tail_ID() int {
	db, err := sql.Open("sqlite3", "dataset.db")
	if err != nil {
		log.Println(err)
	}
	query := "SELECT * FROM snipets ORDER BY id DESC"
	row := db.QueryRow(query)
	var id int
	var detail string
	var snippet string
	err = row.Scan(&id, &detail, &snippet)
	if err != nil {
		log.Println(err)
	}
	return id
}
