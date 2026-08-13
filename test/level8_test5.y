
    cns C_fQZjDl = 98
    v_oSatMO = unknown
    v_oSatMO = C_fQZjDl + 55
    if (v_oSatMO > 0) {
        v_oSatMO = v_oSatMO * 2
    } el {
        v_oSatMO = 0
    }
    pt(v_oSatMO)
    

    cns C_ZuRRyB = 19
    v_zKqjqe = unknown
    v_zKqjqe = C_ZuRRyB + 65
    if (v_zKqjqe > 0) {
        v_zKqjqe = v_zKqjqe * 2
    } el {
        v_zKqjqe = 0
    }
    pt(v_zKqjqe)
    

    cl Base_mmnmuZ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_QXoROG <- Base_mmnmuZ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_ahcJPQ = Child_QXoROG(73, 32)
    pt(obj_ahcJPQ.get_id())
    

    m_FstDWT = math.sqrt(37)
    h_jkILZk = hash.md5("test_oNqmaK")
    log.inf("hash:", h_jkILZk)
    pt(m_FstDWT)
    

    val_ZmpvLO = 65 % 3
    on (val_ZmpvLO) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_ZmpvLO)
    

    val_JQzzwP = 34 % 3
    on (val_JQzzwP) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_JQzzwP)
    

    use "dummy.y"
    
    e_yitNiD = enc.b64("hello drylang")
    pt(e_yitNiD)
    
    j_ZFmdLm = json(`{"test": 123}`)
    pt(j_ZFmdLm)
    
    // Test get() for type info
    val_zbbmsA = get("hello")
    pt(val_zbbmsA)
    

    try {
        throw_PigezT = unknown
        throw_PigezT()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_hwErQI(x) {
        rev x * 2
    }
    uni async_task_hwErQI(40)
    
    asn fn worker_AqrGJH(y) {
        pt("working on", y)
    }
    mul 2 worker_AqrGJH(66)
    
    awt
    

    sum_kssoXN = 0
    lp 5 {
        sum_kssoXN = sum_kssoXN + 1
        if (sum_kssoXN > 100) { done }
    }
    pt(sum_kssoXN)
    

    use "dummy.y"
    
    e_qGwfyK = enc.b64("hello drylang")
    pt(e_qGwfyK)
    
    j_zYDWUp = json(`{"test": 123}`)
    pt(j_zYDWUp)
    
    // Test get() for type info
    val_peYqsc = get("hello")
    pt(val_peYqsc)
    

    val_QRnSwR = 79 % 3
    on (val_QRnSwR) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_QRnSwR)
    

    try {
        throw_MDJovP = unknown
        throw_MDJovP()
    } err(e) {
        pt("caught error")
    }
    

    m_cfAbTD = math.sqrt(29)
    h_XEOpEq = hash.md5("test_xzgMTA")
    log.inf("hash:", h_XEOpEq)
    pt(m_cfAbTD)
    

    asn fn async_task_lAMQIc(x) {
        rev x * 2
    }
    uni async_task_lAMQIc(88)
    
    asn fn worker_SvIiKj(y) {
        pt("working on", y)
    }
    mul 2 worker_SvIiKj(92)
    
    awt
    
