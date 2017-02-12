package main

import (
	_ "github.com/go-sql-driver/mysql"
)

/*

CREATE TABLE IF NOT EXISTS answers (
  uid int(10) unsigned NOT NULL,
  answer varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  points int(10) unsigned NOT NULL,
  INDEX (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS questions (
  uid int(10) unsigned NOT NULL,
  question TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
  INDEX (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS results (
  uid int(10) unsigned NOT NULL,
  image varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  description LONGTEXT COLLATE utf8mb4_unicode_ci NOT NULL,
  start int(10) unsigned NOT NULL,
  end int(10) unsigned NOT NULL,
  INDEX (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=1;

ALTER TABLE  answers ADD FOREIGN KEY (uid) REFERENCES  questions (uid) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE  questions ADD FOREIGN KEY (uid) REFERENCES  tests (uid) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE  results ADD FOREIGN KEY (uid) REFERENCES  tests (uid) ON DELETE CASCADE ON UPDATE CASCADE;
*/

type Test struct {
	Uid int64  `db:"uid" json:"-"`
	Name string `db:"name" json:"name"`
	Description  string `db:"description" json:"description"`
  Image string `db:"image" json:"image"`
  Questions []QuestionModel `json:"questions"`
  ResultInfos []ResultInfo `json:"results"`
}

type QuestionModel struct{
  Question string `db:"question" json:"question"`
  Answers []AnswerModel `json:"answers"`
}

type AnswerModel struct{
  Answer string `db:"answer" json:"answer"`
  Points int `db:"points" json:"points"`
}

type ResultInfo struct{
  Uid int64 `db:"uid" json:"-"`
  Start int `db:"start" json:"start"`
  End int `db:"end" json:"end"`
  Description string `db:"description" json:"description"`
  Image string `db:"image" json:"image"`
}

type Filter struct{
  Uid int64 `db:"uid" json:"-"`
  Title string `db:"title" json:"title"`
  Description string `db:"description" json:"description"`
  TestsString string `db:"tests" json:"-"`
  Tests []int64 `db:"-" json:"ids"`
}
