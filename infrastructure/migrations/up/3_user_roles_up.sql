DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM public.migrations WHERE name = 'user_roles') THEN
            CREATE TABLE public.user_roles
                (
                    user_id integer,
                    role_id integer,
                    CONSTRAINT user_role PRIMARY KEY (user_id, role_id),
                    CONSTRAINT "user" FOREIGN KEY (user_id)
                        REFERENCES public.users (id) MATCH SIMPLE
                        ON UPDATE CASCADE
                        ON DELETE CASCADE
                        NOT VALID,
                    CONSTRAINT role FOREIGN KEY (role_id)
                        REFERENCES public.roles (id) MATCH SIMPLE
                        ON UPDATE NO ACTION
                        ON DELETE NO ACTION
                        NOT VALID
                );

            ALTER TABLE IF EXISTS public.user_roles
                OWNER to postgres;

            INSERT INTO public.migrations (name) VALUES ('user_roles');
        END IF;
END $$;