package remotecgorepo

// #cgo LDFLAGS: -lm
// double callC();
import "C"

func Call() float64 {
	return float64(C.callC())
}
