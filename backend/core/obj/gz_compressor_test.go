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

// TODO: verify that go's implementation of gzip is different than
// Unix's implementation of gzip, because the byte slice outputs
// are different and that causes this test to fail
//
// func TestCompressor(t *testing.T) {
// 	t.Parallel()
// 	tests := []struct {
// 		name          string
// 		file          string
// 		expErrToBeNil bool
// 	}{
// 		{"successfully_compresses_gif_video", "a.gif", true},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			uf, err := os.Open(filepath.Join("testdata", t.Name(), tt.file))
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			st, err := uf.Stat()
// 			if err != nil {
// 				t.Log(err)
// 			} else {
// 				t.Logf("uncompressed file size: %v", st.Size())
// 			}

// 			c := obj.NewCompressor(uf)
// 			defer c.Close()

// 			var b bytes.Buffer
// 			_, err = io.Copy(&b, c)
// 			if !tt.expErrToBeNil {
// 				if err == nil {
// 					t.Fatal("expected erorr to be non-nil, was nil")
// 				}

// 				return
// 			}

// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			expBS, err := ioutil.ReadFile(filepath.Join("testdata", t.Name(), "compress.gz"))
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			if actBS := b.Bytes(); !cmp.Equal(actBS, expBS) {
// 				t.Logf("expected byte slice len: %v", len(expBS))
// 				t.Logf("actual byte slice len: %v", len(actBS))
// 				t.Fatal("expected compressed data not equal to actual compressed data")
// 			}

// 		})
// 	}
// }

func TestCompressionFull(t *testing.T) {
	tests := []struct {
		name string
		file string
	}{
		{"comps_decomps_gif_video", "a.gif"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			uf, err := os.Open(filepath.Join("testdata", t.Name(), tt.file))
			if err != nil {
				t.Fatal(err)
			}

			st, err := uf.Stat()
			if err != nil {
				t.Log(err)
			} else {
				t.Logf("uncompressed file size:\t%v", st.Size())
			}

			expBS, err := ioutil.ReadAll(uf)
			if err != nil {
				t.Fatal(err)
			}

			if _, err := uf.Seek(0, 0); err != nil {
				t.Fatal(err)
			}

			c := obj.NewCompressor(uf)
			defer c.Close()

			var zb bytes.Buffer
			if _, err = io.Copy(&zb, c); err != nil {
				t.Fatal(err)
			}

			t.Logf("compressed file size:\t%v", len(zb.Bytes()))

			d, err := obj.NewDecompressor(ioutil.NopCloser(&zb))
			if err != nil {
				t.Fatal(err)
			}
			defer d.Close()

			var ub bytes.Buffer
			if _, err := io.Copy(&ub, d); err != nil {
				t.Fatal(err)
			}

			if actBS := ub.Bytes(); !cmp.Equal(expBS, actBS) {
				t.Fatal("expected file contents not the same as actual file contents")
			}
		})
	}
}
