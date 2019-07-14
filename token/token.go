// Copyright 2019 The go-yaml Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"strconv"
)

// Token is the set of lexical single tokens of the YAML language.
type Token int

// The list of tokens.
const (
	// Special tokens
	ILLEGAL Token = iota // Uninitialized token
	EOF

	// begin:kind
	kindBegin
	STREAM_START
	STREAM_END

	VERSION_DIRECTIVE
	TAG_DIRECTIVE

	// begin:block
	blockBegin
	// Block
	BLOCK_SEQUENCE_START
	BLOCK_MAPPING_START
	BLOCK_ENTRY
	BLOCK_END
	blockEnd
	// end:block

	FLOW_ENTRY
	FLOW_SEQUENCE_START
	FLOW_SEQUENCE_END
	FLOW_MAPPING_START
	FLOW_MAPPING_END

	KEY
	VALUE
	SCALAR
	BLOCKSCALAR
	ALIAS
	ANCHOR
	TAG
	kindEnd
	// end:kind

	// begin:literal
	literalBegin
	// Identifiers and basic type literals
	IDENT  // main
	INT    // 12345
	BOOL   // true, false
	FLOAT  // 123.45
	CHAR   // 'a'
	STRING // "abc"
	literalEnd
	// end:literal

	// begin:delimiter
	delimiterBegin
	// Delimiters

	// begin:document
	documentBegin
	// document
	DOC_START // ---
	DOC_END   // ...
	documentEnd
	// end:document
	delimiterEnd
	// end:delimiter
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT:  "IDENT",
	INT:    "INT",
	BOOL:   "BOOL",
	FLOAT:  "FLOAT",
	CHAR:   "CHAR",
	STRING: "STRING",

	DOC_START: "---",
	DOC_END:   "...",
}

// String returns the string corresponding to the token tok.
// For, delimiters the string is the actual token character sequence.
//
// For all other tokens the string corresponds to the
// token constant name (e.g. for the token IDENT, the string is "IDENT").
func (tok Token) String() string {
	s := ""
	if (0 <= tok) && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}

	return s
}

var delimiters map[string]Token

func init() {
	delimiters = make(map[string]Token)

	for i := delimiterBegin + 1; i < delimiterEnd; i++ {
		delimiters[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT if not a keyword.
func Lookup(del string) Token {
	if tok, isDelimiter := delimiters[del]; isDelimiter {
		return tok
	}
	return IDENT
}

// IsKind returns whether the tokens corresponding to kinds.
func (tok Token) IsKind() bool { return kindBegin < tok && tok < kindEnd }

// IsBlock returns whether the tokens into the block statement.
func (tok Token) IsBlock() bool { return blockBegin < tok && tok < blockEnd }

// IsLiteral returns whether the tokens corresponding to basic type literals.
func (tok Token) IsLiteral() bool { return literalBegin < tok && tok < literalEnd }

// IsDelimiter returns whether the tokens corresponding to delimiters.
func (tok Token) IsDelimiter() bool { return delimiterBegin < tok && tok < delimiterEnd }

// IsDocument returns whether the tokens into the document statement.
func (tok Token) IsDocument() bool { return documentBegin < tok && tok < documentEnd }
