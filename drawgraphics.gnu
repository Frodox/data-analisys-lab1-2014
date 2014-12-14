set style line 1 lc rgb '#0060ad' lt 1 lw 2 pt 7 ps 1.5   # --- blue
set style line 2 lc rgb '#dd181f' lt 1 lw 2 pt 5 ps 1.5   # --- red
set style line 3 lc rgb '#09782b' lt 1 lw 2 pt 3 ps 1.5   # --- red

set grid xtics lt 0 lw 1 lc rgb "#bbbbbb"
set grid ytics lt 0 lw 1 lc rgb "#bbbbbb"


set title "График зависимости приёмочного числа от объёма выборки"
set xlabel 'Объём контролируемой выборки (n)'
set ylabel 'Приёмочное число (с)'

# set term "postscript" eps
# set output "c_n.eps"

plot 'data.txt'    u 1:2 with linespoints ls 1 title 'c1 (P1)', \
		'' u 1:4 with linespoints ls 2 title 'c2 (P2)', \
		'' u 1:6 with linespoints ls 3 title 'c3 (P3)'

pause -1


set title "График зависимости вероятности ошибки второго рода от объёма выборки"
set xlabel 'Объём контролируемой выборки (n)'
set ylabel 'Вероятность ошибки второго рода (betta)'

plot 'data.txt'    u 1:3 with linespoints ls 1 title 'betta1 (P1)', \
		'' u 1:5 with linespoints ls 2 title 'betta2 (P2)', \
		'' u 1:7 with linespoints ls 3 title 'betta3 (P3)'

pause -1
