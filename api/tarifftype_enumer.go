// Code generated by "enumer -type TariffType -trimprefix TariffType -transform=lower"; DO NOT EDIT.

package api

import (
	"fmt"
	"strings"
)

const _TariffTypeName = "pricestaticpricedynamicco2"

var _TariffTypeIndex = [...]uint8{0, 11, 23, 26}

const _TariffTypeLowerName = "pricestaticpricedynamicco2"

func (i TariffType) String() string {
	i -= 1
	if i < 0 || i >= TariffType(len(_TariffTypeIndex)-1) {
		return fmt.Sprintf("TariffType(%d)", i+1)
	}
	return _TariffTypeName[_TariffTypeIndex[i]:_TariffTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TariffTypeNoOp() {
	var x [1]struct{}
	_ = x[TariffTypePriceStatic-(1)]
	_ = x[TariffTypePriceDynamic-(2)]
	_ = x[TariffTypeCo2-(3)]
}

var _TariffTypeValues = []TariffType{TariffTypePriceStatic, TariffTypePriceDynamic, TariffTypeCo2}

var _TariffTypeNameToValueMap = map[string]TariffType{
	_TariffTypeName[0:11]:       TariffTypePriceStatic,
	_TariffTypeLowerName[0:11]:  TariffTypePriceStatic,
	_TariffTypeName[11:23]:      TariffTypePriceDynamic,
	_TariffTypeLowerName[11:23]: TariffTypePriceDynamic,
	_TariffTypeName[23:26]:      TariffTypeCo2,
	_TariffTypeLowerName[23:26]: TariffTypeCo2,
}

var _TariffTypeNames = []string{
	_TariffTypeName[0:11],
	_TariffTypeName[11:23],
	_TariffTypeName[23:26],
}

// TariffTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TariffTypeString(s string) (TariffType, error) {
	if val, ok := _TariffTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TariffTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TariffType values", s)
}

// TariffTypeValues returns all values of the enum
func TariffTypeValues() []TariffType {
	return _TariffTypeValues
}

// TariffTypeStrings returns a slice of all String values of the enum
func TariffTypeStrings() []string {
	strs := make([]string, len(_TariffTypeNames))
	copy(strs, _TariffTypeNames)
	return strs
}

// IsATariffType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TariffType) IsATariffType() bool {
	for _, v := range _TariffTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
