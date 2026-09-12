package iomanager



type IoManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error 
}