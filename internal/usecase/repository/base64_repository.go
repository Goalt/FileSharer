package usecase_repository

type Base64Repository interface {
	Encode(string) string
	Decode(string) (string, error)
}
