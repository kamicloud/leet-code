//Given a Weather table, write a SQL query to find all dates' Ids with higher temperature compared to its previous (yesterday's) dates. 
//
// 
//+---------+------------------+------------------+
//| Id(INT) | RecordDate(DATE) | Temperature(INT) |
//+---------+------------------+------------------+
//|       1 |       2015-01-01 |               10 |
//|       2 |       2015-01-02 |               25 |
//|       3 |       2015-01-03 |               20 |
//|       4 |       2015-01-04 |               30 |
//+---------+------------------+------------------+
// 
//
// For example, return the following Ids for the above Weather table: 
//
// 
//+----+
//| Id |
//+----+
//|  2 |
//|  4 |
//+----+
// 
//


select Id from Weather as a
where exists (
	select * from Weather as b where date_sub(a.RecordDate , interval 1 day) = b.RecordDate and a.Temperature > b.Temperature
	)

