package main

import (
	"Goslang/evaluator"
	"Goslang/lexer"
	"Goslang/object"
	"Goslang/parser"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	//input := "var y=3;fn add(x:int){return x+3;}fn main() {return y+add(3);}main();"
	//repl.Start(os.Stdin, os.Stdout)
	startTime := time.Now()
	out := os.Stdout
	argFilePath := os.Args[1]
	file, err := os.OpenFile(argFilePath, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	var input strings.Builder
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	defer file.Close()

	for _, line := range fileLines {
		input.WriteString(line)
	}
	input.WriteString("main();")
	env := object.NewEnvironment()
	l := lexer.New(input.String())
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		endTimeWithErrors := time.Now()
		timeDiffWithErrors := endTimeWithErrors.Sub(startTime)
		compileTimeCommentWithError := fmt.Sprintf("Compilation Time: %d Milliseconds\n", timeDiffWithErrors.Milliseconds())
		printParserErrors(out, p.Errors(), compileTimeCommentWithError)
		return
	}
	evaluated := evaluator.Eval(program, env)
	if evaluated.Type() == object.ERROR_OBJ {
		endTimeWithErrObj := time.Now()
		timeDiffWithErrObj := endTimeWithErrObj.Sub(startTime)
		compileTimeCommentWithErrObj := fmt.Sprintf("Compilation Time: %d Milliseconds\n", timeDiffWithErrObj.Milliseconds())
		io.WriteString(out, compileTimeCommentWithErrObj)
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
		io.WriteString(out, "PROGRAM EXITED WITH CODE 1")
		return
	}
	endTimeWithoutErrors := time.Now()
	timeDiffWithoutErrors := endTimeWithoutErrors.Sub(startTime)
	compileTimeCommentWithoutError := fmt.Sprintf("Compilation Time: %d Milliseconds\n", timeDiffWithoutErrors.Milliseconds())
	if evaluated != nil {
		io.WriteString(out, compileTimeCommentWithoutError+"\nPretty fast, huh?\n")
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
		io.WriteString(out, "PROGRAM EXITED WITH CODE 0")
	}

}

func printParserErrors(out io.Writer, errors []string, comment string) {
	io.WriteString(out, comment)
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
	io.WriteString(out, "PROGRAM EXITED WITH CODE 1")
}
func add(x int, y int) int {
	w := x + y
	return w
}
