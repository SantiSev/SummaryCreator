package pdfManager_test

import (
	"gpt_summary_creator/fileManagement/pdfManager"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPdfInRange(t *testing.T) {
	pdfManager := pdfManager.NewPdfReader("../../TheWitcher.pdf", 2, 2)

	secondPageOfBook := pdfManager.Read()

	// FOR MORE ACCURACY USE THE PRINT LINE
	// fmt.Println(secondPageOfBook)

	require.True(t, (len(secondPageOfBook) > 0 &&
		strings.HasPrefix(secondPageOfBook, "This book is a work of fiction.")) &&
		strings.HasSuffix(secondPageOfBook, "First eBook Edition: May 2008\n\nISBN: 978-0-316-05508-6\n\n"))

}
