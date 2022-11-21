CREATE KEYSPACE ffv WITH REPLICATION={'class':'SimpleStrategy', 'replication_factor': 1};
CREATE TABLE IF NOT EXISTS ffv.user(
    user_id uuid,
    mastadon_id text,
    mastadon_username text,
    PRIMARY KEY (user_id)
);
CREATE TYPE IF NOT EXISTS ffv.post(
    url text,
    tags set<text>,
    source text,
    created_at timestamp
);
CREATE TABLE IF NOT EXISTS ffv.tag_list(
    user_id uuid,
    tag_id uuid,
    source text,
    weight float,
    tag_name text,
    seen set<frozen<post>>,
    PRIMARY KEY (user_id, tag_id)
);
CREATE TABLE IF NOT EXISTS ffv.tag(
    tag_id uuid,
    name text,
    source text,
    buffer set<frozen<post>>,
    PRIMARY KEY(tag_id)
);