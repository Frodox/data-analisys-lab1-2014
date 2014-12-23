## Лабораторная работа №1 по Анализу данных


*Дата*: осень 2014 (9й семестр)  
*Тема*: одноступенчатый приёмочный контроль по альтернативному признаку


### Что это ?

В этом проекте находится программный код (использован язык `Go`gs),
позволяющий выполнить задание лабораторной для
произвольных валидных входных данных.


### Задание

Значения объёма контролируемой партии (`N`), числа дефектных изделий поставщика (`Dп`),
числа дефектных изделий заказчика (`Dз`), минимальный (`n_min`)
и максимальный объем выборки (`n_max`), шаг изменения объёма выборки (`step`),
вероятность ошибки первого рода (`α`) приведены в таблице
(таблица вариантов здесь не приведена).

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


### Предварительная подготовка

1. [Установить интерпретатор языка Go](https://golang.org/doc/install)
2. [Настроить окружение](https://golang.org/doc/code.html) для проектов на `Go`
3. Скопировать текущий проект  
`git clone https://github.com/Frodox/data-analysis-lab1-2014.git $GOPATH/src/github.com/Frodox/data-analisys-lab1`
4. Установить программу `gnuplot`


### Как запустить расчёты ?

4. Создать файл `data/lab1.data.varX.txt` на основе файла
`data/lab1.data.var6.txt` со своими входными данными
5. `cd $GOPATH/src/github.com/Frodox/data-analisys-lab1`
6. Сборка проекта: `go build -o lab1 .`
7. Запуск расчётов: `./lab1 --filename data/lab1.data.varX.txt`
8. Запуск расчётов с форматированным выводом:  
`./lab1 --filename data/lab1.data.varX.txt --plotable`

По-умолчанию берутся данные для *6*го варианта.


### Как построить графики ?

1. Подготовка данных для создания графиков:  
`./lab1 --filename data/lab1.data.varX.txt --plotable >data/outputdata.varX.txt`
2. Рисуем графики:  
`gnuplot -e "inputdatafile='data/outputdata.varX.txt'" drawgraphics.gnu`
3. Смотрим файлы `betta_from_n.pdf` и `c_from_n.pdf`
