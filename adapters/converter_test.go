package adapters

import (
	"github.com/ralugr/language-service/logger"
	"reflect"
	"testing"
)

func TestConvertFromByteToStringArray(t *testing.T) {
	input := []struct {
		b        []byte
		expected []string
		err      bool
	}{
		{[]byte("[\"Welcome\"]"), []string{"Welcome"}, false},
		{[]byte("[]"), []string{}, false},
		{[]byte("1234"), nil, true},
	}

	for tc, tt := range input {
		actual, err := ConvertFromByteToStringArray(tt.b)
		logger.Info.Println(" Starting test case ", tc)

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}

		if tt.err != (err != nil) {
			t.Errorf("\nExpected %v \nActual   %v", tt.err, err)
		}
	}
}
