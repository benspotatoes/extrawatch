CREATE TABLE rounds (
  id varchar(64) primary key UNIQUE NOT NULL,
  count integer NOT NULL,
  mode integer NOT NULL,
  time_left integer,
  percent_diff integer,
  points_taken integer,
  notes text,
)

GRANT SELECT, INSERT, UPDATE, DELETE ON matches TO app;
CREATE INDEX index_matches_mode ON matches USING btree (mode);
CREATE INDEX index_matches_time_left ON matches USING btree (time_left);
CREATE INDEX index_matches_percent_diff ON matches USING btree (percent_diff);
CREATE INDEX index_matches_points_taken ON matches USING btree (points_taken);
