
    asn fn async_task_GbhHVI(x) {
        rev x * 2
    }
    uni async_task_GbhHVI(66)
    
    asn fn worker_hScMGp(y) {
        pt("working on", y)
    }
    mul 2 worker_hScMGp(81)
    
    awt
    

    asn fn async_task_qWpKny(x) {
        rev x * 2
    }
    uni async_task_qWpKny(90)
    
    asn fn worker_bkGrfq(y) {
        pt("working on", y)
    }
    mul 2 worker_bkGrfq(75)
    
    awt
    

    sum_GvNEjI = 0
    lp 14 {
        sum_GvNEjI = sum_GvNEjI + 1
        if (sum_GvNEjI > 100) { done }
    }
    pt(sum_GvNEjI)
    

    asn fn async_task_NRMHwb(x) {
        rev x * 2
    }
    uni async_task_NRMHwb(23)
    
    asn fn worker_bOBnGv(y) {
        pt("working on", y)
    }
    mul 2 worker_bOBnGv(53)
    
    awt
    

    sum_CVHbNJ = 0
    lp 10 {
        sum_CVHbNJ = sum_CVHbNJ + 1
        if (sum_CVHbNJ > 100) { done }
    }
    pt(sum_CVHbNJ)
    

    cl Base_lHUZtk {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_WyfaTE <- Base_lHUZtk {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_dcZRqq = Child_WyfaTE(1, 65)
    pt(obj_dcZRqq.get_id())
    
