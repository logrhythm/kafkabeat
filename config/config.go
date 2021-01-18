// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Brokers                 []string `config:"brokers"`
	Topics                  []string `config:"topics"`
	Group                   string   `config:"group"`
	TLSEnabled              bool     `config:"tls_enabled"`
	TLSCertificateAuthority string   `config:"tls_certificate_authorities"`
	TLSCertificate          string   `config:"tls_certificate"`
	TLSCertificateKey       string   `config:"tls_certificate_key"`
}

var DefaultConfig = Config{
	Brokers:                 []string{"localhost:9092"},
	Topics:                  []string{"beat-topic"},
	Group:                   "kafkabeat",
	TLSEnabled:              true,
	TLSCertificateAuthority: "",
	TLSCertificate:          "",
	TLSCertificateKey:       "",
}
