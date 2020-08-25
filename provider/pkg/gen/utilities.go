// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

import (
	"io"
	"strings"
	"unicode"
)

// isLegalIdentifierStart returns true if it is legal for c to be the first character of an identifier that we can
// use for schema properties.
func isLegalIdentifierStart(c rune) bool {
	return c == '_' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of an identifier (besides the first
// character) that we can use for schema properties.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || unicode.In(c, unicode.Mn, unicode.Mc, unicode.Nd, unicode.Pc)
}

// isLegalIdentifier returns true if s is a legal JavaScript identifier that we can use for schema properties.
func isLegalIdentifier(s string) bool {
	reader := strings.NewReader(s)
	c, _, _ := reader.ReadRune()
	if !isLegalIdentifierStart(c) {
		return false
	}
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			return err == io.EOF
		}
		if !isLegalIdentifierPart(c) {
			return false
		}
	}
}

// makeLegalIdentifier removes characters that are not allowed in identifiers.
func makeLegalIdentifier(s string) string {
	replacer := strings.NewReplacer("-", "", "[", "", "]", "")
	return replacer.Replace(s)
}

// firstToLower returns a string with the first character lowercased (`HelloWorld` => `helloWorld`).
func firstToLower(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	return string(append([]rune{unicode.ToLower(runes[0])}, runes[1:]...))
}

// toLowerCamel converts a string to lowerCamelCase.
// The code is adopted from https://github.com/iancoleman/strcase but changed in several ways to handle
// all the cases that are found in Azure in a most user-friendly way.
func toLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if uppercaseAcronym[s] {
		s = strings.ToLower(s)
	}
	if r := rune(s[0]); r == '_' {
		s = s[1:]
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

func toCamelInitCase(s string, initCase bool) string {
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' || v == '.' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

var uppercaseAcronym = map[string]bool{
	"ID": true,
	"TTL": true,
}