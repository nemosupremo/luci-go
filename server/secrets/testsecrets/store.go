// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package testsecrets provides a dumb in-memory secret store to use in unit
// tests. Use secrets.Set(c, &testsecrets.Store{...}) to inject it into
// the context.
package testsecrets

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/luci/luci-go/server/secrets"
	"golang.org/x/net/context"
)

// Store implements secrets.Store in the simplest way possible using memory as
// a backend and very dumb deterministic "randomness" source for secret key
// autogeneration. Useful in unit tests. Can be modified directly (use lock if
// doing it concurrently). NEVER use it outside of tests.
type Store struct {
	sync.Mutex

	Secrets        map[secrets.Key]secrets.Secret // current map of all secrets
	NoAutogenerate bool                           // if true, GetSecret will NOT generate secrets
	SecretLen      int                            // length of generated secret, 8 bytes default
	Rand           *rand.Rand                     // used to generate missing secrets

	counter int // increased with each autogenerated key
}

// GetSecret is a part of Store interface.
func (t *Store) GetSecret(k secrets.Key) (secrets.Secret, error) {
	t.Lock()
	defer t.Unlock()

	if s, ok := t.Secrets[k]; ok {
		return s.Clone(), nil
	}

	if t.NoAutogenerate {
		return secrets.Secret{}, secrets.ErrNoSuchSecret
	}

	// Initialize defaults.
	if t.Secrets == nil {
		t.Secrets = map[secrets.Key]secrets.Secret{}
	}
	if t.SecretLen == 0 {
		t.SecretLen = 8
	}
	if t.Rand == nil {
		t.Rand = rand.New(rand.NewSource(0))
	}

	// Generate deterministic secret.
	t.counter++
	secret := make([]byte, t.SecretLen)
	for i := range secret {
		secret[i] = byte(t.Rand.Int31n(256))
	}
	t.Secrets[k] = secrets.Secret{
		Current: secrets.NamedBlob{
			ID:   fmt.Sprintf("secret_%d", t.counter),
			Blob: secret,
		},
	}
	return t.Secrets[k].Clone(), nil
}

// Use installs default testing store into the context.
func Use(c context.Context) context.Context {
	return secrets.Set(c, &Store{})
}
