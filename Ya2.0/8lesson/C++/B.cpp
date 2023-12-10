#include <iostream>

#include <unordered_map>
#include <string>
#include <sstream>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int N;
    cin >> N;

    unordered_map<string, string> tree;

    for (int i = 0; i != N - 1; ++i) {
        string child, ancs;
        cin >> child >> ancs;

        tree[child] = ancs;
    }

    string line;
    getline(cin, line);  // skip '\n'
    getline(cin, line);

    while (!line.empty()) {
        istringstream ss(line);
        
        string name1, name2;
        ss >> name1 >> name2;

        string it = name1;
        while (tree.count(it) && it != name2) {
            it = tree[it];
        }
        if (it == name2) {
            cout << 2 << " ";
            getline(cin, line);
            continue;
        }
        
        it = name2;
        while (tree.count(it) && it != name1) {
            it = tree[it];
        }
        if (it == name1) {
            cout << 1 << " ";
            getline(cin, line);
            continue;
        }

        cout << 0 << " ";
        getline(cin, line);
    }
    
    return 0;
}