package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 09")
}

func checkErr(e error) {
	if e != nil {
		log.Panicf("err: %+v", e)
	}
}

func main() {
	Practice1()
	Practice2()
	Practice3()
}

func Practice1() {
	// https://gobyexample.com/writing-files
	data1 := []byte("hello\nworld\n")
	errWrite := ioutil.WriteFile("data1", data1, 0644)
	checkErr(errWrite)

	file1, errCreate := os.Create("data2")
	checkErr(errCreate)
	defer file1.Close()

	data2 := []byte{115, 111, 109, 101, 10}
	number1, errWrite := file1.Write(data2)
	checkErr(errWrite)
	log.Printf("wrote %d bytes\n", number1)

	number2, errWrite := file1.WriteString("writes\n")
	checkErr(errWrite)
	log.Printf("wrote %d bytes\n", number2)

	file1.Sync()

	w := bufio.NewWriter(file1)
	number3, errWrite := w.WriteString("buffered\n")
	checkErr(errWrite)
	log.Printf("wrote %d bytes\n", number3)

	w.Flush()
}

func Practice2() {
	// https://blog.csdn.net/flyingshineangel/article/details/62889256
	bytes, errRead := ioutil.ReadFile("data2")
	checkErr(errRead)
	log.Println(string(bytes))

	file, errOpen := os.Open("data2")
	checkErr(errOpen)
	defer file.Close()
	bytes, errRead = ioutil.ReadAll(file)
	checkErr(errRead)
	log.Println(string(bytes))

	file, errOpen = os.Open("data2")
	checkErr(errOpen)
	defer file.Close()
	buffer := make([]byte, 9)
	_, errRead = file.Read(buffer)
	checkErr(errRead)
	log.Println(string(buffer))

	file, errOpen = os.Open("data2")
	checkErr(errOpen)
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer = make([]byte, 9)
	_, errRead = reader.Read(buffer)
	checkErr(errRead)
	log.Println(string(buffer))
}

func Practice3() {
	// https://segmentfault.com/a/1190000015591319
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	for _, p := range proverbs {
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			log.Println("failed to write data")
			os.Exit(1)
		}
	}

	// file scanner
	file, errOpen := os.Open("data2")
	checkErr(errOpen)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		input := fileScanner.Text()
		log.Printf("got: %s", input)
	}
	if errScan := fileScanner.Err(); errScan != nil {
		log.Println("scanner err: %+v", errScan)
	}
}
