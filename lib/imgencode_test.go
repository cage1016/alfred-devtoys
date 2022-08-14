package lib

import (
	"testing"

	"github.com/cage1016/alfred-devtoys/testdata"
	"github.com/stretchr/testify/assert"
)

func TestImageEncode(t *testing.T) {
	type fields struct {
		fn func(image string) (string, string, error)
	}

	type args struct {
		n map[string]string
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "test1",
			prepare: func(f *fields) {
				f.fn = ImageEncode
			},
			args: args{
				n: map[string]string{
					"1.png":     "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGL6P40BEAAA//8ENQGYWUG67AAAAABJRU5ErkJggg==",
					"demo.jpeg": `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`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				if got, _, _ := f.fn(testdata.Path(k)); got != v {
					t.Errorf("ImageEncode() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestDownloadImage(t *testing.T) {
	type fields struct {
		fn func(url, dataDir string) (string, error)
	}

	type args struct {
		n map[string]error
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "test1",
			prepare: func(f *fields) {
				f.fn = Download
			},
			args: args{
				n: map[string]error{
					"https://globalgrasshopper.com/wp-content/uploads/2015/04/Queenstown-views.jpg": nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				_, got := f.fn(k, testdata.Path("."))
				assert.Equal(t, v, got)
			}
		})
	}
}
