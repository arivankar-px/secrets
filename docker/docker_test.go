package docker

import (
	"io/ioutil"
	"testing"

	"github.com/libopenstorage/secrets"
	"github.com/libopenstorage/secrets/test"
)

func TestAll(t *testing.T) {
	// Set the relevant environmnet fields for vault.
	vs, err := NewDockerSecretTest(nil)
	if err != nil {
		t.Fatalf("Unable to create a Vault Secret instance: %v", err)
		return
	}
	test.Run(vs, t)
}

type dockerSecretTest struct {
	s secrets.Secrets
}

func NewDockerSecretTest(secretConfig map[string]interface{}) (test.SecretTest, error) {
	s, err := New(secretConfig)
	if err != nil {
		return nil, err
	}
	return &dockerSecretTest{s}, nil
}

func (v *dockerSecretTest) TestPutSecret(t *testing.T) error {
	return nil
}

func (v *dockerSecretTest) TestGetSecret(t *testing.T) error {
	secretId := "openstorage_secret"
	cipherBlob := []byte{10, 12, 13}
	err := ioutil.WriteFile(getSecretKey(secretId), cipherBlob, 0644)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	secretData, err := v.s.GetSecret(secretId, nil)
	if err != nil {
		t.Errorf("Unexpected error in GetSecret: %v", err)
	}
	if len(secretData) == 0 {
		t.Errorf("GetSecret returned invalid data.")
	}
	secretBlob, ok := secretData[secretId]
	if !ok {
		t.Errorf("GetSecret returned invalid data")
	}

	if len(secretBlob.([]byte)) != len(cipherBlob) {
		t.Errorf("Cipher texts do not match")
	}

	return nil
}
