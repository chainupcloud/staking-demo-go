// description:
// @author renshiwei
// Date: 2024/2/21

package examples

import (
	"context"
	"github.com/chainupcloud/staking-demo-go/eth"
	"github.com/stretchr/testify/require"
	"testing"
)

// eth el endpoint
// @see https://docs.chainupcloud.com/blockchain-api/ethereum
// @see https://www.infura.io/
const evmAddr = "https://api.chainup.net/ethereum/holesky/${YOUR_TOKEN}"

// Sign the tx of the deposit
func TestSignTx(t *testing.T) {
	// /api/v1/validator/create response @see
	txParams := eth.TxParams{
		From:             "",
		To:               "",
		AmountWei:        "",
		ContractCallData: "",
	}

	// your private key
	privateKeyHex := "<YOUR_PRIVATE_KEY>"

	signedTxStr, err := eth.SignTx(context.Background(), evmAddr, txParams, privateKeyHex)
	require.NoError(t, err)

	// signedTxStr as `/api/v1/validator/broadcast` params
	t.Log("signedTxStr:", signedTxStr)
}

// The TX of the deposit is signed and broadcast to the chain
func TestSignAndBroadcastTx(t *testing.T) {
	// /api/v1/validator/create response @see
	txParams := eth.TxParams{
		From:             "",
		To:               "",
		AmountWei:        "",
		ContractCallData: "",
	}

	// your private key
	privateKeyHex := "<YOUR_PRIVATE_KEY>"

	signedTxStr, err := eth.SignTx(context.Background(), evmAddr, txParams, privateKeyHex)
	require.NoError(t, err)

	// signedTxStr as `/api/v1/validator/broadcast` params
	t.Log("signedTxStr:", signedTxStr)

	hash, err := eth.BroadcastSignedTx(context.Background(), evmAddr, signedTxStr, true)
	require.NoError(t, err)

	t.Log("txHash", hash.String())
}
