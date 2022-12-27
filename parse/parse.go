// Package parse otp secret from the migration QR code of Google authenticator
package parse

import (
	"encoding/base32"
	"encoding/base64"
	"errors"
	"github.com/dchaofei/gauth-exporter/internal/googleauth"
	"github.com/golang/protobuf/proto"
	"net/url"
)

// Payload otp payload
type Payload struct {
	Secret    string `json:"secret"`
	Name      string `json:"name"`
	Issuer    string `json:"issuer"`
	Algorithm string `json:"algorithm"`
	Digits    int32  `json:"digits"`
	Type      string `json:"type"`
}

// FromUri parse payload from uri
// uri parameter is the data after decoding the migrated QR code:
// example: FromUri("otpauth-migration://offline?data=xxxxxxxx")
func FromUri(uri string) ([]Payload, error) {
	if uri == "" {
		return nil, errors.New("uri is empty")
	}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	return FromBase64(u.Query().Get("data"))
}

// FromBase64 parse payload from base64
func FromBase64(s string) ([]Payload, error) {
	protoBuffer, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return FromRawBytes(protoBuffer)
}

// FromRawBytes parse from the raw bytes
func FromRawBytes(buffer []byte) ([]Payload, error) {
	var payloads []Payload

	mp := googleauth.MigrationPayload{}
	if err := proto.NewBuffer(buffer).Unmarshal(&mp); err != nil {
		return nil, err
	}

	for _, v := range mp.OtpParameters {
		payloads = append(payloads, Payload{
			Secret:    base32.StdEncoding.EncodeToString(v.Secret),
			Name:      v.Name,
			Issuer:    v.Issuer,
			Algorithm: v.Algorithm.String(),
			Digits:    v.Digits,
			Type:      v.Type.String(),
		})
	}
	return payloads, nil
}
