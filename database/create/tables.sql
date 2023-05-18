PRAGMA foreign_keys=OFF;
CREATE TABLE users (
    userid          INTEGER PRIMARY KEY,    -- User ID
    username        TEXT NOT NULL,          -- User name
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
CREATE TABLE grids (
    gridid          INTEGER PRIMARY KEY,    -- Grid ID
    userid          INTEGER NOT NULL,       -- User who owns the grid
    gridname        TEXT NOT NULL,          -- Grid name
    created         TEXT,                   -- Datetime when created
    modified        TEXT                    -- Datetime last modified
    n               INTEGER                 -- Grid size (width and height)
);
CREATE TABLE cells (
    gridid          INTEGER,                -- Grid ID
    r               INTEGER,                -- Row number (1, 2, ..., n)
    c               INTEGER,                -- Column number (1, 2, ..., n)
    letter          TEXT,                   -- Cell value character
    PRIMARY KEY (gridid, r, c),
    FOREIGN KEY (gridid) REFERENCES grids (gridid) ON DELETE CASCADE
);
CREATE TABLE words (
    gridid          INTEGER,                 -- Grid ID
    r               INTEGER,                -- Row number (1, 2, ..., n)
    c               INTEGER,                -- Column number (1, 2, ..., n)
    dir             TEXT,                   -- Direction (Across="A",Down="D")
    length          INTEGER,                -- Length of word
    clue            TEXT,                   -- Text of clue
    PRIMARY KEY (gridid, r, c, dir),
    FOREIGN KEY (gridid) REFERENCES grids (gridid) ON DELETE CASCADE
);

