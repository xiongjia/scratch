%{

#include <stdio.h>
#include <stdlib.h>

#define YYDEBUG 1

extern int yyparse(void);
extern FILE *yyin;

int yylex(void);
int yyerror(const char *s);

%}

%union {
  int    int_value;
  double double_value;
}

%token <double_value> DOUBLE_LITERAL
%token ADD SUB MUL DIV CR
%type <double_value> expression term primary_expression

%%

line_list
  : line
  | line_list line
  ;

line
  : expression CR { printf(">> %lf\n", $1); }

expression
  : term
  | expression ADD term { $$ = $1 + $3; }
  | expression SUB term { $$ = $1 - $3; }
  ;
term
  : primary_expression
  | term MUL primary_expression { $$ = $1 * $3; }
  | term DIV primary_expression { $$ = $1 / $3; }
  ;
primary_expression
  : DOUBLE_LITERAL
  ;

%%

int main(void) {
  printf("Simple Computer: \n");
  yyin = stdin;
  if (yyparse()) {
    fprintf(stderr, "Error !\n");
    return 1;
  }
  return 0;
}

