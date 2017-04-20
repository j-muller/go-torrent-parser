package gotorrentparser

import (
	"time"

	"github.com/zeebo/bencode"
)

type FileMetadata struct {
	Path   []string `bencode:"path"`
	Length int64    `bencode:"length"`
}

type InfoMetadata struct {
	PieceLength int64  `bencode:"piece length"`
	Pieces      []byte `bencode:"pieces"`

	// single file context
	Name   string `bencode:"name"`
	Length int64  `bencode:"length"`

	Files bencode.RawMessage `bencode:"files"`
	// multi file context
	// Files *FileMetadata `bencode:"files"`
}

type Metadata struct {
	Announce     string   `bencode:"announce"`
	AnnounceList []string `bencode:"announce-list"`
	Comment      string   `bencode:"comment"`
	CreatedBy    string   `bencode:"created by"`
	// Name      string `bencode:"name"`
	CreatedAt int64              `bencode:"creation date"`
	Info      bencode.RawMessage `bencode:"info"`
	// Info      interface{}        `bencode:"info"`
	// InfoHash  string
	// Private   int64   `bencode:"private"`
	// Files     []*File `bencode:"files"`
}

type File struct {
	// Relative path of the file
	Path []string

	// File length
	Length int64
}

type Torrent struct {
	// Announce URL
	Announce string

	// Torrent comment
	Comment string

	// Author
	CreatedBy string

	// Creation time
	CreatedAt time.Time

	// Torrent SHA1
	InfoHash string

	Files []*File
}
