package oauth

// Jwks https://www.rfc-editor.org/rfc/rfc7517#section-2
type Jwks struct {
	// Cryptographic Algorithm
	// e.g. "RSA" or "EC"
	KeyType string `json:"kty"`

	// Intended use of public key (encrypting or verifying signature) and should NOT be used with the "key_ops" field
	// e.g. "sig" or "enc"
	PublicKeyUse string `json:"use"`

	// Identifies operation(s) for which the key is intended to be used and should NOT be used with the "use" field
	//	o  "sign" (compute digital signature or MAC)
	//	o  "verify" (verify digital signature or MAC)
	//	o  "encrypt" (encrypt content)
	//	o  "decrypt" (decrypt content and validate decryption, if applicable)
	//	o  "wrapKey" (encrypt key)
	//	o  "unwrapKey" (decrypt key and validate decryption, if applicable)
	//	o  "deriveKey" (derive key)
	//	o  "deriveBits" (derive bits not to be used as a key)
	KeyOperations []string `json:"key_ops"`

	// Algorithm to be used with the key
	Algorithm string `json:"alg"`

	// Used to match a specific key
	KeyId string `json:"kid"`

	// Refers to a resource for an X509 public key certificate or certificate chain
	X509Url string `json:"x5u"`

	// Contains chain of 1+ PKIX certificates represented as JSON array of certificate value strings
	X509CertificateChain string `json:"x5c"`

	// Base64 Url encoded SHA-1 thumbprint (digest) of the DER encoding of X509 Certificate
	X509Thumbprint string `json:"x5t"`

	// Base64 Url encoded SHA-256 thumbprint (digest) of the DER encoding of X509 Certificate
	X509CertificateThumbprint string `json:"x5t#S256"`
}
