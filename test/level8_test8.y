
    val_htAUVT = 67 % 3
    on (val_htAUVT) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_htAUVT)
    

    try {
        throw_PxplaQ = unknown
        throw_PxplaQ()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_ZjDTzj = unknown
        throw_ZjDTzj()
    } err(e) {
        pt("caught error")
    }
    

    val_kJbokn = 88 % 3
    on (val_kJbokn) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_kJbokn)
    

    m_YRtFrs = math.sqrt(50)
    h_tCiuof = hash.md5("test_xNtuDt")
    log.inf("hash:", h_tCiuof)
    pt(m_YRtFrs)
    

    m_tddvHr = math.sqrt(57)
    h_IkKaUH = hash.md5("test_tlZrWy")
    log.inf("hash:", h_IkKaUH)
    pt(m_tddvHr)
    

    asn fn async_task_vLzdta(x) {
        rev x * 2
    }
    uni async_task_vLzdta(87)
    
    asn fn worker_AsSrbB(y) {
        pt("working on", y)
    }
    mul 2 worker_AsSrbB(7)
    
    awt
    

    try {
        throw_pqTTJd = unknown
        throw_pqTTJd()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_sqVIEw = enc.b64("hello drylang")
    pt(e_sqVIEw)
    
    j_KXpXzY = json(`{"test": 123}`)
    pt(j_KXpXzY)
    
    // Test get() for type info
    val_mfYfbB = get("hello")
    pt(val_mfYfbB)
    

    cns C_Tpwdfh = 71
    v_QEnzWG = unknown
    v_QEnzWG = C_Tpwdfh + 86
    if (v_QEnzWG > 0) {
        v_QEnzWG = v_QEnzWG * 2
    } el {
        v_QEnzWG = 0
    }
    pt(v_QEnzWG)
    

    cl Base_XWqvDo {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_eKqHVS <- Base_XWqvDo {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_EzexIn = Child_eKqHVS(75, 13)
    pt(obj_EzexIn.get_id())
    

    m_YXrFCl = math.sqrt(50)
    h_rtsqgx = hash.md5("test_JaNSZm")
    log.inf("hash:", h_rtsqgx)
    pt(m_YXrFCl)
    

    try {
        throw_YaetZG = unknown
        throw_YaetZG()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_URtafT {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_GAZSsU <- Base_URtafT {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_cjCJwo = Child_GAZSsU(77, 97)
    pt(obj_cjCJwo.get_id())
    

    fn calc_OJpfcb(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_gRFHAS = calc_OJpfcb(53, 77+1)
    pt(res_gRFHAS)
    
