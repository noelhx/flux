all: flux influxql

flux: flux.lm
	colm -o flux flux.lm

influxql: influxql.lm
	colm -o $@ $<
