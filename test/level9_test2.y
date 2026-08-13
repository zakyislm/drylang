
    try {
        throw_yXNblT = unknown
        throw_yXNblT()
    } err(e) {
        pt("caught error")
    }
    

    val_vCUxBc = 50 % 3
    on (val_vCUxBc) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_vCUxBc)
    

    try {
        throw_LdpZYN = unknown
        throw_LdpZYN()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_hVJVYN {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_vxqpxP <- Base_hVJVYN {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_KUgxAr = Child_vxqpxP(70, 91)
    pt(obj_KUgxAr.get_id())
    

    fn calc_sigsQn(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_WEYCKf = calc_sigsQn(46, 22+1)
    pt(res_WEYCKf)
    

    try {
        throw_tccbDz = unknown
        throw_tccbDz()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_ldtJxc = enc.b64("hello drylang")
    pt(e_ldtJxc)
    
    j_XhbgoY = json(`{"test": 123}`)
    pt(j_XhbgoY)
    
    // Test get() for type info
    val_vKdvaO = get("hello")
    pt(val_vKdvaO)
    

    m_vOPmhh = math.sqrt(16)
    h_ewoYZK = hash.md5("test_YzBBuQ")
    log.inf("hash:", h_ewoYZK)
    pt(m_vOPmhh)
    

    cl Base_AgQvJq {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_CYcmCi <- Base_AgQvJq {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_kCQFvy = Child_CYcmCi(67, 17)
    pt(obj_kCQFvy.get_id())
    

    cns C_rWhQOi = 39
    v_TsFGCQ = unknown
    v_TsFGCQ = C_rWhQOi + 39
    if (v_TsFGCQ > 0) {
        v_TsFGCQ = v_TsFGCQ * 2
    } el {
        v_TsFGCQ = 0
    }
    pt(v_TsFGCQ)
    

    asn fn async_task_DirOBl(x) {
        rev x * 2
    }
    uni async_task_DirOBl(72)
    
    asn fn worker_HfpUUh(y) {
        pt("working on", y)
    }
    mul 2 worker_HfpUUh(30)
    
    awt
    

    val_MqPEYd = 17 % 3
    on (val_MqPEYd) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_MqPEYd)
    

    cl Base_rMuDWk {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_MrSyIA <- Base_rMuDWk {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_llAvqk = Child_MrSyIA(83, 85)
    pt(obj_llAvqk.get_id())
    

    val_GPfmQz = 1 % 3
    on (val_GPfmQz) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_GPfmQz)
    

    sum_TZeqDc = 0
    lp 14 {
        sum_TZeqDc = sum_TZeqDc + 1
        if (sum_TZeqDc > 100) { done }
    }
    pt(sum_TZeqDc)
    

    cns C_YiEhdS = 90
    v_vBdTXo = unknown
    v_vBdTXo = C_YiEhdS + 59
    if (v_vBdTXo > 0) {
        v_vBdTXo = v_vBdTXo * 2
    } el {
        v_vBdTXo = 0
    }
    pt(v_vBdTXo)
    
