CREATE TABLE players (
  id varchar(64) primary key UNIQUE NOT NULL,
  name varchar(64) NOT NULL
);

GRANT SELECT, INSERT, UPDATE, DELETE ON players TO app;
CREATE INDEX index_players_name ON players USING btree (name);

CREATE TABLE player_rounds (
  player_id varchar(64) NOT NULL,
  round_id varchar(64) NOT NULL,
  hero integer NOT NULL
);

GRANT SELECT, INSERT, UPDATE, DELETE ON player_rounds TO app;
CREATE INDEX index_player_round_id ON player_rounds USING btree (round_id);
CREATE INDEX index_player_heroes ON player_rounds USING btree (hero);
