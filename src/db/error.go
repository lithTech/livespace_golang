package db

type ConcurrentModificationError struct {}

func (e *ConcurrentModificationError) Error() string {
	return "ConcurrentModificationError"
}