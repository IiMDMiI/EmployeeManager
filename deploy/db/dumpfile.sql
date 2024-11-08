--
-- PostgreSQL database dump
--

-- Dumped from database version 14.13
-- Dumped by pg_dump version 14.13

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: company; Type: TABLE; Schema: public; Owner: griff
--

CREATE TABLE public.company (
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.company OWNER TO griff;

--
-- Name: company_id_seq; Type: SEQUENCE; Schema: public; Owner: griff
--

CREATE SEQUENCE public.company_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.company_id_seq OWNER TO griff;

--
-- Name: company_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: griff
--

ALTER SEQUENCE public.company_id_seq OWNED BY public.company.id;


--
-- Name: department; Type: TABLE; Schema: public; Owner: griff
--

CREATE TABLE public.department (
    name character varying(255) NOT NULL,
    phone character varying(50) NOT NULL,
    company_id integer NOT NULL
);


ALTER TABLE public.department OWNER TO griff;

--
-- Name: employee; Type: TABLE; Schema: public; Owner: griff
--

CREATE TABLE public.employee (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    surname character varying(255) NOT NULL,
    phone character varying(50) NOT NULL,
    company_id integer NOT NULL,
    department_name character varying(255) NOT NULL,
    pass_type character varying(50),
    pass_number character varying(50) NOT NULL
);


ALTER TABLE public.employee OWNER TO griff;

--
-- Name: employee_id_seq; Type: SEQUENCE; Schema: public; Owner: griff
--

CREATE SEQUENCE public.employee_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employee_id_seq OWNER TO griff;

--
-- Name: employee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: griff
--

ALTER SEQUENCE public.employee_id_seq OWNED BY public.employee.id;


--
-- Name: company id; Type: DEFAULT; Schema: public; Owner: griff
--

ALTER TABLE ONLY public.company ALTER COLUMN id SET DEFAULT nextval('public.company_id_seq'::regclass);


--
-- Name: employee id; Type: DEFAULT; Schema: public; Owner: griff
--

ALTER TABLE ONLY public.employee ALTER COLUMN id SET DEFAULT nextval('public.employee_id_seq'::regclass);


--
-- Data for Name: company; Type: TABLE DATA; Schema: public; Owner: griff
--

COPY public.company (id, name) FROM stdin;
1	First Company
2	Second Company
3	Third Company
\.


--
-- Data for Name: department; Type: TABLE DATA; Schema: public; Owner: griff
--

COPY public.department (name, phone, company_id) FROM stdin;
it	333333	1
hr	333444	1
sales	333555	1
it	444333	2
hr	444444	2
sales	444555	2
marketing	777333	3
sales	777555	3
\.


--
-- Data for Name: employee; Type: TABLE DATA; Schema: public; Owner: griff
--

COPY public.employee (id, name, surname, phone, company_id, department_name, pass_type, pass_number) FROM stdin;
1	John	Doe	+15550100	1	it	regular	1001
2	Jane	Smith	+15550101	1	hr	regular	1002
3	Emily	Johnson	+15550102	2	hr	regular	1003
4	Michael	Brown	+15550103	2	sales	regular	1004
5	Sarah	Davis	+15550104	3	marketing	regular	1005
6	David	Wilson	+15550105	3	sales	regular	1006
7	Linda	Martinez	+15550106	1	it	regular	1007
8	James	Garcia	+15550107	2	it	regular	1008
9	Patricia	Rodriguez	+15550108	3	sales	regular	1009
10	Robert	Lee	+15550109	1	sales	regular	1010
\.


--
-- Name: company_id_seq; Type: SEQUENCE SET; Schema: public; Owner: griff
--

SELECT pg_catalog.setval('public.company_id_seq', 3, true);


--
-- Name: employee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: griff
--

SELECT pg_catalog.setval('public.employee_id_seq', 10, true);


--
-- Name: company company_pkey; Type: CONSTRAINT; Schema: public; Owner: griff
--

ALTER TABLE ONLY public.company
    ADD CONSTRAINT company_pkey PRIMARY KEY (id);


--
-- Name: department department_pkey; Type: CONSTRAINT; Schema: public; Owner: griff
--

ALTER TABLE ONLY public.department
    ADD CONSTRAINT department_pkey PRIMARY KEY (company_id, name);


--
-- Name: employee employee_pkey; Type: CONSTRAINT; Schema: public; Owner: griff
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (id);


--
-- Name: employee uniq_employee; Type: CONSTRAINT; Schema: public; Owner: griff
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT uniq_employee UNIQUE (company_id, department_name, pass_number);


--
-- Name: department department_companyid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: griff
--

ALTER TABLE ONLY public.department
    ADD CONSTRAINT department_companyid_fkey FOREIGN KEY (company_id) REFERENCES public.company(id) ON DELETE CASCADE;


--
-- Name: employee employee_companyid_departmentname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: griff
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_companyid_departmentname_fkey FOREIGN KEY (company_id, department_name) REFERENCES public.department(company_id, name) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

