# build scripts for simple compiler

CC = clang
LEX = flex
BISON = bison
RM = rm -rvf

CFLAGS =

all: sc

lex: lex.yy.c
bison: y.tab.c
yacc: bison

lex.yy.c: sc.l
	$(LEX) -o $@ $<

y.tab.h: sc.y
y.tab.c: sc.y
	$(BISON) -dy -o $@ $<

sc: lex yacc
	$(CC) lex.yy.c y.tab.c -o $@ $(CFLAGS) 

clean:
	$(RM) lex.yy.c y.tab.{c,h} sc

