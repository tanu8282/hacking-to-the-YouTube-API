package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
func Test_Sample(t *testing.T) {
	if !(15 == 15) {
		t.Error(`miss`)
	}
}
