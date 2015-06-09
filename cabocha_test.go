package cabocha

import (
	"testing"
)

func TestCabocha(t *testing.T) {
	cabocha := MakeCabocha()
	sentence, err := cabocha.Parse("あなたとJava")
	if err != nil {
		t.Error(err)
	}
	if sentence.Text != "あなたとJava" {
		t.Error("")
	}
	if len(sentence.Chunks) <= 0 {
		t.Fatal("Failed to parse.")
	}
	ftok := sentence.Chunks[0].Tokens[0]
	if ftok.Surface() != "あなた" {
		t.Fatal(sentence)
	}
	if ftok.Base() != "あなた" {
		t.Fatal(ftok.Features)
	}
}
