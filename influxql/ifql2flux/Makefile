COLM=$(HOME)/pkgs/colm-suite/bin/colm

all: ifql2flux parseflux

ifql2flux: ifql2flux.lm flux.lm influxql.lm
	$(COLM) -o $@ $<

parseflux: parseflux.lm flux.lm
	$(COLM) -o $@ $<
