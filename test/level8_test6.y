
    cns C_iRavVz = 70
    v_RFWDCQ = unknown
    v_RFWDCQ = C_iRavVz + 99
    if (v_RFWDCQ > 0) {
        v_RFWDCQ = v_RFWDCQ * 2
    } el {
        v_RFWDCQ = 0
    }
    pt(v_RFWDCQ)
    

    sum_FvVVtS = 0
    lp 15 {
        sum_FvVVtS = sum_FvVVtS + 1
        if (sum_FvVVtS > 100) { done }
    }
    pt(sum_FvVVtS)
    

    use "dummy.y"
    
    e_fYWoDJ = enc.b64("hello drylang")
    pt(e_fYWoDJ)
    
    j_kjBodg = json(`{"test": 123}`)
    pt(j_kjBodg)
    
    // Test get() for type info
    val_qbrCmC = get("hello")
    pt(val_qbrCmC)
    

    asn fn async_task_nCCAAa(x) {
        rev x * 2
    }
    uni async_task_nCCAAa(18)
    
    asn fn worker_KcdZAq(y) {
        pt("working on", y)
    }
    mul 2 worker_KcdZAq(88)
    
    awt
    

    fn calc_tPWwri(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_ewsMQi = calc_tPWwri(12, 3+1)
    pt(res_ewsMQi)
    

    m_YkiJOL = math.sqrt(38)
    h_eYYbWL = hash.md5("test_HjAftO")
    log.inf("hash:", h_eYYbWL)
    pt(m_YkiJOL)
    

    try {
        throw_PQWnrx = unknown
        throw_PQWnrx()
    } err(e) {
        pt("caught error")
    }
    

    arr_SWlzwO = [37, 64+1, 3+2]
    map_PixQro = {"a": arr_SWlzwO[0], "b": arr_SWlzwO[1]}
    pt(map_PixQro["a"])
    

    try {
        throw_sQPsnM = unknown
        throw_sQPsnM()
    } err(e) {
        pt("caught error")
    }
    

    fn calc_WivqYd(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_noAFSb = calc_WivqYd(6, 95+1)
    pt(res_noAFSb)
    

    fn calc_mEidlW(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_hpsLmD = calc_mEidlW(99, 71+1)
    pt(res_hpsLmD)
    

    cl Base_jLktSK {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_qfHNBP <- Base_jLktSK {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_WwqXtj = Child_qfHNBP(91, 91)
    pt(obj_WwqXtj.get_id())
    

    fn calc_xIaATG(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_XPxgXa = calc_xIaATG(64, 63+1)
    pt(res_XPxgXa)
    

    fn calc_kHhsYY(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_EeMzfc = calc_kHhsYY(53, 45+1)
    pt(res_EeMzfc)
    

    sum_EMrXPp = 0
    lp 11 {
        sum_EMrXPp = sum_EMrXPp + 1
        if (sum_EMrXPp > 100) { done }
    }
    pt(sum_EMrXPp)
    
