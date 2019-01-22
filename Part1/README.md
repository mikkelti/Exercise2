# Mutex and Channel basics

### What is an atomic operation?
> Short answer: An operation that cannot be interrupted. Source: Downey.

> Longer answer: An operation where the tasks (threads/processes?) performing it cannot detect any state change other than those performed by themselves. In addition they do not reveal the changes they have made until the operation is completed. Source: BW ch. 7 (includes multiple definitions).

### What is a semaphore?
> A data structure useful for solving synchronization problems. Like an integer, but with the following distinct features:

> 1. After initialization it can only be incremented (signal()) and decremented (wait()), these are atomic operations.

> 2. If the value of the semaphore is negative after one thread decrements, it blocks (tells the scheduler that it cannot proceed). In reality a semaphore is non-negative, hence the decrementing will not be done until another thread has incremented.

> 3. If a thread increments, one waiting thread gets unblocked, if there are any waiting threads. 

### What is a mutex?
> A semaphore that implements mutual exclusion, controlling concurrent access to shared variables.

### What is the difference between a mutex and a binary semaphore?
> For a mutex, on thread calls mutex.wait(), performs an operation on a shared variable, and calls mutex.signal(). Meaning this thread "holds" the mutex until the critical section has been run (the wait- and signal-operations need to be called in that order, by any thread). 

> A binary semaphore is simply a semaphore that can only take the values 0 and 1. The difference is that the wait- and signal-operations don't need to be executed in any particular order. See: https://stackoverflow.com/questions/62814/difference-between-binary-semaphore-and-mutex.   

### What is a critical section?
> A section of statements that must be executed as an atomic operation (without interruption), for the system to work as required.

### What is the difference between race conditions and data races?
 > Race condition: The sequence of operations affect the correctness of a program.
 
 > Data race: Two threads try to access the same memory location concurrently, and the operations to be performed are not reads. Source: https://blog.regehr.org/archives/490. 

### List some advantages of using message passing over lock-based synchronization primitives.
> Easier to visualize, more predictable, no shared variables.

> Scalable.

### List some advantages of using lock-based synchronization primitives over message passing.
> Does not require allocation of memory for messages.

> Leads to more elegant solutions in some cases, but quickly gets very complex.
