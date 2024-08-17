DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM "public"."migrations" WHERE "name" = 'users') THEN
        CREATE TABLE public.users
        (
            id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 ),
            email character varying(255) NOT NULL,
            first_name character varying(255) NOT NULL,
            last_name character varying(255) NOT NULL,
            password_hash character varying(255) NOT NULL,
            is_active boolean NOT NULL,
            created_at date NOT NULL,
            updated_at date,
            PRIMARY KEY (id)
        );

        ALTER TABLE public.users
            OWNER TO postgres;

        INSERT INTO public.migrations (name) VALUES ('users');
    END IF;
END $$;

