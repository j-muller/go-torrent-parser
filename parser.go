package gotorrentparser

import (
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/zeebo/bencode"
)

func Parse(reader io.Reader) (*Torrent, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	metadata := &Metadata{}
	err = bencode.DecodeBytes(data, metadata)
	if err != nil {
		return nil, err
	}

	info := &InfoMetadata{}
	err = bencode.DecodeBytes(metadata.Info, info)
	if err != nil {
		return nil, err
	}

	metadataFiles := make([]*FileMetadata, 0)
	err = bencode.DecodeBytes(info.Files, &metadataFiles)
	if err != nil {
		return nil, err
	}

	files := make([]*File, 0)

	if len(metadataFiles) > 0 {
		for _, f := range metadataFiles {
			files = append(files, &File{
				Path:   append([]string{info.Name}, f.Path...),
				Length: f.Length,
			})
		}
	} else {
		files = append(files, &File{
			Path:   []string{info.Name},
			Length: info.Length,
		})
	}

	return &Torrent{
		Announce:  metadata.Announce,
		Comment:   metadata.Comment,
		CreatedBy: metadata.CreatedBy,
		CreatedAt: time.Unix(metadata.CreatedAt, 0),
		InfoHash:  toSHA1(metadata.Info),
		Files:     files,
	}, nil
}

func ParseFromFile(path string) (*Torrent, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Parse(file)
}
