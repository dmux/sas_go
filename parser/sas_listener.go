// Code generated from SAS.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SAS

import "github.com/antlr4-go/antlr/v4"

// SASListener is a complete listener for a parse tree produced by SASParser.
type SASListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDataStep is called when entering the dataStep production.
	EnterDataStep(c *DataStepContext)

	// EnterPutStatement is called when entering the putStatement production.
	EnterPutStatement(c *PutStatementContext)

	// EnterSetStatement is called when entering the setStatement production.
	EnterSetStatement(c *SetStatementContext)

	// EnterProcStep is called when entering the procStep production.
	EnterProcStep(c *ProcStepContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterConditionalStatement is called when entering the conditionalStatement production.
	EnterConditionalStatement(c *ConditionalStatementContext)

	// EnterRunStatement is called when entering the runStatement production.
	EnterRunStatement(c *RunStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDataStep is called when exiting the dataStep production.
	ExitDataStep(c *DataStepContext)

	// ExitPutStatement is called when exiting the putStatement production.
	ExitPutStatement(c *PutStatementContext)

	// ExitSetStatement is called when exiting the setStatement production.
	ExitSetStatement(c *SetStatementContext)

	// ExitProcStep is called when exiting the procStep production.
	ExitProcStep(c *ProcStepContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitConditionalStatement is called when exiting the conditionalStatement production.
	ExitConditionalStatement(c *ConditionalStatementContext)

	// ExitRunStatement is called when exiting the runStatement production.
	ExitRunStatement(c *RunStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
