package usecase_repository

type CryptoRepository interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
	EncryptString(data string) (string, error)
	DecryptString(data string) (string, error)
}
