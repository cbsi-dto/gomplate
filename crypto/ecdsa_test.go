package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func genECDSAPrivKey() (*ecdsa.PrivateKey, string) {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	der, _ := x509.MarshalECPrivateKey(priv)
	privBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der,
	}
	return priv, string(pem.EncodeToMemory(privBlock))
}

func deriveECPubkey(priv *ecdsa.PrivateKey) string {
	b, _ := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}
	testPubKey := string(pem.EncodeToMemory(pubBlock))
	return testPubKey
}

func TestECDSAGenerateKey(t *testing.T) {
	key, err := ECDSAGenerateKey("P-224")
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(string(key),
		"-----BEGIN EC PRIVATE KEY-----"))
	assert.True(t, strings.HasSuffix(string(key),
		"-----END EC PRIVATE KEY-----\n"))

	key, err = ECDSAGenerateKey("P-256")
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(string(key),
		"-----BEGIN EC PRIVATE KEY-----"))
	assert.True(t, strings.HasSuffix(string(key),
		"-----END EC PRIVATE KEY-----\n"))

	key, err = ECDSAGenerateKey("P-384")
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(string(key),
		"-----BEGIN EC PRIVATE KEY-----"))
	assert.True(t, strings.HasSuffix(string(key),
		"-----END EC PRIVATE KEY-----\n"))

	key, err = ECDSAGenerateKey("P-521")
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(string(key),
		"-----BEGIN EC PRIVATE KEY-----"))
	assert.True(t, strings.HasSuffix(string(key),
		"-----END EC PRIVATE KEY-----\n"))

	_, err = ECDSAGenerateKey("P-999")
	assert.Error(t, err)
}

func TestECDSADerivePublicKey(t *testing.T) {
	_, err := ECDSADerivePublicKey(nil)
	assert.Error(t, err)

	_, err = ECDSADerivePublicKey([]byte(`-----BEGIN FOO-----
	-----END FOO-----`))
	assert.Error(t, err)

	priv, privKey := genECDSAPrivKey()
	expected := deriveECPubkey(priv)

	actual, err := ECDSADerivePublicKey([]byte(privKey))
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}
