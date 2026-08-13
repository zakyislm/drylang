
    sum_lmmRkY = 0
    lp 9 {
        sum_lmmRkY = sum_lmmRkY + 1
        if (sum_lmmRkY > 100) { done }
    }
    pt(sum_lmmRkY)
    

    arr_HaUYxj = [28, 44+1, 26+2]
    map_jZhJfl = {"a": arr_HaUYxj[0], "b": arr_HaUYxj[1]}
    pt(map_jZhJfl["a"])
    

    fn calc_ouyZeX(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_OWuJRI = calc_ouyZeX(25, 66+1)
    pt(res_OWuJRI)
    

    try {
        throw_YHQjZu = unknown
        throw_YHQjZu()
    } err(e) {
        pt("caught error")
    }
    

    val_AcsuRt = 27 % 3
    on (val_AcsuRt) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_AcsuRt)
    

    m_NtFNRX = math.sqrt(70)
    h_aWJXsF = hash.md5("test_CbIkRT")
    log.inf("hash:", h_aWJXsF)
    pt(m_NtFNRX)
    

    use "dummy.y"
    
    e_rBDKjo = enc.b64("hello drylang")
    pt(e_rBDKjo)
    
    j_YBOdHQ = json(`{"test": 123}`)
    pt(j_YBOdHQ)
    
    // Test get() for type info
    val_HpCFuy = get("hello")
    pt(val_HpCFuy)
    

    arr_Umrkvl = [79, 31+1, 81+2]
    map_XFwunD = {"a": arr_Umrkvl[0], "b": arr_Umrkvl[1]}
    pt(map_XFwunD["a"])
    

    asn fn async_task_QsUVVl(x) {
        rev x * 2
    }
    uni async_task_QsUVVl(16)
    
    asn fn worker_RUZDhP(y) {
        pt("working on", y)
    }
    mul 2 worker_RUZDhP(26)
    
    awt
    

    m_HvpuMX = math.sqrt(19)
    h_pqEBWt = hash.md5("test_jwDCYC")
    log.inf("hash:", h_pqEBWt)
    pt(m_HvpuMX)
    

    cl Base_qEpGCk {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_EbKsEs <- Base_qEpGCk {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_jUXIHE = Child_EbKsEs(69, 47)
    pt(obj_jUXIHE.get_id())
    

    cl Base_acnlSN {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_HmbdLr <- Base_acnlSN {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_SlpjwB = Child_HmbdLr(36, 3)
    pt(obj_SlpjwB.get_id())
    

    cl Base_QERkcW {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_IMDFvR <- Base_QERkcW {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_SgRWPY = Child_IMDFvR(99, 78)
    pt(obj_SgRWPY.get_id())
    

    cl Base_rRpgJr {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_sgeEjZ <- Base_rRpgJr {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_TGfJOU = Child_sgeEjZ(13, 38)
    pt(obj_TGfJOU.get_id())
    
