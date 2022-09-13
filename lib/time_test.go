package lib

import (
	"testing"
	"time"
)

func TestTimeDuration(t *testing.T) {
	type fields struct {
		fn func(time.Duration) string
	}

	type args struct {
		n map[time.Duration]string
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "test 1",
			prepare: func(f *fields) {
				f.fn = TimeDuration
			},
			args: args{
				n: map[time.Duration]string{
					1:                                "1sec ago",
					1000:                             "16min 40sec ago",
					10000:                            "2hr 46min 40sec ago",
					(12*Hour + 12*Minute):            "12hr 12min 0sec ago",
					(2*Day + 2*Hour + 2*Minute):      "2d 2hr 2min 0sec ago",
					(2*Month + 2*Day + 2*Minute + 1): "2mo 2d 21hr 0min 7sec ago",
					(3*Year + 2*Month + 2*Day + 9*Hour + 54*Minute + 1): "3yr 2mo 2d 0hr 18min 25sec ago",
					-1000:       "16min 40sec from now",
					0:           "now",
					10000000000: "long ago",
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
				if got := f.fn(k * time.Second); got != v {
					t.Errorf("%s() = %v, want %v", tt.name, got, v)
				}
			}
		})
	}
}
