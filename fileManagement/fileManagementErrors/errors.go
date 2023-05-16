package fileManagementErrors

type ErrorFileExtensionNotSupported struct{}

func (e ErrorFileExtensionNotSupported) Error() string {
	return "File extension not supported"
}
