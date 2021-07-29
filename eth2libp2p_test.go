package eth2libp2p

import (
	"testing"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEth2LibP2p(t *testing.T) {
	privKeysHex := []string{
		"e6ee963a92cf6969f84d31f987e2335fafc0deae205eb558d793c0ca3bd798c5",
		"609ed8a006b549df11ed87ae343220832d28ae5c49e710732817f8f3e61dec7d",
		"2caf621cc0c168e17bbed4218fc8c7c9b3252eb15f8d680309f89bb14eeb1bbe",
		"76be4211e7e0c5045c0f76d2ebd79a311318891e1db0ccc54b87fd559864a34d",
		"d4950396b5431646ae5a397dc39eb4dcd96a6f922e30898097202a5dfb85c0ff",
	}
	for _, privKeyHex := range privKeysHex {
		privKey, err := ethCrypto.HexToECDSA(privKeyHex)
		require.NoError(t, err)
		wallet, err := NewLibP2PIdentityFromEthPrivKey(privKey)
		require.NoError(t, err)
		id, err := P2PIDFromEthPubKey(&privKey.PublicKey)
		require.NoError(t, err)
		assert.Equal(t, wallet.ID, id)
	}
}
