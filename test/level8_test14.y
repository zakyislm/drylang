
    cl Base_BQPPWI {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_veGQlA <- Base_BQPPWI {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_lGOdaK = Child_veGQlA(7, 71)
    pt(obj_lGOdaK.get_id())
    

    asn fn async_task_nJqzDQ(x) {
        rev x * 2
    }
    uni async_task_nJqzDQ(57)
    
    asn fn worker_sGZyuO(y) {
        pt("working on", y)
    }
    mul 2 worker_sGZyuO(50)
    
    awt
    

    sum_zvzczK = 0
    lp 15 {
        sum_zvzczK = sum_zvzczK + 1
        if (sum_zvzczK > 100) { done }
    }
    pt(sum_zvzczK)
    

    sum_VKBouq = 0
    lp 9 {
        sum_VKBouq = sum_VKBouq + 1
        if (sum_VKBouq > 100) { done }
    }
    pt(sum_VKBouq)
    

    m_jiHniL = math.sqrt(61)
    h_vOryAG = hash.md5("test_wcdLTj")
    log.inf("hash:", h_vOryAG)
    pt(m_jiHniL)
    

    cns C_kfpFmJ = 37
    v_VvrEtf = unknown
    v_VvrEtf = C_kfpFmJ + 70
    if (v_VvrEtf > 0) {
        v_VvrEtf = v_VvrEtf * 2
    } el {
        v_VvrEtf = 0
    }
    pt(v_VvrEtf)
    

    cl Base_fDvYbn {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_zzcMZp <- Base_fDvYbn {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_csWSOa = Child_zzcMZp(11, 29)
    pt(obj_csWSOa.get_id())
    

    fn calc_FFlvWa(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_htBLlH = calc_FFlvWa(55, 76+1)
    pt(res_htBLlH)
    

    cl Base_PoxGdF {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_ivKxCe <- Base_PoxGdF {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_duHuNN = Child_ivKxCe(52, 53)
    pt(obj_duHuNN.get_id())
    

    arr_XfLbcm = [77, 26+1, 4+2]
    map_knsXjg = {"a": arr_XfLbcm[0], "b": arr_XfLbcm[1]}
    pt(map_knsXjg["a"])
    

    m_NKApSG = math.sqrt(58)
    h_RCLDrz = hash.md5("test_dGmlur")
    log.inf("hash:", h_RCLDrz)
    pt(m_NKApSG)
    

    val_ABFtNv = 52 % 3
    on (val_ABFtNv) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_ABFtNv)
    

    m_Pzvtja = math.sqrt(55)
    h_oSOfRm = hash.md5("test_RAUnbQ")
    log.inf("hash:", h_oSOfRm)
    pt(m_Pzvtja)
    

    asn fn async_task_bimhQp(x) {
        rev x * 2
    }
    uni async_task_bimhQp(16)
    
    asn fn worker_TqNXUm(y) {
        pt("working on", y)
    }
    mul 2 worker_TqNXUm(38)
    
    awt
    

    use "dummy.y"
    
    e_IiDeOr = enc.b64("hello drylang")
    pt(e_IiDeOr)
    
    j_YcfXQn = json(`{"test": 123}`)
    pt(j_YcfXQn)
    
    // Test get() for type info
    val_wZSMHM = get("hello")
    pt(val_wZSMHM)
    
