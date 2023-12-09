#include <iostream>
#include <unordered_map>
#include <sstream>
#include <string>
#include <utility>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string str;
    getline(cin, str);
    istringstream stream(str);

    unordered_map<int, int> a;

    int x;
    while (stream >> x) {
        if (a.count(x)) {
            ++a[x];
        } else {
            a.insert(pair<int, int>(x, 1));
        }
    }

    istringstream stream2(str);
    int y; // блять я ебашил гоу до 4 утра, я ваще сейчас не в состоянии разбираться с этим стримом блядским
    while (stream2 >> y) {
        if (a[y] == 1) {
            cout << y << " ";
        }
    }

    cout << endl;

    return 0;
}