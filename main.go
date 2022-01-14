package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
)

func main() {

	file := "test.txt"

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("%s file not found", file)
	}
	//This is a file description:
	// P - orientation, mm - unit metrics, A4 - sheet size, "" font type
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "Bold", 14)

	pdf.MultiCell(0, 5, string(content), "", "", false)
	//line break
	pdf.Ln(-1)

	pdf.Cell(0, 5, "(end of excerpt)")

	_ = pdf.OutputFileAndClose("test.pdf")

	fmt.Println("PDF file created")
}
