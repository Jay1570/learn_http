package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderParser(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\nFooFoo: BarBar\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	host, _ := headers.Get("Host")
	foo, _ := headers.Get("FooFoo")
	assert.Equal(t, "localhost:42069", host)
	assert.Equal(t, "BarBar", foo)
	assert.Equal(t, 41, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Invalid Token
	headers = NewHeaders()
	data = []byte("H©st: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Valid multiple value header
	headers = NewHeaders()
	data = []byte("Host: localhost:42069\r\nHost: localhost:42068\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	host, _ = headers.Get("Host")
	assert.Equal(t, "localhost:42069,localhost:42068",host)
	assert.Equal(t, 48, n)
	assert.True(t, done)
}
