# Time yanking

SELECT a FROM cpu WHERE host = 'server01' AND time > 1d;
SELECT a FROM cpu WHERE time > 1d AND host = 'server01';

SELECT a FROM cpu WHERE host = 'server01' OR time > 1d;
SELECT a FROM cpu WHERE time > 1d OR host = 'server01';

SELECT a FROM cpu WHERE ( time > 1d ) OR ( host = 'server01' );
SELECT a FROM cpu WHERE ( time > 2d OR time > 1d ) OR ( host = 'server01' );
