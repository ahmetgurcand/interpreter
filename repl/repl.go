package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ahmetgurcand/interpreter/evaluator"
	"github.com/ahmetgurcand/interpreter/lexer"
	"github.com/ahmetgurcand/interpreter/object"
	"github.com/ahmetgurcand/interpreter/parser"
)

const PROMPT = ">> "
const MONKEY_FACE = `
	    __,__
    .--. .-"     "-. .--.
   / .. \/ .-. .-. \/ .. \
  | | ' |  /  Y   \  |' |  |
  | \   \ \ 0 | 0 / /   / |
   \ '- ,\.-"""""""-./, -'/
    ''-' /_  ^ ^   _\ '-''
	|   \._ _./   |
	\   \ '~' /   /
	 '._ '-=-' _.'
	    '-----'
`


/* Start reads from the input source until encountering a newline,
takes the just read line and passes it to an instance of our lexer 
and finally prints all the tokens the lexer gives us until we encounter EOF.
*/
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}