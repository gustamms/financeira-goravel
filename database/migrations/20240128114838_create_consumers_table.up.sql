CREATE TABLE consumers (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  name VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  cpf VARCHAR(11) NOT NULL,
  KEY idx_consumers_created_at (created_at),
  KEY idx_consumers_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
