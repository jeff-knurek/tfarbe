package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	. "github.com/logrusorgru/aurora"
)

func main() {
	iterateInput(os.Stdin)
}

func iterateInput(input io.Reader) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ln := scanner.Text()
		processLine(ln, os.Stdout)
	}
	return
}

// process in-coming text
func processLine(raw string, out io.Writer) {
	trimmed := strings.TrimSpace(raw)
	if len(trimmed) < 1 {
		fmt.Fprintln(out, raw)
		return
	}

	firstChar := string(trimmed[0])
	var toPrint interface{}
	switch firstChar {
	case "~":
		ch := strings.Replace(raw, "~", "", 1)
		sp := strings.SplitAfter(ch, " -> ")
		if len(sp) != 2 {
			toPrint = Yellow("~" + ch)
			break
		}
		new := sp[1]
		sp2 := strings.SplitAfter(sp[0], " = ")
		if len(sp2) != 2 {
			toPrint = Yellow("~" + ch)
			break
		}
		ch = "~" + sp2[0]
		old := sp2[1]
		toPrint = fmt.Sprintf("%s%s%s", Yellow(ch), Red(old), Green(new))
	case "+":
		new := strings.Replace(raw, "+", "", 1)
		toPrint = Green("+" + new)
	case "-":
		new := strings.Replace(raw, "-", "", 1)
		toPrint = Red("-" + new)
	case "#":
		toPrint = trimmed
	default:
		toPrint = raw
	}
	fmt.Fprintln(out, toPrint)
	return
}
