CREATE TABLE rounds (
  id varchar(64) primary key UNIQUE NOT NULL,
  match_id varchar(64) NOT NULL,
  count integer NOT NULL,
  mode integer NOT NULL,
  time_left integer,
  percent_diff integer,
  points_taken integer,
  win integer NOT NULL DEFAULT 0,
  notes text
);

GRANT SELECT, INSERT, UPDATE, DELETE ON rounds TO app;
CREATE INDEX index_matches_mode ON rounds USING btree (mode);
CREATE INDEX index_matches_time_left ON rounds USING btree (time_left);
CREATE INDEX index_matches_percent_diff ON rounds USING btree (percent_diff);
CREATE INDEX index_matches_points_taken ON rounds USING btree (points_taken);
