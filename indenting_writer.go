/*
Package formatting is a library of text formatting utilities
*/
package formatting

import (
	"fmt"
	"io"
	"strings"
)

// IndentingWriter is a Writer implementation that inserts variable amounts
// of whitespace at the begginging of a line. This is generally most useful
// for formatting source code.
type IndentingWriter struct {
	doIndent, indentRequired bool
	currentIndent            int
	indentStr                string
	Writer                   io.Writer
}

// NewIndentingWriter returns an initialised IndentingWriter
func NewIndentingWriter(writer io.Writer) *IndentingWriter {
	wr := &IndentingWriter{
		doIndent:       true,
		indentRequired: true,
		currentIndent:  0,
		indentStr:      "\t",
		Writer:         writer,
	}

	return wr
}

// SetIndentString specifies the indent to use at the beginning of a line.
// The default value is '\t'
func (w *IndentingWriter) SetIndentString(s string) {
	w.indentStr = s
}

// IncrIndent increases the indent amount by 1
func (w *IndentingWriter) IncrIndent() {
	w.currentIndent++
}

// DecrIndent decreases the indent amount by 1
func (w *IndentingWriter) DecrIndent() {
	if w.currentIndent > 0 {
		w.currentIndent--
	}
}

func (wr *IndentingWriter) writeIndent() {
	for i := 0; i < wr.currentIndent; i++ {
		fmt.Fprint(wr.Writer, wr.indentStr)
	}
}

// Printf prints formatted text to the Writer, adding indents as needed
func (wr *IndentingWriter) Printf(format string, a ...interface{}) (n int, err error) {
	if wr.doIndent && wr.indentRequired {
		wr.writeIndent()
	}
	wr.indentRequired = wr.doIndent && endsInNewline(format)
	return fmt.Fprintf(wr.Writer, format, a...)
}

// Print prints text to the writer, adding indents as needed
func (wr *IndentingWriter) Print(a ...interface{}) (n int, err error) {
	if wr.doIndent && wr.indentRequired {
		wr.writeIndent()
	}
	wr.indentRequired = false
	return fmt.Fprint(wr.Writer, a...)
}

// Println prints text and a newline to the writer, adding indents as needed
func (wr *IndentingWriter) Println(a ...interface{}) (n int, err error) {
	if wr.doIndent && wr.indentRequired {
		wr.writeIndent()
	}
	wr.indentRequired = wr.doIndent && true
	return fmt.Fprintln(wr.Writer, a...)
}

// endsInNewline checks to see if a string ends in a newline.
func endsInNewline(s string) bool {
	trimmed := strings.TrimRight(s, "\t ")
	c := trimmed[len(trimmed)-1]
	return byte(c) == '\n' || byte(c) == '\r'
}
