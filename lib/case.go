package lib

import (
	"reflect"
	"regexp"
	"strings"
	"unicode"

	changecase "github.com/ku/go-change-case"
)

type ChangeCaseFn func(string) string

type ChangeCaser interface {
	Fn(string) string
	Subtitle() string
}

type Case struct {
	fn       ChangeCaseFn
	subtitle string
}

func (c *Case) Fn(str string) string {
	return c.fn(str)
}

func (c *Case) Subtitle() string {
	return c.subtitle
}

func NewCase(fn ChangeCaseFn, sub string) *Case {
	return &Case{
		fn:       fn,
		subtitle: sub,
	}
}

type ChangeCase struct {
	Camel    ChangeCaser
	Pascal   ChangeCaser
	Constant ChangeCaser
	Dot      ChangeCaser
	Lower    ChangeCaser
	Lcfirst  ChangeCaser
	No       ChangeCaser
	Param    ChangeCaser
	Path     ChangeCaser
	Sentence ChangeCaser
	Snake    ChangeCaser
	Swap     ChangeCaser
	Title    ChangeCaser
	Upper    ChangeCaser
	Ucfirst  ChangeCaser
	Hashtag  ChangeCaser
}

func NewChangeCase() *ChangeCase {
	return &ChangeCase{
		Camel:    NewCase(Camel, "Convert to a string with the separators denoted by having the next letter capitalized"),
		Pascal:   NewCase(Pascal, "Convert to a string denoted in the same fashion as camelCase, but with the first letter also capitalized"),
		Constant: NewCase(changecase.Constant, "Convert to an upper case, underscore separated string"),
		Dot:      NewCase(changecase.Dot, "Convert to a lower case, period separated string"),
		Lower:    NewCase(changecase.Lower, "Convert to a string in lower case"),
		Lcfirst:  NewCase(changecase.LcFirst, "Convert to a string with the first character lower cased"),
		No:       NewCase(changecase.No, "Convert the string without any casing (lower case, space separated)"),
		Param:    NewCase(changecase.Param, "Convert to a lower case, dash separated string"),
		Path:     NewCase(changecase.Path, "Convert to a lower case, slash separated string"),
		Sentence: NewCase(Sentence, "Convert to a lower case, space separated string with the first character upper case"),
		Snake:    NewCase(changecase.Snake, "Convert to a lower case, underscore separated string"),
		Swap:     NewCase(changecase.Swap, "Convert to a string with every character case reversed"),
		Title:    NewCase(changecase.Title, "Convert to a space separated string with the first character of every word upper case"),
		Upper:    NewCase(changecase.Upper, "Convert to a string in upper case"),
		Ucfirst:  NewCase(changecase.UcFirst, "Convert to a string with the first character upper cased"),
		Hashtag:  NewCase(Hashtag, "Convert to a string, space separated string with hashtag symbols"),
	}
}

// Get returns the ChangeCaser for the given input name
func (c *ChangeCase) Get(input string) ChangeCaser {
	val := reflect.ValueOf(*c)
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		if strings.ToLower(name) == strings.ToLower(input) {
			return val.Field(i).Interface().(ChangeCaser)
		}
	}
	return nil
}

// Camel converts a string to a string with the separators denoted by having the next letter capitalized
func Camel(s string) string {
	words := strings.Fields(s) // Split string by whitespace
	if len(words) == 0 {
		return ""
	}

	// Convert the first word to lowercase
	words[0] = strings.ToLower(words[0])

	// Capitalize the rest of the words
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}

	return strings.Join(words, "")
}

// Pascal converts a string to pascal case
func Pascal(s string) string {
	words := strings.Fields(s) // Split string by whitespace
	if len(words) == 0 {
		return ""
	}

	// Convert the first word to lowercase
	words[0] = strings.Title(strings.ToLower(words[0]))

	// Capitalize the rest of the words
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}

	return strings.Join(words, "")
}

// Hashtag converts a string to a space separated string with hashtag symbols
func Hashtag(s string) string {
	array := regexp.MustCompile(" +").Split(s, -1)
	for i := 0; i < len(array); i++ {
		if len(array[i]) > 0 && array[i][0:1] != "#" {
			array[i] = "#" + strings.ToLower(array[i])
		}
	}
	return strings.Join(array, " ")
}

// Sentence converts a string to a lower case, space separated string with the first character upper case
func Sentence(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}
