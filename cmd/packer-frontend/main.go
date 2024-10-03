package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-resty/resty/v2"

	"packer/internal/model"
)

type PackRequestResponse struct {
	Status string `json: "status"`
	Packs  []int  `json: "packs"`
}

func main() {
	a := app.New()
	w := a.NewWindow("Packer App")
	w.Resize(fyne.NewSize(200, 300))

	// elements
	resultField := widget.NewLabel("Packages: ")
	entryField := widget.NewEntry()
	packButton := widget.NewButton("Pack", func() {
		fmt.Println("Packing was requested for items: " + entryField.Text + " with packages: " + os.Getenv("BUCKETS"))

		// prepare request paramters
		var packs model.Packs
		packs.Value, _ = strconv.Atoi(entryField.Text)
		str := strings.Split(os.Getenv("BUCKETS"), ",")
		for i := range str {
			val, _ := strconv.Atoi(str[i])
			packs.Buckets = append(packs.Buckets, val)
		}

		fmt.Printf("Request: %v", packs)

		var response PackRequestResponse

		// Send request to API
		client := resty.New()
		_, err := client.R().
			SetHeader("Accept", "application/json").
			SetBody(&packs).
			SetResult(&response).
			Post("http://localhost:8080/pack")
		if err != nil {
			fmt.Printf("ERROR: %v", err)
		} else {
			fmt.Printf("Response: %v", response)
			resultField.SetText("Packages: " + fmt.Sprintf("%v", response.Packs))
		}
	})

	// main window content using elements
	w.SetContent(
		container.NewBorder(
			container.NewCenter(
				resultField,
			),

			container.NewGridWithColumns(
				2,
				entryField,
				packButton,
			),

			nil, // Right
			nil, // Left

			// the rest will take all the rest of the space
			container.NewCenter(
				widget.NewLabel("Buckets: "+os.Getenv("BUCKETS")),
			),
		),
	)
	w.ShowAndRun()
}
