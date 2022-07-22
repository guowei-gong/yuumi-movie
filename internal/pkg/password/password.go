package passwords

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"unicode"
)

const (
	// Cost 进行哈希的次数-数字越大生成bcrypt的速度越慢，成本越大。
	// 同样也意味着如果密码库被盗，攻击者想通过暴力破解的方法猜测出用户密码的成本变得越昂贵。
	Cost = 14
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), Cost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetHashingCost(hashedPassword []byte) (int, error) {
	cost, err := bcrypt.Cost(hashedPassword)
	if err != nil {
		return 0, err
	}
	return cost, nil
}

// ParsePassword parses a single password
func ParsePassword(s string, min, max int) (err error) {
	var (
		isMin   bool
		special bool
		number  bool
		upper   bool
		lower   bool
	)

	if len(s) < min || len(s) > max {
		isMin = false
		err = errors.New("length should be " + strconv.Itoa(min) + " to " + strconv.Itoa(max))
	}

	for _, c := range s {
		// 优化性能，如果在到达终点之前所有都成为 true
		if special && number && upper && lower && isMin {
			break
		}

		switch {
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	if !special {
		err = errors.New("应该至少包含一个特殊字符")
	}
	if !number {
		err = errors.New("应该包含至少一个数字")
	}
	if !lower {
		err = errors.New("应该包含至少一个小写字母")
	}
	if !upper {
		err = errors.New("应该包含至少一个大写字母")
	}
	return err
}
