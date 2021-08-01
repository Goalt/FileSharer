CREATE TABLE IF NOT EXISTS `files_info` (
    token_id VARCHAR (50),
    file_name TEXT NOT NULL,
    file_name_origin TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    primary (token_id)
)