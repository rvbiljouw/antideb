package antideb

import (
	"time"
	"unsafe"
)

func goDetectDebugEnv() {
	go func() {
		select {
		case <-startCh:
		case <-startEnv:
		case <-time.After(time.Second * 5):
		}
		r := DetectDebugEnv()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resEnv <- r
		close(resEnv)
	}()
}

func goDetectParent() {
	go func() {
		select {
		case <-startCh:
		case <-startParent:
		case <-time.After(time.Second * 5):
		}
		r := DetectParent()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resParent <- r
		close(resParent)
	}()
}

func goDetectParent2() {
	go func() {
		select {
		case <-startCh:
		case <-startParent2:
		case <-time.After(time.Second * 5):
		}
		r := DetectParent2()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resParent2 <- r
		close(resParent2)
	}()
}

func goDetectInt3() {
	go func() {
		select {
		case <-startCh:
		case <-startInt3:
		case <-time.After(time.Second * 5):
		}
		r := DetectInt3()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resInt3 <- r
		close(resInt3)
	}()
}

func goDetectPtrace(runLast bool) {
	go func() {
		select {
		case <-startCh:
		case <-startPtrace:
		case <-time.After(time.Second * 5):
		}
		if runLast {
			<-startVDSO
			<-startNoASLR
			<-startHeap
			<-startParent
			<-startInt3
		}
		r := DetectPtrace()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resPtrace <- r
		close(resPtrace)
		resPtrace = nil
	}()
}

func goDetectTotal() {
	go func() {
		select {
		case <-startTotal:
		case <-time.After(time.Second * 5):
		}
		res = res || total > 0
	}()
}
