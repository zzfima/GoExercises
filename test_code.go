package main

import "fmt"

func main() {
	//vars
	var age = 34
	var weight int = 85
	fmt.Println("Hello")
	fmt.Println("My age is ", age, ", my weight is ", weight)

	//if
	if age < 16 {
		fmt.Println("Go to school")
	} else {
		fmt.Println("Go to army")
	}

	//switch
	switch age {
	case 11:
		fmt.Println("you are eleven")
	case 22:
		fmt.Println("you are twenty two")
	default:
		fmt.Println("...strange age...")
	}

	//loops
	var i = 0
	for i < age {
		i += 10
		fmt.Println("your age is ", i)
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	//array
	var arr1 [2]int
	arr1[0] = 1
	arr1[1] = 11

	arr2 := []float64{1.11, 2.22, 3.33}
	fmt.Println(arr2[1])

	//foreach
	for i2, val2 := range arr2 {
		fmt.Println("index: ", i2, ", value: ", val2)
	}

	//hash table
	hash1 := make(map[string]int)
	hash1["baa"] = 15
	hash1["aaa"] = 25
	hash1["caa"] = 35
	for i3, val3 := range hash1 {
		fmt.Println("index: ", i3, ", value: ", val3)
	}

	//func calls
	fmt.Println(f1(4))
	v1, v2 := f2(5)
	fmt.Println(v1, v2)

	//closure
	v3 := func() int {
		return 33 + v1
	}
	fmt.Println(v3())

	//defer
	defer two()
	one()
	one()

	//func pointer
	var x1 = 121
	plus_plus(x1)
	fmt.Println(x1)

	plus_plus_by_ref(&x1)
	fmt.Println(x1)

	//objects
	alex := Human{"Alex Babushkin", 3245435}
	fmt.Println(alex.id)

	//class func
	alex.print_name()

	//interface
	var h2 IMakeVoice = &alex
	h2.do_voice()
}

func f1(v int) int {
	return v + 1
}

func f2(v int) (int, int) {
	return v + 1, v + 2
}

func one() {
	fmt.Println("one")
}

func two() {
	fmt.Println("two")
}

func plus_plus(x int) {
	x++
}

func plus_plus_by_ref(x *int) {
	*x++
}

type Human struct {
	name string
	id   int64
}

func (h *Human) print_name() {
	fmt.Println(h.name)
}

type IMakeVoice interface {
	do_voice()
}

func (h *Human) do_voice() {
	fmt.Println(h.name, " screams")
}
