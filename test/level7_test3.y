
    val_XwyGZS = 17 % 3
    on (val_XwyGZS) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_XwyGZS)
    

    sum_hBDCUo = 0
    lp 5 {
        sum_hBDCUo = sum_hBDCUo + 1
        if (sum_hBDCUo > 100) { done }
    }
    pt(sum_hBDCUo)
    

    cl Base_bAvZCx {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_BzjCJG <- Base_bAvZCx {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_FFuJQo = Child_BzjCJG(55, 39)
    pt(obj_FFuJQo.get_id())
    

    fn calc_appZjb(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_IqIjMZ = calc_appZjb(25, 18+1)
    pt(res_IqIjMZ)
    

    fn calc_bjFvCg(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_bWvfoL = calc_bjFvCg(16, 71+1)
    pt(res_bWvfoL)
    

    arr_lWBePE = [60, 13+1, 43+2]
    map_GkAAGE = {"a": arr_lWBePE[0], "b": arr_lWBePE[1]}
    pt(map_GkAAGE["a"])
    

    m_JBNWJc = math.sqrt(29)
    h_uZyRbH = hash.md5("test_qDXpcp")
    log.inf("hash:", h_uZyRbH)
    pt(m_JBNWJc)
    

    cl Base_rjsGwm {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_BjRgYP <- Base_rjsGwm {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_hQXTWr = Child_BjRgYP(89, 61)
    pt(obj_hQXTWr.get_id())
    

    arr_KgSZGA = [62, 33+1, 14+2]
    map_hvsEZs = {"a": arr_KgSZGA[0], "b": arr_KgSZGA[1]}
    pt(map_hvsEZs["a"])
    

    fn calc_OIfdXO(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_MJOhRB = calc_OIfdXO(52, 2+1)
    pt(res_MJOhRB)
    

    asn fn async_task_gOZUYZ(x) {
        rev x * 2
    }
    uni async_task_gOZUYZ(80)
    
    asn fn worker_gYMIgp(y) {
        pt("working on", y)
    }
    mul 2 worker_gYMIgp(45)
    
    awt
    

    asn fn async_task_xwdoYM(x) {
        rev x * 2
    }
    uni async_task_xwdoYM(1)
    
    asn fn worker_XsyFPo(y) {
        pt("working on", y)
    }
    mul 2 worker_XsyFPo(28)
    
    awt
    

    asn fn async_task_RYWYdu(x) {
        rev x * 2
    }
    uni async_task_RYWYdu(86)
    
    asn fn worker_TqDNPD(y) {
        pt("working on", y)
    }
    mul 2 worker_TqDNPD(87)
    
    awt
    

    fn calc_JQpTMK(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_GzMbFw = calc_JQpTMK(58, 57+1)
    pt(res_GzMbFw)
    
