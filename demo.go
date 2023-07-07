package main

import "fmt"

func main() {
	// Mảng
	var a [5]int //int array with length 5
	fmt.Println(a)
	b := [3]int{90, 92, 93} // short hand declaration to create array
	fmt.Println(b)
	c := [3]int{0: 9, 2: 12}
	fmt.Println(c)
	d := [...]int{9, 10, 11} // ... makes the compiler determine the length
	fmt.Println(d)
	fmt.Println("---------------------------------------------------")
	// Mảng trong Go không phải loại tham chiếu
	e := [...]string{"USA", "China", "India", "Germany", "France"}
	f := e // a copy of a is assigned to b
	f[0] = "Singapore"
	fmt.Println("e is ", e)
	fmt.Println("f is ", f)
	fmt.Println("---------------------------------------------------")
	// Độ dài của mảng
	fmt.Println("length of e is", len(e))
	fmt.Println("---------------------------------------------------")
	for i := 0; i < len(e); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of e is %s\n", i, e[i])
	}
	fmt.Println("---------------------------------------------------")
	sum := 0
	for i, v := range b { //range returns both the index and value
		fmt.Printf("%d the element of a is %d\n", i, v)
		sum += v
	}
	fmt.Println("sum of all elements of a", sum)
	fmt.Println("---------------------------------------------------")

	// Mảng đa chiều
	g := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	for _, v1 := range g {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
	fmt.Println("---------------------------------------------------")

	// Slice
	// Khai báo slice
	sliceDemo := make([]int, 3, 5) // Tạo 1 slice có length = 3 và capacity bằng 5
	fmt.Println(sliceDemo, "slice demo")
	fmt.Println("---------------------------------------------------")

	h := [5]string{"Son", "Vuong", "Chuan", "Yen", "Ha"}
	var i []string = h[1:4] // Tạo 1 slice tham chiếu từ mảng h bắt đầu từ phần tử vị trí 1 đến vị trí 3
	fmt.Println(i, "before")
	h[3] = "Thu"
	fmt.Println(i, "after") // Sau khi gán lại phần tử ở mảng h thì slice thay đổi theo
	fmt.Println("---------------------------------------------------")

	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numslice := numbers[2:5]
	fmt.Println("array before", numbers)
	for i := range numslice { // Ngược lại, sau khi thay đổi giá trị phần tử slice thì mảng gốc thay đổi theo
		numslice[i]++
	}
	fmt.Println("array after", numbers)
	fmt.Println("---------------------------------------------------")

	// Capacity của mảng
	animals := [...]string{"Dog", "Cat", "Chicken", "Duck", "Bird", "Tiger", "Elephant"}
	animalslice := animals[1:3]
	fmt.Println(animalslice, "before")
	fmt.Printf("length of slice %d capacity %d\n", len(animalslice), cap(animalslice)) //length of is 2 and capacity is 5
	animalslice = animalslice[:cap(animalslice)]                                       //re-slicing animalslice till its capacity
	fmt.Println(animalslice, "after")
	fmt.Println("After re-slicing length is", len(animalslice), "and capacity is", cap(animalslice))
	fmt.Println("---------------------------------------------------")

	// Thêm phần tử vào 1 slice
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	fmt.Println("---------------------------------------------------")

	var names []string //zero value of a slice is nil
	fmt.Println("length nil:", len(names), "cap nil:", cap(names))
	fmt.Println("---------------------------------------------------")

	// Truyền slice dưới dạng tham số qua 1 hàm
	numTest := []int{5, 10, 15, 20}
	fmt.Println("slice before function call", numTest)
	doubleUp(numTest)
	fmt.Println("slice after function call", numTest) // Giá trị của các phần tử trong slice đã thay đổi
	fmt.Println("---------------------------------------------------")

	// Copy slice
	// thêm phần tử 0 vào cuối slice a
	var k = []int{1, 2, 3}
	k = append(k, 0)
	// lùi những phần tử từ i trở về sau
	copy(k[2:], k[1:])
	fmt.Println(k, "before")
	// gán vị trí thứ i bằng x
	k[1] = 10
	fmt.Println(k, "after")
	fmt.Println("---------------------------------------------------")
	var l = k
	fmt.Println(l, "l before")
	l = append(l, k...)
	copy(l[2+len(k):], l[2:])
	copy(l[2:], k)
	fmt.Println(l, "insert by length")
	fmt.Println("---------------------------------------------------")

	// Xóa phần tử
	m := []int{1, 2, 3, 4, 5, 6}
	// xóa một phần tử ở cuối
	m = m[:len(m)-1]
	fmt.Println(m, "after deleting last element")
	n := []int{1, 2, 3, 4, 5, 6}
	// xóa N phần tử ở cuối
	n = n[:len(n)-3]
	fmt.Println(n, "after deleting 3 element from the last")
	n = []int{1, 2, 3, 4, 5, 6}
	// xóa phần tử ở vị trí 2
	n = append(n[:2], n[2+1:]...)
	fmt.Println(n, "after deleting element at index 2")
	n = []int{1, 2, 3, 4, 5, 6}
	// xóa 3 phần tử từ vị trí 2
	n = append(n[:2], n[2+3:]...)
	fmt.Println(n, "after deleting 3 element start at index 2")
	n = []int{1, 2, 3, 4, 5, 6}
	n = n[:3+copy(n[3:], n[4:])]
	fmt.Println(n, "after deleting element at index 3")
	o := make([]int, 4, 6)
	fmt.Println(len(o[4:]), "o1:")
	fmt.Println(len(o[2:]), "o2:")
	fmt.Println(copy(o[4:], o[1:]))
	fmt.Println("---------------------------------------------------")
	// Map
	// Khai báo map
	var country map[int]string
	country = make(map[int]string)
	country[1] = "India"
	country[2] = "China"
	country[3] = "Pakistan"
	country[4] = "Germany"
	country[5] = "Australia"
	country[6] = "Indonesia"
	for i, j := range country {
		fmt.Printf("Key: %d Value: %s\n", i, j)
	}
	// Độ dài map sễ bằng số phần tử được khai báo
	fmt.Println(len(country), "<== map length")
	fmt.Println("---------------------------------------------------")
	// Kiểm tra sự tồn tại của cặp key-value
	value, exists := country[1]
	if exists {
		fmt.Println(exists)
		fmt.Println(value)
	}
	fmt.Println("---------------------------------------------------")
	// Sử dụng range với map
	dictionary := map[string]string{
		"Sky":    "Bầu trời",
		"Land":   "Mặt đất",
		"Hello":  "Xin chào",
		"Bye":    "Tạm biệt",
		"Thanks": "Cảm ơn",
	}
	for key, value := range dictionary {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
	fmt.Println("---------------------------------------------------")
	// Xóa 1 cặp key-value
	delete(dictionary, "Bye")
	// Truyền mảng vào phương thức
	removeDic(dictionary, "Thanks")
	for key, value := range dictionary {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
	fmt.Println("---------------------------------------------------")
}
func doubleUp(numbers []int) {
	for i := range numbers {
		numbers[i] *= 2
	}
}
func removeDic(dic map[string]string, key string) {
	delete(dic, key)
}
