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
	iterateInput(os.Stdin, os.Stdout)
}

func iterateInput(input io.Reader, out io.Writer) {
	skipLines := false
	key := ""
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ln := scanner.Text()
		if strings.TrimSpace(ln) == key {
			skipLines = false
		}
		if skipLines {
			fmt.Fprintln(out, ln)
		} else {
			processLine(ln, out)
		}
		if strings.Contains(ln, "<<~") {
			skipLines = true
			key = after(ln, "<<~")
		}
	}
	return
}

// process in-coming text
func processLine(raw string, out io.Writer) {
	var toPrint interface{}
	cleaned := cleanRawInput(raw)
	if len(cleaned) < 1 && len(raw) > 0 {
		// filtered text shouldn't print a new line
		return
	}

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
		toPrint = trimmed
	default:
		toPrint = raw
	}
	fmt.Fprintln(out, toPrint)
	return
}

func cleanRawInput(raw string) string {
	ansi := "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	re := regexp.MustCompile(ansi)
	nocolor := re.ReplaceAllString(raw, "")

	refreshing := "Refreshing state... "
	if strings.Contains(nocolor, refreshing) {
		return ""
	}
	return nocolor
}

// Get substring after a string.
func after(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}
