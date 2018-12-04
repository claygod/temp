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
