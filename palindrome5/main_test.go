package palindrome5

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one symbol",
			args: args{s: "a"},
			want: true,
		},

		{
			name: "full string, odd length 3",
			args: args{s: "aba"},
			want: true,
		},
		{
			name: "full string, odd length 3",
			args: args{s: "abc"},
			want: false,
		},

		{
			name: "full string, odd length 5",
			args: args{s: "abcba"},
			want: true,
		},
		{
			name: "full string, odd length 5",
			args: args{s: "abcxa"},
			want: false,
		},

		{
			name: "full string, even length 2",
			args: args{s: "aa"},
			want: true,
		},
		{
			name: "full string, even length 2",
			args: args{s: "ab"},
			want: false,
		},

		{
			name: "full string, even length 4",
			args: args{s: "abba"},
			want: true,
		},
		{
			name: "full string, even length 4",
			args: args{s: "abbx"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome([]rune(tt.args.s)); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one symbol",
			args: args{s: "abc"},
			want: "a",
		},

		{
			name: "full string, odd length 3",
			args: args{s: "aba"},
			want: "aba",
		},
		{
			name: "right shifted",
			args: args{s: "1 bob"},
			want: "bob",
		},
		{
			name: "right shifted",
			args: args{s: "bob 1"},
			want: "bob",
		},
		{
			name: "right and left shifted",
			args: args{s: "1xbobx2"},
			want: "xbobx",
		},
		{
			name: "2 palindromes",
			args: args{s: "there is bob and a racecar nearby"},
			want: " racecar ",
		},
		{
			name: "leetcode1",
			args: args{s: "cbbd"},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
