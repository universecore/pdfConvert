package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
)

func main() {
	// Read text file
	filename := "test.txt"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		pdf.SetError(err)
	}

	pdf.SetFont("Times", "", 14)
	// Output justified text
	pdf.MultiCell(190, 5, string(content), "0", "0", false)
	// Line break
	pdf.Ln(-1)

	// pdf.Cell(0, 5, "(end of excerpt)")

	_ = pdf.OutputFileAndClose("test.pdf")

	fmt.Println("PDF file created")
}
