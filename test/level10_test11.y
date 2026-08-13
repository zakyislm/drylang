
    cl Base_nHnqIr {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_CvQICE <- Base_nHnqIr {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_zFQaMC = Child_CvQICE(62, 96)
    pt(obj_zFQaMC.get_id())
    

    cns C_QkdZnn = 23
    v_lAHjoT = unknown
    v_lAHjoT = C_QkdZnn + 37
    if (v_lAHjoT > 0) {
        v_lAHjoT = v_lAHjoT * 2
    } el {
        v_lAHjoT = 0
    }
    pt(v_lAHjoT)
    

    try {
        throw_aZGyUr = unknown
        throw_aZGyUr()
    } err(e) {
        pt("caught error")
    }
    

    sum_cvJwCf = 0
    lp 13 {
        sum_cvJwCf = sum_cvJwCf + 1
        if (sum_cvJwCf > 100) { done }
    }
    pt(sum_cvJwCf)
    

    cns C_OhhXKt = 92
    v_nVAxfn = unknown
    v_nVAxfn = C_OhhXKt + 45
    if (v_nVAxfn > 0) {
        v_nVAxfn = v_nVAxfn * 2
    } el {
        v_nVAxfn = 0
    }
    pt(v_nVAxfn)
    

    try {
        throw_bNPgqt = unknown
        throw_bNPgqt()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_naJCzq {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_kfKUTJ <- Base_naJCzq {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_toCcEe = Child_kfKUTJ(79, 44)
    pt(obj_toCcEe.get_id())
    

    asn fn async_task_PBStoB(x) {
        rev x * 2
    }
    uni async_task_PBStoB(58)
    
    asn fn worker_OfYpGy(y) {
        pt("working on", y)
    }
    mul 2 worker_OfYpGy(21)
    
    awt
    

    sum_SRPubN = 0
    lp 10 {
        sum_SRPubN = sum_SRPubN + 1
        if (sum_SRPubN > 100) { done }
    }
    pt(sum_SRPubN)
    

    sum_fpdPVj = 0
    lp 12 {
        sum_fpdPVj = sum_fpdPVj + 1
        if (sum_fpdPVj > 100) { done }
    }
    pt(sum_fpdPVj)
    

    try {
        throw_gvcsyw = unknown
        throw_gvcsyw()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_LAgmmW {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_CvTZOE <- Base_LAgmmW {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_hbHKYf = Child_CvTZOE(86, 16)
    pt(obj_hbHKYf.get_id())
    

    asn fn async_task_GyoeCv(x) {
        rev x * 2
    }
    uni async_task_GyoeCv(54)
    
    asn fn worker_NoShlB(y) {
        pt("working on", y)
    }
    mul 2 worker_NoShlB(74)
    
    awt
    

    use "dummy.y"
    
    e_MIlmNV = enc.b64("hello drylang")
    pt(e_MIlmNV)
    
    j_iSEXEK = json(`{"test": 123}`)
    pt(j_iSEXEK)
    
    // Test get() for type info
    val_RZiBCv = get("hello")
    pt(val_RZiBCv)
    

    sum_KLBUhJ = 0
    lp 6 {
        sum_KLBUhJ = sum_KLBUhJ + 1
        if (sum_KLBUhJ > 100) { done }
    }
    pt(sum_KLBUhJ)
    

    fn calc_sjkkBG(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_QralzA = calc_sjkkBG(69, 5+1)
    pt(res_QralzA)
    

    sum_tVMkVG = 0
    lp 10 {
        sum_tVMkVG = sum_tVMkVG + 1
        if (sum_tVMkVG > 100) { done }
    }
    pt(sum_tVMkVG)
    
