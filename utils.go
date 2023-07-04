package utils

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/sqweek/dialog"
)

var (
	source                 = rand.NewSource(time.Now().UnixNano())
	random                 = rand.New(source)
	symbols         string = `"!@#$%^&*()-+?_=,<>/"`
	uppercase       string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase       string = "abcdefghijklmnopqrstuvwxyz"
	digits          string = "0123456789"
	today           string = time.Now().Format("Mon Jan 2")
	editedcombopath string
)

func init() {
	os.MkdirAll("Results/"+today, 0777)
}

func ExtractCombo(combo string) (string, string) {
	return strings.Split(combo, ":")[0], strings.Split(combo, ":")[1]
}

func ExtractLeftRight(main, left, right string) string {
	return strings.Split((strings.Split(main, left)[1]), right)[0]
}

func ExtractStringIndex(main string, split string, index int) string {
	return strings.Split(main, split)[index]
}
func Length(s string) int {
	return len(s)
}

func RandomReturn() bool {
	return random.Intn(10) > random.Intn(10)
}

func RandomItem(slice []string) string {
	return slice[random.Intn(len(slice))]
}

func CheckString(main, character string) bool {
	return strings.Contains(main, character)
}

func CheckSymbols(password string) bool {
	return strings.ContainsAny(password, symbols)
}

func CheckUpper(password string) bool {
	return strings.ContainsAny(password, uppercase)
}

func CheckLower(password string) bool {
	return strings.ContainsAny(password, lowercase)
}

func CheckDigits(password string) bool {
	return strings.ContainsAny(password, digits)
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
