package signer

// Wallet defines a structure to hold the public and private keys.
type Wallet struct {
	PublicKey  string
	PrivateKey string
}

// UnsignedTx represents an unsigned transaction.
type UnsignedTx struct {
	Data []byte
}

// SignedTx represents a signed transaction.
type SignedTx struct {
	Data []byte
}

// Signer is an interface for signing unsigned transactions.
type Signer interface {
	// SignTransaction signs an unsigned transaction with the given private key and protocol.
	SignTransaction(privateKey string, unsignedTx *UnsignedTx) (*SignedTx, error)
}

func New(protocol string) Signer {
	if protocol == "polygon" {
		return &EthereumWallet{}
	}

	return nil
}
