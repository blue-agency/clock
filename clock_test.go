package clock_test

import (
	"testing"
	"time"

	"github.com/kyohei-shimada/clock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		loc *time.Location
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"When 'loc' is nil",
			args{loc: nil},
			true,
		},
		{
			"When 'loc' is not nil",
			args{loc: time.UTC},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := clock.New(tt.args.loc)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_clock_Now(t *testing.T) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fail()
		return
	}

	type fields struct {
		location *time.Location
	}
	tests := []struct {
		name             string
		fields           fields
		wantLocationName string
	}{
		{
			"Return a current time with location initialized by New()",
			fields{location: jst},
			"Asia/Tokyo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := clock.New(tt.fields.location)
			if err != nil {
				t.Fail()
				return
			}

			now := c.Now()
			assert.Equal(t, tt.wantLocationName, now.Location().String())
		})
	}
}
