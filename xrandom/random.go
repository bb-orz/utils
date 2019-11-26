package xrandom

import (
	"crypto/md5"
	"fmt"
	"github.com/gofuncchan/util/xdata"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// 生成简单伪随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成随机40位哈希
func GenHash(src string) string {
	// 1.获取当前时间戳
	unix := time.Now().Unix()
	// 2.将文件名和时间戳一起计算md5等到前32位十六进制字符
	hash := md5.New()
	io.WriteString(hash, src)
	io.WriteString(hash, strconv.Itoa(int(unix)))
	hb := hash.Sum(nil)

	// 获取时间戳前8位字符
	ub := strconv.Itoa(int(unix))[:8]

	// 组合输出40位哈希字符
	s := fmt.Sprintf("%x%s", hb, ub)

	return s
}

// 生成范围在[start,end), 类型为int的随机数
func GenRandomInt(start int, end int) int {
	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成随机数
	num := r.Intn((end - start)) + start

	return num
}

// 生成范围在[start,end), 类型为int的n个不重复随机数
func GenRandomIntList(start int, end int, count int) []int {
	// 范围检查
	if end < start || (end-start) < count {
		return nil
	}

	// 存放结果的slice
	nums := make([]int, 0)
	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		// 生成随机数
		num := r.Intn((end - start)) + start

		// 查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

// 生成范围在[start,end), 类型为int64的随机数
func GenRandomInt64(start int64, end int64) int64 {
	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成随机数
	num := r.Int63n((end - start)) + start

	return num
}

// 生成范围在[start,end), 类型为int64的n个不重复随机数
func GenRandomInt64List(start int64, end int64, count int) []int64 {
	// 范围检查
	if end < start || (end-start) < xdata.IntToInt64(count) {
		return nil
	}

	// 存放结果的slice
	nums := make([]int64, 0)
	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		// 生成随机数
		num := r.Int63n((end - start)) + start

		// 查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}


