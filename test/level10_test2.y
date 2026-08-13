
    use "dummy.y"
    
    e_xHGkOL = enc.b64("hello drylang")
    pt(e_xHGkOL)
    
    j_TXHExe = json(`{"test": 123}`)
    pt(j_TXHExe)
    
    // Test get() for type info
    val_yZYfbc = get("hello")
    pt(val_yZYfbc)
    

    sum_vvYNyS = 0
    lp 12 {
        sum_vvYNyS = sum_vvYNyS + 1
        if (sum_vvYNyS > 100) { done }
    }
    pt(sum_vvYNyS)
    

    use "dummy.y"
    
    e_kpEHSH = enc.b64("hello drylang")
    pt(e_kpEHSH)
    
    j_ELBpnT = json(`{"test": 123}`)
    pt(j_ELBpnT)
    
    // Test get() for type info
    val_JFQMdm = get("hello")
    pt(val_JFQMdm)
    

    sum_UhqQbC = 0
    lp 10 {
        sum_UhqQbC = sum_UhqQbC + 1
        if (sum_UhqQbC > 100) { done }
    }
    pt(sum_UhqQbC)
    

    val_FCJYzy = 54 % 3
    on (val_FCJYzy) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_FCJYzy)
    

    val_LfFVeq = 77 % 3
    on (val_LfFVeq) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_LfFVeq)
    

    sum_uYTBBy = 0
    lp 11 {
        sum_uYTBBy = sum_uYTBBy + 1
        if (sum_uYTBBy > 100) { done }
    }
    pt(sum_uYTBBy)
    

    use "dummy.y"
    
    e_IZjfrt = enc.b64("hello drylang")
    pt(e_IZjfrt)
    
    j_jkgMcL = json(`{"test": 123}`)
    pt(j_jkgMcL)
    
    // Test get() for type info
    val_skJTQE = get("hello")
    pt(val_skJTQE)
    

    sum_lxjyWv = 0
    lp 6 {
        sum_lxjyWv = sum_lxjyWv + 1
        if (sum_lxjyWv > 100) { done }
    }
    pt(sum_lxjyWv)
    

    cns C_tBEVah = 50
    v_iahxKN = unknown
    v_iahxKN = C_tBEVah + 93
    if (v_iahxKN > 0) {
        v_iahxKN = v_iahxKN * 2
    } el {
        v_iahxKN = 0
    }
    pt(v_iahxKN)
    

    use "dummy.y"
    
    e_dQfhsZ = enc.b64("hello drylang")
    pt(e_dQfhsZ)
    
    j_VYRsZN = json(`{"test": 123}`)
    pt(j_VYRsZN)
    
    // Test get() for type info
    val_wEebFc = get("hello")
    pt(val_wEebFc)
    

    fn calc_TdvNOD(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_zSKTWO = calc_TdvNOD(50, 75+1)
    pt(res_zSKTWO)
    

    val_QStbIz = 9 % 3
    on (val_QStbIz) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_QStbIz)
    

    arr_mHTUyH = [80, 8+1, 45+2]
    map_gLCFHG = {"a": arr_mHTUyH[0], "b": arr_mHTUyH[1]}
    pt(map_gLCFHG["a"])
    

    cns C_sJMdyi = 46
    v_rHoBqs = unknown
    v_rHoBqs = C_sJMdyi + 68
    if (v_rHoBqs > 0) {
        v_rHoBqs = v_rHoBqs * 2
    } el {
        v_rHoBqs = 0
    }
    pt(v_rHoBqs)
    

    cl Base_rfSFVa {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_RxlyQf <- Base_rfSFVa {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_EhxbFs = Child_RxlyQf(70, 7)
    pt(obj_EhxbFs.get_id())
    

    asn fn async_task_bfBJDu(x) {
        rev x * 2
    }
    uni async_task_bfBJDu(45)
    
    asn fn worker_OVvKvq(y) {
        pt("working on", y)
    }
    mul 2 worker_OVvKvq(6)
    
    awt
    
