package shortuuid

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

type encodeDecodeTest struct {
	uuid      string
	shortUUID string
}

var encodeDecodeTestCases = []encodeDecodeTest{
	encodeDecodeTest{
		uuid:      "00009272-f1a9-4c18-a964-b78ac3e826ae",
		shortUUID: "09WquDd4uiDt9eYIxCG",
	},
	encodeDecodeTest{
		uuid:      "0002377a-e9e3-477f-ba9a-2578a47d6160",
		shortUUID: "yEuM5wguKI3G9Jed8o01",
	},
}

func (testCase encodeDecodeTest) TestEncode(t *testing.T) {
	uuid, err := uuid.FromString(testCase.uuid)
	if err != nil {
		t.Errorf("Bad encode test string %s.", testCase.uuid)
	}
	s := FromUUID(uuid)
	if s.String() != testCase.shortUUID {
		t.Errorf("Encoding error for %s. Got %s expected %s", testCase.uuid, s.String(), testCase.shortUUID)
	}
}

func (testCase encodeDecodeTest) TestDecode(t *testing.T) {
	s := ShortUUID(testCase.shortUUID)
	u := s.UUID().String()
	if u != testCase.uuid {
		t.Errorf("Decoding error for %s. Got %s expected %s", testCase.shortUUID, u, testCase.uuid)
	}
}

func TestUUID(t *testing.T) {
	for _, testCase := range encodeDecodeTestCases {
		testCase.TestDecode(t)
	}
}

func TestFromUUID(t *testing.T) {
	for _, testCase := range encodeDecodeTestCases {
		testCase.TestEncode(t)
	}
}

func TestNewV1(t *testing.T) {
	s := NewV1()
	l := len(s.String())
	if l < 21 || 22 < l {
		t.Errorf("Shortened V1 incorrect length.")
	}
}

func TestNewV4(t *testing.T) {
	s := NewV4()
	l := len(s.String())
	if l < 21 || 22 < l {
		t.Errorf("Shortened V4 incorrect length.")
	}
}

func TestNewV5(t *testing.T) {
	s := NewV5(uuid.NewV4(), "test")
	l := len(s.String())
	if l < 21 || 22 < l {
		t.Errorf("Shortened V5 incorrect length.")
	}
}
