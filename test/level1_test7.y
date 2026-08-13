
    sum_wOxhHg = 0
    lp 11 {
        sum_wOxhHg = sum_wOxhHg + 1
        if (sum_wOxhHg > 100) { done }
    }
    pt(sum_wOxhHg)
    

    cl Base_oHkZXM {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_DLARTC <- Base_oHkZXM {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_tXMPOy = Child_DLARTC(10, 83)
    pt(obj_tXMPOy.get_id())
    

    asn fn async_task_GWnKkB(x) {
        rev x * 2
    }
    uni async_task_GWnKkB(80)
    
    asn fn worker_VcStDN(y) {
        pt("working on", y)
    }
    mul 2 worker_VcStDN(62)
    
    awt
    
