CREATE TABLE users
(
  id_user int(5) not null
  auto_increment,
  username varchar
  (255) not null,
  first_name varchar
  (255) not null,
  last_name varchar
  (255) not null,
  actif int
  (1) not null,
  created_at datetime not null,
  updated_at datetime not null,
  primary key
  (id_user)
);

  INSERT INTO users
    (username, first_name, last_name, actif)
  VALUES
    ("username@user.fr", "Test", "TEST", 1);
  INSERT INTO users
    (username, first_name, last_name, actif)
  VALUES
    ("username2@user.fr", "Test2", "TEST", 1);
  INSERT INTO users
    (username, first_name, last_name, actif)
  VALUES
    ("username3@user.fr", "Test3", "TEST", 0);