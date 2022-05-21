
CREATE TABLE users (  
  name VARCHAR ( 50 ) UNIQUE NOT NULL,
  password VARCHAR (50) NOT NULL,
  coockie jsonb NULL
);