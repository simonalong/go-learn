grammar sql;
import SqlRules;

statement: SELECT fields FROM topic WHERE conditions;


// 字段
// 1. 所有字段：*
// 2. 单个属性：name
// 3. 多个属性：name,age,test
// 4. 函数：count()
// 5. 表达式：name as n
fields
    : '*'
    | field (',' field)*
    ;

// 某个字段
// 1. 表达式 xxx as xx
// 2. 属性 xx
// 3. 函数 xxx()
field
    : express_as
    | attr
    | function
    ;

// 表达式
// 1. 属性 as xxx
// 2. 函数 as xxx
express_as
    : attr AS string_attr
    | function AS string_attr
    ;

// 属性
attr
    : string_attr('.'string_attr)*
    ;

// 定义字符串规则
// 支持：下划线或者中划线或者字符开头的字符串，非头部支持字符串、数字、中划线、下划线
string_attr
    : UNDERSCORE*LETTER(LETTER|DIGIT|HYPHEN|UNDERSCORE)*
    ;


// 函数
// 1. 无参数
// 2. 一个或者多个参数
function
    : string_attr'('(string_attr(','WHITESPACE* string_attr))?')'
    ;

// 主题：两种格式都支持：这里不解析了，我们直接透传
//topic
//    : '\''NON_WHITESPACE'\''
//    | '"'NON_WHITESPACE'"'
//    ;

topic
    : '"' (ESC|.)*? '"';
fragment ESC : '\\"' | '\\\\';

// 单行注释(以//开头，换行结束)
LINE_COMMENT : '//' .*? '\r'?'\n' -> skip;
// 多行注释(/* */包裹的所有字符)
COMMENT : '/*' .*? '*/' -> skip;


// 条件表达式
// 关系：
//      比较操作：< <= > >= = <>
//      逻辑操作：and or not
//      列表：in
//      匹配：like
// 字段：
//      属性
//      函数
conditions
    : express_compare
    | express_logic
    | express_in
    | express_like
    ;

// 表达式元素
express_element
    : attr
    | express_element MATH_OPERATOR express_element
    | '('express_element MATH_OPERATOR express_element')'
    ;

// 比较操作
express_compare
    : express_element COMPARE_OPERATOR DIGIT
    | '(' express_element COMPARE_OPERATOR DIGIT ')'
    | DIGIT COMPARE_OPERATOR express_element
    | '(' DIGIT COMPARE_OPERATOR express_element ')'
    ;

// 逻辑操作
express_logic
    : LOGIC_OPERATOR_NOT express_logic
    | '('LOGIC_OPERATOR_NOT express_logic')'
    | express_logic (LOGIC_OPERATOR_AND|LOGIC_OPERATOR_OR) express_logic
    | '('express_logic (LOGIC_OPERATOR_AND|LOGIC_OPERATOR_OR) express_logic')'
    | express_compare
    | express_in
    | express_like
    | '('express_compare')'
    | '('express_in')'
    | '('express_like')'
    | '('express_logic')'
    ;

// 列表表达式
express_in
    : attr LIST_OPERATOR_IN '('NON_WHITESPACE(','NON_WHITESPACE)')';

// 匹配表达式
// 1. 匹配某个字符
// 2. 匹配通配符%，这个只可以在最前或者最后两边
express_like
    : attr FIND_OPERATOR_LIKE '\''FIND_OPERATOR_LIKE?NON_WHITESPACE FIND_OPERATOR_LIKE?'\''
    | attr FIND_OPERATOR_NONE_LIKE '\''FIND_OPERATOR_LIKE?NON_WHITESPACE FIND_OPERATOR_LIKE?'\''
    | '('attr FIND_OPERATOR_LIKE '\''FIND_OPERATOR_LIKE?NON_WHITESPACE FIND_OPERATOR_LIKE?'\''')'
    | '('attr FIND_OPERATOR_NONE_LIKE '\''FIND_OPERATOR_LIKE?NON_WHITESPACE FIND_OPERATOR_LIKE?'\''')';


// 空白字符
fragment WHITESPACE:[ \t\r\n]+;
// 非空格的字符集
fragment NON_WHITESPACE: ~[ \t\r\n];
// 匹配任何大小写字母
fragment LETTER: [a-zA-Z] ;
 // 匹配任何数字
fragment DIGIT: [0-9];
// 匹配中划线字符
fragment HYPHEN: '-';
// 匹配下划线字符
fragment UNDERSCORE: '_';
// 数学表达式
fragment MATH_OPERATOR
    :(MATH_OPERATOR_PLUS|MATH_OPERATOR_REDUCE|MATH_OPERATOR_MULTI|MATH_OPERATOR_DIVICE|MATH_OPERATOR_REMAINDER);
// 比较表达式
fragment COMPARE_OPERATOR
    :(COMPARE_OPERATOR_GREATER|COMPARE_OPERATOR_GREATER_EQ|COMPARE_OPERATOR_LESS|COMPARE_OPERATOR_LESS_EQ|COMPARE_OPERATOR_EQUAL|COMPARE_OPERATOR_NONE_EQUAL);

// 匹配整数（正整数或负整数）
fragment INT
    : '-'? [1-9] [0-9]*
    |'0'
    ;
// 匹配小数（正小数或负小数）
fragment DECIMAL : '-'? [0-9]* '.' [0-9]+;
// 匹配正数或负数，可以是整数或小数
fragment NUMBER : INT | DECIMAL;

