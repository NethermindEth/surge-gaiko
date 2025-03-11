package witness

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

var emptyEthDepositHash = common.HexToHash(
	"569e75fc77c1a856f6daaf9e69d8a9566ca34aa47f9133711ce065a571af0cfd",
)

const (
	anchorGasLimit         = 250_000
	blockMaxTxListBytes    = params.BlobTxBytesPerFieldElement * params.BlobTxFieldElementsPerBlob
	calldataMaxTxListBytes = (params.BlobTxBytesPerFieldElement - 1) * params.BlobTxFieldElementsPerBlob
	blockMaxGasLimit       = 240_000_000
)
