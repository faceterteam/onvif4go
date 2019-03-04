package soap

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"io"
	"time"
)

type wsSecurity struct {
	XMLName       xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`
	UsernameToken wssUsernameToken
}

type wssUsernameToken struct {
	XMLName  xml.Name    `xml:"UsernameToken"`
	Username string      `xml:"Username"`
	Password wssPassword `xml:"Password"`
	Nonce    wssNonce    `xml:"Nonce"`
	Created  string      `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Created"`
}

type wssPassword struct {
	Type     string `xml:"Type,attr"`
	Password string `xml:",chardata"`
}

type wssNonce struct {
	Type  string `xml:"EncodingType,attr"`
	Nonce string `xml:",chardata"`
}

func newUUIDVer4() ([]byte, error) {
	u := new([16]byte)
	if _, err := io.ReadFull(rand.Reader, u[:]); err != nil {
		return u[:], err
	}
	// u.SetVersion(V4)
	// u.SetVariant(VariantRFC4122)

	return u[:], nil
}

func MakeWSSecurity(username string, password string, timeDiff time.Duration) wsSecurity {
	created := time.Now().UTC().Add(timeDiff).Format(time.RFC3339Nano)

	nonce, _ := newUUIDVer4()
	nonce64 := base64.StdEncoding.EncodeToString(nonce)
	hasher := sha1.New()
	hasher.Write(nonce)
	hasher.Write([]byte(created + password))
	shaToken := hasher.Sum(nil)
	shaDigest64 := base64.StdEncoding.EncodeToString(shaToken)

	return wsSecurity{
		UsernameToken: wssUsernameToken{
			Username: username,
			Created:  created,
			Password: wssPassword{
				Type:     "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest",
				Password: shaDigest64,
			},
			Nonce: wssNonce{
				Type:  "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary",
				Nonce: nonce64,
			},
		},
	}
}
