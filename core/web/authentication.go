package web

import (
	"chainlink/core/store"
	"chainlink/core/store/models"
	"chainlink/core/store/orm"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	// ExternalInitiatorAccessKeyHeader is the header name for the access key
	// used by external initiators to authenticate
	ExternalInitiatorAccessKeyHeader = "X-Chainlink-EA-AccessKey"
	// ExternalInitiatorSecretHeader is the header name for the secret used by
	// external initiators to authenticate
	ExternalInitiatorSecretHeader = "X-Chainlink-EA-Secret"
	// APIKey ...
	APIKey = "X-API-KEY"
	// APISecret ...
	APISecret = "X-API-SECRET"
)

var (
	// ErrorAuthFailed is a generic authentication failed - but not because of
	// some system failure on our behalf (i.e. HTTP 5xx), more detail is not
	// given
	ErrorAuthFailed = errors.New("Authentication failed")
)

func authenticatedUser(c *gin.Context) (*models.User, bool) {
	obj, ok := c.Get(SessionUserKey)
	if !ok {
		return nil, false
	}
	return obj.(*models.User), ok
}

func tokenAuth(store *store.Store, c *gin.Context) error {
	eia := &models.AuthToken{
		AccessKey: c.GetHeader(ExternalInitiatorAccessKeyHeader),
		Secret:    c.GetHeader(ExternalInitiatorSecretHeader),
	}

	ei, err := store.FindExternalInitiator(eia)
	if errors.Cause(err) == orm.ErrorNotFound {
		return ErrorAuthFailed
	} else if err != nil {
		return errors.Wrap(err, "finding external intiator")
	}

	ok, err := models.AuthenticateExternalInitiator(eia, ei)
	if err != nil {
		return err
	}

	if !ok {
		return ErrorAuthFailed
	}
	c.Set(SessionExternalInitiatorKey, ei)

	return nil
}

func authenticatedEI(c *gin.Context) (*models.ExternalInitiator, bool) {
	obj, ok := c.Get(SessionExternalInitiatorKey)
	if !ok {
		return nil, false
	}
	return obj.(*models.ExternalInitiator), ok
}

func sessionAuth(store *store.Store, c *gin.Context) error {
	session := sessions.Default(c)
	sessionID, ok := session.Get(SessionIDKey).(string)
	if !ok {
		return ErrorAuthFailed
	}

	user, err := store.AuthorizedUserWithSession(sessionID)
	if err != nil {
		return err
	}
	c.Set(SessionUserKey, &user)
	return nil
}

func sessionAuthRequired(store *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := sessionAuth(store, c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}
	}
}

// sessionOrTokenAuthRequired first tries session authentication, then falls back to
// token authentication, strictly for External Initiators
func sessionOrTokenAuthRequired(store *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := sessionAuth(store, c)
		if err == ErrorAuthFailed {
			err = tokenAuth(store, c)
		}

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}
	}
}
