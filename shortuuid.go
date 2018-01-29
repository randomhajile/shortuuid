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

func inBase62Alphabet(r rune) bool {
	return '0' <= r && r <= '9' || 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z'
}

func FromString(s string) (ShortUUID, error) {
	for _, r := range s {
		if !inBase62Alphabet(r) {
			return ShortUUID(""), fmt.Errorf("Rune '%c' not in base62 alphabet", r)
		}
	}
	return ShortUUID(s), nil
}

func NewV1() (ShortUUID, error) {
	u, err := uuid.NewV1()
	return encode(u.String(), base62Alphabet), err
}

func NewV2(domain byte) (ShortUUID, error) {
	u, err := uuid.NewV2(domain)
	return encode(u.String(), base62Alphabet), err
}

func NewV3(ns uuid.UUID, name string) ShortUUID {
	return encode(uuid.NewV3(ns, name).String(), base62Alphabet)
}

func NewV4() (ShortUUID, error) {
	u, err := uuid.NewV4()
	return encode(u.String(), base62Alphabet), err
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
	extraBytes := make([]byte, 16-len(bytes))
	res, _ := uuid.FromBytes(append(extraBytes, bytes...))

	return res
}

func FromUUID(u uuid.UUID) ShortUUID {
	return encode(u.String(), base62Alphabet)
}
