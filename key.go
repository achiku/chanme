package chanme

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/ripemd160"

	"github.com/btcsuite/btcutil/base58"
)

// bas58 encode prefix
const (
	VerMainPub  = 0
	VerMainPriv = 128
	VerTestPub  = 111
	VerTestPriv = 239
)

// PrivateKey private key
type PrivateKey struct {
	*ecdsa.PrivateKey
}

// GenerateNewPrivateKey generate private key
func GenerateNewPrivateKey(seed io.Reader) (*PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), seed)
	if err != nil {
		return nil, err
	}
	pk := &PrivateKey{
		PrivateKey: priv,
	}
	return pk, nil
}

// Hex hex format of private key
func (pk *PrivateKey) Hex() []byte {
	return []byte(fmt.Sprintf("%x", pk.D))
}

// RawKey raw private key
func (pk *PrivateKey) RawKey() []byte {
	return []byte(fmt.Sprintf("%s", pk.D))
}

// Base58 baes58 private key
func (pk *PrivateKey) Base58() string {
	return base58.Encode(pk.D.Bytes())
}

// Base58Check baes58 private key
func (pk *PrivateKey) Base58Check(ver byte) string {
	return base58.CheckEncode(pk.D.Bytes(), ver)
}

// PublicKey public key
type PublicKey struct {
	*ecdsa.PublicKey
}

// NewPublicKey new public key
func NewPublicKey(priv *PrivateKey) *PublicKey {
	p := priv.Public().(*ecdsa.PublicKey)
	pub := &PublicKey{
		PublicKey: p,
	}
	return pub
}

// BitcoinAddress bitcoin address
func (pub *PublicKey) BitcoinAddress(ver byte) string {
	sh := sha256.Sum256(pub.X.Bytes())
	r := ripemd160.New()
	var buf []byte
	for _, b := range sh {
		buf = append(buf, b)
	}
	b := r.Sum(buf)
	return base58.CheckEncode(b, ver)
}
