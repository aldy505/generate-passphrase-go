package generatepassphrase

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
	"strings"
)

// Options for passphrase instance
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
	randomIndex++
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
	var wordsArray [5000]string
	useWords := getRandomNumber(54)

	switch useWords {
	case 0:
		wordsArray = words0

	case 1:
		wordsArray = words1
	case 2:
		wordsArray = words2
	case 3:
		wordsArray = words3
	case 4:
		wordsArray = words4
	case 5:
		wordsArray = words5
	case 6:
		wordsArray = words6
	case 7:
		wordsArray = words7
	case 8:
		wordsArray = words8
	case 9:
		wordsArray = words9
	case 10:
		wordsArray = words10
	case 11:
		wordsArray = words11
	case 12:
		wordsArray = words12
	case 13:
		wordsArray = words13
	case 14:
		wordsArray = words14
	case 15:
		wordsArray = words15
	case 16:
		wordsArray = words16
	case 17:
		wordsArray = words17
	case 18:
		wordsArray = words18
	case 19:
		wordsArray = words19
	case 20:
		wordsArray = words20
	case 21:
		wordsArray = words21
	case 22:
		wordsArray = words22
	case 23:
		wordsArray = words23
	case 24:
		wordsArray = words24
	case 25:
		wordsArray = words25
	case 26:
		wordsArray = words26
	case 27:
		wordsArray = words27
	case 28:
		wordsArray = words28
	case 29:
		wordsArray = words29
	case 30:
		wordsArray = words30
	case 31:
		wordsArray = words31
	case 32:
		wordsArray = words32
	case 33:
		wordsArray = words33
	case 34:
		wordsArray = words34
	case 35:
		wordsArray = words35
	case 36:
		wordsArray = words36
	case 37:
		wordsArray = words37
	case 38:
		wordsArray = words38
	case 39:
		wordsArray = words39
	case 40:
		wordsArray = words40
	case 41:
		wordsArray = words41
	case 42:
		wordsArray = words42
	case 43:
		wordsArray = words43
	case 44:
		wordsArray = words44
	case 45:
		wordsArray = words45
	case 46:
		wordsArray = words46
	case 47:
		wordsArray = words47
	case 48:
		wordsArray = words48
	case 49:
		wordsArray = words49
	case 50:
		wordsArray = words50
	case 51:
		wordsArray = words51
	case 52:
		wordsArray = words52
	case 53:
		wordsArray = words53
	case 54:
		wordsArray = words54
	}

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

// Generate generates 1 random passphrase based on options provided
func Generate(options Options) (string, error) {
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
				return "", errors.New("failed to get a word")
			}
			if options.Uppercase {
				passphraseArray = append(passphraseArray, strings.ToUpper(word))
			} else if options.Titlecase {
				passphraseArray = append(passphraseArray, strings.Title(strings.ToLower(word)))
			} else {
				passphraseArray = append(passphraseArray, word)
			}
		} else {
			return "", errors.New("unknown pattern found. Use N or W instead")
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

// Generate generates multiple random passphrase based on options provided
func GenerateMultiple(amount int, options Options) ([]string, error) {
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
