package application

import (
	"context"
	"encoding/gob"
	"fmt"
	"strings"
	"time"

	"flamingo.me/flamingo-commerce/cart/domain/cart"
	authApplication "flamingo.me/flamingo/core/auth/application"
	"flamingo.me/flamingo/core/auth/domain"
	"flamingo.me/flamingo/framework/flamingo"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type (
	// CartCache describes a cart caches methods
	CartCache interface {
		GetCart(context.Context, *sessions.Session, CartCacheIdentifier) (*cart.Cart, error)
		CacheCart(context.Context, *sessions.Session, CartCacheIdentifier, *cart.Cart) error
		Invalidate(context.Context, *sessions.Session, CartCacheIdentifier) error
		Delete(context.Context, *sessions.Session, CartCacheIdentifier) error
		DeleteAll(context.Context, *sessions.Session) error
		BuildIdentifier(context.Context, *sessions.Session) (CartCacheIdentifier, error)
	}

	// CartCacheIdentifier identifies Cart Caches
	CartCacheIdentifier struct {
		GuestCartID    string
		IsCustomerCart bool
		CustomerID     string
	}

	// CartSessionCache defines a Cart Cache
	CartSessionCache struct {
		authManager     *authApplication.AuthManager
		userService     *authApplication.UserService
		logger          flamingo.Logger
		lifetimeSeconds float64
	}

	// CachedCartEntry defines a single Cart Cache Entry
	CachedCartEntry struct {
		IsInvalid bool
		Entry     cart.Cart
		ExpiresOn time.Time
	}
)

const (
	// CartSessionCacheCacheKeyPrefix is a string prefix for Cart Cache Keys
	CartSessionCacheCacheKeyPrefix = "cart.sessioncache."
)

var (
	_ CartCache = (*CartSessionCache)(nil)
	// ErrCacheIsInvalid sets generalized invalid Cache Error
	ErrCacheIsInvalid = errors.New("cache is invalid")
)

func init() {
	gob.Register(CachedCartEntry{})
}

// CacheKey creates a Cache Key Identifier string
func (ci *CartCacheIdentifier) CacheKey() string {
	return fmt.Sprintf(
		"cart_%v_%v",
		ci.CustomerID,
		ci.GuestCartID,
	)
}

// BuildIdentifierFromCart creates a Cache Identifier from Cart Data
func BuildIdentifierFromCart(cart *cart.Cart) (*CartCacheIdentifier, error) {
	if cart == nil {
		return nil, errors.New("no cart")
	}

	if cart.BelongsToAuthenticatedUser {
		return &CartCacheIdentifier{
			CustomerID:     cart.AuthenticatedUserId,
			IsCustomerCart: true,
		}, nil
	}

	return &CartCacheIdentifier{
		GuestCartID:    cart.ID,
		CustomerID:     cart.AuthenticatedUserId,
		IsCustomerCart: false,
	}, nil
}

// Inject the dependencies
func (cs *CartSessionCache) Inject(
	authManager *authApplication.AuthManager,
	userService *authApplication.UserService,
	logger flamingo.Logger,
	config *struct {
		LifetimeSeconds float64 `inject:"config:cart.cacheLifetime"` // in seconds
	},
) {
	cs.authManager = authManager
	cs.userService = userService
	cs.logger = logger
	if config != nil {
		cs.lifetimeSeconds = config.LifetimeSeconds
	}
}

// auth tries to retrieve the authentication context for a active session
func (cs *CartSessionCache) auth(c context.Context, session *sessions.Session) domain.Auth {
	ts, _ := cs.authManager.TokenSource(c, session)
	idToken, _ := cs.authManager.IDToken(c, session)

	return domain.Auth{
		TokenSource: ts,
		IDToken:     idToken,
	}
}

// BuildIdentifier creates a CartCacheIdentifier based on the login state
func (cs *CartSessionCache) BuildIdentifier(ctx context.Context, session *sessions.Session) (CartCacheIdentifier, error) {
	if cs.userService.IsLoggedIn(ctx, session) {
		return CartCacheIdentifier{
			CustomerID:     cs.auth(ctx, session).IDToken.Subject,
			IsCustomerCart: true,
		}, nil
	}

	guestCartID, ok := session.Values[GuestCartSessionKey]
	if !ok {
		return CartCacheIdentifier{}, errors.New("Fatal - ShouldHaveGuestCart returned true but got no GuestCartSessionKey?")
	}

	guestCartIDString, ok := guestCartID.(string)
	if !ok {
		return CartCacheIdentifier{}, errors.New("Fatal - ShouldHaveGuestCart returned true but got no GuestCartSessionKey string")
	}

	return CartCacheIdentifier{
		GuestCartID: guestCartIDString,
	}, nil
}

// GetCart fetches a Cart from the Cache
func (cs *CartSessionCache) GetCart(ctx context.Context, session *sessions.Session, id CartCacheIdentifier) (*cart.Cart, error) {
	if cache, ok := session.Values[CartSessionCacheCacheKeyPrefix+id.CacheKey()]; ok {
		if cachedCartsEntry, ok := cache.(CachedCartEntry); ok {
			cs.logger.WithField(flamingo.LogKeyCategory, "CartSessionCache").Debug("Found cached cart %v", id.CacheKey())

			if cachedCartsEntry.IsInvalid {
				return &cachedCartsEntry.Entry, ErrCacheIsInvalid
			}

			if time.Now().After(cachedCartsEntry.ExpiresOn) {
				err := cs.Invalidate(ctx, session, id)
				if err != nil {
					return nil, err
				}

				return nil, ErrCacheIsInvalid
			}

			return &cachedCartsEntry.Entry, nil
		}
		cs.logger.WithField(flamingo.LogKeyCategory, "CartSessionCache").Error("Cannot Cast Cache Entry %v", id.CacheKey())

		return nil, errors.New("cart cache contains invalid data at cache key")
	}
	cs.logger.WithField(flamingo.LogKeyCategory, "CartSessionCache").Debug("Did not Found cached cart %v", id.CacheKey())

	return nil, errors.New("no cart in cache")
}

// CacheCart adds a Cart to the Cache
func (cs *CartSessionCache) CacheCart(ctx context.Context, session *sessions.Session, id CartCacheIdentifier, cartForCache *cart.Cart) error {
	if cartForCache == nil {
		return errors.New("no cart given to cache")
	}

	entry := CachedCartEntry{
		Entry:     *cartForCache,
		ExpiresOn: time.Now().Add(time.Duration(cs.lifetimeSeconds * float64(time.Second))),
	}

	cs.logger.WithField(flamingo.LogKeyCategory, "CartSessionCache").Debug("Caching cart %v", id.CacheKey())
	session.Values[CartSessionCacheCacheKeyPrefix+id.CacheKey()] = entry

	return nil
}

// Invalidate a Cache Entry
func (cs *CartSessionCache) Invalidate(ctx context.Context, session *sessions.Session, id CartCacheIdentifier) error {
	if cache, ok := session.Values[CartSessionCacheCacheKeyPrefix+id.CacheKey()]; ok {
		if cachedCartsEntry, ok := cache.(CachedCartEntry); ok {
			cachedCartsEntry.IsInvalid = true
			session.Values[CartSessionCacheCacheKeyPrefix+id.CacheKey()] = cachedCartsEntry

			return nil
		}
	}

	return errors.New("not found for invalidate")
}

// Delete a Cache entry
func (cs *CartSessionCache) Delete(ctx context.Context, session *sessions.Session, id CartCacheIdentifier) error {
	if _, ok := session.Values[CartSessionCacheCacheKeyPrefix+id.CacheKey()]; ok {
		delete(session.Values, CartSessionCacheCacheKeyPrefix+id.CacheKey())

		// ok deleted something
		return nil
	}

	return errors.New("not found for delete")
}

// DeleteAll empties the Cache
func (cs *CartSessionCache) DeleteAll(ctx context.Context, session *sessions.Session) error {
	deleted := false
	for k := range session.Values {
		if stringKey, ok := k.(string); ok {
			if strings.Contains(stringKey, CartSessionCacheCacheKeyPrefix) {
				delete(session.Values, k)
				deleted = true
			}
		}
	}

	if deleted {
		// successfully deleted something
		return nil
	}

	return errors.New("not found for delete")
}
