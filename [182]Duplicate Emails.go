//Write a SQL query to find all duplicate emails in a table named Person. 
//
// 
//+----+---------+
//| Id | Email   |
//+----+---------+
//| 1  | a@b.com |
//| 2  | c@d.com |
//| 3  | a@b.com |
//+----+---------+
// 
//
// For example, your query should return the following for the above table: 
//
// 
//+---------+
//| Email   |
//+---------+
//| a@b.com |
//+---------+
// 
//
// Note: All emails are in lowercase. 
//

# select distinct Email from Person as p1

# where exists (
#     select Email from Person as p2
#     where p1.Email = p2.Email
#     limit 1,1
#     )
SELECT DISTINCT a.Email FROM Person a
LEFT JOIN (SELECT Id, Email from Person GROUP BY Email) b
ON (a.email = b.email) AND (a.Id = b.Id)
WHERE b.Email IS NULL