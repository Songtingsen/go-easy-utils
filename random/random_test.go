package random

import (
	"fmt"
	"testing"
)

// 测试RandBool函数
func TestRandBool(t *testing.T) {
	result := RandBool()
	fmt.Printf("RandBool()实际结果: %v\n", result)
	if result != true && result != false {
		t.Errorf("RandBool()返回值不符合预期，期望返回true或者false，实际返回: %v", result)
	}
}

// 测试RandBoolSlice函数
func TestRandBoolSlice(t *testing.T) {
	slice := RandBoolSlice(5)
	fmt.Printf("RandBoolSlice(5)实际生成的切片: %v\n", slice)
	if len(slice) != 5 {
		t.Errorf("RandBoolSlice(5)生成的切片长度不符合预期，期望长度为5，实际长度: %d", len(slice))
	}
	for _, b := range slice {
		if b != true && b != false {
			t.Errorf("RandBoolSlice(5)切片中的元素不符合预期，期望元素为true或者false，实际元素: %v", b)
		}
	}
}

// 测试RandInt函数
func TestRandInt(t *testing.T) {
	result := RandInt(1, 10)
	fmt.Printf("RandInt(1, 10)实际结果: %d\n", result)
	if result < 1 || result >= 10 {
		t.Errorf("RandInt(1, 10)返回值不符合预期，期望返回值在[1, 10)之间，实际返回: %d", result)
	}
}

// 测试RandIntSlice函数
func TestRandIntSlice(t *testing.T) {
	slice := RandIntSlice(3, 5, 10)
	fmt.Printf("RandIntSlice(3, 5, 10)实际生成的切片: %v\n", slice)
	if len(slice) != 3 {
		t.Errorf("RandIntSlice(3, 5, 10)生成的切片长度不符合预期，期望长度为3，实际长度: %d", len(slice))
	}
	for _, num := range slice {
		if num < 5 || num >= 10 {
			t.Errorf("RandIntSlice(3, 5, 10)切片中的元素不符合预期，期望元素在[5, 10)之间，实际元素: %d", num)
		}
	}
}

// 测试RandUniqueIntSlice函数
func TestRandUniqueIntSlice(t *testing.T) {
	slice := RandUniqueIntSlice(5, 1, 10)
	fmt.Printf("RandUniqueIntSlice(5, 1, 10)实际生成的切片: %v\n", slice)
	if len(slice) != 5 {
		t.Errorf("RandUniqueIntSlice(5, 1, 10)生成的切片长度不符合预期，期望长度为5，实际长度: %d", len(slice))
	}
	numMap := make(map[int]struct{})
	for _, num := range slice {
		if num < 1 || num >= 10 {
			t.Errorf("RandUniqueIntSlice(5, 1, 10)切片中的元素不符合预期，期望元素在[1, 10)之间，实际元素: %d", num)
		}
		if _, ok := numMap[num]; ok {
			t.Errorf("RandUniqueIntSlice(5, 1, 10)生成的切片存在重复元素，不符合预期")
		}
		numMap[num] = struct{}{}
	}
}

// 测试RandFloat函数
func TestRandFloat(t *testing.T) {
	result := RandFloat(1.0, 10.0, 2)
	fmt.Printf("RandFloat(1.0, 10.0, 2)实际结果: %f\n", result)
	if result < 1.0 || result >= 10.0 {
		t.Errorf("RandFloat(1.0, 10.0, 2)返回值不符合预期，期望返回值在[1.0, 10.0)之间，实际返回: %f", result)
	}
}

// 测试RandFloats函数
func TestRandFloats(t *testing.T) {
	slice := RandFloats(3, 1.0, 10.0, 2)
	fmt.Printf("RandFloats(3, 1.0, 10.0, 2)实际生成的切片: %v\n", slice)
	if len(slice) != 3 {
		t.Errorf("RandFloats(3, 1.0, 10.0, 2)生成的切片长度不符合预期，期望长度为3，实际长度: %d", len(slice))
	}
	for _, num := range slice {
		if num < 1.0 || num >= 10.0 {
			t.Errorf("RandFloats(3, 1.0, 10.0, 2)切片中的元素不符合预期，期望元素在[1.0, 10.0)之间，实际元素: %f", num)
		}
	}
}

// 测试RandBytes函数
func TestRandBytes(t *testing.T) {
	bytes := RandBytes(5)
	fmt.Printf("RandBytes(5)实际生成的字节切片: %v\n", bytes)
	if len(bytes) != 5 {
		t.Errorf("RandBytes(5)生成的字节切片长度不符合预期，期望长度为5，实际长度: %d", len(bytes))
	}
}

// 测试RandString函数
func TestRandString(t *testing.T) {
	str := RandString(5)
	fmt.Printf("RandString(5)实际生成的字符串: %v\n", str)
	if len(str) != 5 {
		t.Errorf("RandString(5)生成的字符串长度不符合预期，期望长度为5，实际长度: %d", len(str))
	}
}

// 测试RandStringSlice函数
func TestRandStringSlice(t *testing.T) {
	slice := RandStringSlice(Letters, 3, 5)
	fmt.Printf("RandStringSlice(Letters, 3, 5)实际生成的字符串切片: %v\n", slice)
	if len(slice) != 3 {
		t.Errorf("RandStringSlice(Letters, 3, 5)生成的字符串切片长度不符合预期，期望长度为3，实际长度: %d", len(slice))
	}
	for _, s := range slice {
		fmt.Printf("RandStringSlice(Letters, 3, 5)字符串切片中的元素: %v\n", s)
		if len(s) != 5 {
			t.Errorf("RandStringSlice(Letters, 3, 5)生成的字符串切片中元素长度不符合预期，期望元素长度为5，实际长度: %d", len(s))
		}
	}
}

// 测试RandFromGivenSlice函数
func TestRandFromGivenSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	result := RandFromGivenSlice(slice)
	fmt.Printf("RandFromGivenSlice从切片中取元素实际结果: %v\n", result)
	if result != 1 && result != 2 && result != 3 {
		t.Errorf("RandFromGivenSlice从切片中取元素不符合预期，期望返回切片中的元素，实际返回: %v", result)
	}
}

// 测试RandSliceFromGivenSlice函数（可重复情况）
func TestRandSliceFromGivenSliceRepeatable(t *testing.T) {
	slice := []int{1, 2, 3}
	result := RandSliceFromGivenSlice(slice, 5, true)
	fmt.Printf("RandSliceFromGivenSlice(可重复，取5个元素)实际生成的切片: %v\n", result)
	if len(result) != 5 {
		t.Errorf("RandSliceFromGivenSlice(可重复，取5个元素)生成的切片长度不符合预期，期望长度为5，实际长度: %d", len(result))
	}
}

// 测试RandSliceFromGivenSlice函数（不可重复情况）
func TestRandSliceFromGivenSliceNonRepeatable(t *testing.T) {
	slice := []int{1, 2, 3}
	result := RandSliceFromGivenSlice(slice, 3, false)
	fmt.Printf("RandSliceFromGivenSlice(不可重复，取3个元素)实际生成的切片: %v\n", result)
	if len(result) != 3 {
		t.Errorf("RandSliceFromGivenSlice(不可重复，取3个元素)生成的切片长度不符合预期，期望长度为3，实际长度: %d", len(result))
	}
	numMap := make(map[int]struct{})
	for _, num := range result {
		if _, ok := numMap[num]; ok {
			t.Errorf("RandSliceFromGivenSlice(不可重复，取3个元素)生成的切片存在重复元素，不符合预期")
		}
		numMap[num] = struct{}{}
	}
}

// 测试RandUpper函数
func TestRandUpper(t *testing.T) {
	str := RandUpper(5)
	fmt.Printf("RandUpper(5)实际生成的大写字符串: %v\n", str)
	if len(str) != 5 {
		t.Errorf("RandUpper(5)生成的大写字符串长度不符合预期，期望长度为5，实际长度: %d", len(str))
	}
	for _, char := range str {
		if char < 'A' || char > 'Z' {
			t.Errorf("RandUpper(5)生成的大写字符串中存在不符合预期的字符，期望为大写字母，实际字符: %c", char)
		}
	}
}

// 测试RandLower函数
func TestRandLower(t *testing.T) {
	str := RandLower(5)
	fmt.Printf("RandLower(5)实际生成的小写字符串: %v\n", str)
	if len(str) != 5 {
		t.Errorf("RandLower(5)生成的小写字符串长度不符合预期，期望长度为5，实际长度: %d", len(str))
	}
	for _, char := range str {
		if char < 'a' || char > 'z' {
			t.Errorf("RandLower(5)生成的小写字符串中存在不符合预期的字符，期望为小写字母，实际字符: %c", char)
		}
	}
}

// 测试RandNumeral函数
func TestRandNumeral(t *testing.T) {
	str := RandNumeral(5)
	fmt.Printf("RandNumeral(5)实际生成的数字字符串: %v\n", str)
	if len(str) != 5 {
		t.Errorf("RandNumeral(5)生成的数字字符串长度不符合预期，期望长度为5，实际长度: %d", len(str))
	}
	for _, char := range str {
		if char < '0' || char > '9' {
			t.Errorf("RandNumeral(5)生成的数字字符串中存在不符合预期的字符，期望为数字，实际字符: %c", char)
		}
	}
}

// 测试RandNumeralOrLetter函数
func TestRandNumeralOrLetter(t *testing.T) {
	str := RandNumeralOrLetter(5)
	fmt.Printf("RandNumeralOrLetter(5)实际生成的数字或字母字符串: %v\n", str)
	if len(str) != 5 {
		t.Errorf("RandNumeralOrLetter(5)生成的数字或字母字符串长度不符合预期，期望长度为5，实际长度: %d", len(str))
	}
}

// 测试RandSymbolChar函数
func TestRandSymbolChar(t *testing.T) {
	str := RandSymbolChar(5)
	fmt.Printf("RandSymbolChar(5)实际生成的符号字符串: %v\n", str)
	if len(str) != 5 {
		t.Errorf("RandSymbolChar(5)生成的符号字符串长度不符合预期，期望长度为5，实际长度: %d", len(str))
	}
}

// 测试nearestPowerOfTwo函数
func TestNearestPowerOfTwo(t *testing.T) {
	result := nearestPowerOfTwo(5)
	fmt.Printf("nearestPowerOfTwo(5)实际结果: %d\n", result)
	expected := 8
	if result != expected {
		t.Errorf("nearestPowerOfTwo(5)返回值不符合预期，期望返回最近的2的整数次幂为8，实际返回: %d", result)
	}
}

// 测试UUIdV4函数
func TestUUIdV4(t *testing.T) {
	uuid, err := UUIdV4()
	fmt.Printf("UUIdV4()实际生成的UUID: %v\n", uuid)
	if err != nil {
		t.Errorf("UUIdV4()生成UUID出错: %v", err)
	}
	if len(uuid) != 36 {
		t.Errorf("UUIdV4()生成的UUID长度不符合预期，期望长度为36，实际长度: %d", len(uuid))
	}
}

// 测试RandNumberOfLength函数
func TestRandNumberOfLength(t *testing.T) {
	result := RandNumberOfLength(3)
	fmt.Printf("RandNumberOfLength(3)实际结果: %d\n", result)
	if result < 100 || result >= 1000 {
		t.Errorf("RandNumberOfLength(3)返回值不符合预期，期望返回一个三位数，实际返回: %d", result)
	}
}
