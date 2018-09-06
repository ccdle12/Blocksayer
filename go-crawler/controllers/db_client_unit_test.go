// +build unit

package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type DBClientSuite struct {
	suite.Suite
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteUnitDBClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(DBClientSuite))
}

// ===========================================================
// Unit Tests
// ===========================================================
// TestInitDBClient will test that a DBClient can be initialized.
func (suite *DBClientSuite) TestInit() {
	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	suite.NoError(err, "There should be no error")
	suite.NotNil(dbClient, "DBClient is not nil when attempting to initialize")
}

// TestConfigInit will test that the cfg DBConfig struct was initialized by the
// constructor.
func (suite *DBClientSuite) TestConfigInit() {
	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	suite.NoError(err, "There should be no error")
	suite.NotNil(dbClient.cfg, "cfg is not nil when attempting to initialize")
	suite.EqualValues(dbClient.cfg.DBHost, testutils.DBHost, "Same host was initialized in cfg.DBHost")
}

// TestEmptyStringInit will test that we are passing an empty string as one of the parameters for
// initializing a DBClient.
func (suite *DBClientSuite) TestEmptyStringInit() {
	dbClient, err := NewDBClient(testutils.DBHost, "", testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	suite.Error(err, "Error should be returned")
	suite.EqualValues(err, utils.ErrPassingEmptyString, "The error return ErrPassingEmptyString")
	suite.Nil(dbClient, "DBClient is not nil when attempting to initialize")
}

// TestUseCaseInit will test that the DB usecase was initialized.
func (suite *DBClientSuite) TestUseCaseInit() {
	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	suite.NoError(err, "No error should be returned")
	suite.NotNil(dbClient.usecase, "usecase is not nil and initialized")
}

// TODO (ccdle12): For Integration Tests
// TestDBClientConnection
// func TestDBClientConnection(t *testing.T) {
// 	suite. := suite.New(t)

// 	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
// 		testutils.DBPassword, testutils.DBName, testutils.DBType)

// 	suite.NoError(err, "No error should be returned")
// 	suite.NotNil(dbClient.usecase, "usecase is not nil and initialized")

// 	dbClient.Connect()
// }

// TODO (ccdle12): For Integration Tests
// TestDBClientClose
// func TestDBClientClose(t *testing.T) {
// 	suite. := suite.New(t)

// 	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
// 		testutils.DBPassword, testutils.DBName, testutils.DBType)

// 	suite.NoError(err, "No error should be returned")
// 	suite.NotNil(dbClient.usecase, "usecase is not nil and initialized")

// 	dbClient.Close()
// }