#include <iostream>
#include <unordered_set>
#include <sstream>
#include <string>


using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string str;
    unordered_set<int> a;

    getline(cin, str);
    istringstream stream(str);

    int x;
    while (stream >> x) {
        if (a.count(x)) {
            cout << "YES" << endl;
        } else {
            a.insert(x);
            cout << "NO" << endl;
        }
    }

    
    return 0;
}