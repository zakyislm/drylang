
    sum_MCiAGO = 0
    lp 14 {
        sum_MCiAGO = sum_MCiAGO + 1
        if (sum_MCiAGO > 100) { done }
    }
    pt(sum_MCiAGO)
    

    cl Base_cIQebX {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_mQYmtn <- Base_cIQebX {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_CtmrXl = Child_mQYmtn(74, 45)
    pt(obj_CtmrXl.get_id())
    

    val_tgQUEm = 75 % 3
    on (val_tgQUEm) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_tgQUEm)
    

    cns C_SgkSdV = 2
    v_COEsKM = unknown
    v_COEsKM = C_SgkSdV + 92
    if (v_COEsKM > 0) {
        v_COEsKM = v_COEsKM * 2
    } el {
        v_COEsKM = 0
    }
    pt(v_COEsKM)
    

    use "dummy.y"
    
    e_IMdLfW = enc.b64("hello drylang")
    pt(e_IMdLfW)
    
    j_NDXIZL = json(`{"test": 123}`)
    pt(j_NDXIZL)
    
    // Test get() for type info
    val_vhCXAh = get("hello")
    pt(val_vhCXAh)
    

    sum_kiETxf = 0
    lp 15 {
        sum_kiETxf = sum_kiETxf + 1
        if (sum_kiETxf > 100) { done }
    }
    pt(sum_kiETxf)
    

    asn fn async_task_gJCent(x) {
        rev x * 2
    }
    uni async_task_gJCent(3)
    
    asn fn worker_bLbjcX(y) {
        pt("working on", y)
    }
    mul 2 worker_bLbjcX(19)
    
    awt
    

    use "dummy.y"
    
    e_HOlImB = enc.b64("hello drylang")
    pt(e_HOlImB)
    
    j_IVogLw = json(`{"test": 123}`)
    pt(j_IVogLw)
    
    // Test get() for type info
    val_Xmktly = get("hello")
    pt(val_Xmktly)
    
