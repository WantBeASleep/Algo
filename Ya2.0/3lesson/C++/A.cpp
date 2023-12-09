#include <iostream>
#include <unordered_set>
#include <string>
#include <sstream>

using namespace std;

int main()
{
    unordered_set<int> a;

    string str1;
    getline(cin, str1);

    istringstream stream1(str1);

    int x1;
    while (stream1 >> x1) a.insert(x1);

    int count = 0;

    string str2;
    getline(cin, str2);

    istringstream stream2(str2);
    int x2;

    while (stream2 >> x2) {
        // cout << x2 << endl;
        if (a.count(x2)) ++count;
    }

    cout << count << endl;

    return 0;
}