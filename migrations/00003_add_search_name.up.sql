ALTER TABLE IF EXISTS public.players
    ADD COLUMN search_name character varying;

COMMENT ON COLUMN public.players.search_name
    IS '検索用選手名';