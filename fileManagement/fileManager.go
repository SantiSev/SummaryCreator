package fileManagement

import (
	docx "gpt_summary_creator/fileManagement/docxManager"
	err "gpt_summary_creator/fileManagement/fileManagementErrors"
	pdf "gpt_summary_creator/fileManagement/pdfManager"
)

// CreateFile is a function that creates a file based on the file extension
// Parameters: fileName string, fileExtension string, fileContent string
// Returns: error
func CreateFile(fileName string, fileExtension string, fileContent string) error {
	switch fileExtension {
	case "pdf":
		pdf.CreatePdf(fileName, fileContent)
		return nil
	case "docx":
		docx.CreateDocx(fileName, fileContent)
		return nil
	default:
		return err.ErrorFileExtensionNotSupported{}
	}
}
