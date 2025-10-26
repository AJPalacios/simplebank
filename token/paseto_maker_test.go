package token

import (
	"math/rand"
	"testing"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/stretchr/testify/require"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func randomString(n int) string {
	var sb []byte
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb = append(sb, c)
	}

	return string(sb)
}

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(randomString(chacha20poly1305.KeySize))
	require.NoError(t, err)

	username := randomString(10)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(randomString(chacha20poly1305.KeySize))
	require.NoError(t, err)

	token, err := maker.CreateToken(randomString(10), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(randomString(chacha20poly1305.KeySize))
	require.NoError(t, err)

	payload, err := maker.VerifyToken("invalid-token")
	require.Error(t, err)
	require.Nil(t, payload)
}

func TestInvalidPasetoKeySize(t *testing.T) {
	maker, err := NewPasetoMaker(randomString(10))
	require.Error(t, err)
	require.EqualError(t, err, "invalid key size: must be exactly 32 characters")
	require.Nil(t, maker)
}
