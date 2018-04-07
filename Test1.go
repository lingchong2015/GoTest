package main

import (
	"fmt"
	"strings"
	//"runtime"
	//"mathoperator"
	"math"
	"io"
)

type Vertex struct {
	lat, lon float64
}

type MyFloat64 float64

func main() {
	//fmt.Println("My favorite number is", rand.Intn(10))
	//
	//fmt.Println(add(42, 13))

	//fmt.Println(add(2))
	//
	//fmt.Print("Go runs on ")
	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	// freebsd, openbsd,
	//	// plan9, windows...
	//	fmt.Printf("%s.", os)
	//}

	//m := make(map[string]Vertex)
	//fmt.Println(m)
	//m["lingchong"] = Vertex{
	//	1, 2,
	//}
	//fmt.Println(m)

	//fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

	//s := "lingchong ling chong lingchong chong"
	//results := WordCount(s)
	//for k, v := range results {
	//	fmt.Printf("Word: %s, Count: %d\n", k, v)
	//}

	//hypot := func(x, y float64) float64 {
	//	return math.Sqrt(x * x + y * y)
	//}
	//
	//fmt.Println(hypot(5, 12))
	//
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))

	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(pos(i), neg(-2 * i))
	//}

	//fn := Fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(fn())
	//}

	//v := Vertex{3, 4}
	//fmt.Println(v.Abs())

	//f := MyFloat64(-1)
	//fmt.Println(f.Abs())

	//hosts := map[string]IPAddr{
	//	"loopback":  {127, 0, 0, 1},
	//	"googleDNS": {8, 8, 8, 8},
	//}
	//for name, ip := range hosts {
	//	fmt.Printf("%v: %v\n", name, ip)
	//}

	//str2 := "hello"
	//data2 := []byte(str2)
	//fmt.Println(data2)
	//str2 = string(data2[:])
	//fmt.Println(str2)

	//fmt.Println(mathoperator.Sqrt(-2))

	s := strings.NewReader("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r := rot13Reader{s}
	r.Read()
}

/*
	type Reader interface {
 	       Read(p []byte) (n int, err error)
	}

	r := strings.NewReader("Hello, Reader!")

		b := make([]byte, 8)
		for {
			n, err := r.Read(b)
			fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
			fmt.Printf("b[:n] = %q\n", b[:n])
			if err == io.EOF {
				break
			}
		}
 */

type rot13Reader struct {
	r io.Reader
	i int32
}

func (rot13 rot13Reader) Read() (int32, error) {

	b := make([]byte, 1)
	for {
		_, err := rot13.r.Read(b)
		fmt.Printf("err = %v b = %v a = %v \n", err, b, string(b))
		if b[0] >= 65 && b[0] <= 90 {
			t:= b[0] + 13
			if t > 90 {
				t = t - 90 + 65 - 1
			}
			fmt.Printf("t = %v a = %v \n", t, string(t))
		} else if b[0] >= 97 && b[0] <= 122 {
			t:= b[0] + 13
			if t > 122 {
				t = t - 122 + 97 - 1
			}
			fmt.Printf("t = %v a = %v \n", t, string(t))
		}

		if err == io.EOF {
			break
		}
	}

	return 0, nil
}

func WordCount(s string) map[string]int32 {
	words := strings.Fields(s)
	wordMap := make(map[string]int32)
	for _, w := range words {
		_, ok := wordMap[w]
		if !ok {
			wordMap[w] = 1
		} else {
			wordMap[w] = wordMap[w] + 1
		}
	}

	return wordMap
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Fibonacci() func() int {
	i := -1
	j := 1
	return func() int {
		result := i + j
		i = j
		j = result
		return result
	}
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.lat * v.lat + v.lon * v.lon)
}

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

type IPAddr [4]byte

func (ipAddr IPAddr) String() string{
	result := ""
	for _, v := range ipAddr {
		result += fmt.Sprintf("%v.", v)
	}
	return result[:len(result) - 1]
}
