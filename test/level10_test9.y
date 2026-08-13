
    try {
        throw_oozuSc = unknown
        throw_oozuSc()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_ZylKOv = unknown
        throw_ZylKOv()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_sXorhU(x) {
        rev x * 2
    }
    uni async_task_sXorhU(38)
    
    asn fn worker_cfkuHd(y) {
        pt("working on", y)
    }
    mul 2 worker_cfkuHd(3)
    
    awt
    

    use "dummy.y"
    
    e_PrYJyi = enc.b64("hello drylang")
    pt(e_PrYJyi)
    
    j_ponoBk = json(`{"test": 123}`)
    pt(j_ponoBk)
    
    // Test get() for type info
    val_VMgkqv = get("hello")
    pt(val_VMgkqv)
    

    use "dummy.y"
    
    e_NZmttE = enc.b64("hello drylang")
    pt(e_NZmttE)
    
    j_xLlARz = json(`{"test": 123}`)
    pt(j_xLlARz)
    
    // Test get() for type info
    val_TNPxNm = get("hello")
    pt(val_TNPxNm)
    

    sum_auztGm = 0
    lp 8 {
        sum_auztGm = sum_auztGm + 1
        if (sum_auztGm > 100) { done }
    }
    pt(sum_auztGm)
    

    val_JMWupn = 62 % 3
    on (val_JMWupn) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_JMWupn)
    

    use "dummy.y"
    
    e_BAZQPE = enc.b64("hello drylang")
    pt(e_BAZQPE)
    
    j_LkMUFA = json(`{"test": 123}`)
    pt(j_LkMUFA)
    
    // Test get() for type info
    val_vMdzUF = get("hello")
    pt(val_vMdzUF)
    

    sum_bgHsyX = 0
    lp 15 {
        sum_bgHsyX = sum_bgHsyX + 1
        if (sum_bgHsyX > 100) { done }
    }
    pt(sum_bgHsyX)
    

    use "dummy.y"
    
    e_iXvPMk = enc.b64("hello drylang")
    pt(e_iXvPMk)
    
    j_eDqytX = json(`{"test": 123}`)
    pt(j_eDqytX)
    
    // Test get() for type info
    val_ZVkQZn = get("hello")
    pt(val_ZVkQZn)
    

    use "dummy.y"
    
    e_zXUyxS = enc.b64("hello drylang")
    pt(e_zXUyxS)
    
    j_AYnvDr = json(`{"test": 123}`)
    pt(j_AYnvDr)
    
    // Test get() for type info
    val_ZjPkcL = get("hello")
    pt(val_ZjPkcL)
    

    val_TqEceH = 19 % 3
    on (val_TqEceH) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_TqEceH)
    

    try {
        throw_TZkwIv = unknown
        throw_TZkwIv()
    } err(e) {
        pt("caught error")
    }
    

    arr_MaFEBF = [25, 19+1, 57+2]
    map_eSjLcB = {"a": arr_MaFEBF[0], "b": arr_MaFEBF[1]}
    pt(map_eSjLcB["a"])
    

    sum_Kepmwn = 0
    lp 15 {
        sum_Kepmwn = sum_Kepmwn + 1
        if (sum_Kepmwn > 100) { done }
    }
    pt(sum_Kepmwn)
    

    cl Base_iNStXq {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_IeHmQo <- Base_iNStXq {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_TTOdjr = Child_IeHmQo(74, 60)
    pt(obj_TTOdjr.get_id())
    

    cl Base_ChdPqH {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_nSoYmw <- Base_ChdPqH {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_rrYQfZ = Child_nSoYmw(90, 22)
    pt(obj_rrYQfZ.get_id())
    
