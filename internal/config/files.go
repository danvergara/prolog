package config

import (
	"os"
	"path/filepath"
)

var (
	// CAFile defines the path to the ca file.
	CAFile = configFile("ca.pem")
	// ServerCertFile defines the path to the server cert file.
	ServerCertFile = configFile("server.pem")
	// ServerKeyFile defines the path to the server key file.
	ServerKeyFile = configFile("server-key.pem")
	// ClientCertFile defines the path to the client cert file.
	ClientCertFile = configFile("client.pem")
	// ClientKeyFile defines the path to the client key file.
	ClientKeyFile = configFile("client-key.pem")
	// RootClientCertFile defines the path to the root client cert file.
	RootClientCertFile = configFile("root-client.pem")
	// RootClientKeyFile defines the path to the root client key file.
	RootClientKeyFile = configFile("root-client-key.pem")
	// NobodyClientCertFile defines the path to the client cert file.
	NobodyClientCertFile = configFile("nobody-client.pem")
	// NobodyClientKeyFile defines the path to the client key file.
	NobodyClientKeyFile = configFile("nobody-client-key.pem")
	// ACLModelFile defines the path to the ACL model file.
	ACLModelFile = configFile("model.conf")
	// ACLPolicyFile defines the path to the ACL policy file.
	ACLPolicyFile = configFile("policy.csv")
)

func configFile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(homeDir, ".prolog", filename)
}
