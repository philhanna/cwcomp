PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE users (
    userid          INTEGER PRIMARY KEY,    -- User ID
    username        TEXT NOT NULL UNIQUE,   -- User name
    password        BLOB NOT NULL,          -- Encrypted password
    created         TEXT,                   -- Datetime this user was created
    email           TEXT,                   -- User email
    confirmed       TEXT,                   -- Datetime email confirmed
    author_name     TEXT,                   -- Name as author
    address_line_1  TEXT,                   -- Address line 1
    address_line_2  TEXT,                   -- Address line 2, if required
    address_city    TEXT,                   -- Author city
    address_state   TEXT,                   -- Author state code
    address_zip     TEXT                    -- Author zip code
);
INSERT INTO users VALUES(
    1,
    'saspeh',
    X'5fab5f43d0b2f97526e7daacb1e42ffd178e2067c59a0036d0d984a3bec289fb',
    '1953-12-04T08:30:00',
    'ph1204@gmail.com',
    '2020-07-07T03:33:01.264065',
    'Phil and Mary Hanna',
    '7500 Cadbury Court',
    'Apt. 202',
    'Raleigh',
    'NC',
    '27615'
    );

CREATE TABLE puzzles (
    id              INTEGER PRIMARY KEY,    -- Puzzle ID
    userid          INTEGER NOT NULL,       -- User who owns the puzzle
    puzzlename      TEXT NOT NULL UNIQUE,   -- Puzzle name
    created         TEXT,                   -- Datetime when created
    modified        TEXT,                   -- Datetime last modified
    n               INTEGER                 -- Puzzle size (width and height)
);
CREATE TABLE cells (
    id              INTEGER,                -- Puzzle ID
    r               INTEGER,                -- Row number (1, 2, ..., n)
    c               INTEGER,                -- Column number (1, 2, ..., n)
    letter          TEXT,                   -- Cell value character
    PRIMARY KEY (id, r, c),
    FOREIGN KEY (id) REFERENCES puzzles (id) ON DELETE CASCADE
);
CREATE TABLE words (
    id              INTEGER,                -- Puzzle ID
    r               INTEGER,                -- Row number (1, 2, ..., n)
    c               INTEGER,                -- Column number (1, 2, ..., n)
    dir             TEXT,                   -- Direction (Across="A",Down="D")
    length          INTEGER,                -- Length of word
    clue            TEXT,                   -- Text of clue
    PRIMARY KEY (id, r, c, dir),
    FOREIGN KEY (id) REFERENCES puzzles (id) ON DELETE CASCADE
);
COMMIT;