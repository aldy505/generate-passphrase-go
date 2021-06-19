package generatepassphrase_test

import (
	"regexp"
	"strings"
	"testing"

	passphrase "github.com/aldy505/generate-passphrase-go"
)

func TestGenerate(t *testing.T) {
	t.Run("should generate passphrase without options", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{})
		if err != nil {
			t.Errorf(err.Error())
		}
		t.Log(got)
		split := strings.Split(got, "-")
		if len(split) != 4 {
			t.Errorf("Strings does not equal to 4 section: %v", got)
		}
	})
	t.Run("should generate a passphrase with size length", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Length: 10,
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		if len(split) != 10 {
			t.Errorf("Strings does not equal to 10 section: %v", got)
		}
	})
	t.Run("should generate all word pattern with numbers: false", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Numbers: false,
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		validRegex := regexp.MustCompile(`[a-zA-Z]`)
		if !(validRegex.MatchString(got)) {
			t.Errorf("Strings has numbers: %v", got)
		}
	})
	t.Run("should output error for unknown pattern", func(t *testing.T) {
		_, err := passphrase.Generate(passphrase.Options{
			Pattern: "AAA",
		})
		if err == nil {
			t.Error("Error was not thrown")
		}
	})
	t.Run("should generate all word pattern with pattern: WWWWW", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Pattern: "WWWWW",
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		if len(split) != 5 {
			t.Errorf("Strings does not equal to 5 section: %v", got)
		}
		for i := 0; i < len(split); i++ {
			validRegex := regexp.MustCompile(`[a-zA-Z]`)
			if !(validRegex.MatchString(got)) {
				t.Errorf("Strings has numbers: %v", got)
			}
		}
	})
	t.Run("should generate all number pattern with pattern: NNNNN", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Pattern: "NNNNN",
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		if len(split) != 5 {
			t.Errorf("Strings does not equal to 5 section: %v", got)
		}
		for i := 0; i < len(split); i++ {
			validRegex := regexp.MustCompile(`[0-9]`)
			if !(validRegex.MatchString(got)) {
				t.Errorf("Strings has words: %v", got)
			}
		}
	})
	t.Run("should generate all uppercase word pattern", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Numbers:   false,
			Uppercase: true,
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		for i := 0; i < len(split); i++ {
			validRegex := regexp.MustCompile(`[A-Z]`)
			if !(validRegex.MatchString(got)) {
				t.Errorf("Strings is not uppercase: %v", got)
			}
		}
	})
	t.Run("should generate all titlecase word pattern", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Numbers:   false,
			Titlecase: true,
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		for i := 0; i < len(split); i++ {
			perWord := strings.Split(split[i], "")
			upperCaseRegex := regexp.MustCompile(`[A-Z]`)
			lowerCaseRegex := regexp.MustCompile(`[a-z]`)
			if !(upperCaseRegex.MatchString(perWord[0])) {
				t.Errorf("Strings is not uppercase: %v", got)
			}
			if !(lowerCaseRegex.MatchString(perWord[1])) {
				t.Errorf("Strings is not lowercase: %v", got)
			}
		}
	})
	t.Run("should have different separator", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Separator: "_",
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		separatorRegex, err := regexp.MatchString("_.", got)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !separatorRegex {
			t.Errorf("Separator doesn't exists: %v", got)
		}
	})
	t.Run("should use pattern if length is also provided", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Length:  10,
			Pattern: "WWNWWW",
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		if len(split) != 6 {
			t.Errorf("Strings is not equal to 6 sections: %v", got)
		}
	})
	t.Run("should still be uppercase if titlecase is also true", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Uppercase: true,
			Titlecase: true,
			Numbers:   false,
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		for i := 0; i < len(split); i++ {
			validRegex := regexp.MustCompile(`[A-Z]`)
			if !(validRegex.MatchString(got)) {
				t.Errorf("Strings is not uppercase: %v", got)
			}
		}
	})
	t.Run("should have all uppercase words and numbers", func(t *testing.T) {
		got, err := passphrase.Generate(passphrase.Options{
			Uppercase: true,
			Titlecase: true,
			Numbers:   true,
			Pattern:   "WWWNWWNWWN",
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		for i := 0; i < len(split); i++ {
			validRegex := regexp.MustCompile(`[0-9A-Z]`)
			if !(validRegex.MatchString(got)) {
				t.Errorf("Strings is not uppercase: %v", got)
			}
		}
	})
}

func TestGenerateMultiple(t *testing.T) {
	t.Run("should generate 5 multiple passphrase without options", func(t *testing.T) {
		got, err := passphrase.GenerateMultiple(5, passphrase.Options{})
		if err != nil {
			t.Errorf(err.Error())
		}
		if len(got) != 5 {
			t.Error("Slice length is not 5")
		}
	})
	t.Run("should generate 25 multiple passphrase without options", func(t *testing.T) {
		got, err := passphrase.GenerateMultiple(25, passphrase.Options{})
		if err != nil {
			t.Errorf(err.Error())
		}
		if len(got) != 25 {
			t.Error("Slice length is not 25")
		}
	})
	t.Run("should generate 10 multiple passphrase with size length", func(t *testing.T) {
		got, err := passphrase.GenerateMultiple(10, passphrase.Options{
			Length: 10,
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		if len(got) != 10 {
			t.Error("Slice length is not 10")
		} else {
			for i := 0; i < len(got); i++ {
				split := strings.Split(got[i], "-")
				if len(split) != 10 {
					t.Errorf("Strings does not equal to 10 section: %v", got[i])
				}
			}
		}
	})
	t.Run("should generate 10 multiple passphrase with all word pattern with numbers: false", func(t *testing.T) {
		got, err := passphrase.GenerateMultiple(10, passphrase.Options{
			Numbers: false,
		})
		if err != nil {
			t.Errorf(err.Error())
		}
		if len(got) != 10 {
			t.Error("Slice length is not 10")
		} else {
			for i := 0; i < len(got); i++ {
				validRegex := regexp.MustCompile(`[a-zA-Z]`)
				if !(validRegex.MatchString(got[i])) {
					t.Errorf("Strings has numbers: %v", got[i])
				}
			}
		}
	})
	t.Run("should output error for unknown pattern on multiple passphrase", func(t *testing.T) {
		_, err := passphrase.GenerateMultiple(10, passphrase.Options{
			Pattern: "AAA",
		})
		if err == nil {
			t.Error("Error was not thrown")
		}
	})
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		passphrase.Generate(passphrase.Options{})
	}
}

func BenchmarkMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		passphrase.GenerateMultiple(10, passphrase.Options{})
	}
}
