package main

import ("fmt"
	"strings"
	"os"
	"io/ioutil"
	"path/filepath")

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
	
	file, err := os.Open("test.txt")
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
		fmt.Println("Error on reading")
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	} 	

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
