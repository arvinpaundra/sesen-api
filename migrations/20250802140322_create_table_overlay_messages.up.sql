BEGIN;

CREATE TABLE IF NOT EXISTS overlay_messages (
    id UUID PRIMARY KEY,
    overlay_id UUID NOT NULL,
    text_color VARCHAR(7) NOT NULL,
    background_color VARCHAR(7) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT overlay_messages_overlay_id_fkey FOREIGN KEY (overlay_id) REFERENCES overlays(id)
);

COMMIT;