CREATE TABLE IF NOT EXISTS journal (
    time TIMESTAMP,
    login TEXT,
    content TEXT
);
CREATE INDEX IF NOT EXISTS journal_time ON journal(time);
