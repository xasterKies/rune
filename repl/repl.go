package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/xasterKies/rune/lexer"
	"github.com/xasterKies/rune/parser"
)

const PROMPT = ">> "

const RUNE_LANG = `

	RRRRRRR  
	RR     RR 
	RR     RR 
	RRRRRRR  
	RR   RR  
	RR    RR 
	RR     RR
	
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
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


		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, RUNE_LANG)
	io.WriteString(out, "Woops! We ran into some runic business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}