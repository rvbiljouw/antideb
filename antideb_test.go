package antideb

import (
	"fmt"
	"testing"
)

// TestDetect - detect a debugger / anti-debugger
func TestDetect(t *testing.T) {
	fmt.Printf("DetectDebugEnv: %v\n", DetectDebugEnv())
	fmt.Printf("DetectParent: %v\n", DetectParent())
	fmt.Printf("DetectParent2: %v\n", DetectParent2())
	fmt.Printf("DetectPtrace: %v\n", DetectPtrace())
	fmt.Printf("Detect: %v\n", DetectAll(false))
}
