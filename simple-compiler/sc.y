%{

#include <stdio.h>
#include <stdlib.h>
int yylex(void);
int yyerror(const char *s);

%}

%token HI BYE

%%

program:
  hi
  ;

hi:
  HI  { printf("Hello\n"); }
  ;

