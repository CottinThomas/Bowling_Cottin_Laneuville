package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(input []Frame, expectedScore int, expectedError error) error {
	score, err := GetScore(input)

	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}
	return nil
}

// Test if the table contains 10 Frame
func TestTableLen(t *testing.T) {
	input := []Frame{{0, 2}, {5, 1}}
	expectedScore := -1
	expectedError := fmt.Errorf("Wrong table size")
	if err := scoreChecker(input, expectedScore, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expectedScore = 0
	expectedError = nil
	if err := scoreChecker(input, expectedScore, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

//Test if the Tuple are well formed
func TestTupleForm(t *testing.T) {
	input := []Frame{{7, 5}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expectedScore := -1
	expectedError := fmt.Errorf("Wrong Frame construction")
	if err := scoreChecker(input, expectedScore, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{-1, 5}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expectedScore = -1
	expectedError = fmt.Errorf("Wrong Frame construction")
	if err := scoreChecker(input, expectedScore, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// No longer needed : already checked with TestTableLen (contains 10 Frame)
/*
func TestNullScore(t *testing.T){
	input := []Frame{}
	expected := 0
	if err:=scoreChecker(input, expected, nil); err != nil{
		t.Fatalf("%+v\n",err)
	}
}
*/

// Test if the result is correct (here, no spare, no strike)
func TestScore(t *testing.T) {
	input := []Frame{{1, 5}, {2, 7}, {4, 2}, {3, 2}, {7, 1}, {1, 5}, {7, 2}, {8, 1}, {5, 4}, {2, 3}}
	expected := 72
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}

	input = []Frame{{0, 1}, {2, 3}, {4, 5}, {0, 0}, {0, 1}, {2, 3}, {4, 5}, {0, 0}, {1, 1}, {1, 1}}
	expected = 34
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// Test if score is correct with Spares
func TestSpare(t *testing.T) {
	input := []Frame{{1, 5}, {2, 8}, {4, 2}, {3, 2}, {7, 1}, {1, 5}, {7, 2}, {8, 1}, {5, 5}, {2, 3}}
	expected := 80
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// Test if score is correct with Strikes
func TestStrike(t *testing.T) {
	input := []Frame{{1, 5}, {10, 0}, {4, 2}, {3, 2}, {7, 1}, {1, 5}, {7, 2}, {8, 1}, {4, 5}, {2, 3}}
	expected := 79
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// Test if score is correct with a strike as 10th throw
func TestEndsWithStrike(t *testing.T){
	input := []Frame{{1, 5}, {1, 0}, {4, 2}, {3, 2}, {7, 1}, {1, 5}, {7, 2}, {8, 1}, {4, 5}, {10,0}, {8,1}}
	expected := 87
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// Test if score is correct with a spare as 10th throw
func TestEndsWithSpare(t *testing.T){
	input := []Frame{{1, 5}, {1, 0}, {4, 2}, {3, 2}, {7, 1}, {1, 5}, {7, 2}, {8, 1}, {4, 5}, {7,3}, {8,0}}
	expected := 85
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{1, 5}, {1, 0}, {4, 2}, {3, 2}, {7, 1}, {1, 5}, {7, 2}, {8, 1}, {4, 5}, {7,3}, {8,1}}
	expected = -1
	expectedError := fmt.Errorf("Wrong Frame construction")
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}