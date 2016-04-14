GOMON is a package that you can drop into your go application for monitoring common runtime stats.

# Design Considerations:

 - Log runtime stats
 - Take logging/reporting adaptors so that these stats can be reported as metrics

Some stat collection will cause a performance hit (like GC etc)

# Usage:
```golang
import "github.com/urjitbhatia/gomon"
...

func monitorSystem() {
	mon := gomon.New()
	mon.CaptureGoRoutineStats()
	mon.Start()
}
```