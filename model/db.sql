CREATE TABLE tantan;

\c tantan;

CREATE TYPE relationship AS (
    user_id VARCHAR(64),
    state CHAR(64),

)