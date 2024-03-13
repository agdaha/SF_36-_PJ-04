package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		configPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "config ok",
			args: args{configPath: "./testdata/config.json"},
			want: &Config{
				URLS:        []string{"https"},
				Period:      1,
				DatabaseURL: "host",
				BindAddr:    ":80",
			},
			wantErr: false,
		},
		{
			name: "config default",
			args: args{configPath: "./testdata/config2.json"},
			want: &Config{
				URLS:        []string{"https"},
				Period:      3,
				DatabaseURL: "host",
				BindAddr:    ":8383",
			},
			wantErr: false,
		},
		{
			name:    "config bad",
			args:    args{configPath: "./testdata/configbad.json"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "config bad2",
			args:    args{configPath: "./testdata/configdab.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
