package eth2libp2p

import (
	"crypto/ecdsa"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	libp2p "github.com/libp2p/go-libp2p"
	p2pCrypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

type LibP2PIdentity struct {
	PrivateKey     p2pCrypto.PrivKey
	PublicKey      p2pCrypto.PubKey
	ID             peer.ID
	IdentityOption libp2p.Option // Used for host creation using this identity
}

func NewLibP2PIdentityFromEthPrivKey(ethPrivKey *ecdsa.PrivateKey) (wallet LibP2PIdentity, err error) {
	privKeyData := ethCrypto.FromECDSA(ethPrivKey)
	privKey, err := p2pCrypto.UnmarshalSecp256k1PrivateKey(privKeyData)
	if err != nil {
		return
	}
	pubKey := privKey.GetPublic()
	id, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		return
	}
	wallet = LibP2PIdentity{
		PrivateKey:     privKey,
		PublicKey:      pubKey,
		ID:             id,
		IdentityOption: libp2p.Identity(privKey),
	}
	return
}

func P2PIDFromEthPubKey(ethPubKey *ecdsa.PublicKey) (id peer.ID, err error) {
	pubKeyData := ethCrypto.FromECDSAPub(ethPubKey)
	pubKey, err := p2pCrypto.UnmarshalSecp256k1PublicKey(pubKeyData)
	if err != nil {
		return
	}
	id, err = peer.IDFromPublicKey(pubKey)
	return
}
