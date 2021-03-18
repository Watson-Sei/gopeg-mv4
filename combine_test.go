package gopeg_mv4

import (
	"reflect"
	"testing"
)

func TestWalkDir(t *testing.T) {
	got := dirwalk("./test_input")
	want := []string{"test_input/test.mp4", "test_input/test1.mp4", "test_input/test2.mp4"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}