
    cl Base_mFOGgH {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_rSnmvE <- Base_mFOGgH {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_SmSllI = Child_rSnmvE(98, 88)
    pt(obj_SmSllI.get_id())
    

    m_JeTSPi = math.sqrt(57)
    h_GOSucT = hash.md5("test_Thlxsh")
    log.inf("hash:", h_GOSucT)
    pt(m_JeTSPi)
    

    asn fn async_task_EPrCvI(x) {
        rev x * 2
    }
    uni async_task_EPrCvI(58)
    
    asn fn worker_cwiYwd(y) {
        pt("working on", y)
    }
    mul 2 worker_cwiYwd(37)
    
    awt
    

    m_NzzdLQ = math.sqrt(32)
    h_DBtOEP = hash.md5("test_bSoInh")
    log.inf("hash:", h_DBtOEP)
    pt(m_NzzdLQ)
    

    use "dummy.y"
    
    e_TzuGJE = enc.b64("hello drylang")
    pt(e_TzuGJE)
    
    j_LvEBTe = json(`{"test": 123}`)
    pt(j_LvEBTe)
    
    // Test get() for type info
    val_wyXotl = get("hello")
    pt(val_wyXotl)
    

    val_xJOJAZ = 80 % 3
    on (val_xJOJAZ) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_xJOJAZ)
    

    asn fn async_task_GHYEUR(x) {
        rev x * 2
    }
    uni async_task_GHYEUR(85)
    
    asn fn worker_IRXghv(y) {
        pt("working on", y)
    }
    mul 2 worker_IRXghv(6)
    
    awt
    

    cns C_SZdIXk = 13
    v_NALdGC = unknown
    v_NALdGC = C_SZdIXk + 82
    if (v_NALdGC > 0) {
        v_NALdGC = v_NALdGC * 2
    } el {
        v_NALdGC = 0
    }
    pt(v_NALdGC)
    

    use "dummy.y"
    
    e_taMAge = enc.b64("hello drylang")
    pt(e_taMAge)
    
    j_ntcluk = json(`{"test": 123}`)
    pt(j_ntcluk)
    
    // Test get() for type info
    val_avblZZ = get("hello")
    pt(val_avblZZ)
    

    m_nIkCWz = math.sqrt(14)
    h_SdmktG = hash.md5("test_XtgbqS")
    log.inf("hash:", h_SdmktG)
    pt(m_nIkCWz)
    

    use "dummy.y"
    
    e_kasvRC = enc.b64("hello drylang")
    pt(e_kasvRC)
    
    j_GnajiJ = json(`{"test": 123}`)
    pt(j_GnajiJ)
    
    // Test get() for type info
    val_kceClB = get("hello")
    pt(val_kceClB)
    

    m_ZcrwDS = math.sqrt(92)
    h_DePWnP = hash.md5("test_HMlQcm")
    log.inf("hash:", h_DePWnP)
    pt(m_ZcrwDS)
    

    val_nIBJRr = 29 % 3
    on (val_nIBJRr) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_nIBJRr)
    

    fn calc_vVuuxA(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_JNvCtX = calc_vVuuxA(11, 82+1)
    pt(res_JNvCtX)
    
