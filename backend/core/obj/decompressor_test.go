package obj_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/briand787b/piqlit/core/obj"
	"github.com/google/go-cmp/cmp"
)

func TestDecompressor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		file          string
		expErrToBeNil bool
	}{
		{"successfully_decompresses_jpg_pic", "a.jpg", true},
		{"successfully_decompresses_gif_video", "a.gif", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cF, err := os.Open(filepath.Join("testdata", t.Name(), "compress.gz"))
			if err != nil {
				t.Fatal(err)
			}

			dc, err := obj.NewDecompressor(cF)
			if err != nil {
				t.Fatal(err)
			}
			defer dc.Close()

			var b bytes.Buffer
			if _, err := io.Copy(&b, dc); err != nil {
				t.Fatal(err)
			}

			expBS, err := ioutil.ReadFile(filepath.Join("testdata", t.Name(), tt.file))
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(expBS, b.Bytes()) {
				t.Fatal("expected decompressed data and actual decompressed data are different")
			}
		})
	}
}
