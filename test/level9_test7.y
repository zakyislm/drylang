
    asn fn async_task_PSgNrA(x) {
        rev x * 2
    }
    uni async_task_PSgNrA(24)
    
    asn fn worker_DKWlpz(y) {
        pt("working on", y)
    }
    mul 2 worker_DKWlpz(85)
    
    awt
    

    cl Base_iSnnQA {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_yDcFpO <- Base_iSnnQA {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_QWrBIc = Child_yDcFpO(82, 35)
    pt(obj_QWrBIc.get_id())
    

    cns C_fAUEap = 57
    v_FAvjSb = unknown
    v_FAvjSb = C_fAUEap + 18
    if (v_FAvjSb > 0) {
        v_FAvjSb = v_FAvjSb * 2
    } el {
        v_FAvjSb = 0
    }
    pt(v_FAvjSb)
    

    arr_XnHSFm = [8, 59+1, 50+2]
    map_mbguYE = {"a": arr_XnHSFm[0], "b": arr_XnHSFm[1]}
    pt(map_mbguYE["a"])
    

    use "dummy.y"
    
    e_zTYXxm = enc.b64("hello drylang")
    pt(e_zTYXxm)
    
    j_zKeATp = json(`{"test": 123}`)
    pt(j_zKeATp)
    
    // Test get() for type info
    val_ohvHlz = get("hello")
    pt(val_ohvHlz)
    

    asn fn async_task_SWbdyh(x) {
        rev x * 2
    }
    uni async_task_SWbdyh(63)
    
    asn fn worker_lnNkUt(y) {
        pt("working on", y)
    }
    mul 2 worker_lnNkUt(15)
    
    awt
    

    arr_eXchyN = [89, 86+1, 15+2]
    map_feMyud = {"a": arr_eXchyN[0], "b": arr_eXchyN[1]}
    pt(map_feMyud["a"])
    

    fn calc_gVpxLk(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_YlHsOy = calc_gVpxLk(19, 16+1)
    pt(res_YlHsOy)
    

    fn calc_MdLrfd(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_tHNwrA = calc_MdLrfd(77, 68+1)
    pt(res_tHNwrA)
    

    cl Base_svehPh {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_XPhjbE <- Base_svehPh {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_xwQSsM = Child_XPhjbE(70, 66)
    pt(obj_xwQSsM.get_id())
    

    val_buuNza = 50 % 3
    on (val_buuNza) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_buuNza)
    

    arr_mdsvHR = [48, 69+1, 16+2]
    map_PKwjRp = {"a": arr_mdsvHR[0], "b": arr_mdsvHR[1]}
    pt(map_PKwjRp["a"])
    

    cl Base_evENip {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_SiAZKy <- Base_evENip {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_xHPtOZ = Child_SiAZKy(94, 7)
    pt(obj_xHPtOZ.get_id())
    

    arr_afqfIe = [11, 2+1, 64+2]
    map_kYDaSN = {"a": arr_afqfIe[0], "b": arr_afqfIe[1]}
    pt(map_kYDaSN["a"])
    

    arr_BAYoEM = [6, 65+1, 37+2]
    map_DJuIci = {"a": arr_BAYoEM[0], "b": arr_BAYoEM[1]}
    pt(map_DJuIci["a"])
    

    use "dummy.y"
    
    e_tjrevL = enc.b64("hello drylang")
    pt(e_tjrevL)
    
    j_vUpveT = json(`{"test": 123}`)
    pt(j_vUpveT)
    
    // Test get() for type info
    val_BwRRrQ = get("hello")
    pt(val_BwRRrQ)
    
