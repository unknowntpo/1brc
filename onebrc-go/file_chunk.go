package main

import (
	"bytes"
	"errors"
	"io"
	"math"
	"os"
)

type FileChunkReader struct {
	// f is the underlining file that need to be read.
	f      *os.File
	chunks []chunk
}

type chunk struct {
	buf *bytes.Buffer
}

func NewFileChunkReader(f *os.File) *FileChunkReader {
	return &FileChunkReader{
		f:      f,
		chunks: make([]chunk, 0, 100),
	}
}

// NumChunks returns number of chunks in FileChunkReader.
func (fr *FileChunkReader) NumChunks() int {
	return len(fr.chunks)
}

func (fr *FileChunkReader) GetChunk(idx int) (chunk, error) {
	if idx >= len(fr.chunks) || idx < 0 {
		return chunk{}, errors.New("Invalid index")
	}
	return fr.chunks[idx], nil
}

const CHUNK_SIZE = 8192

func (fr *FileChunkReader) ReadAll() ([]byte, error) {
	info, err := fr.f.Stat()
	if err != nil {
		return nil, err
	}
	//fmt.Println(info.Size())
	totalSize := info.Size()
	numOfChunks := int(math.Ceil(float64(totalSize) / float64(CHUNK_SIZE)))
	//fmt.Println("numOfChunks", numOfChunks)
	fr.chunks = make([]chunk, numOfChunks)

	for i := 0; i < numOfChunks; i++ {
		offset := i * CHUNK_SIZE
		// What might go wrong ?
		dataBytes := make([]byte, CHUNK_SIZE)
		n, err := fr.f.ReadAt(dataBytes, int64(offset))
		if err != nil {
			switch err {
			case io.EOF:
				break
			default:
				return nil, err
			}
		}
		if n < CHUNK_SIZE {
			// shrink the data
			dataBytes = dataBytes[:n]
		}
		fr.chunks[i] = chunk{buf: bytes.NewBuffer(dataBytes)}
	}
	var buf bytes.Buffer
	for _, ck := range fr.chunks {
		buf.WriteString(ck.buf.String())
	}
	return buf.Bytes(), nil
}
