
CREATE TABLE `bw_sql_Log_Detail` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `sqls` text,
  `mapperMethod` varchar(100) DEFAULT NULL,
  `parameters` varchar(1000) DEFAULT NULL,
  `costTime` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4
CREATE TABLE `bw_sql_log_counts` (
  `mapperMethod` varchar(100) NOT NULL,
  `counts` int(10) DEFAULT NULL,
  PRIMARY KEY (`mapperMethod`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4






CREATE TABLE T_sql_log_counts(
  mapperMethod varchar2(100) PRIMARY key,
  counts int 

)


CREATE TABLE T_sql_Log_Detail (
  id int PRIMARY key,
  sqls varchar2(4000),
  mapperMethod varchar2(100) ,
  parameters varchar(1000) ,
  costTime int
) 


CREATE SEQUENCE T_sql_Log_Detail_SEQ
START WITH 1
INCREMENT BY 1
MAXVALUE 99999999
MINVALUE 1
NOCYCLE
nocache