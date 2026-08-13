
    val_wDIusm = 4 % 3
    on (val_wDIusm) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_wDIusm)
    

    asn fn async_task_EonguL(x) {
        rev x * 2
    }
    uni async_task_EonguL(100)
    
    asn fn worker_kLkPWn(y) {
        pt("working on", y)
    }
    mul 2 worker_kLkPWn(80)
    
    awt
    

    try {
        throw_NTmdca = unknown
        throw_NTmdca()
    } err(e) {
        pt("caught error")
    }
    

    fn calc_YqNaYD(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_EMDsSG = calc_YqNaYD(68, 12+1)
    pt(res_EMDsSG)
    

    m_IFDtWo = math.sqrt(59)
    h_TiFvAk = hash.md5("test_FGOuzO")
    log.inf("hash:", h_TiFvAk)
    pt(m_IFDtWo)
    

    cl Base_hKwCkQ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_XlnFAs <- Base_hKwCkQ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_ziLGms = Child_XlnFAs(98, 59)
    pt(obj_ziLGms.get_id())
    

    val_UApvCE = 9 % 3
    on (val_UApvCE) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_UApvCE)
    

    use "dummy.y"
    
    e_oZbkCo = enc.b64("hello drylang")
    pt(e_oZbkCo)
    
    j_JdtjTe = json(`{"test": 123}`)
    pt(j_JdtjTe)
    
    // Test get() for type info
    val_gxzcIO = get("hello")
    pt(val_gxzcIO)
    

    fn calc_uJSbeb(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_ejLDLj = calc_uJSbeb(45, 39+1)
    pt(res_ejLDLj)
    

    use "dummy.y"
    
    e_houjsV = enc.b64("hello drylang")
    pt(e_houjsV)
    
    j_BmhIrx = json(`{"test": 123}`)
    pt(j_BmhIrx)
    
    // Test get() for type info
    val_nGxaLp = get("hello")
    pt(val_nGxaLp)
    

    fn calc_dvqTjX(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_JdTbar = calc_dvqTjX(43, 22+1)
    pt(res_JdTbar)
    

    try {
        throw_XCJdQw = unknown
        throw_XCJdQw()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_EQJtuq {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_aFeEYE <- Base_EQJtuq {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_HITWpt = Child_aFeEYE(45, 51)
    pt(obj_HITWpt.get_id())
    

    cl Base_tjXsqf {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_bsENbC <- Base_tjXsqf {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_bfETtY = Child_bsENbC(77, 96)
    pt(obj_bfETtY.get_id())
    

    sum_ZsPPHI = 0
    lp 11 {
        sum_ZsPPHI = sum_ZsPPHI + 1
        if (sum_ZsPPHI > 100) { done }
    }
    pt(sum_ZsPPHI)
    

    sum_kcQiNi = 0
    lp 15 {
        sum_kcQiNi = sum_kcQiNi + 1
        if (sum_kcQiNi > 100) { done }
    }
    pt(sum_kcQiNi)
    

    sum_fhvONJ = 0
    lp 7 {
        sum_fhvONJ = sum_fhvONJ + 1
        if (sum_fhvONJ > 100) { done }
    }
    pt(sum_fhvONJ)
    
