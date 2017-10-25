SET search_path TO my_schema, public;

alter user admin with superuser;

CREATE TABLE yamls(
  fileName varchar(50),
  lastUpdated Date
);

CREATE TABLE logs(
  fileName varchar(50),
  logDate Date,
  logContent varchar(10000));

INSERT INTO yamls VALUES ('hello.yml', '10-12-2015');
INSERT INTO logs VALUES ('hello.yml', '10-15-2017', 'this is log content');
select * from yamls;
select * from logs;
