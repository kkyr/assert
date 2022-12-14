package assert_test

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/kkyr/assert"
)

type mockT struct {
	*testing.T

	// if true will print failures to stdout
	print bool
}

func (t mockT) Error(args ...any) {
	if t.print {
		str := args[0].(string)
		// properly format string so example test outputs can be properly validated.
		str = strings.Join(strings.Fields(str), " ")
		fmt.Println(str)
	}
}

func (t mockT) Fatal(args ...any) {
	t.Error(args...)
}

func (_ mockT) Helper() {}

func ExampleAssert_Equal() {
	t := mockT{print: true}
	slice1, slice2 := []string{"go", "gopher"}, []string{"go", "slice"}

	assert.New(t).Equal(slice1, slice2)
	assert.New(t).Require().Field("terms").Equal(slice1, slice2)

	// Output:
	// (-want, +got): []string{ "go", - "gopher", + "slice", }
	// terms: (-want, +got): []string{ "go", - "gopher", + "slice", }
}

func ExampleAssert_Nil() {
	t := mockT{print: true}
	slice := make([]string, 0)

	assert.New(t).Nil(slice)
	assert.New(t).Field("users").Nil(slice)

	// Output:
	// want <nil>, got []
	// users: want <nil>, got []
}

func ExampleAssert_NotNil() {
	t := mockT{print: true}
	var slice []string

	assert.New(t).NotNil(slice)
	assert.New(t).Field("users").NotNil(slice)

	// Output:
	// want <non-nil>, got <nil>
	// users: want <non-nil>, got <nil>
}

func ExampleAssert_ErrorIs() {
	t := mockT{print: true}
	err := fmt.Errorf("boom")
	target := fmt.Errorf("wrapped %w", fmt.Errorf("failure"))

	assert.New(t).ErrorIs(err, target)
	assert.New(t).Field("err").ErrorIs(err, target)

	// Output:
	// no error in err's chain matches target
	// err: no error in err's chain matches target
}

func ExampleAssert_Zero() {
	t := mockT{print: true}
	date, _ := time.Parse("2006-01-02", "2022-02-01")

	assert.New(t).Zero(date)
	assert.New(t).Field("updated_at").Zero(date)

	// Output:
	// want zero value, got 2022-02-01 00:00:00 +0000 UTC
	// updated_at: want zero value, got 2022-02-01 00:00:00 +0000 UTC
}

func ExampleAssert_NotZero() {
	t := mockT{print: true}
	duration := time.Duration(0)

	assert.New(t).NotZero(duration)
	assert.New(t).Field("timeout").NotZero(duration)

	// Output:
	// want non-zero value, got 0s
	// timeout: want non-zero value, got 0s
}

func ExampleAssert_Len() {
	t := mockT{print: true}
	str := "hello world"
	num := 2

	assert.New(t).Len(str, 5)
	assert.New(t).Len(num, 1)

	// Output:
	// want len() = 5, got 11
	// could not apply len() to int
}

func TestAssert_Equal(t *testing.T) {
	number := 5

	for _, tc := range []struct {
		want   any
		got    any
		result bool
	}{
		// equal cases
		{nil, nil, true},
		{true, true, true},
		{number, 5, true},
		{"string", "string", true},
		{[]int{1}, []int{1}, true},
		{[2]string{"go", "gopher"}, [2]string{"go", "gopher"}, true},

		// non-equal cases
		{nil, []int{}, false},
		{nil, ([]string)(nil), false},
		{true, false, false},
		{number, &number, false},
		{"go", "gopher", false},
		{[]int{1}, [1]int{1}, false},
		{[]int{}, [0]int{}, false},
	} {
		t.Run(fmt.Sprintf("want=%v, got=%v", tc.want, tc.got), func(t *testing.T) {
			assert := assert.New(mockT{})

			if eqResult := assert.Equal(tc.want, tc.got); eqResult != tc.result {
				t.Fatalf("Equal(%#v, %#v)=%t, want %t", tc.want, tc.got, eqResult, tc.result)
			} else if neResult := assert.NotEqual(tc.want, tc.got); neResult == eqResult {
				t.Fatalf("Equal()=NotEqual()")
			}
		})
	}
}

func TestAssert_Nil(t *testing.T) {
	for _, tc := range []struct {
		value  any
		result bool
	}{
		// nil values
		{nil, true},
		{(*struct{})(nil), true},
		{([]string)(nil), true},

		// non-nil cases
		{0, false},
		{"", false},
		{[]string{}, false},
	} {
		t.Run(fmt.Sprintf("value=%v", tc.value), func(t *testing.T) {
			assert := assert.New(mockT{})

			if nResult := assert.Nil(tc.value); nResult != tc.result {
				t.Fatalf("Nil(%#v)=%t, want %t", tc.value, nResult, tc.result)
			} else if nnResult := assert.NotNil(tc.value); nnResult == nResult {
				t.Fatalf("Nil()=NotNil()")
			}
		})
	}
}

func TestAssert_ErrorIs(t *testing.T) {
	for _, tc := range []struct {
		err    error
		target error
		result bool
	}{
		// matching errors
		{nil, nil, true},
		{io.EOF, io.EOF, true},
		{fmt.Errorf("wrapped %w", io.EOF), io.EOF, true},

		// non-matching errors
		{nil, io.EOF, false},
		{io.EOF, io.ErrUnexpectedEOF, false},
		{fmt.Errorf("not wrapped %v", io.EOF), io.EOF, false},
	} {
		t.Run(fmt.Sprintf("err=%v, target=%v", tc.err, tc.target), func(t *testing.T) {
			assert := assert.New(mockT{})

			if result := assert.ErrorIs(tc.err, tc.target); result != tc.result {
				t.Fatalf("ErrorIs(%#v, %#v)=%t, want %t", tc.err, tc.target, result, tc.result)
			}
		})
	}
}

func TestAssert_Zero(t *testing.T) {
	for _, tc := range []struct {
		value  any
		result bool
	}{
		// zero values
		{nil, true},
		{0, true},
		{"", true},
		{[]string{}, true},
		{&[]string{}, true},
		{make(chan int), true},
		{(*struct{})(nil), true},
		{time.Time{}, true},

		// non-zero cases
		{1, false},
		{"x", false},
		{[]string{"s"}, false},
	} {
		t.Run(fmt.Sprintf("value=%v", tc.value), func(t *testing.T) {
			assert := assert.New(mockT{})

			if zResult := assert.Zero(tc.value); zResult != tc.result {
				t.Fatalf("Zero(%#v)=%t, want %t", tc.value, zResult, tc.result)
			} else if nzResult := assert.NotZero(tc.value); nzResult == zResult {
				t.Fatalf("Zero()=NotZero()")
			}
		})
	}
}

func TestAssert_Len(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	for _, tc := range []struct {
		value  any
		n      int
		result bool
	}{
		// matching
		{"go", 2, true},
		{[]int{}, 0, true},
		{map[string]string{"x": "y"}, 1, true},
		{&[1]string{}, 1, true},
		{make(chan int), 0, true},
		{ch, 2, true},

		// not matching
		{5, 0, false},
		{nil, 0, false},
		{[]string{"single"}, 2, false},
	} {
		t.Run(fmt.Sprintf("value=%v, n=%d", tc.value, tc.n), func(t *testing.T) {
			assert := assert.New(mockT{})

			if result := assert.Len(tc.value, tc.n); result != tc.result {
				t.Fatalf("Len(%#v,%d)=%t, want %t", tc.value, tc.n, result, tc.result)
			}
		})
	}
}

func TestAssert_Field(t *testing.T) {
	t.Run("returns copy of struct", func(t *testing.T) {
		assert := assert.New(&mockT{})

		cpy := assert.Field("")

		if assert == cpy {
			t.Fatalf("Field() returns struct with same reference")
		}
	})
}

func TestAssert_Require(t *testing.T) {
	t.Run("returns copy of struct", func(t *testing.T) {
		assert := assert.New(&mockT{})

		cpy := assert.Require()

		if assert == cpy {
			t.Fatalf("Require() returns struct with same reference")
		}
	})
}
