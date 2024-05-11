// Code generated by define_gene; DO NOT EDIT.
package tls

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"crypto/tls"
	_ "github.com/ZenLiuCN/engine/golang/context"
	_ "github.com/ZenLiuCN/engine/golang/crypto"
	_ "github.com/ZenLiuCN/engine/golang/crypto/x509"
	_ "github.com/ZenLiuCN/engine/golang/io"
	_ "github.com/ZenLiuCN/engine/golang/net"
	_ "github.com/ZenLiuCN/engine/golang/time"
)

var (
	//go:embed crypto_tls.d.ts
	CryptoTlsDefine   []byte
	CryptoTlsDeclared = map[string]any{
		"QUICEncryptionLevelEarly":                      tls.QUICEncryptionLevelEarly,
		"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA":              tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
		"cipherSuites":                                  tls.CipherSuites,
		"client":                                        tls.Client,
		"dialWithDialer":                                tls.DialWithDialer,
		"PSSWithSHA384":                                 tls.PSSWithSHA384,
		"QUICEncryptionLevelApplication":                tls.QUICEncryptionLevelApplication,
		"PKCS1WithSHA256":                               tls.PKCS1WithSHA256,
		"QUICSetReadSecret":                             tls.QUICSetReadSecret,
		"ECDSAWithP256AndSHA256":                        tls.ECDSAWithP256AndSHA256,
		"TLS_AES_128_GCM_SHA256":                        tls.TLS_AES_128_GCM_SHA256,
		"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA":          tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		"TLS_RSA_WITH_AES_128_CBC_SHA":                  tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		"x509KeyPair":                                   tls.X509KeyPair,
		"VersionSSL30":                                  tls.VersionSSL30,
		"VersionTLS12":                                  tls.VersionTLS12,
		"VersionTLS13":                                  tls.VersionTLS13,
		"cipherSuiteName":                               tls.CipherSuiteName,
		"Ed25519":                                       tls.Ed25519,
		"listen":                                        tls.Listen,
		"TLS_CHACHA20_POLY1305_SHA256":                  tls.TLS_CHACHA20_POLY1305_SHA256,
		"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305":        tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305":          tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		"TLS_FALLBACK_SCSV":                             tls.TLS_FALLBACK_SCSV,
		"TLS_RSA_WITH_AES_128_CBC_SHA256":               tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
		"dial":                                          tls.Dial,
		"PKCS1WithSHA1":                                 tls.PKCS1WithSHA1,
		"PSSWithSHA512":                                 tls.PSSWithSHA512,
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256":         tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":         tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		"ECDSAWithP384AndSHA384":                        tls.ECDSAWithP384AndSHA384,
		"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256":       tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384":       tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		"ECDSAWithP521AndSHA512":                        tls.ECDSAWithP521AndSHA512,
		"NoClientCert":                                  tls.NoClientCert,
		"RenegotiateFreelyAsClient":                     tls.RenegotiateFreelyAsClient,
		"TLS_RSA_WITH_AES_256_CBC_SHA":                  tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		"VerifyClientCertIfGiven":                       tls.VerifyClientCertIfGiven,
		"TLS_RSA_WITH_AES_256_GCM_SHA384":               tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		"X25519":                                        tls.X25519,
		"CurveP521":                                     tls.CurveP521,
		"PKCS1WithSHA512":                               tls.PKCS1WithSHA512,
		"quicClient":                                    tls.QUICClient,
		"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256":   tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
		"TLS_ECDHE_RSA_WITH_RC4_128_SHA":                tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
		"TLS_RSA_WITH_3DES_EDE_CBC_SHA":                 tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		"newListener":                                   tls.NewListener,
		"quicServer":                                    tls.QUICServer,
		"RequestClientCert":                             tls.RequestClientCert,
		"RequireAnyClientCert":                          tls.RequireAnyClientCert,
		"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA":           tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
		"CurveP256":                                     tls.CurveP256,
		"ECDSAWithSHA1":                                 tls.ECDSAWithSHA1,
		"QUICHandshakeDone":                             tls.QUICHandshakeDone,
		"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256":       tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		"TLS_RSA_WITH_RC4_128_SHA":                      tls.TLS_RSA_WITH_RC4_128_SHA,
		"loadX509KeyPair":                               tls.LoadX509KeyPair,
		"parseSessionState":                             tls.ParseSessionState,
		"QUICNoEvent":                                   tls.QUICNoEvent,
		"QUICWriteData":                                 tls.QUICWriteData,
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":            tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		"QUICEncryptionLevelHandshake":                  tls.QUICEncryptionLevelHandshake,
		"server":                                        tls.Server,
		"VersionTLS11":                                  tls.VersionTLS11,
		"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":         tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		"TLS_RSA_WITH_AES_128_GCM_SHA256":               tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		"newLRUClientSessionCache":                      tls.NewLRUClientSessionCache,
		"QUICEncryptionLevelInitial":                    tls.QUICEncryptionLevelInitial,
		"QUICRejectedEarlyData":                         tls.QUICRejectedEarlyData,
		"RenegotiateOnceAsClient":                       tls.RenegotiateOnceAsClient,
		"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA":          tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		"insecureCipherSuites":                          tls.InsecureCipherSuites,
		"PKCS1WithSHA384":                               tls.PKCS1WithSHA384,
		"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256": tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
		"versionName":                                   tls.VersionName,
		"VersionTLS10":                                  tls.VersionTLS10,
		"CurveP384":                                     tls.CurveP384,
		"PSSWithSHA256":                                 tls.PSSWithSHA256,
		"QUICSetWriteSecret":                            tls.QUICSetWriteSecret,
		"RequireAndVerifyClientCert":                    tls.RequireAndVerifyClientCert,
		"TLS_AES_256_GCM_SHA384":                        tls.TLS_AES_256_GCM_SHA384,
		"newResumptionState":                            tls.NewResumptionState,
		"QUICTransportParameters":                       tls.QUICTransportParameters,
		"QUICTransportParametersRequired":               tls.QUICTransportParametersRequired,
		"RenegotiateNever":                              tls.RenegotiateNever,
		"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,

		"emptyCertificate": func() (v tls.Certificate) {
			return v
		},
		"refCertificate": func() *tls.Certificate {
			var x tls.Certificate
			return &x
		},
		"refOfCertificate": func(x tls.Certificate) *tls.Certificate {
			return &x
		},
		"emptyCertificateRequestInfo": func() (v tls.CertificateRequestInfo) {
			return v
		},
		"refCertificateRequestInfo": func() *tls.CertificateRequestInfo {
			var x tls.CertificateRequestInfo
			return &x
		},
		"refOfCertificateRequestInfo": func(x tls.CertificateRequestInfo) *tls.CertificateRequestInfo {
			return &x
		},
		"emptySessionState": func() (v tls.SessionState) {
			return v
		},
		"refSessionState": func() *tls.SessionState {
			var x tls.SessionState
			return &x
		},
		"refOfSessionState": func(x tls.SessionState) *tls.SessionState {
			return &x
		},
		"emptyConnectionState": func() (v tls.ConnectionState) {
			return v
		},
		"refConnectionState": func() *tls.ConnectionState {
			var x tls.ConnectionState
			return &x
		},
		"refOfConnectionState": func(x tls.ConnectionState) *tls.ConnectionState {
			return &x
		},
		"emptyQUICConfig": func() (v tls.QUICConfig) {
			return v
		},
		"refQUICConfig": func() *tls.QUICConfig {
			var x tls.QUICConfig
			return &x
		},
		"refOfQUICConfig": func(x tls.QUICConfig) *tls.QUICConfig {
			return &x
		},
		"emptyCipherSuite": func() (v tls.CipherSuite) {
			return v
		},
		"refCipherSuite": func() *tls.CipherSuite {
			var x tls.CipherSuite
			return &x
		},
		"refOfCipherSuite": func(x tls.CipherSuite) *tls.CipherSuite {
			return &x
		},
		"emptyClientHelloInfo": func() (v tls.ClientHelloInfo) {
			return v
		},
		"refClientHelloInfo": func() *tls.ClientHelloInfo {
			var x tls.ClientHelloInfo
			return &x
		},
		"refOfClientHelloInfo": func(x tls.ClientHelloInfo) *tls.ClientHelloInfo {
			return &x
		},
		"emptyConfig": func() (v tls.Config) {
			return v
		},
		"refConfig": func() *tls.Config {
			var x tls.Config
			return &x
		},
		"refOfConfig": func(x tls.Config) *tls.Config {
			return &x
		},
		"emptyConn": func() (v tls.Conn) {
			return v
		},
		"refConn": func() *tls.Conn {
			var x tls.Conn
			return &x
		},
		"refOfConn": func(x tls.Conn) *tls.Conn {
			return &x
		},
		"emptyDialer": func() (v tls.Dialer) {
			return v
		},
		"refDialer": func() *tls.Dialer {
			var x tls.Dialer
			return &x
		},
		"refOfDialer": func(x tls.Dialer) *tls.Dialer {
			return &x
		},
		"emptyQUICSessionTicketOptions": func() (v tls.QUICSessionTicketOptions) {
			return v
		},
		"refQUICSessionTicketOptions": func() *tls.QUICSessionTicketOptions {
			var x tls.QUICSessionTicketOptions
			return &x
		},
		"refOfQUICSessionTicketOptions": func(x tls.QUICSessionTicketOptions) *tls.QUICSessionTicketOptions {
			return &x
		},
		"emptyClientSessionState": func() (v tls.ClientSessionState) {
			return v
		},
		"refClientSessionState": func() *tls.ClientSessionState {
			var x tls.ClientSessionState
			return &x
		},
		"refOfClientSessionState": func(x tls.ClientSessionState) *tls.ClientSessionState {
			return &x
		},
		"emptyQUICConn": func() (v tls.QUICConn) {
			return v
		},
		"refQUICConn": func() *tls.QUICConn {
			var x tls.QUICConn
			return &x
		},
		"refOfQUICConn": func(x tls.QUICConn) *tls.QUICConn {
			return &x
		},
		"emptyQUICEvent": func() (v tls.QUICEvent) {
			return v
		},
		"refQUICEvent": func() *tls.QUICEvent {
			var x tls.QUICEvent
			return &x
		},
		"refOfQUICEvent": func(x tls.QUICEvent) *tls.QUICEvent {
			return &x
		}}
)

func init() {
	engine.RegisterModule(CryptoTlsModule{})
}

type CryptoTlsModule struct{}

func (S CryptoTlsModule) Identity() string {
	return "golang/crypto/tls"
}
func (S CryptoTlsModule) TypeDefine() []byte {
	return CryptoTlsDefine
}
func (S CryptoTlsModule) Exports() map[string]any {
	return CryptoTlsDeclared
}
