package assert_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/kkyr/assert"
)

type mockT struct {
	*testing.T

	// print failures to stdout
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

func (_ mockT) Helper() {}

func ExampleAssert_Equal() {
	t := mockT{print: true}
	slice1, slice2 := []string{"go", "gopher"}, []string{"go", "slice"}

	assert.New(t).Equal(slice1, slice2)
	assert.New(t).Field("terms").Equal(slice1, slice2)

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
		t.Run(fmt.Sprintf("want=%#v, got=%#v", tc.want, tc.got), func(t *testing.T) {
			assert := assert.New(mockT{})

			if eqResult := assert.Equal(tc.want, tc.got); eqResult != tc.result {
				t.Fatalf("Equal(%#v, %#v)=%t, want %t", tc.want, tc.got, eqResult, tc.result)
			} else if neResult := assert.NotEqual(tc.want, tc.got); neResult == eqResult {
				t.Fatalf("Equal()=NotEqual()")
			}
		})
	}
}
