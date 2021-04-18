#include <bits/stdc++.h>
#include <list>
using namespace std;
#define N 7 //vertex number
#define M 9 //edge number
#define _N 10
#define _M 10
struct Node {
    int index;
    struct node * next;
};
struct node {
    int index;
    int weight;
    struct node * next;
};
struct Node * v[N];
struct Node * k[_N];
void CreateGraph(void);
void Create_another_graph(void);
void bfs(int);
void dfs(int);
void _dfs(int, stack<int>&, int*);
void topsort(void);
void calc_in_number(int*);
int main(){
    CreateGraph();

    printf("create successfully\n");
    //ttt();
    //v[3]->next->next->index;
    //BFS
    printf("bfs result: ");
    bfs(2);
    printf("\n");
    //DFS
    printf("dfs result: ");
    dfs(2);
    printf("\n");

    //TOP SORT
    Create_another_graph();
    topsort();

}

void bfs(int s){
    //O(N+E)
    list<int> l;
    l.push_back(s);
    int visited[N] = {0};
    visited[s] = 1;
    
    while(l.size() > 0){
        int temp = l.front();
        printf("%d ", temp);
        struct node * t = v[temp]->next;
        while(t != nullptr){
            if(visited[t->index] == 0){
                l.push_back(t->index);
                visited[t->index] = 1;
            }
            t = t->next;
        }
        l.pop_front();
    }
    
}

void dfs(int s){
    //O(N+E)
    stack<int> m;
    int visited[N] = {0};
    _dfs(s, m, visited);
}

void _dfs(int s, stack<int>& m, int* visited){
    m.push(s);
    printf("%d ", s);
    struct node * t = v[s]->next;
    visited[s] = 1;
    while(t != nullptr && m.size() > 0){
        if(visited[t->index] == 0){
            //printf("%d ", t->index);
            m.push(t->index);
            _dfs(t->index, m, visited);
            //visited[t->index] = 0;
            m.pop();
        }
        t = t->next;

    }
}

void topsort(void){
    int _list[_N] = {0};
    calc_in_number(_list);
    queue<int> q;
    for(int i = 0; i < _N; i++)
        if(_list[i] == 0) q.push(i);
    struct node * temp;
    while(!q.empty()){
        printf("%d ", q.front());
        if(k[q.front()]->next != nullptr){
            temp = k[q.front()]->next;
            while(temp != nullptr){
                _list[temp->index]--;
                if(_list[temp->index] == 0)
                    q.push(temp->index);
                temp = temp->next;
            }
        }
        q.pop();
    }
}

void calc_in_number(int* _list){
    struct node * temp;
    for(int i = 0; i < _N; i++)
    {
        if(k[i]->next == nullptr)
            continue;
        temp = k[i]->next;
        while(temp){
            _list[temp->index]++;
            temp = temp->next;
        }
    }
    printf("successfully calculate in number...\n");
}

void Create_another_graph(void){
    for(int i = 0; i < _N; i++)
        k[i] = new Node;
    int s[_M] = {0,0,1,2,2,2,3,4,7,9};
    int t[_M] = {1,2,3,3,4,6,5,7,8,2};
    for(int i = 0; i < _M; i++){
        struct node * temp;
        if(k[s[i]]->next == nullptr){
            k[s[i]]->next = new node;
            k[s[i]]->next->index = t[i];
            k[s[i]]->next->next = nullptr;
        }else{
            temp = k[s[i]]->next;
            while(temp->next) temp = temp->next;
            temp->next = new node;
            temp->next->index = t[i];
            temp->next->next = nullptr;
        }
    }
    printf("successfully create the graph!\n");
}

void CreateGraph(void){
    for(int i = 0; i < N; i++){
        v[i] = new Node;
        v[i]->next = nullptr;
    }
    int s[M] = {0,0,0,0,1,3,2,5,2};
    int t[M] = {1,4,2,6,3,5,6,6,4};
    for(int i = 0; i < M; i++){
        //int s, t;
        //printf("input connected vertex index, and ensure that there is no repeated edge: ");
        //scanf("%d %d", &s, &t);
        struct node * temp;
        
        if(v[s[i]]->next == nullptr){
            v[s[i]]->next = new node;
            v[s[i]]->next->index = t[i];
            v[s[i]]->next->next = nullptr;
        }
        else{
            temp = v[s[i]]->next;
            while(temp->next) temp = temp->next;
            temp->next = new node;
            temp->next->index = t[i];
            temp->next->next = nullptr;
        }
        if(v[t[i]]->next == nullptr){
            v[t[i]]->next = new node;
            v[t[i]]->next->index = s[i];
            v[t[i]]->next->next = nullptr;
        }
        else{
            temp = v[t[i]]->next;
            while(temp->next) temp = temp->next;
            temp->next = new node;
            temp->next->index = s[i];
            temp->next->next = nullptr;
        }
        
    }
    

    /*
    v[0]->next = new node;
    struct node* temp;
    temp = v[0]->next;
    temp->index = 1;
    //temp->weight = 0;
    temp->next = new node;
    temp = temp->next;
    temp->index = 2;
    temp->next = new node;
    temp = temp->next;
    temp->index = 3;
    temp->next = nullptr;

    v[1]->next = new node;
    temp = v[1]->next;
    temp->index = 0;
    temp->next = new node;
    temp = temp->next;
    temp->index = 2;
    temp->next = nullptr;

    v[2]->next = new node;
    temp = v[2]->next;
    temp->index = 0;
    temp->next = new node;
    temp = temp->next;
    temp->index = 1;
    temp->next = new node;
    temp = temp->next;
    temp->index = 3;
    temp->next = nullptr;

    v[3]->next = new node;
    temp = v[3]->next;
    temp->index = 0;
    temp->next = new node;
    temp = temp->next;
    temp->index = 2;
    temp->next = nullptr;
    */

}