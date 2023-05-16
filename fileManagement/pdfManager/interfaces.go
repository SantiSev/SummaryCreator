package pdfManager

type PdfReaderInterface interface {
	// Read reads the content of a pdf file and returns it as a string
	Read() string
}
