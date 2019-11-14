package models

import (
	"crypto/subtle"
	"strings"

	"chainlink/core/utils"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// ExternalInitiatorRequest is the incoming record used to create an ExternalInitiator.
type ExternalInitiatorRequest struct {
	Name string `json:"name"`
	URL  WebURL `json:"url"`
}

// ExternalInitiator represents a user that can initiate runs remotely
type ExternalInitiator struct {
	*gorm.Model
	Name           string `gorm:"not null;unique"`
	URL            WebURL `gorm:"not null"`
	AccessKey      string `gorm:"not null"`
	Salt           string `gorm:"not null"`
	HashedSecret   string `gorm:"not null"`
	OutgoingSecret string `gorm:"not null"`
	OutgoingToken  string `gorm:"not null"`
}

// NewExternalInitiator generates an ExternalInitiator from an
// AuthToken, hashing the password for storage
func NewExternalInitiator(
	eia *AuthToken,
	eir *ExternalInitiatorRequest,
) (*ExternalInitiator, error) {
	salt := utils.NewSecret(utils.DefaultSecretSize)
	hashedSecret, err := HashedSecret(eia, salt)
	if err != nil {
		return nil, errors.Wrap(err, "error hashing secret for external initiator")
	}

	return &ExternalInitiator{
		Name:           strings.ToLower(eir.Name),
		URL:            eir.URL,
		AccessKey:      eia.AccessKey,
		HashedSecret:   hashedSecret,
		Salt:           salt,
		OutgoingToken:  utils.NewSecret(utils.DefaultSecretSize),
		OutgoingSecret: utils.NewSecret(utils.DefaultSecretSize),
	}, nil
}

// AuthenticateExternalInitiator compares an auth against an initiator and
// returns true if the password hashes match
func AuthenticateExternalInitiator(eia *AuthToken, ea *ExternalInitiator) (bool, error) {
	hashedSecret, err := HashedSecret(eia, ea.Salt)
	if err != nil {
		return false, err
	}
	return subtle.ConstantTimeCompare([]byte(hashedSecret), []byte(ea.HashedSecret)) == 1, nil
}
