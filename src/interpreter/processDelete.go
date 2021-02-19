/*
	processDelete Function.
*/

package interpreter

import (
	"../fract"
	"../fract/name"
	"../objects"
	"../utilities/vector"
)

// processDelete Process delete keyword.
// tokens Tokens to process.
func (i *Interpreter) processDelete(tokens *vector.Vector) {
	// Value is not defined?
	if len(tokens.Vals) < 2 {
		first := tokens.Vals[0].(objects.Token)
		fract.ErrorCustom(first.File.Path, first.Line, first.Column+len(first.Value),
			"Value is not found!")
	}

	comma := false
	for index := 1; index < len(tokens.Vals); index++ {
		current := tokens.Vals[index].(objects.Token)

		if comma {
			if current.Type != fract.TypeComma {
				fract.Error(current, "Comma is not found!")
			}
			comma = false
			continue
		}

		// Token is not a deletable object?
		if current.Type != fract.TypeName {
			fract.Error(current, "This is not deletable object!")
		}

		position := name.VarIndexByName(i.vars, current.Value)

		// Name is not defined?
		if position == -1 {
			fract.Error(current, "Name is not defined!")
		}

		i.vars.RemoveRange(position, 1)
		comma = true
	}
	//
	if !comma {
		fract.Error(tokens.Vals[len(tokens.Vals)-1].(objects.Token), "The extra comma defined!")
	}
}
