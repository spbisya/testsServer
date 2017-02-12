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
