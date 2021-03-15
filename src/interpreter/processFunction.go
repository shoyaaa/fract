/*
	processFunction Function.
*/

package interpreter

import (
	"github.com/fract-lang/fract/src/fract"
	"github.com/fract-lang/fract/src/grammar"
	"github.com/fract-lang/fract/src/objects"
	"github.com/fract-lang/fract/src/utils/vector"
)

// processFunction Process function.
// tokens Tokens to process.
func (i *Interpreter) processFunction(tokens vector.Vector) {
	tokenLen := len(tokens.Vals)
	_name := tokens.Vals[1].(objects.Token)

	// Name is not name?
	if _name.Type != fract.TypeName {
		fract.Error(_name, "This is not a valid name!")
	}

	// Name is already defined?
	if i.functionIndexByName(_name.Value) != -1 {
		fract.Error(_name, "Already defined function in this name!: "+_name.Value)
	}

	// Function parentheses are not defined?
	if tokenLen < 4 {
		fract.Error(_name, "Where is the function parentheses?")
	}

	i.index++
	function := objects.Function{
		Name:       _name.Value,
		Start:      i.index,
		Parameters: []string{},
	}

	dtToken := tokens.Vals[tokenLen-1].(objects.Token)
	if dtToken.Type != fract.TypeBrace ||
		dtToken.Value != grammar.TokenRParenthes {
		fract.Error(dtToken, "Invalid syntax!")
	}

	paramList := tokens.Sublist(3, tokenLen-4)

	// Decompose function parameters.
	paramName := true
	for _, current := range paramList.Vals {
		current := current.(objects.Token)
		if paramName {
			if current.Type != fract.TypeName {
				fract.Error(current, "Parameter name is not found!")
			}
			function.Parameters = append(function.Parameters, current.Value)
			paramName = false
		} else {
			if current.Type != fract.TypeComma {
				fract.Error(current, "Comma is not found!")
			}
			paramName = true
		}
	}

	i.skipBlock()
	function.Tokens = i.tokens.Sublist(function.Start, i.index-function.Start+1).Vals
	i.funcs.Vals = append(i.funcs.Vals, function)
}
