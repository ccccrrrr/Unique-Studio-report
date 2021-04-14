#include <iostream>
#include <cstdio>
#include <bits/stdc++.h>
using namespace std;
void printlist(int* list, int N){
    for(int i = 0; i < N; i++){
        printf("%d ", list[i]);
    }
    printf("\n");
}
void insertsort(int* list, int N){
    if(N == 0){
        printf("list is empty. back...\n");
    }
	for(int i = 1; i < N; i++){
		int temp = list[i];
        int j = i - 1;
		for(; j >= 0 && list[j] > temp; j--){
			list[j+1] = list[j];
		}
		list[j+1] = temp;
	}
}

void _swap(int* a, int* b);
void HeapSort(int* list, int N);
void BuildHeap(int* list, int N);
void HeapAdjust(int* list,int root, int N);

void QuickSort(int* list, int N);
void quicksort(int* list, int N, int left, int right);

void MergeSort(int* list, int N);
int main(){
    int list[] = {9, 77,2,4,5,1,9,10,22,34};
    /** sort
     */
    //insertsort(list, 8);
    //HeapSort(list, 8);
    QuickSort(list, 10);

    return 0;

    
}

void HeapSort(int* list, int N){
    if(N == 0)
        printf("list is empty! back...\n");
    BuildHeap(list, N);
    printf("build heap: ");
    printlist(list, N);
    int len = N;
    while(len--)
    {
        _swap(&list[0], &list[len]);
        HeapAdjust(list, 0, len);
        printf("sort time %d : ", N - len);
        printlist(list, N);
    }
    printf("heap sort output: ");
    printlist(list, N);
}
void BuildHeap(int* list, int N){
    // ***
    for(int i = (N - 1) / 2; i >= 0; i--){
            HeapAdjust(list, i, N);
    }
}
void HeapAdjust(int* list, int root, int N){
    int child = root * 2 + 1;
    int parent = root;
    while(child < N){
        if(child + 1 < N && list[child] < list[child+1]) child++;
        if(list[parent] < list[child])
        {
            _swap(&list[parent], &list[child]);
        }
        parent = child;
        child  = parent * 2 + 1;
    }
}

void QuickSort(int* list, int N){
    if(N <= 0) {
        printf("error input....\n");
        return;
    }
    quicksort(list, N, 0, N - 1);
    printf("quick sort output ");
    printlist(list, N);
}
void quicksort(int* list, int N, int left, int right){
    if(left >= right)
        return;
    if(left + 1 == right){
        if(list[left] > list[right]){
            _swap(&list[left], &list[right]);
        }
        return;
    }
    int pivot = list[left];
    int i = left + 1, j = right;
    while(i != j){
        //what is import is that you should move right pointer first, or you will get wrong answer.
        while(i < j && list[j] >= pivot) j--;
        while(i < j && list[i] <= pivot) i++;
        _swap(&list[i], &list[j]);
        //printlist(list, 10);
    }
    _swap(&list[left], &list[i]);
    //printlist(list, 10);
    quicksort(list, N, left, i - 1);
    quicksort(list, N, i + 1, right);
}


void MergeSort(int* list, int N){

}
void _swap(int* a, int *b){
    int temp = *a;
    *a = *b;
    *b = temp;
}