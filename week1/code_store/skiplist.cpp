#include <iostream>
#include <stdlib.h>
#include <cstdio>
#include <time.h>
#include <random>
using namespace std;
#define N 4
#define MAXLEVEL 7
struct node {
    int key;
    int value;
    struct node * forward[MAXLEVEL]; // just store the head node in highest level
};
int randomLevel(int MaxLevel);
struct ListStructure {
    int level;
    struct node * head;
    ListStructure(int level){
        this->level = level;
    }
    int Search(int key){
        struct node * x = this->head;
        
        for(int i = this->level - 1; i >= 0; i--){
            while(x->forward[i]->key < key){
                x = x->forward[i];
            }
        }
        
        x = x->forward[0];
        if(x->key == key){
            printf("sucessful search: key: %d value: %d\n", x->key, x->value);
            return x->value;
        }
        else{
            printf("key: %d not found!\n", key);
            return -1;
        }
    }
    void _insert(int key, int value){
        struct node * update[MAXLEVEL];
        struct node * x = this->head;
        for(int i = this->level - 1; i >= 0; i--){
            while(x->forward[i]->key < key){
                x = x->forward[i];
            }
            update[i] = x;
        }
        x = x->forward[0];
        if(x->key == key){
            printf("already have key: %d, insert operation change into update opeartion\n", key);
            x->value = value;
            return;
        }else{
            int newLevel = randomLevel(this->level);
            //let's assume that newlevel will not be greater than present max level...
            
            if(newLevel > this->level){
                for(int i = this->level; i < newLevel; i++){
                    update[i] = this->head;
                this->level = newLevel;
                }
            }
            
            x = new node;
            x->key = key;
            x->value = value;
            for(int i = newLevel - 1; i >= 0; i--){
                x->forward[i] = update[i]->forward[i];
                update[i]->forward[i] = x;
            }

            printf("sucessfully add [%d, %d], in level%d\n", key, value, newLevel);
            
        }
    }
    void _delete(int key){
        struct node * update[this->level];
        struct node * x = this->head;
        for(int i = this->level - 1; i >= 0; i--){
            while(x->forward[i]->key < key)
                x = x->forward[i];
            update[i] = x;
        }
        x = x->forward[0];
        if(x->key == key){
            for(int i = 0; i < this->level; i++){
                if(update[i]->forward[i] != x)
                    break;
                update[i]->forward[i] = x->forward[i];
            }
            printf("delete [%d, %d]\n", key, x->value);
            /*
            if(this->head->forward[this->level] == nullptr){
                printf("null!");
            }
            */
            /*
            while(this->level >= 0 && this->head->forward[this->level]->key == 0xFFFF){
                this->level--;
            }
            */
            free(x);

            
        }
        else{
            printf("[warning] key: %d not found\n", key);
        }
    }
    void update(int key, int new_value){
        struct node * x = this->head;
        for(int i = this->level - 1; i >= 0; i--){
            while(x->forward[i]->key < key){
                x = x->forward[i];
            }
        }
        x = x->forward[0];
        if(x->key == key){
            printf("sucessful update [%d, %d] => [%d, %d]\n", x->key, x->value, x->key, new_value);
            x->value = new_value;
        }
        else{
            printf("key: %d not found!\n", key);
        }
    }
    void search_range(int left_key, int right_key){
        struct node * update[this->level];
        struct node * x = this->head;
        for(int i = this->level - 1; i >= 0; i--){
            while(x->forward[i]->key < left_key)
                x = x->forward[i];
            update[i] = x;
        }
        x = x->forward[0];
        printf("key range %d - %d result: ", left_key, right_key);
        while(x->key <= right_key && x->key != 0xFFFF){
            printf("[%d, %d]  ", x->key, x->value);
            x = x->forward[0];
        }
        printf("\n");
    }
};
struct ListStructure * CreateSkipList(void);
int main(){
    srand(time(0));
    struct ListStructure * skipList = CreateSkipList();
    skipList->Search(6);
    skipList->Search(26);
    skipList->Search(9);
    skipList->Search(10);
    skipList->_insert(10, 10);
    //printf("level: %d\n", skipList->level);
    skipList->_insert(11, 11);
    //printf("level: %d\n", skipList->level);
    skipList->_insert(12, 12);
    //printf("level: %d\n", skipList->level);
    skipList->_insert(13, 13);
    //printf("level: %d\n", skipList->level);
    skipList->_insert(14, 14);
    //printf("level: %d\n", skipList->level);
    
    skipList->_delete(14);
    //printf("level: %d\n", skipList->level);
    skipList->_delete(13);
    //printf("level: %d\n", skipList->level);
    skipList->_delete(12);
    //printf("level: %d\n", skipList->level);
    skipList->_delete(11);
    //printf("level: %d\n", skipList->level);
    skipList->_delete(10);
    //printf("level: %d\n", skipList->level);
    skipList->_delete(222);
    //printf("level: %d\n", skipList->level);
    
    //skipList->Search(10);
    //skipList->_delete(10);
    //skipList->Search(10);
    //skipList->Search(9);
    skipList->search_range(1, 20);
    return 0;
}
struct ListStructure * CreateSkipList(void){
    struct ListStructure * skiplist = new ListStructure(N);
    struct node * head = new node;
    struct node * nodeList[11];
    nodeList[10] = new node;
    nodeList[10]->key = nodeList[10]->value = 0xFFFF;
    for(int i = 0; i < 11; i++){
        nodeList[i] = new node;
        for(int j = 0; j < MAXLEVEL; j++)
            nodeList[i]->forward[j] = nodeList[10];  
    }
    for(int i = 0; i < MAXLEVEL; i++)
        head->forward[i] = nodeList[10];
    head->forward[0] = nodeList[0];
    head->forward[1] = head->forward[2] = head->forward[3] = nodeList[1];
    head->key = head->value = 0;
    nodeList[0]->key = nodeList[0]->value = 3;
        nodeList[0]->forward[0] = nodeList[1];
    nodeList[1]->key = nodeList[1]->value = 6;
        nodeList[1]->forward[0] = nodeList[2];
        nodeList[1]->forward[1] = nodeList[3];
        nodeList[1]->forward[2] = nodeList[8];
    nodeList[2]->key = nodeList[2]->value = 7;
        nodeList[2]->forward[0] = nodeList[3];
    nodeList[3]->key = nodeList[3]->value = 9;
        nodeList[3]->forward[0] = nodeList[4];
        nodeList[3]->forward[1] = nodeList[5];
    nodeList[4]->key = nodeList[4]->value = 12;
        nodeList[4]->forward[0] = nodeList[5];
    nodeList[5]->key = nodeList[5]->value = 17;
        nodeList[5]->forward[0] = nodeList[6];
        nodeList[5]->forward[1] = nodeList[8];
    nodeList[6]->key = nodeList[6]->value = 19;
        nodeList[6]->forward[0] = nodeList[7];
    nodeList[7]->key = nodeList[7]->value = 21;
        nodeList[7]->forward[0] = nodeList[8];
    nodeList[8]->key = nodeList[8]->value = 25;
        nodeList[8]->forward[0] = nodeList[9];
    nodeList[9]->key = nodeList[9]->value = 26;
    skiplist->head = head;
    printf("create sucessfully!\n");
    return skiplist;
}

int randomLevel(int MaxLevel){
    int level = 1;
    int temp = 0;
    while(true){
        temp = rand() % 100 + 1;
        if(temp <= 50 || level >= MaxLevel)
            return level;
        else
            level++;
    }
    return level;
}