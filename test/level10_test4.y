
    val_AGnbMi = 40 % 3
    on (val_AGnbMi) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_AGnbMi)
    

    asn fn async_task_nQLbRG(x) {
        rev x * 2
    }
    uni async_task_nQLbRG(98)
    
    asn fn worker_GBTFTJ(y) {
        pt("working on", y)
    }
    mul 2 worker_GBTFTJ(66)
    
    awt
    

    try {
        throw_ODkBHh = unknown
        throw_ODkBHh()
    } err(e) {
        pt("caught error")
    }
    

    arr_eLvTnc = [38, 47+1, 2+2]
    map_IrumAo = {"a": arr_eLvTnc[0], "b": arr_eLvTnc[1]}
    pt(map_IrumAo["a"])
    

    sum_draKYR = 0
    lp 8 {
        sum_draKYR = sum_draKYR + 1
        if (sum_draKYR > 100) { done }
    }
    pt(sum_draKYR)
    

    use "dummy.y"
    
    e_PfVPUV = enc.b64("hello drylang")
    pt(e_PfVPUV)
    
    j_sJFHpI = json(`{"test": 123}`)
    pt(j_sJFHpI)
    
    // Test get() for type info
    val_rfekOp = get("hello")
    pt(val_rfekOp)
    

    m_Dihbsx = math.sqrt(35)
    h_FHMNfe = hash.md5("test_qJwFBd")
    log.inf("hash:", h_FHMNfe)
    pt(m_Dihbsx)
    

    try {
        throw_YVAEol = unknown
        throw_YVAEol()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_AecHcg(x) {
        rev x * 2
    }
    uni async_task_AecHcg(42)
    
    asn fn worker_SPsMlY(y) {
        pt("working on", y)
    }
    mul 2 worker_SPsMlY(44)
    
    awt
    

    m_RXLJsU = math.sqrt(24)
    h_hbhgNJ = hash.md5("test_hmjxPD")
    log.inf("hash:", h_hbhgNJ)
    pt(m_RXLJsU)
    

    cl Base_rSLeVJ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_dnYxeW <- Base_rSLeVJ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_tzBKRS = Child_dnYxeW(43, 30)
    pt(obj_tzBKRS.get_id())
    

    asn fn async_task_EFiYex(x) {
        rev x * 2
    }
    uni async_task_EFiYex(9)
    
    asn fn worker_iUxEBf(y) {
        pt("working on", y)
    }
    mul 2 worker_iUxEBf(43)
    
    awt
    

    val_mrYZNp = 91 % 3
    on (val_mrYZNp) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_mrYZNp)
    

    cns C_ObknBU = 8
    v_kqUAeq = unknown
    v_kqUAeq = C_ObknBU + 9
    if (v_kqUAeq > 0) {
        v_kqUAeq = v_kqUAeq * 2
    } el {
        v_kqUAeq = 0
    }
    pt(v_kqUAeq)
    

    val_kQptyi = 88 % 3
    on (val_kQptyi) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_kQptyi)
    

    asn fn async_task_bOgGTA(x) {
        rev x * 2
    }
    uni async_task_bOgGTA(32)
    
    asn fn worker_kRRNEq(y) {
        pt("working on", y)
    }
    mul 2 worker_kRRNEq(1)
    
    awt
    

    m_lMCyKL = math.sqrt(70)
    h_ZCKUhw = hash.md5("test_EGAbZb")
    log.inf("hash:", h_ZCKUhw)
    pt(m_lMCyKL)
    
