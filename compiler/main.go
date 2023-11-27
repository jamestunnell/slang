package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
)

func main() {
	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal().Err(err).Str("file", file).Msg("failed to open file")
		}

		l := lexer.New(bufio.NewReader(f))
		toks := parsing.NewTokenSeq(l)
		p := parsing.NewFileParser()

		p.Run(toks)

		for _, parseErr := range p.GetErrors() {
			fmt.Printf("%s %s: %v\n", file, parseErr.Token.Location, parseErr.Error)
		}

		for _, stmt := range p.Statements {
			d, err := json.Marshal(stmt)

			if err != nil {
				log.Warn().Err(err).Msg("failed to marshal statement")

				continue
			}

			fmt.Println(string(d))
		}
	}
}
