package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileChunkReader(t *testing.T) {
	t.Run("read and print file content", func(t *testing.T) {
		fileNames := []string{"./data/testdata.txt", "./data/weather_stations.csv"}
		for _, fileName := range fileNames {
			f, err := os.Open(fileName)

			require.NoError(t, err)

			wantDataB, err := io.ReadAll(f)
			require.NoError(t, err)
			fr := NewFileChunkReader(f)
			gotDataB, err := fr.ReadAll()
			require.NoError(t, err)
			assert.Equalf(t, string(wantDataB), string(gotDataB), "file: %s, incorrect data got from FileChunkReader", f.Name())
		}
	})
}
