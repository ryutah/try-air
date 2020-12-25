package main

import "testing"

func TestFoo(t *testing.T) {
	type in struct {
		a, b int
	}
	tests := []struct {
		name     string
		in       in
		expected int
	}{
		{
			name: "test1",
			in: in{
				a: 1,
				b: 2,
			},
			expected: 3,
		},
		{
			name: "test2",
			in: in{
				a: 2,
				b: 3,
			},
			expected: 5,
		},
		{
			name: "test3",
			in: in{
				a: 3,
				b: 4,
			},
			expected: 7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.in.a + test.in.b
			if test.expected != got {
				t.Errorf("want: %v, got: %v", test.expected, got)
			}
		})
	}
}
