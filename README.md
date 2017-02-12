# testsServer
Server for psychological tests with filters.

CREATE TABLE IF NOT EXISTS tests (
  uid int(10) unsigned NOT NULL,
  name varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  description TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
  image varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=1;

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

CREATE TABLE IF NOT EXISTS filters (
  uid int(10) unsigned NOT NULL,
  title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  description TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
  tests TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=1;

ALTER TABLE  answers ADD FOREIGN KEY (uid) REFERENCES  questions (uid) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE  questions ADD FOREIGN KEY (uid) REFERENCES  tests (uid) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE  results ADD FOREIGN KEY (uid) REFERENCES  tests (uid) ON DELETE CASCADE ON UPDATE CASCADE;

How to add data:

curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Who R u from Hex?\", \"Description\": \"Who are you, waifu bro?\", \"image\": \"http://lorempixel.com/400/200/\", \"questions\": [{\"questionText\": \"Question 2\", \"answers\":[{\"answer\":\"Answer 2\", \"points\": 10}]}], \"results\": [{\"start\":15, \"end\":20, \"description\": \"Cat likes milk\", \"image\": \"http://lorempixel.com/400/200/\"}] }" http://localhost:8080/api/tests

curl -i -X POST -H "Content-Type: application/json" -d "{ \"title\": \"Winx Filter\", \"description\": \"Fei winx\", \"ids\":[1,2,3]}" http://localhost:8080/api/filters
