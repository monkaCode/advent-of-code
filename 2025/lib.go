package lib

import (
	"bufio"
	"log"
	"os"
)

func Read(path string) []string {
	var array []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return array
}

func Write(path string, lines []string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		_, err = w.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
}
