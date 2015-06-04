package cabocha

import (
	"encoding/xml"
	"io"
	"os/exec"
	"strings"
)

type Sentence struct {
	XMLName xml.Name `xml:"sentence"`
	Chunks  []Chunk  `xml:"chunk"`
}

type Chunk struct {
	XMLName xml.Name `xml:"chunk"`
	Tokens  []Token  `xml:"tok"`
	Id      int      `xml:"id,attr"`
	Link    int      `xml:"link,attr"`
	Rel     string   `xml:"rel,attr"`
	Score   float32  `xml:"score,attr"`
	Head    int      `xml:"head,attr"`
	Func    int      `xml:"func,attr"`
}

type Token struct {
	XMLName  xml.Name `xml:"tok"`
	Id       int      `xml:"id,attr"`
	Features string   `xml:"feature,attr"`
	Body     string   `xml:",chardata"`
}

func (self *Sentence) Chunk(id int) *Chunk {
	for i := range self.Chunks {
		if self.Chunks[i].Id == id {
			return &self.Chunks[i]
		}
	}
	return nil
}
func (self *Sentence) Token(id int) *Token {
	for i := range self.Chunks {
		for j := range self.Chunks[i].Tokens {
			if self.Chunks[i].Tokens[j].Id == id {
				return &self.Chunks[i].Tokens[j]
			}
		}
	}
	return nil
}
func (self *Token) Contains(feature string) bool {
	return strings.Contains(self.Features, feature)
}

func (self *Chunk) Body() string {
	strs := make([]string, 0)
	for _, tok := range self.Tokens {
		strs = append(strs, tok.Body)
	}
	return strings.Join(strs, " ")
}

type Cabocha struct {
	path string
}

func MakeCabocha() *Cabocha {
	self := &Cabocha{}
	self.path = "cabocha"
	return self
}
func MakeCabochaWithPath(path string) *Cabocha {
	self := &Cabocha{}
	self.path = path
	return self
}

func (self *Cabocha) Parse(sentence string) (*Sentence, error) {
	var err error
	cmd := exec.Command(self.path, "-f", "3")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	io.WriteString(stdin, sentence)
	err = stdin.Close()
	if err != nil {
		return nil, err
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	res := &Sentence{}
	err = xml.Unmarshal(out, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
