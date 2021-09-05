package repository

import "testing"

func TestCrypto(t *testing.T) {
	key := []byte("G-KaPdSgVkYp2s5r8y/B?E(H+MAQeThq")
	tests := [][]byte{
		[]byte(""),
		[]byte("a"),
		[]byte("abcdef"),
		[]byte("wp4894261-black-"),
		[]byte("wp4894261-black-m"),
		[]byte("wp4894261-black-minimalism-wallpapers.png"),
	}

	checkEncryptDecrypt := func(t *testing.T, s []byte, key []byte) {
		t.Helper()
		crypto, err := NewAESCrypto(key)
		if err != nil {
			t.Error(err)
			return
		}

		encrypted, err := crypto.Encrypt(s)
		if err != nil {
			t.Error(err)
			return
		}

		decrypted, err := crypto.Decrypt(encrypted)
		if err != nil {
			t.Error(err)
			return
		}

		if len(s) != len(decrypted) {
			t.Errorf("len(encrypted) != len(decrypted)")
			return
		}

		for i := range decrypted {
			if s[i] != decrypted[i] {
				t.Errorf("s[%v] != decrypted[%v]", i, i)
				return
			}
		}
	}

	for _, test := range tests {
		checkEncryptDecrypt(t, test, key)
	}
}
