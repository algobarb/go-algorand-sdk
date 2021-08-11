package abi

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func testValidFramework(funcName string, expected interface{}, actual interface{}, t *testing.T) {
	// add equal method for struct Type
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s testing error: expect %s, got %s", funcName, expected, actual)
	}
}

func TestMakeUintTypeValid(t *testing.T) {
	for i := 8; i <= 512; i += 8 {
		uintType, _ := MakeUintType(uint16(i))
		expected := "uint" + strconv.Itoa(i)
		actual := uintType.String()
		testValidFramework("MakeUintType", expected, actual, t)
	}
}

func TestMakeUintTypeInvalid(t *testing.T) {
	// TODO something need to be added
	// uintType, err := MakeUintType(uint16(7))
	// require.Error(t, err)
	// require.Equal(t, "expected message", err.Error())
}

func TestTypeFromStringUintTypeValid(t *testing.T) {
	for i := 8; i <= 512; i += 8 {
		expected, _ := MakeUintType(uint16(i))
		actual, err := TypeFromString(expected.String())
		if err != nil {
			t.Errorf("TypeFromString testing error: Parsing error for %s", expected.String())
		}
		testValidFramework("TypeFromString uint", expected, actual, t)
	}
}

func TestTypeFromStringUintTypeInvalid(t *testing.T) {
	// TODO something need to be added

	// testCases := []struct {
	// 	input string
	// 	expected Type
	// }{
	// 	{
	// 		input: "uint64",
	// 		expected: MakeUintType(uint16(64)),
	// 	},
	// 	{
	// 		input: "(uint64, (uint128, ufixed512x100))",
	// 		expected: MakeUintType(uint16(64)),
	// 	}
	// }

	// for index, test := range ... {
	// 	t.Run(fmt.Sprintf("test %d", index), func(t *testing.T) {
	// 		t.Error("abc")
	// 	})
	// }

}

func TestMakeByteTypeValid(t *testing.T) {
	byteType := MakeByteType()
	expected := byteType.String()
	actual := "byte"
	testValidFramework("MakeByteType", expected, actual, t)
}

func TestTypeFromStringByteTypeValid(t *testing.T) {
	expected := MakeByteType()
	actual, err := TypeFromString(expected.String())
	if err != nil {
		t.Errorf("TypeFromString testing error: Parsing error for %s", expected.String())
	}
	testValidFramework("TypeFromString byte", expected, actual, t)
}

func TestMakeUfixedTypeValid(t *testing.T) {
	for i := 8; i <= 512; i += 8 {
		for j := 1; j <= 160; j++ {
			ufixedType, _ := MakeUFixedType(uint16(i), uint16(j))
			expected := "ufixed" + strconv.Itoa(i) + "x" + strconv.Itoa(j)
			actual := ufixedType.String()
			testValidFramework("MakeUfixedType", expected, actual, t)
		}
	}
}

func TestTypeFromStringUfixedTypeValid(t *testing.T) {
	for i := 8; i <= 512; i += 8 {
		for j := 1; j <= 160; j++ {
			expected, _ := MakeUFixedType(uint16(i), uint16(j))
			actual, err := TypeFromString("ufixed" + strconv.Itoa(i) + "x" + strconv.Itoa(j))
			if err != nil {
				fmt.Println(err)
				t.Errorf("TypeFromString testing error: Parsing error for %s", expected.String())
			}
			testValidFramework("TypeFromString ufixed", expected, actual, t)
		}
	}
}

func TestMakeAddressTypeValid(t *testing.T) {
	addressType := MakeAddressType()
	expected := addressType.String()
	actual := "address"
	testValidFramework("MakeAddressType", expected, actual, t)
}

func TestTypeFromStringAddressTypeValid(t *testing.T) {
	expected := MakeAddressType()
	actual, err := TypeFromString(expected.String())
	if err != nil {
		t.Errorf("TypeFromString testing error: Parsing error for %s", expected.String())
	}
	testValidFramework("TypeFromString address", expected, actual, t)
}

func TestMakeStringTypeValid(t *testing.T) {
	stringType := MakeStringType()
	expected := stringType.String()
	actual := "string"
	testValidFramework("MakeStringType", expected, actual, t)
}

func TestTypeFromStringStringTypeValid(t *testing.T) {
	expected := MakeStringType()
	actual, err := TypeFromString(expected.String())
	if err != nil {
		t.Errorf("TypeFromString testing error: Parsing error for %s", expected.String())
	}
	testValidFramework("TypeFromString string", expected, actual, t)
}

func TestMakeBoolTypeValid(t *testing.T) {
	boolType := MakeBoolType()
	expected := boolType.String()
	actual := "bool"
	testValidFramework("MakeBoolType", expected, actual, t)
}

func TestTypeFromStringBoolTypeValid(t *testing.T) {
	expected := MakeBoolType()
	actual, err := TypeFromString(expected.String())
	if err != nil {
		t.Errorf("TypeFromString testing error: Parsing error for %s", expected.String())
	}
	testValidFramework("TypeFromString bool", expected, actual, t)
}
