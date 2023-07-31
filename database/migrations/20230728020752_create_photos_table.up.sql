CREATE TABLE photos (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  user_id bigint(20),
  path text,
  likes bigint(20),
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_photos_created_at (created_at),
  KEY idx_photos_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
