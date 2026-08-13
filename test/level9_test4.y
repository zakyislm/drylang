
    fn calc_BFGiXs(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_OFShOp = calc_BFGiXs(62, 56+1)
    pt(res_OFShOp)
    

    sum_latsuV = 0
    lp 10 {
        sum_latsuV = sum_latsuV + 1
        if (sum_latsuV > 100) { done }
    }
    pt(sum_latsuV)
    

    m_XoamiA = math.sqrt(72)
    h_JBkNVY = hash.md5("test_qvbKoM")
    log.inf("hash:", h_JBkNVY)
    pt(m_XoamiA)
    

    fn calc_ivaiJv(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_fYVnSZ = calc_ivaiJv(61, 15+1)
    pt(res_fYVnSZ)
    

    val_VKeRdm = 25 % 3
    on (val_VKeRdm) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_VKeRdm)
    

    cl Base_LXQkXq {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_fuCpyu <- Base_LXQkXq {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_WwsYlJ = Child_fuCpyu(79, 90)
    pt(obj_WwsYlJ.get_id())
    

    val_KhUhpy = 25 % 3
    on (val_KhUhpy) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_KhUhpy)
    

    use "dummy.y"
    
    e_EcuxBq = enc.b64("hello drylang")
    pt(e_EcuxBq)
    
    j_yuOXQD = json(`{"test": 123}`)
    pt(j_yuOXQD)
    
    // Test get() for type info
    val_zQQMJJ = get("hello")
    pt(val_zQQMJJ)
    

    cl Base_mMmkdY {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_SDHMiH <- Base_mMmkdY {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_OUSMod = Child_SDHMiH(87, 51)
    pt(obj_OUSMod.get_id())
    

    val_VxYZqp = 83 % 3
    on (val_VxYZqp) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_VxYZqp)
    

    use "dummy.y"
    
    e_TajScQ = enc.b64("hello drylang")
    pt(e_TajScQ)
    
    j_lZkYRD = json(`{"test": 123}`)
    pt(j_lZkYRD)
    
    // Test get() for type info
    val_trfEBn = get("hello")
    pt(val_trfEBn)
    

    m_vxLIFP = math.sqrt(16)
    h_bAfKwS = hash.md5("test_EcfSqg")
    log.inf("hash:", h_bAfKwS)
    pt(m_vxLIFP)
    

    asn fn async_task_ptTMae(x) {
        rev x * 2
    }
    uni async_task_ptTMae(26)
    
    asn fn worker_aoSzPc(y) {
        pt("working on", y)
    }
    mul 2 worker_aoSzPc(34)
    
    awt
    

    try {
        throw_TINfag = unknown
        throw_TINfag()
    } err(e) {
        pt("caught error")
    }
    

    fn calc_SHqCik(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_BKQxJf = calc_SHqCik(88, 93+1)
    pt(res_BKQxJf)
    

    arr_UGgqYu = [61, 17+1, 22+2]
    map_USAfYi = {"a": arr_UGgqYu[0], "b": arr_UGgqYu[1]}
    pt(map_USAfYi["a"])
    
