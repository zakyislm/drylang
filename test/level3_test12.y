
    fn calc_dizshO(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_HgujDX = calc_dizshO(53, 75+1)
    pt(res_HgujDX)
    

    cl Base_QjhwVl {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_yaoDQT <- Base_QjhwVl {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_PdgXhb = Child_yaoDQT(56, 85)
    pt(obj_PdgXhb.get_id())
    

    cns C_SEdezC = 80
    v_RvofGS = unknown
    v_RvofGS = C_SEdezC + 26
    if (v_RvofGS > 0) {
        v_RvofGS = v_RvofGS * 2
    } el {
        v_RvofGS = 0
    }
    pt(v_RvofGS)
    

    m_JOdqUB = math.sqrt(91)
    h_ONoMkd = hash.md5("test_albBfg")
    log.inf("hash:", h_ONoMkd)
    pt(m_JOdqUB)
    

    use "dummy.y"
    
    e_ezJppe = enc.b64("hello drylang")
    pt(e_ezJppe)
    
    j_bKsQln = json(`{"test": 123}`)
    pt(j_bKsQln)
    
    // Test get() for type info
    val_GWiker = get("hello")
    pt(val_GWiker)
    
