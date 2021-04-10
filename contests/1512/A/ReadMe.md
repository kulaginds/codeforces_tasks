# [A. Шпион обнаружен!](https://codeforces.com/contest/1512/problem/A)
> ограничение по времени на тест: 2 секунды  
> ограничение по памяти на тест: 256 мегабайт  
> ввод: стандартный ввод  
> вывод: стандартный вывод  

Вам дан массив `a`, состоящий из `n` (`n>=3`) целых положительных чисел.
Известно, что в этом массиве, все числа, кроме одного, совпадают (например, в массиве `[4,11,4,4]` все числа, кроме одного, равны `4`).

Выведите номер элемента, который не совпадает с остальными. Числа в массиве пронумерованы с единицы.

## Входные данные
В первой строке содержится одно целое число `t` (`1<=t<=100`). Далее следуют `t` наборов входных данных.

Первая строка каждого набора входных данных содержит одно целое число `n` (3<=n<=100) — длина массива `a`.

Вторая строка каждого набора входных данных содержит `n` целых чисел `a1, a2, ..., an` (`1<=ai<=100`).

## Выходные данные
Для каждого набора входных данных выведите одно целое число — номер элемента, который не совпадает с остальными.

### Пример
входные данные:
```text
4
4
11 13 11 11
5
1 4 4 4 4
10
3 3 3 3 10 3 3 3 3 3
3
20 20 10
```

выходные данные:
```text
2
1
5
3
```