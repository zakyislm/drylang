
    cl Base_xvOwWI {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_fWQTDJ <- Base_xvOwWI {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_ewOWYl = Child_fWQTDJ(44, 33)
    pt(obj_ewOWYl.get_id())
    

    use "dummy.y"
    
    e_YavzPa = enc.b64("hello drylang")
    pt(e_YavzPa)
    
    j_WJewPJ = json(`{"test": 123}`)
    pt(j_WJewPJ)
    
    // Test get() for type info
    val_BDjPjB = get("hello")
    pt(val_BDjPjB)
    

    cl Base_JVswKi {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_LxYUjq <- Base_JVswKi {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_zlmHxZ = Child_LxYUjq(76, 93)
    pt(obj_zlmHxZ.get_id())
    

    cns C_EgtnoU = 62
    v_WsdlRS = unknown
    v_WsdlRS = C_EgtnoU + 2
    if (v_WsdlRS > 0) {
        v_WsdlRS = v_WsdlRS * 2
    } el {
        v_WsdlRS = 0
    }
    pt(v_WsdlRS)
    

    fn calc_Zmzarl(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_SrzYpn = calc_Zmzarl(98, 16+1)
    pt(res_SrzYpn)
    

    val_jMiKJf = 26 % 3
    on (val_jMiKJf) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_jMiKJf)
    

    cns C_tuqOlI = 2
    v_ijcRDL = unknown
    v_ijcRDL = C_tuqOlI + 40
    if (v_ijcRDL > 0) {
        v_ijcRDL = v_ijcRDL * 2
    } el {
        v_ijcRDL = 0
    }
    pt(v_ijcRDL)
    

    fn calc_UWQhGU(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_BHAgri = calc_UWQhGU(66, 12+1)
    pt(res_BHAgri)
    

    cl Base_uwvZrI {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_xbLbVK <- Base_uwvZrI {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_EmfKTe = Child_xbLbVK(69, 51)
    pt(obj_EmfKTe.get_id())
    

    val_hivKGi = 45 % 3
    on (val_hivKGi) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_hivKGi)
    

    sum_TwYAKI = 0
    lp 8 {
        sum_TwYAKI = sum_TwYAKI + 1
        if (sum_TwYAKI > 100) { done }
    }
    pt(sum_TwYAKI)
    

    fn calc_OvKkAf(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_OoWCou = calc_OvKkAf(17, 73+1)
    pt(res_OoWCou)
    

    fn calc_dXvUsR(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_neiGKW = calc_dXvUsR(41, 56+1)
    pt(res_neiGKW)
    

    m_yEARuq = math.sqrt(7)
    h_RBHaon = hash.md5("test_upbbZh")
    log.inf("hash:", h_RBHaon)
    pt(m_yEARuq)
    

    cl Base_VjlBwP {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_fDuRcY <- Base_VjlBwP {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_NoUYeO = Child_fDuRcY(66, 68)
    pt(obj_NoUYeO.get_id())
    

    val_hVGaHT = 100 % 3
    on (val_hVGaHT) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_hVGaHT)
    
