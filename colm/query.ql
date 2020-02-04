SELECT "foo" FROM bar WHERE baz = 'yes';

SELECT mean("ENERGY_Power") FROM "consumer"
	WHERE ( "topic" = 'power_meter/solar/SENSOR' AND time >= now() -3m )
	GROUP BY time(1m);

SELECT ENERGY_Power FROM "mqtt_consumer"
	WHERE ( "topic" = 'power_meter/solar/SENSOR' ) AND time >= now() - 3m;
