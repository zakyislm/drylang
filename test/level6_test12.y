
    sum_Zjpule = 0
    lp 7 {
        sum_Zjpule = sum_Zjpule + 1
        if (sum_Zjpule > 100) { done }
    }
    pt(sum_Zjpule)
    

    val_sEMyaK = 93 % 3
    on (val_sEMyaK) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_sEMyaK)
    

    val_bZTHUA = 78 % 3
    on (val_bZTHUA) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_bZTHUA)
    

    sum_fESeGI = 0
    lp 9 {
        sum_fESeGI = sum_fESeGI + 1
        if (sum_fESeGI > 100) { done }
    }
    pt(sum_fESeGI)
    

    m_UWaLKZ = math.sqrt(73)
    h_caNXpt = hash.md5("test_ZlWyCy")
    log.inf("hash:", h_caNXpt)
    pt(m_UWaLKZ)
    

    fn calc_FMTRwd(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_toZLtw = calc_FMTRwd(88, 93+1)
    pt(res_toZLtw)
    

    val_xKKBXS = 93 % 3
    on (val_xKKBXS) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_xKKBXS)
    

    cl Base_mwTXZt {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_ZCRiBA <- Base_mwTXZt {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_OgoMFm = Child_ZCRiBA(67, 48)
    pt(obj_OgoMFm.get_id())
    
