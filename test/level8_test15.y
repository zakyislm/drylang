
    asn fn async_task_OvZlRU(x) {
        rev x * 2
    }
    uni async_task_OvZlRU(77)
    
    asn fn worker_TszREM(y) {
        pt("working on", y)
    }
    mul 2 worker_TszREM(23)
    
    awt
    

    cl Base_yzcLqt {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_kkJvWt <- Base_yzcLqt {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_afDSCU = Child_kkJvWt(5, 66)
    pt(obj_afDSCU.get_id())
    

    sum_IxsgoN = 0
    lp 8 {
        sum_IxsgoN = sum_IxsgoN + 1
        if (sum_IxsgoN > 100) { done }
    }
    pt(sum_IxsgoN)
    

    use "dummy.y"
    
    e_LfFEIX = enc.b64("hello drylang")
    pt(e_LfFEIX)
    
    j_NcbhHG = json(`{"test": 123}`)
    pt(j_NcbhHG)
    
    // Test get() for type info
    val_oQOmDq = get("hello")
    pt(val_oQOmDq)
    

    use "dummy.y"
    
    e_BkaGeo = enc.b64("hello drylang")
    pt(e_BkaGeo)
    
    j_TsNAFE = json(`{"test": 123}`)
    pt(j_TsNAFE)
    
    // Test get() for type info
    val_euzxFu = get("hello")
    pt(val_euzxFu)
    

    fn calc_kuWqAe(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_bjQwYF = calc_kuWqAe(79, 36+1)
    pt(res_bjQwYF)
    

    fn calc_IvNfUf(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_kDBEFD = calc_IvNfUf(9, 9+1)
    pt(res_kDBEFD)
    

    use "dummy.y"
    
    e_DZDSGe = enc.b64("hello drylang")
    pt(e_DZDSGe)
    
    j_qlTbgR = json(`{"test": 123}`)
    pt(j_qlTbgR)
    
    // Test get() for type info
    val_MKDABL = get("hello")
    pt(val_MKDABL)
    

    m_UbcXPd = math.sqrt(59)
    h_aiuFoc = hash.md5("test_GNnltm")
    log.inf("hash:", h_aiuFoc)
    pt(m_UbcXPd)
    

    cns C_PyHIMB = 30
    v_zEhbcE = unknown
    v_zEhbcE = C_PyHIMB + 100
    if (v_zEhbcE > 0) {
        v_zEhbcE = v_zEhbcE * 2
    } el {
        v_zEhbcE = 0
    }
    pt(v_zEhbcE)
    

    fn calc_WIkBHD(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_VXHSYQ = calc_WIkBHD(63, 31+1)
    pt(res_VXHSYQ)
    

    asn fn async_task_ABufoD(x) {
        rev x * 2
    }
    uni async_task_ABufoD(91)
    
    asn fn worker_rkxbHU(y) {
        pt("working on", y)
    }
    mul 2 worker_rkxbHU(86)
    
    awt
    

    val_Wxflst = 23 % 3
    on (val_Wxflst) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_Wxflst)
    

    try {
        throw_dqotbG = unknown
        throw_dqotbG()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_tsGLln {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_tAOfBL <- Base_tsGLln {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_gceugU = Child_tAOfBL(9, 11)
    pt(obj_gceugU.get_id())
    
