package signer

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type SolanaWallet struct{}

func (sw *SolanaWallet) SignTransaction(privateKeys []string, unsignedTx *UnsignedTx) (*SignedTx, error) {
	signers := make([]solana.PrivateKey, 0, len(privateKeys))

	for _, privateKey := range privateKeys {
		signer, err := solana.PrivateKeyFromBase58(privateKey)
		if err != nil {
			return nil, fmt.Errorf("error getting private key: %w", err)
		}

		signers = append(signers, signer)
	}

	data := base58.Decode(string(unsignedTx.Data))

	// parse transaction:
	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(data))
	if err != nil {
		panic(err)
	}

	// clear signatures: https://github.com/gagliardetto/solana-go/issues/186
	tx.Signatures = nil

	if _, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		for _, candidate := range signers {
			if candidate.PublicKey().Equals(key) {
				return &candidate
			}
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("error signing transaction: %w", err)
	}

	marshaledTx, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("error marshaling transaction: %w", err)
	}

	base58EncodedTx := base58.Encode(marshaledTx)

	return &SignedTx{
		Data: []byte(base58EncodedTx),
	}, nil
}
