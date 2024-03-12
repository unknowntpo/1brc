package main

import (
	"errors"
	"io"
)

type FileChunkReader struct {
	chunks []chunk
}

type chunk struct {
	idx   uint
	start int
	size  int
	data  []byte
}

func NewFileChunkReader() *FileChunkReader {
	return &FileChunkReader{chunks: make([]chunk, 0, 100)}
}

func (fr *FileChunkReader) ReadFile(r io.Reader) error {
	return nil
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
