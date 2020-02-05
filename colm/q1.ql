# testing conjunction nomralization

SELECT "foo" FROM bar WHERE
	( ( ( 
	baz = 'yes' AND ( a OR b AND c ) AND ( a AND b ) AND 1 * 2
	) ) );

