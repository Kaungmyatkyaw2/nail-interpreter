package repl

import (
	"bufio"
	"fmt"
	"io"
	"nail/lexer"
	"nail/token"
)

const PROMPT = ">> "

func Start(in io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
