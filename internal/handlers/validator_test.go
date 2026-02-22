package handlers

import "testing"

// check validate code
func TestValidateShortCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid short code 1", args{code: "1234567890"}, true},
		{"Valid short code 2", args{code: "abcDE67890"}, true},
		{"Valid short code 3", args{code: "__________"}, true},
		{"Valid short code 4", args{code: "AsDfGhJkLq"}, true},
		{"Invalid short code 1", args{code: ""}, false},
		{"Invalid short code 2", args{code: "aads1_"}, false},
		{"Invalid short code 3", args{code: "dsafHjjkhdks"}, false},
		{"Invalid short code 4", args{code: "lkjsdfljk!"}, false},
		{"Invalid short code 5", args{code: "lkjsdfljk-"}, false},
		{"Invalid short code 1", args{code: "lkjs fljk-"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateShortCode(tt.args.code); got != tt.want {
				t.Errorf("validateShortCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateURL(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid URL 1", args{rawURL: "https://google.com"}, true},
		{"Valid URL 2", args{rawURL: "http://example.com"}, true},
		{"Valid URL 3", args{rawURL: "https://ya.ru/user?id=1"}, true},
		{"Valid URL 4", args{rawURL: "https://sub.example.com/user?id=1"}, true},
		{"Invalid URL 1", args{rawURL: ""}, false},
		{"Invalid URL 2", args{rawURL: "ht!ps://google.com"}, false},
		{"Invalid URL 3", args{rawURL: "google.com"}, false},
		{"Invalid URL 4", args{rawURL: "httpsgoogle.com"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateURL(tt.args.rawURL); got != tt.want {
				t.Errorf("validateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
