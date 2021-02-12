/*
	processTokens Functions.
*/

package interpreter

import (
	"fmt"

	"../fract"
	"../grammar"
	"../objects"
	"../utilities/vector"
)

// processTokens Process tokens and returns true if block end, returns false if not.
// tokens Tokens to process.
// do Do processes?
func (i *Interpreter) processTokens(tokens *vector.Vector, do bool) {
	// Skip this loop if tokens are empty.
	if !tokens.Any() {
		return
	}

	first := tokens.At(0).(objects.Token)

	if first.Type == fract.TypeValue || first.Type == fract.TypeBrace ||
		first.Type == fract.TypeName || first.Type == fract.TypeBooleanTrue ||
		first.Type == fract.TypeBooleanFalse {
		// Check variable set statement?
		if first.Type == fract.TypeName {
			for index := 1; index < tokens.Len(); index++ {
				current := tokens.At(index).(objects.Token)
				if current.Type == fract.TypeOperator &&
					current.Value == grammar.Setter { // Variable setting.
					i.processVariableSet(tokens)
					return
				}
			}
		}

		// Println
		value := i.processValue(tokens)
		if value.Type == fract.VTIntegerArray || value.Type == fract.VTFloatArray {
			fmt.Println(value.Content)
		} else {
			fmt.Println(value.Content[0])
		}
	} else if first.Type == fract.TypeVariable { // Variable definition.
		i.processVariableDefinition(tokens)
	} else if first.Type == fract.TypeDelete { // Delete from memory.
		i.processDelete(tokens)
	} else if first.Type == fract.TypeIf { // if-elif-else.
		i.processIf(tokens, do)
	} else if first.Type == fract.TypeLoop { // Loop.
		i.processLoop(tokens, do)
	} else {
		fract.Error(first, "What is this?: "+first.Value)
	}
}
