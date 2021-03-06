// +build unit

package nodeclient

import (
	"github.com/ccdle12/Blocksage/crawler/injector"
	"github.com/ccdle12/Blocksage/crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type NodeClientSuite struct {
	suite.Suite
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteUnitNodeClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(NodeClientSuite))
}

// ===========================================================
// Unit Tests
// ===========================================================
// TestInitializingNodeClient will test that the Node Client class will be
// initialized correctly.
func (suite *NodeClientSuite) TestInitNodeClient() {
	nodeClient := New(injector.DefaultHTTPClient(), testutils.NodeAddress, testutils.Username, testutils.Password)

	// nodeClient should be initialized
	suite.NotNil(nodeClient, "Should be able to initialize Client.")
}

// TestGetBlock will call GetBlock() and return a Block Struct. The test will be using a mock test server.
func (suite *NodeClientSuite) TestMockSendNodeRequest() {
	// Create Test Server and pass it to New()
	server := testutils.TestServer(testutils.NodeResCorrectBlockNoTx0)
	defer server.Close()

	nodeClient := New(server.Client(), server.URL, testutils.Username, testutils.Password)
	suite.EqualValues(server.URL, nodeClient.address)

	block, err := nodeClient.GetBlock("0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4")

	suite.NoError(err, "There should be no error when getting NodeResCorrectBlockNoTx0")
	suite.NotNil(block, "Returned block should not be nil")
	suite.EqualValues("000000000000000000000000000000000000000002eb51495ec06b0a5427f048", block.Chainwork, "Chain Work should equal the chainwork found in block-respones.go")
	suite.EqualValues(6727225469722.534, block.Difficulty, "Difficulty should be a float64 that matches the model in block_responses.go")
}

// TestSendRequestToMalformedServer will attempt to send a request to a server that is offline or not online at all.
func (suite *NodeClientSuite) TestSendRequestToMalformedServer() {
	// Create Test Server and pass it to New()
	server := testutils.TestServer(testutils.NodeResCorrectBlockNoTx0)
	defer server.Close()

	nodeClient := New(server.Client(), "http://localhost:3421", testutils.Username, testutils.Password)
	_, err := nodeClient.GetBlock("0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4")

	suite.Error(err, "There should be an error")
}
