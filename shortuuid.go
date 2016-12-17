package shortuuid

import (
	"fmt"
	"math/big"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type ShortUUID string

func (s ShortUUID) String() string {
	return string(s)
}

func NewV1() ShortUUID {
	return encode(uuid.NewV1().String(), base62Alphabet)
}

func NewV2(domain byte) ShortUUID {
	return encode(uuid.NewV2(domain).String(), base62Alphabet)
}

func NewV3(ns uuid.UUID, name string) ShortUUID {
	return encode(uuid.NewV3(ns, name).String(), base62Alphabet)
}

func NewV4() ShortUUID {
	return encode(uuid.NewV4().String(), base62Alphabet)
}

func NewV5(ns uuid.UUID, name string) ShortUUID {
	return encode(uuid.NewV5(ns, name).String(), base62Alphabet)
}

func encode(s, alphabet string) ShortUUID {
	runes := []rune(alphabet)
	i := big.NewInt(0)
	alphabetLength := big.NewInt(62)
	zero := new(big.Int)
	s = "0x" + strings.ToLower(strings.Replace(s, "-", "", -1))

	fmt.Sscan(s, i)
	output := []string{}
	for i.Cmp(zero) == 1 {
		prevNumber := new(big.Int)
		prevNumber.Set(i)
		i.Div(i, alphabetLength)
		digit := new(big.Int)
		digit.Mod(prevNumber, alphabetLength)
		output = append(output, string(runes[digit.Int64()]))
	}

	return ShortUUID(strings.Join(output, ""))
}

func (s ShortUUID) UUID() uuid.UUID {
	runes := []rune(string(s))
	N := new(big.Int)
	alphabetLength := big.NewInt(int64(len(base62Alphabet)))
	for i := range runes {
		currentChar := runes[len(s)-(i+1)]
		N.Mul(N, alphabetLength)
		N.Add(N, big.NewInt(int64(strings.Index(base62Alphabet, string(currentChar)))))
	}

	bytes := N.Bytes()
	extraBytes := []byte{}
	for i := 0; i < 16-len(bytes); i++ {
		extraBytes = append(extraBytes, byte(0))
	}
	res, _ := uuid.FromBytes(append(extraBytes, bytes...))

	return res
}

func FromUUID(u uuid.UUID) ShortUUID {
	return encode(u.String(), base62Alphabet)
}
