// 通用的词法规则，注意是 lexer grammar
lexer grammar CommonLexerRules;
// 匹配标识符(+表示匹配一次或者多次)
ID : [a-zA-Z]+;
// 匹配整数
INT : [0-9]+;
// 匹配换行符(?表示匹配零次或者一次)
NEWLINE : '\r'?'\n';
// 丢弃空白字符
WS : [ \t]+ -> skip;
