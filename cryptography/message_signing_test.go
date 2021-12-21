package cryptography_test

import (
	"crypto/ecdsa"
	//"crypto/elliptic"
	"fmt"
	"encoding/base64"
	"testing"
	//"strings"
	"crypto/rand"

	"github.com/ebfe/keccak"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/require"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func TestSignAndVerify_BCH_p2pkh_address(t *testing.T) {
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

func TestSignAndVerify_BCH_segwit_address(t *testing.T) {
	pkBytes := []byte("some random string of chars... this is probably a weak private key seed")

	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	witnessProg := btcutil.Hash160(pubKey.SerializeCompressed())
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	require.NoError(t, err)
	address := addressWitnessPubKeyHash.EncodeAddress()

	fmt.Printf("PrivKey: %s\n", privKey.Serialize())
	fmt.Printf("Segwit address: %s\n", address)

	message := "some_message_to_sign"

	signedBytes, err := btcec.SignCompact(btcec.S256(), privKey, []byte(message), true)
	require.NoError(t, err)

	signedMessage := base64.StdEncoding.EncodeToString(signedBytes)

	fmt.Printf("signedMessage: %s\n", signedMessage)


	decodedSignedMessage, err := base64.StdEncoding.DecodeString(signedMessage)
	require.NoError(t, err)

	recoveredPubKey, _, err := btcec.RecoverCompact(btcec.S256(), decodedSignedMessage, []byte(message))
	require.NoError(t, err)

	recoveredWitnessProg := btcutil.Hash160(recoveredPubKey.SerializeCompressed())
	recoveredAddressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(recoveredWitnessProg, &chaincfg.MainNetParams)
	require.NoError(t, err)
	recoveredAddress := recoveredAddressWitnessPubKeyHash.EncodeAddress()

	fmt.Println("Recovered address", recoveredAddress)

	require.True(t, address==recoveredAddress, "Message verification failed")
}

const (
	magicPrefix = "\x19Ethereum Signed Message:\n"
)

func magicHash(message string) ([]byte, error) {
	msg := fmt.Sprintf("%s%d%s", magicPrefix, len(message), message)
	h := keccak.New256()
	h.Write([]byte(msg))
	return h.Sum(nil), nil
}

func TestSignAndVerify_ETH_using_go_ethereum(t *testing.T) {
	privKey, err := ecdsa.GenerateKey(
		secp256k1.S256(),
		//strings.NewReader("some random other dfdsstring of chars... this is probably a weak private key seed other data"),
		rand.Reader,
	)
	require.NoError(t, err)
	fmt.Printf("PrivKey: %x\n", crypto.FromECDSA(privKey))

	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	require.True(t, ok)
	fmt.Printf("PubKey: %x\n", crypto.FromECDSAPub(pubKeyECDSA))

	address := crypto.PubkeyToAddress(*pubKeyECDSA)
	fmt.Printf("Address: %s\n", address.Hex())

	//messageToSign := "\x19Ethereum Signed Message:\nsome_message_to_sign"
	messageToSign, err := magicHash("some_other_message_to_sign")
	require.NoError(t, err)

	signature, err := crypto.Sign(
		//crypto.Keccak256([]byte(messageToSign)),
		messageToSign,
		privKey,
	)
	require.NoError(t, err)
	fmt.Printf("Signature: %x\n", signature)

	// crypto.VerifySignature takes the signature as [[ R || S ]] whereas crypto.Sign
	// returns the signture as [[ R || S || V ]] (where V is either 0 or 1) so we remove
	// V here to use with VerifySignature
	sigWithoutVSuffix := signature[:len(signature)-1]
	verified := crypto.VerifySignature(
		crypto.FromECDSAPub(pubKeyECDSA),
		//crypto.Keccak256([]byte(messageToSign)),
		messageToSign,
		sigWithoutVSuffix,
	)
	require.True(t, verified, "Message verification failed (using crypto.VerifySignature)")


	recoveredPubKey, err := crypto.SigToPub(
		//crypto.Keccak256([]byte(messageToSign)),
		messageToSign,
		signature,
	)
	require.NoError(t, err)

	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey)
	fmt.Printf("recovered Address: %s\n", recoveredAddress.Hex())

	require.True(t, address==recoveredAddress, "Message verification failed (using crypto.SigToPub)")
}

func TestSignAndVerify_ETH_using_btcec(t *testing.T) {

	//privKey, pubKey := btcec.PrivKeyFromBytes(
	//	btcec.S256(),
	//	pkBytes,
	//)
}
