package helpers

import (
	"crypto/rand"
	"math/big"
)

const numchars = "0123456789"

func GenerateOTPcode(length int) (string, error) {
	var str string = ""

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(numchars))))

		if err != nil {
			return "", err
		}

		str += string(numchars[n.Int64()])

	}

	return str, nil
}
