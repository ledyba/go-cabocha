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
	if len(sentence.Chunks) <= 0 {
		t.Fatal("Failed to parse.")
	}
	if sentence.Chunks[0].Tokens[0].Surface() != "あなた" {
		t.Fatal(sentence)
	}
}
