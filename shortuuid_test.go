package shortuuid_test

import (
	"testing"

	uuid "github.com/gofrs/uuid"
	"github.com/randomhajile/shortuuid/v2"
	"github.com/stretchr/testify/suite"
)

type shortUUIDTestSuite struct {
	suite.Suite

	uuid      uuid.UUID
	shortUUID shortuuid.ShortUUID
}

func TestShortUUIDTestSuite(t *testing.T) {
	suite.Run(t, new(shortUUIDTestSuite))
}

func (suite *shortUUIDTestSuite) SetupTest() {
	suite.uuid, _ = uuid.FromString("00009272-f1a9-4c18-a964-b78ac3e826ae")
	suite.shortUUID, _ = shortuuid.FromString("09WquDd4uiDt9eYIxCG")
}

func (suite *shortUUIDTestSuite) TestEncode() {
	uuid, err := uuid.FromString(suite.uuid.String())
	suite.NoError(err)

	shortUUID := shortuuid.FromUUID(uuid)
	suite.Equal(suite.shortUUID, shortUUID)
}

func (suite *shortUUIDTestSuite) TestEncodeError() {
	_, err := shortuuid.FromString("!@#$%^&*()")
	suite.Error(err)
}

func (suite *shortUUIDTestSuite) TestDecode() {
	suite.Equal(suite.uuid, suite.shortUUID.UUID())
}

func (suite *shortUUIDTestSuite) TestNewV1() {
	s, err := shortuuid.NewV1()
	suite.NoError(err)

	l := len(s.String())
	suite.True(l == 21 || l == 22)
}

func (suite *shortUUIDTestSuite) TestNewV3() {
	s := shortuuid.NewV3(suite.uuid, "name")

	l := len(s.String())
	suite.True(l == 21 || l == 22)
}

func (suite *shortUUIDTestSuite) TestNewV4() {
	s, err := shortuuid.NewV4()
	suite.NoError(err)

	l := len(s.String())
	suite.True(l == 21 || l == 22)
}

func (suite *shortUUIDTestSuite) TestNewV5() {
	s := shortuuid.NewV5(suite.uuid, "name")

	l := len(s.String())
	suite.True(l == 21 || l == 22)
}
