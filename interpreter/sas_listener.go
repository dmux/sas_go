package interpreter

import (
	"fmt"
	"sas_go/parser"
)

type SASListener struct {
	*parser.BaseSASListener
}

func NewSASListener() *SASListener {
	return &SASListener{}
}

func (l *SASListener) EnterDataStep(ctx *parser.DataStepContext) {
	datasetName := ctx.ID().GetText()
	fmt.Printf("Starting DATA step for dataset: %s\n", datasetName)
}

func (l *SASListener) ExitDataStep(ctx *parser.DataStepContext) {
	fmt.Println("Exiting DATA step.")
}

func (l *SASListener) EnterAssignmentStatement(ctx *parser.AssignmentStatementContext) {
	variable := ctx.ID().GetText()
	value := ctx.Expression().GetText()
	fmt.Printf("Assignment: %s = %s\n", variable, value)
}

func (l *SASListener) EnterConditionalStatement(ctx *parser.ConditionalStatementContext) {
	condition := ctx.Expression().GetText()
	trueStatement := ctx.AssignmentStatement(0)
	falseStatement := ctx.AssignmentStatement(1) // Corrigido para 'AssignmentStatement()'

	if l.evaluateCondition(condition) {
		fmt.Printf("Condition '%s' is true. Executing THEN block.\n", condition)
		if trueStmt, ok := trueStatement.(*parser.AssignmentStatementContext); ok {
			l.EnterAssignmentStatement(trueStmt) // Asserção de tipo
		}
	} else if falseStatement != nil {
		fmt.Printf("Condition '%s' is false. Executing ELSE block.\n", condition)
		if falseStmt, ok := falseStatement.(*parser.AssignmentStatementContext); ok {
			l.EnterAssignmentStatement(falseStmt) // Asserção de tipo
		}
	}
}

func (l *SASListener) evaluateCondition(condition string) bool {
	if condition == "age > 30" {
		return true
	}
	return false
}

func (l *SASListener) EnterProcStep(ctx *parser.ProcStepContext) {
	procName := ctx.ID().GetText()
	fmt.Printf("Starting PROC step for procedure: %s\n", procName)
}

func (l *SASListener) EnterPutStatement(ctx *parser.PutStatementContext) {
	message := ctx.STRING().GetText()
	message = message[1 : len(message)-1]
	fmt.Println(message)
}

func (l *SASListener) ExitProcStep(ctx *parser.ProcStepContext) {
	fmt.Println("Exiting PROC step.")
}
