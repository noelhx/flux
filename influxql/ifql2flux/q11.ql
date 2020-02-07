SELECT usage_user
	FROM cpu
	WHERE time > "2020-02-07T12:42:21Z" AND ( cpu = "cpu0" OR cpu = "cpu1" );

