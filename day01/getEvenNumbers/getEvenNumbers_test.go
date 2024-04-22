package getevennumbers_test

import (
	"SoftwareGoDay1/getevennumbers"
	"testing"
)

func TestGetEvenNumbersCasePositive(t *testing.T) {
	array := []int{1, 0, 19, 17, 16, 8, 13, 24}
	expected := "0 - 8 - 16 - 24"
	actual := getevennumbers.GetEvenNumbers(array)
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestGetEvenNumbersCaseNegative(t *testing.T) {
	array := []int{-10, -15, -2, -6, -60, -43, -17, -24}
	expected := "-60 - -24 - -10 - -6 - -2"
	actual := getevennumbers.GetEvenNumbers(array)
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestGetEvenNumbersCasePositiveAndNegative(t *testing.T) {
	array := []int{1, -10, 19, -17, 16, -8, 13, 24}
	expected := "-10 - -8 - 16 - 24"
	actual := getevennumbers.GetEvenNumbers(array)
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}