package cache

import (
	"github.com/uncut-fm/uncut-common/model"
	"sync"
)

type InmemoryCache struct {
	emailByIDMap       map[int]string
	emailByWalletMap   map[string]string
	emailByIDMutex     *sync.RWMutex
	emailByWalletMutex *sync.RWMutex
}

func NewInMemoryCache() *InmemoryCache {
	return &InmemoryCache{
		emailByIDMap:       make(map[int]string),
		emailByWalletMap:   make(map[string]string),
		emailByIDMutex:     new(sync.RWMutex),
		emailByWalletMutex: new(sync.RWMutex),
	}
}

func (c InmemoryCache) GetEmailByIDCache(userID int) (string, bool) {
	c.emailByIDMutex.RLock()
	defer c.emailByIDMutex.RUnlock()
	email, ok := c.emailByIDMap[userID]

	return email, ok
}

func (c InmemoryCache) GetEmailByWalletCache(wallet string) (string, bool) {
	c.emailByWalletMutex.RLock()
	defer c.emailByWalletMutex.RUnlock()
	email, ok := c.emailByWalletMap[wallet]

	return email, ok
}

func (c InmemoryCache) SetUserEmailToCache(user *model.User) {
	c.setEmailToCacheByID(user.ID, user.Email)
	c.setEmailToCacheByWallet(user.WalletAddresses, user.Email)
}

func (c InmemoryCache) setEmailToCacheByID(userID int, email string) {
	c.emailByIDMutex.Lock()
	defer c.emailByIDMutex.Unlock()
	c.emailByIDMap[userID] = email
}

func (c InmemoryCache) setEmailToCacheByWallet(wallets []string, email string) {
	c.emailByWalletMutex.Lock()
	defer c.emailByWalletMutex.Unlock()
	for _, wallet := range wallets {
		c.emailByWalletMap[wallet] = email
	}
}
