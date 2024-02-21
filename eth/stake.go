// description:
// @author renshiwei
// Date: 2024/2/21

package eth

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
)

const (
	BatchDepositMethod = "deposit"
)

type TxParams struct {
	From             string `json:"from"`
	To               string `json:"to"`
	AmountWei        string `json:"amount_wei"`
	ContractCallData string `json:"contract_call_data"`
}

func SignTx(ctx context.Context, evmAddr string, txParams TxParams, privateKeyHex string) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", errors.Wrap(err, "crypto.HexToECDSA for privateKeyHex.")
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	to := common.HexToAddress(txParams.To)
	data, err := hexutil.Decode(txParams.ContractCallData)
	if err != nil {
		return "", errors.Wrap(err, "ContractCallData decode.")
	}
	value, ok := new(big.Int).SetString(txParams.AmountWei, 0)
	if !ok {
		return "", errors.Errorf("AmountWei:%s params err.", txParams.AmountWei)
	}

	// Get the nonce for your account
	client, err := ethclient.Dial(evmAddr)
	if err != nil {
		return "", errors.Wrap(err, "eth client dial err.")
	}
	nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(address))
	if err != nil {
		return "", errors.Wrap(err, "get nonce.")
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get chainID.")
	}
	s := types.NewEIP155Signer(chainID)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get gasPrice.")
	}

	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  common.HexToAddress(txParams.From),
		To:    &to,
		Data:  data,
		Value: value,
	})
	if err != nil {
		return "", errors.Wrap(err, "EstimateGas.")
	}
	gasLimit = gasLimit + gasLimit*20/100

	notSignedTx := &types.LegacyTx{
		To:    &to,
		Value: value,
		Data:  data,

		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
	}

	signedTx, err := types.SignNewTx(privateKey, s, notSignedTx)
	if err != nil {
		return "", errors.Wrap(err, "SignNewTx.")
	}

	txBinary, err := signedTx.MarshalBinary()
	if err != nil {
		return "", errors.Wrap(err, "signedTx MarshalBinary.")
	}

	return hexutil.Encode(txBinary), nil
}

// BroadcastSignedTx Broadcast signature transactions
// @param evmAddr eth endpoints
// @param signedTransaction @see `SignTx`
func BroadcastSignedTx(ctx context.Context, evmAddr string, signedTransaction string, isWaitResult bool) (common.Hash, error) {
	signedTxByte, err := hexutil.Decode(signedTransaction)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "singedTx decode err.")
	}

	singedTx := &types.Transaction{}
	err = singedTx.UnmarshalBinary(signedTxByte)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "singedTx UnmarshalJSON err.")
	}
	txHash := singedTx.Hash()

	client, err := ethclient.Dial(evmAddr)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "eth client dial err.")
	}

	err = client.SendTransaction(ctx, singedTx)
	if err != nil {
		return txHash, errors.Wrap(err, "SendTransaction err.")
	}

	if isWaitResult {
		receipt, err := bind.WaitMined(ctx, client, singedTx)
		if err != nil {
			return txHash, errors.Wrapf(err, "tx wait mined err. txHash:%s", txHash.String())
		}
		if receipt.Status == types.ReceiptStatusFailed {
			return txHash, errors.Wrapf(err, "tx status is ReceiptStatusFailed. txHash:%s", txHash.String())
		}
	}

	return txHash, nil
}
