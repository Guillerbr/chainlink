package models

import (
	"chainlink/core/utils"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

// AuthToken ...
type AuthToken struct {
	AccessKey string
	Secret    string
}

// GetID returns the ID of this structure for jsonapi serialization.
func (ta *AuthToken) GetID() string {
	return ta.AccessKey
}

// GetName returns the pluralized "type" of this structure for jsonapi serialization.
func (ta *AuthToken) GetName() string {
	return "external_initiators"
}

// SetID returns the ID of this structure for jsonapi serialization.
func (ta *AuthToken) SetID(id string) error {
	ta.AccessKey = id
	return nil
}

// NewAuthToken ...
func NewAuthToken() *AuthToken {
	return &AuthToken{
		AccessKey: utils.NewBytes32ID(),
		Secret:    utils.NewSecret(utils.DefaultSecretSize),
	}
}

func hashInput(ta *AuthToken, salt string) []byte {
	return []byte(fmt.Sprintf("v0-%s-%s-%s", ta.AccessKey, ta.Secret, salt))
}

// HashedSecret generates a hashed password for an external initiator
// authentication
func HashedSecret(ta *AuthToken, salt string) (string, error) {
	hasher := sha3.New256()
	_, err := hasher.Write(hashInput(ta, salt))
	if err != nil {
		return "", errors.Wrap(err, "error writing external initiator authentication to hasher")
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
