
    try {
        throw_Heeeyg = unknown
        throw_Heeeyg()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_RIUQBI {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_HYBvrM <- Base_RIUQBI {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_NcMDcS = Child_HYBvrM(53, 30)
    pt(obj_NcMDcS.get_id())
    

    val_tzqcUE = 66 % 3
    on (val_tzqcUE) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_tzqcUE)
    

    use "dummy.y"
    
    e_onkXNm = enc.b64("hello drylang")
    pt(e_onkXNm)
    
    j_cSJxvm = json(`{"test": 123}`)
    pt(j_cSJxvm)
    
    // Test get() for type info
    val_UDbysZ = get("hello")
    pt(val_UDbysZ)
    

    fn calc_jpPFFJ(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_rpcNhk = calc_jpPFFJ(18, 59+1)
    pt(res_rpcNhk)
    

    arr_oXkuTB = [6, 38+1, 40+2]
    map_pinktf = {"a": arr_oXkuTB[0], "b": arr_oXkuTB[1]}
    pt(map_pinktf["a"])
    

    m_CgqcLt = math.sqrt(35)
    h_eCRHHQ = hash.md5("test_hGSbzR")
    log.inf("hash:", h_eCRHHQ)
    pt(m_CgqcLt)
    

    arr_McgrKx = [100, 58+1, 73+2]
    map_KFWozA = {"a": arr_McgrKx[0], "b": arr_McgrKx[1]}
    pt(map_KFWozA["a"])
    

    sum_KHhMBS = 0
    lp 8 {
        sum_KHhMBS = sum_KHhMBS + 1
        if (sum_KHhMBS > 100) { done }
    }
    pt(sum_KHhMBS)
    

    use "dummy.y"
    
    e_jCnIxc = enc.b64("hello drylang")
    pt(e_jCnIxc)
    
    j_uHXpTH = json(`{"test": 123}`)
    pt(j_uHXpTH)
    
    // Test get() for type info
    val_PzjaTH = get("hello")
    pt(val_PzjaTH)
    

    arr_aaNQkR = [21, 26+1, 91+2]
    map_zrjmKS = {"a": arr_aaNQkR[0], "b": arr_aaNQkR[1]}
    pt(map_zrjmKS["a"])
    

    m_mSowWC = math.sqrt(12)
    h_fpOtPF = hash.md5("test_xPyYyJ")
    log.inf("hash:", h_fpOtPF)
    pt(m_mSowWC)
    

    sum_RouciL = 0
    lp 7 {
        sum_RouciL = sum_RouciL + 1
        if (sum_RouciL > 100) { done }
    }
    pt(sum_RouciL)
    

    sum_uEJzrt = 0
    lp 10 {
        sum_uEJzrt = sum_uEJzrt + 1
        if (sum_uEJzrt > 100) { done }
    }
    pt(sum_uEJzrt)
    
