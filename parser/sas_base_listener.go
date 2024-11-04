// Code generated from SAS.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SAS

import "github.com/antlr4-go/antlr/v4"

// BaseSASListener is a complete listener for a parse tree produced by SASParser.
type BaseSASListener struct{}

var _ SASListener = &BaseSASListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSASListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSASListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSASListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSASListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseSASListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseSASListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSASListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSASListener) ExitStatement(ctx *StatementContext) {}

// EnterDataStep is called when production dataStep is entered.
func (s *BaseSASListener) EnterDataStep(ctx *DataStepContext) {}

// ExitDataStep is called when production dataStep is exited.
func (s *BaseSASListener) ExitDataStep(ctx *DataStepContext) {}

// EnterPutStatement is called when production putStatement is entered.
func (s *BaseSASListener) EnterPutStatement(ctx *PutStatementContext) {}

// ExitPutStatement is called when production putStatement is exited.
func (s *BaseSASListener) ExitPutStatement(ctx *PutStatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseSASListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseSASListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterProcStep is called when production procStep is entered.
func (s *BaseSASListener) EnterProcStep(ctx *ProcStepContext) {}

// ExitProcStep is called when production procStep is exited.
func (s *BaseSASListener) ExitProcStep(ctx *ProcStepContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseSASListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseSASListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterConditionalStatement is called when production conditionalStatement is entered.
func (s *BaseSASListener) EnterConditionalStatement(ctx *ConditionalStatementContext) {}

// ExitConditionalStatement is called when production conditionalStatement is exited.
func (s *BaseSASListener) ExitConditionalStatement(ctx *ConditionalStatementContext) {}

// EnterRunStatement is called when production runStatement is entered.
func (s *BaseSASListener) EnterRunStatement(ctx *RunStatementContext) {}

// ExitRunStatement is called when production runStatement is exited.
func (s *BaseSASListener) ExitRunStatement(ctx *RunStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSASListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSASListener) ExitExpression(ctx *ExpressionContext) {}
