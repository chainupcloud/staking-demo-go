// description:
// @author renshiwei
// Date: 2024/2/21

package eth

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func CreateEth1Key() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", errors.Wrap(err, "publicKey is not *ecdsa.PublicKey")
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	priv := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("publicKey is not *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return priv, address, nil
}
