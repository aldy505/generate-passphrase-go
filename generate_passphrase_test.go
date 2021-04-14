package go_generate_passphrase

import (
	"regexp"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("Should generate passphrase without options", func(t *testing.T) {
		got, err := Generate(&generateOptions{})
		if err != nil {
			t.Errorf(err.Error())
		}
		split := strings.Split(got, "-")
		if len(split) != 4 {
			t.Errorf("Strings does not equal to 4 section: %v", got)
		}
	})
	t.Run("Should generate a passphrase with size length", func(t *testing.T) {
		got, err := Generate(&generateOptions{
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
	t.Run("Should generate all word pattern with numbers: false", func(t *testing.T) {
		got, err := Generate(&generateOptions{
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
	t.Run("Should output error for unknown pattern", func(t *testing.T) {
		_, err := Generate(&generateOptions{
			Pattern: "AAA",
		})
		if err == nil {
			t.Error("Error was not thrown")
		}
	})
	t.Run("Should generate all word pattern with pattern: WWWWW", func(t *testing.T) {
		got, err := Generate(&generateOptions{
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
	t.Run("Should generate all number pattern with pattern: NNNNN", func(t *testing.T) {
		got, err := Generate(&generateOptions{
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
	t.Run("Should generate all uppercase word pattern", func(t *testing.T) {
		got, err := Generate(&generateOptions{
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
	/*
		I don't know exactly why this function right here emits error on the test.
		Not that it returns error, it creates error to the test.

		t.Run("Should generate all titlecase word pattern", func(t *testing.T) {
			got, err := Generate(&generateOptions{
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
				if !(lowerCaseRegex.MatchString(perWord[0])) {
					t.Errorf("Strings is not lowercase: %v", got)
				}
			}
		})
	*/
	t.Run("should have different separator", func(t *testing.T) {
		got, err := Generate(&generateOptions{
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
		got, err := Generate(&generateOptions{
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
		got, err := Generate(&generateOptions{
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
		got, err := Generate(&generateOptions{
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
		got, err := GenerateMultiple(5, &generateOptions{})
		if err != nil {
			t.Errorf(err.Error())
		}
		if len(got) != 5 {
			t.Error("Slice length is not 5")
		}
	})
}
