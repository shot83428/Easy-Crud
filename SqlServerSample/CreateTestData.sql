
DROP DATABASE Booker
DROP TABLE TestBook.Booked
DROP SCHEMA  TestBook
GO

CREATE DATABASE Booker
Go
CREATE SCHEMA TestBook;
GO

CREATE TABLE TestBook.Booked (
  Id       INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name     NVARCHAR(50),
  Train    NVARCHAR(50)
);
GO

INSERT INTO TestBook.Booked (Name, Train) VALUES
  (N'Jared',  N'2234'),
  (N'Nikita', N'1234'),
  (N'Tom',    N'7777');
GO

SELECT * FROM TestBook.Booked;
GO