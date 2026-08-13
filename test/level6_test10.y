
    asn fn async_task_Dcfygy(x) {
        rev x * 2
    }
    uni async_task_Dcfygy(72)
    
    asn fn worker_RYArdd(y) {
        pt("working on", y)
    }
    mul 2 worker_RYArdd(95)
    
    awt
    

    try {
        throw_ktPpIn = unknown
        throw_ktPpIn()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_UMcQyX(x) {
        rev x * 2
    }
    uni async_task_UMcQyX(100)
    
    asn fn worker_baEPKz(y) {
        pt("working on", y)
    }
    mul 2 worker_baEPKz(78)
    
    awt
    

    cl Base_bVDjnW {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_tRsVoY <- Base_bVDjnW {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_rPCmen = Child_tRsVoY(39, 85)
    pt(obj_rPCmen.get_id())
    

    sum_BvUlfs = 0
    lp 5 {
        sum_BvUlfs = sum_BvUlfs + 1
        if (sum_BvUlfs > 100) { done }
    }
    pt(sum_BvUlfs)
    

    use "dummy.y"
    
    e_SypiYw = enc.b64("hello drylang")
    pt(e_SypiYw)
    
    j_XPuvNF = json(`{"test": 123}`)
    pt(j_XPuvNF)
    
    // Test get() for type info
    val_yTCsTP = get("hello")
    pt(val_yTCsTP)
    

    cl Base_aQbOIm {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_pXvHjb <- Base_aQbOIm {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_GDfItf = Child_pXvHjb(26, 10)
    pt(obj_GDfItf.get_id())
    

    sum_AITuYG = 0
    lp 10 {
        sum_AITuYG = sum_AITuYG + 1
        if (sum_AITuYG > 100) { done }
    }
    pt(sum_AITuYG)
    
