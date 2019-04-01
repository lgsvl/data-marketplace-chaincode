package utils

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"golang.org/x/oauth2/jws"
)

// ExtendedJwsClaimSet extends JwsClaimSet
type ExtendedJwsClaimSet struct {
	*jws.ClaimSet
	AaccessTokenHash string `json:"at_hash"`
	CustomRoles      string `json:"custom:roles"`
	EventID          string `json:"event_id"`
	TokenUse         string `json:"token_use"`
	AuthTime         int    `json:"auth_time"`
	CognitoUserName  string `json:"cognito:username"`
}
type JOSEHeader struct {
	Kid string `json:"kid"`
	Alg string `json:"alg"`
}

// decodeAlgKidJwsToken takes a JWS Token and decodes it to fetch the signature or the encription algoritm and the token type
func decodeJOSEHeaderFromToken(jwsToken string) (JOSEHeader, error) {
	if jwsToken == "" {
		return JOSEHeader{}, fmt.Errorf("Empty ID token")
	}
	parts := strings.Split(jwsToken, ".")
	if len(parts) < 2 {
		return JOSEHeader{}, fmt.Errorf("There are no two periods in token : ")
	}
	if m := len(parts[0]) % 4; m != 0 {
		parts[0] += strings.Repeat("=", 4-m)
	}

	JOSEHeaderAsArray, err := base64.URLEncoding.DecodeString(parts[0])

	// JOSEHeaderAsArray, err := base64.URLEncoding.DecodeString(parts[0])
	if err != nil {
		return JOSEHeader{}, fmt.Errorf("Impossible to decode token: %v", err)
	}

	var header JOSEHeader
	err = json.Unmarshal(JOSEHeaderAsArray, &header)
	if err != nil {
		return JOSEHeader{}, fmt.Errorf("Impossible to decode JOSE Header in token: %v", err)
	}

	return header, nil
}

// decodeJwsClaimSetFromToken takes a JWS Token and decodes it to fetch the ClaimSet within
func decodeJwsClaimSetFromToken(jwsToken string) (jws.ClaimSet, error) {
	if jwsToken == "" {
		return jws.ClaimSet{}, fmt.Errorf("Empty ID token")
	}
	parts := strings.Split(jwsToken, ".")
	if len(parts) < 2 {
		return jws.ClaimSet{}, fmt.Errorf("There are no two periods in token : ")
	}
	jwsClaimSetAsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return jws.ClaimSet{}, fmt.Errorf("Impossible to decode JWS Claim Set in token: %v", err)
	}
	var set jws.ClaimSet
	err = json.Unmarshal(jwsClaimSetAsBytes, &set)
	if err != nil {
		return jws.ClaimSet{}, fmt.Errorf("Impossible to unmarshal token: %v", err)
	}
	return set, nil
}

// decodeExtendedJwsClaimSetFromToken takes a JWS Token and decodes it to fetch the ExtendedClaimSet within
func decodeExtendedJwsClaimSetFromToken(jwsToken string) (ExtendedJwsClaimSet, error) {
	if jwsToken == "" {
		return ExtendedJwsClaimSet{}, fmt.Errorf("Empty ID token")
	}
	// Check that the padding is correct for a base64decode
	parts := strings.Split(jwsToken, ".")
	if len(parts) < 2 {
		return ExtendedJwsClaimSet{}, fmt.Errorf("There are no two periods in token : ")
	}
	// Decode the ID token
	extendedJwsClaimSetAsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return ExtendedJwsClaimSet{}, fmt.Errorf("Impossible to decode Extended Jws Claim Set in token: %v", err)
	}
	var set ExtendedJwsClaimSet
	err = json.Unmarshal(extendedJwsClaimSetAsBytes, &set)
	if err != nil {
		return ExtendedJwsClaimSet{}, fmt.Errorf("Impossible to unmarshal token: %v", err)
	}
	return set, nil
}
func decodeJwsSignatureFromToken(jwsToken string) (string, error) {
	if jwsToken == "" {
		return "", fmt.Errorf("Empty ID token")
	}
	parts := strings.Split(jwsToken, ".")
	signature := parts[2]
	if len(parts) < 2 {
		return "", fmt.Errorf("There are no two periods in token : ")
	}
	if m := len(signature) % 4; m != 0 {
		signature += strings.Repeat("=", 4-m)
	}
	signatureAsBytes, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return "", fmt.Errorf("Impossible to decode signature in token: %v", err)
	}
	return string(signatureAsBytes[:]), nil
}

func getIssuerURL(issuer string) (string, error) {
	fmt.Print(issuer + "/.well-known/jwks.json\n")
	return issuer + "/.well-known/jwks.json", nil
}

func decodeJws(jwsToken string) (JOSEHeader, ExtendedJwsClaimSet, string, error) {
	headerOauth, err := decodeJOSEHeaderFromToken(jwsToken)
	if err != nil {
		return JOSEHeader{}, ExtendedJwsClaimSet{}, "", err
	}

	extendedJwsClaimSet, err := decodeExtendedJwsClaimSetFromToken(jwsToken)
	if err != nil {
		return JOSEHeader{}, ExtendedJwsClaimSet{}, "", err
	}

	signature, err := decodeJwsSignatureFromToken(jwsToken)
	if err != nil {
		return JOSEHeader{}, ExtendedJwsClaimSet{}, "", err
	}
	return headerOauth, extendedJwsClaimSet, signature, nil
}

func getCertificateFromIssuer(issuerURL string) ([]byte, error) {
	resp, err := http.Get(issuerURL)
	if err != nil {
		return nil, err
	}
	certs, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return certs, nil
}

func getJwsTokenHash(jwsToken string) (string, error) {
	if jwsToken == "" {
		return "", fmt.Errorf("Empty ID token")
	}
	parts := strings.Split(jwsToken, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("There are no two periods in token : ")
	}
	return parts[0] + "." + parts[1], nil
}

func getPublicKey(certs []byte, kid string) (rsa.PublicKey, error) {
	var goCertificate interface{}
	err := json.Unmarshal(certs, &goCertificate)
	k := goCertificate.(map[string]interface{})["keys"]
	j := k.([]interface{})
	x := j[0]

	if j[0].(map[string]interface{})["kid"] == kid {
		x = j[0]
	} else {
		if j[1].(map[string]interface{})["kid"] == kid {
			x = j[1]
		} else {
			errMsg := "Token is not valid, kid from token and certificate don't match"
			return rsa.PublicKey{}, fmt.Errorf(errMsg)
		}
	}
	h := x.(map[string]interface{})["n"]
	g := x.(map[string]interface{})["e"]

	//build the google pub key
	nStr := h.(string)
	// decN, err := base64.URLEncoding.DecodeString(nStr)
	decN, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		return rsa.PublicKey{}, err
	}

	n := big.NewInt(0)
	n.SetBytes(decN)
	eStr := g.(string)
	decE, err := base64.URLEncoding.DecodeString(eStr)
	if err != nil {
		return rsa.PublicKey{}, err
	}

	var eBytes []byte
	if len(decE) < 8 {
		eBytes = make([]byte, 8-len(decE), 8)
		eBytes = append(eBytes, decE...)
	} else {
		eBytes = decE
	}

	eReader := bytes.NewReader(eBytes)
	var e uint64
	err = binary.Read(eReader, binary.BigEndian, &e)
	if err != nil {
		return rsa.PublicKey{}, err
	}
	pKey := rsa.PublicKey{N: n, E: int(e)}
	return pKey, nil
}

func verifySignature(signature string, toHash string, publicKey rsa.PublicKey) error {
	hasherOauth := sha256.New()
	hasherOauth.Write([]byte(toHash))
	return rsa.VerifyPKCS1v15(&publicKey, crypto.SHA256, hasherOauth.Sum(nil), []byte(signature))

}

func verifyJwsToken(jwsToken string) error {
	headerOauth, extendedJwsClaimSet, signature, err := decodeJws(jwsToken)
	if err != nil {
		return err
	}
	issuerURL, err := getIssuerURL(extendedJwsClaimSet.Iss)
	if err != nil {
		return err
	}
	certs, err := getCertificateFromIssuer(issuerURL)
	if err != nil {
		return err
	}
	toHash, err := getJwsTokenHash(jwsToken)
	if err != nil {
		return err
	}
	publicKey, err := getPublicKey(certs, headerOauth.Kid)
	if err != nil {
		return err
	}
	err = verifySignature(signature, toHash, publicKey)
	if err != nil {
		return err
	}
	return nil
}

func CheckAuth(logger *shim.ChaincodeLogger, auth string) error {
	logger.Debug("entering-CheckAuth")
	defer logger.Info("exiting-CheckAuth")
	parts := strings.Split(auth, " ")
	if len(parts) != 2 {
		errMsg := "Token is not valid"
		return fmt.Errorf(errMsg)
	}
	return verifyJwsToken(parts[1])
}

func GetAccountIDFromToken(logger *shim.ChaincodeLogger, auth string) (string, error) {
	logger.Debug("entering-CheckAuth")
	defer logger.Info("exiting-CheckAuth")
	logger.Debugf("token: %s", auth)
	return auth, nil
}
