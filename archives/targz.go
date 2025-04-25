package archives

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path"
	"time"

	"github.com/jamestunnell/slang"
	"github.com/psanford/memfs"
	"github.com/rs/zerolog/log"
)

type TarGz struct {
	Meta   slang.PackageMeta
	Data   []byte
	Digest string
}

const (
	FormatTarGz  = "targz"
	DigestSHA256 = "sha256"
)

var errDigestMismatch = errors.New("digests do not match")

func NewTarGz(meta slang.PackageMeta) *TarGz {
	return &TarGz{
		Meta:   meta,
		Data:   []byte{},
		Digest: "",
	}
}

func (archive *TarGz) GetMeta() slang.PackageMeta {
	return archive.Meta
}

func (archive *TarGz) GetDataFormat() string {
	return FormatTarGz
}

func (archive *TarGz) GetDigestType() string {
	return DigestSHA256
}

func (archive *TarGz) Pack(root fs.FS) error {
	var tarBuf bytes.Buffer

	tw := tar.NewWriter(&tarBuf)

	if err := tw.AddFS(root); err != nil {
		return fmt.Errorf("failed to add package to archive: %w", err)
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("failed to close tar writer: %w", err)
	}

	metaData, err := json.Marshal(archive.Meta)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata JSON: %w", err)
	}

	var tgzBuf bytes.Buffer

	zw := gzip.NewWriter(&tgzBuf)

	// Setting the Header fields is optional.
	zw.Name = archive.Meta.Path
	zw.Extra = metaData
	zw.ModTime = time.Now()

	// Copy our data to the gzip writer, which compresses it
	if _, err = io.Copy(zw, &tarBuf); err != nil {
		return fmt.Errorf("failed to write gzip data: %w", err)
	}

	if err = zw.Close(); err != nil {
		return fmt.Errorf("failed to close gzip writer: %w", err)
	}

	archive.Data = tgzBuf.Bytes()
	archive.Digest = makeSHA256Digest(archive.Data)

	return nil
}

func makeSHA256Digest(data []byte) string {
	hasher := sha256.New()

	hasher.Write(data)

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (archive *TarGz) Unpack() (fs.FS, error) {
	digest := makeSHA256Digest(archive.Data)
	if digest != archive.Digest {
		return nil, errDigestMismatch
	}

	tgzReader := bytes.NewReader(archive.Data)

	zr, err := gzip.NewReader(tgzReader)
	if err != nil {
		return nil, fmt.Errorf("failed to make gzip reader: %w", err)
	}

	defer func() {
		if closeErr := zr.Close(); closeErr != nil {
			log.Warn().Err(err).Msg("failed to close gzip reader")
		}
	}()

	var meta slang.PackageMeta

	if err = json.Unmarshal(zr.Extra, &meta); err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata JSON from extra: %w", err)
	}

	if meta.Path != archive.Meta.Path {
		return nil, fmt.Errorf("name %s does not match expected %s", meta.Path, archive.Meta.Path)
	}

	if meta.Version != archive.Meta.Version {
		return nil, fmt.Errorf("version %s does not match expected %s", meta.Version, archive.Meta.Version)
	}

	var tarData []byte

	if tarData, err = io.ReadAll(zr); err != nil {
		return nil, fmt.Errorf("failed to read gzip data: %w", err)
	}

	root := memfs.New()
	tr := tar.NewReader(bytes.NewReader(tarData))

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}

		if err != nil {
			return nil, fmt.Errorf("failed to read next archive file: %w", err)
		}

		fpath := hdr.Name

		if err = root.MkdirAll(path.Dir(fpath), 0755); err != nil {
			return nil, fmt.Errorf("failed to create archive file dirs: %w", err)
		}

		var buf bytes.Buffer

		if _, err := io.Copy(&buf, tr); err != nil {
			return nil, fmt.Errorf("failed to get read archive file: %w", err)
		}

		root.WriteFile(fpath, buf.Bytes(), 0755)
	}

	return root, nil
}
