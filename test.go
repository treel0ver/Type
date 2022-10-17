package main

import (
	"time"
    "strings"
    "strconv"
    "bufio"
    "os"
    "fmt"

	//"strconv"
	//"reflect"
)

var almacenar_tops_en_archivo [3000]string

func escribir_tops_a_base_de_datos() {
almacenar_tops_en_archivo[	0	] = 		tops[	0	]	[	0	]
almacenar_tops_en_archivo[	1	] = 		tops[	0	]	[	1	]
almacenar_tops_en_archivo[	2	] = 		tops[	0	]	[	2	]
almacenar_tops_en_archivo[	3	] = 		tops[	0	]	[	3	]
almacenar_tops_en_archivo[	4	] = 		tops[	0	]	[	4	]
almacenar_tops_en_archivo[	5	] = 		tops[	0	]	[	5	]
almacenar_tops_en_archivo[	6	] = 		tops[	1	]	[	0	]
almacenar_tops_en_archivo[	7	] = 		tops[	1	]	[	1	]
almacenar_tops_en_archivo[	8	] = 		tops[	1	]	[	2	]
almacenar_tops_en_archivo[	9	] = 		tops[	1	]	[	3	]
almacenar_tops_en_archivo[	10	] = 		tops[	1	]	[	4	]
almacenar_tops_en_archivo[	11	] = 		tops[	1	]	[	5	]
almacenar_tops_en_archivo[	12	] = 		tops[	2	]	[	0	]
almacenar_tops_en_archivo[	13	] = 		tops[	2	]	[	1	]
almacenar_tops_en_archivo[	14	] = 		tops[	2	]	[	2	]
almacenar_tops_en_archivo[	15	] = 		tops[	2	]	[	3	]
almacenar_tops_en_archivo[	16	] = 		tops[	2	]	[	4	]
almacenar_tops_en_archivo[	17	] = 		tops[	2	]	[	5	]
almacenar_tops_en_archivo[	18	] = 		tops[	3	]	[	0	]
almacenar_tops_en_archivo[	19	] = 		tops[	3	]	[	1	]
almacenar_tops_en_archivo[	20	] = 		tops[	3	]	[	2	]
almacenar_tops_en_archivo[	21	] = 		tops[	3	]	[	3	]
almacenar_tops_en_archivo[	22	] = 		tops[	3	]	[	4	]
almacenar_tops_en_archivo[	23	] = 		tops[	3	]	[	5	]
almacenar_tops_en_archivo[	24	] = 		tops[	4	]	[	0	]
almacenar_tops_en_archivo[	25	] = 		tops[	4	]	[	1	]
almacenar_tops_en_archivo[	26	] = 		tops[	4	]	[	2	]
almacenar_tops_en_archivo[	27	] = 		tops[	4	]	[	3	]
almacenar_tops_en_archivo[	28	] = 		tops[	4	]	[	4	]
almacenar_tops_en_archivo[	29	] = 		tops[	4	]	[	5	]
almacenar_tops_en_archivo[	30	] = 		tops[	5	]	[	0	]
almacenar_tops_en_archivo[	31	] = 		tops[	5	]	[	1	]
almacenar_tops_en_archivo[	32	] = 		tops[	5	]	[	2	]
almacenar_tops_en_archivo[	33	] = 		tops[	5	]	[	3	]
almacenar_tops_en_archivo[	34	] = 		tops[	5	]	[	4	]
almacenar_tops_en_archivo[	35	] = 		tops[	5	]	[	5	]
almacenar_tops_en_archivo[	36	] = 		tops[	6	]	[	0	]
almacenar_tops_en_archivo[	37	] = 		tops[	6	]	[	1	]
almacenar_tops_en_archivo[	38	] = 		tops[	6	]	[	2	]
almacenar_tops_en_archivo[	39	] = 		tops[	6	]	[	3	]
almacenar_tops_en_archivo[	40	] = 		tops[	6	]	[	4	]
almacenar_tops_en_archivo[	41	] = 		tops[	6	]	[	5	]
almacenar_tops_en_archivo[	42	] = 		tops[	7	]	[	0	]
almacenar_tops_en_archivo[	43	] = 		tops[	7	]	[	1	]
almacenar_tops_en_archivo[	44	] = 		tops[	7	]	[	2	]
almacenar_tops_en_archivo[	45	] = 		tops[	7	]	[	3	]
almacenar_tops_en_archivo[	46	] = 		tops[	7	]	[	4	]
almacenar_tops_en_archivo[	47	] = 		tops[	7	]	[	5	]
almacenar_tops_en_archivo[	48	] = 		tops[	8	]	[	0	]
almacenar_tops_en_archivo[	49	] = 		tops[	8	]	[	1	]
almacenar_tops_en_archivo[	50	] = 		tops[	8	]	[	2	]
almacenar_tops_en_archivo[	51	] = 		tops[	8	]	[	3	]
almacenar_tops_en_archivo[	52	] = 		tops[	8	]	[	4	]
almacenar_tops_en_archivo[	53	] = 		tops[	8	]	[	5	]
almacenar_tops_en_archivo[	54	] = 		tops[	9	]	[	0	]
almacenar_tops_en_archivo[	55	] = 		tops[	9	]	[	1	]
almacenar_tops_en_archivo[	56	] = 		tops[	9	]	[	2	]
almacenar_tops_en_archivo[	57	] = 		tops[	9	]	[	3	]
almacenar_tops_en_archivo[	58	] = 		tops[	9	]	[	4	]
almacenar_tops_en_archivo[	59	] = 		tops[	9	]	[	5	]
almacenar_tops_en_archivo[	60	] = 		tops[	10	]	[	0	]
almacenar_tops_en_archivo[	61	] = 		tops[	10	]	[	1	]
almacenar_tops_en_archivo[	62	] = 		tops[	10	]	[	2	]
almacenar_tops_en_archivo[	63	] = 		tops[	10	]	[	3	]
almacenar_tops_en_archivo[	64	] = 		tops[	10	]	[	4	]
almacenar_tops_en_archivo[	65	] = 		tops[	10	]	[	5	]
almacenar_tops_en_archivo[	66	] = 		tops[	11	]	[	0	]
almacenar_tops_en_archivo[	67	] = 		tops[	11	]	[	1	]
almacenar_tops_en_archivo[	68	] = 		tops[	11	]	[	2	]
almacenar_tops_en_archivo[	69	] = 		tops[	11	]	[	3	]
almacenar_tops_en_archivo[	70	] = 		tops[	11	]	[	4	]
almacenar_tops_en_archivo[	71	] = 		tops[	11	]	[	5	]
almacenar_tops_en_archivo[	72	] = 		tops[	12	]	[	0	]
almacenar_tops_en_archivo[	73	] = 		tops[	12	]	[	1	]
almacenar_tops_en_archivo[	74	] = 		tops[	12	]	[	2	]
almacenar_tops_en_archivo[	75	] = 		tops[	12	]	[	3	]
almacenar_tops_en_archivo[	76	] = 		tops[	12	]	[	4	]
almacenar_tops_en_archivo[	77	] = 		tops[	12	]	[	5	]
almacenar_tops_en_archivo[	78	] = 		tops[	13	]	[	0	]
almacenar_tops_en_archivo[	79	] = 		tops[	13	]	[	1	]
almacenar_tops_en_archivo[	80	] = 		tops[	13	]	[	2	]
almacenar_tops_en_archivo[	81	] = 		tops[	13	]	[	3	]
almacenar_tops_en_archivo[	82	] = 		tops[	13	]	[	4	]
almacenar_tops_en_archivo[	83	] = 		tops[	13	]	[	5	]
almacenar_tops_en_archivo[	84	] = 		tops[	14	]	[	0	]
almacenar_tops_en_archivo[	85	] = 		tops[	14	]	[	1	]
almacenar_tops_en_archivo[	86	] = 		tops[	14	]	[	2	]
almacenar_tops_en_archivo[	87	] = 		tops[	14	]	[	3	]
almacenar_tops_en_archivo[	88	] = 		tops[	14	]	[	4	]
almacenar_tops_en_archivo[	89	] = 		tops[	14	]	[	5	]
almacenar_tops_en_archivo[	90	] = 		tops[	15	]	[	0	]
almacenar_tops_en_archivo[	91	] = 		tops[	15	]	[	1	]
almacenar_tops_en_archivo[	92	] = 		tops[	15	]	[	2	]
almacenar_tops_en_archivo[	93	] = 		tops[	15	]	[	3	]
almacenar_tops_en_archivo[	94	] = 		tops[	15	]	[	4	]
almacenar_tops_en_archivo[	95	] = 		tops[	15	]	[	5	]
almacenar_tops_en_archivo[	96	] = 		tops[	16	]	[	0	]
almacenar_tops_en_archivo[	97	] = 		tops[	16	]	[	1	]
almacenar_tops_en_archivo[	98	] = 		tops[	16	]	[	2	]
almacenar_tops_en_archivo[	99	] = 		tops[	16	]	[	3	]
almacenar_tops_en_archivo[	100	] = 		tops[	16	]	[	4	]
almacenar_tops_en_archivo[	101	] = 		tops[	16	]	[	5	]
almacenar_tops_en_archivo[	102	] = 		tops[	17	]	[	0	]
almacenar_tops_en_archivo[	103	] = 		tops[	17	]	[	1	]
almacenar_tops_en_archivo[	104	] = 		tops[	17	]	[	2	]
almacenar_tops_en_archivo[	105	] = 		tops[	17	]	[	3	]
almacenar_tops_en_archivo[	106	] = 		tops[	17	]	[	4	]
almacenar_tops_en_archivo[	107	] = 		tops[	17	]	[	5	]
almacenar_tops_en_archivo[	108	] = 		tops[	18	]	[	0	]
almacenar_tops_en_archivo[	109	] = 		tops[	18	]	[	1	]
almacenar_tops_en_archivo[	110	] = 		tops[	18	]	[	2	]
almacenar_tops_en_archivo[	111	] = 		tops[	18	]	[	3	]
almacenar_tops_en_archivo[	112	] = 		tops[	18	]	[	4	]
almacenar_tops_en_archivo[	113	] = 		tops[	18	]	[	5	]
almacenar_tops_en_archivo[	114	] = 		tops[	19	]	[	0	]
almacenar_tops_en_archivo[	115	] = 		tops[	19	]	[	1	]
almacenar_tops_en_archivo[	116	] = 		tops[	19	]	[	2	]
almacenar_tops_en_archivo[	117	] = 		tops[	19	]	[	3	]
almacenar_tops_en_archivo[	118	] = 		tops[	19	]	[	4	]
almacenar_tops_en_archivo[	119	] = 		tops[	19	]	[	5	]
almacenar_tops_en_archivo[	120	] = 		tops[	20	]	[	0	]
almacenar_tops_en_archivo[	121	] = 		tops[	20	]	[	1	]
almacenar_tops_en_archivo[	122	] = 		tops[	20	]	[	2	]
almacenar_tops_en_archivo[	123	] = 		tops[	20	]	[	3	]
almacenar_tops_en_archivo[	124	] = 		tops[	20	]	[	4	]
almacenar_tops_en_archivo[	125	] = 		tops[	20	]	[	5	]
almacenar_tops_en_archivo[	126	] = 		tops[	21	]	[	0	]
almacenar_tops_en_archivo[	127	] = 		tops[	21	]	[	1	]
almacenar_tops_en_archivo[	128	] = 		tops[	21	]	[	2	]
almacenar_tops_en_archivo[	129	] = 		tops[	21	]	[	3	]
almacenar_tops_en_archivo[	130	] = 		tops[	21	]	[	4	]
almacenar_tops_en_archivo[	131	] = 		tops[	21	]	[	5	]
almacenar_tops_en_archivo[	132	] = 		tops[	22	]	[	0	]
almacenar_tops_en_archivo[	133	] = 		tops[	22	]	[	1	]
almacenar_tops_en_archivo[	134	] = 		tops[	22	]	[	2	]
almacenar_tops_en_archivo[	135	] = 		tops[	22	]	[	3	]
almacenar_tops_en_archivo[	136	] = 		tops[	22	]	[	4	]
almacenar_tops_en_archivo[	137	] = 		tops[	22	]	[	5	]
almacenar_tops_en_archivo[	138	] = 		tops[	23	]	[	0	]
almacenar_tops_en_archivo[	139	] = 		tops[	23	]	[	1	]
almacenar_tops_en_archivo[	140	] = 		tops[	23	]	[	2	]
almacenar_tops_en_archivo[	141	] = 		tops[	23	]	[	3	]
almacenar_tops_en_archivo[	142	] = 		tops[	23	]	[	4	]
almacenar_tops_en_archivo[	143	] = 		tops[	23	]	[	5	]
almacenar_tops_en_archivo[	144	] = 		tops[	24	]	[	0	]
almacenar_tops_en_archivo[	145	] = 		tops[	24	]	[	1	]
almacenar_tops_en_archivo[	146	] = 		tops[	24	]	[	2	]
almacenar_tops_en_archivo[	147	] = 		tops[	24	]	[	3	]
almacenar_tops_en_archivo[	148	] = 		tops[	24	]	[	4	]
almacenar_tops_en_archivo[	149	] = 		tops[	24	]	[	5	]
almacenar_tops_en_archivo[	150	] = 		tops[	25	]	[	0	]
almacenar_tops_en_archivo[	151	] = 		tops[	25	]	[	1	]
almacenar_tops_en_archivo[	152	] = 		tops[	25	]	[	2	]
almacenar_tops_en_archivo[	153	] = 		tops[	25	]	[	3	]
almacenar_tops_en_archivo[	154	] = 		tops[	25	]	[	4	]
almacenar_tops_en_archivo[	155	] = 		tops[	25	]	[	5	]
almacenar_tops_en_archivo[	156	] = 		tops[	26	]	[	0	]
almacenar_tops_en_archivo[	157	] = 		tops[	26	]	[	1	]
almacenar_tops_en_archivo[	158	] = 		tops[	26	]	[	2	]
almacenar_tops_en_archivo[	159	] = 		tops[	26	]	[	3	]
almacenar_tops_en_archivo[	160	] = 		tops[	26	]	[	4	]
almacenar_tops_en_archivo[	161	] = 		tops[	26	]	[	5	]
almacenar_tops_en_archivo[	162	] = 		tops[	27	]	[	0	]
almacenar_tops_en_archivo[	163	] = 		tops[	27	]	[	1	]
almacenar_tops_en_archivo[	164	] = 		tops[	27	]	[	2	]
almacenar_tops_en_archivo[	165	] = 		tops[	27	]	[	3	]
almacenar_tops_en_archivo[	166	] = 		tops[	27	]	[	4	]
almacenar_tops_en_archivo[	167	] = 		tops[	27	]	[	5	]
almacenar_tops_en_archivo[	168	] = 		tops[	28	]	[	0	]
almacenar_tops_en_archivo[	169	] = 		tops[	28	]	[	1	]
almacenar_tops_en_archivo[	170	] = 		tops[	28	]	[	2	]
almacenar_tops_en_archivo[	171	] = 		tops[	28	]	[	3	]
almacenar_tops_en_archivo[	172	] = 		tops[	28	]	[	4	]
almacenar_tops_en_archivo[	173	] = 		tops[	28	]	[	5	]
almacenar_tops_en_archivo[	174	] = 		tops[	29	]	[	0	]
almacenar_tops_en_archivo[	175	] = 		tops[	29	]	[	1	]
almacenar_tops_en_archivo[	176	] = 		tops[	29	]	[	2	]
almacenar_tops_en_archivo[	177	] = 		tops[	29	]	[	3	]
almacenar_tops_en_archivo[	178	] = 		tops[	29	]	[	4	]
almacenar_tops_en_archivo[	179	] = 		tops[	29	]	[	5	]
almacenar_tops_en_archivo[	180	] = 		tops[	30	]	[	0	]
almacenar_tops_en_archivo[	181	] = 		tops[	30	]	[	1	]
almacenar_tops_en_archivo[	182	] = 		tops[	30	]	[	2	]
almacenar_tops_en_archivo[	183	] = 		tops[	30	]	[	3	]
almacenar_tops_en_archivo[	184	] = 		tops[	30	]	[	4	]
almacenar_tops_en_archivo[	185	] = 		tops[	30	]	[	5	]
almacenar_tops_en_archivo[	186	] = 		tops[	31	]	[	0	]
almacenar_tops_en_archivo[	187	] = 		tops[	31	]	[	1	]
almacenar_tops_en_archivo[	188	] = 		tops[	31	]	[	2	]
almacenar_tops_en_archivo[	189	] = 		tops[	31	]	[	3	]
almacenar_tops_en_archivo[	190	] = 		tops[	31	]	[	4	]
almacenar_tops_en_archivo[	191	] = 		tops[	31	]	[	5	]
almacenar_tops_en_archivo[	192	] = 		tops[	32	]	[	0	]
almacenar_tops_en_archivo[	193	] = 		tops[	32	]	[	1	]
almacenar_tops_en_archivo[	194	] = 		tops[	32	]	[	2	]
almacenar_tops_en_archivo[	195	] = 		tops[	32	]	[	3	]
almacenar_tops_en_archivo[	196	] = 		tops[	32	]	[	4	]
almacenar_tops_en_archivo[	197	] = 		tops[	32	]	[	5	]
almacenar_tops_en_archivo[	198	] = 		tops[	33	]	[	0	]
almacenar_tops_en_archivo[	199	] = 		tops[	33	]	[	1	]
almacenar_tops_en_archivo[	200	] = 		tops[	33	]	[	2	]
almacenar_tops_en_archivo[	201	] = 		tops[	33	]	[	3	]
almacenar_tops_en_archivo[	202	] = 		tops[	33	]	[	4	]
almacenar_tops_en_archivo[	203	] = 		tops[	33	]	[	5	]
almacenar_tops_en_archivo[	204	] = 		tops[	34	]	[	0	]
almacenar_tops_en_archivo[	205	] = 		tops[	34	]	[	1	]
almacenar_tops_en_archivo[	206	] = 		tops[	34	]	[	2	]
almacenar_tops_en_archivo[	207	] = 		tops[	34	]	[	3	]
almacenar_tops_en_archivo[	208	] = 		tops[	34	]	[	4	]
almacenar_tops_en_archivo[	209	] = 		tops[	34	]	[	5	]
almacenar_tops_en_archivo[	210	] = 		tops[	35	]	[	0	]
almacenar_tops_en_archivo[	211	] = 		tops[	35	]	[	1	]
almacenar_tops_en_archivo[	212	] = 		tops[	35	]	[	2	]
almacenar_tops_en_archivo[	213	] = 		tops[	35	]	[	3	]
almacenar_tops_en_archivo[	214	] = 		tops[	35	]	[	4	]
almacenar_tops_en_archivo[	215	] = 		tops[	35	]	[	5	]
almacenar_tops_en_archivo[	216	] = 		tops[	36	]	[	0	]
almacenar_tops_en_archivo[	217	] = 		tops[	36	]	[	1	]
almacenar_tops_en_archivo[	218	] = 		tops[	36	]	[	2	]
almacenar_tops_en_archivo[	219	] = 		tops[	36	]	[	3	]
almacenar_tops_en_archivo[	220	] = 		tops[	36	]	[	4	]
almacenar_tops_en_archivo[	221	] = 		tops[	36	]	[	5	]
almacenar_tops_en_archivo[	222	] = 		tops[	37	]	[	0	]
almacenar_tops_en_archivo[	223	] = 		tops[	37	]	[	1	]
almacenar_tops_en_archivo[	224	] = 		tops[	37	]	[	2	]
almacenar_tops_en_archivo[	225	] = 		tops[	37	]	[	3	]
almacenar_tops_en_archivo[	226	] = 		tops[	37	]	[	4	]
almacenar_tops_en_archivo[	227	] = 		tops[	37	]	[	5	]
almacenar_tops_en_archivo[	228	] = 		tops[	38	]	[	0	]
almacenar_tops_en_archivo[	229	] = 		tops[	38	]	[	1	]
almacenar_tops_en_archivo[	230	] = 		tops[	38	]	[	2	]
almacenar_tops_en_archivo[	231	] = 		tops[	38	]	[	3	]
almacenar_tops_en_archivo[	232	] = 		tops[	38	]	[	4	]
almacenar_tops_en_archivo[	233	] = 		tops[	38	]	[	5	]
almacenar_tops_en_archivo[	234	] = 		tops[	39	]	[	0	]
almacenar_tops_en_archivo[	235	] = 		tops[	39	]	[	1	]
almacenar_tops_en_archivo[	236	] = 		tops[	39	]	[	2	]
almacenar_tops_en_archivo[	237	] = 		tops[	39	]	[	3	]
almacenar_tops_en_archivo[	238	] = 		tops[	39	]	[	4	]
almacenar_tops_en_archivo[	239	] = 		tops[	39	]	[	5	]
almacenar_tops_en_archivo[	240	] = 		tops[	40	]	[	0	]
almacenar_tops_en_archivo[	241	] = 		tops[	40	]	[	1	]
almacenar_tops_en_archivo[	242	] = 		tops[	40	]	[	2	]
almacenar_tops_en_archivo[	243	] = 		tops[	40	]	[	3	]
almacenar_tops_en_archivo[	244	] = 		tops[	40	]	[	4	]
almacenar_tops_en_archivo[	245	] = 		tops[	40	]	[	5	]
almacenar_tops_en_archivo[	246	] = 		tops[	41	]	[	0	]
almacenar_tops_en_archivo[	247	] = 		tops[	41	]	[	1	]
almacenar_tops_en_archivo[	248	] = 		tops[	41	]	[	2	]
almacenar_tops_en_archivo[	249	] = 		tops[	41	]	[	3	]
almacenar_tops_en_archivo[	250	] = 		tops[	41	]	[	4	]
almacenar_tops_en_archivo[	251	] = 		tops[	41	]	[	5	]
almacenar_tops_en_archivo[	252	] = 		tops[	42	]	[	0	]
almacenar_tops_en_archivo[	253	] = 		tops[	42	]	[	1	]
almacenar_tops_en_archivo[	254	] = 		tops[	42	]	[	2	]
almacenar_tops_en_archivo[	255	] = 		tops[	42	]	[	3	]
almacenar_tops_en_archivo[	256	] = 		tops[	42	]	[	4	]
almacenar_tops_en_archivo[	257	] = 		tops[	42	]	[	5	]
almacenar_tops_en_archivo[	258	] = 		tops[	43	]	[	0	]
almacenar_tops_en_archivo[	259	] = 		tops[	43	]	[	1	]
almacenar_tops_en_archivo[	260	] = 		tops[	43	]	[	2	]
almacenar_tops_en_archivo[	261	] = 		tops[	43	]	[	3	]
almacenar_tops_en_archivo[	262	] = 		tops[	43	]	[	4	]
almacenar_tops_en_archivo[	263	] = 		tops[	43	]	[	5	]
almacenar_tops_en_archivo[	264	] = 		tops[	44	]	[	0	]
almacenar_tops_en_archivo[	265	] = 		tops[	44	]	[	1	]
almacenar_tops_en_archivo[	266	] = 		tops[	44	]	[	2	]
almacenar_tops_en_archivo[	267	] = 		tops[	44	]	[	3	]
almacenar_tops_en_archivo[	268	] = 		tops[	44	]	[	4	]
almacenar_tops_en_archivo[	269	] = 		tops[	44	]	[	5	]
almacenar_tops_en_archivo[	270	] = 		tops[	45	]	[	0	]
almacenar_tops_en_archivo[	271	] = 		tops[	45	]	[	1	]
almacenar_tops_en_archivo[	272	] = 		tops[	45	]	[	2	]
almacenar_tops_en_archivo[	273	] = 		tops[	45	]	[	3	]
almacenar_tops_en_archivo[	274	] = 		tops[	45	]	[	4	]
almacenar_tops_en_archivo[	275	] = 		tops[	45	]	[	5	]
almacenar_tops_en_archivo[	276	] = 		tops[	46	]	[	0	]
almacenar_tops_en_archivo[	277	] = 		tops[	46	]	[	1	]
almacenar_tops_en_archivo[	278	] = 		tops[	46	]	[	2	]
almacenar_tops_en_archivo[	279	] = 		tops[	46	]	[	3	]
almacenar_tops_en_archivo[	280	] = 		tops[	46	]	[	4	]
almacenar_tops_en_archivo[	281	] = 		tops[	46	]	[	5	]
almacenar_tops_en_archivo[	282	] = 		tops[	47	]	[	0	]

}

func writeLines() error {
    file, err := os.Create("database/TOPS")
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range almacenar_tops_en_archivo {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

var textos [1000]string = [1000]string{
	"Escucha, las reglas propias... se tratan de decidir conseguir algo usando medios y maneras propias para conseguirlo. Por eso decimos que son nuestras reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si fracasamos, hay que retomar la práctica y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello, creas tus propias reglas.",
	"Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.",
	"Tal vez solo nos espere un camino oscuro por delante. Pero aún así necesitas creer y seguir adelante. Cree en que las estrellas iluminarán tu camino, incluso un poco. Vamos, ¡emprendamos una aventura!",
	"Sólo soy verdaderamente libre cuando todos los seres humanos que me rodean, hombres y mujeres, son igualmente libres, de manera que cuanto más numerosos son los hombres libres que me rodean y más profunda y más amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.",
	"Cada tipo que crece podría ser un repetidor, un pequeño maestro, el desencadenante de una reacción en cadena que en sí misma es capaz de cambiar el mundo.",
}
var display_textos [1000]string = [1000]string{
	"**Escucha,​ las reglas​ propias... se tratan de decidir conseguir algo​ usando medios y maneras propias​ para conseguirlo. Por eso decimos que son nuestras​ reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si​ fracasamos, hay que retomar la práctica​ y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello​, creas tus propias reglas.​**",
	"**Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.**",
	"**Tal​ vez solo nos espere un camino oscuro por​ delante. Pero aún así necesitas​ creer y seguir adelante. Cree en que las estrellas​ iluminarán tu camino, incluso un poco. Vamos,​ ¡emprendamos una aventura!**",
	"**Sólo soy​ verdaderamente libre cuando todos los seres humanos​ que me rodean, hombres y mujeres,​ son igualmente libres, de manera que cuanto​ más numerosos son los hombres libres​ que me rodean y más profunda y más​ amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.​**",
	"**Cada​ tipo que crece podría​ ser un repetidor, un pequeño maestro,​ el desencadenante de una reacción en cadena que en sí misma es capaz​ de cambiar el mundo.**",
}

var tops [5000][6]string = [5000][6]string{
	{"1", "0", "0", "0", "0", "0"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},

	{"1", "0", "0", "0", "0", "0"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},

	{"1", "0", "0", "0", "0", "0"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},

	{"1", "0", "0", "0", "0", "0"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},

	{"1", "0", "0", "0", "0", "0"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},
}

const INVISIBLE string = "​"

func how_many_texts() int {
	var cuenta int
	for i := 1; i <= len(textos); i++ {
	    if textos[i] != "" {
	    	cuenta++
	    } else { return cuenta+1 }
	}
	return 0
}

func split_curr() string {
	arrayed := strings.Split(current_text, " ")
	return arrayed[0] + " " + arrayed[1] + " " + arrayed[2] + " " + arrayed[3] + " " + arrayed[4] + " " + arrayed[5] + "..."
}

func is_illegal(s string) bool{
	x := strings.Contains(s, INVISIBLE)
	return x
}

func split(tosplit string, sep rune) []string {
	var fields []string

	last := 0
	for i, c := range tosplit {
		if c == sep {
			fields = append(fields, string(tosplit[last:i]))
			last = i + 1
		}
	}
	fields = append(fields, string(tosplit[last:]))

	return fields
}

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func average_word_length(s string) float64 {
	arrayed := strings.Split(s, " ")
	var total = 0
	for i := 0; i < len(arrayed); i++ {
		total += len([]rune(arrayed[i]))
	}
	var x float64 = float64(total)/float64(len(arrayed))
	return x
}

var random int
var current_text = textos[0]
var is_started bool = false
var time_elapsed float64 = 0

var wpm float64 = 0

var start_author string

var result_splited[]string

func start() {
	time_elapsed = 0
	if is_started == false {
		is_started = true
		for is_started {
			time.Sleep(1 * time.Millisecond)
			time_elapsed = time_elapsed + 1
		}
	}

}

func calculate_wpm() {
	var length = len([]rune(current_text))/* - (CountWords(current_text)) */
	var length_as_a_float float64 = float64(length)
	wpm = length_as_a_float / 5 / time_elapsed * 60000
}

var errores int = 0
var errores_s string

var lista_errores string

func calculate_errors(sent string, current string) {
	/* reseting */
	errores = 0
	lista_errores = ""

	A := sent
	sent_arrayed := strings.Split(A, " ")

	B := current
	text_arrayed := strings.Split(B, " ")

	if len(text_arrayed) == len(sent_arrayed) {
		for i := 0; i < len(text_arrayed); i++ {
				if text_arrayed[i] != sent_arrayed[i] {
					if lista_errores != "" {
						lista_errores = lista_errores + ", " + sent_arrayed[i]
						errores++
					} else {lista_errores = sent_arrayed[i]}
				}
			}
	}

	if  len(text_arrayed) > len(sent_arrayed) {

	}


	errores_s = strconv.FormatInt(int64(errores), 10)
}
