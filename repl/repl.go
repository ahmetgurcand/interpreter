package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/ahmetgurcand/interpreter/lexer"
	"github.com/ahmetgurcand/interpreter/token"

)

const PROMPT = ">> "

/* Start reads from the input source until encountering a newline,
takes the just read line and passes it to an instance of our lexer 
and finally prints all the tokens the lexer gives us until we encounter EOF.
*/
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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}