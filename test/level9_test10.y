
    fn calc_wVJuak(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_FCMFPy = calc_wVJuak(63, 38+1)
    pt(res_FCMFPy)
    

    use "dummy.y"
    
    e_zNWkvO = enc.b64("hello drylang")
    pt(e_zNWkvO)
    
    j_DormXm = json(`{"test": 123}`)
    pt(j_DormXm)
    
    // Test get() for type info
    val_hLPOHl = get("hello")
    pt(val_hLPOHl)
    

    use "dummy.y"
    
    e_RytZsH = enc.b64("hello drylang")
    pt(e_RytZsH)
    
    j_mIpUhh = json(`{"test": 123}`)
    pt(j_mIpUhh)
    
    // Test get() for type info
    val_KYIeWy = get("hello")
    pt(val_KYIeWy)
    

    try {
        throw_ILAxcL = unknown
        throw_ILAxcL()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_JlXNZL = enc.b64("hello drylang")
    pt(e_JlXNZL)
    
    j_Axgtcg = json(`{"test": 123}`)
    pt(j_Axgtcg)
    
    // Test get() for type info
    val_tgfuMl = get("hello")
    pt(val_tgfuMl)
    

    sum_OBoLWo = 0
    lp 9 {
        sum_OBoLWo = sum_OBoLWo + 1
        if (sum_OBoLWo > 100) { done }
    }
    pt(sum_OBoLWo)
    

    m_fbTLMT = math.sqrt(49)
    h_IzrrQC = hash.md5("test_cIBRSi")
    log.inf("hash:", h_IzrrQC)
    pt(m_fbTLMT)
    

    val_VHonyk = 44 % 3
    on (val_VHonyk) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_VHonyk)
    

    m_fwAWPs = math.sqrt(32)
    h_AsjVzr = hash.md5("test_sfWUyK")
    log.inf("hash:", h_AsjVzr)
    pt(m_fwAWPs)
    

    fn calc_wACmwj(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_xgwKYv = calc_wACmwj(92, 32+1)
    pt(res_xgwKYv)
    

    cl Base_kKHOES {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_XjqdOw <- Base_kKHOES {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_eXHfTr = Child_XjqdOw(20, 49)
    pt(obj_eXHfTr.get_id())
    

    cl Base_rCpqFu {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_BbyFzG <- Base_rCpqFu {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_UBNTtX = Child_BbyFzG(7, 20)
    pt(obj_UBNTtX.get_id())
    

    cns C_gTrNHP = 8
    v_cLawvp = unknown
    v_cLawvp = C_gTrNHP + 66
    if (v_cLawvp > 0) {
        v_cLawvp = v_cLawvp * 2
    } el {
        v_cLawvp = 0
    }
    pt(v_cLawvp)
    

    m_djLRWS = math.sqrt(21)
    h_eWtmBu = hash.md5("test_gujEos")
    log.inf("hash:", h_eWtmBu)
    pt(m_djLRWS)
    

    val_tVdfZX = 73 % 3
    on (val_tVdfZX) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_tVdfZX)
    

    cns C_TXMYqu = 64
    v_ZIWwyd = unknown
    v_ZIWwyd = C_TXMYqu + 28
    if (v_ZIWwyd > 0) {
        v_ZIWwyd = v_ZIWwyd * 2
    } el {
        v_ZIWwyd = 0
    }
    pt(v_ZIWwyd)
    
