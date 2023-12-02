package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func InputToIntList(input io.ReadCloser, calibration func(string) (int64, error)) []int64 {
	results := []int64{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		c, err := calibration(s.Text())
		if err != nil {
			continue
		}
		results = append(results, c)
	}
	return results
}

func Sum(l []int64) int64 {
	var sum int64
	for _, v := range l {
		sum += v
	}
	return sum
}

func GetCalibrationValue(s string) (int64, error) {
	f := FindFirstDigit(s)
	l := FindLastDigit(s)

	calString := f + l
	calValue, err := strconv.ParseInt(calString, 10, 64)
	if err != nil {
		return -1, err
	}

	return calValue, nil
}

func FindFirstDigit(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return string(s[i])
		}
	}
	return ""
}

func FindLastDigit(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return string(s[i])
		}
	}
	return ""
}

func GetCalibrationValue2(s string) (int64, error) {
	f := FindFirstDigit2(s)
	l := FindLastDigit2(s)

	calString := f + l
	calValue, err := strconv.ParseInt(calString, 10, 64)
	if err != nil {
		return -1, err
	}

	return calValue, nil
}

func FindFirstDigit2(s string) string {
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if r == 'o' {
			// check for one
			if i+2 < len(s) && s[i+1] == 'n' && s[i+2] == 'e' {
				return "1"
			}
		}
		if r == 't' {
			if i+2 < len(s) && s[i+1] == 'w' && s[i+2] == 'o' {
				return "2"
			}
			if i+4 < len(s) && s[i+1] == 'h' && s[i+2] == 'r' && s[i+3] == 'e' && s[i+4] == 'e' {
				return "3"
			}
		}
		if r == 'f' {
			// check for four and five
			if i+3 < len(s) && s[i+1] == 'o' && s[i+2] == 'u' && s[i+3] == 'r' {
				return "4"
			}
			if i+3 < len(s) && s[i+1] == 'i' && s[i+2] == 'v' && s[i+3] == 'e' {
				return "5"
			}
		}
		if r == 's' {
			// check for six and seven
			if i+2 < len(s) && s[i+1] == 'i' && s[i+2] == 'x' {
				return "6"
			}
			if i+4 < len(s) && s[i+1] == 'e' && s[i+2] == 'v' && s[i+3] == 'e' && s[i+4] == 'n' {
				return "7"
			}
		}
		if r == 'e' {
			// check for eight
			if i+4 < len(s) && s[i+1] == 'i' && s[i+2] == 'g' && s[i+3] == 'h' && s[i+4] == 't' {
				return "8"
			}
		}
		if r == 'n' {
			// check for nine
			if i+3 < len(s) && s[i+1] == 'i' && s[i+2] == 'n' && s[i+3] == 'e' {
				return "9"
			}
		}
		if r >= '0' && r <= '9' {
			return string(r)
		}
	}
	return ""
}

func FindLastDigit2(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if r == 'e' {
			// check for one, three, five, nine
			if i-2 >= 0 && s[i-1] == 'n' && s[i-2] == 'o' {
				return "1"
			}
			if i-4 >= 0 && s[i-1] == 'e' && s[i-2] == 'r' && s[i-3] == 'h' && s[i-4] == 't' {
				return "3"
			}
			if i-3 >= 0 && s[i-1] == 'v' && s[i-2] == 'i' && s[i-3] == 'f' {
				return "5"
			}
			if i-3 >= 0 && s[i-1] == 'n' && s[i-2] == 'i' && s[i-3] == 'n' {
				return "9"
			}
		}
		if r == 'o' {
			// check for two
			if i-2 >= 0 && s[i-1] == 'w' && s[i-2] == 't' {
				return "2"
			}
		}
		if r == 'r' {
			// check for four
			if i-3 >= 0 && s[i-1] == 'u' && s[i-2] == 'o' && s[i-3] == 'f' {
				return "4"
			}
		}
		if r == 'x' {
			// check for six
			if i-2 >= 0 && s[i-1] == 'i' && s[i-2] == 's' {
				return "6"
			}
		}
		if r == 'n' {
			// check for seven
			if i-4 >= 0 && s[i-1] == 'e' && s[i-2] == 'v' && s[i-3] == 'e' && s[i-4] == 's' {
				return "7"
			}
		}
		if r == 't' {
			// check for eight
			if i-4 >= 0 && s[i-1] == 'h' && s[i-2] == 'g' && s[i-3] == 'i' && s[i-4] == 'e' {
				return "8"
			}
		}
		if r >= '0' && r <= '9' {
			return string(r)
		}
	}
	return ""
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	l := InputToIntList(f, GetCalibrationValue)
	fmt.Printf("%d\n", Sum(l))

	f, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	l2 := InputToIntList(f, GetCalibrationValue2)
	fmt.Printf("%d\n", Sum(l2))
}
