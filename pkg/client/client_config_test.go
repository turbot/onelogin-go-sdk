package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	notSupporttedRegion = "Test"
)

func TestClientIDValidation(t *testing.T) {
	invalidClientID := ""
	validClientID := "test"
	clientSecret := "test"

	t.Run("invalidClientID", func(t *testing.T) {
		config, err := NewConfig(
			invalidClientID,
			clientSecret,
			USRegion,
			DefaultTimeout,
		)
		assert.NotNil(t, err)
		assert.Equal(t, APIClientConfig{}, config)
	})

	t.Run("validClientID", func(t *testing.T) {
		config, err := NewConfig(
			validClientID,
			clientSecret,
			USRegion,
			DefaultTimeout,
		)
		assert.Nil(t, err)
		assert.Equal(t, config.clientID, validClientID)
	})
}

func TestClientSecretValidation(t *testing.T) {
	invalidClientSecret := ""
	validClientSecret := "test"
	clientID := "test"

	t.Run("invalidClientSecret", func(t *testing.T) {
		config, err := NewConfig(
			clientID,
			invalidClientSecret,
			USRegion,
			DefaultTimeout,
		)
		assert.NotNil(t, err)
		assert.Equal(t, APIClientConfig{}, config)
	})

	t.Run("validClientSecret", func(t *testing.T) {
		config, err := NewConfig(
			clientID,
			validClientSecret,
			USRegion,
			DefaultTimeout,
		)
		assert.Nil(t, err)
		assert.Equal(t, config.clientSecret, validClientSecret)
	})

}

func TestTimeoutBehavior(t *testing.T) {
	clientID := "test"
	clientSecret := "test"
	testTimeout := 2
	timeoutDefaultTrigger := 0

	tests := map[string]struct {
		timeout         int
		expectedTimeout int
	}{
		"has the provided timeout": {
			timeout:         testTimeout,
			expectedTimeout: testTimeout,
		},
		"uses the expected default timeout": {
			timeout:         timeoutDefaultTrigger,
			expectedTimeout: DefaultTimeout,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cc, err := NewConfig(
				clientID,
				clientSecret,
				USRegion,
				test.timeout,
			)
			assert.Nil(t, err)
			assert.Equal(t, test.expectedTimeout, cc.timeout)
		})
	}
}

func TestRegionValidation(t *testing.T) {
	tests := map[string]struct {
		region             string
		expectedValidation bool
	}{
		"us region": {
			region:             USRegion,
			expectedValidation: true,
		},
		"eu region": {
			region:             EURegion,
			expectedValidation: true,
		},
		"not supported": {
			region:             notSupporttedRegion,
			expectedValidation: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			isValid := isSupportedRegion(test.region)
			assert.Equal(t, test.expectedValidation, isValid)
		})
	}
}