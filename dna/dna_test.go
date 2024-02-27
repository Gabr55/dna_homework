package dna

import "testing"

func TestCompress(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name           string
		args           args
		wantCompressed string
		wantErr        bool
	}{
		{
			name:           "aaaabbсaa -> a4b2с1a2",
			args:           args{source: "aaaabbсaa"},
			wantCompressed: "a4b2с1a2",
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCompressed, err := Compress(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCompressed != tt.wantCompressed {
				t.Errorf("Compress() gotCompressed = %v, want %v", gotCompressed, tt.wantCompressed)
			}
		})
	}
}

func TestDecompress(t *testing.T) {
	type args struct {
		compressed string
	}
	tests := []struct {
		name       string
		args       args
		wantSource string
		wantErr    bool
	}{
		{
			name:       "a4b2с1a2 -> aaaabbсaa",
			args:       args{compressed: "a4b2с1a2"},
			wantSource: "aaaabbсaa",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSource, err := Decompress(tt.args.compressed)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSource != tt.wantSource {
				t.Errorf("Decompress() gotSource = %v, want %v", gotSource, tt.wantSource)
			}
		})
	}
}
