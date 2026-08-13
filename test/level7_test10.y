
    val_lOPraO = 42 % 3
    on (val_lOPraO) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_lOPraO)
    

    m_FjelNo = math.sqrt(46)
    h_pYvekO = hash.md5("test_xzLgWN")
    log.inf("hash:", h_pYvekO)
    pt(m_FjelNo)
    

    arr_JQdEKZ = [35, 16+1, 99+2]
    map_peIUAS = {"a": arr_JQdEKZ[0], "b": arr_JQdEKZ[1]}
    pt(map_peIUAS["a"])
    

    try {
        throw_tynHSK = unknown
        throw_tynHSK()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_jQaxoS = unknown
        throw_jQaxoS()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_Eoogsw(x) {
        rev x * 2
    }
    uni async_task_Eoogsw(100)
    
    asn fn worker_dcgNUX(y) {
        pt("working on", y)
    }
    mul 2 worker_dcgNUX(99)
    
    awt
    

    cl Base_QJhAFT {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_miHcLb <- Base_QJhAFT {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_XhKGuJ = Child_miHcLb(65, 74)
    pt(obj_XhKGuJ.get_id())
    

    use "dummy.y"
    
    e_xfOeEW = enc.b64("hello drylang")
    pt(e_xfOeEW)
    
    j_bRlnNM = json(`{"test": 123}`)
    pt(j_bRlnNM)
    
    // Test get() for type info
    val_nuUOYF = get("hello")
    pt(val_nuUOYF)
    

    arr_VscseL = [56, 79+1, 84+2]
    map_SYBjoN = {"a": arr_VscseL[0], "b": arr_VscseL[1]}
    pt(map_SYBjoN["a"])
    

    use "dummy.y"
    
    e_xjhNBx = enc.b64("hello drylang")
    pt(e_xjhNBx)
    
    j_BKGKEd = json(`{"test": 123}`)
    pt(j_BKGKEd)
    
    // Test get() for type info
    val_LUCDLZ = get("hello")
    pt(val_LUCDLZ)
    

    cl Base_rSOUGT {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_HQGsQD <- Base_rSOUGT {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_MIUtPc = Child_HQGsQD(95, 66)
    pt(obj_MIUtPc.get_id())
    

    use "dummy.y"
    
    e_OhEeCl = enc.b64("hello drylang")
    pt(e_OhEeCl)
    
    j_fCTLqi = json(`{"test": 123}`)
    pt(j_fCTLqi)
    
    // Test get() for type info
    val_oEuZMc = get("hello")
    pt(val_oEuZMc)
    

    arr_ILpqWp = [30, 99+1, 98+2]
    map_xWLcMg = {"a": arr_ILpqWp[0], "b": arr_ILpqWp[1]}
    pt(map_xWLcMg["a"])
    

    sum_bqbtnZ = 0
    lp 5 {
        sum_bqbtnZ = sum_bqbtnZ + 1
        if (sum_bqbtnZ > 100) { done }
    }
    pt(sum_bqbtnZ)
    
