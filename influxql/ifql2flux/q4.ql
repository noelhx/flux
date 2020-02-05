# fail on disjunction of time

SELECT "foo" FROM bar
	WHERE 3 AND ( ( time > 1m AND time < 2m ) OR 1 );

