package main

import (
	"flag"
 	"fmt"
 	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
    "log"
    "time"
    "math/rand"
    "strings"
	"strconv"
	"reflect"
	"github.com/bwmarrin/discordgo"
)



																	func check(e error) {
																	    if e != nil {
																	        panic(e)
																	    }
																	}

/* Variables used for command line parameters */
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	 var fileName = "database/TOPS"
 	fileBytes, err := ioutil.ReadFile(fileName)

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	sliceData := strings.Split(string(fileBytes), "\n")

 	fmt.Println(sliceData[21])
 	fmt.Println(how_many_texts())

	tops[0][0] = sliceData[0]	; tops[0][1] = sliceData[1]	; tops[0][2] = sliceData[2]	; tops[0][3] = sliceData[3]	; tops[0][4] = sliceData[4]	; tops[0][5] = sliceData[5]	; tops[1][0] = sliceData[6]	; tops[1][1] = sliceData[7]	; tops[1][2] = sliceData[8]	; tops[1][3] = sliceData[9]	; tops[1][4] = sliceData[10]	; tops[1][5] = sliceData[11]	; tops[2][0] = sliceData[12]	; tops[2][1] = sliceData[13]	; tops[2][2] = sliceData[14]	; tops[2][3] = sliceData[15]	; tops[2][4] = sliceData[16]	; tops[2][5] = sliceData[17]	; tops[3][0] = sliceData[18]	; tops[3][1] = sliceData[19]	; tops[3][2] = sliceData[20]	; tops[3][3] = sliceData[21]	; tops[3][4] = sliceData[22]	; tops[3][5] = sliceData[23]	; tops[4][0] = sliceData[24]	; tops[4][1] = sliceData[25]	; tops[4][2] = sliceData[26]	; tops[4][3] = sliceData[27]	; tops[4][4] = sliceData[28]	; tops[4][5] = sliceData[29]	; tops[5][0] = sliceData[30]	; tops[5][1] = sliceData[31]	; tops[5][2] = sliceData[32]	; tops[5][3] = sliceData[33]	; tops[5][4] = sliceData[34]	; tops[5][5] = sliceData[35]	; tops[6][0] = sliceData[36]	; tops[6][1] = sliceData[37]	; tops[6][2] = sliceData[38]	; tops[6][3] = sliceData[39]	; tops[6][4] = sliceData[40]	; tops[6][5] = sliceData[41]	; tops[7][0] = sliceData[42]	; tops[7][1] = sliceData[43]	; tops[7][2] = sliceData[44]	; tops[7][3] = sliceData[45]	; tops[7][4] = sliceData[46]	; tops[7][5] = sliceData[47]	; tops[8][0] = sliceData[48]	; tops[8][1] = sliceData[49]	; tops[8][2] = sliceData[50]	; tops[8][3] = sliceData[51]	; tops[8][4] = sliceData[52]	; tops[8][5] = sliceData[53]	; tops[9][0] = sliceData[54]	; tops[9][1] = sliceData[55]	; tops[9][2] = sliceData[56]	; tops[9][3] = sliceData[57]	; tops[9][4] = sliceData[58]	; tops[9][5] = sliceData[59]	; tops[10][0] = sliceData[60]	; tops[10][1] = sliceData[61]	; tops[10][2] = sliceData[62]	; tops[10][3] = sliceData[63]	; tops[10][4] = sliceData[64]	; tops[10][5] = sliceData[65]	; tops[11][0] = sliceData[66]	; tops[11][1] = sliceData[67]	; tops[11][2] = sliceData[68]	; tops[11][3] = sliceData[69]	; tops[11][4] = sliceData[70]	; tops[11][5] = sliceData[71]	; tops[12][0] = sliceData[72]	; tops[12][1] = sliceData[73]	; tops[12][2] = sliceData[74]	; tops[12][3] = sliceData[75]	; tops[12][4] = sliceData[76]	; tops[12][5] = sliceData[77]	; tops[13][0] = sliceData[78]	; tops[13][1] = sliceData[79]	; tops[13][2] = sliceData[80]	; tops[13][3] = sliceData[81]	; tops[13][4] = sliceData[82]	; tops[13][5] = sliceData[83]	; tops[14][0] = sliceData[84]	; tops[14][1] = sliceData[85]	; tops[14][2] = sliceData[86]	; tops[14][3] = sliceData[87]	; tops[14][4] = sliceData[88]	; tops[14][5] = sliceData[89]	; tops[15][0] = sliceData[90]	; tops[15][1] = sliceData[91]	; tops[15][2] = sliceData[92]	; tops[15][3] = sliceData[93]	; tops[15][4] = sliceData[94]	; tops[15][5] = sliceData[95]	; tops[16][0] = sliceData[96]	; tops[16][1] = sliceData[97]	; tops[16][2] = sliceData[98]	; tops[16][3] = sliceData[99]	; tops[16][4] = sliceData[100]	; tops[16][5] = sliceData[101]	; tops[17][0] = sliceData[102]	; tops[17][1] = sliceData[103]	; tops[17][2] = sliceData[104]	; tops[17][3] = sliceData[105]	; tops[17][4] = sliceData[106]	; tops[17][5] = sliceData[107]	; tops[18][0] = sliceData[108]	; tops[18][1] = sliceData[109]	; tops[18][2] = sliceData[110]	; tops[18][3] = sliceData[111]	; tops[18][4] = sliceData[112]	; tops[18][5] = sliceData[113]	; tops[19][0] = sliceData[114]	; tops[19][1] = sliceData[115]	; tops[19][2] = sliceData[116]	; tops[19][3] = sliceData[117]	; tops[19][4] = sliceData[118]	; tops[19][5] = sliceData[119]	; tops[20][0] = sliceData[120]	; tops[20][1] = sliceData[121]	; tops[20][2] = sliceData[122]	; tops[20][3] = sliceData[123]	; tops[20][4] = sliceData[124]	; tops[20][5] = sliceData[125]	; tops[21][0] = sliceData[126]	; tops[21][1] = sliceData[127]	; tops[21][2] = sliceData[128]	; tops[21][3] = sliceData[129]	; tops[21][4] = sliceData[130]	; tops[21][5] = sliceData[131]	; tops[22][0] = sliceData[132]	; tops[22][1] = sliceData[133]	; tops[22][2] = sliceData[134]	; tops[22][3] = sliceData[135]	; tops[22][4] = sliceData[136]	; tops[22][5] = sliceData[137]	; tops[23][0] = sliceData[138]	; tops[23][1] = sliceData[139]	; tops[23][2] = sliceData[140]	; tops[23][3] = sliceData[141]	; tops[23][4] = sliceData[142]	; tops[23][5] = sliceData[143]	; tops[24][0] = sliceData[144]	; tops[24][1] = sliceData[145]	; tops[24][2] = sliceData[146]	; tops[24][3] = sliceData[147]	; tops[24][4] = sliceData[148]	; tops[24][5] = sliceData[149]	; tops[25][0] = sliceData[150]	; tops[25][1] = sliceData[151]	; tops[25][2] = sliceData[152]	; tops[25][3] = sliceData[153]	; tops[25][4] = sliceData[154]	; tops[25][5] = sliceData[155]	; tops[26][0] = sliceData[156]	; tops[26][1] = sliceData[157]	; tops[26][2] = sliceData[158]	; tops[26][3] = sliceData[159]	; tops[26][4] = sliceData[160]	; tops[26][5] = sliceData[161]	; tops[27][0] = sliceData[162]	; tops[27][1] = sliceData[163]	; tops[27][2] = sliceData[164]	; tops[27][3] = sliceData[165]	; tops[27][4] = sliceData[166]	; tops[27][5] = sliceData[167]	; tops[28][0] = sliceData[168]	; tops[28][1] = sliceData[169]	; tops[28][2] = sliceData[170]	; tops[28][3] = sliceData[171]	; tops[28][4] = sliceData[172]	; tops[28][5] = sliceData[173]	; tops[29][0] = sliceData[174]	; tops[29][1] = sliceData[175]	; tops[29][2] = sliceData[176]	; tops[29][3] = sliceData[177]	; tops[29][4] = sliceData[178]	; tops[29][5] = sliceData[179]	; tops[30][0] = sliceData[180]	; tops[30][1] = sliceData[181]	; tops[30][2] = sliceData[182]	; tops[30][3] = sliceData[183]	; tops[30][4] = sliceData[184]	; tops[30][5] = sliceData[185]	; tops[31][0] = sliceData[186]	; tops[31][1] = sliceData[187]	; tops[31][2] = sliceData[188]	; tops[31][3] = sliceData[189]	; tops[31][4] = sliceData[190]	; tops[31][5] = sliceData[191]	; tops[32][0] = sliceData[192]	; tops[32][1] = sliceData[193]	; tops[32][2] = sliceData[194]	; tops[32][3] = sliceData[195]	; tops[32][4] = sliceData[196]	; tops[32][5] = sliceData[197]	; tops[33][0] = sliceData[198]	; tops[33][1] = sliceData[199]	; tops[33][2] = sliceData[200]	; tops[33][3] = sliceData[201]	; tops[33][4] = sliceData[202]	; tops[33][5] = sliceData[203]	; tops[34][0] = sliceData[204]	; tops[34][1] = sliceData[205]	; tops[34][2] = sliceData[206]	; tops[34][3] = sliceData[207]	; tops[34][4] = sliceData[208]	; tops[34][5] = sliceData[209]	; tops[35][0] = sliceData[210]	; tops[35][1] = sliceData[211]	; tops[35][2] = sliceData[212]	; tops[35][3] = sliceData[213]	; tops[35][4] = sliceData[214]	; tops[35][5] = sliceData[215]	; tops[36][0] = sliceData[216]	; tops[36][1] = sliceData[217]	; tops[36][2] = sliceData[218]	; tops[36][3] = sliceData[219]	; tops[36][4] = sliceData[220]	; tops[36][5] = sliceData[221]	; tops[37][0] = sliceData[222]	; tops[37][1] = sliceData[223]	; tops[37][2] = sliceData[224]	; tops[37][3] = sliceData[225]	; tops[37][4] = sliceData[226]	; tops[37][5] = sliceData[227]	; tops[38][0] = sliceData[228]	; tops[38][1] = sliceData[229]	; tops[38][2] = sliceData[230]	; tops[38][3] = sliceData[231]	; tops[38][4] = sliceData[232]	; tops[38][5] = sliceData[233]	; tops[39][0] = sliceData[234]	; tops[39][1] = sliceData[235]	; tops[39][2] = sliceData[236]	; tops[39][3] = sliceData[237]	; tops[39][4] = sliceData[238]	; tops[39][5] = sliceData[239]	; tops[40][0] = sliceData[240]	; tops[40][1] = sliceData[241]	; tops[40][2] = sliceData[242]	; tops[40][3] = sliceData[243]	; tops[40][4] = sliceData[244]	; tops[40][5] = sliceData[245]	; tops[41][0] = sliceData[246]	; tops[41][1] = sliceData[247]	; tops[41][2] = sliceData[248]	; tops[41][3] = sliceData[249]	; tops[41][4] = sliceData[250]	; tops[41][5] = sliceData[251]	; tops[42][0] = sliceData[252]	; tops[42][1] = sliceData[253]	; tops[42][2] = sliceData[254]	; tops[42][3] = sliceData[255]	; tops[42][4] = sliceData[256]	; tops[42][5] = sliceData[257]	; tops[43][0] = sliceData[258]	; tops[43][1] = sliceData[259]	; tops[43][2] = sliceData[260]	; tops[43][3] = sliceData[261]	; tops[43][4] = sliceData[262]	; tops[43][5] = sliceData[263]	; tops[44][0] = sliceData[264]	; tops[44][1] = sliceData[265]	; tops[44][2] = sliceData[266]	; tops[44][3] = sliceData[267]	; tops[44][4] = sliceData[268]	; tops[44][5] = sliceData[269]	; tops[45][0] = sliceData[270]	; tops[45][1] = sliceData[271]	; tops[45][2] = sliceData[272]	; tops[45][3] = sliceData[273]	; tops[45][4] = sliceData[274]	; tops[45][5] = sliceData[275]	; tops[46][0] = sliceData[276]	; tops[46][1] = sliceData[277]	; tops[46][2] = sliceData[278]	; tops[46][3] = sliceData[279]	; tops[46][4] = sliceData[280]	; tops[46][5] = sliceData[281]	; tops[47][0] = sliceData[282]	; tops[47][1] = sliceData[283]	; tops[47][2] = sliceData[284]	; tops[47][3] = sliceData[285]	; tops[47][4] = sliceData[286]	; tops[47][5] = sliceData[287]	; tops[48][0] = sliceData[288]	; tops[48][1] = sliceData[289]	; tops[48][2] = sliceData[290]	; tops[48][3] = sliceData[291]	; tops[48][4] = sliceData[292]	; tops[48][5] = sliceData[293]	; tops[49][0] = sliceData[294]	; tops[49][1] = sliceData[295]	; tops[49][2] = sliceData[296]	; tops[49][3] = sliceData[297]	; tops[49][4] = sliceData[298]	; tops[49][5] = sliceData[299]	; tops[50][0] = sliceData[300]	; tops[50][1] = sliceData[301]	; tops[50][2] = sliceData[302]	; tops[50][3] = sliceData[303]	; tops[50][4] = sliceData[304]	; tops[50][5] = sliceData[305]	; tops[51][0] = sliceData[306]	; tops[51][1] = sliceData[307]	; tops[51][2] = sliceData[308]	; tops[51][3] = sliceData[309]	; tops[51][4] = sliceData[310]	; tops[51][5] = sliceData[311]	; tops[52][0] = sliceData[312]	; tops[52][1] = sliceData[313]	; tops[52][2] = sliceData[314]	; tops[52][3] = sliceData[315]	; tops[52][4] = sliceData[316]	; tops[52][5] = sliceData[317]	; tops[53][0] = sliceData[318]	; tops[53][1] = sliceData[319]	; tops[53][2] = sliceData[320]	; tops[53][3] = sliceData[321]	; tops[53][4] = sliceData[322]	; tops[53][5] = sliceData[323]	; tops[54][0] = sliceData[324]	; tops[54][1] = sliceData[325]	; tops[54][2] = sliceData[326]	; tops[54][3] = sliceData[327]	; tops[54][4] = sliceData[328]	; tops[54][5] = sliceData[329]	; tops[55][0] = sliceData[330]	; tops[55][1] = sliceData[331]	; tops[55][2] = sliceData[332]	; tops[55][3] = sliceData[333]	; tops[55][4] = sliceData[334]	; tops[55][5] = sliceData[335]	; tops[56][0] = sliceData[336]	; tops[56][1] = sliceData[337]	; tops[56][2] = sliceData[338]	; tops[56][3] = sliceData[339]	; tops[56][4] = sliceData[340]	; tops[56][5] = sliceData[341]	; tops[57][0] = sliceData[342]	; tops[57][1] = sliceData[343]	; tops[57][2] = sliceData[344]	; tops[57][3] = sliceData[345]	; tops[57][4] = sliceData[346]	; tops[57][5] = sliceData[347]	; tops[58][0] = sliceData[348]	; tops[58][1] = sliceData[349]	; tops[58][2] = sliceData[350]	; tops[58][3] = sliceData[351]	; tops[58][4] = sliceData[352]	; tops[58][5] = sliceData[353]	; tops[59][0] = sliceData[354]	; tops[59][1] = sliceData[355]	; tops[59][2] = sliceData[356]	; tops[59][3] = sliceData[357]	; tops[59][4] = sliceData[358]	; tops[59][5] = sliceData[359]	; tops[60][0] = sliceData[360]	; tops[60][1] = sliceData[361]	; tops[60][2] = sliceData[362]	; tops[60][3] = sliceData[363]	; tops[60][4] = sliceData[364]	; tops[60][5] = sliceData[365]	; tops[61][0] = sliceData[366]	; tops[61][1] = sliceData[367]	; tops[61][2] = sliceData[368]	; tops[61][3] = sliceData[369]	; tops[61][4] = sliceData[370]	; tops[61][5] = sliceData[371]	; tops[62][0] = sliceData[372]	; tops[62][1] = sliceData[373]	; tops[62][2] = sliceData[374]	; tops[62][3] = sliceData[375]	; tops[62][4] = sliceData[376]	; tops[62][5] = sliceData[377]	; tops[63][0] = sliceData[378]	; tops[63][1] = sliceData[379]	; tops[63][2] = sliceData[380]	; tops[63][3] = sliceData[381]	; tops[63][4] = sliceData[382]	; tops[63][5] = sliceData[383]	; tops[64][0] = sliceData[384]	; tops[64][1] = sliceData[385]	; tops[64][2] = sliceData[386]	; tops[64][3] = sliceData[387]	; tops[64][4] = sliceData[388]	; tops[64][5] = sliceData[389]	; tops[65][0] = sliceData[390]	; tops[65][1] = sliceData[391]	; tops[65][2] = sliceData[392]	; tops[65][3] = sliceData[393]	; tops[65][4] = sliceData[394]	; tops[65][5] = sliceData[395]	; tops[66][0] = sliceData[396]	; tops[66][1] = sliceData[397]	; tops[66][2] = sliceData[398]	; tops[66][3] = sliceData[399]	; tops[66][4] = sliceData[400]	; tops[66][5] = sliceData[401]	; tops[67][0] = sliceData[402]	; tops[67][1] = sliceData[403]	; tops[67][2] = sliceData[404]	; tops[67][3] = sliceData[405]	; tops[67][4] = sliceData[406]	; tops[67][5] = sliceData[407]	; tops[68][0] = sliceData[408]	; tops[68][1] = sliceData[409]	; tops[68][2] = sliceData[410]	; tops[68][3] = sliceData[411]	; tops[68][4] = sliceData[412]	; tops[68][5] = sliceData[413]	; tops[69][0] = sliceData[414]	; tops[69][1] = sliceData[415]	; tops[69][2] = sliceData[416]	; tops[69][3] = sliceData[417]	; tops[69][4] = sliceData[418]	; tops[69][5] = sliceData[419]	; tops[70][0] = sliceData[420]	; tops[70][1] = sliceData[421]	; tops[70][2] = sliceData[422]	; tops[70][3] = sliceData[423]	; tops[70][4] = sliceData[424]	; tops[70][5] = sliceData[425]	; tops[71][0] = sliceData[426]	; tops[71][1] = sliceData[427]	; tops[71][2] = sliceData[428]	; tops[71][3] = sliceData[429]	; tops[71][4] = sliceData[430]	; tops[71][5] = sliceData[431]	; tops[72][0] = sliceData[432]	; tops[72][1] = sliceData[433]	; tops[72][2] = sliceData[434]	; tops[72][3] = sliceData[435]	; tops[72][4] = sliceData[436]	; tops[72][5] = sliceData[437]	; tops[73][0] = sliceData[438]	; tops[73][1] = sliceData[439]	; tops[73][2] = sliceData[440]	; tops[73][3] = sliceData[441]	; tops[73][4] = sliceData[442]	; tops[73][5] = sliceData[443]	; tops[74][0] = sliceData[444]	; tops[74][1] = sliceData[445]	; tops[74][2] = sliceData[446]	; tops[74][3] = sliceData[447]	; tops[74][4] = sliceData[448]	; tops[74][5] = sliceData[449]	; tops[75][0] = sliceData[450]	; tops[75][1] = sliceData[451]	; tops[75][2] = sliceData[452]	; tops[75][3] = sliceData[453]	; tops[75][4] = sliceData[454]	; tops[75][5] = sliceData[455]	; tops[76][0] = sliceData[456]	; tops[76][1] = sliceData[457]	; tops[76][2] = sliceData[458]	; tops[76][3] = sliceData[459]	; tops[76][4] = sliceData[460]	; tops[76][5] = sliceData[461]	; tops[77][0] = sliceData[462]	; tops[77][1] = sliceData[463]	; tops[77][2] = sliceData[464]	; tops[77][3] = sliceData[465]	; tops[77][4] = sliceData[466]	; tops[77][5] = sliceData[467]	; tops[78][0] = sliceData[468]	; tops[78][1] = sliceData[469]	; tops[78][2] = sliceData[470]	; tops[78][3] = sliceData[471]	; tops[78][4] = sliceData[472]	; tops[78][5] = sliceData[473]	; tops[79][0] = sliceData[474]	; tops[79][1] = sliceData[475]	; tops[79][2] = sliceData[476]	; tops[79][3] = sliceData[477]	; tops[79][4] = sliceData[478]	; tops[79][5] = sliceData[479]	; tops[80][0] = sliceData[480]	; tops[80][1] = sliceData[481]	; tops[80][2] = sliceData[482]	; tops[80][3] = sliceData[483]	; tops[80][4] = sliceData[484]	; tops[80][5] = sliceData[485]	; tops[81][0] = sliceData[486]	; tops[81][1] = sliceData[487]	; tops[81][2] = sliceData[488]	; tops[81][3] = sliceData[489]	; tops[81][4] = sliceData[490]	; tops[81][5] = sliceData[491]	; tops[82][0] = sliceData[492]	; tops[82][1] = sliceData[493]	; tops[82][2] = sliceData[494]	; tops[82][3] = sliceData[495]	; tops[82][4] = sliceData[496]	; tops[82][5] = sliceData[497]	; tops[83][0] = sliceData[498]	; tops[83][1] = sliceData[499]	; tops[83][2] = sliceData[500]	; tops[83][3] = sliceData[501]	; tops[83][4] = sliceData[502]	; tops[83][5] = sliceData[503]	; tops[84][0] = sliceData[504]	; tops[84][1] = sliceData[505]	; tops[84][2] = sliceData[506]	; tops[84][3] = sliceData[507]	; tops[84][4] = sliceData[508]	; tops[84][5] = sliceData[509]	; tops[85][0] = sliceData[510]	; tops[85][1] = sliceData[511]	; tops[85][2] = sliceData[512]	; tops[85][3] = sliceData[513]	; tops[85][4] = sliceData[514]	; tops[85][5] = sliceData[515]	; tops[86][0] = sliceData[516]	; tops[86][1] = sliceData[517]	; tops[86][2] = sliceData[518]	; tops[86][3] = sliceData[519]	; tops[86][4] = sliceData[520]	; tops[86][5] = sliceData[521]	; tops[87][0] = sliceData[522]	; tops[87][1] = sliceData[523]	; tops[87][2] = sliceData[524]	; tops[87][3] = sliceData[525]	; tops[87][4] = sliceData[526]	; tops[87][5] = sliceData[527]	; tops[88][0] = sliceData[528]	; tops[88][1] = sliceData[529]	; tops[88][2] = sliceData[530]	; tops[88][3] = sliceData[531]	; tops[88][4] = sliceData[532]	; tops[88][5] = sliceData[533]	; tops[89][0] = sliceData[534]	; tops[89][1] = sliceData[535]	; tops[89][2] = sliceData[536]	; tops[89][3] = sliceData[537]	; tops[89][4] = sliceData[538]	; tops[89][5] = sliceData[539]	; tops[90][0] = sliceData[540]	; tops[90][1] = sliceData[541]	; tops[90][2] = sliceData[542]	; tops[90][3] = sliceData[543]	; tops[90][4] = sliceData[544]	; tops[90][5] = sliceData[545]	; tops[91][0] = sliceData[546]	; tops[91][1] = sliceData[547]	; tops[91][2] = sliceData[548]	; tops[91][3] = sliceData[549]	; tops[91][4] = sliceData[550]	; tops[91][5] = sliceData[551]	; tops[92][0] = sliceData[552]	; tops[92][1] = sliceData[553]	; tops[92][2] = sliceData[554]	; tops[92][3] = sliceData[555]	; tops[92][4] = sliceData[556]	; tops[92][5] = sliceData[557]	; tops[93][0] = sliceData[558]	; tops[93][1] = sliceData[559]	; tops[93][2] = sliceData[560]	; tops[93][3] = sliceData[561]	; tops[93][4] = sliceData[562]	; tops[93][5] = sliceData[563]	; tops[94][0] = sliceData[564]	; tops[94][1] = sliceData[565]	; tops[94][2] = sliceData[566]	; tops[94][3] = sliceData[567]	; tops[94][4] = sliceData[568]	; tops[94][5] = sliceData[569]	; tops[95][0] = sliceData[570]	; tops[95][1] = sliceData[571]	; tops[95][2] = sliceData[572]	; tops[95][3] = sliceData[573]	; tops[95][4] = sliceData[574]	; tops[95][5] = sliceData[575]	; tops[96][0] = sliceData[576]	; tops[96][1] = sliceData[577]	; tops[96][2] = sliceData[578]	; tops[96][3] = sliceData[579]	; tops[96][4] = sliceData[580]	; tops[96][5] = sliceData[581]	; tops[97][0] = sliceData[582]	; tops[97][1] = sliceData[583]	; tops[97][2] = sliceData[584]	; tops[97][3] = sliceData[585]	; tops[97][4] = sliceData[586]	; tops[97][5] = sliceData[587]	; tops[98][0] = sliceData[588]	; tops[98][1] = sliceData[589]	; tops[98][2] = sliceData[590]	; tops[98][3] = sliceData[591]	; tops[98][4] = sliceData[592]	; tops[98][5] = sliceData[593]	; tops[99][0] = sliceData[594]	; tops[99][1] = sliceData[595]	; tops[99][2] = sliceData[596]	; tops[99][3] = sliceData[597]	; tops[99][4] = sliceData[598]	; tops[99][5] = sliceData[599]	; tops[100][0] = sliceData[600]	; tops[100][1] = sliceData[601]	; tops[100][2] = sliceData[602]	; tops[100][3] = sliceData[603]	; tops[100][4] = sliceData[604]	; tops[100][5] = sliceData[605]	; tops[101][0] = sliceData[606]	; tops[101][1] = sliceData[607]	; tops[101][2] = sliceData[608]	; tops[101][3] = sliceData[609]	; tops[101][4] = sliceData[610]	; tops[101][5] = sliceData[611]	; tops[102][0] = sliceData[612]	; tops[102][1] = sliceData[613]	; tops[102][2] = sliceData[614]	; tops[102][3] = sliceData[615]	; tops[102][4] = sliceData[616]	; tops[102][5] = sliceData[617]	; tops[103][0] = sliceData[618]	; tops[103][1] = sliceData[619]	; tops[103][2] = sliceData[620]	; tops[103][3] = sliceData[621]	; tops[104][0] = sliceData[622]	; tops[104][1] = sliceData[623]	; tops[104][2] = sliceData[624]	; tops[104][3] = sliceData[625]	; tops[104][4] = sliceData[626]	; tops[104][5] = sliceData[627]	; tops[105][0] = sliceData[628]	; tops[105][1] = sliceData[629]	; tops[105][2] = sliceData[630]	; tops[105][3] = sliceData[631]	; tops[105][4] = sliceData[632]	; tops[105][5] = sliceData[633]	; tops[106][0] = sliceData[634]	; tops[106][1] = sliceData[635]	; tops[106][2] = sliceData[636]	; tops[106][3] = sliceData[637]	; tops[106][4] = sliceData[638]	; tops[106][5] = sliceData[639]	; tops[107][0] = sliceData[640]	; tops[107][1] = sliceData[641]	; tops[107][2] = sliceData[642]	; tops[107][3] = sliceData[643]	; tops[107][4] = sliceData[644]	; tops[107][5] = sliceData[645]	; tops[108][0] = sliceData[646]	; tops[108][1] = sliceData[647]	; tops[108][2] = sliceData[648]	; tops[108][3] = sliceData[649]	; tops[108][4] = sliceData[650]	; tops[108][5] = sliceData[651]	; tops[109][0] = sliceData[652]	; tops[109][1] = sliceData[653]	; tops[109][2] = sliceData[654]	; tops[109][3] = sliceData[655]	; tops[109][4] = sliceData[656]	; tops[109][5] = sliceData[657]	; tops[110][0] = sliceData[658]	; tops[110][1] = sliceData[659]	; tops[110][2] = sliceData[660]	; tops[110][3] = sliceData[661]	; tops[110][4] = sliceData[662]	; tops[110][5] = sliceData[663]	; tops[111][0] = sliceData[664]	; tops[111][1] = sliceData[665]	; tops[111][2] = sliceData[666]	; tops[111][3] = sliceData[667]	; tops[111][4] = sliceData[668]	; tops[111][5] = sliceData[669]	; tops[112][0] = sliceData[670]	; tops[112][1] = sliceData[671]	; tops[112][2] = sliceData[672]	; tops[112][3] = sliceData[673]	; tops[112][4] = sliceData[674]	; tops[112][5] = sliceData[675]	; tops[113][0] = sliceData[676]	; tops[113][1] = sliceData[677]	; tops[113][2] = sliceData[678]	; tops[113][3] = sliceData[679]	; tops[113][4] = sliceData[680]	; tops[113][5] = sliceData[681]	; tops[114][0] = sliceData[682]	; tops[114][1] = sliceData[683]	; tops[114][2] = sliceData[684]	; tops[114][3] = sliceData[685]	; tops[114][4] = sliceData[686]	; tops[114][5] = sliceData[687]	; tops[115][0] = sliceData[688]	; tops[115][1] = sliceData[689]	; tops[115][2] = sliceData[690]	; tops[115][3] = sliceData[691]	; tops[115][4] = sliceData[692]	; tops[115][5] = sliceData[693]	; tops[116][0] = sliceData[694]	; tops[116][1] = sliceData[695]	; tops[116][2] = sliceData[696]	; tops[116][3] = sliceData[697]	; tops[116][4] = sliceData[698]	; tops[116][5] = sliceData[699]	; tops[117][0] = sliceData[700]	; tops[117][1] = sliceData[701]	; tops[117][2] = sliceData[702]	; tops[117][3] = sliceData[703]	; tops[117][4] = sliceData[704]	; tops[117][5] = sliceData[705]	; tops[118][0] = sliceData[706]	; tops[118][1] = sliceData[707]	; tops[118][2] = sliceData[708]	; tops[118][3] = sliceData[709]	; tops[118][4] = sliceData[710]	; tops[118][5] = sliceData[711]	; tops[119][0] = sliceData[712]	; tops[119][1] = sliceData[713]	; tops[119][2] = sliceData[714]	; tops[119][3] = sliceData[715]	; tops[119][4] = sliceData[716]	; tops[119][5] = sliceData[717]	; tops[120][0] = sliceData[718]	; tops[120][1] = sliceData[719]	; tops[120][2] = sliceData[720]	; tops[120][3] = sliceData[721]	; tops[120][4] = sliceData[722]	; tops[120][5] = sliceData[723]	; tops[121][0] = sliceData[724]	; tops[121][1] = sliceData[725]	; tops[121][2] = sliceData[726]	; tops[121][3] = sliceData[727]	; tops[121][4] = sliceData[728]	; tops[121][5] = sliceData[729]	; tops[122][0] = sliceData[730]	; tops[122][1] = sliceData[731]	; tops[122][2] = sliceData[732]	; tops[122][3] = sliceData[733]	; tops[122][4] = sliceData[734]	; tops[122][5] = sliceData[735]	; tops[123][0] = sliceData[736]	; tops[123][1] = sliceData[737]	; tops[123][2] = sliceData[738]	; tops[123][3] = sliceData[739]	; tops[123][4] = sliceData[740]	; tops[123][5] = sliceData[741]	; tops[124][0] = sliceData[742]	; tops[124][1] = sliceData[743]	; tops[124][2] = sliceData[744]	; tops[124][3] = sliceData[745]	; tops[124][4] = sliceData[746]	; tops[124][5] = sliceData[747]	; tops[125][0] = sliceData[748]	; tops[125][1] = sliceData[749]	; tops[125][2] = sliceData[750]	; tops[125][3] = sliceData[751]	; tops[125][4] = sliceData[752]	; tops[125][5] = sliceData[753]	; tops[126][0] = sliceData[754]	; tops[126][1] = sliceData[755]	; tops[126][2] = sliceData[756]	; tops[126][3] = sliceData[757]	; tops[126][4] = sliceData[758]	; tops[126][5] = sliceData[759]	; tops[127][0] = sliceData[760]	; tops[127][1] = sliceData[761]	; tops[127][2] = sliceData[762]	; tops[127][3] = sliceData[763]	; tops[127][4] = sliceData[764]	; tops[127][5] = sliceData[765]	; tops[128][0] = sliceData[766]	; tops[128][1] = sliceData[767]	; tops[128][2] = sliceData[768]	; tops[128][3] = sliceData[769]	; tops[128][4] = sliceData[770]	; tops[128][5] = sliceData[771]	; tops[129][0] = sliceData[772]	; tops[129][1] = sliceData[773]	; tops[129][2] = sliceData[774]	; tops[129][3] = sliceData[775]	; tops[129][4] = sliceData[776]	; tops[129][5] = sliceData[777]	; tops[130][0] = sliceData[778]	; tops[130][1] = sliceData[779]	; tops[130][2] = sliceData[780]	; tops[130][3] = sliceData[781]	; tops[130][4] = sliceData[782]	; tops[130][5] = sliceData[783]	; tops[131][0] = sliceData[784]	; tops[131][1] = sliceData[785]	; tops[131][2] = sliceData[786]	; tops[131][3] = sliceData[787]	; tops[131][4] = sliceData[788]	; tops[131][5] = sliceData[789]	; tops[132][0] = sliceData[790]	; tops[132][1] = sliceData[791]	; tops[132][2] = sliceData[792]	; tops[132][3] = sliceData[793]	; tops[132][4] = sliceData[794]	; tops[132][5] = sliceData[795]	; tops[133][0] = sliceData[796]	; tops[133][1] = sliceData[797]	; tops[133][2] = sliceData[798]	; tops[133][3] = sliceData[799]	; tops[133][4] = sliceData[800]	; tops[133][5] = sliceData[801]	; tops[134][0] = sliceData[802]	; tops[134][1] = sliceData[803]	; tops[134][2] = sliceData[804]	; tops[134][3] = sliceData[805]	; tops[134][4] = sliceData[806]	; tops[134][5] = sliceData[807]	; tops[135][0] = sliceData[808]	; tops[135][1] = sliceData[809]	; tops[135][2] = sliceData[810]	; tops[135][3] = sliceData[811]	; tops[135][4] = sliceData[812]	; tops[135][5] = sliceData[813]	; tops[136][0] = sliceData[814]	; tops[136][1] = sliceData[815]	; tops[136][2] = sliceData[816]	; tops[136][3] = sliceData[817]	; tops[136][4] = sliceData[818]	; tops[136][5] = sliceData[819]	; tops[137][0] = sliceData[820]	; tops[137][1] = sliceData[821]	; tops[137][2] = sliceData[822]	; tops[137][3] = sliceData[823]	; tops[138][0] = sliceData[824]	; tops[138][1] = sliceData[825]	; tops[138][2] = sliceData[826]	; tops[138][3] = sliceData[827]	; tops[138][4] = sliceData[828]	; tops[138][5] = sliceData[829]	; tops[139][0] = sliceData[830]	; tops[139][1] = sliceData[831]	; tops[139][2] = sliceData[832]	; tops[139][3] = sliceData[833]	; tops[139][4] = sliceData[834]	; tops[139][5] = sliceData[835]	; tops[140][0] = sliceData[836]	; tops[140][1] = sliceData[837]	; tops[140][2] = sliceData[838]	; tops[140][3] = sliceData[839]	; tops[140][4] = sliceData[840]	; tops[140][5] = sliceData[841]	; tops[141][0] = sliceData[842]	; tops[141][1] = sliceData[843]	; tops[141][2] = sliceData[844]	; tops[141][3] = sliceData[845]	; tops[141][4] = sliceData[846]	; tops[141][5] = sliceData[847]	; tops[142][0] = sliceData[848]	; tops[142][1] = sliceData[849]	; tops[142][2] = sliceData[850]	; tops[142][3] = sliceData[851]	; tops[142][4] = sliceData[852]	; tops[142][5] = sliceData[853]	; tops[143][0] = sliceData[854]	; tops[143][1] = sliceData[855]	; tops[143][2] = sliceData[856]	; tops[143][3] = sliceData[857]	; tops[143][4] = sliceData[858]	; tops[143][5] = sliceData[859]	; tops[144][0] = sliceData[860]	; tops[144][1] = sliceData[861]	; tops[144][2] = sliceData[862]	; tops[144][3] = sliceData[863]	; tops[144][4] = sliceData[864]	; tops[144][5] = sliceData[865]	; tops[145][0] = sliceData[866]	; tops[145][1] = sliceData[867]	; tops[145][2] = sliceData[868]	; tops[145][3] = sliceData[869]	; tops[145][4] = sliceData[870]	; tops[145][5] = sliceData[871]	; tops[146][0] = sliceData[872]	; tops[146][1] = sliceData[873]	; tops[146][2] = sliceData[874]	; tops[146][3] = sliceData[875]	; tops[146][4] = sliceData[876]	; tops[146][5] = sliceData[877]	; tops[147][0] = sliceData[878]	; tops[147][1] = sliceData[879]	; tops[147][2] = sliceData[880]	; tops[147][3] = sliceData[881]	; tops[147][4] = sliceData[882]	; tops[147][5] = sliceData[883]	; tops[148][0] = sliceData[884]	; tops[148][1] = sliceData[885]	; tops[148][2] = sliceData[886]	; tops[148][3] = sliceData[887]	; tops[148][4] = sliceData[888]	; tops[148][5] = sliceData[889]	; tops[149][0] = sliceData[890]	; tops[149][1] = sliceData[891]	; tops[149][2] = sliceData[892]	; tops[149][3] = sliceData[893]	; tops[149][4] = sliceData[894]	; tops[149][5] = sliceData[895]	; tops[150][0] = sliceData[896]	; tops[150][1] = sliceData[897]	; tops[150][2] = sliceData[898]	; tops[150][3] = sliceData[899]	; tops[150][4] = sliceData[900]	; tops[150][5] = sliceData[901]	; tops[151][0] = sliceData[902]	; tops[151][1] = sliceData[903]	; tops[151][2] = sliceData[904]	; tops[151][3] = sliceData[905]	; tops[151][4] = sliceData[906]	; tops[151][5] = sliceData[907]	; tops[152][0] = sliceData[908]	; tops[152][1] = sliceData[909]	; tops[152][2] = sliceData[910]	; tops[152][3] = sliceData[911]	; tops[152][4] = sliceData[912]	; tops[152][5] = sliceData[913]	; tops[153][0] = sliceData[914]	; tops[153][1] = sliceData[915]	; tops[153][2] = sliceData[916]	; tops[153][3] = sliceData[917]	; tops[153][4] = sliceData[918]	; tops[153][5] = sliceData[919]	; tops[154][0] = sliceData[920]	; tops[154][1] = sliceData[921]	; tops[154][2] = sliceData[922]	; tops[154][3] = sliceData[923]	; tops[154][4] = sliceData[924]	; tops[154][5] = sliceData[925]	; tops[155][0] = sliceData[926]	; tops[155][1] = sliceData[927]	; tops[155][2] = sliceData[928]	; tops[155][3] = sliceData[929]	; tops[155][4] = sliceData[930]	; tops[155][5] = sliceData[931]	; tops[156][0] = sliceData[932]	; tops[156][1] = sliceData[933]	; tops[156][2] = sliceData[934]	; tops[156][3] = sliceData[935]	; tops[156][4] = sliceData[936]	; tops[156][5] = sliceData[937]	; tops[157][0] = sliceData[938]	; tops[157][1] = sliceData[939]	; tops[157][2] = sliceData[940]	; tops[157][3] = sliceData[941]	; tops[157][4] = sliceData[942]	; tops[157][5] = sliceData[943]	; tops[158][0] = sliceData[944]	; tops[158][1] = sliceData[945]	; tops[158][2] = sliceData[946]	; tops[158][3] = sliceData[947]	; tops[158][4] = sliceData[948]	; tops[158][5] = sliceData[949]	; tops[159][0] = sliceData[950]	; tops[159][1] = sliceData[951]	; tops[159][2] = sliceData[952]	; tops[159][3] = sliceData[953]	; tops[159][4] = sliceData[954]	; tops[159][5] = sliceData[955]	; tops[160][0] = sliceData[956]	; tops[160][1] = sliceData[957]	; tops[160][2] = sliceData[958]	; tops[160][3] = sliceData[959]	; tops[160][4] = sliceData[960]	; tops[160][5] = sliceData[961]	; tops[161][0] = sliceData[962]	; tops[161][1] = sliceData[963]	; tops[161][2] = sliceData[964]	; tops[161][3] = sliceData[965]	; tops[161][4] = sliceData[966]	; tops[161][5] = sliceData[967]	; tops[162][0] = sliceData[968]	; tops[162][1] = sliceData[969]	; tops[162][2] = sliceData[970]	; tops[162][3] = sliceData[971]	; tops[162][4] = sliceData[972]	; tops[162][5] = sliceData[973]	; tops[163][0] = sliceData[974]	; tops[163][1] = sliceData[975]	; tops[163][2] = sliceData[976]	; tops[163][3] = sliceData[977]	; tops[163][4] = sliceData[978]	; tops[163][5] = sliceData[979]	; tops[164][0] = sliceData[980]	; tops[164][1] = sliceData[981]	; tops[164][2] = sliceData[982]	; tops[164][3] = sliceData[983]	; tops[164][4] = sliceData[984]	; tops[164][5] = sliceData[985]	; tops[165][0] = sliceData[986]	; tops[165][1] = sliceData[987]	; tops[165][2] = sliceData[988]	; tops[165][3] = sliceData[989]	; tops[165][4] = sliceData[990]	; tops[165][5] = sliceData[991]	; tops[166][0] = sliceData[992]	; tops[166][1] = sliceData[993]	; tops[166][2] = sliceData[994]	; tops[166][3] = sliceData[995]	; tops[166][4] = sliceData[996]	; tops[166][5] = sliceData[997]	; tops[167][0] = sliceData[998]	; tops[167][1] = sliceData[999]	; tops[167][2] = sliceData[1000]	; tops[167][3] = sliceData[1001]	; tops[167][4] = sliceData[1002]	; tops[167][5] = sliceData[1003]	; tops[168][0] = sliceData[1004]	; tops[168][1] = sliceData[1005]	; tops[168][2] = sliceData[1006]	; tops[168][3] = sliceData[1007]	; tops[168][4] = sliceData[1008]	; tops[168][5] = sliceData[1009]	; tops[169][0] = sliceData[1010]	; tops[169][1] = sliceData[1011]	; tops[169][2] = sliceData[1012]	; tops[169][3] = sliceData[1013]	; tops[169][4] = sliceData[1014]	; tops[169][5] = sliceData[1015]	; tops[170][0] = sliceData[1016]	; tops[170][1] = sliceData[1017]	; tops[170][2] = sliceData[1018]	; tops[170][3] = sliceData[1019]	; tops[170][4] = sliceData[1020]	; tops[170][5] = sliceData[1021]	; tops[171][0] = sliceData[1022]	; tops[171][1] = sliceData[1023]	; tops[171][2] = sliceData[1024]	; tops[171][3] = sliceData[1025]	; tops[171][4] = sliceData[1026]	; tops[171][5] = sliceData[1027]	; tops[172][0] = sliceData[1028]	; tops[172][1] = sliceData[1029]	; tops[172][2] = sliceData[1030]	; tops[172][3] = sliceData[1031]	; tops[172][4] = sliceData[1032]	; tops[172][5] = sliceData[1033]	; tops[173][0] = sliceData[1034]	; tops[173][1] = sliceData[1035]	; tops[173][2] = sliceData[1036]	; tops[173][3] = sliceData[1037]	; tops[173][4] = sliceData[1038]	; tops[173][5] = sliceData[1039]	; tops[174][0] = sliceData[1040]	; tops[174][1] = sliceData[1041]	; tops[174][2] = sliceData[1042]	; tops[174][3] = sliceData[1043]	; tops[174][4] = sliceData[1044]	; tops[174][5] = sliceData[1045]	; tops[175][0] = sliceData[1046]	; tops[175][1] = sliceData[1047]	; tops[175][2] = sliceData[1048]	; tops[175][3] = sliceData[1049]	; tops[175][4] = sliceData[1050]	; tops[175][5] = sliceData[1051]	; tops[176][0] = sliceData[1052]	; tops[176][1] = sliceData[1053]	; tops[176][2] = sliceData[1054]	; tops[176][3] = sliceData[1055]	; tops[176][4] = sliceData[1056]	; tops[176][5] = sliceData[1057]	; tops[177][0] = sliceData[1058]	; tops[177][1] = sliceData[1059]	; tops[177][2] = sliceData[1060]	; tops[177][3] = sliceData[1061]	; tops[177][4] = sliceData[1062]	; tops[177][5] = sliceData[1063]	; tops[178][0] = sliceData[1064]	; tops[178][1] = sliceData[1065]	; tops[178][2] = sliceData[1066]	; tops[178][3] = sliceData[1067]	; tops[178][4] = sliceData[1068]	; tops[178][5] = sliceData[1069]	; tops[179][0] = sliceData[1070]	; tops[179][1] = sliceData[1071]	; tops[179][2] = sliceData[1072]	; tops[179][3] = sliceData[1073]	; tops[179][4] = sliceData[1074]	; tops[179][5] = sliceData[1075]	; tops[180][0] = sliceData[1076]	; tops[180][1] = sliceData[1077]	; tops[180][2] = sliceData[1078]	; tops[180][3] = sliceData[1079]	; tops[180][4] = sliceData[1080]	; tops[180][5] = sliceData[1081]	; tops[181][0] = sliceData[1082]	; tops[181][1] = sliceData[1083]	; tops[181][2] = sliceData[1084]	; tops[181][3] = sliceData[1085]	; tops[181][4] = sliceData[1086]	; tops[181][5] = sliceData[1087]	; tops[182][0] = sliceData[1088]	; tops[182][1] = sliceData[1089]	; tops[182][2] = sliceData[1090]	; tops[182][3] = sliceData[1091]	; tops[182][4] = sliceData[1092]	; tops[182][5] = sliceData[1093]	; tops[183][0] = sliceData[1094]	; tops[183][1] = sliceData[1095]	; tops[183][2] = sliceData[1096]	; tops[183][3] = sliceData[1097]	; tops[183][4] = sliceData[1098]	; tops[183][5] = sliceData[1099]	; tops[184][0] = sliceData[1100]	; tops[184][1] = sliceData[1101]	; tops[184][2] = sliceData[1102]	; tops[184][3] = sliceData[1103]	; tops[184][4] = sliceData[1104]	; tops[184][5] = sliceData[1105]	; tops[185][0] = sliceData[1106]	; tops[185][1] = sliceData[1107]	; tops[185][2] = sliceData[1108]	; tops[185][3] = sliceData[1109]	; tops[185][4] = sliceData[1110]	; tops[185][5] = sliceData[1111]	; tops[186][0] = sliceData[1112]	; tops[186][1] = sliceData[1113]	; tops[186][2] = sliceData[1114]	; tops[186][3] = sliceData[1115]	; tops[186][4] = sliceData[1116]	; tops[186][5] = sliceData[1117]	; tops[187][0] = sliceData[1118]	; tops[187][1] = sliceData[1119]	; tops[187][2] = sliceData[1120]	; tops[187][3] = sliceData[1121]	; tops[187][4] = sliceData[1122]	; tops[187][5] = sliceData[1123]	; tops[188][0] = sliceData[1124]	; tops[188][1] = sliceData[1125]	; tops[188][2] = sliceData[1126]	; tops[188][3] = sliceData[1127]	; tops[188][4] = sliceData[1128]	; tops[188][5] = sliceData[1129]	; tops[189][0] = sliceData[1130]	; tops[189][1] = sliceData[1131]	; tops[189][2] = sliceData[1132]	; tops[189][3] = sliceData[1133]	; tops[189][4] = sliceData[1134]	; tops[189][5] = sliceData[1135]	; tops[190][0] = sliceData[1136]	; tops[190][1] = sliceData[1137]	; tops[190][2] = sliceData[1138]	; tops[190][3] = sliceData[1139]	; tops[190][4] = sliceData[1140]	; tops[190][5] = sliceData[1141]	; tops[191][0] = sliceData[1142]	; tops[191][1] = sliceData[1143]	; tops[191][2] = sliceData[1144]	; tops[191][3] = sliceData[1145]	; tops[191][4] = sliceData[1146]	; tops[191][5] = sliceData[1147]	; tops[192][0] = sliceData[1148]	; tops[192][1] = sliceData[1149]	; tops[192][2] = sliceData[1150]	; tops[192][3] = sliceData[1151]	; tops[192][4] = sliceData[1152]	; tops[192][5] = sliceData[1153]	; tops[193][0] = sliceData[1154]	; tops[193][1] = sliceData[1155]	; tops[193][2] = sliceData[1156]	; tops[193][3] = sliceData[1157]	; tops[193][4] = sliceData[1158]	; tops[193][5] = sliceData[1159]	; tops[194][0] = sliceData[1160]	; tops[194][1] = sliceData[1161]	; tops[194][2] = sliceData[1162]	; tops[194][3] = sliceData[1163]	; tops[194][4] = sliceData[1164]	; tops[194][5] = sliceData[1165]	; tops[195][0] = sliceData[1166]	; tops[195][1] = sliceData[1167]	; tops[195][2] = sliceData[1168]	; tops[195][3] = sliceData[1169]	; tops[195][4] = sliceData[1170]	; tops[195][5] = sliceData[1171]	; tops[196][0] = sliceData[1172]	; tops[196][1] = sliceData[1173]	; tops[196][2] = sliceData[1174]	; tops[196][3] = sliceData[1175]	; tops[196][4] = sliceData[1176]	; tops[196][5] = sliceData[1177]	; tops[197][0] = sliceData[1178]	; tops[197][1] = sliceData[1179]	; tops[197][2] = sliceData[1180]	; tops[197][3] = sliceData[1181]	; tops[197][4] = sliceData[1182]	; tops[197][5] = sliceData[1183]	; tops[198][0] = sliceData[1184]	; tops[198][1] = sliceData[1185]	; tops[198][2] = sliceData[1186]	; tops[198][3] = sliceData[1187]	; tops[198][4] = sliceData[1188]	; tops[198][5] = sliceData[1189]	; tops[199][0] = sliceData[1190]	; tops[199][1] = sliceData[1191]	; tops[199][2] = sliceData[1192]	; tops[199][3] = sliceData[1193]	; tops[199][4] = sliceData[1194]	; tops[199][5] = sliceData[1195]	; tops[200][0] = sliceData[1196]	; tops[200][1] = sliceData[1197]	; tops[200][2] = sliceData[1198]	; tops[200][3] = sliceData[1199]	; tops[200][4] = sliceData[1200]	; tops[200][5] = sliceData[1201]	; tops[201][0] = sliceData[1202]	; tops[201][1] = sliceData[1203]	; tops[201][2] = sliceData[1204]	; tops[201][3] = sliceData[1205]	; tops[201][4] = sliceData[1206]	; tops[201][5] = sliceData[1207]	; tops[202][0] = sliceData[1208]	; tops[202][1] = sliceData[1209]	; tops[202][2] = sliceData[1210]	; tops[202][3] = sliceData[1211]	; tops[202][4] = sliceData[1212]	; tops[202][5] = sliceData[1213]	; tops[203][0] = sliceData[1214]	; tops[203][1] = sliceData[1215]	; tops[203][2] = sliceData[1216]	; tops[204][0] = sliceData[1217]	; tops[204][1] = sliceData[1218]	; tops[204][2] = sliceData[1219]	; tops[204][3] = sliceData[1220]	; tops[204][4] = sliceData[1221]	; tops[204][5] = sliceData[1222]	; tops[205][0] = sliceData[1223]	; tops[205][1] = sliceData[1224]	; tops[205][2] = sliceData[1225]	; tops[205][3] = sliceData[1226]	; tops[205][4] = sliceData[1227]	; tops[205][5] = sliceData[1228]	; tops[206][0] = sliceData[1229]	; tops[206][1] = sliceData[1230]	; tops[206][2] = sliceData[1231]	; tops[206][3] = sliceData[1232]	; tops[206][4] = sliceData[1233]	; tops[206][5] = sliceData[1234]	; tops[207][0] = sliceData[1235]	; tops[207][1] = sliceData[1236]	; tops[207][2] = sliceData[1237]	; tops[207][3] = sliceData[1238]	; tops[207][4] = sliceData[1239]	; tops[207][5] = sliceData[1240]	; tops[208][0] = sliceData[1241]	; tops[208][1] = sliceData[1242]	; tops[208][2] = sliceData[1243]	; tops[208][3] = sliceData[1244]	; tops[208][4] = sliceData[1245]	; tops[208][5] = sliceData[1246]	; tops[209][0] = sliceData[1247]	; tops[209][1] = sliceData[1248]	; tops[209][2] = sliceData[1249]	; tops[209][3] = sliceData[1250]	; tops[209][4] = sliceData[1251]	; tops[209][5] = sliceData[1252]	; tops[210][0] = sliceData[1253]	; tops[210][1] = sliceData[1254]	; tops[210][2] = sliceData[1255]	; tops[210][3] = sliceData[1256]	; tops[210][4] = sliceData[1257]	; tops[210][5] = sliceData[1258]	; tops[211][0] = sliceData[1259]	; tops[211][1] = sliceData[1260]	; tops[211][2] = sliceData[1261]	; tops[211][3] = sliceData[1262]	; tops[211][4] = sliceData[1263]	; tops[211][5] = sliceData[1264]	; tops[212][0] = sliceData[1265]	; tops[212][1] = sliceData[1266]	; tops[212][2] = sliceData[1267]	; tops[212][3] = sliceData[1268]	; tops[212][4] = sliceData[1269]	; tops[212][5] = sliceData[1270]	; tops[213][0] = sliceData[1271]	; tops[213][1] = sliceData[1272]	; tops[213][2] = sliceData[1273]	; tops[213][3] = sliceData[1274]	; tops[213][4] = sliceData[1275]	; tops[213][5] = sliceData[1276]	; tops[214][0] = sliceData[1277]	; tops[214][1] = sliceData[1278]	; tops[214][2] = sliceData[1279]	; tops[214][3] = sliceData[1280]	; tops[214][4] = sliceData[1281]	; tops[214][5] = sliceData[1282]	; tops[215][0] = sliceData[1283]	; tops[215][1] = sliceData[1284]	; tops[215][2] = sliceData[1285]	; tops[215][3] = sliceData[1286]	; tops[215][4] = sliceData[1287]	; tops[215][5] = sliceData[1288]	; tops[216][0] = sliceData[1289]	; tops[216][1] = sliceData[1290]	; tops[216][2] = sliceData[1291]	; tops[216][3] = sliceData[1292]	; tops[216][4] = sliceData[1293]	; tops[216][5] = sliceData[1294]	; tops[217][0] = sliceData[1295]	; tops[217][1] = sliceData[1296]	; tops[217][2] = sliceData[1297]	; tops[217][3] = sliceData[1298]	; tops[217][4] = sliceData[1299]	; tops[217][5] = sliceData[1300]	; tops[218][0] = sliceData[1301]	; tops[218][1] = sliceData[1302]	; tops[218][2] = sliceData[1303]	; tops[218][3] = sliceData[1304]	; tops[218][4] = sliceData[1305]	; tops[218][5] = sliceData[1306]	; tops[219][0] = sliceData[1307]	; tops[219][1] = sliceData[1308]	; tops[219][2] = sliceData[1309]	; tops[219][3] = sliceData[1310]	; tops[219][4] = sliceData[1311]	; tops[219][5] = sliceData[1312]	; tops[220][0] = sliceData[1313]	; tops[220][1] = sliceData[1314]	; tops[220][2] = sliceData[1315]	; tops[220][3] = sliceData[1316]	; tops[220][4] = sliceData[1317]	; tops[220][5] = sliceData[1318]	; tops[221][0] = sliceData[1319]	; tops[221][1] = sliceData[1320]	; tops[221][2] = sliceData[1321]	; tops[221][3] = sliceData[1322]	; tops[221][4] = sliceData[1323]	; tops[221][5] = sliceData[1324]	; tops[222][0] = sliceData[1325]	; tops[222][1] = sliceData[1326]	; tops[222][2] = sliceData[1327]	; tops[222][3] = sliceData[1328]	; tops[222][4] = sliceData[1329]	; tops[222][5] = sliceData[1330]	; tops[223][0] = sliceData[1331]	; tops[223][1] = sliceData[1332]	; tops[223][2] = sliceData[1333]	; tops[223][3] = sliceData[1334]	; tops[223][4] = sliceData[1335]	; tops[223][5] = sliceData[1336]	; tops[224][0] = sliceData[1337]	; tops[224][1] = sliceData[1338]	; tops[224][2] = sliceData[1339]	; tops[224][3] = sliceData[1340]	; tops[224][4] = sliceData[1341]	; tops[224][5] = sliceData[1342]	; tops[225][0] = sliceData[1343]	; tops[225][1] = sliceData[1344]	; tops[225][2] = sliceData[1345]	; tops[225][3] = sliceData[1346]	; tops[225][4] = sliceData[1347]	; tops[225][5] = sliceData[1348]	; tops[226][0] = sliceData[1349]	; tops[226][1] = sliceData[1350]	; tops[226][2] = sliceData[1351]	; tops[226][3] = sliceData[1352]	; tops[226][4] = sliceData[1353]	; tops[226][5] = sliceData[1354]	; tops[227][0] = sliceData[1355]	; tops[227][1] = sliceData[1356]	; tops[227][2] = sliceData[1357]	; tops[227][3] = sliceData[1358]	; tops[227][4] = sliceData[1359]	; tops[227][5] = sliceData[1360]	; tops[228][0] = sliceData[1361]	; tops[228][1] = sliceData[1362]	; tops[228][2] = sliceData[1363]	; tops[228][3] = sliceData[1364]	; tops[228][4] = sliceData[1365]	; tops[228][5] = sliceData[1366]	; tops[229][0] = sliceData[1367]	; tops[229][1] = sliceData[1368]	; tops[229][2] = sliceData[1369]	; tops[229][3] = sliceData[1370]	; tops[229][4] = sliceData[1371]	; tops[229][5] = sliceData[1372]	; tops[230][0] = sliceData[1373]	; tops[230][1] = sliceData[1374]	; tops[230][2] = sliceData[1375]	; tops[230][3] = sliceData[1376]	; tops[230][4] = sliceData[1377]	; tops[230][5] = sliceData[1378]	; tops[231][0] = sliceData[1379]	; tops[231][1] = sliceData[1380]	; tops[231][2] = sliceData[1381]	; tops[231][3] = sliceData[1382]	; tops[231][4] = sliceData[1383]	; tops[231][5] = sliceData[1384]	; tops[232][0] = sliceData[1385]	; tops[232][1] = sliceData[1386]	; tops[232][2] = sliceData[1387]	; tops[232][3] = sliceData[1388]	; tops[232][4] = sliceData[1389]	; tops[232][5] = sliceData[1390]	; tops[233][0] = sliceData[1391]	; tops[233][1] = sliceData[1392]	; tops[233][2] = sliceData[1393]	; tops[233][3] = sliceData[1394]	; tops[233][4] = sliceData[1395]	; tops[233][5] = sliceData[1396]	; tops[234][0] = sliceData[1397]	; tops[234][1] = sliceData[1398]	; tops[234][2] = sliceData[1399]	; tops[234][3] = sliceData[1400]	; tops[234][4] = sliceData[1401]	; tops[234][5] = sliceData[1402]	; tops[235][0] = sliceData[1403]	; tops[235][1] = sliceData[1404]	; tops[235][2] = sliceData[1405]	; tops[235][3] = sliceData[1406]	; tops[235][4] = sliceData[1407]	; tops[235][5] = sliceData[1408]	; tops[236][0] = sliceData[1409]	; tops[236][1] = sliceData[1410]	; tops[236][2] = sliceData[1411]	; tops[236][3] = sliceData[1412]	; tops[236][4] = sliceData[1413]	; tops[236][5] = sliceData[1414]	; tops[237][0] = sliceData[1415]	; tops[237][1] = sliceData[1416]	; tops[237][2] = sliceData[1417]	; tops[237][3] = sliceData[1418]	; tops[237][4] = sliceData[1419]	; tops[237][5] = sliceData[1420]	; tops[238][0] = sliceData[1421]	; tops[238][1] = sliceData[1422]	; tops[238][2] = sliceData[1423]	; tops[238][3] = sliceData[1424]	; tops[238][4] = sliceData[1425]	; tops[238][5] = sliceData[1426]	; tops[239][0] = sliceData[1427]	; tops[239][1] = sliceData[1428]	; tops[239][2] = sliceData[1429]	; tops[239][3] = sliceData[1430]	; tops[239][4] = sliceData[1431]	; tops[239][5] = sliceData[1432]	; tops[240][0] = sliceData[1433]	; tops[240][1] = sliceData[1434]	; tops[240][2] = sliceData[1435]	; tops[240][3] = sliceData[1436]	; tops[240][4] = sliceData[1437]	; tops[240][5] = sliceData[1438]	; tops[241][0] = sliceData[1439]	; tops[241][1] = sliceData[1440]	; tops[241][2] = sliceData[1441]	; tops[241][3] = sliceData[1442]	; tops[241][4] = sliceData[1443]	; tops[241][5] = sliceData[1444]	; tops[242][0] = sliceData[1445]	; tops[242][1] = sliceData[1446]	; tops[242][2] = sliceData[1447]	; tops[242][3] = sliceData[1448]	; tops[242][4] = sliceData[1449]	; tops[242][5] = sliceData[1450]	; tops[243][0] = sliceData[1451]	; tops[243][1] = sliceData[1452]	; tops[243][2] = sliceData[1453]	; tops[243][3] = sliceData[1454]	; tops[243][4] = sliceData[1455]	; tops[243][5] = sliceData[1456]	; tops[244][0] = sliceData[1457]	; tops[244][1] = sliceData[1458]	; tops[244][2] = sliceData[1459]	; tops[244][3] = sliceData[1460]	; tops[244][4] = sliceData[1461]	; tops[244][5] = sliceData[1462]	; tops[245][0] = sliceData[1463]	; tops[245][1] = sliceData[1464]	; tops[245][2] = sliceData[1465]	; tops[245][3] = sliceData[1466]	; tops[245][4] = sliceData[1467]	; tops[245][5] = sliceData[1468]	; tops[246][0] = sliceData[1469]	; tops[246][1] = sliceData[1470]	; tops[246][2] = sliceData[1471]	; tops[246][3] = sliceData[1472]	; tops[246][4] = sliceData[1473]	; tops[246][5] = sliceData[1474]	; tops[247][0] = sliceData[1475]	; tops[247][1] = sliceData[1476]	; tops[247][2] = sliceData[1477]	; tops[247][3] = sliceData[1478]	; tops[247][4] = sliceData[1479]	; tops[247][5] = sliceData[1480]	; tops[248][0] = sliceData[1481]	; tops[248][1] = sliceData[1482]	; tops[248][2] = sliceData[1483]	; tops[248][3] = sliceData[1484]	; tops[248][4] = sliceData[1485]	; tops[248][5] = sliceData[1486]	; tops[249][0] = sliceData[1487]	; tops[249][1] = sliceData[1488]	; tops[249][2] = sliceData[1489]	; tops[249][3] = sliceData[1490]	; tops[249][4] = sliceData[1491]	; tops[249][5] = sliceData[1492]	; tops[250][0] = sliceData[1493]	; tops[250][1] = sliceData[1494]	; tops[250][2] = sliceData[1495]	; tops[250][3] = sliceData[1496]	; tops[250][4] = sliceData[1497]	; tops[250][5] = sliceData[1498]	; tops[251][0] = sliceData[1499]	; tops[251][1] = sliceData[1500]	; tops[251][2] = sliceData[1501]	; tops[251][3] = sliceData[1502]	; tops[251][4] = sliceData[1503]	; tops[251][5] = sliceData[1504]	; tops[252][0] = sliceData[1505]	; tops[252][1] = sliceData[1506]	; tops[252][2] = sliceData[1507]	; tops[252][3] = sliceData[1508]	; tops[252][4] = sliceData[1509]	; tops[252][5] = sliceData[1510]	; tops[253][0] = sliceData[1511]	; tops[253][1] = sliceData[1512]	; tops[253][2] = sliceData[1513]	; tops[253][3] = sliceData[1514]	; tops[253][4] = sliceData[1515]	; tops[253][5] = sliceData[1516]	; tops[254][0] = sliceData[1517]	; tops[254][1] = sliceData[1518]	; tops[254][2] = sliceData[1519]	; tops[254][3] = sliceData[1520]	; tops[254][4] = sliceData[1521]	; tops[254][5] = sliceData[1522]	; tops[255][0] = sliceData[1523]	; tops[255][1] = sliceData[1524]	; tops[255][2] = sliceData[1525]	; tops[255][3] = sliceData[1526]	; tops[255][4] = sliceData[1527]	; tops[255][5] = sliceData[1528]	; tops[256][0] = sliceData[1529]	; tops[256][1] = sliceData[1530]	; tops[256][2] = sliceData[1531]	; tops[256][3] = sliceData[1532]	; tops[256][4] = sliceData[1533]	; tops[256][5] = sliceData[1534]	; tops[257][0] = sliceData[1535]	; tops[257][1] = sliceData[1536]	; tops[257][2] = sliceData[1537]	; tops[257][3] = sliceData[1538]	; tops[257][4] = sliceData[1539]	; tops[257][5] = sliceData[1540]	; tops[258][0] = sliceData[1541]	; tops[258][1] = sliceData[1542]	; tops[258][2] = sliceData[1543]	; tops[258][3] = sliceData[1544]	; tops[258][4] = sliceData[1545]	; tops[258][5] = sliceData[1546]	; tops[259][0] = sliceData[1547]	; tops[259][1] = sliceData[1548]	; tops[259][2] = sliceData[1549]	; tops[259][3] = sliceData[1550]	; tops[259][4] = sliceData[1551]	; tops[259][5] = sliceData[1552]	; tops[260][0] = sliceData[1553]	; tops[260][1] = sliceData[1554]	; tops[260][2] = sliceData[1555]	; tops[260][3] = sliceData[1556]	; tops[260][4] = sliceData[1557]	; tops[260][5] = sliceData[1558]	; tops[261][0] = sliceData[1559]	; tops[261][1] = sliceData[1560]	; tops[261][2] = sliceData[1561]	; tops[261][3] = sliceData[1562]	; tops[261][4] = sliceData[1563]	; tops[261][5] = sliceData[1564]	; tops[262][0] = sliceData[1565]	; tops[262][1] = sliceData[1566]	; tops[262][2] = sliceData[1567]	; tops[262][3] = sliceData[1568]	; tops[262][4] = sliceData[1569]	; tops[262][5] = sliceData[1570]	; tops[263][0] = sliceData[1571]	; tops[263][1] = sliceData[1572]	; tops[263][2] = sliceData[1573]	; tops[263][3] = sliceData[1574]	; tops[263][4] = sliceData[1575]	; tops[263][5] = sliceData[1576]	; tops[264][0] = sliceData[1577]	; tops[264][1] = sliceData[1578]	; tops[264][2] = sliceData[1579]	; tops[264][3] = sliceData[1580]	; tops[264][4] = sliceData[1581]	; tops[264][5] = sliceData[1582]	; tops[265][0] = sliceData[1583]	; tops[265][1] = sliceData[1584]	; tops[265][2] = sliceData[1585]	; tops[265][3] = sliceData[1586]	; tops[265][4] = sliceData[1587]	; tops[265][5] = sliceData[1588]	; tops[266][0] = sliceData[1589]	; tops[266][1] = sliceData[1590]	; tops[266][2] = sliceData[1591]	; tops[266][3] = sliceData[1592]	; tops[266][4] = sliceData[1593]	; tops[266][5] = sliceData[1594]	; tops[267][0] = sliceData[1595]	; tops[267][1] = sliceData[1596]	; tops[267][2] = sliceData[1597]	; tops[267][3] = sliceData[1598]	; tops[267][4] = sliceData[1599]	; tops[267][5] = sliceData[1600]	; tops[268][0] = sliceData[1601]	; tops[268][1] = sliceData[1602]	; tops[268][2] = sliceData[1603]	; tops[268][3] = sliceData[1604]	; tops[268][4] = sliceData[1605]	; tops[268][5] = sliceData[1606]	; tops[269][0] = sliceData[1607]	; tops[269][1] = sliceData[1608]	; tops[269][2] = sliceData[1609]	; tops[269][3] = sliceData[1610]	; tops[269][4] = sliceData[1611]	; tops[269][5] = sliceData[1612]	; tops[270][0] = sliceData[1613]	; tops[270][1] = sliceData[1614]	; tops[270][2] = sliceData[1615]	; tops[270][3] = sliceData[1616]	; tops[270][4] = sliceData[1617]	; tops[270][5] = sliceData[1618]	; tops[271][0] = sliceData[1619]	; tops[271][1] = sliceData[1620]	; tops[271][2] = sliceData[1621]	; tops[271][3] = sliceData[1622]	; tops[271][4] = sliceData[1623]	; tops[271][5] = sliceData[1624]	; tops[272][0] = sliceData[1625]	; tops[272][1] = sliceData[1626]	; tops[272][2] = sliceData[1627]	; tops[272][3] = sliceData[1628]	; tops[272][4] = sliceData[1629]	; tops[272][5] = sliceData[1630]	; tops[273][0] = sliceData[1631]	; tops[273][1] = sliceData[1632]	; tops[273][2] = sliceData[1633]	; tops[273][3] = sliceData[1634]	; tops[273][4] = sliceData[1635]	; tops[273][5] = sliceData[1636]	; tops[274][0] = sliceData[1637]	; tops[274][1] = sliceData[1638]	; tops[274][2] = sliceData[1639]	; tops[274][3] = sliceData[1640]	; tops[274][4] = sliceData[1641]	; tops[274][5] = sliceData[1642]	; tops[275][0] = sliceData[1643]	; tops[275][1] = sliceData[1644]	; tops[275][2] = sliceData[1645]	; tops[275][3] = sliceData[1646]	; tops[275][4] = sliceData[1647]	; tops[275][5] = sliceData[1648]	; tops[276][0] = sliceData[1649]	; tops[276][1] = sliceData[1650]	; tops[276][2] = sliceData[1651]	; tops[276][3] = sliceData[1652]	; tops[276][4] = sliceData[1653]	; tops[276][5] = sliceData[1654]	; tops[277][0] = sliceData[1655]	; tops[277][1] = sliceData[1656]	; tops[277][2] = sliceData[1657]	; tops[277][3] = sliceData[1658]	; tops[277][4] = sliceData[1659]	; tops[277][5] = sliceData[1660]	; tops[278][0] = sliceData[1661]	; tops[278][1] = sliceData[1662]	; tops[278][2] = sliceData[1663]	; tops[278][3] = sliceData[1664]	; tops[278][4] = sliceData[1665]	; tops[278][5] = sliceData[1666]	; tops[279][0] = sliceData[1667]	; tops[279][1] = sliceData[1668]	; tops[279][2] = sliceData[1669]	; tops[279][3] = sliceData[1670]	; tops[279][4] = sliceData[1671]	; tops[279][5] = sliceData[1672]	; tops[280][0] = sliceData[1673]	; tops[280][1] = sliceData[1674]	; tops[280][2] = sliceData[1675]	; tops[280][3] = sliceData[1676]	; tops[280][4] = sliceData[1677]	; tops[280][5] = sliceData[1678]	; tops[281][0] = sliceData[1679]	; tops[281][1] = sliceData[1680]	; tops[281][2] = sliceData[1681]	; tops[281][3] = sliceData[1682]	; tops[281][4] = sliceData[1683]	; tops[281][5] = sliceData[1684]	; tops[282][0] = sliceData[1685]	; tops[282][1] = sliceData[1686]	; tops[282][2] = sliceData[1687]	; tops[282][3] = sliceData[1688]	; tops[282][4] = sliceData[1689]	; tops[282][5] = sliceData[1690]	; tops[283][0] = sliceData[1691]	; tops[283][1] = sliceData[1692]	; tops[283][2] = sliceData[1693]	; tops[283][3] = sliceData[1694]	; tops[283][4] = sliceData[1695]	; tops[283][5] = sliceData[1696]	; tops[284][0] = sliceData[1697]	; tops[284][1] = sliceData[1698]	; tops[284][2] = sliceData[1699]	; tops[284][3] = sliceData[1700]	; tops[284][4] = sliceData[1701]	; tops[284][5] = sliceData[1702]	; tops[285][0] = sliceData[1703]	; tops[285][1] = sliceData[1704]	; tops[285][2] = sliceData[1705]	; tops[285][3] = sliceData[1706]	; tops[285][4] = sliceData[1707]	; tops[285][5] = sliceData[1708]	; tops[286][0] = sliceData[1709]	; tops[286][1] = sliceData[1710]	; tops[286][2] = sliceData[1711]	; tops[286][3] = sliceData[1712]	; tops[286][4] = sliceData[1713]	; tops[286][5] = sliceData[1714]	; tops[287][0] = sliceData[1715]	; tops[287][1] = sliceData[1716]	; tops[287][2] = sliceData[1717]	; tops[287][3] = sliceData[1718]	; tops[287][4] = sliceData[1719]	; tops[287][5] = sliceData[1720]	; tops[288][0] = sliceData[1721]	; tops[288][1] = sliceData[1722]	; tops[288][2] = sliceData[1723]	; tops[288][3] = sliceData[1724]	; tops[288][4] = sliceData[1725]	; tops[288][5] = sliceData[1726]	; tops[289][0] = sliceData[1727]	; tops[289][1] = sliceData[1728]	; tops[289][2] = sliceData[1729]	; tops[289][3] = sliceData[1730]	; tops[289][4] = sliceData[1731]	; tops[289][5] = sliceData[1732]	; tops[290][0] = sliceData[1733]	; tops[290][1] = sliceData[1734]	; tops[290][2] = sliceData[1735]	; tops[290][3] = sliceData[1736]	; tops[290][4] = sliceData[1737]	; tops[290][5] = sliceData[1738]	; tops[291][0] = sliceData[1739]	; tops[291][1] = sliceData[1740]	; tops[291][2] = sliceData[1741]	; tops[291][3] = sliceData[1742]	; tops[291][4] = sliceData[1743]	; tops[291][5] = sliceData[1744]	; tops[292][0] = sliceData[1745]	; tops[292][1] = sliceData[1746]	; tops[292][2] = sliceData[1747]	; tops[292][3] = sliceData[1748]	; tops[292][4] = sliceData[1749]	; tops[292][5] = sliceData[1750]	; tops[293][0] = sliceData[1751]	; tops[293][1] = sliceData[1752]	; tops[293][2] = sliceData[1753]	; tops[293][3] = sliceData[1754]	; tops[293][4] = sliceData[1755]	; tops[293][5] = sliceData[1756]	; tops[294][0] = sliceData[1757]	; tops[294][1] = sliceData[1758]	; tops[294][2] = sliceData[1759]	; tops[294][3] = sliceData[1760]	; tops[294][4] = sliceData[1761]	; tops[294][5] = sliceData[1762]	; tops[295][0] = sliceData[1763]	; tops[295][1] = sliceData[1764]	; tops[295][2] = sliceData[1765]	; tops[295][3] = sliceData[1766]	; tops[295][4] = sliceData[1767]	; tops[295][5] = sliceData[1768]	; tops[296][0] = sliceData[1769]	; tops[296][1] = sliceData[1770]	; tops[296][2] = sliceData[1771]	; tops[296][3] = sliceData[1772]	; tops[296][4] = sliceData[1773]	; tops[296][5] = sliceData[1774]	; tops[297][0] = sliceData[1775]	; tops[297][1] = sliceData[1776]	; tops[297][2] = sliceData[1777]	; tops[297][3] = sliceData[1778]	; tops[297][4] = sliceData[1779]	; tops[297][5] = sliceData[1780]	; tops[298][0] = sliceData[1781]	; tops[298][1] = sliceData[1782]	; tops[298][2] = sliceData[1783]	; tops[298][3] = sliceData[1784]	; tops[298][4] = sliceData[1785]	; tops[298][5] = sliceData[1786]	; tops[299][0] = sliceData[1787]	; tops[299][1] = sliceData[1788]	; tops[299][2] = sliceData[1789]	; tops[299][3] = sliceData[1790]	; tops[299][4] = sliceData[1791]	; tops[299][5] = sliceData[1792]	; tops[300][0] = sliceData[1793]	; tops[300][1] = sliceData[1794]	; tops[300][2] = sliceData[1795]	; tops[300][3] = sliceData[1796]	; tops[300][4] = sliceData[1797]	; tops[300][5] = sliceData[1798]	;
	
	/* Create a new Discord session using the provided bot token. */
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	/* Register the messageCreate func as a callback for MessageCreate events. */
	dg.AddHandler(messageCreate)

	/* In this example, we only care about receiving message events. */
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	/* Open a websocket connection to Discord and begin listening. */
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	/* Wait here until CTRL-C or other term signal is received. */
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	/* Cleanly close down the Discord session. */
	dg.Close()
}

/* This function will be called (due to AddHandler above) every time a new  */
/* message is created on any channel that the authenticated bot has access to. */
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	/* Ignore all messages created by the bot itself */
	/* This isn't required in this specific example but it's a good practice. */
	if m.Author.ID == s.State.User.ID {
		return
	}

	if is_illegal(m.Content) {
		s.ChannelMessageSend("1031077892748234762", "<@" + m.Author.ID + "> ha hecho trampas\t" + time.Now().Format("01-02-2006 15:04:05") + "\t" + split_curr())
	}
	if m.Content == ".tt" {
		rand.Seed(time.Now().UnixNano())

		
		random = rand.Intn(how_many_texts())
		random = 0

		start_author = m.Author.ID
		s.ChannelMessageSend(m.ChannelID, "Preparados...") 
		time.Sleep(2 * time.Second)
/*
		chanl, err := s.Channel(m.ChannelID)
		if err != nil {
			return
		}

		s.ChannelMessageEdit(m.ChannelID, m.Message.ID, display_textos [random])
		fmt.Println(m.Message.ID)
*/
		s.ChannelMessageSend(m.ChannelID, display_textos [random])
		current_text = textos[random]
		start()
		
	}

	if m.Content == current_text && is_started {
		calculate_wpm()
		/*is_started = false*/

		wpm_stringed := fmt.Sprint(wpm)

		// CALCULATE IF GREATER OR SMALER
		
		var AAA float64

		/* IF TOP 5 */
		/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+4][2], 64); err == nil {AAA = AA}
		if wpm > AAA {
			if tops[random*5+3][5] == m.Author.ID {
				var temp string
				var temp2 string

				temp = tops[random*5+4][2]
				temp2 = tops[random*5+4][3]

				wpm_string_temp := fmt.Sprintf("%f", wpm)  
				tops[random*5+4][2] = wpm_string_temp
				tops[random*5+4][1] = m.Author.Username
				tops[random*5+4][5] = m.Author.ID
				dt := time.Now()
				tops[random*5+4][3] = dt.Format("01-02-2006 15:04:05")
				tops[random*5+4][0] = "5"

				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
			}
		}

		/* IF TOP 4 */
		/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+3][2], 64); err == nil {AAA = AA}
		if wpm > AAA {
			if tops[random*5+3][5] == m.Author.ID {
				var temp string
				var temp2 string

				temp = tops[random*5+3][2]
				temp2 = tops[random*5+3][3]

				wpm_string_temp := fmt.Sprintf("%f", wpm)  
				tops[random*5+3][2] = wpm_string_temp
				tops[random*5+3][1] = m.Author.Username
				tops[random*5+3][5] = m.Author.ID
				dt := time.Now()
				tops[random*5+3][3] = dt.Format("01-02-2006 15:04:05")
				tops[random*5+3][0] = "4"

				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
			} else if tops[random*5+4][5] == m.Author.ID {
				var temp string
				var temp2 string

				temp = tops[random*5+4][2]
				temp2 = tops[random*5+4][3]

				tops[random*5+4][2] = tops[random*5+3][2]
				tops[random*5+4][1] = tops[random*5+3][1]
				tops[random*5+4][5] = tops[random*5+3][5]
				tops[random*5+4][3] = tops[random*5+3][3]
				tops[random*5+4][0] = "5"

				wpm_string_temp := fmt.Sprintf("%f", wpm)  
				tops[random*5+3][2] = wpm_string_temp
				tops[random*5+3][1] = m.Author.Username
				tops[random*5+3][5] = m.Author.ID
				dt := time.Now()
				tops[random*5+3][3] = dt.Format("01-02-2006 15:04:05")
				tops[random*5+3][0] = "4"

				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
			}
		}

		/* IF TOP 3 */
		/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+2][2], 64); err == nil {AAA = AA}
		if wpm > AAA {
			if tops[random*5+2][5] == m.Author.ID {
				var temp string
				var temp2 string

				temp = tops[random*5+2][2]
				temp2 = tops[random*5+2][3]

				wpm_string_temp := fmt.Sprintf("%f", wpm)  
				tops[random*5+2][2] = wpm_string_temp
				tops[random*5+2][1] = m.Author.Username
				tops[random*5+2][5] = m.Author.ID
				dt := time.Now()
				tops[random*5+2][3] = dt.Format("01-02-2006 15:04:05")
				tops[random*5+2][0] = "3"

				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
			}
		}

		/* IF TOP 2 */
		/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+1][2], 64); err == nil {AAA = AA}
		if wpm > AAA {
			if tops[random*5+1][5] == m.Author.ID {
				var temp string
				var temp2 string

				temp = tops[random*5+1][2]
				temp2 = tops[random*5+1][3]

				wpm_string_temp := fmt.Sprintf("%f", wpm)  
				tops[random*5+1][2] = wpm_string_temp
				tops[random*5+1][1] = m.Author.Username
				tops[random*5+1][5] = m.Author.ID
				dt := time.Now()
				tops[random*5+1][3] = dt.Format("01-02-2006 15:04:05")
				tops[random*5+1][0] = "2"

				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
			}
		}

		/* IF TOP 1 */
		/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5][2], 64); err == nil {AAA = AA}
		if wpm > AAA {
			if tops[random*5][5] == m.Author.ID {
				var temp string
				var temp2 string

				temp = tops[random*5][2]
				temp2 = tops[random*5][3]

				wpm_string_temp := fmt.Sprintf("%f", wpm)  
				tops[random*5][2] = wpm_string_temp
				tops[random*5][1] = m.Author.Username
				tops[random*5][5] = m.Author.ID
				dt := time.Now()
				tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
				tops[random*5][0] = "1"

				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
			}
		}

		s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has terminado la carrera!\nWPM: " + wpm_stringed + 
		"\nTops:\n" + tops[random*5][0] + ". " + tops[random*5][1] + " (" + tops[random*5][2] + " wpm) " + tops[random*5][3] +
		"\n" + tops[random*5+1][0] + ". " + tops[random*5+1][1] + " (" + tops[random*5+1][2] + " wpm) " + tops[random*5+1][3] +
		"\n" + tops[random*5+2][0] + ". " + tops[random*5+2][1] + " (" + tops[random*5+2][2] + " wpm) " + tops[random*5+2][3] +
		"\n" + tops[random*5+3][0] + ". " + tops[random*5+3][1] + " (" + tops[random*5+3][2] + " wpm) " + tops[random*5+4][3] +
		"\n" + tops[random*5+4][0] + ". " + tops[random*5+4][1] + " (" + tops[random*5+4][2] + " wpm) " + tops[random*5+4][3] +

		"\n--------------------------------------------------------------------")

		var random_s = strconv.FormatInt(int64(random)+1, 10)

		f, err := os.OpenFile("database/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		dt := time.Now()
		if _, err := f.Write([]byte(m.Author.ID + "\t" + dt.Format("01-02-2006 15:04:05") + "\t" + wpm_stringed + "\ttexto: (" + random_s + ") " + split_curr() + "\n")); err != nil {
			f.Close() 
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

		/* WRITE TO FILE */
	    if err != nil {
	        log.Fatalf("readLines: %s", err)
	    }

	    escribir_tops_a_base_de_datos()


		if err := writeLines(); err != nil {
       		 log.Fatalf("writeLines: %s", err)
	  	}


 	   /* WRITE TO FILE */


	} else if CountWords(m.Content) > CountWords(current_text)-3 {
		if is_illegal(m.Content) {
			for i := 0; i < 4; i++ {
				fmt.Println("¡¡<@" + m.Author.ID + "> HA COPY PASTE!!")
			}
			

			A := m.Content
			sent_arrayed := strings.Split(A, " ")

			B := current_text
			text_arrayed := strings.Split(B, " ")

			if (sent_arrayed[0] == text_arrayed[0]) /*&& (sent_arrayed[1] == text_arrayed[1] && (sent_arrayed[len(sent_arrayed)-1] == text_arrayed[len(text_arrayed)-1])*/ {
				calculate_wpm()
				calculate_errors(m.Content, current_text)
				/*is_started = false*/
				wpm_stringed := fmt.Sprint(wpm)
				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\nHAS INTENTADO COPY PASTE. (TU WPM HUBIERA SIDO: " + wpm_stringed + ")\n--------------------------------------------------------------------")

				s.ChannelMessageSend("1031077892748234762", errores_s + ": " + lista_errores)
			}
		} else {

			A := m.Content
			sent_arrayed := strings.Split(A, " ")

			B := current_text
			text_arrayed := strings.Split(B, " ")

			if (sent_arrayed[0] == text_arrayed[0]) /*&& (sent_arrayed[1] == text_arrayed[1] && (sent_arrayed[len(sent_arrayed)-1] == text_arrayed[len(text_arrayed)-1])*/ {
				calculate_wpm()
				calculate_errors(m.Content, current_text)
				/*is_started = false*/
				wpm_stringed := fmt.Sprint(wpm)
				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\nNos has terminado la carrera correctamente.\nHas cometido " + errores_s + " errores: " + lista_errores + "\nWPM raw: " + wpm_stringed + "\n--------------------------------------------------------------------")
			}
		}
	}

	if m.Content == ".info" {
		sec := fmt.Sprint(time_elapsed)
		sec2 := fmt.Sprint(time_elapsed/1000)

		var average_word_length_of_current_text_stringed string = fmt.Sprintf("%f", (average_word_length(current_text)))
	
	    dat, err := os.ReadFile("database/test")
    	check(err)
    	s.ChannelMessageSend(m.ChannelID, string(dat) + "\n[" + m.Author.ID + "]\nmilliseconds: " + sec + "\nseconds: " + sec2 + "\nstart_author: " + start_author + "\naverage_word_length_of_current_text_stringed: " + average_word_length_of_current_text_stringed)
	}

	if m.Content == ".test" {
		f, err := os.OpenFile("database/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		dt := time.Now()
		if _, err := f.Write([]byte(m.Author.ID + "\t" + dt.Format("02-21-2006 15:04:05") + "\tappended some data\n")); err != nil {
			f.Close() 
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	if m.Content == ".A" {
		ff := string(tops[random*5][2])
		fmt.Println(ff)
		fmt.Println(tops[random*5][2])
		f, err := strconv.ParseFloat(ff, 8)

		
		fmt.Println(f, err, reflect.TypeOf(f))


	}

	if m.Content == ".help" {
		s.ChannelMessageSend(m.ChannelID, "```.tt       empieza un test de velocidad en idioma español\n.info     enseña información aleatoria para desarrolladores\n.mapache  pone el gif de un mapache```")
	}

	/* FUN COMMANDS */
	/**/	
	/**/	if m.Content == ".mapache" {
	/**/		s.ChannelMessageSend(m.ChannelID, "https://tenor.com/view/froze-stop-moving-not-moving-still-standing-gif-16669739")
	/**/	}
	/**/
	/**/	if m.Content == ".logo" {
	/**/		s.ChannelMessageSend(m.ChannelID, "𝗧𝗬𝗣𝗘 𝗕𝗢𝗧")
	/**/	}	
	/**/
	/**/	if m.Content == ".go" {
	/**/		s.ChannelMessageSend(m.ChannelID, "https://camo.githubusercontent.com/833cfd306ac2bef74ddf0560ee3b4112321c5b6939e52a1629f0aed8aec46922/687474703a2f2f692e696d6775722e636f6d2f485379686177742e6a7067")
	/**/	}
	/**/
	/**/	if m.Content == ".ch" {
	/**/		s.ChannelMessageSend(m.ChannelID, "https://thumbs.gfycat.com/AdmirableLikableFlea-mobile.mp4")
	/**/	}
	/**/
	/**/
	/**/
	/**/
	/**/
	/**/
	/**/
	/**/
			

}

