
    m_YiKWQU = math.sqrt(65)
    h_GmxBgQ = hash.md5("test_WSeifO")
    log.inf("hash:", h_GmxBgQ)
    pt(m_YiKWQU)
    

    try {
        throw_wnHRWi = unknown
        throw_wnHRWi()
    } err(e) {
        pt("caught error")
    }
    

    val_PYuvrG = 20 % 3
    on (val_PYuvrG) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_PYuvrG)
    

    arr_REmDEN = [38, 59+1, 16+2]
    map_AaVdVB = {"a": arr_REmDEN[0], "b": arr_REmDEN[1]}
    pt(map_AaVdVB["a"])
    

    sum_prWelx = 0
    lp 6 {
        sum_prWelx = sum_prWelx + 1
        if (sum_prWelx > 100) { done }
    }
    pt(sum_prWelx)
    

    cns C_hMrWcU = 7
    v_wYteRR = unknown
    v_wYteRR = C_hMrWcU + 96
    if (v_wYteRR > 0) {
        v_wYteRR = v_wYteRR * 2
    } el {
        v_wYteRR = 0
    }
    pt(v_wYteRR)
    

    cl Base_bYVRoM {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_mSdgvn <- Base_bYVRoM {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_EzdIUj = Child_mSdgvn(38, 75)
    pt(obj_EzdIUj.get_id())
    

    m_loajRv = math.sqrt(97)
    h_DCIvwJ = hash.md5("test_kFcOWQ")
    log.inf("hash:", h_DCIvwJ)
    pt(m_loajRv)
    

    fn calc_vZbMtp(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_XRxErm = calc_vZbMtp(53, 98+1)
    pt(res_XRxErm)
    

    fn calc_OwaMzI(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_FDcQoN = calc_OwaMzI(12, 25+1)
    pt(res_FDcQoN)
    

    cl Base_iCGAqK {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_bANQfh <- Base_iCGAqK {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_yVbWpf = Child_bANQfh(92, 77)
    pt(obj_yVbWpf.get_id())
    

    arr_LvQiyD = [55, 94+1, 36+2]
    map_NXWhOK = {"a": arr_LvQiyD[0], "b": arr_LvQiyD[1]}
    pt(map_NXWhOK["a"])
    

    cl Base_MQrwvG {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_CLvwnP <- Base_MQrwvG {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_cyhsFC = Child_CLvwnP(83, 30)
    pt(obj_cyhsFC.get_id())
    

    cns C_cRkVDa = 83
    v_eDsfCi = unknown
    v_eDsfCi = C_cRkVDa + 52
    if (v_eDsfCi > 0) {
        v_eDsfCi = v_eDsfCi * 2
    } el {
        v_eDsfCi = 0
    }
    pt(v_eDsfCi)
    

    use "dummy.y"
    
    e_PodBKZ = enc.b64("hello drylang")
    pt(e_PodBKZ)
    
    j_rFHLlO = json(`{"test": 123}`)
    pt(j_rFHLlO)
    
    // Test get() for type info
    val_wnrqvh = get("hello")
    pt(val_wnrqvh)
    
