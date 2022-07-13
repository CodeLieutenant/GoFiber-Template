package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLogFile(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	path := "./test-logs/log.json"

	defer os.RemoveAll("./test-logs")

	file, err := CreateLogFile(path)

	assert.NoError(err)
	assert.NotNil(file)
	assert.FileExists(path)
}

func TestFileExists(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("FileExists", func(t *testing.T) {
		file, err := os.Create("./test-exists.txt")

		assert.NoError(err)

		defer os.Remove("./test-exists.txt")
		defer file.Close()

		assert.True(FileExists("./test-exists.txt"))
	})

	t.Run("File_Does_Not_Exists", func(t *testing.T) {
		assert.False(FileExists("./test-does-not-exists.txt"))
	})
}

func TestCreateDirectory(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	defer os.RemoveAll("./test-dir")

	path, err := CreateDirectory("./test-dir", 0o744)

	assert.NoError(err)
	abs, _ := filepath.Abs("./test-dir")

	assert.Equal(abs, path)
}

func TestRandomString(t *testing.T) {
	t.Parallel()

	l := 32

	str := RandomString(int32(l))

	if base64.RawURLEncoding.EncodedLen(l) != len(str) {
		t.Fatalf("Expected length: %d Given %d", l, len(str))
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(32)
	}
}

func random(b *testing.B) string {
	buffer := make([]byte, 32)

	n, err := rand.Read(buffer)
	if err != nil {
		b.Errorf("error while generating random buffer %v", err)
		return ""
	}

	if n != 32 {
		b.Errorf("expected length 32, given %d", n)
		return ""
	}

	return base64.RawURLEncoding.EncodeToString(buffer)
}

func BenchmarkCryptoRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random(b)
	}
}

func TestIsSuccess(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	type Data struct {
		Value    int
		Expected bool
	}

	data := [400]Data{}

	for i := 0; i < 400; i++ {
		data[i] = Data{
			Value:    100 + i,
			Expected: 100+i >= 200 && 100+i < 300,
		}
	}

	for _, item := range data {
		assert.Equal(item.Expected, IsSuccess(item.Value))
	}
}

func TestUnsafeBytes(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	bytes := []byte("Hello World")

	unsafeBytes := UnsafeBytes("Hello World")

	assert.EqualValues(bytes, unsafeBytes)
}

func TestUnsafeString(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	bytes := []byte("Hello World")

	str := UnsafeString(bytes)

	assert.EqualValues("Hello World", str)
}

func TestGetenv(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("DefaultValue", func(t *testing.T) {
		value := Getenv("HELLO_ENV")
		assert.Empty(value)

		value = Getenv("HELLO_ENV", "some_default_value")

		assert.NotEmpty(value)
		assert.Equal("some_default_value", value)
	})

	t.Run("WithEnvSet", func(t *testing.T) {
		os.Setenv("HELLO_ENV", "value")

		value := Getenv("HELLO_ENV")
		assert.NotEmpty(value)
		assert.Equal("value", value)

		value = Getenv("HELLO_ENV", "hello_world")
		assert.NotEmpty(value)
		assert.Equal("value", value)
	})
}

func TestIsInt(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("IsInt", func(t *testing.T) {
		assert.True(IsInt("23445555"))
	})

	t.Run("NotAnInt", func(t *testing.T) {
		assert.False(IsInt("fjkhadskjdfhasjd"))
	})

	t.Run("CannotStartWith0", func(t *testing.T) {
		assert.False(IsInt("023355"))
	})
}
