# testing conjunction normalization

SELECT "foo" FROM bar WHERE
	( ( ( ( 
	( baz = 'yes' ) AND ( a AND b AND c ) AND ( a AND b ) AND 1 * 2
	) ) ) OR badness ) AND k;
