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
		break
	case 1:
		wordsArray = words1
		break
	case 2:
		wordsArray = words2
		break
	case 3:
		wordsArray = words3
		break
	case 4:
		wordsArray = words4
		break
	case 5:
		wordsArray = words5
		break
	case 6:
		wordsArray = words6
		break
	case 7:
		wordsArray = words7
		break
	case 8:
		wordsArray = words8
		break
	case 9:
		wordsArray = words9
		break
	case 10:
		wordsArray = words10
		break
	case 11:
		wordsArray = words11
		break
	case 12:
		wordsArray = words12
		break
	case 13:
		wordsArray = words13
		break
	case 14:
		wordsArray = words14
		break
	case 15:
		wordsArray = words15
		break
	case 16:
		wordsArray = words16
		break
	case 17:
		wordsArray = words17
		break
	case 18:
		wordsArray = words18
		break
	case 19:
		wordsArray = words19
		break
	case 20:
		wordsArray = words20
		break
	case 21:
		wordsArray = words21
		break
	case 22:
		wordsArray = words22
		break
	case 23:
		wordsArray = words23
		break
	case 24:
		wordsArray = words24
		break
	case 25:
		wordsArray = words25
		break
	case 26:
		wordsArray = words26
		break
	case 27:
		wordsArray = words27
		break
	case 28:
		wordsArray = words28
		break
	case 29:
		wordsArray = words29
		break
	case 30:
		wordsArray = words30
		break
	case 31:
		wordsArray = words31
		break
	case 32:
		wordsArray = words32
		break
	case 33:
		wordsArray = words33
		break
	case 34:
		wordsArray = words34
		break
	case 35:
		wordsArray = words35
		break
	case 36:
		wordsArray = words36
		break
	case 37:
		wordsArray = words37
		break
	case 38:
		wordsArray = words38
		break
	case 39:
		wordsArray = words39
		break
	case 40:
		wordsArray = words40
		break
	case 41:
		wordsArray = words41
		break
	case 42:
		wordsArray = words42
		break
	case 43:
		wordsArray = words43
		break
	case 44:
		wordsArray = words44
		break
	case 45:
		wordsArray = words45
		break
	case 46:
		wordsArray = words46
		break
	case 47:
		wordsArray = words47
		break
	case 48:
		wordsArray = words48
		break
	case 49:
		wordsArray = words49
		break
	case 50:
		wordsArray = words50
		break
	case 51:
		wordsArray = words51
		break
	case 52:
		wordsArray = words52
		break
	case 53:
		wordsArray = words53
		break
	case 54:
		wordsArray = words54
		break
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
