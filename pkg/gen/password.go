package gen

import (
	"github.com/pkg/errors"
	"github.com/sethvargo/go-diceware/diceware"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Password struct {
	WordCount        int
	Separator        string
	DigitCount       int
	ShouldCapitalize bool

	generator *diceware.Generator
	sb        strings.Builder
}

func NewPassword() (*Password, error) {
	generator, err := diceware.NewGenerator(&diceware.GeneratorInput{
		WordList: diceware.WordListEffLarge(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "gen: creating diceware generator")
	}

	pw := &Password{generator: generator}
	return pw, nil
}

func (p *Password) Generate() (string, error) {
	if p.WordCount < 1 {
		return "", errors.New("gen: Password.WordCount must be greater than 0")
	}
	if p.DigitCount < 0 {
		return "", errors.New("gen: Password.DigitCount must be greater than or equal to 0")
	}

	words, err := p.generator.Generate(p.WordCount)
	if err != nil {
		return "", err
	}

	p.sb.Reset()
	for i, word := range words {
		if p.ShouldCapitalize {
			word = strings.Title(word)
		}
		p.sb.WriteString(word)
		if i < p.WordCount-1 || p.DigitCount > 0 {
			p.sb.WriteString(p.Separator)
		}
	}

	// Doing this numerically does not account for overflow
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < p.DigitCount; i++ {
		p.sb.WriteString(strconv.Itoa(rnd.Intn(10)))
	}

	return p.sb.String(), nil
}
