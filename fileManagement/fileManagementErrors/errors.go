package fileManagementErrors

// ErrorFileExtensionNotSupported is an error that occurs when a file extension is not supported
type ErrorFileExtensionNotSupported struct{}

func (e ErrorFileExtensionNotSupported) Error() string {
	return "File extension not supported"
}
