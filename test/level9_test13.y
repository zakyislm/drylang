
    val_moUivL = 51 % 3
    on (val_moUivL) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_moUivL)
    

    use "dummy.y"
    
    e_bPwTyO = enc.b64("hello drylang")
    pt(e_bPwTyO)
    
    j_kOzNJm = json(`{"test": 123}`)
    pt(j_kOzNJm)
    
    // Test get() for type info
    val_MWJozY = get("hello")
    pt(val_MWJozY)
    

    use "dummy.y"
    
    e_tqfeBZ = enc.b64("hello drylang")
    pt(e_tqfeBZ)
    
    j_nbIXAx = json(`{"test": 123}`)
    pt(j_nbIXAx)
    
    // Test get() for type info
    val_TCRMTK = get("hello")
    pt(val_TCRMTK)
    

    sum_xdiYBC = 0
    lp 13 {
        sum_xdiYBC = sum_xdiYBC + 1
        if (sum_xdiYBC > 100) { done }
    }
    pt(sum_xdiYBC)
    

    use "dummy.y"
    
    e_RVEwdp = enc.b64("hello drylang")
    pt(e_RVEwdp)
    
    j_cuzEdS = json(`{"test": 123}`)
    pt(j_cuzEdS)
    
    // Test get() for type info
    val_frzilu = get("hello")
    pt(val_frzilu)
    

    m_geqlNm = math.sqrt(40)
    h_wQTzbc = hash.md5("test_hPNJtt")
    log.inf("hash:", h_wQTzbc)
    pt(m_geqlNm)
    

    cl Base_EjnaSA {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_UPFemi <- Base_EjnaSA {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_pJpkxC = Child_UPFemi(63, 69)
    pt(obj_pJpkxC.get_id())
    

    cl Base_CEVDPz {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_TnsVKR <- Base_CEVDPz {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_GSFgdn = Child_TnsVKR(16, 58)
    pt(obj_GSFgdn.get_id())
    

    m_QHNGhF = math.sqrt(64)
    h_EBRgEV = hash.md5("test_qoZXZH")
    log.inf("hash:", h_EBRgEV)
    pt(m_QHNGhF)
    

    use "dummy.y"
    
    e_tiBryb = enc.b64("hello drylang")
    pt(e_tiBryb)
    
    j_tMXdQW = json(`{"test": 123}`)
    pt(j_tMXdQW)
    
    // Test get() for type info
    val_vSeJtA = get("hello")
    pt(val_vSeJtA)
    

    cns C_NoKsMB = 74
    v_mAJEJx = unknown
    v_mAJEJx = C_NoKsMB + 83
    if (v_mAJEJx > 0) {
        v_mAJEJx = v_mAJEJx * 2
    } el {
        v_mAJEJx = 0
    }
    pt(v_mAJEJx)
    

    cl Base_RezWxj {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_hZlWgm <- Base_RezWxj {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_xiTHGD = Child_hZlWgm(69, 82)
    pt(obj_xiTHGD.get_id())
    

    cns C_ymeBBC = 34
    v_JQfqpN = unknown
    v_JQfqpN = C_ymeBBC + 56
    if (v_JQfqpN > 0) {
        v_JQfqpN = v_JQfqpN * 2
    } el {
        v_JQfqpN = 0
    }
    pt(v_JQfqpN)
    

    try {
        throw_DTiqlY = unknown
        throw_DTiqlY()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_trViov(x) {
        rev x * 2
    }
    uni async_task_trViov(84)
    
    asn fn worker_FOTNIU(y) {
        pt("working on", y)
    }
    mul 2 worker_FOTNIU(21)
    
    awt
    

    fn calc_UIYhxu(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_URsWkQ = calc_UIYhxu(67, 50+1)
    pt(res_URsWkQ)
    
