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


/*
	package wallet

import (
    "crypto/ecdsa"
    "fmt"
    "wallet-service/models"

    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "golang.org/x/crypto/sha3"
)

func GenerateWallet() (*models.WalletResponse, error) {
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        return nil, err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("error casting public key to ECDSA")
    }

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

    // Derive address manually to verify consistency
    hash := sha3.NewKeccak256()
    hash.Write(publicKeyBytes[1:])
    derivedAddress := hexutil.Encode(hash.Sum(nil)[12:])

    // Go-ethereum method (recommended)
    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

    if address != derivedAddress {
        return nil, fmt.Errorf("address mismatch: %s != %s", address, derivedAddress)
    }

    return &models.WalletResponse{
        Address: address,
    }, nil
}

*/


