package go_generate_passphrase

import (
	"crypto/rand"
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Options struct {
	Length    int
	Separator string
	Numbers   bool
	Uppercase bool
	Titlecase bool
	Pattern   string
}

var randomBytes []byte
var randomIndex int

func getRandomValue() int {
	if randomIndex >= 0 || randomIndex >= len(randomBytes) {
		randomIndex = 0
		randomBytes = make([]byte, 256)
		rand.Read(randomBytes)
	}
	randomIndex += 1
	return int(randomBytes[randomIndex])
}

func getRandomNumber(max int) int {
	var rand int
	for {
		rand = getRandomValue()
		if !(rand >= 256-(256%max)) {
			break
		}
	}
	return rand % max
}

func getRandomWord() (string, error) {
	vendorPath, err := filepath.Abs("./vendor/github.com/aldy505/go-generate-passphrase/words.txt")
	if err != nil {
		return "", err
	}
	gopathPath := filepath.Join(os.Getenv("GOPATH"), "/pkg/mod/github.com/aldy505/go-generate-passphrase@v0.0.1")
	words, err := ioutil.ReadFile(gopathPath)
	if err != nil {
		words, err = ioutil.ReadFile(vendorPath)
		if err != nil {
			words, err = ioutil.ReadFile("./words.txt")
			if err != nil {
				return "", err
			}
		}
	}
	wordsString := string(words)
	wordsArray := strings.Split(wordsString, "\n")
	randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordsArray))))
	if err != nil {
		return "", err
	}
	return wordsArray[randomInt.Int64()], nil
}

func getRandomPattern(length int, numbers bool) string {
	var pool []string
	var pattern strings.Builder
	if numbers {
		pool = append(pool, "N", "W", "W")
	} else {
		pool = append(pool, "W", "W", "W")
	}
	for i := 0; i < length; i++ {
		pattern.WriteString(pool[getRandomNumber(2)])
	}
	return pattern.String()
}

func Generate(options *Options) (string, error) {
	var length int
	if options.Length <= 0 {
		length = 4
	} else {
		length = options.Length
	}

	var numbers bool
	if options.Numbers {
		numbers = true
	} else {
		numbers = false
	}

	var passphraseArray []string
	var pattern string

	if len(options.Pattern) > 0 {
		pattern = strings.ToUpper(options.Pattern)
	} else {
		pattern = getRandomPattern(length, numbers)
	}

	eachPattern := strings.Split(pattern, "")

	for i := 0; i < len(eachPattern); i++ {
		if eachPattern[i] == "N" {
			passphraseArray = append(passphraseArray, strconv.Itoa(getRandomValue()))
		} else if eachPattern[i] == "W" {
			word, err := getRandomWord()
			if err != nil {
				return "", errors.New("Failed to get a word.")
			}
			if options.Uppercase {
				passphraseArray = append(passphraseArray, strings.ToUpper(word))
			} else if options.Titlecase {
				passphraseArray = append(passphraseArray, strings.Title(strings.ToLower(word)))
			} else {
				passphraseArray = append(passphraseArray, word)
			}
		} else {
			return "", errors.New("Unknown pattern found. Use N or W instead.")
		}
	}

	var separator string
	if len(options.Separator) > 0 {
		separator = options.Separator
	} else {
		separator = "-"
	}

	passphrase := strings.Join(passphraseArray, separator)
	return passphrase, nil
}

func GenerateMultiple(amount int, options *Options) ([]string, error) {
	var passphrase []string
	for i := 0; i < amount; i++ {
		generated, err := Generate(options)
		if err != nil {
			return passphrase, err
		}
		passphrase = append(passphrase, generated)
	}
	return passphrase, nil
}
