:pensive: Ну, нарешивали нарешивали, а толку то? Вообщем по задумке тут бы классно все струтурировать, как получится хз.  

![FaceOfRepa](/img/me.jpg)  

## List

*Темы без ссылок не нуждаются в просмотре*

* Ya 2.0
    * Линейный поиск
    * Множества
    * [Сортировка подсчетом](#сортировка-подсчетом)
    * Словари
    * [Префиксные суммы](#префиксные-суммы)
    * [Два указателя](#два-указателя)
    * [Бинарный поиск](#бинарный-поиск)
    * [Сортировка событий](#сортировка-событий) *to-do*
    * Деревья
* Ya 3.0
    * [Стеки](#стеки)
    * [Очереди(реализация)](#очереди-и-диквэй)
    * [Куча](#куча)
    * [Одномерное ДП](#одномерное-дп)
    * [Двумерное ДП](#двумерное-дп)



#### Сортировка подсчетом
Пусть дан массив из N элементов таких, что их разброс невелик. Обычная сортировка займет $O(NLogN)$  

Однако, если разбор небольшой проще сосчитать кол-во каждого элемента, а затем восстановить массив  
![NumSort](/img/NumSort.png)

#### Префиксные суммы
Подсчет *чего-либо* на отрезке, быстрый ответ на запрос ~~суммы~~ *операции* на отрезке 

![PreffSum](/img/PreffSum.png)

#### Два указателя
Много задач, формата 2 массива, станцуйте с бубном и т.д. Обычно 2 указателя там очевидны

Далее - метод плавающего окна. Тригер к 2 указателям соблюдение инварианта.

Должно выполнять 2 свойвства

* если отрезок [L,R] хороший, то любой вложенный в него отрезок также хороший
    * если отрезок [L,R] хороший, то любой отрезок, который его содержит также хороший
* Во вторых, вы должны уметь пересчитывать вашу функцию (проверять хороший отрезок или плохой), при увеличении левой или правой границы на единицу вправо.

![2pointers1](/img/2pointers1.png)

![2pointers2](/img/2pointers2.png)

Код всегда максимально похожий  

```
L = 0
for R = 0..n-1
    add(a[R])
    while not good():
        remove(a[L])
        L++
```

Есть еще неплохая темка с очередью на 2 стеках, помогает в задачах, с сложным пересчетом функции. Это нужно, если при удалении/добавлении элемента в наше "окно", самого элемента недостаточно, и нужно учитывать предыдушие, но в таком случае требуется транзитивность, как например хранение максимума на окне, хранения НОД всех чисел и прочее

![2pointersStack](/img/2pointersStackk.png)

#### Бинарный поиск
Сразу не занимаемся херней с поиском на отрезке, там нужно думать, что бы правая и левая граница сошлась, а думать это для дебилов, д - думать, д - дебил  

Главный тригер к бинарке - монотонная функция 2 видов, поиск элемента в массиве, минимакс и т.д.

![BinSearchFunc](/img/BinSearch1.png)

Поиск по таким монотонным функциям места "перегиба" - бинпоиск по ответу

Пусть у нас монотонная, возрастающая функция, тогда при поиске сохраняем инвариант полуинтервала [L, R]:
* L - указывает на "левую" часть функции
* R - правая часть функции
```golang
L := 0 // Левая граница
// крутые математики, сразу могут определить при каком R функция будет true, но д - думать, д - дебил, делаем так
R := 1
for !func(R) {
    R *= 2
}

for R != L + 1 {
    MID := (L + R) / 2
    if good(MID) {
        R = MID
    } else {
        L = MID
    }
}

func good(X int, ...) bool {
    res := // проверяем, удовлетворению условию при X
    return res
}
```

Вещественный бинпоиск - все тоже самое, тот же инвариант, но, условие for-а теперь так просто не задашь, глобально есть 2 варианта
```golang
// 1 eps - точность
eps := 1e-6
for (R - L > eps)

// 2 через итерации
for (i := 0; i != 150; i++)
// 150 итераций - 2^150, вообщем ответ получим с максимальной точностью, у нас на итерации ~60-той, уже будут соседние числа в float64 ~~double~~
```

Минимакс - развал кабины  
Задачи двух видов: $min(max(x_i))$ или $max(min(x_i))$  

Вместо попыток в минимизацию, мы можем просто ограничить функцию $max(x_i)$ сверху, а далее бин поиском подбирать ограничение, это будет та же мемезация за $~O(LogN)$  

Затем мы получим $max(x_i) <= t; \forall x_i <= t$ Обычно это можно проверить за $~O(N)$ 

При $max(min(x_i))$ тоже самое в обратную сторону, ограничиваем снизу, затем решаем задачу $min(x_i) > t$

Максимальное/минимальное среднее - развал кабинx2  
Пускай у нас есть какое-то множество, нам нужно выбрать подмножество, обладающее какими-то свойствами, и в этом подмножестве нужно максимизировать/минимизировать среднее арифметическое элементов.

То бишь 
```math
\frac{\sum_{l}^r a_i}{r - l} \to max
```

в таком случае max/min заменяем на бинарку с ограничением снизу/сверху  

$$\frac{\sum_{l}^r a_i}{r - l} >= x$$  

$$\sum_{l}^r a_i >= x * (r - l)$$  

$$\sum_{l}^r (a_i - x) >= 0$$

подобное ищется префиксными суммами 
```math
  \begin{cases}
    p_r > p_l\\
    r - l >= D 
  \end{cases}
```
В данном случае достаточно хранить индекс минимальной префиксной суммы на отрезке [0, R-L]

**TO-DO** Поиск *K* элемента, разобрать

#### Сортировка событий

**TO-DO**

#### Стеки

Задача на правильную скобочную последовательность, признак стека - *вложенность*

Задача на преобразование в постфиксную запись, опять *некая вложенность*

![преобразование в постфикс](/img/StackPostfix.png)

Задача о поиске ближайшего меньшего справа 

![меньшийсправа](/img/StackLessRight.png)

Тут сложнее. Если нам прилетает число X, то для всех чисел в стеке больше X, чисто X будет ответом, более они не нуждаются в рассмотрении. Также стек будет упорядочен, ведь если $a_i > a_i-1$ в стеке, то $a_i$ вытолкнестся, ответ для него уже найден

#### Очереди, и диквэй

Реализация очереди - как в C++, но, также можно через кольцевой буфер

![dequeLoop](/img/LoopBuff.png)

#### Куча

Очередь с приоритетом/куча  

Должна уметь 
* add(x) - *Добавить элемент*  
* pop() - *удалить min/max*  

индексы детей: 2 * pos + 1; 2 * pos + 2

Всегда заполнена по слоям

![heap](/img/Heap.png)

Каждый ребенок меньше/больше его родителя. Из-за заполнености по слоям, нет пропуска, можно деражть её на массиве например, но лучше дек. (В случае *Go* сосем бибу и делаем *Slice*)

Добавление идет по слоям --> добавляем в последний слой и делаем просеивание вверх

![HeapInsert](/img/HeapIns.png)

Удаление элемента

Выкидываем элемент на вершине и заменяем его самым последнем элементов *(это будет последний эелемент на последнем слое)*. Далее идет просеивание вниз, меняем элемент в вершине с меньшим из его детей *(случай кучи на максимум)*.

Есть тут одна ~~хуе~~*фича*. Значится при просеивании вниз, есть 4 случая:
* нет детей
* левый ребенок
* правый ребенок
* оба ребенка

Рассматривать 4 случая лень, можно удалить последний элемент из кучи *который мы посставили вверх* в конце операции, тогда, у каждого узла будет или 2 ребенка, или ни одного. 

С нет детей и двумя все понятно, всего в дереве может быть только 1 узел с 1 ребенком, и при удалении, если таковой имеется, именно его мы будем закидывать вверх, тогда рассматривать его на просеивании смысла нет, и мы свели задачу к 2 вариантам.

![HeapDel](/img/HeapDel.png)

Зачем эта куча нужна по итогу? 

Задача LRU Кеш
* Словарь - хранимые в кеши элементы
* Очередь с приоритетами - хранит приоритеты элементов, позволяет выкидывать самые старые элементы

![LRU](/img/LRUCash.png)

Задача о медиане в окне

![MedWin1](/img/MedWin1.png)

У нас есть 2 кучи, одна на минимум, другая на максимум. Таким образом искать медиану просто, достаточно пересчитывать кучки при сдвиге элементов

Пирамидальная сортировка

Идея в том, что бы поддерживать правильную кучу для каждого элемента

![Heapify](/img/Heapify.png)

Мы берем $\frac{N}{2}$ элементов, считаем их готовыми кучами, затем последовательно элемент за элементом просеиваем вниз, сохраняя инвариант готовой кучи

Тут не требуется доп память, в этом и прикол

#### Одномерное ДП

Вот вычисление чисел фибоначи через рекурсию в виде дерева: 

$F(n) = F(n-1) + F(n-2)$

![Fibonachi](/img/Fib1.png)

Проблема в том, что мы повторно будем вычислять числа, что не круто. Например при вычислении F(5), мы рекурсивно посчитаем F(4) и F(3), но, для F(4) мы будем тоже считать F(3).  

Рекурсия с мемоизацией *(c латыни meme - шутка, прикол, смешная картинка)*

![meme](/img/Memezation.png)

Классические задачи типа Ступенек, клеток с деньгами и т.д. *Обычно задачи типо : максимизировать, минимизировать, посчитать кол-во вариаций*

На каждом шаге рассматриваем варианты что было до и т.д. Кстати, тут частенько сложности $~O(N^2)$

**TO-DO** чекнуть задачу в конце на представление числа

#### Двумерное ДП

Параметры одной природы - тоже самое что и обычное, одномерное ДП, ничего интересного тут нет.

Тут типовые задачи на таблицу, дойти из одной клетку в другую max/min путь. Тут 2 параметра особой роли не играют, скорее выступают как "поле"

тут также можно абузить прием для сдвигов dx, dy, что бы не писать иф на 8 условий, например при переходе в таблице можно

```golang
var dx = []int{1, 0, 0, -1}
var dy = []int{0, 1, -1, 0}

а потом такие i, j = i + dx[k], j + dy[k]
```

