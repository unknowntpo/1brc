package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileChunkReader(t *testing.T) {
	fileNames := []string{"./data/testdata.txt", "./data/weather_stations.csv"}
	for _, fileName := range fileNames {
		t.Run(fmt.Sprintf("read and print file content: %v", fileName), func(t *testing.T) {
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
}

func BenchmarkFileChunkReader(b *testing.B) {
	fileNames := []string{"./data/weather_stations.csv"}
	for _, fileName := range fileNames {
		b.Run(fmt.Sprintf("read and print file content: %v", fileName), func(b *testing.B) {
			for _, fileName := range fileNames {
				f, err := os.Open(fileName)
				require.NoError(b, err)
				fr := NewFileChunkReader(f)
				_, err = fr.ReadAll()
				require.NoError(b, err)
			}
		})
	}
}
