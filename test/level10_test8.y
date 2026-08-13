
    cns C_KPgkjR = 36
    v_gKIiSI = unknown
    v_gKIiSI = C_KPgkjR + 12
    if (v_gKIiSI > 0) {
        v_gKIiSI = v_gKIiSI * 2
    } el {
        v_gKIiSI = 0
    }
    pt(v_gKIiSI)
    

    cns C_hrChXf = 86
    v_WKWvdr = unknown
    v_WKWvdr = C_hrChXf + 59
    if (v_WKWvdr > 0) {
        v_WKWvdr = v_WKWvdr * 2
    } el {
        v_WKWvdr = 0
    }
    pt(v_WKWvdr)
    

    arr_kcjuUf = [13, 37+1, 85+2]
    map_AbwvjC = {"a": arr_kcjuUf[0], "b": arr_kcjuUf[1]}
    pt(map_AbwvjC["a"])
    

    val_RrsLZg = 29 % 3
    on (val_RrsLZg) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_RrsLZg)
    

    sum_kaCAJn = 0
    lp 11 {
        sum_kaCAJn = sum_kaCAJn + 1
        if (sum_kaCAJn > 100) { done }
    }
    pt(sum_kaCAJn)
    

    try {
        throw_JmcrQL = unknown
        throw_JmcrQL()
    } err(e) {
        pt("caught error")
    }
    

    arr_cWcEgx = [97, 14+1, 9+2]
    map_dFqHie = {"a": arr_cWcEgx[0], "b": arr_cWcEgx[1]}
    pt(map_dFqHie["a"])
    

    asn fn async_task_iuHaLG(x) {
        rev x * 2
    }
    uni async_task_iuHaLG(64)
    
    asn fn worker_ZOYrwM(y) {
        pt("working on", y)
    }
    mul 2 worker_ZOYrwM(60)
    
    awt
    

    val_oJzHYW = 24 % 3
    on (val_oJzHYW) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_oJzHYW)
    

    use "dummy.y"
    
    e_ImwNDC = enc.b64("hello drylang")
    pt(e_ImwNDC)
    
    j_Naafaq = json(`{"test": 123}`)
    pt(j_Naafaq)
    
    // Test get() for type info
    val_pHrbHH = get("hello")
    pt(val_pHrbHH)
    

    cl Base_IFOHXB {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_ivbWBo <- Base_IFOHXB {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_AqeCzM = Child_ivbWBo(4, 26)
    pt(obj_AqeCzM.get_id())
    

    cns C_gJpWwM = 19
    v_tfNSpT = unknown
    v_tfNSpT = C_gJpWwM + 87
    if (v_tfNSpT > 0) {
        v_tfNSpT = v_tfNSpT * 2
    } el {
        v_tfNSpT = 0
    }
    pt(v_tfNSpT)
    

    cns C_sidQmw = 51
    v_nxGrtF = unknown
    v_nxGrtF = C_sidQmw + 2
    if (v_nxGrtF > 0) {
        v_nxGrtF = v_nxGrtF * 2
    } el {
        v_nxGrtF = 0
    }
    pt(v_nxGrtF)
    

    arr_OlYsXI = [40, 33+1, 90+2]
    map_YDKkAj = {"a": arr_OlYsXI[0], "b": arr_OlYsXI[1]}
    pt(map_YDKkAj["a"])
    

    arr_uLunXR = [38, 20+1, 68+2]
    map_pjLDAR = {"a": arr_uLunXR[0], "b": arr_uLunXR[1]}
    pt(map_pjLDAR["a"])
    

    use "dummy.y"
    
    e_LgTvtT = enc.b64("hello drylang")
    pt(e_LgTvtT)
    
    j_PJkLWQ = json(`{"test": 123}`)
    pt(j_PJkLWQ)
    
    // Test get() for type info
    val_AsLChw = get("hello")
    pt(val_AsLChw)
    

    cns C_UkqyMd = 25
    v_GSJcPL = unknown
    v_GSJcPL = C_UkqyMd + 15
    if (v_GSJcPL > 0) {
        v_GSJcPL = v_GSJcPL * 2
    } el {
        v_GSJcPL = 0
    }
    pt(v_GSJcPL)
    
