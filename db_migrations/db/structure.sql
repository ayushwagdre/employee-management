SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP INDEX IF EXISTS public.index_merchant_partner_configs_on_merchant_id_and_client_id;
DROP INDEX IF EXISTS public.index_merchant_partner_configs_on_merchant_id;
DROP INDEX IF EXISTS public.index_merchant_partner_configs_on_client_id;
ALTER TABLE IF EXISTS ONLY public.schema_migrations DROP CONSTRAINT IF EXISTS schema_migrations_pkey;
ALTER TABLE IF EXISTS ONLY public.merchant_partner_configs DROP CONSTRAINT IF EXISTS merchant_partner_configs_pkey;
ALTER TABLE IF EXISTS ONLY public.ar_internal_metadata DROP CONSTRAINT IF EXISTS ar_internal_metadata_pkey;
DROP TABLE IF EXISTS public.schema_migrations;
DROP TABLE IF EXISTS public.merchant_partner_configs;
DROP TABLE IF EXISTS public.ar_internal_metadata;
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS pgcrypto;
-- *not* dropping schema, since initdb creates it
--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: ar_internal_metadata; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ar_internal_metadata (
    key character varying NOT NULL,
    value character varying,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


--
-- Name: merchant_partner_configs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.merchant_partner_configs (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    merchant_id uuid,
    client_id uuid,
    app_configs jsonb DEFAULT '{}'::jsonb,
    created_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying NOT NULL
);


--
-- Name: ar_internal_metadata ar_internal_metadata_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ar_internal_metadata
    ADD CONSTRAINT ar_internal_metadata_pkey PRIMARY KEY (key);


--
-- Name: merchant_partner_configs merchant_partner_configs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.merchant_partner_configs
    ADD CONSTRAINT merchant_partner_configs_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: index_merchant_partner_configs_on_client_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_merchant_partner_configs_on_client_id ON public.merchant_partner_configs USING btree (client_id);


--
-- Name: index_merchant_partner_configs_on_merchant_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_merchant_partner_configs_on_merchant_id ON public.merchant_partner_configs USING btree (merchant_id);


--
-- Name: index_merchant_partner_configs_on_merchant_id_and_client_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_merchant_partner_configs_on_merchant_id_and_client_id ON public.merchant_partner_configs USING btree (merchant_id, client_id);


--
-- PostgreSQL database dump complete
--

SET search_path TO "$user", public;

INSERT INTO "schema_migrations" (version) VALUES
('20230726000000'),
('20230727000000');


