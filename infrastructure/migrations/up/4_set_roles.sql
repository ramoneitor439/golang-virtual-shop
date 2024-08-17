DO $$
    BEGIN 
        IF NOT EXISTS (SELECT 1 FROM public.migrations WHERE name = 'set_roles') THEN
            INSERT INTO public.roles (name, normalized_name) VALUES ('Admin', 'ADMIN'),('User', 'USER');
            INSERT INTO public.migrations (name) VALUES ('set_roles');
        END IF;
END $$;