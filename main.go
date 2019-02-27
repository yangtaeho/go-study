package main

/*
https://blog.wonhada.com/?p=1930
단숨에 Go언어 배우기 따라치기
19-02-27
3 년 전 todoList 에 있던 걸 한 번 끄집어내봄
*/

import (
	"fmt"
	. "fmt"
	"math/rand"
	"os"
	tOs "os"
	"strconv"
	"strings"
	"time"
)

var print = fmt.Println

// ===================================

// var print = fmt.Println

// 3.들여쓰기는 반드시 탭(Tab) 키를 사용.
// 4. 주석은 C 스타일 (// 와 /**/)
// 5. 변수와 상수
var site1 string = "test.com"
var site2 = "test.com"

func rename(name string) {
	site2 := name
	print("site2 => " + site2)
}

var userCount1 int = 1

const userCount2 int = 3

var (
	site      string = "test"
	author    string = "test2"
	userCount int    = 100
)

// 6. 대입 연산자
func getInfo() string {
	res := site + ":" + author + ":" + strconv.Itoa(userCount)
	print(res)
	return res
}

// 7. 채널 Channel
func add(a int, b int, c chan int) {
	c <- a + b
}

// 8. 포인터
func add2(a int, b int, c *int) {
	*c = a + b
}

// 9. 조건문은 C 스타일 (if, switch)
func get(d *int) {
	switch *d {
	case 5:
		println("***** 5")
		fallthrough
	case 6:
		println(*d)
		fallthrough
	case 7:
		println("======")
	}
}

// 10. 반복문은 for만 존재, C 스타일
func iterate() {
	println("[s] iterate")
	rand.Seed(int64(time.Now().Nanosecond()))
	i := rand.Intn(10)
	if i != 5 {
		println("no loop i =>" + strconv.Itoa(i))
	}
	for i == 5 {
		println("무한도전" + strconv.Itoa(i))
		i++
	}
	println("[e] iterate" + strconv.Itoa(i))
}

// 11. 배열
func arr() {
	var arr [3]string = [3]string{"hell", "Golang", "world"}
	var slice []string = []string{"hell", "Glang"}
	slice = append(slice, "World")
	slice = slice[1:]
	fmt.Println(arr[0], strings.Join(slice, " "))
}

// 12. 맵
func getMap() map[string]string {
	var _map map[string]string = map[string]string{
		"name":    "taeho yang",
		"country": "Korea",
		"company": "http://daum.net",
	}
	fmt.Println("map", _map["name"])
	return _map
}
func getChildren() []string {
	return []string{"jenny", "rose"}
}

// 13. 함수
func getCompany() string {
	return getMap()["company"]
}
func getName() (string, string) {
	res := strings.Split(getMap()["name"], " ")
	return res[0], res[1]
}
func getCountOfChildren() (count int) {
	count = len(getChildren())
	return
}

// 14. 구조체, 인터페이스
type Circle struct {
	x float64
	y float64
	r float64
}

type Primitive interface {
	getWidth() float64
	getHeight() float64
}
type Rectangle struct {
	width  float64
	height float64
}

func (rect Rectangle) getWidth() float64 {
	return rect.width
}
func (rect Rectangle) getHeight() float64 {
	return rect.height
}

// ===================================

func main() {
	print("Hello World")

	// 일시 정지를 위해
	var b []byte = make([]byte, 4)

	block := false
	if block {
		os.Stdin.Read(b)

		var c = make(chan int)
		go add(1, 3, c)
		print("chan c -> " + strconv.Itoa(<-c))

		// 일시 정지를 위해
		os.Stdin.Read(b)

		// 8. 포인터
		// – C와 유사함 (단, new 키워드 사용 가능)
		var c2 int
		add2(1, 3, &c2)
		Println("pointer c -> " + strconv.Itoa(c2)) // 닉네임 생략

		// 일시 정지를 위해
		os.Stdin.Read(b)

		d := new(int)
		add2(1, 4, d)
		Println("new d -> " + strconv.Itoa(*d))

		tOs.Stdin.Read(b) //os에 닉네임 tOs 부여

		get(d)
		os.Stdin.Read(b)

		add2(1, *d, d)
		get(d)
		os.Stdin.Read(b)

		add2(1, *d, d)
		get(d)
		os.Stdin.Read(b)

		// 9. 조건문은 C 스타일 (if, switch)
		// – 단, switch문이 좀 특이함
		// – switch문에 break 생략 가능
		// – case를 연속으로 여러개 사용 불가 (break를 생략할 수 있는 이유)
		// – fallthrough 키워드: 조건에 맞는 case 다음 case 무조건 실행
		print("main -> getInfo")
		getInfo()
		print("main -> rename")
		rename(getInfo())
		os.Stdin.Read(b)

		// 10. 10. 반복문은 for만 존재, C 스타일
		// – for i := 0; i < 10; i++ { } - while문 대신 아래 코드처럼 사용
		print("main -> iterate")
		iterate()
		os.Stdin.Read(b)

		// 11. 배열
		print("main -> arr")
		arr()
		os.Stdin.Read(b)

		// 12. 맵
		print("main -> getMap")
		getMap()
		os.Stdin.Read(b)

		// 13. 함수
		print("main -> func")
		print(getCompany())
		print(getName())
		print(getCountOfChildren())
		os.Stdin.Read(b)

		// 14. 구조체, 인터페이스
		// – 클래스는 없고 구조체를 이용함
		// – 인터페이스의 사용법이 특이함
		var ccr Circle = Circle{x: 1, y: 2, r: 3}
		print(ccr.r)

		ccr2 := Circle{2, 4, 8}
		print(ccr2.r)

		rect := Rectangle{100, 200}
		print(rect.getWidth())
		// print(rect.getHeight())

		p := Primitive(rect)
		print(p.getHeight())

		var p2 Primitive = rect
		print(p2.getWidth())

		os.Stdin.Read(b)
	}
}
