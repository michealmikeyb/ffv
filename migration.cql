CREATE KEYSPACE ffv WITH REPLICATION={'class':'SimpleStrategy', 'replication_factor': 1};
CREATE TABLE IF NOT EXISTS ffv.user(
    user_id uuid,
    mastodon_id text,
    mastodon_username text,
    PRIMARY KEY (mastodon_id)
);
CREATE TYPE IF NOT EXISTS ffv.post(
    url text,
    tags set<text>,
    source text,
    created_at timestamp,
    author text,
    likes int,
    content text
);
CREATE TABLE IF NOT EXISTS ffv.tag_list(
    user_id uuid,
    source text,
    weight int,
    tag_name text,
    seen set<frozen<post>>,
    PRIMARY KEY (user_id, tag_name, source)
);
CREATE TABLE IF NOT EXISTS ffv.tag(
    name text,
    source text,
    buffer list<frozen<post>>,
    PRIMARY KEY(name, source)
);