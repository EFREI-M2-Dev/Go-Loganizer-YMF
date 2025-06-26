package analyzer

type UnreachableFileError struct {
	Path string
	Err error
	Type string
}

func (e *UnreachableFileError) Error() string {
	return "impossible d'accéder au fichier " + e.Path + ": " + e.Err.Error()
}

type ParsingError struct {
	Path string
	Err error
	Type string
}

func (e *ParsingError) Error() string {
	return "échec du parsing du fichier " + e.Path + ": " + e.Err.Error()
}

type UnsupportedFileTypeError struct {
	Path string
	Type string
}

func (e *UnsupportedFileTypeError) Error() string {
	return "type de fichier non supporté: " + e.Type + " pour le fichier " + e.Path
}

