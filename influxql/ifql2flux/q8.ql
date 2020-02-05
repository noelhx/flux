# Also from #7530
#
# Seems as though there is a hack in the parser/engine that rearranges the
# parse tree. Not sure if this is supposed to work.
#
# https://github.com/influxdata/influxdb/issues/7530#issuecomment-326363141

SELECT a FROM cpu
WHERE host = 'server01' OR host = 'server02' AND time > now() - 1d;
