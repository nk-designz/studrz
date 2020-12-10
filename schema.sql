CREATE TABLE public.consumes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    questionair_result_id bigint,
    name text,
    rate smallint
);

CREATE TABLE public.questionair_results (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    score smallint,
    filling_date timestamp with time zone
);

CREATE TABLE public.statistics (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    calc text,
    value numeric,
    key text,
    findings bigint
);

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text,
    password text,
    birthday timestamp with time zone,
    role smallint
);

ALTER TABLE ONLY public.consumes
    ADD CONSTRAINT consumes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.questionair_results
    ADD CONSTRAINT questionair_results_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.statistics
    ADD CONSTRAINT statistics_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);

CREATE INDEX idx_consumes_deleted_at ON public.consumes USING btree (deleted_at);
CREATE INDEX idx_questionair_results_deleted_at ON public.questionair_results USING btree (deleted_at);
CREATE INDEX idx_questionair_results_score ON public.questionair_results USING btree (score);
CREATE INDEX idx_statistics_deleted_at ON public.statistics USING btree (deleted_at);
CREATE INDEX idx_users_birthday ON public.users USING btree (birthday);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);

ALTER TABLE ONLY public.consumes
    ADD CONSTRAINT fk_consumes_questionair_result FOREIGN KEY (questionair_result_id) REFERENCES public.questionair_results(id);

ALTER TABLE ONLY public.questionair_results
    ADD CONSTRAINT fk_questionair_results_user FOREIGN KEY (user_id) REFERENCES public.users(id);
