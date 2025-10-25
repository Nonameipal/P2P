ALTER TABLE bookings
    DROP CONSTRAINT IF EXISTS bookings_item_id_fkey,
    ALTER COLUMN item_id TYPE UUID USING item_id::text::uuid,
    ADD CONSTRAINT bookings_item_id_fkey FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE;