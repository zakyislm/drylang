
    sum_ysrpYp = 0
    lp 13 {
        sum_ysrpYp = sum_ysrpYp + 1
        if (sum_ysrpYp > 100) { done }
    }
    pt(sum_ysrpYp)
    

    val_IMtzqk = 93 % 3
    on (val_IMtzqk) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_IMtzqk)
    

    val_oxSaIe = 99 % 3
    on (val_oxSaIe) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_oxSaIe)
    

    use "dummy.y"
    
    e_LdOPKK = enc.b64("hello drylang")
    pt(e_LdOPKK)
    
    j_nFqEuK = json(`{"test": 123}`)
    pt(j_nFqEuK)
    
    // Test get() for type info
    val_vtnzQB = get("hello")
    pt(val_vtnzQB)
    

    m_asCRjn = math.sqrt(17)
    h_AiLNgY = hash.md5("test_QjeofA")
    log.inf("hash:", h_AiLNgY)
    pt(m_asCRjn)
    

    try {
        throw_bWkYrn = unknown
        throw_bWkYrn()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_AQgWyo = enc.b64("hello drylang")
    pt(e_AQgWyo)
    
    j_DhoaDE = json(`{"test": 123}`)
    pt(j_DhoaDE)
    
    // Test get() for type info
    val_FOTvnD = get("hello")
    pt(val_FOTvnD)
    

    m_ixBOgm = math.sqrt(86)
    h_ThOcVz = hash.md5("test_zXOiUF")
    log.inf("hash:", h_ThOcVz)
    pt(m_ixBOgm)
    

    fn calc_zCSJBE(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_HSEpBS = calc_zCSJBE(89, 57+1)
    pt(res_HSEpBS)
    

    val_nbPVkY = 56 % 3
    on (val_nbPVkY) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_nbPVkY)
    

    cns C_akbIds = 9
    v_LxfNFK = unknown
    v_LxfNFK = C_akbIds + 66
    if (v_LxfNFK > 0) {
        v_LxfNFK = v_LxfNFK * 2
    } el {
        v_LxfNFK = 0
    }
    pt(v_LxfNFK)
    

    cl Base_psUIKv {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_BaamcW <- Base_psUIKv {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_uuHWMZ = Child_BaamcW(47, 60)
    pt(obj_uuHWMZ.get_id())
    

    sum_xrfRyZ = 0
    lp 11 {
        sum_xrfRyZ = sum_xrfRyZ + 1
        if (sum_xrfRyZ > 100) { done }
    }
    pt(sum_xrfRyZ)
    

    cl Base_GJWQQy {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_mGPfsH <- Base_GJWQQy {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_QchIGy = Child_mGPfsH(64, 22)
    pt(obj_QchIGy.get_id())
    

    cl Base_jLdlFp {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_aeNUbE <- Base_jLdlFp {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_GpRMib = Child_aeNUbE(54, 91)
    pt(obj_GpRMib.get_id())
    

    arr_IHjvqU = [9, 43+1, 38+2]
    map_IHGxGD = {"a": arr_IHjvqU[0], "b": arr_IHjvqU[1]}
    pt(map_IHGxGD["a"])
    
