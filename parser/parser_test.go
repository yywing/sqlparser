package parser

import "testing"

func TestIsHive(t *testing.T) {
	sqls := map[string]bool{
		"create external table if not exists t (i int, d decimal) stored as orc;": true,
		"create view v1 as select c from t":                                       true,
		"drop table if exists t":                                                  true,
		"SELECT A + - B * C ^ D + E | F OR G AND H || I - (SELECT COUNT(*) FROM DB.TB) OR J BETWEEN K AND L OR M IS TRUE FROM DB.TB1": true,
		`WITH A AS (
			SELECT *, 1e1 as c, 1E1 as d, 1.0 as e FROM TB
		)
		SELECT A,B,C
		FROM A`: true,
		"show extended tables in db where t = 't'": true,
	}
	for sql, expect := range sqls {
		result := IsHive(sql)
		if result != expect {
			t.Errorf("IsHive(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsMysql(t *testing.T) {
	sqls := map[string]bool{
		"SELECT name from aaa":                                             true,
		"select /* back-quote num */ `2a` from t":                          true,
		"select * from t1 where k like 'Müller' collate latin1_german2_ci": true,

		"INSERT INTO aaa (name, age) VALUES ('a', 1)": true,
		"update /* bool in update */ a set b = true":  true,
		"delete /* where */ from a where a = b":       true,

		"alter table a reorganize partition b into (partition c values less than (?), partition d values less than (maxvalue))": true,
		"create table a (\n\t`a` int\n)":    true,
		"drop view a":                       true,
		"SELECT name from aaa; drop view a": true,

		"I am a good select body gg where": false,
	}
	for sql, expect := range sqls {
		result := IsMysql(sql)
		if result != expect {
			t.Errorf("IsMysql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsPostgresql(t *testing.T) {
	sqls := map[string]bool{
		"SELECT name from aaa":                                             true,
		"select /* back-quote num */ \"2a\" from t":                        true,
		"select * from t1 where k like 'Müller' collate latin1_german2_ci": true,

		"INSERT INTO aaa (name, age) VALUES ('a', 1)": true,
		"update /* bool in update */ a set b = true":  true,
		"delete /* where */ from a where a = b":       true,

		"ALTER TABLE a DETACH PARTITION b;": true,
		"create table a (\n\t\"a\" int\n)":  true,
		"drop view a":                       true,
		"SELECT name from aaa; drop view a": true,

		"I am a good select body gg where": false,
	}
	for sql, expect := range sqls {
		result := IsPostgresql(sql)
		if result != expect {
			t.Errorf("IsPostgresql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsSQLite(t *testing.T) {
	sqls := map[string]bool{
		"select * from x  ": true,
		"Select Name, printf('%,d',Bytes) Size,     FIRST_VALUE(Name) OVER (         ORDER BY Bytes     ) AS SmallestTrack FROM     tracks WHERE     AlbumId = 1; ": true,
		"CREATE TABLE f (     id INT PRIMARY KEY NOT NULL,     i  TEXT            NOT NULL,     j  INT             NOT NULL,     k  CHAR(50),     m  REAL ); ":      true,
		"SELECT * FROM Song JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song NATURAL JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song LEFT JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song LEFT OUTER JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song RIGHT JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song RIGHT OUTER JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song FULL JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song FULL OUTER JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song INNER JOIN Album ON Song.albumId = Album.id;  SELECT * FROM Song CROSS JOIN Album ON Song.albumId = Album.id;": true,
		"SELECT * FROM foo WHERE id = 10;  SELECT * FROM foo WHERE id < 25;  SELECT * FROM foo WHERE id > 25;  SELECT * FROM foo WHERE id <= 25;  SELECT * FROM foo WHERE id >= 25;  SELECT * FROM foo WHERE id != 10;  SELECT * FROM foo where id IS 10;  SELECT * FROM foo where id IS NOT 25;  SELECT * FROM foo where id IS NOT DISTINCT FROM 10;  SELECT * FROM foo where id IS DISTINCT FROM 25;  SELECT * FROM foo where id IN (1, 2, 3);  SELECT * FROM foo where id = 10 AND id != 25;  SELECT * FROM foo where id = 10 OR id = 11; ":                                                                                                                                               true,
		"SELECT 名, 色 FROM 猫  SELECT * FROM 本  SELECT * FROM msg WHERE data = 'こんにちわ'  SELECT * FROM msg WHERE data = :データ": true,
		"CREATE TABLE foo (bar text, baz text); ALTER TABLE foo DROP COLUMN bar;":                                          true,
		`CREATE TABLE "Person" (
			"PersonId"	    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"GivenName"	    NVARCHAR(255) NOT NULL,
			"FamilyName"	NVARCHAR(255) NULL, -- #2915
			"Deleted"	    INTEGER DEFAULT NULL
		)`: true,
		"SELECT (VALUES(1),(2),(3),(4)); VALUES(1) UNION VALUES(2); VALUES(1),(2),(3) EXCEPT VALUES(2) VALUES(1),(2),(3) EXCEPT VALUES(1),(3); VALUES(1),(2),(3),(4) UNION ALL SELECT 5 LIMIT 99; VALUES(1),(2),(3),(4) UNION ALL SELECT 5 LIMIT 3; VALUES('0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ'); VALUES(1,8),(2,11),(3,1),(2,15),(1,4),(1,99) VALUES('g.txt', 1000000002, 5, '12345', 0) VALUES(9223372036854775807) ": true,
		"SELECT FIRST_VALUE(c2) OVER ( ORDER BY c3 ROWS UNBOUNDED PRECEDING EXCLUDE NO OTHERS ) AS c FROM t; SELECT FIRST_VALUE(c2) OVER ( ORDER BY c3 ROWS UNBOUNDED PRECEDING EXCLUDE CURRENT ROW ) AS c FROM t; SELECT FIRST_VALUE(c2) OVER ( ORDER BY c3 ROWS UNBOUNDED PRECEDING EXCLUDE GROUP ) AS c FROM t; SELECT FIRST_VALUE(c2) OVER ( ORDER BY c3 ROWS UNBOUNDED PRECEDING EXCLUDE TIES ) AS c FROM t; ":                                 true,
		"SELECT     Name,     printf('%,d',Bytes) Size,     FIRST_VALUE(Name) OVER (         ORDER BY Bytes     ) AS SmallestTrack FROM     tracks  WHERE     AlbumId = 1;  SELECT     AlbumId,     Name,     printf('%,d',Bytes) Size,     FIRST_VALUE(Name) OVER (         PARTITION BY AlbumId         ORDER BY Bytes DESC 		ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING     ) AS LargestTrack FROM     tracks; select * from COMPANY;  SELECT C.ID, C.NAME, C.AGE, D.DEPT         FROM COMPANY AS C, DEPARTMENT AS D         WHERE  C.ID = D.EMP_ID;  SELECT C.ID AS COMPANY_ID, C.NAME AS COMPANY_NAME, C.AGE, D.DEPT         FROM COMPANY AS C, DEPARTMENT AS D         WHERE  C.ID = D.EMP_ID;  SELECT  	Value,  	CUME_DIST()  	OVER ( 		ORDER BY value 	) CumulativeDistribution FROM 	CumeDistDemo; SELECT 	Val, 	DENSE_RANK () OVER (  		ORDER BY Val )  	ValRank  FROM 	DenseRankDemo;   SELECT 	CustomerId, 	Year, 	Total, 	LAG ( Total, 1, 0 ) OVER (  		ORDER BY Year  	) PreviousYearTotal  FROM 	CustomerInvoices  WHERE 	CustomerId = 4;  SELECT 	CustomerId, 	Year, 	Total, 	LAG ( Total,1,0) OVER (  		PARTITION BY CustomerId 		ORDER BY Year ) PreviousYearTotal  FROM 	CustomerInvoices;   SELECT     Name,     printf ( '%.f minutes',                      Milliseconds / 1000 / 60 )                      AS Length,     LAST_VALUE ( Name ) OVER (         ORDER BY Milliseconds          RANGE BETWEEN UNBOUNDED PRECEDING AND          UNBOUNDED FOLLOWING     ) AS LongestTrack  FROM     tracks  WHERE     AlbumId = 4;   SELECT     AlbumId,     Name,     printf ( '%.f minutes',                      Milliseconds / 1000 / 60 )                      AS Length,     LAST_VALUE ( Name ) OVER (         PARTITION BY AlbumId         ORDER BY Milliseconds DESC         RANGE BETWEEN UNBOUNDED PRECEDING AND          UNBOUNDED FOLLOWING     ) AS ShortestTrack  FROM     tracks; 	 SELECT 	CustomerId, 	Year, 	Total, 	LEAD ( Total,1,0) OVER ( ORDER BY Year ) NextYearTotal FROM 	CustomerInvoices  WHERE 	CustomerId = 1;  SELECT 	CustomerId, 	Year, 	Total, 	LEAD ( Total, 1, 0 ) OVER ( 		PARTITION BY CustomerId  		ORDER BY Year  	) NextYearTotal  FROM 	CustomerInvoices;  SELECT     AlbumId,     Name,     Milliseconds Length,     NTH_VALUE ( Name,2 ) OVER (         PARTITION BY AlbumId         ORDER BY Milliseconds DESC         RANGE BETWEEN              UNBOUNDED PRECEDING AND              UNBOUNDED FOLLOWING     ) AS SecondLongestTrack  FROM     tracks; 	 SELECT     Name,     Milliseconds Length,     NTH_VALUE(name,2) OVER (         ORDER BY Milliseconds DESC     ) SecondLongestTrack FROM     tracks; 	 SELECT 	Name, 	Milliseconds, 	NTILE ( 4 ) OVER (  		ORDER BY Milliseconds ) LengthBucket FROM 	tracks  WHERE 	AlbumId = 1; 	 SELECT 	AlbumId, 	Name, 	Milliseconds, 	NTILE ( 3 ) OVER (  		PARTITION BY AlbumId 		ORDER BY Bytes ) SizeBucket FROM 	tracks;  SELECT     Name,     Milliseconds,     PERCENT_RANK() OVER(          ORDER BY Milliseconds      ) LengthPercentRank FROM     tracks  WHERE     AlbumId = 1; 	 SELECT     Name,     Milliseconds,     printf('%.2f',PERCENT_RANK() OVER(          ORDER BY Milliseconds      )) LengthPercentRank FROM     tracks  WHERE     AlbumId = 1;  SELECT     AlbumId,     Name,     Bytes,     printf('%.2f',PERCENT_RANK() OVER(          PARTITION BY AlbumId         ORDER BY Bytes      )) SizePercentRank FROM     tracks; 	 SELECT 	Val, 	RANK () OVER (  		ORDER BY Val  	) ValRank FROM 	RankDemo;  SELECT 	Name, 	Milliseconds, 	RANK () OVER (  		ORDER BY Milliseconds DESC 	) LengthRank  FROM 	tracks;  SELECT 	Name, 	Milliseconds, 	AlbumId, 	RANK () OVER (  		PARTITION BY AlbumId 		ORDER BY Milliseconds DESC 	) LengthRank  FROM 	tracks; 	 	 SELECT  	*  FROM ( 	SELECT 		Name, 		Milliseconds, 		AlbumId, 		RANK () OVER (  			PARTITION BY AlbumId 			ORDER BY Milliseconds DESC 		) LengthRank  	FROM 		tracks )  WHERE  	LengthRank = 2; 	 SELECT     ROW_NUMBER () OVER (          ORDER BY Country      ) RowNum,     FirstName,     LastName,     country  FROM     customers;  SELECT     ROW_NUMBER () OVER (          PARTITION BY Country         ORDER BY FirstName     ) RowNum,     FirstName,     LastName,     country  FROM     customers;  SELECT * FROM (     SELECT         ROW_NUMBER () OVER (              ORDER BY FirstName         ) RowNum,         FirstName,         LastName,         Country      FROM         customers ) t WHERE      RowNum > 20 AND RowNum <= 30;   SELECT      Country,     FirstName,     LastName,     Amount FROM (     SELECT          Country,          FirstName,         LastName,         Amount,         ROW_NUMBER() OVER (             PARTITION BY country              ORDER BY Amount DESC         ) RowNum     FROM          Sales ) WHERE     RowNum = 1;": true,
		"INSERT INTO t1(b,c) VALUES(5,99);  INSERT INTO t1(b,c) VALUES(5,99) RETURNING b,c,a,rowid;  INSERT INTO t1 DEFAULT VALUES RETURNING *;  INSERT INTO t1 SELECT * FROM t2 RETURNING *;  INSERT INTO t1 VALUES(3) RETURNING a, (SELECT c FROM t2 WHERE t1.a=t2.b) AS x;  INSERT INTO t4(a,b,c) VALUES(1,22,33)   ON CONFLICT(a) DO UPDATE SET b=44;  INSERT INTO t4(a,b,c) VALUES(1,22,33)   ON CONFLICT(a) DO UPDATE SET b=44   RETURNING *;  INSERT INTO t4(a,b,c) VALUES(2,3,4),(4,5,6),(5,6,7)   ON CONFLICT(a) DO UPDATE SET b=100   RETURNING *, '|';  UPDATE t1 SET c='bellum' WHERE c='pax';  UPDATE t1 SET c='bellum' WHERE c='pax' RETURNING rowid, b, '|';  UPDATE t2 SET b='123' WHERE b='abc' RETURNING (SELECT b FROM t1);  UPDATE t2 SET b='123' WHERE b='abc' RETURNING (SELECT b FROM t1);  UPDATE t2 SET b='123' WHERE b='abc' RETURNING b;  UPDATE t1 SET id=id+y FROM t2 WHERE t1.id=t2.x RETURNING *, '|';  UPDATE t1 SET b=b+y FROM t2 WHERE t2.x=t1.a RETURNING t1.b;  UPDATE t1 AS alias SET b=alias.b+1000 RETURNING t1.b;  DELETE FROM t1 WHERE c='bellum' RETURNING rowid, *, '|';  DELETE FROM t1 RETURNING *;  DELETE FROM t3 WHERE f>100 RETURNING 'D', e, f;  DELETE FROM t1 RETURNING x, affinity(x); ": true,
	}

	for sql, expect := range sqls {
		result := IsSQLite(sql)
		if result != expect {
			t.Errorf("IsSQLite(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
