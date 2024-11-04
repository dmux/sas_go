// Code generated from SAS.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SAS

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SASParser struct {
	*antlr.BaseParser
}

var SASParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sasParserInit() {
	staticData := &SASParserStaticData
	staticData.LiteralNames = []string{
		"", "'DATA'", "';'", "'RUN'", "'PUT'", "'SET'", "'PROC'", "'='", "'IF'",
		"'THEN'", "'ELSE'", "'>'", "'<'", "'>='", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "NUMBER",
		"STRING", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "dataStep", "putStatement", "setStatement",
		"procStep", "assignmentStatement", "conditionalStatement", "runStatement",
		"expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 97, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 5,
		0, 22, 8, 0, 10, 0, 12, 0, 25, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 35, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 41, 8, 2, 10,
		2, 12, 2, 44, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 62, 8, 5, 10, 5, 12,
		5, 65, 9, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 81, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 3, 9, 91, 8, 9, 1, 9, 1, 9, 3, 9, 95, 8, 9, 1, 9, 0, 0,
		10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 1, 2, 0, 7, 7, 11, 14, 99, 0,
		23, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 36, 1, 0, 0, 0, 6, 48, 1, 0, 0, 0,
		8, 52, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 69, 1, 0, 0, 0, 14, 74, 1, 0,
		0, 0, 16, 84, 1, 0, 0, 0, 18, 94, 1, 0, 0, 0, 20, 22, 3, 2, 1, 0, 21, 20,
		1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0,
		24, 26, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 26, 27, 5, 0, 0, 1, 27, 1, 1, 0,
		0, 0, 28, 35, 3, 4, 2, 0, 29, 35, 3, 6, 3, 0, 30, 35, 3, 10, 5, 0, 31,
		35, 3, 12, 6, 0, 32, 35, 3, 14, 7, 0, 33, 35, 3, 16, 8, 0, 34, 28, 1, 0,
		0, 0, 34, 29, 1, 0, 0, 0, 34, 30, 1, 0, 0, 0, 34, 31, 1, 0, 0, 0, 34, 32,
		1, 0, 0, 0, 34, 33, 1, 0, 0, 0, 35, 3, 1, 0, 0, 0, 36, 37, 5, 1, 0, 0,
		37, 38, 5, 15, 0, 0, 38, 42, 5, 2, 0, 0, 39, 41, 3, 2, 1, 0, 40, 39, 1,
		0, 0, 0, 41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43,
		45, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 46, 5, 3, 0, 0, 46, 47, 5, 2, 0,
		0, 47, 5, 1, 0, 0, 0, 48, 49, 5, 4, 0, 0, 49, 50, 5, 17, 0, 0, 50, 51,
		5, 2, 0, 0, 51, 7, 1, 0, 0, 0, 52, 53, 5, 5, 0, 0, 53, 54, 5, 15, 0, 0,
		54, 55, 5, 2, 0, 0, 55, 9, 1, 0, 0, 0, 56, 57, 5, 6, 0, 0, 57, 58, 5, 15,
		0, 0, 58, 63, 5, 2, 0, 0, 59, 62, 3, 12, 6, 0, 60, 62, 3, 14, 7, 0, 61,
		59, 1, 0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0,
		0, 63, 64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 67,
		5, 3, 0, 0, 67, 68, 5, 2, 0, 0, 68, 11, 1, 0, 0, 0, 69, 70, 5, 15, 0, 0,
		70, 71, 5, 7, 0, 0, 71, 72, 3, 18, 9, 0, 72, 73, 5, 2, 0, 0, 73, 13, 1,
		0, 0, 0, 74, 75, 5, 8, 0, 0, 75, 76, 3, 18, 9, 0, 76, 77, 5, 9, 0, 0, 77,
		80, 3, 12, 6, 0, 78, 79, 5, 10, 0, 0, 79, 81, 3, 12, 6, 0, 80, 78, 1, 0,
		0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 5, 2, 0, 0, 83, 15,
		1, 0, 0, 0, 84, 85, 5, 3, 0, 0, 85, 86, 5, 2, 0, 0, 86, 17, 1, 0, 0, 0,
		87, 90, 5, 15, 0, 0, 88, 89, 7, 0, 0, 0, 89, 91, 5, 16, 0, 0, 90, 88, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 95, 1, 0, 0, 0, 92, 95, 5, 16, 0, 0, 93,
		95, 5, 17, 0, 0, 94, 87, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 93, 1, 0,
		0, 0, 95, 19, 1, 0, 0, 0, 8, 23, 34, 42, 61, 63, 80, 90, 94,
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

// SASParserInit initializes any static state used to implement SASParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSASParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SASParserInit() {
	staticData := &SASParserStaticData
	staticData.once.Do(sasParserInit)
}

// NewSASParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSASParser(input antlr.TokenStream) *SASParser {
	SASParserInit()
	this := new(SASParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SASParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SAS.g4"

	return this
}

// SASParser tokens.
const (
	SASParserEOF          = antlr.TokenEOF
	SASParserT__0         = 1
	SASParserT__1         = 2
	SASParserT__2         = 3
	SASParserT__3         = 4
	SASParserT__4         = 5
	SASParserT__5         = 6
	SASParserT__6         = 7
	SASParserT__7         = 8
	SASParserT__8         = 9
	SASParserT__9         = 10
	SASParserT__10        = 11
	SASParserT__11        = 12
	SASParserT__12        = 13
	SASParserT__13        = 14
	SASParserID           = 15
	SASParserNUMBER       = 16
	SASParserSTRING       = 17
	SASParserWS           = 18
	SASParserCOMMENT      = 19
	SASParserLINE_COMMENT = 20
)

// SASParser rules.
const (
	SASParserRULE_program              = 0
	SASParserRULE_statement            = 1
	SASParserRULE_dataStep             = 2
	SASParserRULE_putStatement         = 3
	SASParserRULE_setStatement         = 4
	SASParserRULE_procStep             = 5
	SASParserRULE_assignmentStatement  = 6
	SASParserRULE_conditionalStatement = 7
	SASParserRULE_runStatement         = 8
	SASParserRULE_expression           = 9
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(SASParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *SASParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SASParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33114) != 0 {
		{
			p.SetState(20)
			p.Statement()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
		p.Match(SASParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DataStep() IDataStepContext
	PutStatement() IPutStatementContext
	ProcStep() IProcStepContext
	AssignmentStatement() IAssignmentStatementContext
	ConditionalStatement() IConditionalStatementContext
	RunStatement() IRunStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) DataStep() IDataStepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataStepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataStepContext)
}

func (s *StatementContext) PutStatement() IPutStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPutStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPutStatementContext)
}

func (s *StatementContext) ProcStep() IProcStepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcStepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcStepContext)
}

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) ConditionalStatement() IConditionalStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalStatementContext)
}

func (s *StatementContext) RunStatement() IRunStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SASParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SASParserRULE_statement)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SASParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.DataStep()
		}

	case SASParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.PutStatement()
		}

	case SASParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.ProcStep()
		}

	case SASParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(31)
			p.AssignmentStatement()
		}

	case SASParserT__7:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(32)
			p.ConditionalStatement()
		}

	case SASParserT__2:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(33)
			p.RunStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataStepContext is an interface to support dynamic dispatch.
type IDataStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsDataStepContext differentiates from other interfaces.
	IsDataStepContext()
}

type DataStepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataStepContext() *DataStepContext {
	var p = new(DataStepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_dataStep
	return p
}

func InitEmptyDataStepContext(p *DataStepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_dataStep
}

func (*DataStepContext) IsDataStepContext() {}

func NewDataStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataStepContext {
	var p = new(DataStepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_dataStep

	return p
}

func (s *DataStepContext) GetParser() antlr.Parser { return s.parser }

func (s *DataStepContext) ID() antlr.TerminalNode {
	return s.GetToken(SASParserID, 0)
}

func (s *DataStepContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *DataStepContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DataStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterDataStep(s)
	}
}

func (s *DataStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitDataStep(s)
	}
}

func (p *SASParser) DataStep() (localctx IDataStepContext) {
	localctx = NewDataStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SASParserRULE_dataStep)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(SASParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.Match(SASParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(39)
				p.Statement()
			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(SASParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPutStatementContext is an interface to support dynamic dispatch.
type IPutStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsPutStatementContext differentiates from other interfaces.
	IsPutStatementContext()
}

type PutStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPutStatementContext() *PutStatementContext {
	var p = new(PutStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_putStatement
	return p
}

func InitEmptyPutStatementContext(p *PutStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_putStatement
}

func (*PutStatementContext) IsPutStatementContext() {}

func NewPutStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PutStatementContext {
	var p = new(PutStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_putStatement

	return p
}

func (s *PutStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PutStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(SASParserSTRING, 0)
}

func (s *PutStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PutStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PutStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterPutStatement(s)
	}
}

func (s *PutStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitPutStatement(s)
	}
}

func (p *SASParser) PutStatement() (localctx IPutStatementContext) {
	localctx = NewPutStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SASParserRULE_putStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(SASParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(SASParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetStatementContext is an interface to support dynamic dispatch.
type ISetStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsSetStatementContext differentiates from other interfaces.
	IsSetStatementContext()
}

type SetStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetStatementContext() *SetStatementContext {
	var p = new(SetStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_setStatement
	return p
}

func InitEmptySetStatementContext(p *SetStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_setStatement
}

func (*SetStatementContext) IsSetStatementContext() {}

func NewSetStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetStatementContext {
	var p = new(SetStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_setStatement

	return p
}

func (s *SetStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SetStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(SASParserID, 0)
}

func (s *SetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterSetStatement(s)
	}
}

func (s *SetStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitSetStatement(s)
	}
}

func (p *SASParser) SetStatement() (localctx ISetStatementContext) {
	localctx = NewSetStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SASParserRULE_setStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(SASParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(SASParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProcStepContext is an interface to support dynamic dispatch.
type IProcStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllAssignmentStatement() []IAssignmentStatementContext
	AssignmentStatement(i int) IAssignmentStatementContext
	AllConditionalStatement() []IConditionalStatementContext
	ConditionalStatement(i int) IConditionalStatementContext

	// IsProcStepContext differentiates from other interfaces.
	IsProcStepContext()
}

type ProcStepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcStepContext() *ProcStepContext {
	var p = new(ProcStepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_procStep
	return p
}

func InitEmptyProcStepContext(p *ProcStepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_procStep
}

func (*ProcStepContext) IsProcStepContext() {}

func NewProcStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcStepContext {
	var p = new(ProcStepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_procStep

	return p
}

func (s *ProcStepContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcStepContext) ID() antlr.TerminalNode {
	return s.GetToken(SASParserID, 0)
}

func (s *ProcStepContext) AllAssignmentStatement() []IAssignmentStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentStatementContext); ok {
			tst[i] = t.(IAssignmentStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProcStepContext) AssignmentStatement(i int) IAssignmentStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *ProcStepContext) AllConditionalStatement() []IConditionalStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalStatementContext); ok {
			len++
		}
	}

	tst := make([]IConditionalStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalStatementContext); ok {
			tst[i] = t.(IConditionalStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProcStepContext) ConditionalStatement(i int) IConditionalStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalStatementContext)
}

func (s *ProcStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterProcStep(s)
	}
}

func (s *ProcStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitProcStep(s)
	}
}

func (p *SASParser) ProcStep() (localctx IProcStepContext) {
	localctx = NewProcStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SASParserRULE_procStep)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(SASParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(SASParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SASParserT__7 || _la == SASParserID {
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SASParserID:
			{
				p.SetState(59)
				p.AssignmentStatement()
			}

		case SASParserT__7:
			{
				p.SetState(60)
				p.ConditionalStatement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(SASParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(SASParserID, 0)
}

func (s *AssignmentStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (p *SASParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SASParserRULE_assignmentStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(SASParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(SASParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Expression()
	}
	{
		p.SetState(72)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalStatementContext is an interface to support dynamic dispatch.
type IConditionalStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllAssignmentStatement() []IAssignmentStatementContext
	AssignmentStatement(i int) IAssignmentStatementContext

	// IsConditionalStatementContext differentiates from other interfaces.
	IsConditionalStatementContext()
}

type ConditionalStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalStatementContext() *ConditionalStatementContext {
	var p = new(ConditionalStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_conditionalStatement
	return p
}

func InitEmptyConditionalStatementContext(p *ConditionalStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_conditionalStatement
}

func (*ConditionalStatementContext) IsConditionalStatementContext() {}

func NewConditionalStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalStatementContext {
	var p = new(ConditionalStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_conditionalStatement

	return p
}

func (s *ConditionalStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalStatementContext) AllAssignmentStatement() []IAssignmentStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentStatementContext); ok {
			tst[i] = t.(IAssignmentStatementContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalStatementContext) AssignmentStatement(i int) IAssignmentStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *ConditionalStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterConditionalStatement(s)
	}
}

func (s *ConditionalStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitConditionalStatement(s)
	}
}

func (p *SASParser) ConditionalStatement() (localctx IConditionalStatementContext) {
	localctx = NewConditionalStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SASParserRULE_conditionalStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(SASParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Expression()
	}
	{
		p.SetState(76)
		p.Match(SASParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.AssignmentStatement()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SASParserT__9 {
		{
			p.SetState(78)
			p.Match(SASParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.AssignmentStatement()
		}

	}
	{
		p.SetState(82)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRunStatementContext is an interface to support dynamic dispatch.
type IRunStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRunStatementContext differentiates from other interfaces.
	IsRunStatementContext()
}

type RunStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunStatementContext() *RunStatementContext {
	var p = new(RunStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_runStatement
	return p
}

func InitEmptyRunStatementContext(p *RunStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_runStatement
}

func (*RunStatementContext) IsRunStatementContext() {}

func NewRunStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunStatementContext {
	var p = new(RunStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_runStatement

	return p
}

func (s *RunStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *RunStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RunStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterRunStatement(s)
	}
}

func (s *RunStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitRunStatement(s)
	}
}

func (p *SASParser) RunStatement() (localctx IRunStatementContext) {
	localctx = NewRunStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SASParserRULE_runStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(SASParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(SASParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SASParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SASParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(SASParserID, 0)
}

func (s *ExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SASParserNUMBER, 0)
}

func (s *ExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(SASParserSTRING, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SASListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SASParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SASParserRULE_expression)
	var _la int

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SASParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(SASParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30848) != 0 {
			{
				p.SetState(88)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30848) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(89)
				p.Match(SASParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SASParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(SASParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SASParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Match(SASParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
