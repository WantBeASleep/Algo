#include <iostream>
#include <deque>
#include <string>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string line;
    deque<int> q;
    while (getline(cin, line) && line != "exit") {
        if (line == "pop_front") {
            if (q.empty()) {
                cout << "error" << "\n";
            } else {
                cout << q.front() << "\n";
                q.pop_front();
            }
        } else if (line == "pop_back") {
            if (q.empty()) {
                cout << "error" << "\n";
            } else {
                cout << q.back() << "\n";
                q.pop_back();
            }
        } else if (line == "front") {
            if (q.empty()) {
                cout << "error" << "\n";
            } else {
                cout << q.front() << "\n";
            }
        } else if (line == "back") {
            if (q.empty()) {
                cout << "error" << "\n";
            } else {
                cout << q.back() << "\n";
            }
        } else if (line == "size") {
            cout << q.size() << "\n";
        } else if (line == "clear") {
            q.clear();
            cout << "ok" << "\n";
        } else {
            if (line.substr(0, 9) == "push_back") {
                q.push_back(stoi(line.substr(10)));
            } else {
                q.push_front(stoi(line.substr(11)));
            }
            cout << "ok" << "\n";
        }
    }

    cout << "bye" << "\n";
    return 0;
}