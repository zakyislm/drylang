
    m_NJwxQk = math.sqrt(66)
    h_IRaBKI = hash.md5("test_KMLmyj")
    log.inf("hash:", h_IRaBKI)
    pt(m_NJwxQk)
    

    use "dummy.y"
    
    e_KtTmKl = enc.b64("hello drylang")
    pt(e_KtTmKl)
    
    j_GygWEd = json(`{"test": 123}`)
    pt(j_GygWEd)
    
    // Test get() for type info
    val_eRkEeK = get("hello")
    pt(val_eRkEeK)
    

    m_cwbgav = math.sqrt(22)
    h_BhIWVw = hash.md5("test_kGyKgN")
    log.inf("hash:", h_BhIWVw)
    pt(m_cwbgav)
    

    val_cqgsDQ = 67 % 3
    on (val_cqgsDQ) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_cqgsDQ)
    

    cl Base_intCMl {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_fMDmYG <- Base_intCMl {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_hGmLtf = Child_fMDmYG(23, 39)
    pt(obj_hGmLtf.get_id())
    

    arr_fnJsjm = [3, 72+1, 80+2]
    map_qgcjkJ = {"a": arr_fnJsjm[0], "b": arr_fnJsjm[1]}
    pt(map_qgcjkJ["a"])
    

    sum_KeIyub = 0
    lp 11 {
        sum_KeIyub = sum_KeIyub + 1
        if (sum_KeIyub > 100) { done }
    }
    pt(sum_KeIyub)
    

    fn calc_pASfNN(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_QfOcto = calc_pASfNN(13, 75+1)
    pt(res_QfOcto)
    

    cl Base_ypZKQQ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_ElNGgE <- Base_ypZKQQ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_nqjcHh = Child_ElNGgE(59, 19)
    pt(obj_nqjcHh.get_id())
    

    cns C_IkpqgU = 44
    v_vFjZkt = unknown
    v_vFjZkt = C_IkpqgU + 62
    if (v_vFjZkt > 0) {
        v_vFjZkt = v_vFjZkt * 2
    } el {
        v_vFjZkt = 0
    }
    pt(v_vFjZkt)
    

    val_rOCcbI = 91 % 3
    on (val_rOCcbI) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_rOCcbI)
    

    val_oUBZMH = 62 % 3
    on (val_oUBZMH) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_oUBZMH)
    

    val_OnbrGS = 54 % 3
    on (val_OnbrGS) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_OnbrGS)
    

    arr_CCQuuM = [29, 46+1, 1+2]
    map_Gyjkso = {"a": arr_CCQuuM[0], "b": arr_CCQuuM[1]}
    pt(map_Gyjkso["a"])
    

    sum_TCtOBR = 0
    lp 7 {
        sum_TCtOBR = sum_TCtOBR + 1
        if (sum_TCtOBR > 100) { done }
    }
    pt(sum_TCtOBR)
    

    sum_txxLSN = 0
    lp 7 {
        sum_txxLSN = sum_txxLSN + 1
        if (sum_txxLSN > 100) { done }
    }
    pt(sum_txxLSN)
    

    sum_AxKkVU = 0
    lp 8 {
        sum_AxKkVU = sum_AxKkVU + 1
        if (sum_AxKkVU > 100) { done }
    }
    pt(sum_AxKkVU)
    
