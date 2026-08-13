
    m_QxgEjf = math.sqrt(71)
    h_mXiZCe = hash.md5("test_IlLWKV")
    log.inf("hash:", h_mXiZCe)
    pt(m_QxgEjf)
    

    cl Base_MHMwZV {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_XHaFRn <- Base_MHMwZV {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_yWLxPE = Child_XHaFRn(19, 26)
    pt(obj_yWLxPE.get_id())
    

    cns C_MvvvbB = 4
    v_uZtBWU = unknown
    v_uZtBWU = C_MvvvbB + 2
    if (v_uZtBWU > 0) {
        v_uZtBWU = v_uZtBWU * 2
    } el {
        v_uZtBWU = 0
    }
    pt(v_uZtBWU)
    

    val_wUwDvq = 77 % 3
    on (val_wUwDvq) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_wUwDvq)
    

    try {
        throw_BxzVfj = unknown
        throw_BxzVfj()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_CCCygv(x) {
        rev x * 2
    }
    uni async_task_CCCygv(38)
    
    asn fn worker_bkNrax(y) {
        pt("working on", y)
    }
    mul 2 worker_bkNrax(55)
    
    awt
    

    cns C_oErsHj = 72
    v_vJiLqu = unknown
    v_vJiLqu = C_oErsHj + 78
    if (v_vJiLqu > 0) {
        v_vJiLqu = v_vJiLqu * 2
    } el {
        v_vJiLqu = 0
    }
    pt(v_vJiLqu)
    

    arr_DECIGk = [69, 2+1, 9+2]
    map_ReicLE = {"a": arr_DECIGk[0], "b": arr_DECIGk[1]}
    pt(map_ReicLE["a"])
    

    sum_enRhEX = 0
    lp 9 {
        sum_enRhEX = sum_enRhEX + 1
        if (sum_enRhEX > 100) { done }
    }
    pt(sum_enRhEX)
    

    arr_YhqYFk = [13, 36+1, 74+2]
    map_icoFoq = {"a": arr_YhqYFk[0], "b": arr_YhqYFk[1]}
    pt(map_icoFoq["a"])
    

    fn calc_LuLMMd(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_fmNjhu = calc_LuLMMd(100, 28+1)
    pt(res_fmNjhu)
    

    arr_zXhKju = [25, 53+1, 60+2]
    map_rMenGF = {"a": arr_zXhKju[0], "b": arr_zXhKju[1]}
    pt(map_rMenGF["a"])
    

    m_MRzJld = math.sqrt(16)
    h_BveupC = hash.md5("test_nXkBhd")
    log.inf("hash:", h_BveupC)
    pt(m_MRzJld)
    

    use "dummy.y"
    
    e_SAQhgr = enc.b64("hello drylang")
    pt(e_SAQhgr)
    
    j_NMuYHx = json(`{"test": 123}`)
    pt(j_NMuYHx)
    
    // Test get() for type info
    val_IOTwSL = get("hello")
    pt(val_IOTwSL)
    
