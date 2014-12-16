# -*- coding: koi8-r -*-

set encoding koi8r
set term postscript eps enhanced "Arial,11"

set style line 1 lc rgb '#0060ad' lt 1 lw 2 pt 7 ps 1.5   # --- blue
set style line 2 lc rgb '#dd181f' lt 1 lw 2 pt 5 ps 1.5   # --- red
set style line 3 lc rgb '#09782b' lt 1 lw 2 pt 3 ps 1.5   # --- red

set grid xtics lt 0 lw 1 lc rgb "#bbbbbb"
set grid ytics lt 0 lw 1 lc rgb "#bbbbbb"

### determine input data filename #
if (!exists("inputdatafile")) inputdatafile='outputdata.var6.txt'


# ПЕРВЫЙ ГРАФИК ###############################################################
set title "График зависимости приёмочного числа от объёма выборки"
set xlabel 'Объём контролируемой выборки (n)'
set ylabel 'Приёмочное число (с)'

set output '| ps2pdf -g6000x4300 -dOptimize=true -dEmbedAllFonts=true -dEPSFitPage - c_from_n.pdf'

plot inputdatafile u 1:2 with linespoints ls 1 title 'c1 (P1)', \
		'' u 1:4 with linespoints ls 2 title 'c2 (P2)', \
		'' u 1:6 with linespoints ls 3 title 'c3 (P3)'


# ВТОРОЙ ГРАФИК ###############################################################

set title "График зависимости вероятности ошибки второго рода от объёма выборки"
set xlabel 'Объём контролируемой выборки (n)'
set ylabel 'Вероятность ошибки второго рода (betta)'

set output '| ps2pdf -g6000x4300 -dOptimize=true -dEmbedAllFonts=true -dEPSFitPage - betta_from_n.pdf'

plot inputdatafile u 1:3 with linespoints ls 1 title 'betta1 (P1)', \
		'' u 1:5 with linespoints ls 2 title 'betta2 (P2)', \
		'' u 1:7 with linespoints ls 3 title 'betta3 (P3)'

