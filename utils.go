package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"syscall"

	"github.com/sqweek/dialog"
)

var (
	source                 = rand.NewSource(time.Now().UnixNano())
	random                 = rand.New(source)
	today           string = time.Now().Format("Mon Jan 2")
	editedcombopath string
)

func init() {
	os.MkdirAll("Results/"+today, 0777)
}

func GenFileName(name string) (*os.File, error) {
	editedcombopath = fmt.Sprintf("%s/%s %s.txt", "Results/"+today, name, strings.ReplaceAll(time.Now().Format("15:04:05"), ":", "-"))
	_, err := os.Create(editedcombopath)
	if err != nil {
		fmt.Println(err)
	}
	return os.Create(editedcombopath)
}

func DialogBox() string {
	filepath, err := dialog.File().Title("Load File").Load()
	if err != nil {
		fmt.Println("There was an error while loading file. Try again")
	}
	return filepath

}

func LoadFile() []string {
	var lines []string
	filepath := DialogBox()
	bucket := make(map[string]bool)
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		final := strings.TrimSpace(scanner.Text())
		if _, ok := bucket[final]; !ok {
			bucket[final] = true
			lines = append(lines, final)
		}
	}
	return lines
}

func WriteFile(filepath string, dataSlice []string) {
	file, _ := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()
	for _, data := range dataSlice {
		file.WriteString(data + "\n")
	}
}

func WriteToFile(file *os.File, data string) {
	file.WriteString(data + "\n")
}

func Title(title string) (int, error) {
	var kernel32 = syscall.NewLazyDLL("kernel32.dll")
	var procSetConsoleTitle = kernel32.NewProc("SetConsoleTitleW")

	ptr, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return 0, err
	}

	_, _, err = procSetConsoleTitle.Call(uintptr(unsafe.Pointer(ptr)))
	if err != nil && err.Error() != "The operation completed successfully." {
		return 0, err
	}

	return 1, nil
}
