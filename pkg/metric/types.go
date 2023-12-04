package metric

import "golang.org/x/exp/constraints"

type Accepted interface {
	constraints.Ordered | ~bool
}
