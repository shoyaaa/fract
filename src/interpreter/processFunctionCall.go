/*
	processFunctionCall Function.
*/

package interpreter

import (
	"../fract"
	"../grammar"
	"../objects"
	"../parser"
	"../utilities/vector"
)

// processArgument Process function argument.
// function Function.
// current Current token.
// count Count of appended arguments.
// value Value instance.
func processArgument(function *objects.Function, current *objects.Token,
	count int, value objects.Value) objects.Variable {
	if count >= len(function.Parameters) {
		fract.Error(*current, "Argument overflow!")
	}
	return objects.Variable{
		Name:  function.Parameters[count],
		Const: false,
		Value: value,
	}
}

// processFunctionCall Process function call.
// tokens Tokens to process.
func (i *Interpreter) processFunctionCall(tokens vector.Vector) objects.Value {
	_name := tokens.Vals[0].(objects.Token)

	// Name is not defined?
	nameIndex := i.functionIndexByName(_name.Value)
	if nameIndex == -1 {
		fract.Error(_name, "Function is not defined!: "+_name.Value)
	}

	i.blockCount++

	function := i.funcs.Vals[nameIndex].(objects.Function)
	variableLen := len(i.vars.Vals)

	// Decompose arguments.
	tokens, _ = parser.DecomposeBrace(&tokens, grammar.TokenLParenthes,
		grammar.TokenRParenthes, false)
	braceCount := 0
	lastComma := 0
	count := 0
	for index := range tokens.Vals {
		current := tokens.Vals[index].(objects.Token)
		if current.Type == fract.TypeBrace {
			if current.Value == grammar.TokenLParenthes ||
				current.Value == grammar.TokenLBrace ||
				current.Value == grammar.TokenLBracket {
				braceCount++
			} else if current.Value == grammar.TokenRParenthes ||
				current.Value == grammar.TokenRBrace ||
				current.Value == grammar.TokenRBracket {
				braceCount--
			}
		} else if current.Type == fract.TypeComma && braceCount == 0 {
			valueList := tokens.Sublist(lastComma, index-lastComma)
			if len(valueList.Vals) == 0 {
				fract.Error(current, "Value is not defined!")
			}
			i.vars.Vals = append(i.vars.Vals, processArgument(&function, &current, count,
				i.processValue(valueList)))
			count++
			lastComma = index + 1
		}
	}

	if tokenLen := len(tokens.Vals); lastComma < tokenLen {
		current := tokens.Vals[lastComma].(objects.Token)
		valueList := tokens.Sublist(lastComma, tokenLen-lastComma)
		if len(valueList.Vals) == 0 {
			fract.Error(current, "Value is not defined!")
		}
		i.vars.Vals = append(i.vars.Vals, processArgument(&function, &current, count,
			i.processValue(valueList)))
		count++
	}

	// All parameters is not defined?
	if count != len(function.Parameters) {
		fract.Error(_name, "All parameters is not defined!")
	}

	functionLen := len(i.funcs.Vals)
	returnValue := objects.Value{
		Content: nil,
	}

	nameIndex = i.index
	itokens := i.tokens
	i.tokens = function.Tokens
	i.index = 0

	// Process block.
	i.functions++
	for ; i.index < len(i.tokens.Vals); i.index++ {
		tokens := i.tokens.Vals[i.index].(vector.Vector)
		if tokens.Vals[0].(objects.Token).Type == fract.TypeBlockEnd { // Block is ended.
			break
		} else if i.processTokens(&tokens, true) == fract.FUNCReturn {
			tokens := i.tokens.Vals[i.returnIndex].(vector.Vector)
			i.returnIndex = fract.TypeNone
			valueList := vector.New(tokens.Vals[1:]...)
			if len(valueList.Vals) == 0 {
				break
			}
			returnValue = i.processValue(valueList)
			break
		}
	}

	// Remove temporary variables.
	i.vars.Vals = i.vars.Vals[:variableLen]
	// Remove temporary functions.
	i.funcs.Vals = i.funcs.Vals[:functionLen]

	i.functions--
	i.index = nameIndex
	i.tokens = itokens
	i.subtractBlock(nil)

	return returnValue
}
