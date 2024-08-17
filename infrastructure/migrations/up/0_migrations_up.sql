DO $$
BEGIN
    
    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'migrations') THEN
        
        CREATE TABLE public.migrations
        (
            id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 ),
            name character varying(255),
            PRIMARY KEY (id)
        );

        
        ALTER TABLE public.migrations
            OWNER TO postgres;
    END IF;
END $$;