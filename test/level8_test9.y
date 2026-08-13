
    val_lcLfDM = 74 % 3
    on (val_lcLfDM) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_lcLfDM)
    

    fn calc_VUvFVC(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_avfcIk = calc_VUvFVC(70, 69+1)
    pt(res_avfcIk)
    

    arr_PtHktv = [51, 81+1, 36+2]
    map_xOpkJb = {"a": arr_PtHktv[0], "b": arr_PtHktv[1]}
    pt(map_xOpkJb["a"])
    

    sum_YMYwKZ = 0
    lp 15 {
        sum_YMYwKZ = sum_YMYwKZ + 1
        if (sum_YMYwKZ > 100) { done }
    }
    pt(sum_YMYwKZ)
    

    sum_cRUeZf = 0
    lp 7 {
        sum_cRUeZf = sum_cRUeZf + 1
        if (sum_cRUeZf > 100) { done }
    }
    pt(sum_cRUeZf)
    

    asn fn async_task_DbcJqc(x) {
        rev x * 2
    }
    uni async_task_DbcJqc(31)
    
    asn fn worker_acTrOY(y) {
        pt("working on", y)
    }
    mul 2 worker_acTrOY(88)
    
    awt
    

    cl Base_hcdLUP {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_QFQpab <- Base_hcdLUP {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_fHxnLy = Child_QFQpab(70, 72)
    pt(obj_fHxnLy.get_id())
    

    cl Base_NEdnGf {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_sIwpSW <- Base_NEdnGf {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_RGsjxy = Child_sIwpSW(21, 84)
    pt(obj_RGsjxy.get_id())
    

    try {
        throw_minGlr = unknown
        throw_minGlr()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_vgYqnY = enc.b64("hello drylang")
    pt(e_vgYqnY)
    
    j_iaIKzV = json(`{"test": 123}`)
    pt(j_iaIKzV)
    
    // Test get() for type info
    val_PxORrL = get("hello")
    pt(val_PxORrL)
    

    use "dummy.y"
    
    e_sdOjok = enc.b64("hello drylang")
    pt(e_sdOjok)
    
    j_EoVSkF = json(`{"test": 123}`)
    pt(j_EoVSkF)
    
    // Test get() for type info
    val_OYFIgW = get("hello")
    pt(val_OYFIgW)
    

    fn calc_aRjzIO(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_CYCcTg = calc_aRjzIO(42, 60+1)
    pt(res_CYCcTg)
    

    val_vPcGOR = 28 % 3
    on (val_vPcGOR) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_vPcGOR)
    

    try {
        throw_eWjrxO = unknown
        throw_eWjrxO()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_waLJuy = enc.b64("hello drylang")
    pt(e_waLJuy)
    
    j_Itbbsj = json(`{"test": 123}`)
    pt(j_Itbbsj)
    
    // Test get() for type info
    val_qSOSMi = get("hello")
    pt(val_qSOSMi)
    
