package functions

import (
	"fmt"

	"github.com/fract-lang/fract/pkg/arithmetic"
	"github.com/fract-lang/fract/pkg/fract"
	obj "github.com/fract-lang/fract/pkg/objects"
)

// Int Convert object to integer.
// f Function.
// parameters Parameters.
func Int(f obj.Function, parameters []*obj.Variable) obj.Value {
	switch parameters[1].Value.Content[0].Data { // Cast type.
	case "strcode":
		data := parameters[0].Value.Content[0]
		if data.Type != fract.VALString || data.Data == "" || len(data.Data) > 1 {
			return obj.Value{
				Content: []obj.DataFrame{
					{
						Data: "-1",
						Type: fract.VALInteger,
					},
				},
			}
		}
		return obj.Value{
			Content: []obj.DataFrame{
				{
					Data: fmt.Sprint(parameters[0].Value.Content[0].Data[0]),
					Type: fract.VALInteger,
				},
			},
		}
	case "object":
		fallthrough
	default:
		return obj.Value{
			Content: []obj.DataFrame{
				{
					Data: fmt.Sprint(int64(arithmetic.ToArithmetic(parameters[0].Value.Content[0].Data))),
					Type: fract.VALInteger,
				},
			},
		}
	}
}
