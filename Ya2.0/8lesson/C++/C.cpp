#include <iostream>

#include <string>
#include <sstream>
#include <unordered_map>
#include <unordered_set>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);

    int n;
    cin >> n;
    unordered_map<string, string> tree;
    for (int i = 0; i != n - 1; ++i) {
        string child, ancs;
        cin >> child >> ancs;

        tree[child] = ancs;
    }

    string line;
    getline(cin, line);// прожевываем '\n'
    getline(cin, line);

    while (!line.empty()) {
        istringstream ss(line);

        string name1, name2;
        ss >> name1 >> name2;

        unordered_set<string> path;
        string it = name1;
        while (tree.count(it)) {
            path.insert(it);
            it = tree[it];
        }
        path.insert(it); // добавляем корень

        it = name2;
        while (tree.count(it) && !path.count(it)) {
            it = tree[it];
        }

        cout << it << endl;
        getline(cin, line);
    }

    return 0;
}