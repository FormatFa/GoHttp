// Code generated from /Users/ljh/Documents/personal/antrl4test/src/main/java/online/indigo6a/gohttp/GoHttp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GoHttpLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GoHttpLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gohttplexerLexerInit() {
	staticData := &GoHttpLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'HTTP/'", "'GET'", "'HEAD'", "'POST'", "'PUT'", "'DELETE'", "'CONNECT'",
		"'OPTIONS'", "'TRACE'", "' '", "", "", "", "'.'", "'-'", "'#'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "SP", "ALPHA", "DIGIT", "HEXDIG",
		"Dot", "DASH", "Hashtag", "Colon", "CRLF", "OTHER", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"SP", "ALPHA", "DIGIT", "HEXDIG", "Dot", "DASH", "Hashtag", "Colon",
		"CRLF", "OTHER", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 120, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 3, 12, 103, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 0, 0, 20,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		1, 0, 4, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 3,
		0, 9, 10, 13, 13, 32, 32, 120, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 47, 1, 0,
		0, 0, 5, 51, 1, 0, 0, 0, 7, 56, 1, 0, 0, 0, 9, 61, 1, 0, 0, 0, 11, 65,
		1, 0, 0, 0, 13, 72, 1, 0, 0, 0, 15, 80, 1, 0, 0, 0, 17, 88, 1, 0, 0, 0,
		19, 94, 1, 0, 0, 0, 21, 96, 1, 0, 0, 0, 23, 98, 1, 0, 0, 0, 25, 102, 1,
		0, 0, 0, 27, 104, 1, 0, 0, 0, 29, 106, 1, 0, 0, 0, 31, 108, 1, 0, 0, 0,
		33, 110, 1, 0, 0, 0, 35, 112, 1, 0, 0, 0, 37, 114, 1, 0, 0, 0, 39, 116,
		1, 0, 0, 0, 41, 42, 5, 72, 0, 0, 42, 43, 5, 84, 0, 0, 43, 44, 5, 84, 0,
		0, 44, 45, 5, 80, 0, 0, 45, 46, 5, 47, 0, 0, 46, 2, 1, 0, 0, 0, 47, 48,
		5, 71, 0, 0, 48, 49, 5, 69, 0, 0, 49, 50, 5, 84, 0, 0, 50, 4, 1, 0, 0,
		0, 51, 52, 5, 72, 0, 0, 52, 53, 5, 69, 0, 0, 53, 54, 5, 65, 0, 0, 54, 55,
		5, 68, 0, 0, 55, 6, 1, 0, 0, 0, 56, 57, 5, 80, 0, 0, 57, 58, 5, 79, 0,
		0, 58, 59, 5, 83, 0, 0, 59, 60, 5, 84, 0, 0, 60, 8, 1, 0, 0, 0, 61, 62,
		5, 80, 0, 0, 62, 63, 5, 85, 0, 0, 63, 64, 5, 84, 0, 0, 64, 10, 1, 0, 0,
		0, 65, 66, 5, 68, 0, 0, 66, 67, 5, 69, 0, 0, 67, 68, 5, 76, 0, 0, 68, 69,
		5, 69, 0, 0, 69, 70, 5, 84, 0, 0, 70, 71, 5, 69, 0, 0, 71, 12, 1, 0, 0,
		0, 72, 73, 5, 67, 0, 0, 73, 74, 5, 79, 0, 0, 74, 75, 5, 78, 0, 0, 75, 76,
		5, 78, 0, 0, 76, 77, 5, 69, 0, 0, 77, 78, 5, 67, 0, 0, 78, 79, 5, 84, 0,
		0, 79, 14, 1, 0, 0, 0, 80, 81, 5, 79, 0, 0, 81, 82, 5, 80, 0, 0, 82, 83,
		5, 84, 0, 0, 83, 84, 5, 73, 0, 0, 84, 85, 5, 79, 0, 0, 85, 86, 5, 78, 0,
		0, 86, 87, 5, 83, 0, 0, 87, 16, 1, 0, 0, 0, 88, 89, 5, 84, 0, 0, 89, 90,
		5, 82, 0, 0, 90, 91, 5, 65, 0, 0, 91, 92, 5, 67, 0, 0, 92, 93, 5, 69, 0,
		0, 93, 18, 1, 0, 0, 0, 94, 95, 5, 32, 0, 0, 95, 20, 1, 0, 0, 0, 96, 97,
		7, 0, 0, 0, 97, 22, 1, 0, 0, 0, 98, 99, 7, 1, 0, 0, 99, 24, 1, 0, 0, 0,
		100, 103, 3, 23, 11, 0, 101, 103, 2, 65, 70, 0, 102, 100, 1, 0, 0, 0, 102,
		101, 1, 0, 0, 0, 103, 26, 1, 0, 0, 0, 104, 105, 5, 46, 0, 0, 105, 28, 1,
		0, 0, 0, 106, 107, 5, 45, 0, 0, 107, 30, 1, 0, 0, 0, 108, 109, 5, 35, 0,
		0, 109, 32, 1, 0, 0, 0, 110, 111, 5, 58, 0, 0, 111, 34, 1, 0, 0, 0, 112,
		113, 7, 2, 0, 0, 113, 36, 1, 0, 0, 0, 114, 115, 9, 0, 0, 0, 115, 38, 1,
		0, 0, 0, 116, 117, 7, 3, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 6, 19, 0,
		0, 119, 40, 1, 0, 0, 0, 2, 0, 102, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoHttpLexerInit initializes any static state used to implement GoHttpLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoHttpLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoHttpLexerInit() {
	staticData := &GoHttpLexerLexerStaticData
	staticData.once.Do(gohttplexerLexerInit)
}

// NewGoHttpLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoHttpLexer(input antlr.CharStream) *GoHttpLexer {
	GoHttpLexerInit()
	l := new(GoHttpLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GoHttpLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "GoHttp.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GoHttpLexer tokens.
const (
	GoHttpLexerT__0    = 1
	GoHttpLexerT__1    = 2
	GoHttpLexerT__2    = 3
	GoHttpLexerT__3    = 4
	GoHttpLexerT__4    = 5
	GoHttpLexerT__5    = 6
	GoHttpLexerT__6    = 7
	GoHttpLexerT__7    = 8
	GoHttpLexerT__8    = 9
	GoHttpLexerSP      = 10
	GoHttpLexerALPHA   = 11
	GoHttpLexerDIGIT   = 12
	GoHttpLexerHEXDIG  = 13
	GoHttpLexerDot     = 14
	GoHttpLexerDASH    = 15
	GoHttpLexerHashtag = 16
	GoHttpLexerColon   = 17
	GoHttpLexerCRLF    = 18
	GoHttpLexerOTHER   = 19
	GoHttpLexerWS      = 20
)
