// Package random实现了一些基本函数来生成随机int和string。
package random

import (
	crand "crypto/rand"
	"fmt"
	"github.com/Songtingsen/go-utils/mathutil"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
	"unsafe"
)

const (
	MaximumCapacity = math.MaxInt32>>1 + 1
	Numeral         = "0123456789"
	LowwerLetters   = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SymbolChars     = "!@#$%^&*()_+-=[]{}|;':\",./<>?"
	AllChars        = Numeral + LowwerLetters + UpperLetters + SymbolChars
)

var rn = rand.NewSource(time.Now().UnixNano())

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandBool 生成一个随机的布尔值（真或假）。
func RandBool() bool {
	return rand.Intn(2) == 1
}

// RandBoolSlice 生成指定长度的随机布尔切片。
// https://go.dev/play/p/JXP_ettl56T
func RandBoolSlice(length int) []bool {
	if length <= 0 {
		return []bool{}
	}

	result := make([]bool, length)
	for i := range result {
		result[i] = RandBool()
	}

	return result
}

// RandInt 在[min，max）之间生成随机整数。
func RandInt(min, max int) int {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}

	if min == 0 && max == math.MaxInt {
		return rand.Int()
	}

	return rand.Intn(max-min) + min
}

// RandIntSlice 生成一个随机整数切片。 生成的整数在min和max之间（不包括）。
func RandIntSlice(length, min, max int) []int {
	if length <= 0 || min > max {
		return []int{}
	}

	result := make([]int, length)
	for i := range result {
		result[i] = RandInt(min, max)
	}

	return result
}

// RandUniqueIntSlice 生成一个长度不重复的随机整数切片。
func RandUniqueIntSlice(length, min, max int) []int {
	if min > max {
		return []int{}
	}
	if length > max-min {
		length = max - min
	}

	nums := make([]int, length)
	used := make(map[int]struct{}, length)
	for i := 0; i < length; {
		r := RandInt(min, max)
		if _, use := used[r]; use {
			continue
		}
		used[r] = struct{}{}
		nums[i] = r
		i++
	}

	return nums
}

// RandFloat 生成具有特定精度的[min，max）之间的随机float64数。
func RandFloat(min, max float64, precision int) float64 {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}

	n := rand.Float64()*(max-min) + min

	return mathutil.RoundToFloat(n, precision)
}

// RandFloats 生成一个长度为64个不重复的随机浮点数的切片。
func RandFloats(length int, min, max float64, precision int) []float64 {
	nums := make([]float64, length)
	used := make(map[float64]struct{}, length)
	for i := 0; i < length; {
		r := RandFloat(min, max, precision)
		if _, use := used[r]; use {
			continue
		}
		used[r] = struct{}{}
		nums[i] = r
		i++
	}

	return nums
}

// RandBytes 生成随机字节切片。
func RandBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	b := make([]byte, length)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return nil
	}

	return b
}

// RandString 生成指定长度的随机英文字符串。
func RandString(length int) string {
	return random(Letters, length)
}

// RandStringSlice sliceLen根据字符集生成长度为strLen的随机字符串切片。
// 图表集应该是以下之一：随机。数字，随机。LowWerLetters，随机。UpperLetters
// 随机。信件，随机。符号，随机。全查。或它们的组合。
func RandStringSlice(charset string, sliceLen, strLen int) []string {
	if sliceLen <= 0 || strLen <= 0 {
		return []string{}
	}

	result := make([]string, sliceLen)

	for i := range result {
		result[i] = random(charset, strLen)
	}

	return result
}

// RandFromGivenSlice 从给定的切片生成随机元素。
func RandFromGivenSlice[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return slice[rand.Intn(len(slice))]
}

// RandSliceFromGivenSlice 从给定切片生成长度为num的随机切片。
// -如果可重复为真，则生成的切片可能包含重复元素。
func RandSliceFromGivenSlice[T any](slice []T, num int, repeatable bool) []T {
	if num <= 0 || len(slice) == 0 {
		return slice
	}

	if !repeatable && num > len(slice) {
		num = len(slice)
	}

	result := make([]T, num)
	if repeatable {
		for i := range result {
			result[i] = slice[rand.Intn(len(slice))]
		}
	} else {
		shuffled := make([]T, len(slice))
		copy(shuffled, slice)
		rand.Shuffle(len(shuffled), func(i, j int) {
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		})
		result = shuffled[:num]
	}
	return result
}

// RandUpper 生成指定长度的随机大写字符串。
func RandUpper(length int) string {
	return random(UpperLetters, length)
}

// RandLower 生成指定长度的随机小写字符串。
func RandLower(length int) string {
	return random(LowwerLetters, length)
}

// RandNumeral 生成指定长度的随机数字字符串。
func RandNumeral(length int) string {
	return random(Numeral, length)
}

// RandNumeralOrLetter 生成指定长度的随机数字或字母字符串。
func RandNumeralOrLetter(length int) string {
	return random(Numeral+Letters, length)
}

// RandSymbolChar  generate a random symbol char of specified length.
// symbol chars: !@#$%^&*()_+-=[]{}|;':\",./<>?.
func RandSymbolChar(length int) string {
	return random(SymbolChars, length)
}

// nearestPowerOfTwo 返回一个大于等于cap的最近的2的整数次幂，参考java8的hashmap的tableSizeFor函数
func nearestPowerOfTwo(cap int) int {
	n := cap - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 {
		return 1
	} else if n >= MaximumCapacity {
		return MaximumCapacity
	}
	return n + 1
}

// random 基于给定的字符串范围生成随机字符串。
func random(s string, length int) string {
	// 确保随机数生成器的种子是动态的
	pid := os.Getpid()
	timestamp := time.Now().UnixNano()
	rand.Seed(int64(pid) + timestamp)

	// 仿照strings.Builder
	// 创建一个长度为 length 的字节切片
	bytes := make([]byte, length)
	strLength := len(s)
	if strLength <= 0 {
		return ""
	} else if strLength == 1 {
		for i := 0; i < length; i++ {
			bytes[i] = s[0]
		}
		return *(*string)(unsafe.Pointer(&bytes))
	}
	// s的字符需要使用多少个比特位数才能表示完
	// letterIdBits := int(math.Ceil(math.Log2(strLength))),下面比上面的代码快
	letterIdBits := int(math.Log2(float64(nearestPowerOfTwo(strLength))))
	// 最大的字母id掩码
	var letterIdMask int64 = 1<<letterIdBits - 1
	// 可用次数的最大值
	letterIdMax := 63 / letterIdBits
	// 循环生成随机字符串
	for i, cache, remain := length-1, rn.Int63(), letterIdMax; i >= 0; {
		// 检查随机数生成器是否用尽所有随机数
		if remain == 0 {
			cache, remain = rn.Int63(), letterIdMax
		}
		// 从可用字符的字符串中随机选择一个字符
		if idx := int(cache & letterIdMask); idx < strLength {
			bytes[i] = s[idx]
			i--
		}
		// 右移比特位数，为下次选择字符做准备
		cache >>= letterIdBits
		remain--
	}
	// 仿照strings.Builder用unsafe包返回一个字符串，避免拷贝
	// 将字节切片转换为字符串并返回
	return *(*string)(unsafe.Pointer(&bytes))
}

// UUIdV4 根据RFC 4122生成版本4的随机UUID。
// https://go.dev/play/p/nEGqyG_wpsN
func UUIdV4() (string, error) {
	uuid := make([]byte, 16)

	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// RandNumberOfLength 生成一个长度为len的随机数
func RandNumberOfLength(len int) int {
	m := int(math.Pow10(len) - 1)
	i := int(math.Pow10(len - 1))
	ret := rand.Intn(m-i+1) + i

	return ret
}
