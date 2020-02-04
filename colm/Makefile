all: testflux iql2flux

testflux: testflux.lm flux.lm
	colm -o $@ $<

iql2flux: iql2flux.lm flux.lm influxql.lm
	colm -o $@ $<
