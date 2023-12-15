## Список тем:

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
    * [Обход графов в глубину](#обход-графов-в-глубину)
    * [Обход графов в ширину](#обход-графов-в-ширину)
* ITMO Academy codeforces
    * [Z-функция(поиск подстроки в строке)](#z-функция)
    * [Дерево отрезков](#до)



### Сортировка подсчетом
Пусть дан массив из N элементов таких, что их *разброс невелик*. Обычная сортировка займет *~O(NLogN)*

Однако, если разброс элементов небольшой, проще сосчитать кол-во каждого элемента, а затем восстановить массив

![NumSort](/img/NumSort.png)

### Префиксные суммы
Подсчет *чего-либо* на отрезке, быстрый ответ на запрос *операции* на отрезке 

![PreffSum](/img/PreffSum.png)

### Два указателя
Много задач, формата 2 массива, станцуйте с бубном и т.д. Обычно 2 указателя там очевидны

#### Метод плавающего окна. Тригер к 2 указателям - соблюдение инварианта.

Должно выполнять 2 свойвства

* если отрезок [L,R] хороший, то любой *вложенный в него отрезок* также хороший
    * если отрезок [L,R] хороший, то любой отрезок, *который его содержит* также хороший
* Мы должны уметь пересчитывать функцию *(проверять хороший отрезок или плохой)*, при смещении левой или правой границы вправо.

![2pointers1](/img/2pointers1.png)

![2pointers2](/img/2pointers2.png)

Пример кода:

```golang
L = 0
for R = 0..n-1
    add(a[R])
    while not good():
        remove(a[L])
        L++
```

##### Метод очереди на 2-ух стеках для плавающего окна
Если при пересчете функции нам требуется иметь доступ не только к крайним элементам: например мы считаем разброс между *min* и *max* элементом в окне, удаление/прибавление элемента не ограничивается просмотром крайних.

В таком случае, можно сделать очередь на 2-ух стеках, в правом-левом стеке хранить необходимые значения *(например хранить в стеке будет храниться максимум, тогда при удалении/добавлении элемента, пересчет функии будет ~O(1))*

Абстрактная схема такой очереди:

![2pointersStack](/img/2pointersStackk.png)

### Бинарный поиск
Главный тригер к бинарке - *монотонная функция 2 видов*. Делеие на *хорошую* и *плохую* часть, с помощью бинарки удобно искать этот разрыв за *~O(LogN)*

![BinSearchFunc](/img/BinSearch1.png)

Структура бинпоиска не меняется, в общем виде это:

```golang
L := 0 // Левая граница
// крутые математики, сразу могут определить при каком R функция будет true, но д - думать, д - дебил, делаем так
R := 1
for !good(R, ...) {
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

#### Вещественный бинпоиск
Так же бинарка, однако достаточную *близость* *L* и *R* обрабатываем 2 способами:

```golang
// 1 eps - точность
eps := 1e-6
for (R - L > eps)

// 2 через итерации
for (i := 0; i != 100; i++)
// 100 итераций - 2^100, вообщем ответ получим с максимальной точностью, у нас на итерации ~60-той, уже будут соседние числа в float64
```

#### Минимакс/Максмини 
Задачи двух видов: $min(max(x_i))$ или $max(min(x_i))$  

Внешняя функцию, например *min(...)* можно заменить на ограничение сверху, а это ограничение подбирать бинаркой, получим $max(x_i) \leqslant T$. А это обычная монотонная функция.

Само выражение $max(x_i) \leqslant T$ обычно можно проверить за *~O(N)*.

#### Max/Min среднее
Пускай у нас есть какое-то множество, нам нужно выбрать подмножество, обладающее какими-то свойствами, и в этом подмножестве нужно максимизировать/минимизировать *среднее арифметическое элементов*. $\frac{\sum_{l}^r a_i}{r - l} \to max$

Делаем прием из прошлого блока, подбираем ограничение функции и преобразовываем:

$$
\begin{align}
    \frac{\sum_{l}^r a_i}{r - l} >= x \\
    \sum_{l}^r a_i >= x * (r - l) \\
    \sum_{l}^r (a_i - x) >= 0 \\
\end{align}
$$

Далее применим префиксные суммы:

$$
\begin{cases}
    p_r > p_l \\
    r - l >= D \\ 
\end{cases}
$$

**TO-DO** Поиск *K* элемента https://codeforces.com/edu/course/2/lesson/6/5

### Сортировка событий

**TO-DO**

### Стеки
**Задача на правильную скобочную последовательность** с различными скобками *(например ([)] - не получится обработать префиксной суммой)*. При добавлении новой скобки, нам не требуется пробегаться *по всем* скобкам в массиве, мы можем обработывать лишь скобки на вершине стека, каждая скобка будет обработа *лишь 1 раз*

**Задача на преобразование в постфиксную запись**, опять тот же подход.

Иными словами, признаком стека может являться некая *вложенность*, или возможность поэтапно обрабатывать данные. Строго признака нет, от задаче к задаче.

![преобразование в постфикс](/img/StackPostfix.png)

**Задача о поиске ближайшего меньшего справа**

![меньшийсправа](/img/StackLessRight.png)

Храним в стеке значения элементов массива. Тогда, обработаем число `x`:
1. Поочередно рассматриваем элементы на вершине стека, если `stack.top()` > x, выбрасываем из стека
2. Если `stack.top()` < x, то дальнейшее рассмотрение не имеет смысла, ведь такой элемент, был бы уже обработан заранее.

### Очереди, и диквэй
Помимо стандарной реализации очереди через буферы и бакеты, имеет место быть реализация, через *кольцевой буфер*:

![dequeLoop](/img/LoopBuff.png)

#### Куча
Структура, позволяющая получить *min/max* элемента за *~O(LogN)*
* add(x) - *Добавить элемент*  
* pop() - *удалить min/max*  

Каждый ребенок меньше/больше его родителя. Куча заполняется *по слоям*, из-за этого её хранение в массиве компактно и удобно.

![heap](/img/Heap.png)

**Добавление в кучу** - элемент добавляется в конец очереди и просеивается вверх.

![HeapInsert](/img/HeapIns.png)

**Удаление элемента** - Элемент в вершине кучи заменяется последним элементом кучи, после чего идет просеивание вниз.

*Просеивание вниз* - элемент в узле меняется с меньшим/большим из детей.

Если при замене элемента, удалить последний, после просеивания, можно свести все узлы к ситуации: 0 или 2 детей. *(свап с неудленным элементом в любом случае не требуется)*

![HeapDel](/img/HeapDel.png)

**Синхронизированная куча** - Можно хранить hashMap, *{значение:индекс в массиве кучи}*, таким образом, изменять элемент в куче и просеивать его. **TO-DO** 

**Задача LRU Кеш**: Хранить в кеше 
* Словарь - хранимые в кеши элементы
* Очередь с приоритетами - хранит приоритеты элементов, позволяет выкидывать самые старые элементы

![LRU](/img/LRUCash.png)

**Задача о медиане в окне**: Есть массив из n элементов, нужно найти *медиану*. Например: *{1, 9, 9321}* - медиана 9.

Поддерживаем 2 кучи, на максимум и на минимум. Ответом будет - минимальный элемент в 2-ой куче.

![MedWin1](/img/MedWin1.png)

**Пирамидальная сортировка**: Мы берем $\frac{N}{2}$ элементов, считаем их готовыми кучами, затем последовательно элемент за элементом просеиваем вниз, сохраняя инвариант готовой кучи. Подобная сортировка не требует дополнительной памяти.

![Heapify](/img/Heapify.png)

### Одномерное ДП
При рекрсивных вычислениях, мы повторно запускаем рекурсию от ранее вычисленных значениях.
Пример на числах фибоначи: $F(n) = F(n-1) + F(n-2)$

![Fibonachi](/img/Fib1.png)

**Рекурсия с мемоизацией**: Убираем рекурсию в пользу цикла.

![meme](/img/Memezation.png)

Классические задачи одномерного ДП: максимизировать, минимизировать, посчитать кол-во вариаций. Разбиваем задачи на более маленькие, рассматриваем все варианты.

**TO-DO** задача в конце на представление числа

#### Двумерное ДП

##### Параметры одной природы

Одномерное ДП, но сводящееся к таблице. Типовые задачи на поле: дойти от первой клетки до последний макс/мин при этом параметр.

*Аналог писать ифы для сдвига по нескольким направлениям - писать подобные массивы*

```golang
var dx = []int{1, 0, 0, -1}
var dy = []int{0, 1, -1, 0}

i, j = i + dx[k], j + dy[k]
```
##### Параметры разной природы
**Задача про черепашку без долгов**: 2 параметра - минимальные долги, максимальная прибль. В данном случае минимальные долги, можно подобрать *бинарным поиском*, если хватит *x* монет на старте, то *x+1* тоже хватит.

**Задача о наибольшей общец подпоследовательноти**: Формируем таблицу *первое/второе* слово, *dp[i][j] =* НОП далее придерживаемся правил:
* *X == Y* - *dp[i][j]++*, очевидно мы берем 2 новые буквы в ответ
* *X != Y* - *dp[i][j] = max(dp[i-1][j], dp[i][j-1])* - в этих строках/столбцах мы уже рассмотрели *всевозможные* комбинации новых букв. *Обе в ответ мы не берем*

**Задача о редакционном расстоянии левенштайна**: Есть 2 слова, операции вставки, изменения, удаления. За наименьшее количество операций преобразовать слово *A -> B*. Строим таблицу, сверху/сбоку слова A/B. Если буквы совпадают +0, если отличаются то минимум от +1 с трех сторон. В dp[i][j] храним минимум изменений.

![Levenstain](/img/Levenstain.png)

**Задача о обедах и купонах**: Массив цен на N дней вперед. При покупке на более 100р, получаем купон на бесплатный обед. Посчитать минимальные затраты. Таблица дни/купоны.

##### ДП на подотрезках

**Задача про правильную скобочную последовательность**: Удалить минимальное количество скобок, что бы последовательность стала правильной.

Последовательность формируется как *S =* $\lambda$ *| SS | (S) | [S] | {S}*. Пусть dp[i][j] мин. кол-во скобок требуется удалить на *[L,R]*.

Тогда, если у нас последовательность $[...)$, где скобки не соответствуют, тогда разбиваем её на две, и ищем грубо говоря $[...](...)$, то есть, мы ищем по массиву dp: $min(dp[L, i] + dp[i+1, R])$. Однако, если у нас например вложенная последовательность $(...)$, тогда имеет место выражение $dp[L, R] = dp[L+1, R-1]$. Итого в общем случае:

```
if s[L] соответствует s[R] {
    dp[L][R] = dp[L+1][R-1]
} 
for i = L+1; i <= R-1; i++ {
    dp[L][R] = min(dp[L][R], dp[L][i] + dp[i+1][R])
}
```

сложность $~O(N^3)$

Задача про упаковку символов - задача говна, dp[i][j] = кол-во символов в которые можно упаковать + всегда при [L,R] надо проверить на разбитие на делители, мало ли можно свернуть

Ну и каличь ес честно

#### Обход графов в глубину

Обход в глубину нужен для понимания *структуры* графа!

Графы храним как:
    * таблица смежности
    * список ребер
    * списки смежности

в основном будем юзать *список смежности*

```go
func dfs(g [][]int, visited []bool, now int) {
    visited[now] = true
    for v := range g[now] {
        if !visited[v] {
            dfs(g, visited, v)
        }
    }
}
```

Обход в глубину, это всегда *грубо-говоря* линия

Поиск циклов: раскраска в 2 цвета не работает, мы можем прийти 2 путями. Красим в 3 цвета:
* Белый - непосещенные вершины
* Серый - активная вершине, присутствующая в стеке dfs
* Черный - обработанная вершина.

Важно понимать, что если вершина черная - мы уже обработали всех её детей.

Для поиска компонент связанности - после dfs получаем компоненту связанности.

Для поиска циклов - делаем обычный dfs, если при обходе встречаем серую вершину - мы в цикле

Раскраска в 2 цвета: двудольные графы. Делаем обычный обход поочередно крася вершины в 2 цвета. Если мы полностью обошли граф, и не встретили соседних вершин с одинаковыми цветами - граф двудольный.

Топологическая сортировка: Можно делать обычный dfs, и докидывать в массив вершины, покрашенные в черный цвет. 

#### Обход графов в ширину

Глобально в *старой школе* bfs реализовывался через очередь. Добавления по слоям будет выглядить как то так:

```c++
while (!q.empty()) {
    int layer = q.size();
    for (int i = 0; i != layer; i++) {
        int vertex = q.front();
        for (auto x : graph[vertex]) {
            if (!visited[x]) {
                q.push_back(x);
            }
        }
    }
}
```

Можно делать через массив, индекс которого - шаг bfs-а. 

![BFSARRAY](/img/BFS1.png)

Восстановление пути: массивчик prev для каждой вершины, будет показывать откуда мы пришли

Задача о ребрах и вершинах на кратчайших путях: Делаем bfs от начальной до конечной вершины в 2 стороны. $Start \rightarrow End$ и $Start \leftarrow End$. Далее, если номер слоя вершины *при правом обходе* сложить с номером слоя вершины *при левом обходе*, если мы получаем длину кратчайшего пути, вершина принадлежит кратчайшему пути.

Граф состояний: **TO-DO**. Насколько я понял у нас что-то вроде дерева графов, но выглядит сложновато.

Кратчайший путь в 0-1 графе: Если пишем на очереди, то при 0 ребре просто кидаем в начало очереди. 

Если пишем на массиве тут интересней: **TO-DO**, хзхз

Кратчайший путь в 0-k графе: ~~не занимаемся херней и юзаем дейкстрку~~, можно сделать как то хитро на массиве с heap-ом, но реализация сведется к дейкстре по итогу.

#### Z-функция

Задача о поиске подстроки в строке решается 2 основными способами:
* Предпроцессинг образа (*Z-func, префиксные функции*)
* Предпроцессинг текста (*суффиксное дерево, массив, автомат*)

Z-функцие строки *s* называется такой массив длины *n(n = len(s))*, где *z[i]* = длина наибольшего префикса строки s и её суффикса s[i..n]. Обычно полагают *z[0]=0*

Что можно делать, зная Z-функцию? 

Задача о поиске вхождений: есть строка p, t. Нужно найти все вхождения p в t:
* строим строку `p + $ + t`
* считаем Z-функцию полученной строки
* z[i] = p  $\Longleftrightarrow$, когда t[i..i + (len(p) - 1)]

$~O(N)$

Период строки: если z[i] + i == n && len(z) % i == 0.

Строка с одной ошибкой: 2 прохода в стороны z-funcом, ищем нужные индексы с 2-ух сторон. 

Задача о количестве *различных* подстрок за $~O(N)$.

![Z-func](/img/Z-func.png)

Идем справо-налево добавляя по одной букве. Таким образом добавляя *k-ую* букву, кол-во уникальных подотрезков для *k-1* уже посчитано. При добавлении новой буквы, новый отрезки могут получится только с ней, то бишь будут префиксами различной длинны. Считаем *Z-func* от нашей текущей строки. Пусть её длина, *m*, не учитывая уникальный подстроки, а считая все подряд, мы бы добавили *m* новых строк. *(просто все префиксы)* Тогда, если мы возьмем максимум в нашей *Z-func*, мы получим кол-во уже посчитанных подстрок. Итоге делаем $+m - max(z_i)$

Z-algo, за $~O(N)$

Делаем 2 индекса l, r. Бежим по строке с 1-ой *(z~~xc~~[0] = 0)* позиции. l и r будут обозначать совпадение префикса и подстроки с i-ого индекса так, что r - максимально.

Далее *(только что понял, Z-функция это же чистое dp)* идя по строке, глобально натыкаемся на 2 случая
1. i < r: Отрезок [l, r] совпадает с префиксом, соответственно s[i.. r] == s[i-l..r-l]. Тогда зная это совпадение *(а оно может быть и больше)* можем воспользоваться ранее вычисленной z-func. Тогда `z[i] = min(r-i+1, z[i-l])`. Если z[i-l] *не дойдет* до r, то все впорядке, но если оно выходит за рамки, мы не знаем что находится за границей r, в этом случае нам надо идти вручную поэлементно и проверять.

2. i > r: Полная неизвестносить, идем вручную и проверяем, потом обновим r.

Эти 2 случая можно склеить в один, получим подобное:

*Здесь l, r - ПОЛУИНТЕРВАЛ!*

```golang
var l, r int
	for i := 1; i != len(z); i++ {
		if i < r {
			z[i] = min(r-i, z[i-l])
		}
		for z[i] + i < len(z) && s[i + z[i]] == s[z[i]] {
			z[i]++
		}
		if r < i + z[i] {
			r = i + z[i]
			l = i
		}
	}
```

#### ДО

Дерево отрезков, выглядит эту штука как: 

![DO1](/img/DO1.png)

Базово - представляем из себя дерево, позволяющее выполнять следующие операции за $~O(logN)$
* set(i, v) - присвоить элемету i значение v
* calc(l, r) - посчитать операцию calc на отрезке **Операция обязанна быть ассоциативной**

Очень похоже на префиксные суммы, но с возможностью пересчета за логарифм

set(i, v) - изменяется элемент массив *(до которого дошли рекурсией)*, и по цепочке вверх идет пересчет элементов, за счет *ассоциативности* функции

calc(l, r) - посчитаить функцию от отрезка [L, R]. Каждый элемент в дереве - является результатом функции от подотрезка всего массива, выглядит это все следующим образом: 

![DO2](/img/DO2.png)

Соответственно, если $[l_i, r_i] \in [l, r]$, то углублятся дальше в дерево смысла нет. Иначе следует рассмотреть рекурсивно детей узла. Если подотрезок узла вне искомого отрезка, просто выходим из узла *возвращаем нейтральный элемент для функции*

Пример:

```c++
void set(int i, int v, int x, int lx, int rx) {
    // i, v - параметры фукнции
    // x, lx, rx - номер, левая-правая граница узла в дереве(идем рекурсивно)

    if (rx == lx + 1) {
        // случай когда дошли до последнего слоя, дальше углубляться нет смысла
        // С учетом set'а, в данном блоке мы окажемся, если нашли искомый элемент - можем менять его и выходить из рекурсии
        array[x] = v; // тут важно понимать, array - массив всех узлов, последним слоем которого является сам массив, над которым идет надстройка
        // lx, rx - индексы за что отвечает узел в массиве значений. x - индекс узла!, это разные вещи!
        // например первый элемент в последнем слое будет иметь: lx = 0; rx = 1(полуинтервал), x = 7.
        return;
    }

    // 2*x+1 - левый ребенок, 2*x+2 - правый.
    mid = (lx + rx)/2;
    if (i < lx) {
        set(i, v, 2*x+1, lx, mid);
    } else {
        set(i, v, 2*x+2, mid, rx);
    }

    // изменив элемент в дереве, мы пересчитываем все что выше него!
    array[x] = op_calc(array[2*x+1], array[2*x+2]); // op_calc - сама ассоциативная функция от 2-ух элементов
    return;
}

Type calc(int l, int r, int x, int lx, int rx) {
    // l, r - полуинтервал
    if (rx <= l || lx >= r) { // полностью промахнулись
        return NEUTRAL 
    }
    if (lx >= l && rx <= r) { // полностью попали
        return array[x]
    }

    int mid = (lx + rx)/2;
    Type left = calc(l, r, 2*x+1, lx, mid);
    Type right = calc(l, r, 2*x+2, mid, rx);

    return op_calc(left, right);
}
```

Инициализация дерева: 
```c++
clacc segtree {
    public:
        segtree(int n, vector<Type>& origin) {
            init(n);
            build(0, 0, size, origin);
        }

        calc(int l, int r) {...} // публичные методы calc && set
        set(int i, int v) {...}
    
    private:
        vector<Type> array; // все узлы дерева
        int size; // размер основания дерева
        // основание дерева всегда степерь 2-ки, так тупо удобнее, дерево похоже на бинарное

        void init(int n) { // n - размер искходного массива
            size = 1;
            while (size < n) {
                size *= 2;
            }
            array.assign(2*size-1, START_ELEMENT); // 2*size-1 - количество узлов в дереве
        }

        void build(int x, int lx, int rx, vector<Type>& origin) { // origin - массив входных данных
            // ну тут просто dfs
            if (rx == lx + 1) {
                if (lx < origin.size()) {
                    // основание дерева - степень 2-ки, значит у нас будут *неиспользованые* элементы, lx - индекс в массиве, надо проверить что мы не в таких элементах
                    array[x] = origin[lx];
                    return;
                }
            }

            int mid = (lx + rx)/2;
            build(2*x+1, lx, mid, origin);
            build(2*x+2, mid, rx, origin);
            array[x] = op_calc(array[2*x+1], array[2*x+2])
        }

        calc(int l, int r, int x, int lx, int rx) {...}
        set...
};
```

Примеры задач: 

*Любая, где мы можем считать функцию ассоциативно, проще - зная значения на 2-ух отрезков, можем пересчитать её на общем*
Мимум && кол-во минимальных элементов: обе компоненты (минимум и кол-во считаются ассоциативно)

Отрезок с максимальной суммой: ~~не выебываться и решать префиксами, если нет изменения~~ У нас есть 2 отрезка, что в них нужно хранить, что бы при объединении, можно было пересчитать максимальную сумму?

* сегмент с максимальной суммой, тогда при объединении `maxSeg = max(Seg1, Seg2)`
* preff, suff - храня просто сегменты, мы не учитываем отрезке, получаемые непосредственно в процессе объединения. Тогда: `maxSeg = max(Seg1, Seg2, suff1 + preff2)`

Как пересчитывать preff и suff? `preff = max(preff1, sum1 + preff2)`, с suff - аналогично

K - единица. Просто дерево с операциями 
* set(i, v) - $v \in \{0, 1\}$
* find(k) -  найти индекс k-ой единицы

Абстрактно - есть массив плохих/хороших элементов, можем менять плохо на хороший и обратно. Найти индекс k-го хорошоге элемента, используется *часто*

Первый больше X, (не монотонный массив, без бинар очки). Просто дерево на максимум

Инверсии, Вложенные/Пересекающиеся отрезки - [codeforces](https://codeforces.com/edu/course/2/lesson/4/3)