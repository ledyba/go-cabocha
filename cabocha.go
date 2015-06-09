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
	ID      int      `xml:"id,attr"`
	Link    int      `xml:"link,attr"`
	Rel     string   `xml:"rel,attr"`
	Score   float32  `xml:"score,attr"`
	Head    int      `xml:"head,attr"`
	Func    int      `xml:"func,attr"`
}

type Token struct {
	XMLName  xml.Name `xml:"tok"`
	ID       int      `xml:"id,attr"`
	feature  string   `xml:"feature,attr"`
	Features []string `xml:"-"`
	body     string   `xml:",chardata"`
}

func (self *Sentence) Chunk(id int) *Chunk {
	for i := range self.Chunks {
		if self.Chunks[i].ID == id {
			return &self.Chunks[i]
		}
	}
	return nil
}
func (self *Sentence) Token(id int) *Token {
	for i := range self.Chunks {
		for j := range self.Chunks[i].Tokens {
			if self.Chunks[i].Tokens[j].ID == id {
				return &self.Chunks[i].Tokens[j]
			}
		}
	}
	return nil
}

func (tok *Token) Contains(feature string) bool {
	return strings.Contains(tok.feature, feature)
}

func (tok *Token) Base() string {
	return tok.Features[6]
}

func (tok *Token) Reading() string {
	return tok.Features[7]
}

func (tok *Token) Pron() string {
	return tok.Features[8]
}
func (tok *Token) Surface() string {
	return tok.body
}

func (chunk *Chunk) Body() string {
	var strs []string
	for _, tok := range chunk.Tokens {
		strs = append(strs, tok.Surface())
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

func (cabo *Cabocha) Parse(sentence string) (*Sentence, error) {
	var err error
	cmd := exec.Command(cabo.path, "-f", "3")
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
	for _, chunk := range res.Chunks {
		for _, tok := range chunk.Tokens {
			tok.Features = strings.Split(tok.feature, ",")
		}
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
