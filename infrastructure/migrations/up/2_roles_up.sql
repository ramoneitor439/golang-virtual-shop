DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM "public"."migrations" WHERE "name" = 'roles') THEN
        CREATE TABLE public.roles
        (
            id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 ),
            name character varying(255) NOT NULL,
            normalized_name character varying(255) NOT NULL,
            PRIMARY KEY (id)
        );

        ALTER TABLE public.roles
            OWNER TO postgres;

        INSERT INTO public.migrations (name) VALUES ('roles');
    END IF;
END $$;