package main

import (
	"unicode"
)

var reservedCharacters = map[rune]bool{
	';': true,
	':': true,
	'"': true,
	'`': true,
}

func ParseAndExecute() {
	for scanner.Current() != -1 {
		next := ParseInner()
		if next == nil {
			return
		}
		print(next.String())

		next.Run()
	}
}

func ParseInner() Definition {
	nextChar := scanner.SkipWhitespace()

	if nextChar == -1 {
		return nil
	}

	switch {
	case nextChar == ':':
		return ParseWord()
	case nextChar == '"':
		return ParseString()
	case nextChar == '`':
		return ParseScript()
	case nextChar == '$':
		return ParseEnvironmentVariable()
	default:
		return ParseReference()
	}
}

func ParseWord() Definition {
	scanner.Move() // Move past :

	// get key
	for (scanner.Current() != -1) && !unicode.IsSpace(scanner.Current()) {
		stringBuilder.WriteRune(scanner.Current())
		scanner.Move()
	}
	scanner.Move()

	word := stringBuilder.String()
	stringBuilder.Reset()

	parts := make([]Definition, 0)

	// TODO: handle CRLF
	for scanner.Current() != -1 && !(scanner.Current() == '\n' && scanner.Peek() == '\n') {
		parsed := ParseInner()
		if parsed != nil {
			parts = append(parts, parsed)
		}
	}
	if scanner.Current() == ';' {
		scanner.Move()
	}

	return Word{
		Key:         word,
		Definitions: parts,
	}
}

func ParseString() Definition {
	scanner.Move() // move past "

	for scanner.Current() != -1 && (scanner.Current() != '"' || scanner.Recall() == '\\') {
		stringBuilder.WriteRune(scanner.Current())
		scanner.Move()
	}
	scanner.Move() // Move past last "

	result := String{
		value: stringBuilder.String(),
	}
	stringBuilder.Reset()

	return result
}

func ParseScript() Definition {
	scanner.Move()

	// TODO: format should be a split of strings and the arguments associated
	parts := make([]string, 0)
	for scanner.Current() != -1 && (scanner.Current() != '`' || scanner.Recall() == '\\') {
		if scanner.Current() == '+' {
			parts = append(parts, stringBuilder.String())
			stringBuilder.Reset()

			stringBuilder.WriteRune(scanner.Current())

			scanner.Move()
			for unicode.IsNumber(scanner.Current()) {
				stringBuilder.WriteRune(scanner.Current())
				scanner.Move()
			}

			for scanner.Current() == '.' {
				reference := ParseReference()
				stringBuilder.WriteString(reference.String())
			}

			parts = append(parts, stringBuilder.String())
			stringBuilder.Reset()

			continue
		} else {
			stringBuilder.WriteRune(scanner.Current())
			scanner.Move()
		}
	}

	parts = append(parts, stringBuilder.String())
	stringBuilder.Reset()

	scanner.Move() // Move past the last `

	result := Script{
		Parts: parts,
	}

	stringBuilder.Reset()

	return result
}

func ParseEnvironmentVariable() Definition {
	scanner.Move()

	for scanner.Current() != -1 && !unicode.IsSpace(scanner.Current()) && !reservedCharacters[scanner.Current()] {
		stringBuilder.WriteRune(scanner.Current())
		scanner.Move()
	}

	result := EnvironmentVariable{
		Name: stringBuilder.String(),
	}
	stringBuilder.Reset()

	return result
}

func ParseReference() Definition {
	for scanner.Current() != -1 && !unicode.IsSpace(scanner.Current()) && !reservedCharacters[scanner.Current()] {
		stringBuilder.WriteRune(scanner.Current())
		scanner.Move()
	}

	result := Reference{
		Name: stringBuilder.String(),
	}
	stringBuilder.Reset()

	return result
}
