package signer

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthereumWallet struct{}

func (ew *EthereumWallet) SignTransaction(privateKey string, unsignedTx *UnsignedTx) (*SignedTx, error) {
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	privKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	decodedUnsignedTxBytes, err := hex.DecodeString(string(unsignedTx.Data))
	if err != nil {
		return nil, err
	}

	tx := &types.Transaction{}

	err = tx.UnmarshalBinary(decodedUnsignedTxBytes)
	if err != nil {
		return nil, err
	}

	signer := types.LatestSignerForChainID(tx.ChainId())

	signedTx, err := types.SignTx(tx, signer, privKey)
	if err != nil {
		return nil, err
	}

	signedTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return &SignedTx{
		Data: []byte(hex.EncodeToString(signedTxBytes)),
	}, nil
}
