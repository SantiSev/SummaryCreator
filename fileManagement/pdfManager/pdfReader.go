package pdfManager

import (
	"bufio"
	"log"
	"os"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

// init Initializes the PdfReader by setting the license key for unipdf
// without this function declaration the pdf reader will not work
func init() {
	err := license.SetMeteredKey(getLicence("../../unidocApikey.txt"))
	if err != nil {
		panic(err)
	}
}

type PdfReader struct {
	reader    *model.PdfReader
	fromPage  int
	untilPage int
	fileName  string
}

func NewPdfReader(fileName string, fromPage, untilPage int) PdfReaderInterface {
	r := new(PdfReader)
	r.fromPage = fromPage
	r.untilPage = untilPage
	r.fileName = fileName
	return r
}

func (r *PdfReader) Read() string {
	// TODO Concurrency (parallelize the reading of the pages)

	// Open a PDF file.
	file, err := os.Open(r.fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new PDF reader.
	r.reader, err = model.NewPdfReader(file)
	if err != nil {
		panic(err)
	}

	r.validatePageRange(r.fromPage, r.untilPage)

	ch := make(chan string, 1)

	// Extract text from each page.
	for pageNum := r.fromPage; pageNum <= r.untilPage; pageNum++ {
		page, err := r.reader.GetPage(pageNum)
		if err != nil {
			panic(err)
		}

		textExtractor, err := extractor.New(page)
		if err != nil {
			panic(err)
		}

		text, err := textExtractor.ExtractText()
		if err != nil {
			panic(err)
		}

		ch <- text
	}
	return <-ch
}

func (r *PdfReader) validatePageRange(fromPage, untilPage int) {

	// Get number of pages.
	numPages, err := r.reader.GetNumPages()
	if err != nil {
		panic(err)
	}

	if (fromPage < 1 || untilPage < 1) || (fromPage > untilPage || untilPage > numPages) {
		panic("Invalid page range")
	}
}

// Util

// getLicence returns the licence key from the given file
func getLicence(fromFile string) string {
	file, err := os.Open(fromFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	apiKey := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return apiKey
}
