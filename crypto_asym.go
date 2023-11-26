package engine

import (
	"crypto"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/ecies/go/v2"
	"golang.org/x/crypto/ed25519"
	"hash"
)

type RSA struct {
}

func (RSA) Generate(bits int) *rsa.PrivateKey {
	return fn.Panic1(rsa.GenerateKey(rand.Reader, bits))
}
func (RSA) MarshalPrivateKey(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}
func (RSA) ParsePrivateKey(key []byte) *rsa.PrivateKey {
	return fn.Panic1(x509.ParsePKCS1PrivateKey(key))
}
func (RSA) MarshalPublicKey(key *rsa.PublicKey) []byte {
	return x509.MarshalPKCS1PublicKey(key)
}
func (RSA) ParsePublicKey(key []byte) *rsa.PublicKey {
	return fn.Panic1(x509.ParsePKCS1PublicKey(key))
}

func (RSA) Sign(key *rsa.PrivateKey, hash crypto.Hash, data []byte, opt *rsa.PSSOptions) []byte {
	if opt == nil {
		return fn.Panic1(rsa.SignPKCS1v15(rand.Reader, key, hash, data))
	} else {
		return fn.Panic1(rsa.SignPSS(rand.Reader, key, hash, data, opt))
	}

}
func (RSA) Verify(key *rsa.PublicKey, hash crypto.Hash, data, sign []byte, opt *rsa.PSSOptions) bool {
	if opt == nil {
		err := rsa.VerifyPKCS1v15(key, hash, data, sign)
		if err != nil {
			if errors.Is(err, rsa.ErrVerification) {
				return false
			} else {
				panic(err)
			}
		} else {
			return true
		}
	} else {
		err := rsa.VerifyPSS(key, hash, data, sign, opt)
		if err != nil {
			if errors.Is(err, rsa.ErrVerification) {
				return false
			} else {
				panic(err)
			}
		} else {
			return true
		}
	}
}

func (RSA) EncryptPKCS1v15(pk *rsa.PublicKey, msg []byte) []byte {
	return fn.Panic1(rsa.EncryptPKCS1v15(rand.Reader, pk, msg))
}
func (RSA) DecryptPKCS1v15(pk *rsa.PrivateKey, msg []byte) []byte {
	return fn.Panic1(rsa.DecryptPKCS1v15(rand.Reader, pk, msg))
}
func (RSA) DecryptOAEP(pk *rsa.PrivateKey, hashFn hash.Hash, msg []byte, label []byte) []byte {
	return fn.Panic1(rsa.DecryptOAEP(hashFn, rand.Reader, pk, msg, label))
}
func (RSA) EncryptOAEP(pk *rsa.PublicKey, hashFn hash.Hash, msg []byte, label []byte) []byte {
	return fn.Panic1(rsa.EncryptOAEP(hashFn, rand.Reader, pk, msg, label))
}

type Algorithm int

const (
	algUndefined Algorithm = iota
	algRSA
	algECDH
	algECDSA
	algED25519
)

type AsymmetricPrivateKey struct {
	Alg    Algorithm
	key    crypto.PrivateKey
	pubKey crypto.PublicKey // only for ED25519
}

func (s *AsymmetricPrivateKey) Bytes() []byte {
	return fn.Panic1(x509.MarshalPKCS8PrivateKey(s.key))
}
func (s *AsymmetricPrivateKey) Load(pem []byte) {
	k := fn.Panic1(x509.ParsePKCS8PrivateKey(pem))
	switch v := k.(type) {
	case *rsa.PrivateKey:
		{
			s.key = v
			s.Alg = algRSA
		}
	case *ecdsa.PrivateKey:
		{
			s.key = v
			s.Alg = algECDSA
		}
	case *ecdh.PrivateKey:
		{
			s.key = v
			s.Alg = algECDH
		}
	case ed25519.PrivateKey:
		{
			s.key = v
			s.Alg = algED25519
		}
	default:
		panic(fmt.Errorf("unknown key type : %T", v))
	}
}
func (s *AsymmetricPrivateKey) Equal(other *AsymmetricPrivateKey) bool {
	if s == nil || other == nil {
		return false
	}
	if s.Alg != other.Alg {
		return false
	}
	if s.key == nil || other.key == nil {
		return false
	}
	switch k := s.key.(type) {
	case *rsa.PrivateKey:
		return k.Equal(other.key)
	case *ecdsa.PrivateKey:
		return k.Equal(other.key)
	case *ecdh.PrivateKey:
		return k.Equal(other.key)
	case ed25519.PrivateKey:
		return k.Equal(other.key) && s.pubKey.(ed25519.PublicKey).Equal(other.pubKey)
	default:
		panic(fmt.Errorf("unknown key type : %T", s.key))
	}
}
func (s *AsymmetricPrivateKey) PublicKey() *AsymmetricPublicKey {
	switch k := s.key.(type) {
	case *rsa.PrivateKey:
		return &AsymmetricPublicKey{
			Alg: s.Alg,
			key: &k.PublicKey,
		}
	case *ecdsa.PrivateKey:
		return &AsymmetricPublicKey{
			Alg: s.Alg,
			key: &k.PublicKey,
		}
	case ed25519.PrivateKey:
		return &AsymmetricPublicKey{
			Alg: s.Alg,
			key: s.pubKey,
		}
	case *ecdh.PrivateKey:
		return &AsymmetricPublicKey{
			Alg: s.Alg,
			key: k.PublicKey(),
		}
	default:
		panic(fmt.Errorf("unknown key type : %T", s.key))
	}
}
func (s *AsymmetricPrivateKey) mustRSA() *rsa.PrivateKey {
	if s.Alg != algRSA {
		panic("require a RSA key")
	}
	if k, ok := s.key.(*rsa.PrivateKey); !ok {
		panic("require a RSA key")
	} else {
		return k
	}
}

func (s *AsymmetricPrivateKey) mustECDSA() *ecdsa.PrivateKey {
	if s.Alg != algECDSA {
		panic("require a ECDSA key")
	}
	if k, ok := s.key.(*ecdsa.PrivateKey); !ok {
		panic("require a ECDSA key")
	} else {
		return k
	}
}

func (s *AsymmetricPrivateKey) mustED25519() ed25519.PrivateKey {
	if s.Alg != algED25519 {
		panic("require a ED25519 key")
	}
	if k, ok := s.key.(ed25519.PrivateKey); !ok {
		panic("require a ED25519 key")
	} else {
		return k
	}
}
func (s *AsymmetricPrivateKey) mustECDH() *ecdh.PrivateKey {
	if s.Alg != algECDH {
		panic("require a ECDH key")
	}
	if k, ok := s.key.(*ecdh.PrivateKey); !ok {
		panic("require a ECDH key")
	} else {
		return k
	}
}

func (s *AsymmetricPrivateKey) mustECISE() *eciesgo.PrivateKey {
	k := s.mustECDSA()
	return &eciesgo.PrivateKey{
		PublicKey: &eciesgo.PublicKey{
			Curve: k.Curve,
			X:     k.X,
			Y:     k.Y,
		},
		D: k.D,
	}
}

type AsymmetricPublicKey struct {
	Alg Algorithm
	key crypto.PublicKey
}

func (s *AsymmetricPublicKey) Bytes() []byte {
	return fn.Panic1(x509.MarshalPKIXPublicKey(s.key))
}
func (s *AsymmetricPublicKey) Load(pem []byte) {
	k := fn.Panic1(x509.ParsePKIXPublicKey(pem))
	switch v := k.(type) {
	case *rsa.PublicKey:
		{
			s.key = v
			s.Alg = algRSA
		}
	case *ecdsa.PublicKey:
		{
			s.key = v
			s.Alg = algECDSA
		}
	case *ecdh.PublicKey:
		{
			s.key = v
			s.Alg = algECDH
		}
	case ed25519.PublicKey:
		{
			s.key = v
			s.Alg = algED25519
		}
	default:
		panic(fmt.Errorf("unknown key type : %T", v))
	}
}
func (s *AsymmetricPublicKey) Equal(other *AsymmetricPublicKey) bool {
	if s == nil || other == nil {
		return false
	}
	if s.Alg != other.Alg {
		return false
	}
	if s.key == nil || other.key == nil {
		return false
	}
	switch k := s.key.(type) {
	case *rsa.PublicKey:
		return k.Equal(other.key)
	case *ecdsa.PublicKey:
		return k.Equal(other.key)
	case *ecdh.PublicKey:
		return k.Equal(other.key)
	case ed25519.PublicKey:
		return k.Equal(other.key)
	default:
		panic(fmt.Errorf("unknown key type : %T", s.key))
	}
}
func (s *AsymmetricPublicKey) mustRSA() *rsa.PublicKey {
	if s.Alg != algRSA {
		panic("require a RSA  key")
	}
	if k, ok := s.key.(*rsa.PublicKey); !ok {
		panic("require a RSA  key")
	} else {
		return k
	}
}

func (s *AsymmetricPublicKey) mustECDSA() *ecdsa.PublicKey {
	if s.Alg != algECDSA {
		panic("require a ECDSA key")
	}
	if k, ok := s.key.(*ecdsa.PublicKey); !ok {
		panic("require a ECDSA key")
	} else {
		return k
	}
}
func (s *AsymmetricPublicKey) mustECISE() *eciesgo.PublicKey {
	k := s.mustECDSA()
	return &eciesgo.PublicKey{
		Curve: k.Curve,
		X:     k.X,
		Y:     k.Y,
	}
}
func (s *AsymmetricPublicKey) mustED25519() ed25519.PublicKey {
	if s.Alg != algED25519 {
		panic("require a ED25519 key")
	}
	if k, ok := s.key.(ed25519.PublicKey); !ok {
		panic("require a ED25519 key")
	} else {
		return k
	}
}

func (s *AsymmetricPublicKey) mustECDH() *ecdh.PublicKey {
	if s.Alg != algECDH {
		panic("require a ECDH key")
	}
	if k, ok := s.key.(*ecdh.PublicKey); !ok {
		panic("require a ECDH key")
	} else {
		return k
	}
}

type (
	Asymmetric struct {
	}
	Option struct {
		Bits  int
		Curve string
	}
)

func (Asymmetric) GenerateKey(algorithm Algorithm, option Option) *AsymmetricPrivateKey {
	switch algorithm {
	case algRSA:
		return &AsymmetricPrivateKey{
			Alg: algorithm,
			key: fn.Panic1(rsa.GenerateKey(rand.Reader, option.Bits)),
		}
	case algECDH:
		switch option.Curve {
		case "P256":
			return &AsymmetricPrivateKey{
				Alg: algorithm,
				key: fn.Panic1(ecdh.P256().GenerateKey(rand.Reader)),
			}
		case "P384":
			return &AsymmetricPrivateKey{
				Alg: algorithm,
				key: fn.Panic1(ecdh.P384().GenerateKey(rand.Reader)),
			}
		case "P521":
			return &AsymmetricPrivateKey{
				Alg: algorithm,
				key: fn.Panic1(ecdh.P521().GenerateKey(rand.Reader)),
			}
		case "X25519":
			return &AsymmetricPrivateKey{
				Alg: algorithm,
				key: fn.Panic1(ecdh.X25519().GenerateKey(rand.Reader)),
			}
		default:
			panic(fmt.Errorf("unknown ECDH curve type : %s", option.Curve))
		}
	case algECDSA:
		switch option.Curve {
		case "P256":
			return &AsymmetricPrivateKey{
				Alg: algorithm,
				key: fn.Panic1(ecdsa.GenerateKey(elliptic.P256(), rand.Reader)),
			}
		case "P384":
			return &AsymmetricPrivateKey{
				Alg: algorithm,
				key: fn.Panic1(ecdsa.GenerateKey(elliptic.P384(), rand.Reader)),
			}
		case "P521":
			return &AsymmetricPrivateKey{
				Alg: algorithm,
				key: fn.Panic1(ecdsa.GenerateKey(elliptic.P521(), rand.Reader)),
			}
		case "P224":
			return &AsymmetricPrivateKey{
				Alg: algorithm,
				key: fn.Panic1(ecdsa.GenerateKey(elliptic.P224(), rand.Reader)),
			}
		default:
			panic(fmt.Errorf("unknown ECDSA curve type : %s", option.Curve))
		}
	case algED25519:
		pk, puk := fn.Panic2(ed25519.GenerateKey(rand.Reader))
		return &AsymmetricPrivateKey{
			Alg:    algorithm,
			key:    pk,
			pubKey: puk,
		}
	default:
		panic(fmt.Errorf("unknown algorithm type : %d", algorithm))
	}
}
func (Asymmetric) ParsePrivateKey(pem []byte) *AsymmetricPrivateKey {
	p := new(AsymmetricPrivateKey)
	p.Load(pem)
	return p
}
func (Asymmetric) ParsePublicKey(pem []byte) *AsymmetricPublicKey {
	p := new(AsymmetricPublicKey)
	p.Load(pem)
	return p
}

func (Asymmetric) SignRSA(key *AsymmetricPrivateKey, hash crypto.Hash, data []byte, opt *rsa.PSSOptions) []byte {
	if opt == nil {
		return fn.Panic1(rsa.SignPKCS1v15(rand.Reader, key.mustRSA(), hash, data))
	} else {
		return fn.Panic1(rsa.SignPSS(rand.Reader, key.mustRSA(), hash, data, opt))
	}

}
func (Asymmetric) VerifyRSA(key *AsymmetricPublicKey, hash crypto.Hash, data, sign []byte, opt *rsa.PSSOptions) bool {
	if opt == nil {
		err := rsa.VerifyPKCS1v15(key.mustRSA(), hash, data, sign)
		if err != nil {
			if errors.Is(err, rsa.ErrVerification) {
				return false
			} else {
				panic(err)
			}
		} else {
			return true
		}
	} else {
		err := rsa.VerifyPSS(key.mustRSA(), hash, data, sign, opt)
		if err != nil {
			if errors.Is(err, rsa.ErrVerification) {
				return false
			} else {
				panic(err)
			}
		} else {
			return true
		}
	}
}
func (Asymmetric) SignECDSA(key *AsymmetricPrivateKey, data []byte) []byte {
	return fn.Panic1(ecdsa.SignASN1(rand.Reader, key.mustECDSA(), data))
}
func (Asymmetric) VerifyECDSA(key *AsymmetricPublicKey, data, sign []byte) bool {
	return ecdsa.VerifyASN1(key.mustECDSA(), data, sign)
}

func (Asymmetric) SignED25519(key *AsymmetricPrivateKey, data []byte) []byte {
	return ed25519.Sign(key.mustED25519(), data)
}
func (Asymmetric) VerifyED25519(key *AsymmetricPublicKey, data, sign []byte) bool {
	return ed25519.Verify(key.mustED25519(), data, sign)
}
func (a Asymmetric) Sign(key *AsymmetricPrivateKey, data []byte, hasher crypto.Hash, opt *rsa.PSSOptions) []byte {
	switch key.Alg {
	case algRSA:
		return a.SignRSA(key, hasher, data, opt)
	case algECDSA:
		return a.SignECDSA(key, data)
	case algED25519:
		return a.SignED25519(key, data)
	default:
		panic(fmt.Errorf("unsupported algorithm for sign: %d", key.Alg))
	}
}
func (a Asymmetric) Verify(key *AsymmetricPublicKey, data, sign []byte, hasher crypto.Hash, opt *rsa.PSSOptions) bool {
	switch key.Alg {
	case algRSA:
		return a.VerifyRSA(key, hasher, data, sign, opt)
	case algECDSA:
		return a.VerifyECDSA(key, data, sign)
	case algED25519:
		return a.VerifyED25519(key, data, sign)
	default:
		panic(fmt.Errorf("unsupported algorithm for sign: %d", key.Alg))
	}
}
func (a Asymmetric) Encrypt(key *AsymmetricPublicKey, data []byte, hasher hash.Hash) []byte {
	switch key.Alg {
	case algRSA:
		return fn.Panic1(rsa.EncryptOAEP(hasher, rand.Reader, key.mustRSA(), data, nil))
	case algECDSA:
		return fn.Panic1(eciesgo.Encrypt(key.mustECISE(), data))
	default:
		panic(fmt.Errorf("unsupported algorithm for sign: %d", key.Alg))
	}
}
func (a Asymmetric) Decrypt(key *AsymmetricPrivateKey, secret []byte, hasher hash.Hash) []byte {
	switch key.Alg {
	case algRSA:
		return fn.Panic1(rsa.DecryptOAEP(hasher, rand.Reader, key.mustRSA(), secret, nil))
	case algECDSA:
		return fn.Panic1(eciesgo.Decrypt(key.mustECISE(), secret))
	default:
		panic(fmt.Errorf("unsupported algorithm for sign: %d", key.Alg))
	}
}
