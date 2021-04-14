package jwt

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/cristalhq/jwt/v3"
	"github.com/loadimpact/k6/js/common"
	"github.com/loadimpact/k6/js/modules"
)

func init() {
	modules.Register("k6/x/jwt", new(Jwt))
}

type Jwt struct{}

type Signer struct {
	signer jwt.Signer
}

func (j *Jwt) XSigner(ctxPtr *context.Context, priv []byte) interface{} {
	rt := common.GetRuntime(*ctxPtr)

	privPem, _ := pem.Decode(priv)

	privPemBytes := privPem.Bytes

	var parsedKey interface{}
	parsedKey, _ = x509.ParsePKCS1PrivateKey(privPemBytes)
	privateKey, _ := parsedKey.(*rsa.PrivateKey)

	signer, _ := jwt.NewSignerRS(jwt.RS256, privateKey)

	return common.Bind(rt, &Signer{signer: signer}, ctxPtr)
}

func (c *Signer) Sign(claims interface{}, kid string) string {
	builder := jwt.NewBuilder(c.signer, jwt.WithKeyID(kid))

	token, _ := builder.Build(claims)

	return token.String()
}
