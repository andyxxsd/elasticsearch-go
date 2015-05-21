package query

type Query interface {
	MarshalJSON() ([]byte, error)
}
