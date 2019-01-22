#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;  // What does this do? Set it to 1?

// Note the return type: void*
void* incrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&mutex);
	    i++;
        pthread_mutex_unlock(&mutex);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int j = 0; j < 1000001; j++) {
	pthread_mutex_lock(&mutex);
	i--;
    pthread_mutex_unlock(&mutex);
    }
    return NULL;
}


int main(){
    pthread_t incrementingThread, decrementingThread;
    
    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL); // Wait until incrementingThread has terminated
    pthread_join(decrementingThread, NULL); // Wait until decrementingThread has terminated
    
    printf("The magic number is: %d\n", i);
    return 0;
}
