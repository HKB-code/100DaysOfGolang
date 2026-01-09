package iomanager

type IOManager interface {
	ReadFiles() ([]string, error)
	WriteResult(data any) error
}