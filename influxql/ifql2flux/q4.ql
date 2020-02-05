# fail on disjunction of time

SELECT "foo" FROM bar
	WHERE 3 AND ( ( time > 1 AND time < 2 ) OR 1 );

