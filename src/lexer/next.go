/*
	Next Function
*/

package lexer

import (
	"../fract"
	"../utilities/vector"
)

// Next Lex next line.
func (l *Lexer) Next() *vector.Vector {
	tokens := vector.New()

	// If file is finished?
	if l.Finished {
		return tokens
	}

	// Reset bracket counter.
	l.braceCount = 0

tokenize:

	// Restore to defaults.
	l.Column = 1
	l.lastToken.Type = fract.TypeNone
	l.lastToken.Line = 0
	l.lastToken.Column = 0
	l.lastToken.Value = ""

	// Tokenize line.
	token := l.Generate()
	for token.Value != "" {
		tokens.Append(token)
		l.lastToken = token
		token = l.Generate()
	}

	// Go next line.
	l.Line++

	// Line equals to or bigger then last line.
	l.Finished = l.Line > len(l.File.Lines.Vals)

	// Lexer file is finished?
	if l.Finished {
		/* Check parentheses. */
		if l.braceCount > 0 {
			l.Line-- // Subtract for correct line number.
			l.Error("Parentheses is expected to close...")
			goto tokenize
		} else if l.BlockCount > 0 {
			l.Line--
			l.Error("Block is expected ending...")
		}
	}

	return tokens
}
