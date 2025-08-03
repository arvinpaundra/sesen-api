BEGIN;

CREATE TABLE IF NOT EXISTS overlays (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    ringtone_url TEXT NOT NULL,
    is_tts_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    is_nsfw_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT overlays_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);

COMMIT;