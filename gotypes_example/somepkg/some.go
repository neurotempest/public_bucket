package somepkg

import (
	"github.com/neurotempest/public_bucket/gotypes_example/someotherpkg"

	"github.com/shopspring/decimal"
)

type SomeType struct{
	A otherpkg.SomeOtherType
	B decimal.Decimal
}
