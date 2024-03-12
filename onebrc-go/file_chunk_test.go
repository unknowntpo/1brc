package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileChunkReader(t *testing.T) {
	t.Run("read and print file content", func(t *testing.T) {
		f, err := os.Open("./data/testdata.txt")
		require.NoError(t, err)
		fr := NewFileChunkReader()
		require.NoError(t, fr.ReadFile(f))
	})
}
