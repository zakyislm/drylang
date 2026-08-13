
    m_cWxaVH = math.sqrt(15)
    h_XSCajB = hash.md5("test_mPUuSg")
    log.inf("hash:", h_XSCajB)
    pt(m_cWxaVH)
    

    cl Base_hAesRE {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_xoIxfs <- Base_hAesRE {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_ikxhzW = Child_xoIxfs(60, 33)
    pt(obj_ikxhzW.get_id())
    

    use "dummy.y"
    
    e_MFYAqO = enc.b64("hello drylang")
    pt(e_MFYAqO)
    
    j_WpaZCo = json(`{"test": 123}`)
    pt(j_WpaZCo)
    
    // Test get() for type info
    val_XmrGAW = get("hello")
    pt(val_XmrGAW)
    

    fn calc_GoLgNT(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_DrfSya = calc_GoLgNT(75, 15+1)
    pt(res_DrfSya)
    

    val_wsnFcO = 27 % 3
    on (val_wsnFcO) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_wsnFcO)
    

    arr_JKTAkB = [88, 61+1, 50+2]
    map_HdgaXE = {"a": arr_JKTAkB[0], "b": arr_JKTAkB[1]}
    pt(map_HdgaXE["a"])
    

    fn calc_UGzbWF(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_HUUFQV = calc_UGzbWF(26, 16+1)
    pt(res_HUUFQV)
    

    cns C_KMVnof = 52
    v_VZoZlc = unknown
    v_VZoZlc = C_KMVnof + 37
    if (v_VZoZlc > 0) {
        v_VZoZlc = v_VZoZlc * 2
    } el {
        v_VZoZlc = 0
    }
    pt(v_VZoZlc)
    

    arr_UAdpcI = [22, 41+1, 26+2]
    map_BrBXDo = {"a": arr_UAdpcI[0], "b": arr_UAdpcI[1]}
    pt(map_BrBXDo["a"])
    

    cns C_QqOXOO = 13
    v_oGQcJh = unknown
    v_oGQcJh = C_QqOXOO + 68
    if (v_oGQcJh > 0) {
        v_oGQcJh = v_oGQcJh * 2
    } el {
        v_oGQcJh = 0
    }
    pt(v_oGQcJh)
    

    asn fn async_task_OwZuWS(x) {
        rev x * 2
    }
    uni async_task_OwZuWS(36)
    
    asn fn worker_ODzEAi(y) {
        pt("working on", y)
    }
    mul 2 worker_ODzEAi(55)
    
    awt
    

    fn calc_pCQqyo(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_TsNVDd = calc_pCQqyo(15, 97+1)
    pt(res_TsNVDd)
    

    m_pAiCuP = math.sqrt(68)
    h_pHyosL = hash.md5("test_clwwKd")
    log.inf("hash:", h_pHyosL)
    pt(m_pAiCuP)
    

    val_fWLsPO = 33 % 3
    on (val_fWLsPO) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_fWLsPO)
    

    cl Base_TpIagg {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_urAkbI <- Base_TpIagg {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_IsejNk = Child_urAkbI(78, 70)
    pt(obj_IsejNk.get_id())
    

    m_wbCxNw = math.sqrt(19)
    h_YGIUyr = hash.md5("test_sFLhax")
    log.inf("hash:", h_YGIUyr)
    pt(m_wbCxNw)
    

    try {
        throw_fiUoAo = unknown
        throw_fiUoAo()
    } err(e) {
        pt("caught error")
    }
    
