
    use "dummy.y"
    
    e_MEYtJT = enc.b64("hello drylang")
    pt(e_MEYtJT)
    
    j_ozDkuL = json(`{"test": 123}`)
    pt(j_ozDkuL)
    
    // Test get() for type info
    val_SQoaQK = get("hello")
    pt(val_SQoaQK)
    

    cns C_ICWksa = 100
    v_tTElen = unknown
    v_tTElen = C_ICWksa + 15
    if (v_tTElen > 0) {
        v_tTElen = v_tTElen * 2
    } el {
        v_tTElen = 0
    }
    pt(v_tTElen)
    

    m_eiaoMk = math.sqrt(40)
    h_dgEPfA = hash.md5("test_ShroRg")
    log.inf("hash:", h_dgEPfA)
    pt(m_eiaoMk)
    

    cl Base_QLOuYf {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_chDIlj <- Base_QLOuYf {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_wqmnDI = Child_chDIlj(9, 29)
    pt(obj_wqmnDI.get_id())
    

    try {
        throw_KfZhRk = unknown
        throw_KfZhRk()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_GQBjxm {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_jfBLrw <- Base_GQBjxm {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_CYjGtb = Child_jfBLrw(48, 20)
    pt(obj_CYjGtb.get_id())
    

    arr_iedZbF = [56, 95+1, 33+2]
    map_ELsdZo = {"a": arr_iedZbF[0], "b": arr_iedZbF[1]}
    pt(map_ELsdZo["a"])
    

    m_ahfobd = math.sqrt(16)
    h_OZnKgr = hash.md5("test_IPPTlm")
    log.inf("hash:", h_OZnKgr)
    pt(m_ahfobd)
    

    cl Base_lUFJRi {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_KMgLjN <- Base_lUFJRi {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_ZQaIky = Child_KMgLjN(50, 14)
    pt(obj_ZQaIky.get_id())
    

    try {
        throw_kRvRww = unknown
        throw_kRvRww()
    } err(e) {
        pt("caught error")
    }
    

    fn calc_DCkEXM(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_FgHKan = calc_DCkEXM(79, 90+1)
    pt(res_FgHKan)
    

    arr_LfAsqF = [46, 89+1, 21+2]
    map_Tgxhwd = {"a": arr_LfAsqF[0], "b": arr_LfAsqF[1]}
    pt(map_Tgxhwd["a"])
    

    use "dummy.y"
    
    e_bjAHUo = enc.b64("hello drylang")
    pt(e_bjAHUo)
    
    j_Lcdjvb = json(`{"test": 123}`)
    pt(j_Lcdjvb)
    
    // Test get() for type info
    val_EnpMbS = get("hello")
    pt(val_EnpMbS)
    

    cl Base_mgFuiJ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_YxZZyl <- Base_mgFuiJ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_zbnNtP = Child_YxZZyl(79, 1)
    pt(obj_zbnNtP.get_id())
    

    arr_bIHhGu = [91, 62+1, 7+2]
    map_xFfeOk = {"a": arr_bIHhGu[0], "b": arr_bIHhGu[1]}
    pt(map_xFfeOk["a"])
    

    sum_VXvbxt = 0
    lp 8 {
        sum_VXvbxt = sum_VXvbxt + 1
        if (sum_VXvbxt > 100) { done }
    }
    pt(sum_VXvbxt)
    

    val_oTXozb = 19 % 3
    on (val_oTXozb) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_oTXozb)
    
