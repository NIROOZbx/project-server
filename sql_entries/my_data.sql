--
-- PostgreSQL database dump
--

\restrict dyOdNTWV1hXI4SzxTwNfXekKjic6GdvxJ9NdAvfYRicizuZPaTxRpF9czmiDUGL

-- Dumped from database version 18.0
-- Dumped by pg_dump version 18.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (3, '2025-11-17 20:30:32.501139+05:30', '2025-12-10 10:53:12.795091+05:30', NULL, 'Nirooz vp', '$2a$10$gRscIJn7i3dsYpl0DhC5jeUqQkR07zwOCNB3NfUOVGcC2Aq/5ao2y', 'user', 'huyboom463@gmail.com', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1765188947/products/mydg5q7iheyupvpy8do2.avif', false, '', 23, true);
INSERT INTO public.users VALUES (22, '2025-12-08 15:08:55.733244+05:30', '2025-12-11 23:17:02.951288+05:30', NULL, 'rock', '$2a$10$pC/ohwObb72GnTjR2NeCXedr4SeCFoQTQ9JNM1HyuzMDNk8zZAvwO', 'user', 'kixevi4173@crsay.com', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1765300901/products/dbet0i1zc8ydzktxt9qf.png', false, '', 28, true);
INSERT INTO public.users VALUES (4, '2025-11-21 20:56:13.499675+05:30', '2025-12-08 10:29:37.466491+05:30', NULL, 'rahees', '$2a$10$YljgYsiNANIIs0PCo4FIZeF/pm../h1931v.ofkl4uv.h3VSU2lCS', 'user', 'nemim60558@feralrex.com', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1765169955/products/eacr4nu8hjzmkhin9iyb.png', false, '', 3, true);
INSERT INTO public.users VALUES (5, '2025-11-25 11:36:31.938105+05:30', '2025-11-28 15:37:06.203883+05:30', NULL, 'Admin', '$2a$10$auAISc6h/bDOFQzdaDepv.3VdV50s1tm1312QU6PZK3jsIVAvG76K', 'admin', 'thestoryguy93@gmail.com', '', false, '', 6, true);
INSERT INTO public.users VALUES (21, '2025-12-08 14:21:06.426095+05:30', '2025-12-08 14:26:08.782979+05:30', NULL, 'suhail', '$2a$10$woV6DqiCuHQp7XqHxIEhEeGRQcTh.uxozCICWuAzJHgsl19G/1.Je', 'user', 'lohono9446@bialode.com', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1765184012/products/vms1esojjxg7vuufwsro.jpg', false, '', 2, true);
INSERT INTO public.users VALUES (20, '2025-12-02 23:35:40.346461+05:30', '2025-12-09 20:58:36.260594+05:30', NULL, 'danny', '$2a$10$AA0jjD8Nt1rgRjoK9gAo8Of7kRCe/4IBkncc07iymrI42MVuS33lO', 'user', 'gafas47830@bialode.com', '', false, '', 3, true);
INSERT INTO public.users VALUES (19, '2025-12-02 22:33:19.918516+05:30', '2025-12-08 15:23:47.035029+05:30', NULL, 'nirooz', '$2a$10$tFot2ym2V0b6670ZkRPh5O3za30SjrET.IE1/RHE1sIbKqtE4vEO6', 'user', 'harry@gmail.com', '', false, '', 1, false);


--
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.addresses VALUES (9, '2025-12-08 15:11:02.750049+05:30', '2025-12-09 23:24:04.202406+05:30', NULL, 22, 'ksjnfkjsd', '3423424234', 'fsdf', 'fdsfs', 'fdsfsfs', '423423', 'fdsf', false);
INSERT INTO public.addresses VALUES (14, '2025-12-09 23:20:11.149203+05:30', '2025-12-09 23:24:04.202406+05:30', NULL, 22, 'kjdsfnskjfdnksjdfnjkdsf', '1231234123', 'sokfnsokf', 'on', 'skfsklfmskdf', '434232', 'nfo', true);
INSERT INTO public.addresses VALUES (15, '2025-12-09 23:21:03.642486+05:30', '2025-12-09 23:24:04.202406+05:30', NULL, 22, 'skdnskfd', '1247824378', 'ojfnj', 'fnnfoof', 'ofn', '424323', 'n', false);
INSERT INTO public.addresses VALUES (2, '2025-11-24 09:49:53.086297+05:30', '2025-12-06 22:04:16.707851+05:30', NULL, 3, 'Toronto Office', '+1 416-555-0100', '100 King St W', 'Toronto', 'Ontario', 'M5X 1A9', 'Canada', true);
INSERT INTO public.addresses VALUES (3, '2025-11-24 09:49:56.466063+05:30', '2025-12-06 22:04:16.707851+05:30', NULL, 3, 'Secure Mailbox', '+1 512-555-1234', 'PO Box 98765', 'Austin', 'Texas', '73301', 'USA', false);
INSERT INTO public.addresses VALUES (4, '2025-11-24 09:50:02.154688+05:30', '2025-12-06 22:04:16.707851+05:30', NULL, 3, 'Alice''s Place', '+1 702-555-7777', '777 Lucky Avenue', 'Las Vegas', 'Nevada', '89109', 'USA', false);
INSERT INTO public.addresses VALUES (5, '2025-11-24 09:50:06.73795+05:30', '2025-12-06 22:04:16.707851+05:30', NULL, 3, 'Beach House', '+1 305-555-9999', '404 Ocean Drive', 'Miami', 'Florida', '33139', 'USA', false);
INSERT INTO public.addresses VALUES (1, '2025-11-24 09:49:30.67323+05:30', '2025-12-06 22:04:16.707851+05:30', '2025-12-08 09:51:28.434079+05:30', 3, 'New Main Apartment', '+1 212-555-5000', '5005 Fifth Avenue, Penthouse B', 'New York', 'New York', '10022', 'USA', false);
INSERT INTO public.addresses VALUES (6, '2025-12-08 10:17:41.014623+05:30', '2025-12-08 10:17:41.014623+05:30', NULL, 3, 'John Doe', '9876543210', 'Flat 402, Sunshine Apartments, MG Road', 'Bangalore', 'Karnataka', '560001', 'India', false);
INSERT INTO public.addresses VALUES (7, '2025-12-08 10:24:06.144088+05:30', '2025-12-08 10:28:47.848568+05:30', NULL, 4, 'Alice Smith', '1234567890', '45 West Street, Suite 2B', 'NY', 'New York', '10001', 'USA', true);
INSERT INTO public.addresses VALUES (8, '2025-12-08 10:26:21.182535+05:30', '2025-12-08 10:28:47.848568+05:30', NULL, 4, 'John Doe', '9876543210', 'Flat 402, Sunshine Apartments, MG Road', 'Bangalore', 'Karnataka', '560001', 'India', false);
INSERT INTO public.addresses VALUES (10, '2025-12-09 23:09:41.361028+05:30', '2025-12-09 23:09:41.361028+05:30', '2025-12-09 23:16:38.652699+05:30', 22, 'kmsdfs', '1234563476', 'fsdfsdf', 'fsdfs', 'sdfsfds', '756756', 'fdsfsd', false);
INSERT INTO public.addresses VALUES (11, '2025-12-09 23:14:38.572362+05:30', '2025-12-09 23:14:38.572362+05:30', '2025-12-09 23:19:52.608703+05:30', 22, 'fsmfsd', '5343453453', 'sfsdf', 'fsdfs', 'fdsfsd', '345345', 'fdssd', false);
INSERT INTO public.addresses VALUES (12, '2025-12-09 23:16:13.975038+05:30', '2025-12-09 23:16:13.975038+05:30', '2025-12-09 23:22:28.846505+05:30', 22, 'sdfsmf s', '1231321231', 'sdmf smfd', 'smdf smfd', 'sdmnv skmdv', '123456', 'sflkskmdf', false);
INSERT INTO public.addresses VALUES (16, '2025-12-09 23:22:48.158248+05:30', '2025-12-09 23:22:48.158248+05:30', '2025-12-09 23:23:31.523203+05:30', 22, 'kjandfknsdfkjs', '2345345345', 'n', 'nij', 'jnj', '213123', 'nj', false);
INSERT INTO public.addresses VALUES (13, '2025-12-09 23:17:53.589581+05:30', '2025-12-09 23:17:53.589581+05:30', '2025-12-09 23:24:02.825501+05:30', 22, 'dsfsm', '1234567898', 'fkdslmflsmdf', 'skfmsklfd', 'sfms fmsfd', '423423', 'slkfmsfd', false);


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.products VALUES (2, 'FC Barcelona Home Jersey 2018-19', 'FC Barcelona', 'Laliga', 2019, 10, 65.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471062/barcelona-2018-19-home_owcmqp.jpg', 'Home', 'The official FC Barcelona home jersey for the 2018-19 season. Made with high-quality, lightweight polyester for superior ventilation and fit. Includes the iconic Blaugrana stripes, official crest, and sponsor details for a true fan experience.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (11, 'Chelsea Away Jersey 2016-17', 'Chelsea', 'Premier League', 2016, 10, 72.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471059/chelsea-16-17-away_vkvivy.jpg', 'Away', 'The official Chelsea away jersey for the 2016-17 season. Vibrant blue with abstract patterns, using Dri-FIT material for moisture control. Includes the iconic lion badge and Nike Swoosh for authentic Blues pride.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (14, 'Germany Home Jersey 1990', 'Germany', 'International', 1990, 10, 95.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471062/germany-26-home_sveypo.png', 'Home', 'The official Germany home jersey for the 1990 season, a retro classic from their World Cup victory. Reimagined with modern soft-touch fabric while preserving the original design. Features historic eagle and adidas stripes.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (17, 'France Home Jersey 2022', 'France', 'International', 2022, 10, 88.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471061/france-2022_qnrhhm.jpg', 'Home', 'The official France home jersey for the 2022 season. Navy blue with rooster emblem, crafted from recycled materials for eco-friendliness. Nike''s Dri-FIT ensures comfort during high-stakes games.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (5, 'Manchester United 25-26', 'Manchester United', 'Premier League', 2026, 10, 68.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1765108369/products/oohv0gha6r2zhcqlpodo.avif', 'Third', 'The official Manchester United third jersey for the 2011-12 season. Retro-inspired with modern fabric technology for breathability and flexibility', '0001-01-01 05:53:28+05:53:28', '2025-12-07 17:24:33.441225+05:30', '2025-12-07 17:24:43.770513+05:30', NULL);
INSERT INTO public.products VALUES (10, 'Real Madrid Home Jersey 2017-18', 'Real Madrid', 'Laliga', 2017, 10, 78.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471066/real-madrid-17-18-home_dhqamj.jpg', 'Home', 'The official Real Madrid home jersey for the 2017-18 season. Classic white with gradient fade, constructed with ClimaCool technology for temperature regulation. Embroidered Real Madrid crest and sponsor for a luxurious feel.', NULL, '2025-12-06 12:47:41.002196+05:30', NULL, NULL);
INSERT INTO public.products VALUES (8, 'Inter Milan Away Jersey 2024-25', 'Inter Milan', 'Serie A', 2024, 11, 91.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471064/inter-24-25-away_u6zjs2.jpg', 'Away', 'The official Inter Milan away jersey for the 2024-25 season. Crafted from recycled polyester for sustainability and performance, with snake-print accents nodding to the club''s heritage. Includes Nerazzurri crest for true authenticity.', NULL, '2025-12-10 01:34:58.60995+05:30', NULL, NULL);
INSERT INTO public.products VALUES (6, 'Arsenal Away Jersey 2020-21', 'Arsenal', 'Premier League', 2020, 9, 75.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471059/arsenal-20-21-away_anihio.jpg', 'Away', 'The official Arsenal away jersey for the 2020-21 season. Designed with aerodynamic panels and anti-odor treatment for all-day freshness. Captures the Gunners'' spirit through dynamic patterns and official team details.', NULL, '2025-12-10 09:39:04.431448+05:30', NULL, NULL);
INSERT INTO public.products VALUES (3, 'Germany Home Jersey 2014', 'Germany', 'International', 2014, 7, 84.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471062/germany-14_f9s5ak.jpg', 'Home', 'The official Germany home jersey for the 2014 season, commemorating their World Cup triumph. Constructed from durable, quick-dry material with ergonomic seams for unrestricted movement. Adorned with the classic black-red-gold colors and embroidered eagle emblem.', NULL, '2025-12-11 23:32:24.194341+05:30', NULL, NULL);
INSERT INTO public.products VALUES (13, 'FC Barcelona Home Jersey 2024-25', 'FC Barcelona', 'Laliga', 2024, 9, 94.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471062/barcelona-2024-25-home_pkr1z6.jpg', 'Home', 'The official FC Barcelona home jersey for the 2024-25 season. Updated Blaugrana with golden accents, featuring AEROREADY fabric for sweat absorption. Official crest and Spotify sponsor make it a must-have for Culers.', NULL, '2025-12-10 09:01:18.22399+05:30', NULL, NULL);
INSERT INTO public.products VALUES (4, 'Real Madrid Third Jersey 2016-17', 'Real Madrid', 'Laliga', 2016, 11, 93.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471066/real-madrid-16-17-third-kit_vuca9l.jpg', 'Third', 'The official Real Madrid third jersey for the 2016-17 season. Featuring a bold, alternative design in vibrant hues, built with advanced fabric for sweat management and comfort. Includes authentic club insignia and subtle performance enhancements.', NULL, '2025-12-10 10:39:17.267449+05:30', NULL, NULL);
INSERT INTO public.products VALUES (1, 'Tottenham Home Jersey 2025-26', 'Tottenham Hotspur', 'Premier League', 2025, 8, 88.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471068/Tott-25-26-home_cctozc.jpg', 'Home', 'The official Tottenham Hotspur home jersey for the 2025-26 season. Crafted from premium breathable fabric with moisture-wicking technology, it ensures maximum comfort during intense matches or everyday wear. Features authentic team badges, embroidered sponsor logos, and subtle design elements inspired by the club''s rich history.', NULL, '2025-12-10 09:55:48.52807+05:30', NULL, NULL);
INSERT INTO public.products VALUES (15, 'AC Milan Home Jersey 2025-26', 'AC Milan', 'Serie A', 2025, 9, 92.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/milan-25-26-home_fii9xo.jpg', 'Home', 'The official AC Milan home jersey for the 2025-26 season. Red and black stripes with subtle sheen, built with HEAT.RDY for cooling. Includes the Rossoneri crest and Puma branding for timeless elegance.', NULL, '2025-12-10 11:56:04.060988+05:30', NULL, NULL);
INSERT INTO public.products VALUES (9, 'Bayern Munich Away Jersey 2025-26', 'Bayern Munich', 'Bundesliga', 2025, 1, 89.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471064/bayern-25-26-away_hkigr7.jpg', 'Away', 'The official Bayern Munich away jersey for the 2025-26 season. Bold design with metallic accents, made from ultra-light fabric for speed and agility. Features the Bavarian diamond pattern and official FC Bayern insignia.', NULL, '2025-12-11 14:57:27.881748+05:30', NULL, NULL);
INSERT INTO public.products VALUES (18, 'Portugal Home Jersey 2022', 'Portugal', 'International', 2022, 10, 85.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471066/por-22_zi4i8m.jpg', 'Home', 'The official Portugal home jersey for the 2022 season. Green and red with armillary sphere, made with lightweight, stretchable fabric. Celebrates Ronaldo and the Seleção''s passion.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (19, 'FC Barcelona Home Jersey 2014-15', 'FC Barcelona', 'Laliga', 2014, 10, 60.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471061/barcelona-2014-15-home_tgh3ij.jpg', 'Home', 'The official FC Barcelona home jersey for the 2014-15 season, treble winners. Horizontal stripes with Qatar sponsor, using ClimaCool for ventilation. A collector''s piece with authentic Barca heritage.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (20, 'Spain Home Jersey 2025', 'Spain', 'International', 2025, 10, 91.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471068/spain-25_ztl7uw.jpg', 'Home', 'The official Spain home jersey for the 2025 season. Red with gold accents, featuring Adidas Primegreen sustainable fabric. Includes the Spanish crest for Euro and World Cup pride.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (21, 'Man City Home Jersey 2021-22', 'Manchester City', 'Premier League', 2021, 10, 79.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471059/city-21-22-home_gaifsv.jpg', 'Home', 'The official Manchester City home jersey for the 2021-22 season. Sky blue with etched patterns, built with recycled polyester for performance. Etihad sponsor and eagle badge complete the Cityzens look.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (22, 'Arsenal Third Jersey 2025-26', 'Arsenal', 'Premier League', 2025, 10, 94.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471060/arsenal-2025-26-third_kit_tdhcgo.jpg', 'Third', 'The official Arsenal third jersey for the 2025-26 season. Exotic print inspired by London architecture, with AEROREADY moisture management. Cannon crest adds a unique twist to the Gunners'' collection.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (23, 'Real Madrid Home Jersey 2023-24', 'Real Madrid', 'Laliga', 2023, 10, 89.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471067/real-madrid-23-24-home_j1he0m.jpg', 'Home', 'The official Real Madrid home jersey for the 2023-24 season. Crisp white with golden details, featuring HEAT.RDY cooling zones. Celebrates La Decima and beyond with embroidered madridista pride.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (24, 'Argentina Away Jersey 2021-22', 'Argentina', 'International', 2021, 10, 80.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471058/argentina-21-22-away_ek5mhc.jpg', 'Away', 'The official Argentina away jersey for the 2021-22 season. Light purple with celestial motifs, made from breathable fabric for versatility. Sun of May emblem honors the Albiceleste''s journey.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (25, 'Juventus Home Jersey 2018-19', 'Juventus', 'Serie A', 2018, 10, 77.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471064/juventus-18-19-home_bjtuz4.jpg', 'Home', 'The official Juventus home jersey for the 2018-19 season. Black and white with J scudetto, using Techfit compression for support. Zebra stripes embody the Old Lady''s storied legacy.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (26, 'England Home Jersey 2024', 'England', 'International', 2024, 10, 90.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471060/england-24_tfubop.jpg', 'Home', 'The official England home jersey for the 2024 season. Red with subtle three lions, crafted from Dri-FIT ADV for advanced sweat-wicking. Inspired by 1966 triumph for the Three Lions supporters.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (27, 'AS Roma Home Jersey 2019-20', 'AS Roma', 'Serie A', 2019, 10, 75.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471067/roma-19-20_qlqdek.jpg', 'Home', 'The official AS Roma home jersey for the 2019-20 season. Maroon and gold with wolf emblem, featuring Nike''s recycled yarn for sustainability. Captures the Giallorossi''s passionate Roman spirit.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (28, 'FC Barcelona Home Jersey 2026', 'FC Barcelona', 'Laliga', 2026, 10, 95.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471061/barcelona-26-home_ur3t7w.webp', 'Home', 'The official FC Barcelona home jersey for the 2026 season. Futuristic Blaugrana with tech-infused fabric for peak performance. Official crest and innovative patterns for the next era of Barca dominance.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (29, 'AC Milan Home Jersey 2022-23', 'AC Milan', 'Serie A', 2022, 10, 85.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/milan-22-23-home_jfvlh2.jpg', 'Home', 'The official AC Milan home jersey for the 2022-23 season, Scudetto winners. Red-black with hexagonal patterns, using evoKNIT for flexibility. Rossoneri heritage shines through every stitch.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (30, 'Tottenham Home Jersey 2022-23', 'Tottenham Hotspur', 'Premier League', 2022, 10, 84.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471068/tott-22-2-home_fllkqz.jpg', 'Home', 'The official Tottenham Hotspur home jersey for the 2022-23 season. Navy blue with gold cockerel, featuring Nike Dri-FIT for all-weather comfort. Spurs'' fighting spirit in every detail.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (31, 'Juventus Away Jersey 2018-19', 'Juventus', 'Serie A', 2018, 10, 76.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471064/juventus-18-19-away_jibj30.jpg', 'Away', 'The official Juventus away jersey for the 2018-19 season. Vibrant yellow with black accents, built with breathable mesh panels. Juve crest ensures authenticity for away day warriors.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (32, 'Dortmund Home Jersey 2023-24', 'Borussia Dortmund', 'Bundesliga', 2023, 10, 88.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471061/dortmund-23-24-home_rvf5yh.jpg', 'Home', 'The official Borussia Dortmund home jersey for the 2023-24 season. Yellow wall-inspired black walls, using PUMA''s ULTRAWEAVE for lightness. BVB crest roars with Signal Iduna Park energy.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (33, 'Real Madrid Home Jersey 2025-26', 'Real Madrid', 'Laliga', 2025, 10, 95.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471069/real-madrid-2025-26-home_dmaxhn.jpg', 'Home', 'The official Real Madrid home jersey for the 2025-26 season. Timeless white with futuristic glow, featuring advanced cooling tech. Hala Madrid anthem embodied in premium craftsmanship.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (34, 'Inter Milan Home Jersey 2007-08', 'Inter Milan', 'Serie A', 2007, 10, 62.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471063/inter-07-08-home_mjj9zs.jpg', 'Home', 'The official Inter Milan home jersey for the 2007-08 season, treble era retro. Blue-black stripes with Pirelli sponsor, re-crafted with modern soft fabric for comfort and nostalgia.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (35, 'Mexico Home Jersey 2020-21', 'Mexico', 'International', 2020, 10, 78.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/mexico-20-21_oviw2o.jpg', 'Home', 'The official Mexico home jersey for the 2020-21 season. Green with eagle crest, using Adidas'' Climalite for humidity control. Viva Mexico vibrancy in every Aztec-inspired detail.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (36, 'RB Leipzig Home Jersey 2024-25', 'RB Leipzig', 'Bundesliga', 2024, 10, 89.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471066/rb-24-25_sezd4e.jpg', 'Home', 'The official RB Leipzig home jersey for the 2024-25 season. Red Bull energy in white base, with bio-based materials for sustainability. Bulls logo charges with RasenBallsport intensity.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (37, 'Real Madrid Home Jersey 2009-10', 'Real Madrid', 'Laliga', 2009, 10, 68.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471066/real-madrid-09-10-home_kvopdh.jpg', 'Home', 'The official Real Madrid home jersey for the 2009-10 season, Galacticos revival. White with blue accents, updated with breathable tech. A nod to Ronaldo''s arrival and Bernabeu glory.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (38, 'AC Milan Away Jersey 2017-18', 'AC Milan', 'Serie A', 2017, 10, 74.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/milan-17-18-away_er69pl.jpg', 'Away', 'The official AC Milan away jersey for the 2017-18 season. Grey with red details, featuring moisture-mapped zones. Diavolo away kit for San Siro sojourns with style.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (39, 'FC Barcelona Away Jersey 2016-17', 'FC Barcelona', 'Laliga', 2016, 10, 70.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471061/barcelona-2016-17-away_gmwlzv.jpg', 'Away', 'The official FC Barcelona away jersey for the 2016-17 season. Mustard yellow with blue trim, using Nike''s engineered knit for fit. Mes que un club extends to every away adventure.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (40, 'Spain Home Jersey 2012', 'Spain', 'International', 2012, 10, 68.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471067/spain-12_cbmxyp.jpg', 'Home', 'The official Spain home jersey for the 2012 season, Euro champions. Red with gold V-pattern, reimagined with modern ventilation. Tiki-taka legacy in adidas authenticity.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (41, 'Man United Home Jersey 2021-22', 'Manchester United', 'Premier League', 2021, 10, 81.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471068/united-21-22-home_qkaclp.jpg', 'Home', 'The official Manchester United home jersey for the 2021-22 season. Heritage red with black collar, using Heatmap tech for targeted cooling. Theatre of Dreams design for eternal Reds.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (42, 'Manchester City Third Jersey 2023-24', 'Manchester City', 'Premier League', 2023, 10, 88.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471059/city-23-24-third-kit_wo9av8.jpg', 'Third', 'The official Manchester City third jersey for the 2023-24 season. Pinstripe suit-inspired, with PUMA''s lightweight construction. Etihad elegance for continental clashes.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (43, 'Germany Home Jersey 2018', 'Germany', 'International', 2018, 10, 75.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471062/germany-18_kb5df7.jpg', 'Home', 'The official Germany home jersey for the 2018 season. White with black accents, featuring miDORI bio-yarn for eco-performance. Die Mannschaft resilience in every thread.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (44, 'Liverpool Home Jersey 2019-20', 'Liverpool', 'Premier League', 2019, 10, 78.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471064/liverpool-19-20-home_ifcayl.jpg', 'Home', 'The official Liverpool home jersey for the 2019-20 season, Champions League winners. Red with gold YNWA, using Nike''s VaporKnit for speed. This means more to the Kop.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (45, 'Germany Home Jersey 2022', 'Germany', 'International', 2022, 10, 85.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471063/grmany-22_lbddaz.jpg', 'Home', 'The official Germany home jersey for the 2022 season. Teutonic white with eagle update, crafted from Parley ocean plastic. Sustainable style for Deutschland''s future.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (46, 'Real Madrid Third Jersey 2020-21', 'Real Madrid', 'Laliga', 2020, 10, 80.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471067/real-madrid-20-21-third-kit_kcwxpt.jpg', 'Third', 'The official Real Madrid third jersey for the 2020-21 season. Purple haze design, with ClimaCool for all-conditions play. Alternative kit for Los Blancos'' versatility.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (47, 'Juventus Away Jersey 2025-26', 'Juventus', 'Serie A', 2025, 10, 93.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/juventus-25-26-away_nu5ksf.jpg', 'Away', 'The official Juventus away jersey for the 2025-26 season. Electric blue with gold, featuring REPREVE recycled fibers. Bianconeri boldness away from Turin.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (48, 'Liverpool Away Jersey 2023-24', 'Liverpool', 'Premier League', 2023, 10, 87.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/liverpool-23-24-away_mtwh6g.jpg', 'Away', 'The official Liverpool away jersey for the 2023-24 season. Purple with green accents, using Dri-FIT for endurance. Anfield on the road with Klopp-era flair.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (49, 'Arsenal Home Jersey 2018-19', 'Arsenal', 'Premier League', 2018, 10, 73.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471059/arsenal-18-19-home_fbn8dq.jpg', 'Home', 'The official Arsenal home jersey for the 2018-19 season. Brick red with navy, featuring Flyknit for seamless fit. Highbury homage in modern Gooner gear.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (50, 'Chelsea Away Jersey 2025-26', 'Chelsea', 'Premier League', 2025, 10, 92.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471059/chelsea-25-26-away_o3u75a.jpg', 'Away', 'The official Chelsea away jersey for the 2025-26 season. Teal with white stripes, built with Nike''s AeroSwift for aerodynamics. Stamford Bridge blues in away form.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (51, 'FC Barcelona Away Jersey 2024-25', 'FC Barcelona', 'Laliga', 2024, 10, 91.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471062/barcelona-2024-25-away_vfctmu.jpg', 'Away', 'The official FC Barcelona away jersey for the 2024-25 season. Dark indigo with light stripes, using AEROREADY for dry play. Catalan creativity on the road.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (52, 'Argentina Home Jersey 2025-26', 'Argentina', 'International', 2025, 10, 93.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471058/argentina-25-26-home_oj6b2x.jpg', 'Home', 'The official Argentina home jersey for the 2025-26 season. Classic light blue-white with updated sun, from sustainable yarns. Copa America champions'' enduring style.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (53, 'Man City Home Jersey 2025-26', 'Manchester City', 'Premier League', 2025, 10, 92.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471060/city-25-26-home_zjiwmw.jpg', 'Home', 'The official Manchester City home jersey for the 2025-26 season. Sky blue with electric details, featuring PUMA''s evoKNIT. Etihad skyline for City''s sky-high ambitions.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (54, 'Brazil Away Jersey 2016', 'Brazil', 'International', 2016, 10, 74.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471058/brazil-2016-away_dfo7fk.jpg', 'Away', 'The official Brazil away jersey for the 2016 season. Gold with green trim, using Nike''s Dri-FIT for samba rhythm. CBF crest for Seleção supporters worldwide.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (55, 'Brazil Home Jersey 2022', 'Brazil', 'International', 2022, 10, 87.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471057/brazil-2022_czo5cw.jpg', 'Home', 'The official Brazil home jersey for the 2022 season. Vibrant yellow with green, featuring hexagonal patterns for flair. Jogo Bonito in every pentagonal detail.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (56, 'Portugal Home Jersey 2025', 'Portugal', 'International', 2025, 10, 90.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471066/por-25_cq84nm.jpg', 'Home', 'The official Portugal home jersey for the 2025 season. Green with red accents, from Nike''s VaporKnit. Quinas shield for Ronaldo''s enduring legacy.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (57, 'Belgium Home Jersey 2018', 'Belgium', 'International', 2018, 10, 75.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471063/belgium-18_wcesgu.jpg', 'Home', 'The official Belgium home jersey for the 2018 season. Red devils red with black devil, using adidas tech for bronze medal pride. Diables Rouges intensity.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (58, 'Bayern Munich Away Jersey 2024-25', 'Bayern Munich', 'Bundesliga', 2024, 10, 90.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471063/bayern-24-25-away_pivdur.jpg', 'Away', 'The official Bayern Munich away jersey for the 2024-25 season. Deep purple with cream, featuring TELI insulation. Mia san mia echoes in Allianz Arena absences.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (59, 'Arsenal FC 2024-25 Away Kit', 'Arsenal FC', 'Premier League', 2023, 10, 45.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471060/arsenal-24-25-away_nssi1c.jpg', 'Home', 'The Adidas Arsenal 2024-25 away jersey has a black base, combined with red and green as a homage to former Arsenal players David Rocastle and Ian Wright. Enhanced with breathable mesh and official cannon badge for timeless Gunners tribute.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (60, 'Ajax 2025-26 Home Kit', 'Ajax', 'Bundesliga', 2026, 10, 56.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471058/ajax-25-26_dis420.jpg', 'Home', 'The Adidas AFC Ajax 2025-2026 home football shirt proudly features the club''s iconic traditional white base with a broad central red stripe. The sleeves are white, complemented by red Adidas stripes on the shoulders. Made with sustainable materials for De Godenzonen pride.', NULL, NULL, NULL, NULL);
INSERT INTO public.products VALUES (7, 'Liverpool Home Jersey 2025-26', 'Liverpool', 'Premier League', 2025, 7, 95.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/liverpool-25-26-home_ku2smo.jpg', 'Home', 'The official Liverpool home jersey for the 2025-26 season. Engineered with AeroReady technology to keep you dry, featuring the timeless red kit with enhanced collar for style and support. Authentic badges ensure a premium collector''s item.', NULL, '2025-11-24 23:02:34.806584+05:30', NULL, NULL);
INSERT INTO public.products VALUES (61, 'Arsenal 2025/26 Home Kit', 'Arsenal', 'Premier League', 2026, 10, 89.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1764087336/products/qkcwn5ijdpwegmgxvikh.webp', 'Home', 'The official home jersey for the 25/26 season.', '2025-11-25 21:45:36.229793+05:30', '2025-11-25 21:45:36.229793+05:30', '2025-11-25 23:48:23.169977+05:30', NULL);
INSERT INTO public.products VALUES (12, 'Argentina Home Jersey 2022', 'Argentina', 'International', 2022, 10, 90.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471058/argentina-2022-home_e77dew.jpg', 'Home', 'The official Argentina home jersey for the 2022 season, World Cup winners'' edition. Light blue and white stripes with sun emblem, made from breathable mesh for tropical climates. Celebrates Messi and the squad''s triumph.', NULL, '2025-11-27 14:38:17.236588+05:30', NULL, NULL);
INSERT INTO public.products VALUES (16, 'Manchester United Home Kit 25-26', 'Manchester United', 'Premier League', 2025, 9, 93.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471068/united-25-26-home_l7my2n.jpg', 'Home', 'The official Manchester United home jersey for the 2025-26 season. Iconic red with black accents, using Vaporknit technology for lightweight feel. Authentic devil crest and Adidas details for Old Trafford loyalty.', NULL, '2025-11-27 17:31:50.251918+05:30', NULL, NULL);
INSERT INTO public.products VALUES (63, 'dsfdsf', 'fds', 'Premier League', 13, 3, 983.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1765108611/products/mv6lncl4pnauzpyq3hy9.png', 'Home', 'fdgdfg', '2025-12-07 17:26:52.057862+05:30', '2025-12-07 17:26:52.057862+05:30', '2025-12-07 17:29:43.484475+05:30', NULL);
INSERT INTO public.products VALUES (64, 'man united 25-26 third kit', 'man united', 'Premier League', 2026, 0, 19.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1765187532/products/as0pkwvbq5u9d8xswl79.avif', 'Home', 'kansfd', '2025-12-08 15:22:13.723565+05:30', '2025-12-08 15:22:33.396023+05:30', '2025-12-08 15:22:57.018178+05:30', NULL);
INSERT INTO public.products VALUES (62, 'Arsenal 2025/26 Home Kit', 'Arsenal', 'Premier League', 2026, 9, 89.00, '$  ', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1764243673/products/b0u7yct2hmpbmtfxjtfp.webp', 'Home', 'The official home jersey for the 25/26 season.', '2025-11-27 17:11:14.844005+05:30', '2025-12-09 23:41:43.587067+05:30', NULL, NULL);


--
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.notifications VALUES (1, '2025-12-11 12:08:44.598185+05:30', '2025-12-11 12:34:05.069977+05:30', NULL, 22, 'Hey! Quick Update 😊', 'We’ve updated your account details. Everything looks good now! Let us know if you need help with anything else.', '2025-12-11 12:08:44.598185+05:30', true);
INSERT INTO public.notifications VALUES (2, '2025-12-11 12:49:26.83535+05:30', '2025-12-11 12:53:24.338382+05:30', NULL, 22, 'Important Notice', 'Action is required on your account. Please log in and review the latest update for your security.', '2025-12-11 12:49:26.83535+05:30', true);
INSERT INTO public.notifications VALUES (3, '2025-12-11 23:17:29.213152+05:30', '2025-12-11 23:18:02.487552+05:30', NULL, 22, 'Order Shipped! 🚚', 'Good news! Your Bayern Munich Away Jersey 2025-26 is on its way.', '2025-12-11 23:17:29.213152+05:30', true);
INSERT INTO public.notifications VALUES (4, '2025-12-11 23:32:24.209122+05:30', '2025-12-12 18:32:19.685431+05:30', NULL, 22, 'Order Confirmed! 🎉', 'Your order #ORD-20251211-4973 has been placed successfully. Wait for more updates', '2025-12-11 23:32:24.209122+05:30', true);


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.orders VALUES (17, '2025-12-10 09:56:20.144052+05:30', '2025-12-10 10:39:17.268088+05:30', NULL, 22, 14, 84, 'pending', 'COD', '2025-12-10 09:56:20.144052+05:30', 'ORD-20251210-6606');
INSERT INTO public.orders VALUES (16, '2025-12-10 09:55:48.516436+05:30', '2025-12-10 10:40:44.473553+05:30', NULL, 22, 14, 88, 'delivered', 'COD', '2025-12-10 09:55:48.516436+05:30', 'ORD-20251210-5068');
INSERT INTO public.orders VALUES (18, '2025-12-10 10:50:24.601841+05:30', '2025-12-10 10:50:24.601841+05:30', NULL, 3, 2, 181, 'pending', 'COD', '2025-12-10 10:50:24.601841+05:30', 'ORD-20251210-7729');
INSERT INTO public.orders VALUES (19, '2025-12-10 10:53:01.447533+05:30', '2025-12-10 11:56:07.263808+05:30', NULL, 3, 2, 0, 'cancelled', 'COD', '2025-12-10 10:53:01.447533+05:30', 'ORD-20251210-5273');
INSERT INTO public.orders VALUES (20, '2025-12-11 14:57:27.871479+05:30', '2025-12-11 14:57:27.871479+05:30', NULL, 22, 14, 89, 'pending', 'COD', '2025-12-11 14:57:27.871479+05:30', 'ORD-20251211-1368');
INSERT INTO public.orders VALUES (21, '2025-12-11 23:32:24.169796+05:30', '2025-12-11 23:32:24.169796+05:30', NULL, 22, 14, 84, 'pending', 'COD', '2025-12-11 23:32:24.169796+05:30', 'ORD-20251211-4973');


--
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_items VALUES (24, '2025-12-10 09:56:20.146228+05:30', '2025-12-10 09:56:58.223639+05:30', NULL, 17, 3, 84, 1, 'Germany Home Jersey 2014', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471062/germany-14_f9s5ak.jpg', 'delivered', '', NULL);
INSERT INTO public.order_items VALUES (25, '2025-12-10 09:56:20.147276+05:30', '2025-12-10 10:39:17.266653+05:30', NULL, 17, 4, 93, 1, 'Real Madrid Third Jersey 2016-17', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471066/real-madrid-16-17-third-kit_vuca9l.jpg', 'cancelled', 'Ordered by mistake', NULL);
INSERT INTO public.order_items VALUES (23, '2025-12-10 09:55:48.525398+05:30', '2025-12-10 10:40:44.4715+05:30', NULL, 16, 1, 88, 1, 'Tottenham Home Jersey 2025-26', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471068/Tott-25-26-home_cctozc.jpg', 'delivered', '', NULL);
INSERT INTO public.order_items VALUES (26, '2025-12-10 10:50:24.61156+05:30', '2025-12-10 10:50:56.207188+05:30', NULL, 18, 9, 89, 1, 'Bayern Munich Away Jersey 2025-26', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471064/bayern-25-26-away_hkigr7.jpg', 'shipped', '', NULL);
INSERT INTO public.order_items VALUES (27, '2025-12-10 10:50:24.614622+05:30', '2025-12-10 10:50:57.666212+05:30', NULL, 18, 15, 92, 1, 'AC Milan Home Jersey 2025-26', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/milan-25-26-home_fii9xo.jpg', 'delivered', '', NULL);
INSERT INTO public.order_items VALUES (28, '2025-12-10 10:53:01.448057+05:30', '2025-12-10 11:56:07.247115+05:30', NULL, 19, 15, 92, 1, 'AC Milan Home Jersey 2025-26', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471065/milan-25-26-home_fii9xo.jpg', 'cancelled', '', NULL);
INSERT INTO public.order_items VALUES (29, '2025-12-11 14:57:27.879978+05:30', '2025-12-11 23:17:29.207017+05:30', NULL, 20, 9, 89, 1, 'Bayern Munich Away Jersey 2025-26', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471064/bayern-25-26-away_hkigr7.jpg', 'shipped', '', NULL);
INSERT INTO public.order_items VALUES (30, '2025-12-11 23:32:24.192429+05:30', '2025-12-11 23:32:24.192429+05:30', NULL, 21, 3, 84, 1, 'Germany Home Jersey 2014', 'https://res.cloudinary.com/dhwsfp1hh/image/upload/v1763471062/germany-14_f9s5ak.jpg', 'active', '', NULL);


--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.payments VALUES (12, '2025-12-10 09:55:48.522317+05:30', '2025-12-10 09:55:48.522317+05:30', NULL, 0, 22, 88, 'COD', 'pending', '');
INSERT INTO public.payments VALUES (13, '2025-12-10 09:56:20.146228+05:30', '2025-12-10 09:56:20.146228+05:30', NULL, 0, 22, 177, 'COD', 'pending', '');
INSERT INTO public.payments VALUES (14, '2025-12-10 10:50:24.610041+05:30', '2025-12-10 10:50:24.610041+05:30', NULL, 0, 3, 181, 'COD', 'pending', '');
INSERT INTO public.payments VALUES (15, '2025-12-10 10:53:01.448057+05:30', '2025-12-10 10:53:01.448057+05:30', NULL, 0, 3, 92, 'COD', 'pending', '');
INSERT INTO public.payments VALUES (16, '2025-12-11 14:57:27.878321+05:30', '2025-12-11 14:57:27.878321+05:30', NULL, 0, 22, 89, 'COD', 'pending', '');
INSERT INTO public.payments VALUES (17, '2025-12-11 23:32:24.191371+05:30', '2025-12-11 23:32:24.191371+05:30', NULL, 0, 22, 84, 'COD', 'pending', '');


--
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.reviews VALUES (1, '2025-12-07 13:56:09.546938+05:30', '2025-12-07 13:56:09.546938+05:30', NULL, 3, 7, 'Nice product,good quality');
INSERT INTO public.reviews VALUES (2, '2025-12-07 14:06:45.439381+05:30', '2025-12-07 14:06:45.439381+05:30', NULL, 3, 7, 'Awesome product with speedy delivery very much impressed');
INSERT INTO public.reviews VALUES (3, '2025-12-08 15:12:45.018083+05:30', '2025-12-08 15:12:45.018083+05:30', NULL, 22, 42, 'This is a nice product');
INSERT INTO public.reviews VALUES (4, '2025-12-10 00:51:33.460112+05:30', '2025-12-10 00:51:33.460112+05:30', NULL, 3, 16, 'Nice product');
INSERT INTO public.reviews VALUES (5, '2025-12-10 00:53:04.268887+05:30', '2025-12-10 00:53:04.268887+05:30', NULL, 3, 16, 'Nice product');
INSERT INTO public.reviews VALUES (6, '2025-12-10 00:54:42.422512+05:30', '2025-12-10 00:54:42.422512+05:30', NULL, 3, 7, 'I love this jersey,it feels very comfortable to wear it
');


--
-- Data for Name: wishlists; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.wishlists VALUES (3, 33);
INSERT INTO public.wishlists VALUES (3, 6);
INSERT INTO public.wishlists VALUES (22, 9);


--
-- Name: addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.addresses_id_seq', 16, true);


--
-- Name: notifications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.notifications_id_seq', 4, true);


--
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_items_id_seq', 30, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 21, true);


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payments_id_seq', 17, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 64, true);


--
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reviews_id_seq', 6, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 31, true);


--
-- PostgreSQL database dump complete
--

\unrestrict dyOdNTWV1hXI4SzxTwNfXekKjic6GdvxJ9NdAvfYRicizuZPaTxRpF9czmiDUGL

