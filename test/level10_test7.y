
    cl Base_qQtwnb {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_vOoAQm <- Base_qQtwnb {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_KHRywV = Child_vOoAQm(5, 15)
    pt(obj_KHRywV.get_id())
    

    asn fn async_task_AhYOef(x) {
        rev x * 2
    }
    uni async_task_AhYOef(9)
    
    asn fn worker_vsfzLy(y) {
        pt("working on", y)
    }
    mul 2 worker_vsfzLy(19)
    
    awt
    

    arr_PXvsBs = [98, 89+1, 77+2]
    map_LLHEhd = {"a": arr_PXvsBs[0], "b": arr_PXvsBs[1]}
    pt(map_LLHEhd["a"])
    

    val_YcLkYY = 30 % 3
    on (val_YcLkYY) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_YcLkYY)
    

    fn calc_bWhxTi(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_IYJzvv = calc_bWhxTi(33, 20+1)
    pt(res_IYJzvv)
    

    m_SdOllF = math.sqrt(10)
    h_loJdDY = hash.md5("test_dstSnD")
    log.inf("hash:", h_loJdDY)
    pt(m_SdOllF)
    

    sum_cLMgak = 0
    lp 11 {
        sum_cLMgak = sum_cLMgak + 1
        if (sum_cLMgak > 100) { done }
    }
    pt(sum_cLMgak)
    

    val_mbjfBS = 47 % 3
    on (val_mbjfBS) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_mbjfBS)
    

    try {
        throw_KvUoAi = unknown
        throw_KvUoAi()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_iIsuTP = unknown
        throw_iIsuTP()
    } err(e) {
        pt("caught error")
    }
    

    m_FMVCph = math.sqrt(98)
    h_aSnlQv = hash.md5("test_qXOYDQ")
    log.inf("hash:", h_aSnlQv)
    pt(m_FMVCph)
    

    fn calc_HCwHqa(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_AtFaWy = calc_HCwHqa(71, 32+1)
    pt(res_AtFaWy)
    

    cns C_fAGHRj = 96
    v_IrGwqK = unknown
    v_IrGwqK = C_fAGHRj + 1
    if (v_IrGwqK > 0) {
        v_IrGwqK = v_IrGwqK * 2
    } el {
        v_IrGwqK = 0
    }
    pt(v_IrGwqK)
    

    try {
        throw_Djkjzi = unknown
        throw_Djkjzi()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_lpfxMv {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_pMAidA <- Base_lpfxMv {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_pNcvBb = Child_pMAidA(61, 70)
    pt(obj_pNcvBb.get_id())
    

    val_HbzIGO = 94 % 3
    on (val_HbzIGO) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_HbzIGO)
    

    sum_PxGQDj = 0
    lp 6 {
        sum_PxGQDj = sum_PxGQDj + 1
        if (sum_PxGQDj > 100) { done }
    }
    pt(sum_PxGQDj)
    
