package executor

import "regexp"

const tokenregex string = `\$[A-Za-z0-9_]+`

type Tokenizer interface {
	Tokenize(content string) []string
}

type tokenizer struct {
}

func NewTokenizer() tokenizer {
	tokenizer := tokenizer{}
	return tokenizer
}

func (tokenizer tokenizer) Tokenize(content string) []string {
	reg := regexp.MustCompile(tokenregex)
	tokens := reg.FindAllStringSubmatch(content, -1)

	var mergedTokens []string
	for _, token := range tokens {
		mergedTokens = append(mergedTokens, token...)
	}
	return mergedTokens
}
