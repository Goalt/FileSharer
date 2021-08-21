CREATE TABLE IF NOT EXISTS `file_info` (
    token_id VARCHAR(50) NOT NULL,
    file_name TEXT NOT NULL,
    file_name_origin BLOB NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (token_id)
)