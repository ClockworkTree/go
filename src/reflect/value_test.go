package reflect

import (
	"testing"
	"unsafe"
)

func TestValue_CanSet(t *testing.T) {
	type fields struct {
		typ  *rtype
		ptr  unsafe.Pointer
		flag flag
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				typ:  tt.fields.typ,
				ptr:  tt.fields.ptr,
				flag: tt.fields.flag,
			}
			if got := v.CanSet(); got != tt.want {
				t.Errorf("CanSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
