package cryptography_test

import (
	"fmt"
	"encoding/base64"
	"testing"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/require"
)

func TestSignAndVerify_BCH(t *testing.T) {
	pkBytes := []byte("some random string of chars... this is probably a weak private key seed")

	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	rawAddress, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.MainNetParams)

	address := rawAddress.EncodeAddress()

	fmt.Printf("PrivKey: %s\n", privKey.Serialize())
	fmt.Printf("P2PKH address: %s\n", address)

	message := "some_message_to_sign"

	signedBytes, err := btcec.SignCompact(btcec.S256(), privKey, []byte(message), true)
	require.NoError(t, err)

	signedMessage := base64.StdEncoding.EncodeToString(signedBytes)

	fmt.Printf("signedMessage: %s\n", signedMessage)


	decodedSignedMessage, err := base64.StdEncoding.DecodeString(signedMessage)
	require.NoError(t, err)

	recoveredPubKey, _, err := btcec.RecoverCompact(btcec.S256(), decodedSignedMessage, []byte(message))
	require.NoError(t, err)

	rawRecoveredAddress, err := btcutil.NewAddressPubKey(recoveredPubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	require.NoError(t, err)

	recoveredAddress := rawRecoveredAddress.EncodeAddress()

	fmt.Println("Recovered address", recoveredAddress)

	require.True(t, address==recoveredAddress, "Message verification failed")
}

