// Code generated from SAS.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type SASLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SASLexerLexerStaticData struct {
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

func saslexerLexerInit() {
	staticData := &SASLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'DATA'", "';'", "'RUN'", "'PUT'", "'SET'", "'PROC'", "'='", "'IF'",
		"'THEN'", "'ELSE'", "'>'", "'<'", "'>='", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "NUMBER",
		"STRING", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "ID", "NUMBER", "STRING",
		"WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 143, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 5, 14, 93, 8, 14, 10, 14, 12, 14, 96, 9, 14, 1, 15,
		4, 15, 99, 8, 15, 11, 15, 12, 15, 100, 1, 16, 1, 16, 5, 16, 105, 8, 16,
		10, 16, 12, 16, 108, 9, 16, 1, 16, 1, 16, 1, 17, 4, 17, 113, 8, 17, 11,
		17, 12, 17, 114, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 123,
		8, 18, 10, 18, 12, 18, 126, 9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 5, 19, 137, 8, 19, 10, 19, 12, 19, 140, 9, 19,
		1, 19, 1, 19, 1, 124, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13,
		39, 39, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 148, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 1, 41, 1, 0, 0, 0, 3, 46, 1, 0, 0, 0, 5, 48, 1, 0, 0, 0, 7, 52, 1, 0,
		0, 0, 9, 56, 1, 0, 0, 0, 11, 60, 1, 0, 0, 0, 13, 65, 1, 0, 0, 0, 15, 67,
		1, 0, 0, 0, 17, 70, 1, 0, 0, 0, 19, 75, 1, 0, 0, 0, 21, 80, 1, 0, 0, 0,
		23, 82, 1, 0, 0, 0, 25, 84, 1, 0, 0, 0, 27, 87, 1, 0, 0, 0, 29, 90, 1,
		0, 0, 0, 31, 98, 1, 0, 0, 0, 33, 102, 1, 0, 0, 0, 35, 112, 1, 0, 0, 0,
		37, 118, 1, 0, 0, 0, 39, 132, 1, 0, 0, 0, 41, 42, 5, 68, 0, 0, 42, 43,
		5, 65, 0, 0, 43, 44, 5, 84, 0, 0, 44, 45, 5, 65, 0, 0, 45, 2, 1, 0, 0,
		0, 46, 47, 5, 59, 0, 0, 47, 4, 1, 0, 0, 0, 48, 49, 5, 82, 0, 0, 49, 50,
		5, 85, 0, 0, 50, 51, 5, 78, 0, 0, 51, 6, 1, 0, 0, 0, 52, 53, 5, 80, 0,
		0, 53, 54, 5, 85, 0, 0, 54, 55, 5, 84, 0, 0, 55, 8, 1, 0, 0, 0, 56, 57,
		5, 83, 0, 0, 57, 58, 5, 69, 0, 0, 58, 59, 5, 84, 0, 0, 59, 10, 1, 0, 0,
		0, 60, 61, 5, 80, 0, 0, 61, 62, 5, 82, 0, 0, 62, 63, 5, 79, 0, 0, 63, 64,
		5, 67, 0, 0, 64, 12, 1, 0, 0, 0, 65, 66, 5, 61, 0, 0, 66, 14, 1, 0, 0,
		0, 67, 68, 5, 73, 0, 0, 68, 69, 5, 70, 0, 0, 69, 16, 1, 0, 0, 0, 70, 71,
		5, 84, 0, 0, 71, 72, 5, 72, 0, 0, 72, 73, 5, 69, 0, 0, 73, 74, 5, 78, 0,
		0, 74, 18, 1, 0, 0, 0, 75, 76, 5, 69, 0, 0, 76, 77, 5, 76, 0, 0, 77, 78,
		5, 83, 0, 0, 78, 79, 5, 69, 0, 0, 79, 20, 1, 0, 0, 0, 80, 81, 5, 62, 0,
		0, 81, 22, 1, 0, 0, 0, 82, 83, 5, 60, 0, 0, 83, 24, 1, 0, 0, 0, 84, 85,
		5, 62, 0, 0, 85, 86, 5, 61, 0, 0, 86, 26, 1, 0, 0, 0, 87, 88, 5, 60, 0,
		0, 88, 89, 5, 61, 0, 0, 89, 28, 1, 0, 0, 0, 90, 94, 7, 0, 0, 0, 91, 93,
		7, 1, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 30, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 99, 7,
		2, 0, 0, 98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0,
		100, 101, 1, 0, 0, 0, 101, 32, 1, 0, 0, 0, 102, 106, 5, 39, 0, 0, 103,
		105, 8, 3, 0, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104,
		1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 106, 1, 0,
		0, 0, 109, 110, 5, 39, 0, 0, 110, 34, 1, 0, 0, 0, 111, 113, 7, 4, 0, 0,
		112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 6, 17, 0, 0, 117, 36,
		1, 0, 0, 0, 118, 119, 5, 47, 0, 0, 119, 120, 5, 42, 0, 0, 120, 124, 1,
		0, 0, 0, 121, 123, 9, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0,
		0, 124, 125, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 127, 128, 5, 42, 0, 0, 128, 129, 5, 47, 0, 0, 129, 130,
		1, 0, 0, 0, 130, 131, 6, 18, 0, 0, 131, 38, 1, 0, 0, 0, 132, 133, 5, 47,
		0, 0, 133, 134, 5, 42, 0, 0, 134, 138, 1, 0, 0, 0, 135, 137, 8, 5, 0, 0,
		136, 135, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142,
		6, 19, 0, 0, 142, 40, 1, 0, 0, 0, 7, 0, 94, 100, 106, 114, 124, 138, 1,
		6, 0, 0,
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

// SASLexerInit initializes any static state used to implement SASLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSASLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SASLexerInit() {
	staticData := &SASLexerLexerStaticData
	staticData.once.Do(saslexerLexerInit)
}

// NewSASLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSASLexer(input antlr.CharStream) *SASLexer {
	SASLexerInit()
	l := new(SASLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SASLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SAS.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SASLexer tokens.
const (
	SASLexerT__0         = 1
	SASLexerT__1         = 2
	SASLexerT__2         = 3
	SASLexerT__3         = 4
	SASLexerT__4         = 5
	SASLexerT__5         = 6
	SASLexerT__6         = 7
	SASLexerT__7         = 8
	SASLexerT__8         = 9
	SASLexerT__9         = 10
	SASLexerT__10        = 11
	SASLexerT__11        = 12
	SASLexerT__12        = 13
	SASLexerT__13        = 14
	SASLexerID           = 15
	SASLexerNUMBER       = 16
	SASLexerSTRING       = 17
	SASLexerWS           = 18
	SASLexerCOMMENT      = 19
	SASLexerLINE_COMMENT = 20
)
