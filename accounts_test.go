package main

import (
	"testing"
)

// Not sure how to test addAccountHandler - need to seperate from firestore

func TestCheckIfNameIsValidValid(t *testing.T) {
	if !CheckIfNameIsValid("HciWare") {
		t.Fatalf("Valid name return false")
	}
}

func TestCheckIfNameIsValidInValid(t *testing.T) {
	if CheckIfNameIsValid(" HciWare") {
		t.Fatalf("Valid name return false")
	}
}

func TestCheckIfNameIsValidInEmpty(t *testing.T) {
	if CheckIfNameIsValid("") {
		t.Fatalf("Empty name should be invalid")
	}
}

func TestCheckIfNameIsValidInCracked(t *testing.T) {
	if CheckIfNameIsValid("_87874 7*&8 ````__,mm") {
		t.Fatalf("Empty name should be invalid")
	}
}

func TestCheckIfNameIsValidInLong(t *testing.T) {
	if !CheckIfNameIsValid("abcdefghijklmnopABCEDFKJHOYGS0239402943___9034802943") {
		t.Fatalf("Long name should be valid")
	}
}
