CREATE TABLE matches (
  id varchar(64) primary key UNIQUE NOT NULL,
  map integer NOT NULL,
  win integer NOT NULL DEFAULT 0,
  rank_diff integer NOT NULL,
  ending_rank integer NOT NULL,
  placement boolean NOT NULL DEFAULT false,
  played_on timestamp without time zone NOT NULL
);

GRANT SELECT, INSERT, UPDATE, DELETE ON matches TO app;
CREATE INDEX index_matches_map ON matches USING btree (map);
CREATE INDEX index_matches_win ON matches USING btree (win);
CREATE INDEX index_matches_played_on ON matches USING btree (played_on);
