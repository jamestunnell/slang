package archives_test

import (
	"bytes"
	"io"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/archives"
	"github.com/jamestunnell/slang/examples"
)

func TestTarGz_PackUnpack(t *testing.T) {
	meta := slang.PackageMeta{
		Path:    "examples/garage",
		Version: "0.1.0",
		Commit:  "",
	}
	tgz := archives.NewTarGz(meta)
	exampleFS := examples.Garage()

	expectedData, err := makeFileTreeData(exampleFS)

	require.NoError(t, err)

	err = tgz.Pack(exampleFS)

	require.NoError(t, err)

	resultFS, err := tgz.Unpack()

	require.NoError(t, err)

	actualData, err := makeFileTreeData(resultFS)

	require.NoError(t, err)
	require.ElementsMatch(t, maps.Keys(actualData), maps.Keys(expectedData))

	for path, data := range expectedData {
		assert.Equal(t, actualData[path], data)
	}
}

func makeFileTreeData(root fs.FS) (map[string][]byte, error) {
	treeData := map[string][]byte{}

	err := fs.WalkDir(root, ".", func(path string, d fs.DirEntry, errArg error) error {
		if errArg != nil {
			return errArg
		}

		if f, err := root.Open(path); err != nil {
			return err
		} else {
			var buf bytes.Buffer

			_, _ = io.Copy(&buf, f)

			treeData[path] = buf.Bytes()
		}

		return nil
	})

	return treeData, err
}
