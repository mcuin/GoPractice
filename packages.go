package main

import ("fmt"
	"strings"
	//"os"
	//"io/ioutil"
	//"path/filepath"
	"container/list"
	"sort")

type Person struct {
		Name string
		Age int
	}

type ByName []Person
type ByAge []Person

func(this ByName) Len() int {
	return len(this)
}

func(this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func(this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func(this ByAge) Len() int {
	return len(this)
}

func(this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func(this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaa", "a", "b", 2),
		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
	)
	
	/*file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error in open")
		return
	}
	defer file.Close();

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	bs2, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	str2 := string(bs2)
	fmt.Println(str2)

	file2, err := os.Create("test2.text")
	if err != nil {
		fmt.Println("Error in create")
		return
	}

	defer file2.Close()
	
	file2.WriteString("test")

	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Error directory open")
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		err := errors.New("error message")
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	} 	

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})*/

	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	kids := []Person {
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids) 
	sort.Sort(ByAge(kids))
	fmt.Println(kids)

	
}
