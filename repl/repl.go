package repl

import (
	"Goslang/evaluator"
	"Goslang/lexer"
	"Goslang/object"
	"Goslang/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

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
			if arr, ok := evaluated.(*object.PrintObject); ok {
				for _, element := range arr.Elements {
					switch element := element.(type) {
					case *object.ReturnVal:
						io.WriteString(out, element.Value.Inspect()+"\n")
					case *object.Integer:
						io.WriteString(out, element.Inspect()+"\n")
					case *object.String:
						io.WriteString(out, element.Value+"\n")

					default:
						io.WriteString(out, element.Inspect())
					}
				}
			}

			//io.WriteString(out, "\n")
			/*io.WriteString(out, "PROGRAM EXITED WITH CODE 0")*/
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
