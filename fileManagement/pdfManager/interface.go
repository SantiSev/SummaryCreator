package pdfManager

type PdfManagerInterface interface {
	// Read reads the content of a pdf file and returns it as a string
	Read() string
	// CreatePdf creates a pdf file
	CreatePdf(fileName, fileContent string)
}
