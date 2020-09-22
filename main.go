package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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
	var toPrint interface{}
	cleaned := cleanRawInput(raw)

	trimmed := strings.TrimSpace(cleaned)
	if len(trimmed) < 1 {
		fmt.Fprintln(out, raw)
		return
	}

	firstChar := string(trimmed[0])
	switch firstChar {
	case "~":
		ch := strings.Replace(cleaned, "~", "", 1)
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
		new := strings.Replace(cleaned, "+", "", 1)
		toPrint = Green("+" + new)
	case "-":
		new := strings.Replace(cleaned, "-", "", 1)
		toPrint = Red("-" + new)
	case "#":
		toPrint = strings.TrimSpace(cleaned)
	default:
		toPrint = raw
	}
	fmt.Fprintln(out, toPrint)
	return
}

func cleanRawInput(raw string) string {
	ansi := "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	re := regexp.MustCompile(ansi)
	return re.ReplaceAllString(raw, "")
}
