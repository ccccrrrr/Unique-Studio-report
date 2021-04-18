#include <iostream>
using namespace std;
#include <unordered_map>
#include <list>
struct node {
    int key, value;
    node(int key, int value){
        this->key = key;
        this->value = value;
    }
};
struct LRUCache {
    int capacity;
    unordered_map<int, list<struct node*>::iterator> _map;
    list<struct node *> _list;
    LRUCache(int c){
        this->capacity = c;
    }
    int get(int key){
        if(this->_map.size() == 0){
            printf("[warning] cache is empty\n");
            return -1;
        }
        if(this->_map.find(key) == this->_map.end()){
            printf("[warning] finding error!\n");
            return -1;
        }else{
            auto temp = _map.find(key);
            //STL C++ splice: move the element int list to first of the list
            _list.splice(_list.begin(), _list, temp->second);
            printf("[finding] output is [%d, %d]\n", temp->first, (*temp->second)->value);
            return temp->first;
        }
    }
    void put(int key, int value){
        if(_map.find(key) != _map.end()){
            printf("already exist!\n");
            return;
        }
        
        if(_map.size() >= capacity)
        {
            struct node * temp = _list.back();
            printf("[warning] cache is already full... the data: [%d, %d] is deserted\n", temp->key, temp->value);;
            printf("[%d, %d] is put in cache\n", key, value);
            _map.erase(temp->key);
            _list.pop_back();
            struct node * t = new node(key, value);
            _list.push_front(t);
            _map.emplace(t->key, _list.begin());
            
        }
        else{
            struct node * t = new node(key, value);
            _list.push_front(t);
            _map.emplace(t->key, _list.begin());
            printf("[%d, %d] is put in cache sucessfully\n", key, value);
        }
        
    }
};
int main(){
    LRUCache * cache = new LRUCache(3);
    cache->put(2, 4);
    cache->put(3, 5);
    cache->put(4, 6);
    cache->put(5, 8);
    cache->get(2);
    //cache->get(4);
    //cache->put(6, 7);
    //cache->put(7, 8);
    cache->put(0, 88);
    cache->get(3);
    return 0;
}