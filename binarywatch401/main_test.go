package main

import (
	"reflect"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	type args struct {
		turnedOn int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "turnedOn 1",
			args: args{
				turnedOn: 1,
			},
			want: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			name: "turnedOn 7",
			args: args{
				turnedOn: 7,
			},
			want: []string{
				"3:31", "3:47", "3:55", "3:59", "5:31", "5:47", "5:55", "5:59", "6:31", "6:47", "6:55", "6:59",
				"7:15", "7:23", "7:27", "7:29", "7:30", "7:39", "7:43", "7:45", "7:46", "7:51", "7:53", "7:54",
				"7:57", "7:58", "9:31", "9:47", "9:55", "9:59", "10:31", "10:47", "10:55", "10:59", "11:15",
				"11:23", "11:27", "11:29", "11:30", "11:39", "11:43", "11:45", "11:46", "11:51", "11:53", "11:54",
				"11:57", "11:58",
			},
		},
		{
			name: "turnedOn 9",
			args: args{
				turnedOn: 9,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatch(tt.args.turnedOn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBinaryWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
