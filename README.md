Formatting library
==================

Formatting is a small (one file) library for formatting text. Currently it only
contains the IndentingWriter class, which is used to write text with variable
amounts of whitespace at the beginning of a line. Generally this would be used
to format source code, like this:

    buffer := bytes.NewBuffer(make([]byte, 0))
    ind := formatting.NewIndentingWriter(buffer)

    ind.Println("for {")
	ind.IncrIndent()
	ind.Println("// Do some stuff")
	ind.DecrIndent()
	ind.Println("}")