CREATE TABLE companies (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  nome_fantasia VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  cnpj VARCHAR(14) NOT NULL,
  email VARCHAR(255) NOT NULL,
  KEY idx_companies_created_at (created_at),
  KEY idx_companies_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
