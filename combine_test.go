package gopeg_mv4

import (
	"reflect"
	"testing"
)

func TestCombineMp4(t *testing.T) {
	err := CombineMp4("main.mp4", "./test_input", "./test_output")
	if err != nil {
		t.Errorf("An error occurred when joining Mp4.")
	}
}

func TestWriteTxt(t *testing.T) {
	err := WriteTxt([]string{"test.txt","test1.txt","test2.txt"})
	if err != nil {
		t.Errorf("Text file creation error")
	}
}

func TestRemoveTxt(t *testing.T) {
	err := RemoveTxt()
	if err != nil {
		t.Errorf("The file was not found or there was an error deleting it.")
	}
}

func TestWalkDir(t *testing.T) {
	got := dirwalk("./test_input")
	want := []string{"test_input/test.mp4", "test_input/test1.mp4", "test_input/test2.mp4"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}