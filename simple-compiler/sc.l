%{

#include <stdio.h>
#include "y.tab.h"

extern char *yytext;

int yyerror(char const *errormsg);

%}
  
%%

[ \t] ;
"+"   return ADD;
"-"   return SUB;
"*"   return MUL;
"/"   return DIV;
"\n"  return CR;

([1-9][0-9]*)|0|([0-9]+\.[0-9]*) {
  double temp = 0;
  sscanf(yytext, "%lf", &temp);
  yylval.double_value = temp;
  return DOUBLE_LITERAL;
}

. { yyerror("Unknow char");  }

%%

int yywrap(void) {
  return 0;
}

int yyerror(char const *errormsg) {
  fprintf(stderr, "parser error near %s\n", yytext);
  return 0;
}

