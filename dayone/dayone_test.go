package dayone

import (
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestSendFirstWithLastNumberAsInt(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{"single digit", "dayone", 11},
		{"double digits", "oneight", 18},
		{"complex string containing numbers", "I have oneight apples and twone oranges", 11},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			numChan := make(chan int)
			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func() {
				sendFirstWithLastNumberAsInt(tc.input, numChan)
				wg.Done()
			}()
			go func() {
				wg.Wait()
				close(numChan)
			}()

			received := <-numChan
			if received != tc.expected {
				t.Fatalf("Expected %d but got %d", tc.expected, received)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		want        int
	}{
		{
			name:        "Empty file",
			fileContent: "",
			want:        0,
		},
		{
			name:        "Single line no match",
			fileContent: "Hello World",
			want:        0,
		},
		{
			name:        "Single line with match",
			fileContent: "Hello1World2",
			want:        12,
		},
		{
			name:        "Multiple lines with match",
			fileContent: "Hello1\nWorld2",
			want:        33,
		},
		{
			name:        "Multiple lines with match",
			fileContent: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			want:        142,
		},
		{
			name: "Multiple lines with match",
			fileContent: `two1nine
							eightwothree
							abcone2threexyz
							xtwone3four
							4nineeightseven2
							zoneight234
							7pqrstsixteen`,
			want: 281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Run(strings.NewReader(tt.fileContent))

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
