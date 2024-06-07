## overview 

 The "sync" package in Go provides a WaitGroup that allows waiting for the completion of all goroutines.
It has three primary methods namely,


 wg.Add() specify the number of goroutines to wait for.
-wg.Wait()  indicates it wait for all goroutines to finish,
- wg.Done()  indicates that a goroutine has finishes its task.