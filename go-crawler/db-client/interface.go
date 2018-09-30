package dbclient

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

// Controller is an interface for all DB Controllers.
type Controller interface {
	WriteBlock(block *models.Block) error
	WriteTransaction(tx *models.Transaction) error
	Connect() error
	Close() error
}
