package main

type ArgsAndResult struct {
	res  chan obtainedData
	args string
	get  httpInterface
}

const maxChanelsSched = 100

/*
**A Task Worker is a Goroutine that is listening to a channel waiting for
**a new task to do, in this case the Task would be to download, canculate
**and send to the thread creator of the channel
**the minimum, maximum, total count and the sum of all prices
**
 */

func taskWorker(ch chan ArgsAndResult, workerId int) {
	for true {
		job := <-ch
		GetObtainedData(job.args, job.res, job.get)
	}
}

/*
**Initialize the Task Workers and returns the channel from where
**the workers will listen
**
 */
func startWorkers(workers int) chan ArgsAndResult {
	ch := make(chan ArgsAndResult)
	for i := 0; i < workers; i++ {
		go taskWorker(ch, i)
	}
	return ch
}
