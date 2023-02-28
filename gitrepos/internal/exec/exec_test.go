package exec_test

import (
	"gitrepos/internal/exec"
	"testing"
)

func TestRun(t *testing.T) {
	execPath := "/tmp"
	type args struct {
		c []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ls one", args{c: []string{"ls", "-alh"}}, false},
		{"ls single", args{c: []string{"ls", "/tmp"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := exec.Run(tt.args.c, execPath); (err != nil) != tt.wantErr {
				t.Errorf("execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		bin string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"exist command", args{bin: "ls"}, true},
		{"not exist command", args{bin: "notlssss"}, false},
		{"exist command with path", args{bin: "/bin/ls"}, true},
		{"not exist command", args{bin: "/bin/notlssss"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exec.Exists(tt.args.bin); got != tt.want {
				t.Errorf("ExistsCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
