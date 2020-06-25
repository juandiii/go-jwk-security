package security

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

type JwtKeys struct {
	JwtURL    string
	cachedSet *jwk.Set
}

func (j *JwtKeys) GetKeys() error {
	if j.cachedSet != nil {
		return nil
	}

	fmt.Println("Connecting :: Keycloak")
	set, err := jwk.FetchHTTP(j.JwtURL)

	if err != nil {
		return errors.New("Couldn't connect Keycloak, try again")
	}
	fmt.Println("Connected successfully :: Keycloak")
	j.cachedSet = set

	return nil
}

func (j *JwtKeys) GetKey(token *jwt.Token) (interface{}, error) {

	keyID, ok := token.Header["kid"].(string)

	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	if key := j.cachedSet.LookupKeyID(keyID); len(key) == 1 {
		var raw interface{}
		return raw, key[0].Raw(&raw)
	}

	return nil, fmt.Errorf("unable to find key %q", keyID)
}
