grammar SAS;

program
    : statement* EOF
    ;

statement
    : dataStep
    | putStatement
    | procStep
    | assignmentStatement
    | conditionalStatement
    | runStatement
    ;

dataStep
    : 'DATA' ID ';'
    (statement)*
    'RUN' ';'
    ;

putStatement
    : 'PUT' STRING ';'
    ;

setStatement
    : 'SET' ID ';'
    ;

procStep
    : 'PROC' ID ';'
      (assignmentStatement | conditionalStatement)*
      'RUN' ';'
    ;

assignmentStatement
    : ID '=' expression ';'
    ;

conditionalStatement
    : 'IF' expression 'THEN' assignmentStatement ('ELSE' assignmentStatement)? ';'
    ;

runStatement
    : 'RUN' ';'
    ;

expression
    : ID (('>' | '<' | '=' | '>=' | '<=') NUMBER)?
    | NUMBER
    | STRING
    ;

ID      : [a-zA-Z_][a-zA-Z0-9_]*;
NUMBER  : [0-9]+;
STRING  : '\'' ~('\''|'\r'|'\n')* '\'';

WS      : [ \t\r\n]+ -> skip;
COMMENT  : '/*' .*? '*/' -> skip;
LINE_COMMENT : '/*' ~[\r\n]* -> skip;

