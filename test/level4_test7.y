
    cl Base_jmLEzj {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_wQeXyW <- Base_jmLEzj {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_qBuZpF = Child_wQeXyW(21, 19)
    pt(obj_qBuZpF.get_id())
    

    asn fn async_task_Lgfbsk(x) {
        rev x * 2
    }
    uni async_task_Lgfbsk(46)
    
    asn fn worker_pvTyEl(y) {
        pt("working on", y)
    }
    mul 2 worker_pvTyEl(86)
    
    awt
    

    cns C_UgcBlW = 86
    v_oCqurB = unknown
    v_oCqurB = C_UgcBlW + 67
    if (v_oCqurB > 0) {
        v_oCqurB = v_oCqurB * 2
    } el {
        v_oCqurB = 0
    }
    pt(v_oCqurB)
    

    sum_GtoCvM = 0
    lp 12 {
        sum_GtoCvM = sum_GtoCvM + 1
        if (sum_GtoCvM > 100) { done }
    }
    pt(sum_GtoCvM)
    

    try {
        throw_orVwRK = unknown
        throw_orVwRK()
    } err(e) {
        pt("caught error")
    }
    

    m_uvzAEH = math.sqrt(45)
    h_JqOdlf = hash.md5("test_cePPnM")
    log.inf("hash:", h_JqOdlf)
    pt(m_uvzAEH)
    
