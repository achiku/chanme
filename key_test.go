package chanme

import (
	"crypto/rand"
	"testing"

	"github.com/btcsuite/btcutil/base58"
)

func TestGeneratePrivateKey(t *testing.T) {
	k, err := GenerateNewPrivateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("raw private key: %s", k.RawKey())
	// 85921530139364424125770269749103022238451965595622540103774140775402382608027
	t.Logf("hex format private key: %s", k.Hex())
	// 1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd
	t.Logf("base58 encoded private key: %s", k.Base58())
	// GgUdt6PYJrNLJsdec3LKcQKwZU5MtgsNag91ZGZn9rRb
	t.Logf("base58 check encoded private key: %s", k.Base58Check(VerMainPriv))
	// 5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn
	a, b, err := base58.CheckDecode(k.Base58Check(VerMainPriv))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("hex: %x, version: %d", a, b)
}

func TestGeneratePublicKey(t *testing.T) {
	priv, err := GenerateNewPrivateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	pub := NewPublicKey(priv)
	t.Logf("%d", pub.PublicKey.X)
	t.Logf("%s", pub.BitcoinAddress(VerMainPub))
}
