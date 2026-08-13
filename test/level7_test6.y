
    cl Base_haohMX {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_vBpMhE <- Base_haohMX {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_XLUfhq = Child_vBpMhE(37, 43)
    pt(obj_XLUfhq.get_id())
    

    try {
        throw_bMvKsK = unknown
        throw_bMvKsK()
    } err(e) {
        pt("caught error")
    }
    

    val_HrZRoW = 22 % 3
    on (val_HrZRoW) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_HrZRoW)
    

    m_ucYVmW = math.sqrt(94)
    h_fPlPMi = hash.md5("test_pAsGmL")
    log.inf("hash:", h_fPlPMi)
    pt(m_ucYVmW)
    

    try {
        throw_fHolyy = unknown
        throw_fHolyy()
    } err(e) {
        pt("caught error")
    }
    

    m_gUOtEa = math.sqrt(20)
    h_nAVtJP = hash.md5("test_murVPV")
    log.inf("hash:", h_nAVtJP)
    pt(m_gUOtEa)
    

    sum_QBXqDA = 0
    lp 10 {
        sum_QBXqDA = sum_QBXqDA + 1
        if (sum_QBXqDA > 100) { done }
    }
    pt(sum_QBXqDA)
    

    fn calc_fRdvEB(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_zJtObI = calc_fRdvEB(78, 53+1)
    pt(res_zJtObI)
    

    m_kLHfhb = math.sqrt(60)
    h_QurIzI = hash.md5("test_HfKuaT")
    log.inf("hash:", h_QurIzI)
    pt(m_kLHfhb)
    

    val_okFkps = 97 % 3
    on (val_okFkps) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_okFkps)
    

    m_vtppii = math.sqrt(17)
    h_IZmIeF = hash.md5("test_WHCSew")
    log.inf("hash:", h_IZmIeF)
    pt(m_vtppii)
    

    arr_wOineH = [9, 75+1, 95+2]
    map_XNCygQ = {"a": arr_wOineH[0], "b": arr_wOineH[1]}
    pt(map_XNCygQ["a"])
    

    cns C_kZQEza = 42
    v_SwrEwk = unknown
    v_SwrEwk = C_kZQEza + 23
    if (v_SwrEwk > 0) {
        v_SwrEwk = v_SwrEwk * 2
    } el {
        v_SwrEwk = 0
    }
    pt(v_SwrEwk)
    

    fn calc_yKnVgb(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_TnVrMR = calc_yKnVgb(3, 36+1)
    pt(res_TnVrMR)
    
