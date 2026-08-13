
    m_suUarA = math.sqrt(91)
    h_gPdZTy = hash.md5("test_AhcTrS")
    log.inf("hash:", h_gPdZTy)
    pt(m_suUarA)
    

    sum_DklvJR = 0
    lp 9 {
        sum_DklvJR = sum_DklvJR + 1
        if (sum_DklvJR > 100) { done }
    }
    pt(sum_DklvJR)
    

    asn fn async_task_opoIVI(x) {
        rev x * 2
    }
    uni async_task_opoIVI(41)
    
    asn fn worker_FxSBbW(y) {
        pt("working on", y)
    }
    mul 2 worker_FxSBbW(76)
    
    awt
    

    cns C_kFuozX = 100
    v_RiCDgV = unknown
    v_RiCDgV = C_kFuozX + 34
    if (v_RiCDgV > 0) {
        v_RiCDgV = v_RiCDgV * 2
    } el {
        v_RiCDgV = 0
    }
    pt(v_RiCDgV)
    

    fn calc_kAOsKe(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_tazlOR = calc_kAOsKe(28, 65+1)
    pt(res_tazlOR)
    

    use "dummy.y"
    
    e_UirfCR = enc.b64("hello drylang")
    pt(e_UirfCR)
    
    j_YiNsMX = json(`{"test": 123}`)
    pt(j_YiNsMX)
    
    // Test get() for type info
    val_MywFvn = get("hello")
    pt(val_MywFvn)
    

    cns C_quTAlz = 95
    v_BIsHQn = unknown
    v_BIsHQn = C_quTAlz + 25
    if (v_BIsHQn > 0) {
        v_BIsHQn = v_BIsHQn * 2
    } el {
        v_BIsHQn = 0
    }
    pt(v_BIsHQn)
    

    use "dummy.y"
    
    e_yRuYDP = enc.b64("hello drylang")
    pt(e_yRuYDP)
    
    j_semhaK = json(`{"test": 123}`)
    pt(j_semhaK)
    
    // Test get() for type info
    val_nsXuHx = get("hello")
    pt(val_nsXuHx)
    

    try {
        throw_jOAbYQ = unknown
        throw_jOAbYQ()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_KHqkPW = enc.b64("hello drylang")
    pt(e_KHqkPW)
    
    j_LnoGCL = json(`{"test": 123}`)
    pt(j_LnoGCL)
    
    // Test get() for type info
    val_YOdmnP = get("hello")
    pt(val_YOdmnP)
    

    arr_gFrvUC = [20, 73+1, 66+2]
    map_kVyzVd = {"a": arr_gFrvUC[0], "b": arr_gFrvUC[1]}
    pt(map_kVyzVd["a"])
    

    cl Base_CpXZrq {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_kmLQfH <- Base_CpXZrq {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_HjJQjp = Child_kmLQfH(60, 16)
    pt(obj_HjJQjp.get_id())
    

    cl Base_RlrtOD {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_ImmjKE <- Base_RlrtOD {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_JAxYOr = Child_ImmjKE(86, 9)
    pt(obj_JAxYOr.get_id())
    

    use "dummy.y"
    
    e_bVcReS = enc.b64("hello drylang")
    pt(e_bVcReS)
    
    j_zxEfEj = json(`{"test": 123}`)
    pt(j_zxEfEj)
    
    // Test get() for type info
    val_FzHhRQ = get("hello")
    pt(val_FzHhRQ)
    

    m_wUBLzJ = math.sqrt(60)
    h_hitHJT = hash.md5("test_KGxCmH")
    log.inf("hash:", h_hitHJT)
    pt(m_wUBLzJ)
    

    fn calc_HJJaqL(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_JFRmSf = calc_HJJaqL(44, 93+1)
    pt(res_JFRmSf)
    

    cl Base_QNHGEr {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_zfkgHN <- Base_QNHGEr {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_FmqBLM = Child_zfkgHN(47, 76)
    pt(obj_FmqBLM.get_id())
    
