BEGIN;

CREATE TABLE IF NOT EXISTS overlay_qr (
    id UUID PRIMARY KEY,
    overlay_id UUID NOT NULL,
    code TEXT NOT NULL,
    qr_color VARCHAR(7) NOT NULL,
    background_color VARCHAR(7) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT overlay_qr_overlay_id_fkey FOREIGN KEY (overlay_id) REFERENCES overlays(id)
);

COMMIT;