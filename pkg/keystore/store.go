package keystore

type Keystore interface {
	Get(string) (KeychainInfo, error)
}

// Scheme defines the scheme on which a keychain entry is based.
type Scheme string

const (
	// BIP44 indicates that the keychain scheme is legacy.
	BIP44 Scheme = "BIP44"

	// BIP49 indicates that the keychain scheme is segwit.
	BIP49 Scheme = "BIP49"

	// BIP84 indicates that the keychain scheme is native segwit.
	BIP84 Scheme = "BIP84"
)

type KeychainInfo struct {
	Descriptor              string `json:"descriptor"`
	XPub                    string `json:"xpub"`                       // Extended public key serialized with standard HD version bytes
	SLIP32ExtendedPublicKey string `json:"slip32_extended_public_key"` // Extended public key serialized with SLIP-0132 HD version bytes
	ExternalPublicKey       string `json:"external_public_key"`        // External chain public key at HD tree depth 4
	ExternalChainCode       string `json:"external_chain_code"`        // External chain chain code at HD tree depth 4
	MaxExternalIndex        uint32 `json:"max_external_index"`         // Number of external chain addresses in keychain
	InternalPublicKey       string `json:"internal_public_key"`        // Internal chain public key at HD tree depth 4
	InternalChainCode       string `json:"internal_chain_code"`        // Internal chain chain code at HD tree depth 4
	MaxInternalIndex        uint32 `json:"max_internal_index"`         // Number of external chain addresses in keychain
	LookaheadSize           uint32 `json:"lookahead_size"`             // Numerical size of the lookahead zone
	Scheme                  Scheme `json:"scheme"`                     // String identifier for keychain scheme
}

type derivationToPublicKeyMap map[string]struct {
	PublicKey string `json:"public_key"` // Public key at HD tree depth 5
	Used      bool   `json:"used"`       // Whether any txn history at derivation
}

// KeystoreSchema is a map containing keychain data corresponding to an
// account descriptor.
type KeystoreSchema map[string]struct {
	Main        KeychainInfo             `json:"main"`
	Derivations derivationToPublicKeyMap `json:"derivations"`
	Addresses   map[string][2]uint32     `json:"addresses"` // derivation path at HD tree depth 5
}
