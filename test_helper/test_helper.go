package test_helper

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// H is a wrapper around testing.T which adds assertions of a sort
func H(t *testing.T) Helper {
	t.Helper()
	return Helper{t}
}

// Helper is a test helper, is has a private member t of type *testing.T which is used
// to print erorrs and fail test suites.
type Helper struct {
	t *testing.T
}

func (h Helper) StrEql(expect, result string) {
	h.t.Helper()
	if diff := cmp.Diff(expect, result); diff != "" {
		h.t.Errorf("string equality assertion failed (-expected, +result)\n%s", diff)
	}
}

func (h Helper) ErrEql(got, want error) {
	h.t.Helper()
	if got == nil && want == nil {
		return
	}
	if got != nil && want != nil {
		if got.Error() == want.Error() {
			return
		}
	}
	h.t.Fatalf("error equality assertion failed, got %q wanted %q", got, want)
}

func (h Helper) BoolEql(expect, result bool) {
	h.t.Helper()
	if expect != result {
		h.t.Fatalf("boolean equality assertion failed, expected %t got %t", expect, result)
	}
}

// func (h helper) InterfaceEql(got, want interface{}) {
// 	h.t.Helper()
// 	if diff := cmp.Diff(want, got); diff != "" {
// 		h.t.Errorf("interface{} equality assertion failed (-got +want)\n%s", diff)
// 	}
// }

// func (h helper) IsNil(any interface{}) {
// 	h.t.Helper()
// 	if any != nil {
// 		h.t.Fatalf("wanted not nil, got %v", any)
// 	}
// }

// func (h helper) NotNil(any interface{}) {
// 	h.t.Helper()
// 	if any == nil {
// 		h.t.Fatalf("wanted not nil, got %v", any)
// 	}
// }

// func (h helper) TypeEql(got, want interface{}) {
// 	h.t.Helper()
// 	// check obvious case
// 	if got == nil && want == nil {
// 		return
// 	}
// 	// check for type equality
// 	if strings.Compare(fmt.Sprintf("%T", got), fmt.Sprintf("%T", want)) != 0 {
// 		h.t.Fatalf("type equality assertion failed, got %q wanted %q", fmt.Sprintf("%T", got), fmt.Sprintf("%T", want))
// 	}
// }

// func (h helper) IntEql(got, want int) {
// 	h.t.Helper()
// 	if got != want {
// 		h.t.Fatalf("integer equality assertion failed, got %d wanted %d", got, want)
// 	}
// }
