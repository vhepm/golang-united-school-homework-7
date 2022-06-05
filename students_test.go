package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

func NewTestPeople() People {
	return People{
		Person{
			firstName: "Test1",
			lastName:  "test1",
			birthDay:  time.Now(),
		},
		Person{
			firstName: "Test2",
			lastName:  "test2",
			birthDay:  time.Now(),
		},
	}
}

// WRITE YOUR CODE BELOW

func TestPeople_Len(t *testing.T) {
	people := NewTestPeople()

	got := people.Len()
	expected := 2

	if got != expected {
		t.Errorf(
			"Did not get expected result. Wanted %v, got: %v\n",
			expected,
			got,
		)
	}
}

func TestPeople_Unix_Less(t *testing.T) {
	people := NewTestPeople()

	got := people.Less(0, 1)

	if got != true {
		t.Errorf(
			"Did not get expected result. Wanted %v, got: %v\n",
			true,
			got,
		)
	}
}

func TestPeople_SameFirstName_Less(t *testing.T) {
	people := People{
		Person{
			firstName: "Test1",
			lastName:  "test1",
			birthDay:  time.Now(),
		},
		Person{
			firstName: "Test1",
			lastName:  "test21",
			birthDay:  time.Now(),
		},
	}

	got := people.Less(0, 1)

	if got != true {
		t.Errorf(
			"Did not get expected result. Wanted %v, got: %v\n",
			true,
			got,
		)
	}
}

func TestPeople_DifferentUnix_Less(t *testing.T) {
	people := People{
		Person{
			firstName: "Test1",
			lastName:  "test1",
			birthDay:  time.Unix(0, 566).UTC(),
		},
		Person{
			firstName: "Test1",
			lastName:  "test21",
			birthDay:  time.Unix(456, -67).UTC(),
		},
	}

	got := people.Less(0, 1)

	if got != false {
		t.Errorf(
			"Did not get expected result. Wanted %v, got: %v\n",
			false,
			got,
		)
	}
}

func TestPeople_Swap(t *testing.T) {
	people := NewTestPeople()

	people.Swap(0, 1)

	if people[0].firstName != "Test2" {
		t.Error("Did not get expected result.")
	}

}

func TestNew(t *testing.T) {
	input := "1 2 3 4 5\n2 3 5 9 1"

	_, err := New(input)

	if err != nil {
		t.Errorf(
			"Did not get expected result. Got: %v\n",
			err,
		)
	}
}

func TestNew_RowsNotSameLen(t *testing.T) {
	input := "1 2 3 4 5\n2 3 5 91"

	_, err := New(input)

	if err == nil {
		t.Error("Did not get expected result.")
	}
}

func TestNew_ConvError(t *testing.T) {
	input := "1 2 3 4 5\n2 3 5 test 1"

	_, err := New(input)

	if err == nil {
		t.Error("Did not get expected result.")
	}
}

func TestMatrix_Rows(t *testing.T) {
	matrix, _ := New("1 2 3 4 5\n2 3 5 9 1")

	matrix.Rows()
}

func TestMatrix_Cols(t *testing.T) {
	matrix, _ := New("1 2 3 4 5\n2 3 5 9 1")

	matrix.Cols()
}

func TestMatrix_Set_True(t *testing.T) {
	matrix, _ := New("1 2 3 4 5\n2 3 5 9 1")

	got := matrix.Set(1, 1, 5)

	if got != true {
		t.Errorf(
			"Did not get expected result. Wanted %v, got: %v\n",
			true,
			got,
		)
	}
}

func TestMatrix_Set_False(t *testing.T) {
	matrix, _ := New("1 2 3 4 5\n2 3 5 9 1")

	got := matrix.Set(111, 1, 5)

	if got != false {
		t.Errorf(
			"Did not get expected result. Wanted %v, got: %v\n",
			true,
			got,
		)
	}
}
