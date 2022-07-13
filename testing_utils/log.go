package testing_utils

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/rzajac/zltest"
)

func NewTestLogger(t *testing.T, level zerolog.Level) (zerolog.Logger, *zltest.Tester) {
	tst := zltest.New(t)
	testWriter := zerolog.NewTestWriter(t)

	logger := zerolog.New(zerolog.MultiLevelWriter(tst, testWriter)).Level(level)

	return logger, tst
}
