CREATE USER telegrambot WITH
  SUPERUSER
  CREATEDB
  CREATEROLE
  NOINHERIT
  LOGIN
  ENCRYPTED PASSWORD 'botpassword';


CREATE DATABASE telegrambot WITH
  OWNER = telegrambot
  TEMPLATE = template0
  ENCODING = 'UTF-8'
  LC_COLLATE = 'en_US.UTF-8'
  LC_CTYPE = 'en_US.UTF-8'
  IS_TEMPLATE = false;

\connect telegrambot;

GRANT ALL PRIVILEGES ON DATABASE "telegrambot" to telegrambot;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO telegrambot;
GRANT ALL ON schema public TO telegrambot;
