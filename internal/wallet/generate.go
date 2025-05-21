package wallet

import (
    "crypto/ecdsa"
    "github.com/ethereum/go-ethereum/crypto"
)

func GenerateWalletAddress() (string, *ecdsa.PrivateKey, error) {
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        return "", nil, err
    }

    address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
    return address, privateKey, nil
}
