When a Go program starts up, the Go runtime asks the machine (virtual or
physical) how many operating system threads can run in parallel. This is based on
the number of cores that are available to the program. 
For each thread that can be run in parallel, the runtime creates an operating system
thread (M) and attaches that to a data structure that represents a logical processor
(P) inside the program.

G : application layer Thread   [200 ns context switch]
M Os thread will execute on Hw [1-2 us context switch]

			<Os Schedular>
			M ----> HW
Go routines ==========
			|
			P
<Go schedular>

Every P has Local Run Queue
and there is one Global Run Queue.

GC 
Syscall
	-Blocking
	-Cgo
every syscall is an opportunity to context switch 

Go schedular is cooperative identical to os scedular

Diff bw Goroutine and operating System thread
 	it doesnt have priorities , for hardware interrupts


Work: A set of instructions to be executed for a running application. This is
	accomplished by threads and an application can have 1 to many threads.

Thread: A path of execution that is scheduled and performed. Threads are
	responsible for the execution of instructions on the hardware

Thread States: A thread can be in one of three states: Running, Runnable, or
Waiting. Running means the thread is executing its assigned instructions on the
hardware by having a G placed on the M. Runnable means the thread wants time on
the hardware to execute its assigned instructions and is sitting in a run queue.
Waiting means the thread is waiting for something before it can resume its work.
Waiting threads are not a concern of the scheduler

Concurrency: This means undefined out of order execution. In other words, given
a set of instructions that would be executed in the order provided, they are
executed in a different undefined order, but all executed. The key is, the result of
executing the full set of instructions in any undefined order produces the same
result. I will say work can be done concurrently when the order the work is
executed in doesn’t matter, as long as all the work is completed.

Parallelism: This means doing a lot of things at once. For this to be an option, I
need the ability to physically execute two or more operating system threads at the
same time on the hardware.
CPU Bound Work: This is work that does not cause the thread to naturally move
into a waiting state. Calculating fibonacci numbers would be considered CPU-Bound
work.

I/O Bound Work: This is work that does cause the thread to naturally move into a
waiting state. Fetching data from different URLs would be considered I/O-Bound
work.

Synchronization: When two or more Goroutines will need to access the same
memory location potentially at the same time, they need to be synchronized and
take turns. If this synchronization doesn’t take place, and at least one Goroutine is
performing a write, I can end up with a data race. Data races are a cause of data
corruption bugs that can be difficult to find.
Orchestration: When two or more Goroutines need to signal each other, with or
without data, orchestration is the mechanic required. If orchestration does not take
place, guarantees about concurrent work being performed and completed will be
missed. This can cause all sorts of data corruption bugs.use waitgrp

IO bound work gets converted into CPU bound work

because operating system level thread keeps on running always
we are paying for context switch 200 ns

Go routine is so lightweight, 
we could build a web service that throws every request a go routine

have a mindset of fast Enough

