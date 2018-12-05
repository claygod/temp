	efacez := *(**hchan)(unsafe.Pointer(&ch))
	fmt.Println("=++++=", efacez)
	//fmt.Println("=++++=", efacez.ch)
	fmt.Println("=++++=", efacez.dataqsiz)
	fmt.Println("=++++=", efacez.qcount)
	fmt.Println("=++++=", efacez.dataqsiz)


func chanbufr(c *hchan, i uint) unsafe.Pointer {
	for u := i; u > 0; u-- {
		p1 := (*unsafe.Pointer)(unsafe.Pointer(uintptr(c.buf) + uintptr(u-1)*uintptr(c.elemsize)))
		p2 := (*unsafe.Pointer)(unsafe.Pointer(uintptr(c.buf) + uintptr(u)*uintptr(c.elemsize)))
		*p2 = *p1
	}
	return c.buf
}
// -----------------------------------------------------

package rchan

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestRchan1(t *testing.T) {
	ch := make(chan int64, 7)
	ag := runtime.NewChanAgent(ch)
	ch <- 771
	fmt.Println("-- prior ", ag.GetPriority())
	fmt.Println("-- count ", ag.GetCount())
	fmt.Println("-- size ", ag.GetSize())
	fmt.Println("")
	//ag.SetPriority()
	//ag.SetClean()
	ch <- 772
	fmt.Println("-- prior ", ag.GetPriority())
	fmt.Println("-- count ", ag.GetCount())
	fmt.Println("-- size ", ag.GetSize())
	fmt.Println("")
	ag.SetPriority()
	time.Sleep(1 * time.Millisecond)
	ch <- 773

	fmt.Println("-- prior ", ag.GetPriority())
	fmt.Println("-- count ", ag.GetCount())
	fmt.Println("-- size ", ag.GetSize())

	ch <- 774

	out := <-ch
	fmt.Println(out)
	time.Sleep(1 * time.Millisecond)
	out = <-ch
	fmt.Println(out)
	time.Sleep(1 * time.Millisecond)
	out = <-ch
	fmt.Println(out)
	time.Sleep(1 * time.Millisecond)
	out = <-ch
	fmt.Println(out)
}

// ------------------------------------------------

// chanbuf(c, i) is pointer to the i'th slot in the buffer.
func chanbuf(c *hchan, i uint) unsafe.Pointer {

	if atomic.Cas(c.flagPriority, 1, 0) { // atomic.Cas(c.flagPriority, 1, 0)  atomic.Load(c.flagPriority)
		return chanbufPr(c, i)
	}
	return add(c.buf, uintptr(i)*uintptr(c.elemsize))
}

func chanbufPr(c *hchan, i uint) unsafe.Pointer {
	for u := i; u > 0; u-- {
		p1 := (*unsafe.Pointer)(unsafe.Pointer(uintptr(c.buf) + uintptr(u-1)*uintptr(c.elemsize)))
		p2 := (*unsafe.Pointer)(unsafe.Pointer(uintptr(c.buf) + uintptr(u)*uintptr(c.elemsize)))
		*p2 = *p1
	}
	return c.buf
}

func addr(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

// -------------------------------------------------

func chanClean(c *hchan) {
	if atomic.Cas(c.flagClean, 1, 0) {
		c.qcount = 0 //c.dataqsiz // TODO:
	}
}

// --------------------------------


