# Disjunction of time from Issue referenced by DOC faq.
# https://github.com/influxdata/influxdb/issues/7530

SELECT a FROM cpu WHERE time >= now() - 10m OR (time >= now() - 20m AND region = 'uswest');
