/*
Copyright (c) 2025 Red Hat Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the
License. You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific
language governing permissions and limitations under the License.
*/

package testing

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	. "github.com/onsi/gomega"
)

// MakeTokenObject generates a token with the claims resulting from merging the default claims and the claims explicitly
// given.
func MakeTokenObject(claims jwt.MapClaims) *jwt.Token {
	merged := jwt.MapClaims{}
	for name, value := range MakeClaims() {
		merged[name] = value
	}
	for name, value := range claims {
		if value == nil {
			delete(merged, name)
		} else {
			merged[name] = value
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, merged)
	token.Header["kid"] = "123"
	var err error
	token.Raw, err = token.SignedString(jwtPrivateKey)
	Expect(err).ToNot(HaveOccurred())
	return token
}

// MakeClaims generates a default set of claims to be used to issue a token.
func MakeClaims() jwt.MapClaims {
	iat := time.Now()
	exp := iat.Add(1 * time.Minute)
	return jwt.MapClaims{
		"iss": "https://example.com/auth/realms/my",
		"iat": iat.Unix(),
		"typ": "Bearer",
		"sub": "mysubject",
		"exp": exp.Unix(),
	}
}

// MakeTokenString generates a token issued by the default OpenID server and with the given type and with the given
// life. If the life is zero the token will never expire. If the life is positive the token will be valid, and expire
// after that time. If the life is negative the token will be already expired that time ago.
func MakeTokenString(typ string, life time.Duration) string {
	claims := jwt.MapClaims{}
	claims["typ"] = typ
	if life != 0 {
		claims["exp"] = time.Now().Add(life).Unix()
	}
	token := MakeTokenObject(claims)
	return token.Raw
}

// Public and private key that will be used to sign and verify tokens in the tests:
var (
	jwtPrivateKey *rsa.PrivateKey
)

func init() {
	var err error

	// Generate the keys used to sign and verify tokens:
	jwtPrivateKey, err = rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
}
