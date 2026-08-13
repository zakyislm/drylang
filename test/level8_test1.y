
    try {
        throw_NZaXYk = unknown
        throw_NZaXYk()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_fZigTZ(x) {
        rev x * 2
    }
    uni async_task_fZigTZ(99)
    
    asn fn worker_zSErxh(y) {
        pt("working on", y)
    }
    mul 2 worker_zSErxh(88)
    
    awt
    

    try {
        throw_CSkKWG = unknown
        throw_CSkKWG()
    } err(e) {
        pt("caught error")
    }
    

    cns C_hZYzgI = 27
    v_tcgyHy = unknown
    v_tcgyHy = C_hZYzgI + 3
    if (v_tcgyHy > 0) {
        v_tcgyHy = v_tcgyHy * 2
    } el {
        v_tcgyHy = 0
    }
    pt(v_tcgyHy)
    

    asn fn async_task_MXfZUD(x) {
        rev x * 2
    }
    uni async_task_MXfZUD(71)
    
    asn fn worker_yAoZhD(y) {
        pt("working on", y)
    }
    mul 2 worker_yAoZhD(17)
    
    awt
    

    try {
        throw_PalyCF = unknown
        throw_PalyCF()
    } err(e) {
        pt("caught error")
    }
    

    arr_HamaOI = [41, 66+1, 90+2]
    map_ekEBfy = {"a": arr_HamaOI[0], "b": arr_HamaOI[1]}
    pt(map_ekEBfy["a"])
    

    cns C_ngwPwq = 93
    v_VlejsU = unknown
    v_VlejsU = C_ngwPwq + 8
    if (v_VlejsU > 0) {
        v_VlejsU = v_VlejsU * 2
    } el {
        v_VlejsU = 0
    }
    pt(v_VlejsU)
    

    cns C_HszSqG = 31
    v_GxAkvu = unknown
    v_GxAkvu = C_HszSqG + 5
    if (v_GxAkvu > 0) {
        v_GxAkvu = v_GxAkvu * 2
    } el {
        v_GxAkvu = 0
    }
    pt(v_GxAkvu)
    

    val_nfGdJX = 30 % 3
    on (val_nfGdJX) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_nfGdJX)
    

    cns C_MRbzKH = 77
    v_BQGSwk = unknown
    v_BQGSwk = C_MRbzKH + 98
    if (v_BQGSwk > 0) {
        v_BQGSwk = v_BQGSwk * 2
    } el {
        v_BQGSwk = 0
    }
    pt(v_BQGSwk)
    

    val_QoLbqB = 85 % 3
    on (val_QoLbqB) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_QoLbqB)
    

    cl Base_JLqdum {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_dGTAFL <- Base_JLqdum {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_NGaIgO = Child_dGTAFL(8, 88)
    pt(obj_NGaIgO.get_id())
    

    asn fn async_task_xtSVHs(x) {
        rev x * 2
    }
    uni async_task_xtSVHs(83)
    
    asn fn worker_wJzqjR(y) {
        pt("working on", y)
    }
    mul 2 worker_wJzqjR(8)
    
    awt
    

    use "dummy.y"
    
    e_NmMkBW = enc.b64("hello drylang")
    pt(e_NmMkBW)
    
    j_IsQSAm = json(`{"test": 123}`)
    pt(j_IsQSAm)
    
    // Test get() for type info
    val_Ngycpw = get("hello")
    pt(val_Ngycpw)
    
