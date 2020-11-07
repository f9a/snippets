package snippets

import (
	"crypto/rand"
	"math/big"
)

func RandInt(min, max int64) (num int64, err error) {
	x := big.NewInt(max - min + 1)

	r, err := rand.Int(rand.Reader, x)
	if err != nil {
		return 0, err
	}

	r.Add(r, big.NewInt(min))

	return r.Int64(), nil
}

func RandIntn(max int64) (num int64, err error) {
	return RandInt(0, max)
}
