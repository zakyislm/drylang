
    fn calc_BARgNx(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_KjoZpJ = calc_BARgNx(32, 85+1)
    pt(res_KjoZpJ)
    

    sum_WvhrJg = 0
    lp 11 {
        sum_WvhrJg = sum_WvhrJg + 1
        if (sum_WvhrJg > 100) { done }
    }
    pt(sum_WvhrJg)
    

    val_XqJsjG = 57 % 3
    on (val_XqJsjG) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_XqJsjG)
    

    val_BJGLdx = 22 % 3
    on (val_BJGLdx) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_BJGLdx)
    

    use "dummy.y"
    
    e_XtpUrr = enc.b64("hello drylang")
    pt(e_XtpUrr)
    
    j_fJovxW = json(`{"test": 123}`)
    pt(j_fJovxW)
    
    // Test get() for type info
    val_QLnFXE = get("hello")
    pt(val_QLnFXE)
    

    cns C_zFNAmy = 60
    v_PfNEkJ = unknown
    v_PfNEkJ = C_zFNAmy + 27
    if (v_PfNEkJ > 0) {
        v_PfNEkJ = v_PfNEkJ * 2
    } el {
        v_PfNEkJ = 0
    }
    pt(v_PfNEkJ)
    

    cl Base_DyvsEJ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_kbGWlF <- Base_DyvsEJ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_eFwxGA = Child_kbGWlF(1, 84)
    pt(obj_eFwxGA.get_id())
    

    cns C_eeKQDn = 64
    v_YWvFdK = unknown
    v_YWvFdK = C_eeKQDn + 11
    if (v_YWvFdK > 0) {
        v_YWvFdK = v_YWvFdK * 2
    } el {
        v_YWvFdK = 0
    }
    pt(v_YWvFdK)
    

    asn fn async_task_DyGDNH(x) {
        rev x * 2
    }
    uni async_task_DyGDNH(25)
    
    asn fn worker_fPawVL(y) {
        pt("working on", y)
    }
    mul 2 worker_fPawVL(72)
    
    awt
    

    try {
        throw_aYjCVn = unknown
        throw_aYjCVn()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_aFfAaT = enc.b64("hello drylang")
    pt(e_aFfAaT)
    
    j_KiIRZM = json(`{"test": 123}`)
    pt(j_KiIRZM)
    
    // Test get() for type info
    val_oOcGMN = get("hello")
    pt(val_oOcGMN)
    

    sum_AidZMy = 0
    lp 13 {
        sum_AidZMy = sum_AidZMy + 1
        if (sum_AidZMy > 100) { done }
    }
    pt(sum_AidZMy)
    

    cns C_OMsYPM = 84
    v_lHWFsH = unknown
    v_lHWFsH = C_OMsYPM + 91
    if (v_lHWFsH > 0) {
        v_lHWFsH = v_lHWFsH * 2
    } el {
        v_lHWFsH = 0
    }
    pt(v_lHWFsH)
    

    use "dummy.y"
    
    e_LRUYMI = enc.b64("hello drylang")
    pt(e_LRUYMI)
    
    j_aaCXhO = json(`{"test": 123}`)
    pt(j_aaCXhO)
    
    // Test get() for type info
    val_whEqOA = get("hello")
    pt(val_whEqOA)
    

    asn fn async_task_iQUDBI(x) {
        rev x * 2
    }
    uni async_task_iQUDBI(7)
    
    asn fn worker_pMnmUt(y) {
        pt("working on", y)
    }
    mul 2 worker_pMnmUt(87)
    
    awt
    
