grammar LibExpr;
// 引入 CommonLexerRules.g4 中全部的词法规则
import CommonLexerRules;

prog : stat+;
stat : expr NEWLINE             # printExpr
    | ID '=' expr NEWLINE       # assign
    | NEWLINE                   # blank
    ;
expr : expr op=('*' | '/') expr    # MulDiv
    | expr op=('+' | '-') expr     # AddSub
    | INT                       # int
    | ID                        # id
    | '(' expr ')'              # parens
    | 'clear'                   # clear
    ;

// 为上诉语法中使用的算术符命名
MUL : '*';
DIV : '/';
ADD : '+';
SUB : '-';
