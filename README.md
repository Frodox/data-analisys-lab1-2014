Лабораторная работа №1 по Анализу данных
=======================

*Дата*: осень 2014 (9й семестр)  
*Тема*: одноступенчатый приёмочный контроль по альтернативному признаку


### Что это ?

В этом проекте находится программный код (использован язык `Go`), позволяющий выполнить задание лабораторной
для произвольных валидных входных данных.


### Задание

Значения объёма контролируемой партии (`N`), числа дефектных изделий поставщика (`Dп`),
числа дефектных изделий заказчика (`Dз`), минимальный (`n_min`)
и максимальный объем выборки (`n_max`), шаг изменения объёма выборки (`step`),
вероятность ошибки первого рода (`α`) приведены в таблице.

Необходимо для каждого объёма выборки между `n_min` и `n_max` с шагом `step`
выполнить расчёт приёмочного числа (`c`) и вероятности ошибки второго рода (`β`).
Выполнить расчёт с использованием каждой из формул для закона распределения
числа дефектных изделий.

* P1 -- точная формула. `P (x = l) = C (D_f, l) * C (N-D_f, n-l)   /   C (N, n)`
* P2 -- приближённая формула (биноминальная). `P(x=l) = C(n, l) * po^l * (1 - po)^(n-l); po = D_f / N`
* P3 -- приближёнаня формула. `P(x=l) = lambda^l / (l!)  * exp(-lambda)`

Построить графики зависимости приёмочного числа и вероятности ошибки второго
рода от объёма выборки. На графиках отобразить тремя ломаными значения для каждой
из применённых формул. Сделать вывод о применимости приближенных формул
для закона распределения числа дефектных изделий в данных условиях.


### Как запустить расчёты ?

### Как построить график ?

`$ gnuplot -e inputdatafile='outputdata.var6.txt' drawgraphics.gnu`
