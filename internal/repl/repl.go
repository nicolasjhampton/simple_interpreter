package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/nicolasjhampton/simple_interpreter/internal/token"
	"github.com/nicolasjhampton/simple_interpreter/internal/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return 
		}

		line := scanner.Text()
		l := lexer.New(line)

		for nToken := l.NextToken(); nToken.Type != token.EOF; nToken = l.NextToken() {
			fmt.Printf("%+v\n", nToken)
		}
	}
}