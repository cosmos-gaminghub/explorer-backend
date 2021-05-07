package main

import (
	"encoding/base64"
	"encoding/hex"
	"log"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
)

func main() {
	ed25519Test()
	secp256k1Test()
	getCCNAddrs()
}

func ed25519Test() {
	log.Println("---Generate consensus address using ed25519---")

	edPrivKey := ed25519.GenPrivKey()
	log.Printf("private key: %s", edPrivKey.String())

	edPubKey := edPrivKey.PubKey()
	log.Printf("public key: %s", edPubKey.String())

	edAddress := edPubKey.Address()
	log.Printf("address: %s", edAddress)

	consAddress, err := types.ConsAddressFromHex(edAddress.String())
	if err != nil {
		panic(err)
	}
	log.Printf("consensus address: %s", consAddress)

	consPub, err := types.Bech32ifyPubKey(types.Bech32PubKeyTypeConsPub, edPubKey)
	if err != nil {
		panic(err)
	}
	log.Printf("consensus pub address: %s", consPub)
}

func secp256k1Test() {
	log.Println("---Generate account address and validator address using secp256k1---")

	secpPrivKey := secp256k1.GenPrivKey()
	log.Printf("private key: %s", secpPrivKey.String())

	secpPubKey := secpPrivKey.PubKey()
	log.Printf("public key: %s", secpPubKey.String())

	secpAddress := secpPubKey.Address()
	log.Printf("address: %s", secpAddress)

	accAddress, err := types.AccAddressFromHex(secpAddress.String())
	if err != nil {
		panic(err)
	}
	log.Printf("account address: %s", accAddress)

	valAddress, err := types.ValAddressFromHex(secpAddress.String())
	if err != nil {
		panic(err)
	}
	log.Printf("validator operator address: %s", valAddress)
}

func getCCNAddrs() {
	log.Println("---Get address from each Bech32 addresses---")
	accAddrBech32Str := "cosmos1qaa9zej9a0ge3ugpx3pxyx602lxh3ztqda85ee"
	valAddrBech32Str := "cosmosvaloper1qaa9zej9a0ge3ugpx3pxyx602lxh3ztqgfnp42"
	consPubAddrBech32Str := "cosmosvalconspub1zcjduepqgx5wxl6eygqf6rx4gura2dy5vzelthjgqntk7q9l2hnpjqam6atsq8u0lx"

	accAddress, err := types.AccAddressFromBech32(accAddrBech32Str)
	if err != nil {
		panic(err)
	}
	log.Printf("address: %s", hex.EncodeToString(accAddress))

	valAddress, err := types.ValAddressFromBech32(valAddrBech32Str)
	if err != nil {
		panic(err)
	}
	log.Printf("address: %s", hex.EncodeToString(valAddress))

	consPubKey, err := types.GetPubKeyFromBech32(types.Bech32PubKeyTypeConsPub, consPubAddrBech32Str)
	if err != nil {
		panic(err)
	}
	log.Printf("consensus public key: %s", consPubKey)
	log.Printf("base64 encoded consensus public key: %s", base64.StdEncoding.EncodeToString(consPubKey.Bytes()))

	address := consPubKey.Address()
	log.Printf("address: %s", address.String())

	consAddress := types.GetConsAddress(consPubKey)
	log.Printf("consensus address: %s", consAddress)
}
